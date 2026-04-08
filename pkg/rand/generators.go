package rand

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/lucasjones/reggen"
	"github.com/nyaruka/phonenumbers"
)

const (
	defaultPhoneMaxAttempts = 20
	phonePrefixMinMobile    = 6
	phonePrefixMaxMobile    = 7
	usAreaCodeMin           = 201
	usAreaCodeMax           = 989
	usExchangeCodeMin       = 200
	usExchangeCodeMax       = 999
	maxFourDigits           = 9999
	maxSevenDigits          = 9999999
	maxEightDigits          = 99999999
	maxNineDigits           = 999999999
	pastYearsRange          = -10
	futureYearsRange        = 10
)

var germanMobilePrefixes = []string{
	"0151", "0152", "0155", "0157", "0159", "0160", "0162", "0163",
	"0170", "0171", "0172", "0173", "0174", "0175", "0176", "0177",
	"0178", "0179",
}

func genUUID(_ map[string]string) (string, error) {
	return gofakeit.UUID(), nil
}

func genEmail(opts map[string]string) (string, error) {
	domain, hasDomain := opts["domain"]
	if !hasDomain || domain == "" {
		return gofakeit.Email(), nil
	}
	// Build <username>@<domain>
	username := gofakeit.Username()
	return username + "@" + domain, nil
}

func genPhone(opts map[string]string) (string, error) {
	country := strings.ToUpper(optOrDefault(opts, "country", "US"))
	format := strings.ToLower(optOrDefault(opts, "format", "e164"))

	switch format {
	case "e164", "national", "international":
	default:
		return "", fmt.Errorf("rand:phone unknown format %q: expected e164, national, or international", format)
	}

	for range defaultPhoneMaxAttempts {
		raw := randomPhoneForCountry(country)
		num, err := phonenumbers.Parse(raw, country)
		if err != nil || !phonenumbers.IsValidNumberForRegion(num, country) {
			continue
		}

		switch format {
		case "e164":
			return phonenumbers.Format(num, phonenumbers.E164), nil
		case "national":
			return phonenumbers.Format(num, phonenumbers.NATIONAL), nil
		default:
			return phonenumbers.Format(num, phonenumbers.INTERNATIONAL), nil
		}
	}

	return "", fmt.Errorf("rand:phone exhausted attempts for country %q", country)
}

func randomPhoneForCountry(country string) string {
	switch country {
	case "FR":
		// French mobile numbers: 06/07XXXXXXXX
		return fmt.Sprintf(
			"0%d%08d",
			gofakeit.Number(phonePrefixMinMobile, phonePrefixMaxMobile),
			gofakeit.Number(0, maxEightDigits),
		)
	case "US":
		// Basic NANP-compatible shape
		return fmt.Sprintf(
			"%03d%03d%04d",
			gofakeit.Number(usAreaCodeMin, usAreaCodeMax),
			gofakeit.Number(usExchangeCodeMin, usExchangeCodeMax),
			gofakeit.Number(0, maxFourDigits),
		)
	case "GB":
		// UK mobile numbers: 07XXXXXXXXX
		return fmt.Sprintf("07%09d", gofakeit.Number(0, maxNineDigits))
	case "DE":
		// German mobile numbers use prefixes such as 0151, 0152, 0155, 0157, 0159, 0160, 0162, 0163, 0170-0179.
		prefix := germanMobilePrefixes[gofakeit.Number(0, len(germanMobilePrefixes)-1)]
		// Length varies by carrier; generate 7-8 trailing digits and let phonenumbers validate.
		if gofakeit.Bool() {
			return fmt.Sprintf("%s%07d", prefix, gofakeit.Number(0, maxSevenDigits))
		}
		return fmt.Sprintf("%s%08d", prefix, gofakeit.Number(0, maxEightDigits))
	case "BE":
		// Belgian mobile numbers: 04XXXXXXXX
		return fmt.Sprintf("04%08d", gofakeit.Number(0, maxEightDigits))
	case "ES":
		// Spanish mobile numbers: 6XXXXXXXX or 7XXXXXXXX
		return fmt.Sprintf(
			"%d%08d",
			gofakeit.Number(phonePrefixMinMobile, phonePrefixMaxMobile),
			gofakeit.Number(0, maxEightDigits),
		)
	default:
		return gofakeit.Phone()
	}
}

func genInt(opts map[string]string) (string, error) {
	minVal, err := strconv.Atoi(optOrDefault(opts, "min", "0"))
	if err != nil {
		return "", fmt.Errorf("rand:int invalid min option: %w", err)
	}
	maxVal, err := strconv.Atoi(optOrDefault(opts, "max", "1000"))
	if err != nil {
		return "", fmt.Errorf("rand:int invalid max option: %w", err)
	}
	if minVal > maxVal {
		return "", fmt.Errorf("rand:int min (%d) must be less than or equal to max (%d)", minVal, maxVal)
	}
	return strconv.Itoa(gofakeit.Number(minVal, maxVal)), nil
}

func genDate(opts map[string]string) (string, error) {
	direction := strings.ToLower(optOrDefault(opts, "direction", "past"))
	layout := optOrDefault(opts, "format", time.RFC3339)

	now := time.Now()
	var start, end time.Time

	switch direction {
	case "past":
		start = now.AddDate(pastYearsRange, 0, 0)
		end = now
	case "future":
		start = now
		end = now.AddDate(futureYearsRange, 0, 0)
	case "now":
		return now.Format(layout), nil
	default:
		return "", fmt.Errorf("rand:date unknown direction %q: expected past, future, or now", direction)
	}

	return gofakeit.DateRange(start, end).Format(layout), nil
}

func genWords(opts map[string]string) (string, error) {
	count, err := strconv.Atoi(optOrDefault(opts, "count", "3"))
	if err != nil {
		return "", fmt.Errorf("rand:words invalid count option: %w", err)
	}
	if count < 1 {
		return "", fmt.Errorf("rand:words count must be >= 1, got %d", count)
	}
	return gofakeit.Sentence(count), nil
}

func genRegex(opts map[string]string) (string, error) {
	pattern, ok := opts["pattern"]
	if !ok || pattern == "" {
		return "", errors.New("rand:regex requires a non-empty pattern= option")
	}
	const regexLimit = 10
	result, err := reggen.Generate(pattern, regexLimit)
	if err != nil {
		return "", fmt.Errorf("rand:regex failed to generate string from pattern %q: %w", pattern, err)
	}
	return result, nil
}

func optOrDefault(opts map[string]string, key, defaultVal string) string {
	if v, ok := opts[key]; ok && v != "" {
		return v
	}
	return defaultVal
}
