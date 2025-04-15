package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math error: square root of negative number")
	}
	return 1, nil
}

func process[T any](data []T) error {
	if len(data) == 0 {
		return errors.New("error: data length is zero")
	}
	return nil
}

func main() {
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", result)
	}

	result1, err1 := sqrt(-6)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Result:", result1)
	}

	datapack := make([]int, 0)
	if err := process(datapack); err != nil {
		fmt.Println(err)
	}

	if err1 := eprocess(); err1 != nil {
		fmt.Println(err1)
	}

	if err2 := readData(); err != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Data readed sucessfully")
	}
}

type customError struct {
	message error
}

func (cE *customError) Error() string {
	return fmt.Sprintf("Error: %s\n", cE.message)
}

func eprocess() error {
	return &customError{message: errors.New("unable to process it")}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	} else {
		return nil
	}
}

func readConfig() error {
	return errors.New("config error")
}
