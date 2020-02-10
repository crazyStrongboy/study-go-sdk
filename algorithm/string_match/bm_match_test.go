package string_match

import (
	"fmt"
	"testing"
)

func Test_bm_match(t *testing.T) {
	origin := "abcdfg"
	sub := "dg"
	match := bm_match(origin, sub)
	fmt.Println("match: ", match)
}
