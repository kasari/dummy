# dummy
Set a dummy value recursively

### Example
```go
package main

import (
	"math/rand"
	"time"

	"github.com/k0kubun/pp"
	"github.com/kasari/dummy"
)

type hoge struct {
	ID   int
	Name string
	Fuga fuga
}

type fuga struct {
	Text string
	Map  map[string]int
	Poyo []poyo
}

type poyo struct {
	privateNum uint
	Text       string
}

func main() {
	h := &hoge{}
	dummy.Set(h)
	pp.Print(h)
}
```

#### Output
```go
&main.hoge{
  ID:   2959226547383548094,
  Name: "qZR",
  Fuga: main.fuga{
    Text: "nlvcUwMLPmvwpHsxb",
    Map:  {
      "cJHxGnVvTmPEJqLw": 5893965662154205921,
    },
    Poyo: []main.poyo{
      main.poyo{
        privateNum: 0x8e8db389d89452ba,
        Text:       "t",
      },
      main.poyo{
        privateNum: 0xf15723e657da6a4d,
        Text:       "lBZFNvHzwwHp",
      },
    },
  },
}
```
