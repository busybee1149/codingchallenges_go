package utils

func GetCharacterFrequency(str string) map[rune]int {
	characterMap := make(map[rune]int)
	for _, ch := range str {
		characterMap[ch]++
	}
	return characterMap
}