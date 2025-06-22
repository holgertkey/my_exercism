package resistorcolorduo

var colorCode = map[string]int{
    "black":  0,
    "brown":  1,
    "red":    2,
    "orange": 3,
    "yellow": 4,
    "green":  5,
    "blue":   6,
    "violet": 7,
    "grey":   8,
    "white":  9,
}

func Value(colors []string) int {
    if len(colors) < 2 {
        return 0
    }
    return colorCode[colors[0]] * 10 + colorCode[colors[1]]
}