package mac

import (
	"github.com/andybrewer/mack"
)

func UnSleep() string {
	osa, err := mack.Tell("System Events", "do shell script \"caffeinate -u -t 1\"")
	if err != nil {
		return err.Error()
	}
	return osa
}

func Sleep() string {
	osa, err := mack.Tell("Finder", "sleep")
	if err != nil {
		return err.Error()
	}
	return osa
}
