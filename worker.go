package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	BINARY = "bussim"
)

func RunSimulation(input, id string) error {
	log.Println("Running sim w/ id ", id)
	baseFile := path.Join(DATADIR, id)
	err := os.Mkdir(baseFile, 0755)
	if err != nil {
		log.Println(err)
	}
	inputFile := path.Join(baseFile, INPUTFNAME)
	outputFile := path.Join(baseFile, OUTPUTFNAME)
	metaFile := path.Join(baseFile, METADATAFNAME)

	// Write file
	err = ioutil.WriteFile(inputFile, []byte(input), 0600)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create and run command
	// Exit status is thrown away because the command output is more useful
	cmd := exec.Command(BINARY, "-input", inputFile, "-output", outputFile)
	cmdOutput, _ := cmd.CombinedOutput()
	err = ioutil.WriteFile(metaFile, []byte(cmdOutput), 0600)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
