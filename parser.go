package semijson

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	participle "github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type JValue struct {
	Literal *JLiteral `  @@ `
	Object  *JObject  `| @@ `
	Array   *JArray   `| @@ `
	Date    *JDate    `| @@`
}

func (v JValue) JSON() string {
	if v.Literal != nil {
		return v.Literal.JSON()
	} else if v.Object != nil {
		return v.Object.JSON()
	} else if v.Array != nil {
		return v.Array.JSON()
	} else if v.Date != nil {
		return v.Date.JSON()
	}
	return ""
}

type JLiteral struct {
	Null    string    `@Null`
	Boolean *JBoolean `| @("true" | "false")`
	String  *string   `| @String`
	Decimal *float64  `| @( "-"? Digit+ "." Digit+ )`
	Integer *int64    `| @( "-"? Digit+ )`
}

type JBoolean bool

func (b *JBoolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

func (b *JBoolean) JSON() string {
	return strconv.FormatBool(bool(*b))
}

func (v JLiteral) JSON() string {
	if v.Null != "" {
		return "null"
	} else if v.Boolean != nil {
		return v.Boolean.JSON()
	} else if v.String != nil {
		return fmt.Sprintf("%q", *v.String)
	} else if v.Decimal != nil {
		return fmt.Sprintf("%g", *v.Decimal)
	} else if v.Integer != nil {
		return strconv.Itoa(int(*v.Integer))
	}
	return ""
}

type JObject struct {
	Fields []JField `"{" (@@ ("," @@)*)? "}"`
}

func (v JObject) JSON() string {
	var sb strings.Builder
	sb.WriteRune('{')
	for idx, field := range v.Fields {
		sb.WriteString(field.JSON())
		if idx+1 < len(v.Fields) {
			sb.WriteRune(',')
		}
	}
	sb.WriteRune('}')
	return sb.String()
}

type JField struct {
	Key   string `@Ident ":"`
	Value JValue `@@`
}

func (v JField) JSON() string {
	return fmt.Sprintf("%q:%s", v.Key, v.Value.JSON())
}

type JArray struct {
	Values []JValue `"[" (@@ ("," @@)*)? "]"`
}

func (v JArray) JSON() string {
	var sb strings.Builder
	sb.WriteRune('[')
	for idx, value := range v.Values {
		sb.WriteString(value.JSON())
		if idx+1 < len(v.Values) {
			sb.WriteRune(',')
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

type JDate struct {
	Year       uint16   `"new" "Date" "(" @(Digit+) `
	Month      uint8    `"," @(Digit+)`
	Day        uint8    `"," @(Digit+) `
	MoreValues []uint64 `( "," @(Digit+) )* ")"`
}

func (v JDate) JSON() string {
	var hour, minute, second uint64
	if len(v.MoreValues) > 3 {
		hour = v.MoreValues[3]
	}
	if len(v.MoreValues) > 4 {
		minute = v.MoreValues[4]
	}
	if len(v.MoreValues) > 5 {
		second = v.MoreValues[5]
	}

	return fmt.Sprintf(
		"\"%04d-%02d-%02dT%02d:%02d:%02dZ\"",
		v.Year, v.Month, v.Day,
		hour, minute, second)
}

func newParser() *participle.Parser {
	def := lexer.MustSimple([]lexer.Rule{
		{"Null", `null|undefined`, nil},
		{"Digit", `\d`, nil},
		{"Ident", `[a-zA-Z]\w*`, nil},
		{"String", `"(\\"|[^"])*"|'(\\'|[^'])*'`, nil},

		{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
		{"Whitespace", `[ \t\n\r]+`, nil},
	})

	parser := participle.MustBuild(&JValue{},
		participle.Lexer(def),
		participle.Elide("Whitespace"),
		participle.Unquote("String"),
		participle.UseLookahead(10),
	)
	return parser
}

func ParseString(json string) (JValue, error) {
	var buffer bytes.Buffer
	_, err := buffer.Write([]byte(json))
	if err != nil {
		return JValue{}, err
	}

	return Parse(&buffer)
}

func Parse(reader io.Reader) (JValue, error) {
	parser := newParser()

	value := JValue{}
	err := parser.Parse("", reader, &value)
	if err != nil {
		return JValue{}, err
	}

	return value, nil
}
