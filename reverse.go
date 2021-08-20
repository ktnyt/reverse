package reverse

// An implementation of Interface can be reversed by the routines in this
// package. The methods refer to elements of the underlying collection by
// integer index.
type Interface interface {
	Len() int
	Swap(i, j int)
}

// Reverse reverses data.
func Reverse(data Interface) {
	n := data.Len()
	for i := 0; i < n/2; i++ {
		data.Swap(i, n-i-1)
	}
}

// StringSlice attaches the methods of Interface to a []string.
type StringSlice []string

func (x StringSlice) Len() int      { return len(x) }
func (x StringSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Strings reverses a slice of strings.
func Strings(x []string) {
	Reverse(StringSlice(x))
}

// ByteSlice attaches the methods of Interface to a []byte.
type ByteSlice []byte

func (x ByteSlice) Len() int      { return len(x) }
func (x ByteSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Bytes reverses a slice of strings.
func Bytes(x []byte) {
	Reverse(ByteSlice(x))
}

// IntSlice attaches the methods of Interface to a []int.
type IntSlice []int

func (x IntSlice) Len() int      { return len(x) }
func (x IntSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Ints reverses a slice of ints.
func Ints(x []int) {
	Reverse(IntSlice(x))
}

// Float64Slice attaches the methods of Interface to a []float64.
type Float64Slice []float64

func (x Float64Slice) Len() int      { return len(x) }
func (x Float64Slice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Float64s reverses a slice of float64s.
func Float64s(x []float64) {
	Reverse(Float64Slice(x))
}
