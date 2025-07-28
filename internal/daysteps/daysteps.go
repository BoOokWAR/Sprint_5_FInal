package daysteps

import (
	"errors"
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
	//Разделение.
	parts := strings.Split(datastring, ",")

	//Проверка длины полученного парсинга.
	if len(parts) != 2 {
		return errors.New("полученная длина != 2")
	}
	//Проверка смены типа для шагов.
	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	//Проверка на наличие шагов.
	if ds.Steps <= 0 {
		return errors.New("шаги <= 0")
	}
	//Проверка парсинга длительности.
	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	// Проверяем, что длительность положительная.
	if ds.Duration <= 0 {
		return errors.New("длительность <= 0")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Расчет дистанции.
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	//Расчёт соженных каллорий.
	spentcalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", errors.New("ошибка в вычислении каллорий")
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		spentcalories), nil
}
