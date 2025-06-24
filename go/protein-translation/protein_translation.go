package protein

import (
    "errors"
)

var (
    ErrStop         = errors.New("stop codon")
    ErrInvalidBase  = errors.New("invalid codon")
)

var codonMap = map[string]string{
    "AUG": "Methionine",
    "UUU": "Phenylalanine", "UUC": "Phenylalanine",
    "UUA": "Leucine", "UUG": "Leucine",
    "UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
    "UAU": "Tyrosine", "UAC": "Tyrosine",
    "UGU": "Cysteine", "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromCodon(codon string) (string, error) {
    protein, found := codonMap[codon]
    if !found {
        return "", ErrInvalidBase
    }
    if protein == "STOP" {
        return "", ErrStop
    }
    return protein, nil
}

func FromRNA(rna string) ([]string, error) {
    var proteins []string

    for i := 0; i < len(rna); i += 3 {
        if i + 3 > len(rna) {
            break
        }
        codon := rna[i : i+3]
        protein, err := FromCodon(codon)
        if err == ErrStop {
            break
        }
        if err != nil {
            return nil, err
        }
        proteins = append(proteins, protein)
    }
    
    return proteins, nil
}
