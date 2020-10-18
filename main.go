package main

import (
	"flag"
	"log"
	"os"
	"time"

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

	// Parse the CLI flags
	toggle := flag.Bool("toggle", false, "Toggles the lights instead of flipping on and off")
	flag.Parse()

	// Set the Tuya API's config
	config.SetEnv(common.URLUS, accessID, accessSecret)

	// Get the Device from API
	log.Printf("Getting Device - %v", deviceID)
	getDeviceResponse, err := device.GetDevice(deviceID)
	if err != nil {
		log.Fatalf("Fatal during GetDevice %v", err.Error())
	}

	log.Printf("Retrived Device.Name=%v", getDeviceResponse.Result.Name)

	if *toggle {

		log.Printf("Toggling the lights 💡")

		// Get the current power state of the device
		currentPowerState := GetDevicePowerState(*getDeviceResponse)

		// Set the device power state to the inverse
		SetDevicePower(getDeviceResponse.Result.ID, !currentPowerState)

	} else {

		log.Printf("You earned your sunshine ☀️")

		// Turn on the lights
		SetDevicePower(getDeviceResponse.Result.ID, true)

		log.Print("Sleeping for 10 seconds")
		time.Sleep(10 * time.Second)

		// Turn off the lights
		SetDevicePower(getDeviceResponse.Result.ID, false)

	}

}

// GetDevicePowerState is a helper function to determine a device's power state based on the GetDeviceResponse
func GetDevicePowerState(getDeviceResponse device.GetDeviceResponse) bool {

	devicePowerState := false

	// Iterate through all the Status
	for _, v := range getDeviceResponse.Result.Status {

		// Determine the device's power state
		if v.Code == "switch_led" {

			// Need to cast the interface to bool
			devicePowerState = v.Value.(bool)

			// No need to keep iterating
			break
		}
	}

	return devicePowerState

}

// SetDevicePower sets the power state of the passed device ID with the passed power state
func SetDevicePower(deviceID string, targetPowerState bool) {

	// Construct the intended device commands
	var switchDeviceCommand device.Command
	switchDeviceCommand.Code = "switch_led"
	switchDeviceCommand.Value = targetPowerState

	var deviceCommands []device.Command
	deviceCommands = append(deviceCommands, switchDeviceCommand)

	// Let it rip!
	response, err := device.PostDeviceCommand(deviceID, deviceCommands)
	if err != nil {
		log.Fatalf("Fatal during PostDeviceCommand %v", err.Error())
	}

	log.Printf("Completed SetDevicePower! targetPowerState=%v Response.Code=%v Response.Msg=%v",
		targetPowerState,
		response.Code,
		response.Msg,
	)

}
