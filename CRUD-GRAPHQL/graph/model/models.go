package models

type Ingredient struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	//Foreign key
	RecipeID int `json:"recipeId"`
}

type Recipe struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Procedure string `json:"procedure"`
	//Ingredients owned by a recipe
	Ingredients []Ingredient `json:"ingredients"`
}
