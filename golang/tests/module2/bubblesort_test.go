package module2_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestBubblesort(t *testing.T) {
	assert := assert.New(t)
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Test bubble sort", func(t *testing.T) {
		r, w := helpers.Replacer(`4
4 3 2 1
		`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Bubble_sort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`		3 4 2 1
3 2 4 1
3 2 1 4
2 3 1 4
2 1 3 4
1 2 3 4
`, string(out))
	})
}
