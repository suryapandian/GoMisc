func minDeletions(s string) int {
    charMap := make(map[rune]int)
    for _, r := range s {
        charMap[r] += 1
    }

    var result int
    numMap := make(map[int]rune)
    for k, v := range charMap {
        for i := v; i >= 0; i-- {
            if _, ok := numMap[i]; (i == 0) || !ok {
                numMap[i] = k
                break
            }
            result += 1
        }
    }
    return result
}