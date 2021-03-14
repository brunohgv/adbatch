package adb

import (
	"github.com/brunohgv/adbatch/utils"
)

func GetDevices() []string {
	output, err := utils.ExecuteCommand("adb devices")
	if err != nil {
		panic(err)
	}
	return getDeviceIdsFromDevicesMessage(output)
}