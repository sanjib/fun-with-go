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

- for
- func
- import
- map
- package
- range
- var

**_keywords pool_**

- break default interface select
- case defer go struct
- chan else goto switch
- const fallthrough if type
- continue return

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
- [slice length & capacity - see how underlying array capacity changes when array elements are removed from front of array](./tour/2.basics/3.more_types/11.slice_len_cap/slice_len_cap.go)
- [create slice with make - reinforce understanding of how removing array elements from beginning affect underlying array capacity](./tour/2.basics/3.more_types/13.create_slice_make/create_slice_make.go)

## exercise solutions

- [custom Sqrt function](./tour/2.basics/2.flow_ctrl/8.exe_loops_funcs/exe_loops_funcs.go)
- [implement Pic](./tour/2.basics/3.more_types/18.exe_slices/exe_slices.go)
- [implement WordCount](./tour/2.basics/3.more_types/23.exe_maps/exe_maps.go)
- [implement Fibonacci sequence using closure](./tour/2.basics/3.more_types/26.exe_fib_closure/exe_fib_closure.go)

## ref blog articles

- [Go Slices: usage and internals](https://blog.golang.org/slices-intro)
