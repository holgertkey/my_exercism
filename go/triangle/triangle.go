
package triangle

import "math"

type Kind int

const (
    NaT Kind = iota // not a triangle
    Equ             // equilateral
    Iso             // isosceles
    Sca             // scalene
)

func KindFromSides(a, b, c float64) Kind {
    if a <= 0 || b <= 0 || c <= 0 || a + b <= c || b + c <= a || c + a <= b {
        return NaT
    }

    const epsilon = 1e-9
    equal := func(x, y float64) bool {
        return math.Abs(x - y) < epsilon
    }

    switch {
    case equal(a, b) && equal(b, c):
        return Equ
    case equal(a, b) || equal(b, c) || equal(c, a):
        return Iso
    default:
        return Sca
    }
}
