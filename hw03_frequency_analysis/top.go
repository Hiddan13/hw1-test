package main

import (
	"fmt"
	"sort"
	"strings"
)

const text1 = ``

type Words struct {
	Word string
	Num  int
}

var resSlice = []Words{}

func Top10(text string) []string {
	if text == "" {
		return make([]string, 0)
	}
	WordsMap := make(map[string]int)
	InputText := strings.Fields(text)
	// var SliceWords []string
	SliceWords := make([]string, 10)

	for _, word := range InputText {
		WordsMap[word]++
	}

	for key, value := range WordsMap {
		word := Words{key, value}
		resSlice = append(resSlice, word)
	}
	sort.Slice(resSlice, func(i, j int) bool {
		// если одинаковое количество раз встречается - то сортируем лексеграфически
		if resSlice[i].Num == resSlice[j].Num {
			return resSlice[i].Word < resSlice[j].Word
		}
		return resSlice[i].Num > resSlice[j].Num // иначе просто по каличеству
	})

	for index, wordinslice := range resSlice {
		if index == 10 {
			break
		} else {
			SliceWords[index] = wordinslice.Word
			// SliceWords = append(SliceWords, wordinslice.Word)
		}
	}
	return SliceWords
}

func main() {
	fmt.Println(Top10(text1))
}
