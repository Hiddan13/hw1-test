package main

import (
	"fmt"
	"sort"
	"strings"
)

const text2 = ""

var s_slice []string

var value_slice string

var num_word = 0

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
	var resultSlice = []Words{}
	var res []string
	s_slice = strings.Fields(t)
	for _, x := range s_slice {
		value_slice = x
		for _, c := range s_slice {
			if value_slice == c {
				num_word++
			}
		}
		a := Words{value_slice, num_word}
		resultSlice = append(resultSlice, a)
		num_word = 0
	}
	sort.Slice(resultSlice, func(i, j int) bool {
		if resultSlice[i].Num == resultSlice[j].Num { // если одинаковое количество раз встречается - то сортируем лексеграфически
			return resultSlice[i].Word < resultSlice[j].Word
		} else {
			return resultSlice[i].Num > resultSlice[j].Num // иначе просто по каличеству
		}
	})
	ss := Del_Replay(resultSlice)
	for a, s := range ss {
		if a == 10 {
			break
		} else {
			res = append(res, s.Word)
		}
	}
	return res
}

func Del_Replay(typeSlice []Words) []Words {
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
