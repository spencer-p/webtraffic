package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

const (
	INPUTFNAME    = "input.txt"
	OUTPUTFNAME   = "output.txt"
	METADATAFNAME = "meta.txt"
)

type PageData struct {
	Input, Output, Meta []byte
}

func LoadPageData(dir, id string) *PageData {
	return &PageData{
		Input:  LoadData(dir, id, INPUTFNAME),
		Output: LoadData(dir, id, OUTPUTFNAME),
		Meta:   LoadData(dir, id, METADATAFNAME),
	}
}

func LoadData(dir, id, base string) []byte {
	filename := path.Join(dir, id, base)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		body = []byte(fmt.Sprintf("Failed to load %s for '%s'", base, id))
	}
	return body
}
