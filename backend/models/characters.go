package models

// Character struct represent player in the game. One player can have multiple characters.
type Character struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Experience int    `json:"experience"`
}
