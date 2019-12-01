package rand

import (
	"math/rand"
	"time"
)

type Options struct {
	Lowers   bool
	Uppers   bool
	Numbers  bool
	Specials bool
	Length   int
}

func Make(opts Options) string {
	source := make([]rune, 0)
	generated := make([]rune, 0)
	rand.Seed(time.Now().UnixNano())
	compileSource(opts, &source)
	for i := 0; i < opts.Length; i++ {
		generated = append(generated, source[rand.Intn(len(source))])
	}

	return string(generated)
}

func compileSource(opts Options, dest *[]rune) {
	if opts.Lowers {
		*dest = append(*dest, []rune(lowers)...)
	}

	if opts.Uppers {
		*dest = append(*dest, []rune(uppers)...)
	}

	if opts.Numbers {
		*dest = append(*dest, []rune(numbers)...)
	}

	if opts.Specials {
		*dest = append(*dest, []rune(specials)...)
	}
}
