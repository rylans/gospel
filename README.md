# gospel
Spelling correction in Go

## Quick Start

#### Download and install

    go get github.com/rylans/gospel

#### Create a file `trygospel.go`
```go
package main

import "fmt"
import "github.com/rylans/gospel"

func main(){
  c := gospel.ForEnglish()
  fmt.Println(c.Correct("gospell"))
}
```

#### Build and run

    go build trygospel.go
    ./trygospel


## Requirements

This program relies on a word list located at "/usr/share/dict/words"


## License

gospel source code is licensed under Apache License, Version 2.0
