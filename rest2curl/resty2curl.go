package resty2curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Curl []string

// append appends a string to the Curl
func (c *Curl) append(newSlice ...string) {
	*c = append(*c, newSlice...)
}

// String returns a ready to copy/paste command
func (c *Curl) String() string {
	return strings.Join(*c, " ")
}

func bashEscape(str string) string {
	return `'` + strings.Replace(str, `'`, `'\''`, -1) + `'`
}

func GetCurFromRestyRequest(req resty.Request) (*Curl, error) {
	command := Curl{}
	command.append("curl")
	command.append("-X", bashEscape(req.Method))
	if req.Body != nil {
		var buff bytes.Buffer
		b, _ := json.Marshal(req.Body)
		r := bytes.NewReader(b)
		_, err := buff.ReadFrom(r)
		if err != nil {
			return nil, fmt.Errorf("getCurl: buffer read from body erorr: %w", err)
		}
		if len(buff.String()) > 0 {
			bodyEscaped := bashEscape(buff.String())
			command.append("-d", bodyEscaped)
		}
	}

	var keys []string

	for k := range req.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		command.append("-H", bashEscape(fmt.Sprintf("%s: %s", k, strings.Join(req.Header[k], " "))))
	}

	command.append(bashEscape(req.URL))

	command.append("--compressed")

	return &command, nil
}
