import "maps"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sMap := make(map[rune]int)
    tMap := make(map[rune]int)

    for _, sChar := range s {
        sMap[sChar]++
    } 

    for _, tChar := range t {
        tMap[tChar]++
    } 

    return maps.Equal(sMap, tMap)
}