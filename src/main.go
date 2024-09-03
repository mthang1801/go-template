package main

import "fmt"

type EmailStatus int

const (
	emailBounced EmailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)

func main() {
	fmt.Println(emailBounced, emailInvalid, emailOpened)
}
