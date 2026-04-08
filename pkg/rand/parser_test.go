package rand

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsRandExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid rand prefix", "rand:uuid", true},
		{"valid with options", "rand:email:domain=kil.com", true},
		{"env prefix", "env.API_URL", false},
		{"plain variable", "user_id", false},
		{"empty string", "", false},
		{"partial prefix", "ran:uuid", false},
		{"rand no colon", "randuuid", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsRandExpression(tt.input))
		})
	}
}

func TestParse_NoOptions(t *testing.T) {
	typeName, opts, err := Parse("rand:uuid")
	require.NoError(t, err)
	assert.Equal(t, "uuid", typeName)
	assert.Empty(t, opts)
}

func TestParse_SingleOption(t *testing.T) {
	typeName, opts, err := Parse("rand:email:domain=kil.com")
	require.NoError(t, err)
	assert.Equal(t, "email", typeName)
	assert.Equal(t, "kil.com", opts["domain"])
}

func TestParse_MultipleOptions(t *testing.T) {
	typeName, opts, err := Parse("rand:phone:country=FR,format=e164")
	require.NoError(t, err)
	assert.Equal(t, "phone", typeName)
	assert.Equal(t, "FR", opts["country"])
	assert.Equal(t, "e164", opts["format"])
}

func TestParse_IntOptions(t *testing.T) {
	typeName, opts, err := Parse("rand:int:min=1,max=100")
	require.NoError(t, err)
	assert.Equal(t, "int", typeName)
	assert.Equal(t, "1", opts["min"])
	assert.Equal(t, "100", opts["max"])
}

func TestParse_DateOptions(t *testing.T) {
	typeName, opts, err := Parse("rand:date:direction=future,format=2006-01-02")
	require.NoError(t, err)
	assert.Equal(t, "date", typeName)
	assert.Equal(t, "future", opts["direction"])
	assert.Equal(t, "2006-01-02", opts["format"])
}

func TestParse_WordsOption(t *testing.T) {
	typeName, opts, err := Parse("rand:words:count=5")
	require.NoError(t, err)
	assert.Equal(t, "words", typeName)
	assert.Equal(t, "5", opts["count"])
}

func TestParse_RegexPatternWithColonAndComma(t *testing.T) {
	// Pattern contains both ':' and ','
	typeName, opts, err := Parse(`rand:regex:pattern=[A-Z]{3}-\d{4}`)
	require.NoError(t, err)
	assert.Equal(t, "regex", typeName)
	assert.Equal(t, `[A-Z]{3}-\d{4}`, opts["pattern"])
}

func TestParse_RegexPatternWithComma(t *testing.T) {
	typeName, opts, err := Parse(`rand:regex:pattern=[a-z,A-Z]+`)
	require.NoError(t, err)
	assert.Equal(t, "regex", typeName)
	assert.Equal(t, `[a-z,A-Z]+`, opts["pattern"])
}

func TestParse_InvalidNotRandExpression(t *testing.T) {
	_, _, err := Parse("env.API_URL")
	assert.Error(t, err)
}

func TestParse_EmptyType(t *testing.T) {
	_, _, err := Parse("rand::option=val")
	assert.Error(t, err)
}

func TestParse_MissingPattern(t *testing.T) {
	_, _, err := Parse("rand:regex:nopattern=foo")
	assert.Error(t, err)
}

func TestParse_EmptyPattern(t *testing.T) {
	_, _, err := Parse("rand:regex:pattern=")
	assert.Error(t, err)
}

func TestParse_InvalidOptionFormat(t *testing.T) {
	_, _, err := Parse("rand:email:nodomain")
	assert.Error(t, err)
}
