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
	if steps <= 0 {
		return 0, fmt.Errorf("Steps count error")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("Duration error")
	}

	if height <= 0 {
		return 0, fmt.Errorf("Height error")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("Weight error")
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * avgSpeed * durationInMinutes) / minInH

	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("Steps count error")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("Duration error")
	}

	if height <= 0 {
		return 0, fmt.Errorf("Height error")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("Weight error")
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * avgSpeed * durationInMinutes) / minInH

	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}

	walkDistance := Distance(steps, height)
	avgSpeed := walkDistance / duration.Hours()
	return avgSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceInM := stepLength * float64(steps)
	return distanceInM / mInKm
}
