package golang

import "strings"

func wordPattern(pattern string, s string) bool {
	wordList := make(map[string]string)
	symbolList := make(map[string]string)
	strArray := strings.Split(s, " ")
	if len(strArray) != len(pattern) {
		return false
	}
	for i, v := range pattern {
		symbolString := string(v)
		wordString := strArray[i]
		wordS, ok := wordList[symbolString]
		symbolS, ok2 := symbolList[wordString]
		if ok && ok2 {
			if wordS != wordString || symbolS != symbolString {
				return false
			}
			continue
		}
		if !ok && !ok2 {
			wordList[symbolString] = wordString
			symbolList[wordString] = symbolString
			continue
		}
		return false
	}
	return true
}
