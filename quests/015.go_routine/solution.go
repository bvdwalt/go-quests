package goroutine

import "fmt"

func SendRequest(input string) string {
	responseChan := make(chan string)

	go Server(input, responseChan)

	return <-responseChan
}

func Server(request string, response chan string) {
	response <- fmt.Sprintf("processed: %s", request)
}
