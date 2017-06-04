package kolpa

import (
	"regexp"
	"bufio"
	"os"
	"log"
	"math/rand"
	"strings"
	"time"
)

// Parses and replaces tags in a text with provided values in the map m.
func (g *Generator) parser(text string, m map[string]string) string {
	src := []byte(text)
	search := regexp.MustCompile(`{{(.*?)}}`)

	src = search.ReplaceAllFunc(src, func(s []byte) []byte {
		return []byte(m[string(s)[2:len(s)-2]])
	})

	return string(src)
}


// Concatenates multiple string slices by using append function and returns new slice.
func appendMultiple(slices ...[]string) []string {
	base := slices[0]
	rest := slices[1:]
	
	for _, slice := range rest {
		base = append(base, slice...)
	}

	return base
}


// Reads the file 'fName' and returns its content as a slice of strings.
func (g *Generator) fileToSlice(fName string) []string {
	var res []string
	path := os.Getenv("GOPATH") + "/src/kolpa/data/" + g.Locale + "/" + fName
	file, err := os.Open(path)

	if err != nil {
		return res	
	}
	defer file.Close()
	

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

// Reads the tab separated file 'fName' and returns its content as a map of strings to strings.
func (g *Generator) fileToMap(fName string) map[string]string {
	m := make(map[string]string)
	path := os.Getenv("GOPATH") + "/src/kolpa/data/" + g.Locale + "/" + fName
	file, err := os.Open(path)

	if err != nil {
		return m	
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		m[line[0]] = line[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return m
}

// Returns random item from the given string slice.
func getRandom(options []string) string {
	rand.Seed(time.Now().Unix())
	return options[rand.Intn(len(options))]
}

// Returns random boolean variable.
func randBool() bool {
	rand.Seed(time.Now().Unix())
	val := rand.Float64()
	if val <= 0.5 {
		return true
	} else {
		return false
	}
}