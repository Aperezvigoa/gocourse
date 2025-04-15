package intermediate

import (
	"errors"
	"fmt"
)

func main() {
	if doSomethingError := doSomething(); doSomethingError != nil {
		fmt.Println(doSomethingError)
		// return
	}

	if processSomethingError := processSomething(); processSomethingError != nil {
		fmt.Println(processSomethingError)
	}
}

type customError struct {
	code    int
	message string
	er      error
}

// Error returns the error message. Implementing Error() method of error interface
func (c *customError) Error() string {
	return fmt.Sprintf("Error %d: %s %v", c.code, c.message, c.er)
}

// Functions that returns a custom error
func doSomething() error {
	return &customError{
		code:    500,
		message: "Something went wrong!",
	}
}

func processSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    404,
			message: "Not found",
			er:      err,
		}
	} else {
		return nil
	}
}

func doSomethingElse() error {
	return errors.New("internal error")
}
