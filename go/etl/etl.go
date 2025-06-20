package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    out := map[string]int{}
    for point, arc := range in {
        for _, char := range arc {
            out[strings.ToLower(char)] = point 
        }
    }
    return out
}
