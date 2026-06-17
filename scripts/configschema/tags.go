package configschema

import (
	"reflect"
	"strings"
	"testflowkit/internal/config"
)

func parseYAMLTag(tag string) (name string, skip bool) {
	if tag == "-" {
		return "", true
	}
	if tag == "" {
		return "", false
	}
	parts := strings.Split(tag, ",")
	return parts[0], false
}

func isFieldRequired(field reflect.StructField, validateTag string) bool {
	if field.Type.Kind() == reflect.Pointer {
		return false
	}
	for _, part := range strings.Split(validateTag, ",") {
		if strings.TrimSpace(part) == "required" {
			return true
		}
	}
	return false
}

func dereferenceType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

func parseOneOfEnum(validateTag string) []string {
	for _, part := range strings.Split(validateTag, ",") {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "oneof=") {
			raw := strings.TrimPrefix(part, "oneof=")
			return strings.Fields(raw)
		}
	}
	return nil
}

func hardcodedEnum(typeName string) []string {
	switch typeName {
	case "APIType":
		return []string{string(config.APITypeREST), string(config.APITypeGraphQL)}
	case "SecuritySchemeType":
		return []string{
			string(config.SecurityTypeBearer),
			string(config.SecurityTypeBasic),
			string(config.SecurityTypeAPIKey),
			string(config.SecurityTypeOAuth2),
			string(config.SecurityTypeOIDC),
			string(config.SecurityTypeCertificate),
			string(config.SecurityTypeNone),
		}
	case "APIKeyPlacement":
		return []string{
			string(config.APIKeyPlacementHeader),
			string(config.APIKeyPlacementQuery),
			string(config.APIKeyPlacementCookie),
		}
	case "OAuth2TokenAuthMethod":
		return []string{
			string(config.OAuth2TokenAuthMethodPost),
			string(config.OAuth2TokenAuthMethodBasic),
		}
	default:
		return nil
	}
}
