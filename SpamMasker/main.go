package main

import (
	"fmt"
)

func HideLinks(inputString string) string {
	hidedString := []byte(inputString) // создаю срез для строки со скрытой ссылкой

	hideMark := byte('*')    // маскировочный символ
	prefix := "http://"      // префикс строки, байты после нее и до первого пробела заменяю на *
	prefixLen := len(prefix) // длина префикса 7

	hideMode := false // режим false - обычный символ, иначе замаскированный

	for i := 0; i < len(inputString); i++ { // перебираем исходную строку

		// ищем префикс в строке и если нашли, то переводим в режим маскировки
		if len(inputString)-i >= prefixLen &&
			inputString[i:i+prefixLen] == prefix {

			hideMode = true    // то переводим в режим маскировки
			i += prefixLen - 1 // пропускаем сам префикс
			continue
		}

		// снимаем режим маскировки если встречаем пробел
		if hideMode && inputString[i] == ' ' {
			hideMode = false
		}

		// в режиме маскировки меняем текущий символ на скрытый
		if hideMode {
			hidedString[i] = hideMark
		}
	}

	return string(hidedString)
}

func main() {
	inputString := "Here's my spammy page: http://hehefouls.netHAHAHA see you." // исходная строка
	hidedString := HideLinks(inputString)

	fmt.Printf("Исходная строка %s\n", inputString)
	fmt.Printf("Замаскированная строка %s", string(hidedString))
}
