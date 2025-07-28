package trainings

import (
	"fmt"
	"log"
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

func (t *Training) Parse(datastring string) (err error) {
	//Разделение.
	parts := strings.Split(datastring, ",")
	//Проверка корреткности формата.
	if len(parts) != 3 {
		return fmt.Errorf("incorrect format")
	}
	//Проверка смены типа для шагов.
	t.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	//Проверка на наличие шагов.
	if t.Steps <= 0 {
		return fmt.Errorf("отсутствуют шаги")
	}
	//Вид активности.
	t.TrainingType = parts[1]

	//Продолжительность процесса.
	t.Duration, err = time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if t.Duration <= 0 {
		return fmt.Errorf("не положительная длительность")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {

	// Расчет дистанции.
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)

	//Расчёт средней скорости.
	meanspeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	//Проверка типа тренировки.
	var result string
	var sc float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		t.TrainingType = "Ходьба"
		sc, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			log.Printf("Ошибка: %v", err)
			return "", err
		}
	case "Бег":
		t.TrainingType = "Бег"
		sc, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			log.Printf("Ошибка: %v", err)
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	result = fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanspeed,
		sc)
	return result, nil

}
