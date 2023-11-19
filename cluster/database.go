package cluster

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"github.com/neutrino2211/go-option"
)

const MAX_RECORD_LEN = 0xffffffff

func commitToDiskJob(database *Database) func() {
	done := false
	go func() {
		for !done {
			time.Sleep(10 * time.Second)
			database.Save()
		}
	}()

	return func() {
		done = true
	}
}

func isDBDormant(database *Database, maxIdleTime time.Duration) bool {
	for _, r := range database.records {
		if r.LastPing.Add(maxIdleTime).After(time.Now()) {
			return false
		}
	}

	return true
}

func hibernateDatabaseJob(database *Database, inspectionInterval, maxIdleTime time.Duration) {

	go func() {
		for {
			time.Sleep(inspectionInterval)

			if isDBDormant(database, maxIdleTime) {
				database.Hibernate()
				return
			}
		}
	}()
}

// func resizeBytes(data []byte, size int) []

func uintToByte(integer uint32) []byte {
	r := []byte{}

	r = binary.LittleEndian.AppendUint32(r, integer)

	return r
}

func uint64ToByte(integer uint64) []byte {
	r := []byte{}

	r = binary.LittleEndian.AppendUint64(r, integer)

	return r
}

func int64ToByte(integer int64) []byte {
	r := []byte{}

	r = binary.LittleEndian.AppendUint64(r, uint64(integer))

	return r
}

func byteToUint64(data []byte) uint64 {
	return binary.LittleEndian.Uint64(data)
}

func removeRecord(slice []*Record, s int) []*Record {
	slice[s] = &Record{}
	return slice
}

type Record struct {
	Data       [2048]byte
	Owner      [64]byte
	Permission byte
	LastPing   time.Time
}

func (r *Record) Hash() ([64]byte, error) {
	sha := sha512.New()
	d := r.Data[:]
	d = append(d, r.Owner[:]...)
	_, err := sha.Write(d)

	if err != nil {
		return [64]byte{}, err
	}

	sum := sha.Sum(nil)
	key := [64]byte{}

	copy(key[:], sum[:64])

	return key, nil
}

func (r *Record) ValidateOwnership(token [64]byte) bool {
	return token == r.Owner
}

type Database struct {
	records      []*Record
	deletedList  []uint32
	lookupTable  map[[64]byte]uint32
	isHibernated bool
	name         string
	cancelFunc   func()
}

func (d *Database) coldLookup(hash [64]byte) (bool, error) {
	data, err := os.ReadFile(d.name)

	if err != nil {
		return false, err
	}

	lookupLen := byteToUint64(data[4:12])

	lookupTableBytes := data[28 : 28+lookupLen]

	err = gob.NewDecoder(bytes.NewBuffer(lookupTableBytes)).Decode(&d.lookupTable)

	if err != nil {
		return false, err
	}

	_, ok := d.lookupTable[hash]

	return ok, nil
}

func (d *Database) Size() uint64 {
	return uint64(len(d.records))
}

func (d *Database) Hibernate() {
	d.isHibernated = true
	d.cancelFunc()
	d.CommitToDisk(d.name)

	d.lookupTable = map[[64]byte]uint32{}
	d.records = make([]*Record, 0)
	println(d.name + " hibernated")
}

func (d *Database) WakeUp() {
	d.Init()

	d.isHibernated = false
	println(d.name + " woken up")
}

func (d *Database) Init() {
	d.LoadFromDisk(d.name)
	hibernateDatabaseJob(d, 30*time.Second, 10*time.Second)
	d.cancelFunc = commitToDiskJob(d)
}

func (d *Database) Save() {
	d.CommitToDisk(d.name)
}

func (d *Database) AddRecord(r *Record) [64]byte {
	key, err := r.Hash()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error adding record, continuing...")
		return [64]byte{}
	}

	if _, ok := d.lookupTable[key]; ok {
		fmt.Fprintln(os.Stderr, "Record already existts, continuing...")
		return key
	}

	d.lookupTable[key] = uint32(len(d.records))
	d.records = append(d.records, r)

	return key
}

func (d *Database) GetRecord(hash [64]byte) *Record {
	if d.isHibernated {
		d.WakeUp()
	}

	if index, ok := d.lookupTable[hash]; ok {
		fmt.Println(d.records)
		return d.records[index]
	}

	return nil
}

