package leetcode

import "strings"

func reverseWords(s string) string {
    words := strings.Fields(s)
    for l, r := 0, len(words) - 1; l < r; l, r = l + 1, r - 1 {
        words[l], words[r] = words[r], words[l]
    }
    return strings.Join(words, " ")
}

