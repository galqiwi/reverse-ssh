package wizard

import (
	"fmt"
	"strconv"
)

func RequestValue(description, value string) (string, error) {
	fmt.Printf("Input %v", description)
	if value != "" {
		fmt.Printf(" (default is %v)", value)
	}
	fmt.Print("\n")

	fmt.Print(">>> ")

	output, err := ReadStdinLine()
	if output == "" {
		output = value
		fmt.Printf("\033[F>>> %v\n", value)
	}

	return output, err
}

func RequestIntValue(description string, value int) (int, error) {
	output, err := RequestValue(description, fmt.Sprint(value))
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(output)
}
