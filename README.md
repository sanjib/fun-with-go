# fun-with-go

A personal journey of fun and experiments with the Go language.

Following are some of the functions, variables and commands that
I am personally using along my journey with Go.

## functions and vars

- len(v Type) int
- bufio.NewScanner(r io.Reader) \*Scanner
- bufio: (s \*Scanner).Scan() bool
- bufio: (s \*Scanner).Text() string
- fmt.Println(a ...interface{}) (n int, err error)
- fmt.Sprintf(format string, a ...interface{}) string
- strings.Join(elems []string, sep string) string
- os.Args []string
- os.Stdin NewFile(uintptr(syscall.Stdin), "/dev/stdin")
- runtime.NumCPU() int
- time.Now() Time
- time.Since(t Time) Duration

## keywords

- break
- case
- const // example: const World = "世界"
- default
- defer
- else
- for
- func
- go
- if
- import
- interface
- map
- package
- range
- return
- struct
- switch
- type // example: switch v := i.(type)
- var

**_keywords pool_**

- select
- chan
- goto
- fallthrough
- continue

## tools

- go [command] [arguments]
- go build [packages]
- go doc [pkg] [methodOrField]
- go fmt [filename]
- go get [packages]
- go help [keyword]
- go run [filename]
- go test -bench=. [filename]
- go version

## definitions

- pointers
  - a pointer holds the memory address of a value
  - the & operator generates a pointer to its operand
  - the \* operator denotes a pointer's underlying value

## some common conversion characters in format strings

|            |                      |
| ---------- | -------------------- |
| %d         | int in decimal       |
| %x, %o, %b | int in hex, oct, bin |
| %f, %g, %e | float                |
| %t         | bool                 |
| %c         | rune                 |
| %s         | string               |
| %q         | quoted string, rune  |
| %v         | any value            |
| %T         | type                 |
| %%         | literal percent      |

## helpful custom programs

- [generate random number with crypto/rand based seed](./tour/2.basics/1.pkg_var_func/1.packages/packages.go)

## nuances

- [switch - see switch condition](./tour/2.basics/2.flow_ctrl/9.switch/switch.go)
- [switch evaluation order - see evaluation in case](./tour/2.basics/2.flow_ctrl/10.switch_eval_order/switch_eval_order.go)
- [switch with no condition - see no condition](./tour/2.basics/2.flow_ctrl/11.switch_with_no_cond/switch_with_no_cond.go)
- [pointers - unerstand pointer, &, \*](./tour/2.basics/3.more_types/1.pointers/pointers.go)
- [struct pointers - see struct field change via pointer](./tour/2.basics/3.more_types/4.struct_pointers/struct_pointers.go)
- a slice consists of a pointer to the array, the length of the segment, and its capacity
- [slice length & capacity - see how underlying array capacity changes when array elements are removed from front of array](./tour/2.basics/3.more_types/11.slice_len_cap/slice_len_cap.go)
- [create slice with make - reinforce understanding of how removing array elements from beginning affect underlying array capacity](./tour/2.basics/3.more_types/13.create_slice_make/create_slice_make.go)
- [pointer receivers](./tour/3.methods_and_interfaces/4.methods_pointers/methods_pointers.go)
- [pointers & functions](./tour/3.methods_and_interfaces/5.pointers_and_functions/pointers_and_functions.go)
- [methods and pointer indirection - methods with pointer receivers take either a value (as a convenience) or a pointer](./tour/3.methods_and_interfaces/6.methods_and_pointer_indirection/methods_and_pointer_indirection.go)
- therefore method pointer receivers can modify the received value even if it was passed as a value instead of a pointer
- all methods on a given type should have either value or pointer receivers, but not a mixture of both
- a value of interface type can hold any value that implements those methods
- an interface value holds a value of a specific underlying concrete type
- if the concrete value inside the interface itself is nil, the method will be called with a nil receiver
- calling a method on a nil interface is a run-time error - there has to be some type with a concrete method assigned to the interface value
- an empty interface may hold values of any type
- type assertion provides access to an interface value's underlying concrete value
- type assertion syntax is similar to that of reading from a map
- the error type is a built-in interface similar to fmt.Stringer
- as with fmt.Stringer, the fmt package looks for the error interface when printing values

```go
type Stringer interface {
    String() string
}
type error interface {
    Error() string
}
```

- there are many implementations of the io.Reader interface, including files, network connections, compressors, ciphers, and others

## exercise solutions

- [custom Sqrt function](./tour/2.basics/2.flow_ctrl/8.exe_loops_funcs/exe_loops_funcs.go)
- [implement Pic](./tour/2.basics/3.more_types/18.exe_slices/exe_slices.go)
- [implement WordCount](./tour/2.basics/3.more_types/23.exe_maps/exe_maps.go)
- [implement Fibonacci sequence using closure](./tour/2.basics/3.more_types/26.exe_fib_closure/exe_fib_closure.go)
- [implement stringer](./tour/3.methods_and_interfaces/18.exe_stringers/exe_stringers.go)
- [errors](./tour/3.methods_and_interfaces/20.exe_errors/exe_errors.go)
- [errors](./tour/3.methods_and_interfaces/20.exe_errors/exe_errors.go)
- [implement Read that emits 'A'](./tour/3.methods_and_interfaces/22.exe_readers/exe_readers.go)
- [implement rot13Reader](./tour/3.methods_and_interfaces/23.exe_rot13_reader/exe_rot13_reader.go)
- [implement Image](./tour/3.methods_and_interfaces/25.exe_images/exe_images.go)

## ref blog articles

- [Go Slices: usage and internals](https://blog.golang.org/slices-intro)
