package tests

import (
	"Go_day01-1/nydiamig/internal/parser"
	"testing"
)

func TestJSONParser_Parse(t *testing.T) {
	data := []byte(`{
	  "cake": [
	    {
	      "name": "Test Cake",
	      "time": "45 min",
	      "ingredients": [
	        {"ingredient_name": "Sugar", "ingredient_count": "1", "ingredient_unit": "cup"},
	        {"ingredient_name": "Flour", "ingredient_count": "2"}
	      ]
	    }
	  ]
	}`)

	parser := &parser.JSONParser{}
	recipe, err := parser.Parse(data)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	if len(recipe.Cakes) != 1 {
		t.Errorf("expected 1 cake, got %d", len(recipe.Cakes))
	}

	cake := recipe.Cakes[0]
	if cake.Name != "Test Cake" {
		t.Errorf("expected cake name 'Test Cake', got '%s'", cake.Name)
	}

	if len(cake.Ingredients) != 2 {
		t.Errorf("expected 2 ingredients, got %d", len(cake.Ingredients))
	}

	ing := cake.Ingredients[0]
	if ing.Name != "Sugar" || ing.Count != "1" || ing.Unit != "cup" {
		t.Errorf("unexpected first ingredient: %+v", ing)
	}
}
