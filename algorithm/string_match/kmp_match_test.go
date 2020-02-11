package string_match

import (
	"fmt"
	"testing"
)

func Test_kmp_match(t *testing.T) {
	origin := "acbbcabcbc"
	sub := "abacdababc"
	match := kmp_match(origin, sub)
	fmt.Println("match: ", match)
}
