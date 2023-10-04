package main

import (
	"fmt"
	"os"

	"github.com/andrewfrench/rand/pkg/rand"

	"github.com/spf13/cobra"
	"github.com/atotto/clipboard"
)

const clipboardMessage = "Copied value to clipboard."

var (
	opts             = rand.Options{}
	passwordMode     = false
	passwordModeOpts = rand.Options{
		Uppers:   true,
		Lowers:   true,
		Numbers:  true,
		Specials: true,
		Length:   32,
		ToClipboard: true,
	}
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const info = `A random string and number generation utility.

By default, rand will output an eight-character string containing characters
a-z0-9. If any of the source modifiers -l, -u, -n, or -s are specified, the
source will only include the character types requested.

The length of the output string can be specified using -c without altering the
character source.

In password mode (-p or --password), rand will generate a 32-character string
from all sources. This mode is equivalent to rand -lunsc32. Password mode 
will override all other modifiers.

Examples:
  $ rand
  %s

  $ rand -luns
  %s

  $ rand -nuc20
  %s

  $ rand -c3
  %s

  $ rand -p
  %s
`

var cmd = &cobra.Command{
	Use:   "rand",
	Short: "A random string and number generation utility.",
	Long:  makeInfoText(info),
	PreRun: func(_ *cobra.Command, _ []string) {
		if passwordMode {
			opts = passwordModeOpts

			return
		}

		if !opts.Specials && !opts.Uppers && !opts.Lowers && !opts.Numbers {
			opts.Lowers = true
			opts.Numbers = true
		}
	},
	Run: func(_ *cobra.Command, _ []string) {
		value := rand.Make(opts)

		if opts.ToClipboard {
			err := clipboard.WriteAll(value)
			if err != nil {
				fmt.Println("Failed writing to clipboard: %s", err.Error())

				os.Exit(1)
			}

			fmt.Println("Copied value to clipboard.")

			return 
		}

		fmt.Println(value)
	},
}

func init() {
	cmd.Flags().BoolVarP(&passwordMode, "password", "p", false, "generate a password and write to the clipboard")

	cmd.Flags().IntVarP(&opts.Length, "len", "c", 8, "sets the length, in characters, of the output")
	cmd.Flags().BoolVarP(&opts.Lowers, "lowers", "l", false, "include lowercase letters a-z")
	cmd.Flags().BoolVarP(&opts.Uppers, "uppers", "u", false, "include uppercase letters A-Z")
	cmd.Flags().BoolVarP(&opts.Numbers, "numbers", "n", false, "include numerals 0-9")
	cmd.Flags().BoolVarP(&opts.Specials, "specials", "s", false, "include special characters like !, @, and #")
	cmd.Flags().BoolVarP(&opts.ToClipboard, "clipboard", "b", false, "copy the output to the clipboard")
}

func makeInfoText(info string) string {
	noArgsExample := rand.Make(rand.Options{Length: 8, Numbers: true, Lowers: true})
	lunsExample := rand.Make(rand.Options{Length: 8, Numbers: true, Lowers: true, Specials: true, Uppers: true})
	nuc20Example := rand.Make(rand.Options{Length: 20, Numbers: true, Uppers: true})
	c3Example := rand.Make(rand.Options{Length: 3, Numbers: true, Lowers: true})

	return fmt.Sprintf(info, noArgsExample, lunsExample, nuc20Example, c3Example, clipboardMessage)
}
