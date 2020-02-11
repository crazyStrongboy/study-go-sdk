package string_match

import (
	"fmt"
	"testing"
)

func Test_bm_match(t *testing.T) {
	origin := "acbbcabcbc"
	sub := "abcbc"
	match := bm_match(origin, sub)
	fmt.Println("match: ", match)
}
