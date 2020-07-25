package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seenBrowsers2 := make(map[string]byte)
	users := make([]map[string]interface{}, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		user := make(map[string]interface{})
		// fmt.Printf("%v %v\n", err, line)
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	fmt.Fprintln(out, "found users:")
	for i, user := range users {
		processUser(out, user, seenBrowsers2, i)
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers2))
}

func processUser(out io.Writer, user map[string]interface{}, seenBrowsers map[string]byte, i int) {
	isAndroid := false
	isMSIE := false

	browsers, ok := user["browsers"].([]interface{})
	if !ok {
		return
	}

	for _, browserRaw := range browsers {
		browser, ok := browserRaw.(string)
		if !ok {
			continue
		}
		a := strings.Contains(browser, "Android")
		isAndroid = a || isAndroid
		ie := strings.Contains(browser, "MSIE")
		isMSIE = ie || isMSIE
		if a || ie {
			_, ok := seenBrowsers[browser]
			if !ok {
				seenBrowsers[browser] = 0
			}
		}
	}

	if !(isAndroid && isMSIE) {
		return
	}
	r := regexp.MustCompile("@")
	email := r.ReplaceAllString(user["email"].(string), " [at] ")
	fmt.Fprintf(out, "[%d] %s <%s>\n", i, user["name"], email)
	//foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user["name"], email)
}
