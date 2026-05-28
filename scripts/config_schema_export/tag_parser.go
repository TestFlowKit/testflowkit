package main

import (
	"fmt"
	"strconv"
	"strings"
)

const requiredIfPairLen = 2

type constraints struct {
	required   bool
	requiredIf string
	min        interface{}
	max        interface{}
	enum       []string
}

func parseValidateTag(tag string) constraints {
	c := constraints{}
	for _, part := range strings.Split(tag, ",") {
		part = strings.TrimSpace(part)
		switch {
		case part == "required":
			c.required = true
		case strings.HasPrefix(part, "required_if="):
			vals := strings.Fields(strings.TrimPrefix(part, "required_if="))
			if len(vals) == requiredIfPairLen {
				c.requiredIf = fmt.Sprintf("%s=%s", strings.ToLower(vals[0]), vals[1])
			}
		case strings.HasPrefix(part, "oneof="):
			c.enum = strings.Fields(strings.TrimPrefix(part, "oneof="))
		case strings.HasPrefix(part, "min="):
			c.min = parseNumber(strings.TrimPrefix(part, "min="))
		case strings.HasPrefix(part, "max="):
			c.max = parseNumber(strings.TrimPrefix(part, "max="))
		}
	}
	return c
}

func parseNumber(s string) interface{} {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	return nil
}
