package api

import (
	"fmt"

	"github.com/andybrewer/mack"
)

func Keycode(code string) string {
	osa, err := mack.Tell("System Events", fmt.Sprintf("key code %s", code))
	if err != nil {
		return err.Error()
	}
	return osa
}

func Keystroke(isSpecial bool, key string) string {
	var cmd string
	if isSpecial {
		cmd = fmt.Sprintf("keystroke %s", key)
	} else {
		cmd = fmt.Sprintf("keystroke \"%s\"", key)
	}

	osa, err := mack.Tell("System Events", cmd)
	if err != nil {
		return err.Error()
	}
	return osa
}
