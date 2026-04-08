package rand

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/nyaruka/phonenumbers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---- Generate() public API ----

func TestGenerate_UUID(t *testing.T) {
	val, err := Generate("rand:uuid")
	require.NoError(t, err)
	uuidRe := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	assert.Regexp(t, uuidRe, val)
}

func TestGenerate_UUID_IsDifferentEachCall(t *testing.T) {
	a, _ := Generate("rand:uuid")
	b, _ := Generate("rand:uuid")
	assert.NotEqual(t, a, b)
}

func TestGenerate_Email_NoOption(t *testing.T) {
	val, err := Generate("rand:email")
	require.NoError(t, err)
	assert.Contains(t, val, "@")
}

func TestGenerate_Email_WithDomain(t *testing.T) {
	val, err := Generate("rand:email:domain=kil.com")
	require.NoError(t, err)
	assert.True(t, strings.HasSuffix(val, "@kil.com"), "expected suffix @kil.com, got %s", val)
}

func TestGenerate_Phone_FR_E164(t *testing.T) {
	val, err := Generate("rand:phone:country=FR,format=e164")
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(val, "+33"), "expected +33 prefix for FR e164, got %s", val)

	num, err := phonenumbers.Parse(val, "FR")
	require.NoError(t, err)
	assert.True(t, phonenumbers.IsValidNumberForRegion(num, "FR"), "expected a valid FR number, got %s", val)
}

func TestGenerate_Phone_US_National(t *testing.T) {
	val, err := Generate("rand:phone:country=US,format=national")
	require.NoError(t, err)
	assert.NotEmpty(t, val)
}

func TestGenerate_Phone_US_International(t *testing.T) {
	val, err := Generate("rand:phone:country=US,format=international")
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(val, "+1"), "expected +1 prefix for US international, got %s", val)
}

func TestGenerate_Phone_RegionalE164Validity(t *testing.T) {
	tests := []struct {
		country string
		prefix  string
	}{
		{country: "FR", prefix: "+33"},
		{country: "US", prefix: "+1"},
		{country: "GB", prefix: "+44"},
		{country: "DE", prefix: "+49"},
		{country: "BE", prefix: "+32"},
		{country: "ES", prefix: "+34"},
	}

	for _, tt := range tests {
		t.Run(tt.country, func(t *testing.T) {
			val, err := Generate("rand:phone:country=" + tt.country + ",format=e164")
			require.NoError(t, err)
			assert.True(
				t,
				strings.HasPrefix(val, tt.prefix),
				"expected %s prefix for %s e164, got %s",
				tt.prefix,
				tt.country,
				val,
			)

			num, err := phonenumbers.Parse(val, tt.country)
			require.NoError(t, err)
			assert.True(
				t,
				phonenumbers.IsValidNumberForRegion(num, tt.country),
				"expected a valid %s number, got %s",
				tt.country,
				val,
			)
		})
	}
}

func TestGenerate_Phone_DefaultFormat_IsE164(t *testing.T) {
	val, err := Generate("rand:phone:country=US")
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(val, "+"), "expected E.164 (+) prefix as default, got %s", val)
}

func TestGenerate_Phone_InvalidFormat(t *testing.T) {
	_, err := Generate("rand:phone:country=FR,format=unknown")
	assert.Error(t, err)
}

func TestGenerate_Int_InRange(t *testing.T) {
	val, err := Generate("rand:int:min=5,max=10")
	require.NoError(t, err)
	n, err := strconv.Atoi(val)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, n, 5)
	assert.LessOrEqual(t, n, 10)
}

func TestGenerate_Int_Defaults(t *testing.T) {
	val, err := Generate("rand:int")
	require.NoError(t, err)
	n, err := strconv.Atoi(val)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, n, 0)
	assert.LessOrEqual(t, n, 1000)
}

func TestGenerate_Int_MinGreaterThanMax(t *testing.T) {
	_, err := Generate("rand:int:min=100,max=1")
	assert.Error(t, err)
}

func TestGenerate_Date_Past(t *testing.T) {
	val, err := Generate("rand:date:direction=past")
	require.NoError(t, err)
	parsed, err := time.Parse(time.RFC3339, val)
	require.NoError(t, err)
	assert.True(t, parsed.Before(time.Now()))
}

func TestGenerate_Date_Future(t *testing.T) {
	val, err := Generate("rand:date:direction=future")
	require.NoError(t, err)
	parsed, err := time.Parse(time.RFC3339, val)
	require.NoError(t, err)
	assert.True(t, parsed.After(time.Now()))
}

func TestGenerate_Date_Now(t *testing.T) {
	val, err := Generate("rand:date:direction=now")
	require.NoError(t, err)
	assert.NotEmpty(t, val)
}

func TestGenerate_Date_CustomFormat(t *testing.T) {
	val, err := Generate("rand:date:direction=past,format=2006-01-02")
	require.NoError(t, err)
	_, err = time.Parse("2006-01-02", val)
	require.NoError(t, err)
}

func TestGenerate_Date_InvalidDirection(t *testing.T) {
	_, err := Generate("rand:date:direction=yesterday")
	assert.Error(t, err)
}

func TestGenerate_Words(t *testing.T) {
	val, err := Generate("rand:words:count=4")
	require.NoError(t, err)
	assert.NotEmpty(t, val)
	// gofakeit.Sentence adds a period; words should be at least 4
	words := strings.Fields(val)
	assert.GreaterOrEqual(t, len(words), 4)
}

func TestGenerate_Words_Default(t *testing.T) {
	val, err := Generate("rand:words")
	require.NoError(t, err)
	assert.NotEmpty(t, val)
}

func TestGenerate_Regex(t *testing.T) {
	val, err := Generate(`rand:regex:pattern=[A-Z]{3}-\d{4}`)
	require.NoError(t, err)
	re := regexp.MustCompile(`^[A-Z]{3}-\d{4}$`)
	assert.Regexp(t, re, val, "generated value %q does not match pattern", val)
}

func TestGenerate_UnknownType(t *testing.T) {
	_, err := Generate("rand:foobar")
	assert.Error(t, err)
}
