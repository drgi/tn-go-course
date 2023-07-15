package geom

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.
// 1. Нужен коструктор и так как есть ограничение то конструктор вернет ошибку
// 2. Нужен именно возврат ошибки а не просто печать в консоль
// 3. Лучше использавать указатель в ф-ии CalculateDistance, хотя так как тут нет изменений внутри сруктуры не обязательно
// 4. Возможено стоит раздробить на более мелкие сущности, на точку(Point) из X и Y, а Geom был бы составом из двух точек?Но если брать эту задачу то может и усложнение.

type Geom struct {
	X1, Y1, X2, Y2 float64
}

func New(X1, Y1, X2, Y2 float64) (*Geom, error) {
	if X1 < 0 || X2 < 0 || Y1 < 0 || Y2 < 0 {
		return nil, fmt.Errorf("Координаты не могут быть меньше нуля")
	}
	return &Geom{
		X1: X1,
		Y1: Y1,
		X2: X2,
		Y2: Y2,
	}, nil
}

func (geom *Geom) CalculateDistance() (distance float64) {
	distance = math.Sqrt(math.Pow(geom.X2-geom.X1, 2) + math.Pow(geom.Y2-geom.Y1, 2))
	// возврат расстояния между точками
	return distance
}
