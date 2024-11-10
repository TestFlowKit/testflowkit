package browser

import "time"

func CreateInstance(headlessMode bool, slowMotion time.Duration, incognitoMode bool) Browser {
	return newRodBrowser(headlessMode, slowMotion, incognitoMode)
}