func (d *Database) UpdateRecord(hash [64]byte, record *Record) {

	if index, ok := d.lookupTable[hash]; ok {
		d.records[index] = record
		delete(d.lookupTable, hash)
		d.lookupTable[hash] = index
	}

}

func (d *Database) DeleteRecord(hash [64]byte) {
	if index, ok := d.lookupTable[hash]; ok {
		delete(d.lookupTable, hash)
		d.records = removeRecord(d.records, int(index))
	}
}

func (d *Database) RecordExists(hash [64]byte, shouldWakeUp bool) bool {
	if d.isHibernated {
		ok, err := d.coldLookup(hash)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error performing cold lookup", err)
		}

		if ok && shouldWakeUp {
			d.WakeUp()
		}
		return ok
	}

	_, ok := d.lookupTable[hash]

	return ok
}

func (d *Database) CommitToDisk(filename string) error {
	err := os.WriteFile(filename, d.Serialize(), 0644)

	return err
}

func (d *Database) LoadFromDisk(filename string) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	d.Deserialize(data)
	return nil
}

func (d *Database) Serialize() []byte {
	var buf = []byte{'x', 'H', 'D', 'B'} // Magic bytes
	var mapBytesBuffer = new(bytes.Buffer)
	var recordBytesBuffer = new(bytes.Buffer)
	var deletedListBytesBuffer = new(bytes.Buffer)

	mapEncoder := gob.NewEncoder(mapBytesBuffer)
	recordsEncoder := gob.NewEncoder(recordBytesBuffer)
	deletedListEncoder := gob.NewEncoder(deletedListBytesBuffer)

	option.Some(mapEncoder.Encode(d.lookupTable)).ExpectNil("Failed to serialize database: [lookupTable]")
	option.Some(recordsEncoder.Encode(d.records)).ExpectNil("Failed to serialize database: [records]")
	option.Some(deletedListEncoder.Encode(d.deletedList)).ExpectNil("Failed to serialize database: [deletedList]")

	lookupSize := mapBytesBuffer.Len()              // Size of lookup table
	recordsSize := recordBytesBuffer.Len()          // Size of records
	deletedListSize := deletedListBytesBuffer.Len() // Size of deleted list

	buf = append(buf, uint64ToByte(uint64(lookupSize))...)  // Encode lookup table size
	buf = append(buf, uint64ToByte(uint64(recordsSize))...) // Encode the record size
	buf = append(buf, uint64ToByte(uint64(deletedListSize))...)

	buf = append(buf, mapBytesBuffer.Bytes()...)    // Add encoded lookup table
	buf = append(buf, recordBytesBuffer.Bytes()...) // Add encoded records
	buf = append(buf, deletedListBytesBuffer.Bytes()...)

	return buf
}

func (d *Database) Deserialize(data []byte) {
	// 0..3 is magic bytes, 4..11 is lookup len in bytes, 12..20 is records table in bytes
	lookupLen := byteToUint64(data[4:12])
	recordsLen := byteToUint64(data[12:20])

	lookupTableBytes := data[28 : 28+lookupLen]
	recordsBytes := data[28+lookupLen : 28+lookupLen+recordsLen]
	deletedListBytes := data[28+lookupLen+recordsLen:]

	lookupTableDecoder := gob.NewDecoder(bytes.NewBuffer(lookupTableBytes))
	recordDecoder := gob.NewDecoder(bytes.NewBuffer(recordsBytes))
	deletedListDecoder := gob.NewDecoder(bytes.NewBuffer(deletedListBytes))

	baseRecords := []*Record{}
	baseLookupTable := make(map[[64]byte]uint32)
	baseDeletedList := []uint32{}

	option.Some(lookupTableDecoder.Decode(&baseLookupTable)).ExpectNil("Failed to deserialize database: [lookupTable]")
	option.Some(recordDecoder.Decode(&baseRecords)).ExpectNil("Failed to deserialize database: [records]")
	option.Some(deletedListDecoder.Decode(&baseDeletedList)).ExpectNil("Failed to deserialize database: [deletedList]")

	d.records = baseRecords
	d.lookupTable = baseLookupTable
	d.deletedList = baseDeletedList
}
