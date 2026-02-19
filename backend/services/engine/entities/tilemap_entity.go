package entities

// TileMapEntity representing a tilemap entity in the game engine.
type TileMapEntity struct {
	// Define the fields of the TileMapEntity struct here.
	*Entity                // Extends Entity
	Width      int         `json:"width"`          // Width  of the Entity in pixels
	Height     int         `json:"height"`         // Height of the Entity in pixels
	Properties []*Property `json:"fieldInstances"` // The Properties defined on the Entity
	TileRect   *TileRect   `json:"__tile"`
	Pivot      []float32   `json:"__pivot"` // Pivot position of the Entity (a centered Pivot would be 0.5, 0.5)
}

type Property struct {
	Identifier string `json:"__identifier"`
	Type       string `json:"__type"`  // The Type of the Property.
	Value      any    `json:"__value"` // The value contained within the property.
	// contains filtered or unexported fields
}

type TileRect struct {
	X          int      `json:"x"`
	Y          int      `json:"y"`
	W          int      `json:"w"`
	H          int      `json:"h"`
	TilesetUID int      `json:"tilesetUid"`
	Tileset    *Tileset `json:"-"`
}

type Tileset struct {
	Path       string `json:"relPath"` // Relative path to the tileset image; already is normalized using filepath.FromSlash().
	ID         int    `json:"uid"`
	GridSize   int    `json:"tileGridSize"`
	Spacing    int
	Padding    int
	Width      int `json:"pxWid"`
	Height     int `json:"pxHei"`
	Identifier string
	CustomData map[int]string  `json:"-"` // Key: tileID, Value: custom data string
	Enums      map[int]EnumSet `json:"-"` // Key: enumValueID, Value: tileIDs (tile indices)
}

type EnumSet []string
