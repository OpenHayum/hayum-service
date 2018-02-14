package main

import (
	"fmt"

	"bitbucket.org/hayum/hayum-service/config"
)

func main() {
	fmt.Println("Listening: ", config.Port)
}
