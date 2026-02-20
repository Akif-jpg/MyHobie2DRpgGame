package models

type Inventory struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	CharacterID int    `json:"character_id"`
	ItemList    []int  `json:"item_list"`
	Quantity    int    `json:"quantity"`
	CreatedAt   string `json:"created_at"`
}
