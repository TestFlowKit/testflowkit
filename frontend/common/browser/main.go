package browser

func CreateInstance(headlessMode bool) Browser {
	return newRodBrowser(headlessMode)
}
