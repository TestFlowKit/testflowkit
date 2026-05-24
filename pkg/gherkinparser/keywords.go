package gherkinparser

// DocStringDelimiter is the default Gherkin doc string delimiter.
const DocStringDelimiter = "\"\"\""

const (
	gherkinKeywordGiven = "Given"
	gherkinKeywordWhen  = "When"
	gherkinKeywordThen  = "Then"
	gherkinKeywordAnd   = "And"
	gherkinKeywordBut   = "But"
)

// DocStringExample wraps body content in a Gherkin doc string block.
func DocStringExample(prefix, body string) string {
	return prefix + "\n" + DocStringDelimiter + "\n" + body + "\n" + DocStringDelimiter
}
