### What are, according to you, the three most underrated language features and why? 

Some amazing features of Go include generation of binaries, easy to use concurrent routines, tool chains, reliable garbage collectors, support for embedding C code, quick learning curve, amazing support by Gophers community, simplicity and many more.

However, some features that according to me are underrated include:
 - Go's wide range of built in standard packages
 - strict compile time checks
 - built in testing and profiling
 - defer statements.

The strict compilation checking, prevents us from mistakenly defining or declaring unused variables, including unused packages etc... making Go clean and light.
Built in testing tools and wide range of standard packages make Go easy to learn and use.
Defer statements are widely used. They are very handy to perform clean up operations after  function execution. 

### How can Go still shoot you in the foot?

One main feature of Go that excited me when I started is the ease and safe way by which I could spin up multiple go routines. However, if we get carried away by the power of go routines and overlook the machine capabilities, it could lead to fatal problems. 

For instance, spinning multiple go routines in a web server, each making database calls, could exhaust the available database connections. Thereby forcing the web server to go down by not accepting new requests. Like wise, in-memory usage could shoot if the program is not written considering the machine specifications while spinning go routines.

Though go routines, channels and wait groups are easy to use, we should also be conscious of the infrastructure in hand.

### In your opinion, which cases justify using Unsafe? Explain your points

While converting different structs with the attributes having same datatypes, it is very convenient to use `unsafe` package. Casting is not possible in such cases and depending on the number of attributes in the struct definitions, conventional conversion could be lengthy.

Example: Many of the struct definitions in `unix` and other `sys` packages have many attributes in them. Conversion of such structs is hassle-free through `unsafe`. 

`sync` and `reflect` are very well known and widely used packages that use `unsafe`.

Sometimes, system calls require us to pass pointer to some memory to the kernel to perform tasks. Here we could pass `unsafe.Pointer` and later convert it to `uintptr` which could be used with functions in `syscall` package. 