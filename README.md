# Directory Trimmer

## Installation

```sh
make build
cp out/dirtrim <somewhere_in_your_path>
```

## Usage

```sh
$ echo "foo/bar/baz" | dirtrim
f/b/baz

$ echo $HOME | dirtrim
~

$ echo $HOME/foo/bar | dirtrim
~/f/b/baz
```

## Bash Prompt

```sh
export PS1="\u@\h \$(pwd | dirtrim) $ "
```
