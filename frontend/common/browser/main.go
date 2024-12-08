package browser

import "time"

func CreateInstance(headlessMode bool, timeout, slowMotion time.Duration, incognitoMode bool) Browser {
	return newRodBrowser(headlessMode, timeout, slowMotion, incognitoMode)
}
