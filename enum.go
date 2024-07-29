package main

import (
	"os"
	"strings"
	"text/tabwriter"
	"text/template"
	"unicode"
)

func Run(fileName string, packageName string, typeName string, noPrefix bool, names []string) error {
	data := make(map[string]interface{})
	data["Command"] = CommandLine()
	data["Package"] = packageName
	var n []string
	for _, name := range names {
		n = append(n, FixName(name))
	}
	data["Names"] = n
	data["Strings"] = names
	data["Type"] = typeName
	data["Prefix"] = ""
	if !noPrefix {
		data["Prefix"] = typeName
	}
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	funcMap := template.FuncMap{
		"ToLower":    strings.ToLower,
		"Capitalize": Capitalize,
	}
	// tabw := tabwriter.NewWriter(f, 8, 8, 8, ' ', 0)
	tabw := tabwriter.NewWriter(f, 1, 1, 1, ' ', 0)
	t2 := template.New("t_enum").Funcs(funcMap)
	t2, err = t2.Parse(fileTemplate)
	if err != nil {
		return err
	}
	return t2.Execute(tabw, data)
}

func CommandLine() string {
	return strings.Join(os.Args, " ")
}

func FixName(str string) string {
	str = strings.ReplaceAll(str, "-", "_")
	str = strings.ReplaceAll(str, " ", "_")
	return str
}

func Capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// "encoding/xml"

const fileTemplate = `// Code generated by enum (github.com/mpkondrashin/enum) using following command:
// {{.Command}}
// DO NOT EDIT!

package {{.Package}}

import (
    "encoding/json"
    "errors"
    "fmt"
    "strconv"
    "strings"
)

type {{.Type}} int

const (
{{range .Names}}    {{$.Prefix}}{{Capitalize .}}	{{$.Type }}	= iota
{{end -}})

{{ $strings := .Strings }}
// String - return string representation for {{.Type}} value
func (v {{.Type}})String() string {
    s, ok := map[{{.Type}}]string {
{{ range $index,$name := .Names }}         {{$.Prefix}}{{Capitalize $name}}:	"{{index $strings $index}}",
{{end}}    }[v]
    if ok {
        return s
    }
    return "{{.Type}}(" + strconv.FormatInt(int64(v), 10) + ")"
}

// ErrUnknown{{.Type}} - will be returned wrapped when parsing string
// containing unrecognized value.
var ErrUnknown{{.Type}} = errors.New("unknown {{.Type}}")

{{ $names := .Names }}
var map{{.Type}}FromString = map[string]{{.Type}}{
{{ range $index,$string := .Strings }}    "{{ToLower $string}}":    {{$.Prefix}}{{ $name := index $names $index }}{{Capitalize $name}},
{{end}}}

// UnmarshalJSON implements the Unmarshaler interface of the json package for {{.Type}}.
func (s *{{.Type}}) UnmarshalJSON(data []byte) error {
    var v string
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    result, ok := map{{.Type}}FromString[strings.ToLower(v)]
    if !ok {
        return fmt.Errorf("%w: %s", ErrUnknown{{.Type}}, v)
    }
    *s = result
    return nil
}

// MarshalJSON implements the Marshaler interface of the json package for {{.Type}}.
func (s {{.Type}}) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("\"%v\"", s)), nil
}

// UnmarshalYAML implements the Unmarshaler interface of the yaml.v3 package for {{.Type}}.
func (s *{{.Type}}) UnmarshalYAML(unmarshal func(interface{}) error) error {
    var v string
    if err := unmarshal(&v); err != nil {
        return err
    }
    result, ok := map{{.Type}}FromString[strings.ToLower(v)]		
    if !ok {
        return fmt.Errorf("%w: %s", ErrUnknown{{.Type}}, v)
    }
    *s = result
    return nil
}
`

/*

// MarshalXML implements the Marshaler interface of the xml package for {{.Type}}.
func (s {{.Type}})  MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(s.String(), start)
}

// UnmarshalXML implements the Unmarshaler interface of the xml package for {{.Type}}.
func (s *{{.Type}}) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	result, ok := map{{.Type}}FromString[strings.ToLower(v)]
    if !ok {
        return fmt.Errorf("%w: %s", ErrUnknown{{.Type}}, v)
    }
	*s = result
	return nil
}

`
*/
