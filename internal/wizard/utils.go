package wizard

import "fmt"

func RequestConfigValue(description string, value *string) error {
	var err error
	*value, err = RequestValue(description, *value)
	return err
}

func RequestConfigIntValue(description string, value *int) error {
	var err error
	*value, err = RequestIntValue(description, *value)
	return err
}

func PrintSplitter() {
	fmt.Print("\n------------------------------------------\n\n")
}
