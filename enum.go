package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func Run(packageName string, prefix string, values ...string) error {
	data := make(map[string]interface{})
	data["Package"] = packageName
	data["Items"] = values
	data["Type"] = prefix

	fErr, err := os.Create("enum_error.go")
	if err != nil {
		return err
	}
	defer fErr.Close()

	t1 := template.New("t_err")
	t1, err = t1.Parse(fileErrTemplate)
	if err != nil {
		return err
	}
	if err := t1.Execute(fErr, data); err != nil {
		return err
	}

	fileName := fmt.Sprintf("enum_%s.go", strings.ToLower(prefix))
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	t2 := template.New("t_enum")
	t2.Funcs(map[string]interface{}{"Title": strings.Title})
	t2, err = t2.Parse(fileTemplate)
	if err != nil {
		return err
	}
	return t2.Execute(f, data)
}

const fileErrTemplate = `// Code generated by enum. DO NOT EDIT

package {{.Package}}

import "errors"

var ErrEnumUnknown = errors.New("unknown")
`

const fileTemplate = `// Code generated by enum (github.com/mpkondrashin/enum). DO NOT EDIT

package {{.Package}}

import (
    "encoding/json"
    "fmt"
    "strconv"
)

type {{.Type}} int

const (
{{range .Items}}    {{$.Type}}{{Title .}} {{$.Type }} = iota
{{end -}})

func (v {{.Type}})String() string {
    s, ok := map[{{.Type}}]string {
{{range .Items}}        {{$.Type}}{{Title .}}: "{{.}}",
{{end}}    }[v]
    if ok {
        return s
    }
    return "{{.Type}}(" + strconv.FormatInt(int64(v), 10) + ")"
}

func (s *{{.Type}}) UnmarshalJSON(data []byte) error {
    var v string
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    result, ok := map[string]{{.Type}}{
{{range .Items}}        "{{.}}": {{$.Type}}{{Title .}},
{{end}}    }[v]
    if !ok {
        return fmt.Errorf("%w: %s", ErrEnumUnknown, v)
    }
    *s = result
    return nil
}
`
