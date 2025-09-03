package palindrome

// go through every possible substring and check if they're a palindrome

// go through every possible substring:
// - func subStrings(s string) []string (returns every possible substring)
// check if each substring == palindrome:
// - func isPalindrome(s string) bool (returns reverse(s) == string)

// func PalindromeSubStrings(s string) []string {
// palindromes = []
// substrings = subString(s)
// loop through substrings
//    if isPalindrome(substring) == true
//       palindromes += substring
//   else
//       continue loop
// return palindromes
// }

func substrings(s string) []string {
	// input: "abadd"
	// output:
	//    ["a", "ab", "aba", "abad", "abadd",
	//     "b", "ba", "bad", "badd",
	//     "a", "ad", "add",
	//     "d", "dd",
	//     "d"]

	// go through s, for each letter of s go through all the substrings that start with that letter

	// substrings = []
	// loop through s
	//    for s[i]:
	//      prevSubString = s[i]
	//      loop through s[i+1] --> len(s)-1
	//         for s[j]:
	//              newSubstring = (prevSubstring+s[j])
	//              substrings = substrings + newSubstring
	//              prevSubstring = newSubstring

	var substrings []string

	for i := 0; i < len(s); i++ {
		currentLetter := (string(s[i]))
		substrings = append(substrings, currentLetter)
		prevSubstring := currentLetter

		for j := i + 1; j < len(s); j++ {
			newSubstring := prevSubstring+currentLetter
			substrings = append(substrings, newSubstring)
			prevSubstring = newSubstring
		}
	}

	return substrings
}

func isPalindrome(s string) bool {
	reversed := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	return reversed == s
}

func PalindromeSubStrings(s string) []string {
	palindromes := []string{}

	for _, substring := range substrings(s) {
		if isPalindrome(substring) {
			palindromes = append(palindromes, substring)
		}
	}

	return palindromes
}
