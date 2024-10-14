package point

import "math"

// Point Создаем структуру с инкапсулированными параметрами x и y
type Point struct {
	x float64
	y float64
}

// CreatePoint Функция конструктор для создания новой точки
func CreatePoint(valueX, valueY float64) *Point {
	return &Point{
		x: valueX,
		y: valueY,
	}
}

// SetX метод для установки значения x
func (p *Point) SetX(valueX float64) {
	p.x = valueX
}

// SetY метод для установки значения y
func (p *Point) SetY(valueY float64) {
	p.y = valueY
}

// GetX метод для получения значения x
func (p *Point) GetX() float64 {
	return p.x
}

// GetY метод для получения значения y
func (p *Point) GetY() float64 {
	return p.y
}

// GetDistance метод для вычисления расстояния до другой точки
func (p *Point) GetDistance(otherPoint *Point) float64 {
	return math.Sqrt(math.Pow(otherPoint.x-p.x, 2) + math.Pow(otherPoint.y-p.y, 2))
}
