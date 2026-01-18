package main

import (
	"fmt"

	"github.com/samma9/samRepos/tree/main/modules/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
