// main.go
package main

import (
	"time"
)

func main() {
	input := "1 2 2 3 4 5 6"
	output := "1 2 2 3 4 5 7\nabv cd\nfasdf"
	report := NewReport("1", 10*time.Second, input, input, output)
	report.Print()
}
