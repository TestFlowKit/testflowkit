package form

import (
	"regexp"
	"testing"
)

func TestPrivateIFillTheInputWithSentences(t *testing.T) {
	wildcard := "{string}"
	expectedWildcardNumber := 2

	re := regexp.MustCompile(wildcard)

	for _, sentence := range iFillTheInputWith.Sentences {
		occurs := len(re.FindAllString(sentence, -1))
		if occurs == expectedWildcardNumber {
			continue
		}
		t.Fatalf("all sentencesmust contains %d wildcars but \"%s\" contains %d", expectedWildcardNumber, sentence, occurs)
	}
}
