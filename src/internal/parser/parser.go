package parser

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type DBReader interface {
	Parse(data []byte) (Recipe, error)
}

// JSONParser - структура для парсинга JSON
type JSONParser struct{}

func (jp *JSONParser) Parse(data []byte) (Recipe, error) {
	var recipe Recipe
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(data, &recipe)
	return recipe, err
}

// XMLParser - структура для парсинга XML
type XMLParser struct{}

func (xp *XMLParser) Parse(data []byte) (Recipe, error) {
	var recipe Recipe
	err := xml.Unmarshal(data, &recipe)
	return recipe, err
}

// ParseData - функция для парсинга данных в зависимости от формата
func ParseData(format string, data []byte) (Recipe, error) {
	var parser DBReader

	switch format {
	case "json":
		parser = &JSONParser{}
	case "xml":
		parser = &XMLParser{}
	default:
		return Recipe{}, fmt.Errorf("unsupported format: %s", format)
	}

	return parser.Parse(data)
}

// DetectFormatFromExtension определяет формат по расширению
func DetectFormatFromExtension(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		return "json"
	case ".xml":
		return "xml"
	default:
		return ""
	}
}
func IsSupportedFormat(format string) bool {
	switch format {
	case "json", "xml":
		return true
	default:
		return false
	}
}
