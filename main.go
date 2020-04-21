package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		// inputs provided by Github Actions runtime
		// we should define them in action.yml
		status = os.Getenv("INPUT_STATUS")

		// github environment context
		workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
	)

	fmt.Println("Start printing stuff...")
	fmt.Println("[input] status %v", status)
	fmt.Println("[context] workflow %v", workflow)
	fmt.Println("[context] repo %v", repo)
}
