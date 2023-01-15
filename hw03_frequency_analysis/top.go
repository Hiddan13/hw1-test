package main

import (
	"fmt"
	"sort"
	"strings"
)

const text2 = ""

var sSlice []strin

var valueSlice string

var numWord = 0

type Words struct {
	Word string
	Num  int
}

var mySlice []string

func main() {
	mySlice = Top10(text2)
	fmt.Println(mySlice)
}

func Top10(t string) []string {
	resultSlice := []Words{}
	var res []string
	sSlice = strings.Fields(t)
	for _, x := range sSlice {
		valueSlice = x
		for _, c := range sSlice {
			if valueSlice == c {
				numWord++
			}
		}
		a := Words{valueSlice, numWord}
		resultSlice = append(resultSlice, a)
		numWord = 0
	}
	sort.Slice(resultSlice, func(i, j int) bool {
		// если одинаковое количество раз встречается - то сортируем лексеграфически
		if resultSlice[i].Num == resultSlice[j].Num {
			return resultSlice[i].Word < resultSlice[j].Word
		}
		return resultSlice[i].Num > resultSlice[j].Num // иначе просто по каличеству
	})
	ss := DelReplay(resultSlice)
	for a, s := range ss {
		if a == 10 {
			break
		} else {
			res = append(res, s.Word)
		}
	}
	return res
}

func DelReplay(typeSlice []Words) []Words {
	keys := make(map[Words]bool)
	list := []Words{}
	for _, entry := range typeSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
