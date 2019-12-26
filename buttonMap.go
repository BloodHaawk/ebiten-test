package main

import (
	"github.com/hajimehoshi/ebiten"
)

func makeButtonConfig() map[string]string {
	config := make(map[string]string, 0)

	getConfig("configs/button_config_template", &config) // use defaults
	getConfig("configs/button_config", &config)          // use user-defined

	return config
}

// Gamepad mapping
var buttonMap = map[string]ebiten.GamepadButton{
	"Button0":  ebiten.GamepadButton0,
	"Button1":  ebiten.GamepadButton1,
	"Button2":  ebiten.GamepadButton2,
	"Button3":  ebiten.GamepadButton3,
	"Button4":  ebiten.GamepadButton4,
	"Button5":  ebiten.GamepadButton5,
	"Button6":  ebiten.GamepadButton6,
	"Button7":  ebiten.GamepadButton7,
	"Button8":  ebiten.GamepadButton8,
	"Button9":  ebiten.GamepadButton9,
	"Button10": ebiten.GamepadButton10,
	"Button11": ebiten.GamepadButton11,
	"Button12": ebiten.GamepadButton12,
	"Button13": ebiten.GamepadButton13,
	"Button14": ebiten.GamepadButton14,
	"Button15": ebiten.GamepadButton15,
	"Button16": ebiten.GamepadButton16,
	"Button17": ebiten.GamepadButton17,
	"Button18": ebiten.GamepadButton18,
	"Button19": ebiten.GamepadButton19,
	"Button20": ebiten.GamepadButton20,
	"Button21": ebiten.GamepadButton21,
	"Button22": ebiten.GamepadButton22,
	"Button23": ebiten.GamepadButton23,
	"Button24": ebiten.GamepadButton24,
	"Button25": ebiten.GamepadButton25,
	"Button26": ebiten.GamepadButton26,
	"Button27": ebiten.GamepadButton27,
	"Button28": ebiten.GamepadButton28,
	"Button29": ebiten.GamepadButton29,
	"Button30": ebiten.GamepadButton30,
	"Button31": ebiten.GamepadButton31,
}
