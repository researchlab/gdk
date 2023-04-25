package gdk

type (
	// Signed is a constraint that permits any signed integer type.
	// If future releases of Go add new predeclared signed integer types,
	// this constraint will be modified to include them.
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}

	// Unsigned is a constraint that permits any unsigned integer type.
	// If future releases of Go add new predeclared unsigned integer types,
	// this constraint will be modified to include them.
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	// Integer is a constraint that permits any integer type.
	// If future releases of Go add new predeclared integer types,
	// this constraint will be modified to include them.
	Integer interface {
		Signed | Unsigned
	}

	// Float is a constraint that permits any floating-point type.
	// If future releases of Go add new predeclared floating-point types,
	// this constraint will be modified to include them.
	Float interface {
		~float32 | ~float64
	}

	// Complex is a constraint that permits any complex numeric type.
	// If future releases of Go add new predeclared complex numeric types,
	// this constraint will be modified to include them.
	Complex interface {
		~complex64 | ~complex128
	}

	// Ordered is a constraint that permits any ordered type: any type
	// that supports the operators < <= >= >.
	// If future releases of Go add new ordered types,
	// this constraint will be modified to include them.
	Ordered interface {
		Integer | Float | ~string
	}

	Null struct{}

	Call func()

	// function provides one input argument and one return
	Func[R, T any] func(T) R

	// TFunc provides one input and two return
	TFunc[E, R, T any] func(E) (R, T)

	// two-arity specialization of function
	BiFunc[R, T, U any] func(T, U) R

	// function provides one input argument and no returns
	Consumer[T any] func(T)

	// function provides two input arguments and no returns
	BiConsumer[T, U any] func(T, U)

	// function provides one input and one return
	Supplier[R any] func() R

	// Evaluate use the specified parameter to perform a test that returns true or false
	Evaluate[E any] Func[bool, E]

	// compare function
	CMP[E any] BiFunc[int, E, E]

	// equal function
	EQL[E any] BiFunc[bool, E, E]

	// Comparable is a interface to compare action
	Comparable[E any] interface {
		CompareTo(v E) int
	}
)

var (
	Empty Null // const var for nil usage marker
)
