package main

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if sentence1 == sentence2 {
		return true
	}
	arr1, arr2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	if len(arr1) < len(arr2) {
		arr1, arr2 = arr2, arr1
	}
	diff := len(arr1) - len(arr2)
	for i := 0; i < len(arr1); i++ {
		for j := i + diff; j < len(arr1)+1; j++ {
			if equal(reConduct(arr1, i, j), arr2) {
				return true
			}
		}
	}
	return false
}
func reConduct(arr []string, i, j int) []string {
	res := []string{}
	res = append(res, arr[0:i]...)
	res = append(res, arr[j:]...)
	return res
}
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
