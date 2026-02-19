package entities

import "github.com/Akif-jpg/MyHobieMMORPGGame/services/engine/components"

type Entity struct {
	Identifier string                 `json:"__identifier"` // Name of the Entity
	IID        string                 `json:"iid"`          // IID of the Entity
	Position   []int                  `json:"px"`           // Position of the Entity (x, y)
	Tags       []string               `json:"__tags"`       // Tags (categories) assigned to the Entity
	Components []components.Component `json:"components"`   // Components of the Entity
	Data       any                    `json:"-"`            // Data allows you to attach key custom data to the entity post-parsing
}
