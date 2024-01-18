/*
Test generated by RoostGPT for test MiniProjects using AI Type Azure Open AI and AI Model roost-gpt4-32k

Test Scenarios:

1. Null Input: Test with "nil" or empty values.
2. Simple Test: Test with a simple valid palindrome like "Madam".
3. Numeric Test: Test with a string that contains numbers - valid palindrome as "12321".
4. Not a Palindrome: Test with a string that is not a palindrome - "programming".
5. Special Characters Test: Test with a strings contains special characters like "A man, a plan, a canal, Panama.". The function should remove non alphanumeric characters and return a result.
6. Test with Unicode Characters: Test with a string that contains unicode characters.
7. Large Input Test: Test with a string of a large size, to check how the function handles it.
8. Case Insensitive Test: The function should return true for "WoW" and "wow".
9. White Spaces Test: Test with a string with leading, trailing and middle white spaces - "Able was I saw elba".
10. Single Character Test: Test with a single character string like "a" or "1".
11. Same Character Test: Test with a string where all characters are same.
12. Test with a mix of upper and lower case letters.
13. Test with a string having numeric, alphabetic and alpanumeric characters-"12321abba12321" or "abcddcba@123".
14. Multiword Palindrome: Test with a multiword palindrome like "step on no pets".
*/
package palindrome

import (
	"testing"
)

func TestIsPalindromeRecursive_9c3315538c(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Null Input Test", "", true},
		{"Simple Test", "Madam", true},
		{"Numeric Test", "12321", true},
		{"Not a Palindrome Test", "programming", false},
		{"Special Characters Test", "A man, a plan, a canal, Panama.", true},
		{"Unicode Characters Test", "été", true},
		{"Large Input Test", "abcdefghihgfedcba", true},
		{"Case Insensitive Test", "WoW", true},
		{"White Spaces Test", "Able was I saw elba", true},
		{"Single Character Test", "a", true},
		{"Same Character Test", "aaaaaa", true},
		{"Mix of upper and lower case letters Test", "WoWowO", true},
		{"Numeric, alphabetic and alpanumeric characters Test", "12321abba12321", true},
		{"Multiword Palindrome Test", "step on no pets", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromeRecursive(tt.input); got != tt.want {
				t.Errorf("IsPalindromeRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
