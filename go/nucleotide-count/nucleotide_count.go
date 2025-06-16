package dna

import "errors"

type Histogram map[rune]uint

type DNA string 

func (d DNA) Counts() (Histogram, error) {
    var h Histogram = map[rune]uint{
        'A': 0,
        'C': 0,
        'G': 0,
        'T': 0,
    }

    for _, char := range d {
        switch char {
        case 'A', 'C', 'G', 'T':
            h[char]++
        default:
            return nil, errors.New("Error: ivalid DNA")
        } 
    }

    return h, nil
}
