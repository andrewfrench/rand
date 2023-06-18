# rand

A random string and number generation utility.

By default, `rand` will output an eight-character string containing characters `a-z0-9`. If any of the source modifiers `-l`, `-u`, `-n`, or `-s` are specified, the source will only include the character types requested.

The length of the output string can be specified using `-c` without altering the character source.

In password mode (`-p` or `--password`), `rand` will generate a 32-character string from all sources. This mode is equivalent to `rand -lunsc32`. Password mode will override all other modifiers.

### Usage

`rand [flags]`

#### Examples

```
$ rand
6agqlqrl
```

```
$ rand -p
p8wCbF+V,bFDeG?h7EhomAsfpaU<ed#9
```

```
$ rand -luns
F##iPSJt
```

```
$ rand -nuc20
Q1XH3E6EQ4A0EHEKUVDZ
```

```
$ rand -c3
w1p
```

#### Flags

`-h`, `--help`       help for rand

`-c`, `--len int`    sets the length, in characters, of the output (default 8)

`-l`, `--lowers`     include lowercase letters a-z

`-n`, `--numbers`    include numerals 0-9

`-p`, `--password`   generate a password string

`-s`, `--specials`   include special characters like !, @, and #

`-u`, `--uppers`     include uppercase letters A-Z

### Build

```sh
make build
```

### Install

If the repository is cloned locally:

```sh
make install
```

