[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/tasmota)](https://goreportcard.com/report/github.com/mskrha/tasmota)

## tasmota

### Description
Go library for communication with IoT devices using the Tasmota firmware.

### Installation
`go get github.com/mskrha/tasmota`

### Example usage
```go
package main

import (
	"fmt"
	"os"

	"github.com/mskrha/tasmota"
)

func main() {
	t, err := tasmota.New(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	e, err := t.GetEnergy()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", e)
}
```
