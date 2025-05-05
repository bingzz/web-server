package data

import "web-server/src/schema"

// In-memory store caching
var DataStore = schema.DataStore{}

// Initializing data
var Albums = []schema.Album{
	{
		ID:       1,
		Title:    "Blue Train",
		AuthorID: 1,
		Price:    56.99,
	},
	{
		ID:       2,
		Title:    "Jeru",
		AuthorID: 1,
		Price:    17.99,
	},
	{
		ID:       3,
		Title:    "Sarah Vaughan and Clifford Brown",
		AuthorID: 1,
		Price:    39.99,
	},
}
