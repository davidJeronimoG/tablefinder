package utilities

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

	// ReadValues reads a file and returns a map of key-value pairs.
	func GetValuesFromFile (file string) (map[string]string, error){

    f, err := os.Open(file)

	fieldsMap := make(map[string]string)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, "  ")

		re := regexp.MustCompile(`/*No\b`)

		if re.MatchString(array[len(array)-1]) {
			continue
		} else {

		fieldsMap[array[0]] = array[len(array)-1]
		}

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return fieldsMap, nil
	}