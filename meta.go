package main

import (
	"encoding/json"
	"os"
)

type MetaStore struct {
	Ids      []string
	Filename string
}

func LoadMetaStore(filename string) *MetaStore {
	// Open input
	input, err := os.Open(filename)
	if err == nil {
		// Decode and return
		dec := json.NewDecoder(input)
		var store MetaStore
		err = dec.Decode(&store)
		if err == nil {
			store.Filename = filename
			return &store
		}
	}
	return &MetaStore{Ids: []string{}, Filename: filename}
}

func (ms *MetaStore) Write() error {
	// Create output
	output, err := os.Create(ms.Filename)
	if err != nil {
		return err
	}

	// Encode and return
	enc := json.NewEncoder(output)
	err = enc.Encode(ms)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MetaStore) AddAndWriteId(id string) {
	ms.AddId(id)
	ms.Write()
}

func (ms *MetaStore) AddId(id string) {
	ms.Ids = append(ms.Ids, id)
}
