package string_match

import (
	"fmt"
	"testing"
)

func Test_bf_match(t *testing.T) {
	origin := "abcdfg"
	sub := "dg"
	match := bf_match(origin, sub)
	fmt.Println("match: ", match)
}
