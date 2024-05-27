// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

// MySleep блокирует выполнение на указанное количество миллисекунд.
func MySleep(milliseconds int) {
	<-time.After(time.Duration(milliseconds) * time.Millisecond)
}

func main() {
	fmt.Println("Sleeping for 2 seconds...")
	MySleep(2000)
	fmt.Println("Awake!")
}
