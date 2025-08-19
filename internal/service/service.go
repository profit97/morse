package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertString(input string) (string, error) {
	// Проверка на код Морзе
	if strings.ContainsAny(input, ".-") {
		// Конвертация из кода Морзе в текст
		return morse.ToText(input), nil
	} else {
		// Конвертация из текста в код Морзе
		return morse.ToMorse(input), nil
	}
}
