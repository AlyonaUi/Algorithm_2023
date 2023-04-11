package module2_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestStorage(t *testing.T) {
	assert := assert.New(t)
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Test storage", func(t *testing.T) {
		r, w := helpers.Replacer(`5
		1 50 3 4 3
		16
		1 2 3 4 5 1 3 3 4 5 5 5 5 5 4 5		
		`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Storage()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`		yes
		no
		no
		no
		yes		
`, string(out))
	})
}
