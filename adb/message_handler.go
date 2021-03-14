package adb

import (
	"strings"
)

func getDeviceIdsFromDevicesMessage(message string) []string {
	lines := strings.Split(message, "\n")
	deviceIds := []string{}
	for i := 1 ; i < len(lines) - 2 ; i++ {
		id := strings.Split(lines[i], "\t")[0]
		deviceIds = append(deviceIds, id)
	}
	return deviceIds
}