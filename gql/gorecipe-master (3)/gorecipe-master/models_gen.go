// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gorecipe

type NewIngredient struct {
	Name string `json:"name"`
}

type NewRecipe struct {
	Name      string  `json:"name"`
	Procedure *string `json:"procedure"`
}