package p
type f struct {
Foo int
Foo2 int // trailing foo comment
// Qux is an int.
// With a multi-line comment.
Qux int

Gux int // trailing comment

// Baz is a string.
Baz string

// ID is the ID field of the struct.
ID int64 // trailing comment

lux string

inline struct {
// Qux is an int.
// With a multi-line comment.
Qux int
// Baz is a string.
Baz string
}
inlineInterface interface {

GetMethod(foo, bar int) error
GetMethod2(foo, bar int) error // ok

// GetMethod3 does a thing
GetMethod3(foo, bar int) error
}
}

type g interface {

GetMethod(foo, bar int) error
GetMethod2(foo, bar int) error // ok

// GetMethod3 does a thing
GetMethod3(foo, bar int) error
}
