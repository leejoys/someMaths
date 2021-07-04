package someMaths

import (
"math"	
)

//Разработайте RPC-сервер, который умеет вычислять расстояние между двумя точками на плоскости.
//В качестве аргумента функция должна принимать тип Points, объявленный следующим образом:
// Точка на плоскости.
type Point struct {
	X, Y float64
}

// Аргумент для функции Dist.
type Points struct {
	A, B Point
}

type Distance *int

//Сервер должен экспортировать функцию Dist, вычисляющую
//расстояние между двумя точками и записывающую результат
//вычислений во второй аргумент.
func Dist(p Points, d Distance) {
	*d = int(math.Sqrt((math.Pow(float64(p.B.X-p.A.X), 2)) -
		math.Pow(float64(p.B.Y-p.A.Y), 2)))
}