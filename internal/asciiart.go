package internal

import (
	"errors"
	"fmt"
	"os"
)

const height = 8

func Run() error {
	option, err := parseArgs()
	if err != nil {
		return err
	}

	if option.reverseFile != "" {
		content, err := os.ReadFile(option.reverseFile)
		if err != nil {
			return err
		}

		result, err := Reverse(string(content), "standard.txt")
		if err != nil {
			return err
		}

		fmt.Print(result)
		return nil
	}

	if option.input == "" {
		return nil
	}

	ok, err := ValidInput(option.input)
	if !ok {
		return err
	}

	result, err := Convert(option.input, "standard.txt")
	if err != nil {
		return err
	}

	gotWidth, err := getWidth()
	if err != nil {
		return errors.New("there is no terminal to get width")
	}

	if width(result) > int(gotWidth) {
		return errors.New("please provide shorter text")
	}

	fmt.Print(result)

	return nil
}
