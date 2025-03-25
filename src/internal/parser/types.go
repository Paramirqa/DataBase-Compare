package parser

// Ingredient - структура для хранения информации об ингредиенте
type Ingredient struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit,omitempty" json:"ingredient_unit,omitempty"`
}

// Cake - структура для хранения информации о торте
type Cake struct {
	Name        string       `xml:"name" json:"name"`
	Time        string       `xml:"stovetime" json:"time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

// Recipe - структура для хранения всего рецепта
type Recipe struct {
	Cakes []Cake `xml:"cake" json:"cake"`
}
