package resistorcolor

import "slices"

func Colors() []string {
    colors := []string{
        "black",
        "brown",
        "red",
        "orange",
        "yellow",
        "green",
        "blue",
        "violet",
        "grey",
        "white",
    }
    return colors
}

func ColorCode(color string) int {
    colors := Colors()
    return slices.Index(colors, color)
}
