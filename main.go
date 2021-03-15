package main

import (
	"strings"
	"os"
	"sync"
	"github.com/brunohgv/adbatch/adb"
	"github.com/brunohgv/adbatch/runner"
)

func main() {
    arguments := os.Args[1:]

	command := strings.Join(arguments, " ")

	deviceIds := adb.GetDevices()
	outputDir, err := os.Getwd()
	if err != nil {
		panic("Could not locate the current directory")
	}

	var wg sync.WaitGroup
	for _, deviceId := range deviceIds {
		wg.Add(1)

		go func(deviceId string) {
			defer wg.Done()
			process := runner.Process{
				Command:	command,
				DeviceId:	deviceId,
				FilePath:	outputDir,
			}
			process.Run()
		}(deviceId)
	}

	wg.Wait()
}