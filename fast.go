package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
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

	seenBrowsers2 := make(map[string]byte, 115)

	fmt.Fprintln(out, "found users:")
	scanner := bufio.NewScanner(file)
	sBuff := make([]byte, 0, 800)
	scanner.Buffer(sBuff, 1000)
	i := 0
	user := User{}
	l := jlexer.Lexer{}

	for scanner.Scan() {
		line := scanner.Bytes()
		l.Data = line
		//user := &User{}
		//err := json.Unmarshal(line, &user)

		err := user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}
		processUser(out, &user, seenBrowsers2, i)
		i++
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers2))
}

var r = regexp.MustCompile("@")

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
	email := r.ReplaceAllString(user.Email, " [at] ")
	fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, email)
}

//code produced by easyJson codegenerator. Placed here only for submiting the task
func easyjson9f2eff5fDecodeCourseraHwHw3BenchDto(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeCourseraHwHw3BenchDto(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix)
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeCourseraHwHw3BenchDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeCourseraHwHw3BenchDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeCourseraHwHw3BenchDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeCourseraHwHw3BenchDto(l, v)
}
