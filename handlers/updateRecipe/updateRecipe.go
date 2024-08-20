package main

type Recipe struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}
