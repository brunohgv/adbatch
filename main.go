package main

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
	"github.com/brunohgv/adbatch/adb"
	"github.com/brunohgv/adbatch/utils"
	"github.com/brunohgv/adbatch/files"
)

func main() {
    arguments := os.Args[1:]

	command := strings.Join(arguments, " ")

	deviceIds := adb.GetDevices()
	currentDir, err := os.Getwd()
	if err != nil {
		panic("Could not locate the current directory")
	}

	for _, deviceId := range deviceIds {
		out, _ := utils.ExecuteCommand(fmt.Sprintf("adb -s %s %s", deviceId, command))
		fmt.Printf("Output for %v:\n\n", deviceId)
		fmt.Println(out)
		filePath := filepath.Join(currentDir, fmt.Sprintf("%v-out.txt", deviceId))
		files.WriteFile(filePath, out)
		fmt.Println("================================================")
	}
}