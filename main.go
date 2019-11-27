package goarea

import "math"

// Pi é a proporcao entre o perimetro da circunferencia e seu diametro
const Pi = 3.1416

// Circ é responsável por calcular a área da circunferẽncia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
