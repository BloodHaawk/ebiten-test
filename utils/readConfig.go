package utils

import (
	"log"
	"strconv"
)

// MakeConfig returns a mapped config from a config file
func MakeConfig() map[string]string {
	config := make(map[string]string, 0)

	GetConfig("configs/config_template", &config) // use defaults
	GetConfig("configs/config", &config)          // use user-defined

	return config
}

// ReadGamepadID read a gamepad ID parameter from a config file. Should be an integer.
func ReadGamepadID(config map[string]string) int {
	id, err := strconv.ParseInt(config["gamepadID"], 0, 64)
	if err != nil {
		log.Fatalf("gamepadID parameter (value read: %s) was not set correctly in config file. Edit it with an integer value.", config["gamepadID"])
	}
	return int(id)

}

// ReadDeadZone reads a gamepad's dead zone parameter froma config file. Should be a number between 0 and 1.
func ReadDeadZone(config map[string]string) float64 {
	f, err := strconv.ParseFloat(config["deadZone"], 64)
	if err != nil || f < 0 || f > 1 {
		log.Fatalf("deadZone parameter (value read: %s) was not set correctly in config file. Edit it with a positive number between 0 and 1.", config["deadZone"])
	}
	return f

}
