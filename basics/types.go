bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.


Type casting:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)



i := 42
f := float64(i)
u := uint(f)