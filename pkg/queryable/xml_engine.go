package queryable

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
)

type XMLEngine struct {
	doc *xmlquery.Node
}

func newXMLEngine(body []byte) (Queryable, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML body: %w", err)
	}

	return &XMLEngine{doc: doc}, nil
}

func (e *XMLEngine) Get(path string) (Result, error) {
	if err := validatePathSyntax(path, FormatXML); err != nil {
		return Result{}, err
	}

	node, err := xmlquery.Query(e.doc, path)
	if err != nil {
		return Result{}, fmt.Errorf("invalid XPath '%s': %w", path, err)
	}
	if node == nil {
		return Result{}, nil
	}

	return resultFromXMLNode(node), nil
}

func (e *XMLEngine) GetAll(path string) ([]Result, error) {
	if err := validatePathSyntax(path, FormatXML); err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(e.doc, path)
	if err != nil {
		return nil, fmt.Errorf("invalid XPath '%s': %w", path, err)
	}

	results := make([]Result, 0, len(nodes))
	for _, node := range nodes {
		results = append(results, resultFromXMLNode(node))
	}

	return results, nil
}

func (e *XMLEngine) Exists(path string) (bool, error) {
	if err := validatePathSyntax(path, FormatXML); err != nil {
		return false, err
	}

	node, err := xmlquery.Query(e.doc, path)
	if err != nil {
		return false, fmt.Errorf("invalid XPath '%s': %w", path, err)
	}

	return node != nil, nil
}

func resultFromXMLNode(node *xmlquery.Node) Result {
	raw := strings.TrimSpace(node.InnerText())
	kind := inferXMLKind(node, raw)

	return Result{
		Exists: true,
		Raw:    raw,
		Kind:   kind,
		Value:  raw,
	}
}

func inferXMLKind(node *xmlquery.Node, raw string) string {
	if node == nil {
		return string(TypeNull)
	}

	if hasElementChildren(node) {
		return string(TypeObject)
	}

	if raw == "" {
		return string(TypeString)
	}

	lowerRaw := strings.ToLower(raw)
	if _, err := strconv.ParseBool(lowerRaw); err == nil {
		return string(TypeBoolean)
	}
	if _, err := strconv.ParseInt(raw, 10, 64); err == nil {
		return string(TypeInteger)
	}
	if _, err := strconv.ParseFloat(raw, 64); err == nil {
		return string(TypeNumber)
	}

	return string(TypeString)
}

func hasElementChildren(node *xmlquery.Node) bool {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == xmlquery.ElementNode {
			return true
		}
	}
	return false
}
