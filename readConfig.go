package main

import (
	"log"
	"strconv"
)

func makeConfig() map[string]string {
	config := make(map[string]string, 0)

	getConfig("configs/config_template", &config) // use defaults
	getConfig("configs/config", &config)          // use user-defined

	return config
}

// Read gamepadID parameter. Should be an integer.
func setGamepadID(config map[string]string) {
	id, err := strconv.ParseInt(config["gamepadID"], 0, 64)
	if err != nil {
		log.Fatalf("gamepadID parameter (value read: %s) was not set correctly in config file. Edit it with an integer value.", config["gamepadID"])
	} else {
		gamepadID = int(id)
	}
}

// Read deadZone parameter. Should be a number between 0 and 1.
func setDeadZone(config map[string]string) {
	f, err := strconv.ParseFloat(config["deadZone"], 64)
	if err != nil || f < 0 || f > 1 {
		log.Fatalf("deadZone parameter (value read: %s) was not set correctly in config file. Edit it with a positive number between 0 and 1.", config["deadZone"])
	} else {
		deadZone = f
	}
}
