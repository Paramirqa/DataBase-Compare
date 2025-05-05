package parser

import "strings"

// DiffResult - структура для хранения различий между рецептами
type DiffResult struct {
	CakeDiffs []CakeDiff `json:"cake_diffs"`
}

// CakeDiff - различия между двумя тортами
type CakeDiff struct {
	CakeName        string           `json:"cake_name"`
	TimeDifference  string           `json:"time_difference,omitempty"`
	MissingInFirst  bool             `json:"missing_in_first,omitempty"`
	MissingInSecond bool             `json:"missing_in_second,omitempty"`
	IngredientDiffs []IngredientDiff `json:"ingredient_diffs,omitempty"`
}

// IngredientDiff - различия между двумя ингредиентами
type IngredientDiff struct {
	Name            string `json:"name"`
	CountDifference string `json:"count_difference,omitempty"`
	UnitDifference  string `json:"unit_difference,omitempty"`
	MissingInFirst  bool   `json:"missing_in_first,omitempty"`
	MissingInSecond bool   `json:"missing_in_second,omitempty"`
}

// CompareRecipes сравнивает два рецепта и возвращает различия
func CompareRecipes(r1, r2 Recipe) DiffResult {
	var result DiffResult

	cakeMap2 := make(map[string]Cake)
	for _, cake := range r2.Cakes {
		cakeMap2[cake.Name] = cake
	}

	// Cakes from r1
	for _, cake1 := range r1.Cakes {
		cake2, exists := cakeMap2[cake1.Name]
		if !exists {
			result.CakeDiffs = append(result.CakeDiffs, CakeDiff{
				CakeName:        cake1.Name,
				MissingInSecond: true,
			})
			continue
		}

		cakeDiff := CakeDiff{
			CakeName: cake1.Name,
		}

		if strings.TrimSpace(cake1.Time) != strings.TrimSpace(cake2.Time) {
			cakeDiff.TimeDifference = cake1.Time + " vs " + cake2.Time
		}

		cakeDiff.IngredientDiffs = compareIngredients(cake1.Ingredients, cake2.Ingredients)

		if len(cakeDiff.IngredientDiffs) > 0 || cakeDiff.TimeDifference != "" {
			result.CakeDiffs = append(result.CakeDiffs, cakeDiff)
		}
	}

	// Cakes only in r2
	cakeMap1 := make(map[string]bool)
	for _, cake := range r1.Cakes {
		cakeMap1[cake.Name] = true
	}

	for _, cake2 := range r2.Cakes {
		if _, found := cakeMap1[cake2.Name]; !found {
			result.CakeDiffs = append(result.CakeDiffs, CakeDiff{
				CakeName:       cake2.Name,
				MissingInFirst: true,
			})
		}
	}

	return result
}

// Вспомогательная функция для сравнения ингредиентов
func compareIngredients(ings1, ings2 []Ingredient) []IngredientDiff {
	var diffs []IngredientDiff

	ingMap2 := make(map[string]Ingredient)
	for _, ing := range ings2 {
		ingMap2[ing.Name] = ing
	}

	for _, ing1 := range ings1 {
		ing2, exists := ingMap2[ing1.Name]
		if !exists {
			diffs = append(diffs, IngredientDiff{
				Name:            ing1.Name,
				MissingInSecond: true,
			})
			continue
		}

		if ing1.Count != ing2.Count || ing1.Unit != ing2.Unit {
			diffs = append(diffs, IngredientDiff{
				Name:            ing1.Name,
				CountDifference: ing1.Count + " vs " + ing2.Count,
				UnitDifference:  ing1.Unit + " vs " + ing2.Unit,
			})
		}
	}

	// Ingredients only in 2
	ingMap1 := make(map[string]bool)
	for _, ing := range ings1 {
		ingMap1[ing.Name] = true
	}

	for _, ing2 := range ings2 {
		if _, found := ingMap1[ing2.Name]; !found {
			diffs = append(diffs, IngredientDiff{
				Name:           ing2.Name,
				MissingInFirst: true,
			})
		}
	}

	return diffs
}
