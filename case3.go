package main

import "bytes"

func charCap(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}

	return ch
}

func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if charCap(word[i]) != charCap(word[len(word)-1-i]) {
			return false
		}
	}

	return true
}

func addHashtags(word string) string {
	var buf bytes.Buffer

	buf.WriteByte('#')

	for i := 0; i < len(word); i++ {
		buf.WriteByte(word[i])
		buf.WriteByte('#')
	}

	return buf.String()
}

func longestPalindromSubstring(word string) int {
	hashtagedWord := addHashtags(word)
	lspVals := make([]int, len(hashtagedWord))

	for i := 0; i < len(hashtagedWord); i++ {
		if i == 0 {
			continue
		}

		var j = 1
		for ; i-j >= 0 && i+j+1 <= len(hashtagedWord); j++ {
			if !isPalindrome(hashtagedWord[i-j : i+j+1]) {
				break
			}
		}

		lspVals[i] = j - 1

		if lspVals[i] <= 1 {
			continue
		}

		var k = 1
		for ; k < lspVals[i] && i+k < len(hashtagedWord); k++ {
			if lspVals[i-k]+i <= lspVals[i]+i {
				lspVals[i+k] = lspVals[i-k]
			}
		}

		i += (k - 1)
	}

	maxLspVal := 0
	for _, lspVal := range lspVals {
		if lspVal > maxLspVal {
			maxLspVal = lspVal
		}
	}

	return maxLspVal
}

func IsCircularPalindrome(word string) bool {
	if len(word) == 1 {
		return false
	}

	if isPalindrome(word) {
		return true
	}

	if longestPalindromSubstring(word) > 1 {
		return true
	}

	return false
}
