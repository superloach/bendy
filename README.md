archiving bcus my ink fork is gone and this was kinda garbage

# bendy
A Fantasy Console using the Ink scripting language.

## installing
```
go get github.com/superloach/bendy/cmd/bendy
```

## packaging a game
currently only single-file games are supported for packaging. requires [packr2](https://github.com/gobuffalo/packr/tree/master/v2).
```
cd $GOPATH/src/github.com/superloach/bendy
cp <game file> ./cmd/bendy/builtin.ink
packr2
go build ./cmd/bendy
```
