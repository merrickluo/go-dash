# Go-Dash

## Development

To [Working with generic code](https://github.com/golang/tools/blob/master/gopls/doc/advanced.md#working-with-generic-code):

```bash
go get -v golang.org/dl/gotip
gotip download
gotip get golang.org/x/tools/gopls@master golang.org/x/tools@master
```

Then run test:

```bash
gotip test -v ./...
```
