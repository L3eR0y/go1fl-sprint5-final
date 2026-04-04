package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// format: "3456,Ходьба,3h00m"
func (t *Training) Parse(datastring string) (err error) {
	parsedData := strings.Split(datastring, ",")

	if len(parsedData) != 3 {
		return fmt.Errorf("Date string length error")
	}

	stepsCount, err := strconv.Atoi(parsedData[0])
	if err != nil || stepsCount <= 0 {
		return fmt.Errorf("Steps count error")
	}
	t.Steps = stepsCount

	activity := parsedData[1]
	t.TrainingType = activity

	activityDuration, err := time.ParseDuration(parsedData[2])

	if err != nil || activityDuration <= 0 {
		return fmt.Errorf("Activity duration error")
	}

	t.Duration = activityDuration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	resultString := ""

	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	activityDistance := spentenergy.Distance(t.Steps, t.Height)
	var spentCalories float64

	switch {
	case strings.ToLower(t.TrainingType) == "ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		if err != nil {
			return "", err
		}

		spentCalories = calories

	case strings.ToLower(t.TrainingType) == "бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		if err != nil {
			return "", err
		}

		spentCalories = calories

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	resultString = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), activityDistance, avgSpeed, spentCalories)

	return resultString, nil
}
