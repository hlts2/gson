package gson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	// ErrorIndexOutOfRange represents index out of range error
	ErrorIndexOutOfRange = errors.New("index out of range")

	// ErrorNotArray represents error that target object is not array
	ErrorNotArray = errors.New("not array")

	// ErrorNotMap represents error that target object is not map
	ErrorNotMap = errors.New("not map")

	// ErrorInvalidJSONKey represents error that json key dose not exist
	ErrorInvalidJSONKey = errors.New("invalid json Key")

	// ErrorInvalidSyntax represents invaild syntax error
	ErrorInvalidSyntax = errors.New("invalid syntax")
)

// ResultError represents a conversion error
type ResultError struct {
	Fn     string
	Object interface{}
	Err    error
}

func (e *ResultError) Error() string {
	return "goson." + e.Fn + ": parsing " + Quote(e.Object) + ": " + e.Err.Error()
}

// Quote returns quoted object string
func Quote(object interface{}) string {
	return fmt.Sprintf("\"%v\"", object)
}

// Result represents a json value that is returned from Search() and Path().
type Result struct {
	object interface{}
}

// Goson is goson base struct
type Goson struct {
	jsonObject interface{}
}

// NewGosonFromByte returns Goson instance created from byte array
func NewGosonFromByte(data []byte) (*Goson, error) {
	g := new(Goson)

	if err := decode(bytes.NewReader(data), &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

// NewGosonFromString returns Goson instance created from string
func NewGosonFromString(data string) (*Goson, error) {
	g := new(Goson)

	if err := decode(strings.NewReader(data), &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

// NewGosonFromReader returns Goson instance created from io.Reader
func NewGosonFromReader(reader io.Reader) (*Goson, error) {
	g := new(Goson)

	if err := decode(reader, &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

func decode(reader io.Reader, object *interface{}) error {
	dec := json.NewDecoder(reader)
	if err := dec.Decode(object); err != nil {
		return err
	}
	return nil
}

func isJSON(object interface{}) bool {
	if _, err := json.Marshal(object); err != nil {
		return false
	}
	return true
}

// Indent converts json object to json string
func (g *Goson) Indent(prefix, indent string) (string, error) {
	return indentJSONString(g.jsonObject, prefix, indent)
}

func indentJSONString(object interface{}, prefix, indent string) (string, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, data, prefix, indent); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// HasWithKeys returns bool if there is json value coresponding to keys
func (g *Goson) HasWithKeys(keys ...string) bool {
	var err error

	jsonObject := g.jsonObject
	for _, key := range keys {
		jsonObject, err = search(jsonObject, key)
		if err != nil {
			return false
		}
	}
	return true
}

// HasWithPath returns bool if there is json value coresponding to path
func (g *Goson) HasWithPath(path string) bool {
	var err error

	jsonObject := g.jsonObject
	for _, key := range strings.Split(path, "/")[1:] {
		jsonObject, err = search(jsonObject, key)
		if err != nil {
			return false
		}
	}
	return true
}

// Search returns json value corresponding to keys. keys represents key of hierarchy of json
func (g *Goson) Search(keys ...string) (*Result, error) {
	var err error

	jsonObject := g.jsonObject

	for _, key := range keys {
		if jsonObject, err = search(jsonObject, key); err != nil {
			return nil, err
		}
	}
	return &Result{jsonObject}, nil
}

// Path returns json value corresponding to path.
func (g *Goson) Path(path string) (*Result, error) {
	var err error

	jsonObject := g.jsonObject

	for _, key := range strings.Split(path, "/")[1:] {
		if jsonObject, err = search(jsonObject, key); err != nil {
			return nil, err
		}
	}
	return &Result{jsonObject}, nil
}

func search(object interface{}, key string) (interface{}, error) {
	index, err := strconv.Atoi(key)
	if err == nil {
		switch object.(type) {
		case []interface{}:
		default:
			return nil, ErrorNotArray
		}

		v := object.([]interface{})

		if 0 <= index && index < len(v) {
			return v[index], nil
		}

		return nil, ErrorIndexOutOfRange
	}

	switch object.(type) {
	case map[string]interface{}:
	default:
		return nil, ErrorNotArray
	}

	v, ok := object.(map[string]interface{})[key]
	if !ok {
		return nil, ErrorInvalidJSONKey
	}

	return v, nil
}

// Indent converts json object to json string
func (r *Result) Indent(prefix, indent string) string {
	str, err := indentJSONString(r.object, prefix, indent)
	if err != nil {
		return fmt.Sprintf("%v", r.object)
	}
	return str
}

// Uint8 converts an interface{} to a uint8 and returns an error if types don't match.
func (r *Result) Uint8() (uint8, error) {
	const fn = "Uint8"

	switch r.object.(type) {
	case uint8:
		return r.object.(uint8), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Uint16 converts an interface{} to a uint16 and returns an error if types don't match.
func (r *Result) Uint16() (uint16, error) {
	const fn = "Uint16"

	switch r.object.(type) {
	case uint16:
		return r.object.(uint16), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Uint32 converts an interface{} to a uint32 and returns an error if types don't match.
func (r *Result) Uint32() (uint32, error) {
	const fn = "Uint32"

	switch r.object.(type) {
	case uint32:
		return r.object.(uint32), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Uint64 converts an interface{} to a uint64 and returns an error if types don't match.
func (r *Result) Uint64() (uint64, error) {
	const fn = "Uint64"

	switch r.object.(type) {
	case uint64:
		return r.object.(uint64), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int8 converts an interface{} to a int8 and returns an error if types don't match.
func (r *Result) Int8() (int8, error) {
	const fn = "Int8"

	switch r.object.(type) {
	case int8:
		return r.object.(int8), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int16 converts an interface{} to a int16 and returns an error if types don't match.
func (r *Result) Int16() (int16, error) {
	const fn = "Int16"

	switch r.object.(type) {
	case int16:
		return r.object.(int16), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int32 converts an interface{} to a int32 and returns an error if types don't match.
func (r *Result) Int32() (int32, error) {
	const fn = "Int32"

	switch r.object.(type) {
	case int32:
		return r.object.(int32), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int64 converts an interface{} to a int64 and returns an error if types don't match.
func (r *Result) Int64() (int64, error) {
	const fn = "Int64"

	switch r.object.(type) {
	case int64:
		return r.object.(int64), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int converts an interface{} to a int and returns an error if types don't match.
func (r *Result) Int() (int, error) {
	const fn = "Int"

	switch r.object.(type) {
	case int:
		return r.object.(int), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Float32 converts an interface{} to a float32 and returns an error if types don't match.
func (r *Result) Float32() (float32, error) {
	const fn = "Float32"

	switch r.object.(type) {
	case float32:
		return r.object.(float32), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Float64 converts an interface{} to a float64 and returns an error if types don't match.
func (r *Result) Float64() (float64, error) {
	const fn = "Float64"

	switch r.object.(type) {
	case float64:
		return r.object.(float64), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Byte converts an interface{} to a byte and returns an error if types don't match.
func (r *Result) Byte() (byte, error) {
	const fn = "Byte"

	switch r.object.(type) {
	case byte:
		return r.object.(byte), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Rune converts an interface{} to a rune and returns an error if types don't match.
func (r *Result) Rune() (rune, error) {
	const fn = "Rune"

	switch r.object.(type) {
	case rune:
		return r.object.(rune), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Complex64 converts an interface{} to a complex64 and returns an error if types don't match.
func (r *Result) Complex64() (complex64, error) {
	const fn = "Complex64"

	switch r.object.(type) {
	case complex64:
		return r.object.(complex64), nil
	default:
		return 0 + 0i, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Complex128 converts an interface{} to a complexa128 and returns an error if types don't match.
func (r *Result) Complex128() (complex128, error) {
	const fn = "Complex128"

	switch r.object.(type) {
	case complex128:
		return r.object.(complex128), nil
	default:
		return 0 + 0i, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Bool converts an interface{} to a bool and returns an error if types don't match.
func (r *Result) Bool() (bool, error) {
	const fn = "Bool"

	switch r.object.(type) {
	case bool:
		return r.object.(bool), nil
	default:
		return false, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// MapInterface converts an interface{} to a map[string]interface{} and returns an error if types don't match.
func (r *Result) MapInterface() (map[string]interface{}, error) {
	const fn = "MapInterface"

	switch r.object.(type) {
	case map[string]interface{}:
		return r.object.(map[string]interface{}), nil
	default:
		return map[string]interface{}{}, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// MapInterfaceSlice converts an interface{} to a []map[string]interface{} and returns an error if types don't match.
func (r *Result) MapInterfaceSlice() ([]map[string]interface{}, error) {
	const fn = "MapInterfaceSlice"

	switch r.object.(type) {
	case []map[string]interface{}:
		return r.object.([]map[string]interface{}), nil
	default:
		return []map[string]interface{}{}, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}
