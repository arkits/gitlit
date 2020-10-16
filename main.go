package main

import (
	"log"
	"os"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
	"github.com/TuyaInc/tuya_cloud_sdk_go/api/device"
	"github.com/TuyaInc/tuya_cloud_sdk_go/config"
)

func main() {

	log.Printf("Namaskar Mandali! 🙏")

	// Read private vars from OS env-vars
	deviceID := os.Getenv("P4L_DEVICE_ID")
	accessID := os.Getenv("P4L_ACCESS_ID")
	accessSecret := os.Getenv("P4L_ACCESS_SECRET")

	// Set the Tuya API's config
	config.SetEnv(common.URLUS, accessID, accessSecret)

	// Get the Device from API
	log.Printf("Getting Device - %v", deviceID)
	got, err := device.GetDevice(deviceID)
	if err != nil {
		log.Fatalf("Fatal during GetDevice %v", err.Error())
	}
	log.Printf("Retrived Device.Name=%v", got.Result.Name)

	var targetPowerState bool

	// Iterate through all the Status
	for _, v := range got.Result.Status {

		// Determine the device's power state
		if v.Code == "switch_led" {

			// Need to cast the interface to bool
			currentPowerState := v.Value.(bool)

			// We will invert the power state
			targetPowerState = !currentPowerState

			log.Printf("currentPowerState=%v targetPowerState%v", currentPowerState, targetPowerState)

			// No need to keep iterating
			break
		}
	}

	// Construct the intended device commands
	var switchDeviceCommand device.Command
	switchDeviceCommand.Code = "switch_led"
	switchDeviceCommand.Value = targetPowerState

	var deviceCommands []device.Command
	deviceCommands = append(deviceCommands, switchDeviceCommand)

	// Let it rip!
	response, err := device.PostDeviceCommand(deviceID, deviceCommands)
	log.Printf("Response Code=%v Msg=%v", response.Code, response.Msg)

}
