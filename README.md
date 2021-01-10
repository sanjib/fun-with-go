# fun-with-go
A personal journey of fun and experiments with the Go language

Following are some of the functions, variables and commands that 
I am personally using along my journey with Go.

## functions and vars

- len(v Type) int
- bufio.NewScanner(r io.Reader) *Scanner
- bufio: (s *Scanner).Scan() bool
- bufio: (s *Scanner).Text() string
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

***keywords pool***

- break		default		interface	select
- case		defer		go		struct
- chan		else		goto		switch
- const		fallthrough	if		type
- continue	return       

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
