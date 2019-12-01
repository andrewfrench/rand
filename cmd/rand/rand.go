package main

import (
	"fmt"
	"os"

	"github.com/andrewfrench/rand/pkg/rand"

	"github.com/spf13/cobra"
)

var (
	opts = rand.Options{}
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const randInfo = `A random string and number generation utility.

By default, rand will output an eight-character string containing characters
a-z0-9. If any of the source modifiers -l, -u, -n, or -s are specified, the
source will only include the character types requested.

The length of the output string can be specified using -c without altering the
character source.

Examples:
  $ rand
  rl6gtll2

  $ rand -luns
  3E0*6hp^

  $ rand -nuc 20
  15NTQV4ASIN1UEQF0MXY
`

var cmd = &cobra.Command{
	Use:   "rand",
	Short: "A random string and number generation utility.",
	Long:  randInfo,
	PreRun: func(_ *cobra.Command, _ []string) {
		if !opts.Special && !opts.Uppers && !opts.Lowers && !opts.Numbers {
			opts.Lowers = true
			opts.Numbers = true
		}
	},
	Run: func(_ *cobra.Command, _ []string) {
		generated := rand.Make(opts)
		fmt.Println(generated)
	},
}

func init() {
	cmd.Flags().IntVarP(&opts.Length, "len", "c", 8, "sets the length, in characters, of the output")
	cmd.Flags().BoolVarP(&opts.Lowers, "lowers", "l", false, "include lowercase letters a-z")
	cmd.Flags().BoolVarP(&opts.Uppers, "uppers", "u", false, "include uppercase letters A-Z")
	cmd.Flags().BoolVarP(&opts.Numbers, "numbers", "n", false, "include numerals 0-9")
	cmd.Flags().BoolVarP(&opts.Special, "special", "s", false, "include special characters like !, @, and #")
}
