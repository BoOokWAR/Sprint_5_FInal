package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Проверка на наличие шагов.
	if steps <= 0 {
		return 0, fmt.Errorf("неверное количество шагов")
	}
	//Проверка веса
	if weight <= 0 {
		return 0, fmt.Errorf("вес меньше или равен 0")
	}
	//Проверка роста
	if height <= 0 {
		return 0, fmt.Errorf("рост меньше или равен 0")
	}
	//Проверка наличия длителности больше 0.
	if duration <= 0 {
		return 0, fmt.Errorf("неверная длительность тренировки")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return walkingCaloriesCoefficient * (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Проверка на наличие шагов.
	if steps <= 0 {
		return 0, fmt.Errorf("неверное количество шагов")
	}
	//Проверка веса
	if weight <= 0 {
		return 0, fmt.Errorf("вес меньше или равен 0")
	}
	//Проверка роста
	if height <= 0 {
		return 0, fmt.Errorf("рост меньше или равен 0")
	}
	//Проверка наличия длителности больше 0.
	if duration <= 0 {
		return 0, fmt.Errorf("неверная длительность тренировки")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	//Проверка на наличие шагов.
	if steps <= 0 {
		return 0
	}
	//Проверка роста.
	if height <= 0 {
		return 0
	}
	//Проверка длительности больше 0.
	if duration <= 0 {
		return 0
	}
	//Перевод в часы.
	durationInHours := duration.Hours()

	//Расчёт дистанции.
	d := Distance(steps, height)
	return d / durationInHours
}

func Distance(steps int, height float64) float64 {
	//Проверка на наличие шагов.
	if steps <= 0 {
		return 0
	}
	//Проверка роста.
	if height <= 0 {
		return 0
	}
	//Расчёт длины шага.
	stepLength := float64(stepLengthCoefficient) * height
	return stepLength * float64(steps) / mInKm
}
