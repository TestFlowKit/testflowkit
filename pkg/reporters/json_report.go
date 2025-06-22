package reporters

type jsonReport struct {
	StartDate string               `json:"start_date"`
	Duration  string               `json:"totalDuration"`
	Scenarios []jsonScenarioReport `json:"scenarios"`
}

type jsonScenarioReport struct {
	Title        string                   `json:"title"`
	Duration     string                   `json:"duration"`
	Result       string                   `json:"result"`
	Steps        []jsonScenarioStepReport `json:"steps"`
	ErrorMessage string                   `json:"error_message,omitzero"`
}

type jsonScenarioStepReport struct {
	Title    string `json:"title"`
	Status   string `json:"status"`
	Duration string `json:"duration"`
}
