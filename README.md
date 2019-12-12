# sortkeys-go

To install:

```
go get -u github.com/jeffreylo/sortkeys-go
```

```
λ sortkeys-go -h
Usage of sortkeys-go:
  -file string
    	Filename to be parsed
  -o string
    	Output filename
  -p string
    	Comma-separated ordered list of field names to prioritize
  -w	Write to -file
```

## Example

Given the following:

```go
type Example struct {
	// B is a comment for B.
	B string
	// A is a comment for A.
	A string
}
```

`sortkeys-go` rewrites to:

```go
type Example struct {
	// A is a comment for A.
	A string

	// B is a comment for B.
	B string
}
```

See [docs for more details](./docs).
