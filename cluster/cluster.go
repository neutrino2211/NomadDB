package cluster

import (
	"fmt"
	"math"
	"os"
)

type ClusterFlags struct {
	IsPrimeCluster bool
}

type Cluster struct {
	Databases []*Database
	Name      string
	Peers     []string
	Flags     ClusterFlags
}

func (c *Cluster) HasRecord(hash [64]byte) bool {
	for _, d := range c.Databases {
		if d.RecordExists(hash, false) {
			return true
		}
	}

	return false
}

func (c *Cluster) GetRecord(hash [64]byte) *Record {
	for _, d := range c.Databases {
		if d.RecordExists(hash, true) {
			return d.GetRecord(hash)
		}
	}

	return nil
}

func (c *Cluster) UpdateRecord(hash [64]byte, record *Record) {
	for _, d := range c.Databases {
		if d.RecordExists(hash, true) {
			d.UpdateRecord(hash, record)
			return
		}
	}
}

func (c *Cluster) AddRecord(record *Record) (int, [64]byte) {
	var minSize = math.Inf(1)
	var minDb *Database
	var minIndex = 0

	if h, _ := record.Hash(); c.HasRecord(h) {
		for i, d := range c.Databases {
			if d.RecordExists(h, false) {
				fmt.Fprintf(os.Stderr, "Record already exists in cluster\n")
				return i, h
			}
		}
	}

	for i, d := range c.Databases {

		if d.Size() < uint64(minSize) {
			minSize = float64(d.Size())
			minDb = d
			minIndex = i
		}
	}

	if minDb.isHibernated {
		minDb.WakeUp()
	}

	return minIndex, minDb.AddRecord(record)
}

func (c *Cluster) DeleteRecord(hash [64]byte) {
	for _, d := range c.Databases {
		if d.RecordExists(hash, true) {
			d.DeleteRecord(hash)
			return
		}
	}
}

func NewCluster(name string, nodes uint, prime bool) *Cluster {
	c := &Cluster{
		Databases: []*Database{},
		Name:      name,
		Peers:     []string{},
		Flags: ClusterFlags{
			IsPrimeCluster: false,
		},
	}

	for i := 0; i < int(nodes); i++ {
		dbName := fmt.Sprintf("%s_%d", name, i)

		d := &Database{
			records:     []*Record{},
			lookupTable: make(map[[64]byte]uint32),
			name:        dbName,
		}
		d.Init()
		c.Databases = append(c.Databases, d)
	}

	return c
}
