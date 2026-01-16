package channel

import "fmt"

func ReadFromBoth(ch1 chan string, ch2 chan string) string {
	val1 := <-ch1
	val2 := <-ch2

	return fmt.Sprintf("read: %s & %s", val1, val2)
}

func WriteToBoth(ch1 chan string, ch2 chan string, msg string) {
	go func() {
		formatted := "write: " + msg
		ch1 <- formatted
		ch2 <- formatted
	}()
}

func ReadThenWrite(input chan string, output chan string) {
	val := <-input
	output <- fmt.Sprintf("transform: %s", val)
}
