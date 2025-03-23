package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Ingredient - структура для хранения информации об ингредиенте
type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"ingredient_name"`
	Count string `json:"ingredient_count" xml:"ingredient_count"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"ingredient_unit,omitempty"` // Используем omitempty, чтобы пропустить пустые поля
}

// Cake - структура для хранения информации о торте
type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"time"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients"`
}

// Recipe - структура для хранения всего рецепта
type Recipe struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}
type DBReader interface {
	DBReader(data []byte) (Recipe, error)
}

// ParseData - функция для парсинга данных в зависимости от формата
func parseData(format string, data []byte) (Recipe, error) {
	var parser DBReader

	switch format {
	case "json":
		parser = &JSONParser{}
	case "xml":
		parser = &XMLParser{}
	default:
		return Recipe{}, fmt.Errorf("unsupported format: %s", format)
	}

	return parser.DBReader(data)
}
