package goarea

import "math"

//Pi é uma proporção numérica
const Pi = 3.1416

// Comentario
func Circ(raio float64) float64 {
	return math.Pow(raio,2) * Pi
}

// Comentario
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
