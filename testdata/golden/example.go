package p

type f struct {
	// ID is the ID field of the struct.
	ID int64 // trailing comment

	// Baz is a string.
	Baz string

	Foo  int
	Foo2 int // trailing foo comment
	Gux  int // trailing comment

	// Qux is an int.
	// With a multi-line comment.
	Qux int

	inline struct {
		// Baz is a string.
		Baz string

		// Qux is an int.
		// With a multi-line comment.
		Qux int
	}
	inlineInterface interface {
		GetMethod(foo, bar int) error
		GetMethod2(foo, bar int) error // ok

		// GetMethod3 does a thing
		GetMethod3(foo, bar int) error
	}
	lux string
}

type g interface {
	GetMethod(foo, bar int) error
	GetMethod2(foo, bar int) error // ok

	// GetMethod3 does a thing
	GetMethod3(foo, bar int) error
}
