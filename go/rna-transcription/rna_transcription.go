package strand

var dnaToRna = map[rune]string {
    'G': "C",
    'C': "G",
    'T': "A",
    'A': "U",
}

func ToRNA(dna string) string {
    var rna string
    for _, rune := range dna {
        rna += dnaToRna[rune]
    }
    return rna
}
