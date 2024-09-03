package generate

import (
	"fmt"
	"testing"
)

func Test_generate(t *testing.T) {
	triangle := generate(5)
	for _, row := range triangle {
		fmt.Println(row)
	}
}
