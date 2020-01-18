package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// LogError prints an error and exits
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// PrintFPS prints the current TPS and FPS values
func PrintFPS(screen *ebiten.Image) {
	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)
}

// GetConfig reads a config file and returns a new config map
func GetConfig(configFile string, config *map[string]string) {
	file, err := os.Open(configFile)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var kv = strings.Split(scanner.Text(), " ")
		if len(kv) == 2 {
			(*config)[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	LogError(scanner.Err())
	return
}
