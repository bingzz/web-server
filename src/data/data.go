package data

import "web-server/src/schema"

// In-memory store caching
var DataStore = schema.DataStore{}

// Initializing data
var Albums = []schema.Album{
	{
		ID:     "al-1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "al-2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
	{
		ID:     "al-3",
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	},
}
