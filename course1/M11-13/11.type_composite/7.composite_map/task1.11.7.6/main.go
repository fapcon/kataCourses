package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Word struct {
	Word string
	Pos  int
}

// filterWords Фильтрует текст, заменяя цензурные и повторяющиеся слова
func filterWords(text string, censorMap map[string]string) string {
	// Разделение текста на предложения с помощью splitSentences
	var s, q []string
	var res string
	s = splitSentences(text)

	for i := 0; i < len(s); i++ {

		q = strings.Fields(s[i])
		// Создание пустой карты уникальных слов с помощью make(map[string]Word)
		/*
			uniqueWords := make(map[string]int)
			var unw []string
			for _, val := range q {
				uniqueWords[val]++
				if uniqueWords[val] == 1 {
					unw = append(unw, val)
				}
			}
			fmt.Println(unw)
		*/

		// ниже создание массива ключей ценсормапы

		keysCensor := make([]string, len(censorMap))
		z := 0
		for k := range censorMap {
			keysCensor[z] = k
			z++
		}

		for i := range q {
			for j, val := range keysCensor {
				if strings.ToLower(q[i]) == keysCensor[j] {
					q[i] = CheckUpper(q[i], censorMap[val])
				}
			}
		}

		uniqueWords := make(map[string]int)
		var unw []string
		for _, val := range q {
			uniqueWords[val]++
			if uniqueWords[val] == 1 {
				unw = append(unw, val)
			}
		}

		unwLow := make([]string, len(unw))
		for i := range unw {
			unwLow[i] = strings.ToLower(unw[i])
		}

		uniqueWords1 := make(map[string]int)
		var unwLow1 []string
		for _, val := range unwLow {
			uniqueWords1[val]++
			if uniqueWords1[val] == 1 {
				unwLow1 = append(unwLow1, val)
			}
		}

		unwLow1[0] = CheckUpper(unwLow1[0], q[0])
		res = res + WordsToSentence(unwLow1) + " "

	}
	end := strings.TrimSpace(res)

	//

	/*





		for i := 0; i < len(s); i++ {
			str = str + s[i] + " "
		}

		st := strings.TrimSpace(str)
		//
		// Если предложений больше одного, то обработка каждого предложения рекурсивно
		// Обработка каждого предложения в цикле
		// Рекурсивный вызов функции filterWords





		/*
			if len(s) > 1 {
				for _, val := range s {
					return filterWords(val, censorMap)
				}
				return strings.Join(s, " ")
			}

	*/

	/*



		// Прерывание блока условия "если предложений больше одного" c помощью return strings.Join(sentences, " ")

		// Разделение текста на отдельные слова с помощью strings.Fields(text)
		fmt.Println()
		fmt.Println()
		q = append(q, strings.Fields(st)...)
		fmt.Println(q)
		// Создание пустой карты уникальных слов с помощью make(map[string]Word)
		uniqueWords := make(map[string]Word)
		var unw []string
		for _, val := range q {
			if entry, ok := uniqueWords[val]; ok {
				entry.Pos++
				uniqueWords[val] = entry
				if uniqueWords[val].Pos == 1 {
					unw = append(unw, val)
				}
			}
		}
		fmt.Println(unw)

		// ниже создание массива ключей ценсормапы

		keysCensor := make([]string, len(censorMap))
		z := 0
		for k := range censorMap {
			keysCensor[z] = k
			z++
		}
		fmt.Println()
		fmt.Println()
		fmt.Println(keysCensor)
		fmt.Println()
		fmt.Println()
		for i := range q {
			for j, val := range keysCensor {
				if strings.ToLower(q[i]) == keysCensor[j] {
					q[i] = CheckUpper(q[i], censorMap[val])

				}
			}
		}
		fmt.Println()
		fmt.Println(keysCensor)
		fmt.Println()
		fmt.Println(q)

		res = WordsToSentence(q)

		fmt.Println()
		fmt.Println(res)





	*/

	// Обработка каждого слова в цикле
	// Если слово содержится в карте цензурных слов, то
	// Замена слова на значение из карты, используя CheckUpper

	// Если слово не содержится в карте уникальных слов, то (для проверки ключа в карте уникальных слов, используйте функцию strings.ToLower)
	// Добавление слова в карту уникальных слов
	// Продолжение выполнения цикла с помощью continue
	// Если слово содержится в карте уникальных слов, то нужно его очистить

	// Замена в слайсе слов при помощи карты уникальных слов и их индекса

	// Возвращение предложения из слайса слов, используйте функцию WordsToSentence
	return end
}

// WordsToSentence Удаляет пустые слова из слайса и объединяет их в предложение, добавляя в конце восклицательный знак
func WordsToSentence(words []string) string {
	filtered := make([]string, 0, len(words))

	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}

	return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
}

// CheckUpper Проверяет, нужно ли заменять первую букву на заглавную
func CheckUpper(old, new string) string {
	if len(old) == 0 || len(new) == 0 {
		return new
	}

	chars := []rune(old)

	if unicode.IsUpper(chars[0]) {
		runes := []rune(new)
		new = string(append([]rune{unicode.ToUpper(runes[0])}, runes[1:]...))
	}

	return new
}

// splitSentences Разделяет текст на предложения
func splitSentences(message string) []string {
	// Создание регулярного выражения для поиска знаков препинания
	originSentences := strings.Split(message, "!")
	var orphan string
	var sentences []string

	for i, sentence := range originSentences {
		words := strings.Split(sentence, " ")

		if len(words) == 1 {
			if len(orphan) > 0 {
				orphan += " "
			}

			orphan += words[0] + "!"
			continue
		}

		if orphan != "" {
			originSentences[i] = strings.Join([]string{orphan, " ", sentence}, " ") + "!"
			orphan = ""
		}

		sentences = append(sentences, originSentences[i])
	}

	return sentences
}

func main() {
	text := "Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! Беги, беги, успевай стать финансово независимым с помощью крипты! Крипта будущее финансового мира!"
	censorMap := map[string]string{
		"крипта":   "фрукты",
		"крипту":   "фрукты",
		"крипты":   "фруктов",
		"биткоин":  "яблоки",
		"лайткоин": "яблоки",
		"эфир":     "яблоки",
	}

	filteredText := filterWords(text, censorMap)
	fmt.Println(filteredText) // Внимание! Покупай срочно фрукты только у нас! Яблоки по низким ценам! Беги, успевай стать финансово независимым с помощью фруктов! Фрукты будущее финансового мира!
}
