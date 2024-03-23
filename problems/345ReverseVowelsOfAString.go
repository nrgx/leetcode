package leetcode

func reverseVowels(s string) string {
    bs := []byte(s)
    for l, r := 0, len(bs) - 1; l < r; l, r = l + 1, r - 1 {
        for l < r && !isVowel(bs[l]) {
            l++
        }
        for l < r && !isVowel(bs[r]) {
            r--
        }
        bs[l], bs[r] = bs[r], bs[l]
    }
    return string(bs)
}

func isVowel(b byte) bool {
    for _, v := range "AaEeIiOoUu" {
        if byte(v) == b {
            return true
        }
    }
    return false
}

