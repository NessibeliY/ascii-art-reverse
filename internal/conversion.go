package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Reverse(input, banner string) (string, error) {
	if input == "" {
		return "", nil
	}

	lines := strings.Split(input, "\n")
	filteredLines := make([]string, 0, len(lines))

	for _, line := range lines {
		if strings.HasSuffix(line, "$") {
			filteredLines = append(filteredLines, strings.TrimSuffix(line, "$"))
		}
	}

	if len(filteredLines) < height {
		return "", errors.New("invalid ascii art input")
	}

	alphabet, err := getAlphab(banner)
	if err != nil {
		return "", err
	}

	reverseAlphabet := make(map[string]rune)
	for char, art := range alphabet {
		reverseAlphabet[strings.Join(art, "\n")] = char
	}

	var result strings.Builder
	bigLines := make([][]string, 0, len(filteredLines))
	for i := 0; i < len(filteredLines); {
		if filteredLines[i] == "" {
			bigLines = append(bigLines, []string{})
			i++
			continue
		}

		if i+height > len(filteredLines) {
			return "", errors.New("invalid ascii art input: incomplete character block")
		}

		bigLines = append(bigLines, filteredLines[i:i+height])
		i += height
	}

	for _, bigLine := range bigLines {
		if len(bigLine) == 0 {
			result.WriteRune('\n')
			continue
		}

		maxWidth := len(bigLine[0])
		charStart := 0

		for charStart < maxWidth {
			oneChar := make([]string, height)
			charFound := false

			for charEnd := charStart + 1; charEnd <= maxWidth; charEnd++ {
				for i, line := range bigLine {
					if charStart < len(line) {
						oneChar[i] = line[charStart:charEnd]
					} else {
						oneChar[i] = ""
					}
				}

				joinedArt := strings.Join(oneChar, "\n")
				if char, found := reverseAlphabet[joinedArt]; found {
					result.WriteRune(char)
					charStart = charEnd
					charFound = true
					break
				}
			}

			if !charFound {
				charStart++
			}
		}
	}
	return result.String(), nil
}

func Convert(input, banner string) (string, error) {
	if input == "" {
		return "", nil
	}

	ok, err := ValidInput(input)
	if !ok {
		return "", err
	}

	alphabet, err := getAlphab(banner)
	if err != nil {
		return "", err
	}

	input = strings.ReplaceAll(input, "\\n", "\n")

	words := strings.Split(input, "\n")

	result := make([]string, 0, len(words))

	for indexWord, word := range words {
		if word == "" {
			result = append(result, word)
			continue
		}

		var middleResult string

		for indexHeight := 0; indexHeight < height; indexHeight++ {
			for _, letter := range word {
				middleResult += alphabet[letter][indexHeight]
			}

			if indexHeight != height-1 || indexWord == len(words)-1 {
				middleResult += "\n"
			}
		}

		result = append(result, middleResult)
	}

	result = adjustNewLines(result)

	res := strings.Join(result, "\n")

	return res, nil
}

func getAlphab(banner string) (map[rune][]string, error) {
	cwd, _ := os.Getwd()

	if banner == "shadow.txt" {
		banner = "shadow.txt"
	} else if banner == "standard.txt" {
		banner = "standard.txt"
	} else if banner == "thinkertoy.txt" {
		banner = "thinkertoy.txt"
	} else {
		return nil, fmt.Errorf("error")
	}
	cwd = TrimCwd(cwd) + "/internal/banner/" + banner

	file, err := os.Open(cwd)
	if err != nil {
		return nil, fmt.Errorf("os error:%w", err)
	}

	scanner := bufio.NewScanner(file)
	textFromFile, err := os.ReadFile(cwd)
	if err != nil {
		return nil, fmt.Errorf("os error:%w", err)
	}

	hashStandard := uint64(8250112135784318067)
	hashShadow := uint64(8377067621923326644)
	hashTinkertoy := uint64(4863852022380994373)
	hashRead := strToHash(textFromFile)
	if (banner == "standard.txt" && hashStandard != hashRead) ||
		(banner == "shadow.txt" && hashShadow != hashRead) ||
		(banner == "thinkertoy.txt" && hashTinkertoy != hashRead) {
		return nil, errors.New("the banner was modified")
	}

	alphabet := make(map[rune][]string)

	skip := true

	var indexRune rune = 32

	for scanner.Scan() {
		if skip {
			skip = false

			continue
		}

		letter := make([]string, height)

		for i := 0; i < height; i++ {
			letter[i] = scanner.Text()

			if i != height-1 {
				scanner.Scan()
			}
		}

		alphabet[indexRune] = letter
		indexRune++

		skip = true
	}

	return alphabet, nil
}

func adjustNewLines(result []string) []string {
	onlyNewLine := true

	for _, v := range result {
		if v != "" {
			onlyNewLine = false
			break
		}
	}

	if onlyNewLine {
		return result
	}

	toAdd := false

	for _, v := range result {
		if v == "" {
			toAdd = true
			continue
		}
		toAdd = false
	}

	if toAdd {
		result = append(result, "")
	}

	return result
}
