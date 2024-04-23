package structs

import (
	"math"
)

const pi = 22 / 7

type Segitiga struct {
	Alas   int `json:"alas"`
	Tinggi int `json:"tinggi"`
}
type Persegi struct {
	Sisi int `json:"sisi"`
}
type PersegiPanjang struct {
	Panjang, Lebar int
}
type Lingkaran struct {
	JariJari int
}

func (s Segitiga) Luas() float64 {
	return float64(s.Alas * s.Tinggi / 2)
}
func (s Persegi) Luas() float64 {
	return float64(math.Pow(float64(s.Sisi), 2))
}
func (s PersegiPanjang) Luas() float64 {
	return float64(s.Panjang * s.Lebar)
}
func (s Lingkaran) Luas() float64 {
	return math.Pow(float64(s.JariJari), 2) * pi
}

func (s Segitiga) Keliling() float64 {
	return float64(s.Alas * 3)
}
func (s Persegi) Keliling() float64 {
	return float64(s.Sisi * 4)
}
func (s PersegiPanjang) Keliling() float64 {
	return float64((2 * s.Panjang) + (2 * s.Lebar))
}
func (s Lingkaran) Keliling() float64 {
	return 2 * pi * float64(s.JariJari)
}

type Hitung interface {
	Luas() int
	Keliling() int
}
