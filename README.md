# go error handling

## Result

```bash
$ go run main.go

== Pattern1: Creates on initialize ==
Hmm, he seems to need a cake

== Pattern 2: Creates ad-hok ==
Func2() got error: error - i need cake not banana

== Pattern 3: Define custom error object ==
Hmm, he seems to need cake not strawberry

== Pattern 4: Use Error interface not concrete struct ==
let me give him your cake
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xb code=0x1 addr=0x0 pc=0x309d]

goroutine 1 [running]:
panic(0xdc380, 0xc82000a1c0)
    /usr/local/Cellar/go/1.6.2/libexec/src/runtime/panic.go:481 +0x3e6
main.main()
    /Users/yuichi.murata/go/src/github.com/yuichi1004/go-error-handling/main.go:91 +0xaad
    exit status 2
```

