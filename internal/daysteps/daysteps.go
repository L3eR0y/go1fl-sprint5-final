package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	if len(datastring) < 1 {
		return fmt.Errorf("Data string is empty")
	}

	sData := strings.Split(datastring, ",")

	if len(sData) != 2 {
		return fmt.Errorf("Wrong format of date string")
	}

	stepsCount, err := strconv.Atoi(sData[0])
	if err != nil || stepsCount <= 0 {
		return fmt.Errorf("Steps count error")
	}
	ds.Steps = stepsCount

	duration, err := time.ParseDuration(sData[1])
	if err != nil || duration <= 0 {
		return fmt.Errorf("Duration time error")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", fmt.Errorf("Calories spent error")
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
