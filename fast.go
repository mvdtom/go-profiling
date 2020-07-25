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

type User struct {
	Name     string
	Email    string
	Browsers []string
}

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seenBrowsers2 := make(map[string]byte)

	fmt.Fprintln(out, "found users:")
	scanner := bufio.NewScanner(file)
	i := 0
	user := User{}
	for scanner.Scan() {
		line := scanner.Bytes()
		//user := &User{}
		err := json.Unmarshal(line, &user)
		if err != nil {
			panic(err)
		}
		processUser(out, &user, seenBrowsers2, i)
		i++
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers2))
}

func processUser(out io.Writer, user *User, seenBrowsers map[string]byte, i int) {
	isAndroid := false
	isMSIE := false

	for _, browser := range user.Browsers {
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
	email := r.ReplaceAllString(user.Email, " [at] ")
	fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, email)
}
