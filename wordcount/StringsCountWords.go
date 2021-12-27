package wordcount

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type Request struct {
	Text string `json:"text"`
}

//func CountWords
func CountWords(response http.ResponseWriter, request *http.Request) {
	var inputData Request
	json.NewDecoder(request.Body).Decode(&inputData)
	fmt.Println("Input Text :", inputData.Text)
	//calling repeatedWordAndCount func for getting words and count of all text.
	wc := repeatedWordAndCount(inputData.Text)
	var keys []int
	for _, val := range wc {
		keys = append(keys, val)
	}

	sort.Ints(keys)
	//getting highest 10 numbers data
	data := topTenRecords(keys, inputData.Text)

	for word, count := range data {
		var data = make(map[string]int)
		data[word] = count

	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range data {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	//assigning most highest array value to ss
	ss = ss[len(ss)-10:]

	json.NewEncoder(response).Encode(ss)
}

//func repeatedWordAndCount for getting words and count of all text.
func repeatedWordAndCount(st string) (wc map[string]int) {

	// using strings.Field Function
	input := strings.Fields(st)
	wc = make(map[string]int)
	for _, word := range input {
		_, match := wc[word]
		if match {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

//func topTenRecords for getting highest 10 values from text.
func topTenRecords(intSlice []int, sentence string) map[string]int {

	data := make(map[string]int)
	if len(intSlice) > 10 {
		topfinal := intSlice[len(intSlice)-10:]

		for key, val := range repeatedWordAndCount(sentence) {

			for _, value := range topfinal {

				if value == val {
					data[key] = val
				}
			}

		}

	}

	return data
}
