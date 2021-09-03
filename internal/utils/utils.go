package utils

// Chunks рзделяет слайс на батчи заданного размера.
func Chunks(items []int, size int) [][]int {
	if size <= 0 {
		panic("size param must be positive")
	}

	if items == nil {
		panic("size param must be positive")
	}

	itemsLength := len(items)
	chunksLength := itemsLength / size

	if chunksLength == 0 || itemsLength%size > 0 {
		chunksLength++
	}

	chunks := make([][]int, chunksLength)

	for i := 0; i < chunksLength; i++ {
		start := i * size
		end := i*size + size

		if end > itemsLength {
			end = itemsLength
		}

		chunks[i] = items[start:end]
	}

	return chunks
}

// InvertMap конвертирует отображение (“ключ-значение“) в отображение (“значение-ключ“).
func InvertMap(data map[string]int) map[int]string {
	result := make(map[int]string, len(data))
	for k, v := range data {
		if _, ok := result[v]; ok {
			panic("equal values can not invert to key")
		}

		result[v] = k
	}

	return result
}

// FilterStopWords фильтрует входной слайс на стоп-слова.
func FilterStopWords(words []string) []string {
	stopWords := []string{"Nigger", "Rainbow", "Timati"}
	result := make([]string, 0, len(words))

	for _, word := range words {
		if contains(stopWords, word) {
			continue
		}

		result = append(result, word)
	}

	return result
}

func contains(s []string, item string) bool {
	for _, value := range s {
		if value == item {
			return true
		}
	}

	return false
}
