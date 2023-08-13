package main

import (
	"os"

	"github.com/neutrino2211/commander"
)

type NewRecordRequest struct {
	Block [2048]byte `json:"record"`
}

type PingRecordRequest struct {
	Hash [20]byte `json:"hash"`
}

func main() {
	program := commander.Commander{}
	program.Init("cluster")
	program.Register("database", &DBCommand{})
	program.Register("registry", &RegistryCommand{})
	program.Register("init", &InitCommand{})
	program.Parse(os.Args)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	data, err := io.ReadAll(r.Body)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	req := &NewRecordRequest{}

	// 	err = json.Unmarshal(data, req)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	_, hash := localCluster.AddRecord(&Record{
	// 		Data:     req.Block,
	// 		LastPing: time.Now(),
	// 	})

	// 	w.Write(hash[:])
	// })

	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	data, err := io.ReadAll(r.Body)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	req := &PingRecordRequest{}

	// 	err = json.Unmarshal(data, req)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	record := localCluster.GetRecord(req.Hash)

	// 	if record != nil {
	// 		record.LastPing = time.Now()
	// 		localCluster.UpdateRecord(req.Hash, record)
	// 		w.Write(int64ToByte(record.LastPing.Unix()))
	// 		return
	// 	}

	// 	w.Write([]byte{0})
	// })

	// http.ListenAndServe(":4040", nil)

	// fmt.Printf("0x%x\n", binary.LittleEndian.AppendUint32())
}
