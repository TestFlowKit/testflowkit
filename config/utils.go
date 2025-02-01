package config

func IsElementDefined(elementName string) bool {
	fc := FrontConfig{}
	_, err := fc.GetHTMLElementSelectors(elementName)
	return err == nil
}

func IsPageDefined(pageName string) bool {
	fc := FrontConfig{}
	_, err := fc.GetPageURL(pageName)
	return err == nil
}
