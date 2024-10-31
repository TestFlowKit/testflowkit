package browser

import "time"

func CreateInstance(headlessMode bool, slowMotion time.Duration) Browser {
	return newRodBrowser(headlessMode, slowMotion)
}
