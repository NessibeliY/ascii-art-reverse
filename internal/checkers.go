package internal

import (
	"errors"
	"hash/crc64"
	"os"
	"strings"
)

var errorUsage = errors.New("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")

type parsedArgs struct {
	input       string
	reverseFile string
}

func parseArgs() (*parsedArgs, error) {
	args := os.Args[1:]

	if len(args) != 1 {
		return nil, errors.New("please provide one string")
	}

	if strings.HasPrefix(args[0], "--reverse=") {
		fileName := strings.TrimPrefix(args[0], "--reverse=")
		if fileName == "" {
			return nil, errorUsage
		}
		return &parsedArgs{reverseFile: fileName}, nil
	}

	if strings.HasPrefix(args[0], "--reverse") {
		return nil, errorUsage
	}

	return &parsedArgs{input: args[0]}, nil
}

func ValidInput(input string) (bool, error) {
	for _, letter := range input {
		if letter > 127 {
			return false, errors.New("provide ascii chars only")
		}
	}

	return true, nil
}

func strToHash(bannerText []byte) uint64 {
	crc64Table := crc64.MakeTable(crc64.ECMA)
	hashedData := crc64.Checksum([]byte(bannerText), crc64Table)
	return hashedData
}
