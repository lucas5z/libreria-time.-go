package timecomplemeto

import (
	"fmt"
	"time"
)

func Varibles() {
	//tiempo
	second := time.Second
	fmt.Println("1-", int64(second/time.Millisecond))
	//tiempo opracion
	seconds := 10
	fmt.Println("2-", time.Duration(seconds)*time.Second) // prints 10s
}
