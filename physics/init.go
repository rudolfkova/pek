package physics

import (
	"github.com/rudolfkova/pek/vec"
	"github.com/rudolfkova/pek/entity"
)

// Инициализация динамических объектов
func InitDyn(dyn ...*entity.Object) {
	dynObj = append(dynObj, dyn...)
}

// Инициализация статических объектов
func InitStat(stat ...*entity.Object) {
	statObj = append(statObj, stat...)
}

// Инициализация векторов границ статических объектов
func NewStatVec() {
	for _, s := range statObj {
		//AB
		X1_AB := s.X
		Y1_AB := s.Y
		X2_AB := s.X + float64(s.Width)
		Y2_AB := s.Y
		//BC
		X1_BC := s.X + float64(s.Width)
		Y1_BC := s.Y
		X2_BC := s.X + float64(s.Width)
		Y2_BC := s.Y + float64(s.Height)
		//DC
		X1_DC := s.X
		Y1_DC := s.Y + float64(s.Height)
		X2_DC := s.X + float64(s.Width)
		Y2_DC := s.Y + float64(s.Height)
		//AD
		X1_AD := s.X
		Y1_AD := s.Y
		X2_AD := s.X
		Y2_AD := s.Y + float64(s.Height)
		s.AB = *Line.NewVec(X1_AB, Y1_AB, X2_AB, Y2_AB)
		s.BC = *Line.NewVec(X1_BC, Y1_BC, X2_BC, Y2_BC)
		s.DC = *Line.NewVec(X1_DC, Y1_DC, X2_DC, Y2_DC)
		s.AD = *Line.NewVec(X1_AD, Y1_AD, X2_AD, Y2_AD)
		statVec = append(statVec, &s.AB)
		statVec = append(statVec, &s.BC)
		statVec = append(statVec, &s.DC)
		statVec = append(statVec, &s.AD)
	}
}

// Слайсы для хранения динамических и статических объектов
var dynObj []*entity.Object
var statObj []*entity.Object

// Слайсы для хранения векторов границ статических объектов
var statVec []*Line.Line
