package main

import (
	"fmt"
	"strings"
	"os"
	"sync"
	"github.com/brunohgv/adbatch/adb"
	"github.com/brunohgv/adbatch/runner"
)

func main() {
    arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Adbatch is working!")
		fmt.Println("Run an adb command using adbatch to broadcast it to all connected devices.")
	} else {
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
}