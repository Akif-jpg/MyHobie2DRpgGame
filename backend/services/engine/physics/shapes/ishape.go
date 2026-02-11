package shapes

const (
	CircleType    = "circle"
	RectangleType = "rectangle"
)

type Shape interface {
	GetType() string
	GetCenter() Point
	SetCenter(center Point)

	IntersectsRectangle(other *Rectangle) bool
	IntersectsCircle(other *Circle) bool
	IntersectsLine(other *Line) bool
	IntersectsPoint(point *Point) bool

	ContainsRectangle(other *Rectangle) bool
	ContainsCircle(other *Circle) bool
	ContainsLine(other *Line) bool
	ContainsPoint(point *Point) bool
}
