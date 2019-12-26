package main

import "strconv"

import "log"

func makeConfig() map[string]string {
	config := make(map[string]string, 0)

	getConfig("configs/config_template", &config) // use defaults
	getConfig("configs/config", &config)          // use user-defined

	return config
}

// Read deadZone parameter. Should always be between 0 and 1
func setDeadZone(config map[string]string) {
	f, err := strconv.ParseFloat(config["deadZone"], 64)
	if err != nil || f < 0 || f > 1 {
		log.Fatalf("deadZone parameter (value read: %s) was not set correctly in config file. Edit it with a positive number between 0 and 1.", config["deadZone"])
	} else {
		deadZone = f
	}
}
