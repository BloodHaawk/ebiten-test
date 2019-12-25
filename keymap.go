package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

func getConfig(configFile string, config *map[string]string) {
	file, err := os.Open(configFile)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var kv = strings.Split(scanner.Text(), " ")
		if len(kv) == 2 {
			(*config)[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	logError(scanner.Err())
	file.Close()

	return
}

func makeKeyConfig() map[string]string {
	config := make(map[string]string, 0)

	getConfig("config_template", &config) // use defaults
	getConfig("config", &config)          // use user-defined

	return config

}

var keyMap = map[string]ebiten.Key{
	"0":            ebiten.Key0,
	"1":            ebiten.Key1,
	"2":            ebiten.Key2,
	"3":            ebiten.Key3,
	"4":            ebiten.Key4,
	"5":            ebiten.Key5,
	"6":            ebiten.Key6,
	"7":            ebiten.Key7,
	"8":            ebiten.Key8,
	"9":            ebiten.Key9,
	"A":            ebiten.KeyA,
	"B":            ebiten.KeyB,
	"C":            ebiten.KeyC,
	"D":            ebiten.KeyD,
	"E":            ebiten.KeyE,
	"F":            ebiten.KeyF,
	"G":            ebiten.KeyG,
	"H":            ebiten.KeyH,
	"I":            ebiten.KeyI,
	"J":            ebiten.KeyJ,
	"K":            ebiten.KeyK,
	"L":            ebiten.KeyL,
	"M":            ebiten.KeyM,
	"N":            ebiten.KeyN,
	"O":            ebiten.KeyO,
	"P":            ebiten.KeyP,
	"Q":            ebiten.KeyQ,
	"R":            ebiten.KeyR,
	"S":            ebiten.KeyS,
	"T":            ebiten.KeyT,
	"U":            ebiten.KeyU,
	"V":            ebiten.KeyV,
	"W":            ebiten.KeyW,
	"X":            ebiten.KeyX,
	"Y":            ebiten.KeyY,
	"Z":            ebiten.KeyZ,
	"Apostrophe":   ebiten.KeyApostrophe,
	"Backslash":    ebiten.KeyBackslash,
	"Backspace":    ebiten.KeyBackspace,
	"CapsLock":     ebiten.KeyCapsLock,
	"Comma":        ebiten.KeyComma,
	"Delete":       ebiten.KeyDelete,
	"Down":         ebiten.KeyDown,
	"End":          ebiten.KeyEnd,
	"Enter":        ebiten.KeyEnter,
	"Equal":        ebiten.KeyEqual,
	"Escape":       ebiten.KeyEscape,
	"F1":           ebiten.KeyF1,
	"F2":           ebiten.KeyF2,
	"F3":           ebiten.KeyF3,
	"F4":           ebiten.KeyF4,
	"F5":           ebiten.KeyF5,
	"F6":           ebiten.KeyF6,
	"F7":           ebiten.KeyF7,
	"F8":           ebiten.KeyF8,
	"F9":           ebiten.KeyF9,
	"F10":          ebiten.KeyF10,
	"F11":          ebiten.KeyF11,
	"F12":          ebiten.KeyF12,
	"GraveAccent":  ebiten.KeyGraveAccent,
	"Home":         ebiten.KeyHome,
	"Insert":       ebiten.KeyInsert,
	"KP0":          ebiten.KeyKP0,
	"KP1":          ebiten.KeyKP1,
	"KP2":          ebiten.KeyKP2,
	"KP3":          ebiten.KeyKP3,
	"KP4":          ebiten.KeyKP4,
	"KP5":          ebiten.KeyKP5,
	"KP6":          ebiten.KeyKP6,
	"KP7":          ebiten.KeyKP7,
	"KP8":          ebiten.KeyKP8,
	"KP9":          ebiten.KeyKP9,
	"KPAdd":        ebiten.KeyKPAdd,
	"KPDecimal":    ebiten.KeyKPDecimal,
	"KPDivide":     ebiten.KeyKPDivide,
	"KPEnter":      ebiten.KeyKPEnter,
	"KPEqual":      ebiten.KeyKPEqual,
	"KPMultiply":   ebiten.KeyKPMultiply,
	"KPSubtract":   ebiten.KeyKPSubtract,
	"Left":         ebiten.KeyLeft,
	"LeftBracket":  ebiten.KeyLeftBracket,
	"Menu":         ebiten.KeyMenu,
	"Minus":        ebiten.KeyMinus,
	"NumLock":      ebiten.KeyNumLock,
	"PageDown":     ebiten.KeyPageDown,
	"PageUp":       ebiten.KeyPageUp,
	"Pause":        ebiten.KeyPause,
	"Period":       ebiten.KeyPeriod,
	"PrintScreen":  ebiten.KeyPrintScreen,
	"Right":        ebiten.KeyRight,
	"RightBracket": ebiten.KeyRightBracket,
	"ScrollLock":   ebiten.KeyScrollLock,
	"Semicolon":    ebiten.KeySemicolon,
	"Slash":        ebiten.KeySlash,
	"Space":        ebiten.KeySpace,
	"Tab":          ebiten.KeyTab,
	"Up":           ebiten.KeyUp,
	"Alt":          ebiten.KeyAlt,
	"Control":      ebiten.KeyControl,
	"Shift":        ebiten.KeyShift,
}
