This folder contains code from Xah Lee's Golang Tutorial available at: http://xahlee.info/golang/golang_index.html

len(string)

```text
fmt.Println("a", len("a"))
fmt.Println("Ø", len("Ø"))
fmt.Println("♥", len("\u2665"))
fmt.Println("😂", len("\U0001f602"))

// output:
a 1
Ø 2
♥ 3
😂 4
```