# sortkeys-go

Given the following:

```go
type Example struct {
	// B is a comment for B.
	B string
	// A is a comment for A.
	A string
}
```

`sortkeys-go` rewrites the struct to :

```go
type Example struct {
	// A is a comment for A.
	A string

	// B is a comment for B.
	B string
}
```
