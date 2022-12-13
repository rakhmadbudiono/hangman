package core

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var words = []string{
	"GUM",
	"SIN",
	"FOR",
	"CRY",
	"LUG",
	"BYE",
	"FLY",
	"UGLY",
	"EACH",
	"FROM",
	"WORK",
	"TALK",
	"WITH",
	"SELF",
	"PIZZA",
	"THING",
	"FEIGN",
	"FIEND",
	"ELBOW",
	"FAULT",
	"DIRTY",
	"BUDGET",
	"SPIRIT",
	"QUAINT",
	"MAIDEN",
	"ESCORT",
	"PICKAX",
	"EXAMPLE",
	"TENSION",
	"QUININE",
	"KIDNEY",
	"REPLICA",
	"SLEEPER",
	"TRIANGLE",
	"KANGAROO",
	"MAHOGANY",
	"SERGEANT",
	"SEQUENCE",
	"MOUSTACHE",
	"DANGEROUS",
	"SCIENTIST",
	"DIFFERENT",
	"QUIESCENT",
	"MAGISTRATE",
	"ERRONEOUSLY",
	"LOUDSPEAKER",
	"PHYTOTOXIC",
	"MATRIMONIAL",
	"PARASYMPATHOMIMETIC",
	"THIGMOTROPISM",
}

func shuffleWords() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := range words {
		j := rand.Intn(len(words))
		words[i], words[j] = words[j], words[i]
	}
}

func sanitizeInput(input string) string {
	return strings.ToUpper(string(input[0]))
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
