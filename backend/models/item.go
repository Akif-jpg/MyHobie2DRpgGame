package models

// Item struct represents an item in the game.
type Item struct {
	UUID        string `json:"id"` // Every item has uuid because it is used to identify the item in the game.
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Effects     string `json:"effects"` // Effects store as json data and game's rule engine apply these effects.
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	CreatedAt   string `json:"created_at"`
}
