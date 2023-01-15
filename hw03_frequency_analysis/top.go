package main

import (
	"fmt"
	"sort"
	"strings"
)

const text1 = `К`

type Words struct {
	Word string
	Num  int
}

var resSlice = []Words{}

func Top10(text string) []string {
	ma := make(map[string]int)
	input := strings.Fields(text)

	for _, word := range input {
		_, match := ma[word]
		if match {
			ma[word] += 1
		} else {
			ma[word] = 1
		}

	}

	for key, value := range ma {
		aaa := Words{key, value}
		resSlice = append(resSlice, aaa)
	}
	sort.Slice(resSlice, func(i, j int) bool {
		// если одинаковое количество раз встречается - то сортируем лексеграфически
		if resSlice[i].Num == resSlice[j].Num {
			return resSlice[i].Word < resSlice[j].Word
		}
		return resSlice[i].Num > resSlice[j].Num // иначе просто по каличеству
	})

	var res []string
	for a, s := range resSlice {
		if a == 10 {
			break
		} else {
			res = append(res, s.Word)
		}
	}
	return res
}

func main() {
	fmt.Println(Top10(text1))
}
