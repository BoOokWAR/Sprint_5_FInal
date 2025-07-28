package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	//Перебираем слайс и проверяем ошибку парсинга
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка парсинга: %v", err)
			continue
		}
		// Получаем информацию о тренировке
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка получения информации: %v", err)
			continue
		}
		fmt.Println(info)
	}
}
