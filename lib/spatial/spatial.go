package spatial

import "github.com/limwa/advent-of-code-2024/lib/math"

type Vec2D struct {
	X int
	Y int
}

func (v Vec2D) ManhattanDistance(other Vec2D) int {
	return math.Abs(v.X-other.X) + math.Abs(v.Y-other.Y)
}

func (v Vec2D) EuclideanDistanceSqr(other Vec2D) int {
	return (v.X-other.X)*(v.X-other.X) + (v.Y-other.Y)*(v.Y-other.Y)
}

func (v Vec2D) Negate() Vec2D {
	return Vec2D{-v.X, -v.Y}
}

func (v Vec2D) Add(other Vec2D) Vec2D {
	return Vec2D{v.X + other.X, v.Y + other.Y}
}

func (v Vec2D) Sub(other Vec2D) Vec2D {
	return Vec2D{v.X - other.X, v.Y - other.Y}
}

func (v Vec2D) Dot(other Vec2D) int {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2D) Scale(factor float64) Vec2D {
	return Vec2D{int(float64(v.X) * factor), int(float64(v.Y) * factor)}
}

func (v Vec2D) IsWithinBounds(min Vec2D, max Vec2D) bool {
	return v.X >= min.X && v.X < max.X && v.Y >= min.Y && v.Y < max.Y
}
