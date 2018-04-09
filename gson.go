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
	return "Gson." + e.Fn + ": parsing " + Quote(e.Object) + ": " + e.Err.Error()
}

// Quote returns quoted object string
func Quote(object interface{}) string {
	return fmt.Sprintf("\"%v\"", object)
}

// Result represents a json value that is returned from Search() and Path().
type Result struct {
	object interface{}
}

// Gson is gson base struct
type Gson struct {
	jsonObject interface{}
}

// NewGsonFromByte returns Gson instance created from byte array
func NewGsonFromByte(data []byte) (*Gson, error) {
	g := new(Gson)

	if err := decode(bytes.NewReader(data), &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

// NewGsonFromString returns Gson instance created from string
func NewGsonFromString(data string) (*Gson, error) {
	g := new(Gson)

	if err := decode(strings.NewReader(data), &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

// NewGsonFromReader returns Gson instance created from io.Reader
func NewGsonFromReader(reader io.Reader) (*Gson, error) {
	g := new(Gson)

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
func (g *Gson) Indent(prefix, indent string) (string, error) {
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
func (g *Gson) HasWithKeys(keys ...string) bool {
	var err error

	jsonObject := g.jsonObject
	for _, key := range keys {
		jsonObject, err = getByKey(jsonObject, key)
		if err != nil {
			return false
		}
	}
	return true
}

// HasWithPath returns bool if there is json value coresponding to path
func (g *Gson) HasWithPath(path string) bool {
	var err error

	jsonObject := g.jsonObject
	for _, key := range strings.Split(path, ".") {
		jsonObject, err = getByKey(jsonObject, key)
		if err != nil {
			return false
		}
	}
	return true
}

// GetByKeys returns json value corresponding to keys. keys represents key of hierarchy of json
func (g *Gson) GetByKeys(keys ...string) (*Result, error) {
	var err error
	jsonObject := g.jsonObject

	for _, key := range keys {
		if jsonObject, err = getByKey(jsonObject, key); err != nil {
			return nil, err
		}
	}
	return &Result{jsonObject}, nil
}

// GetByPath returns json value corresponding to path.
func (g *Gson) GetByPath(path string) (*Result, error) {
	keys := strings.Split(path, ".")

	var err error
	jsonObject := g.jsonObject

	for _, key := range keys {
		if jsonObject, err = getByKey(jsonObject, key); err != nil {
			return nil, err
		}
	}
	return &Result{jsonObject}, nil
}

func getByKey(object interface{}, key string) (interface{}, error) {
	index, err := strconv.Atoi(key)
	if err == nil {
		if v, ok := object.([]interface{}); ok {
			if index >= 0 && index < len(v) {
				return v[index], nil
			}
			return nil, ErrorIndexOutOfRange
		}
		return nil, ErrorNotArray
	}

	if m, ok := object.(map[string]interface{}); ok {
		if v, ok := m[key]; ok {
			return v, nil
		}
		return nil, ErrorInvalidJSONKey
	}
	return nil, ErrorNotMap
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
	case int:
		return uint8(r.object.(int)), nil
	}
	return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
}

// Uint16 converts an interface{} to a uint16 and returns an error if types don't match.
func (r *Result) Uint16() (uint16, error) {
	const fn = "Uint16"

	switch r.object.(type) {
	case int:
		return uint16(r.object.(int)), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Uint32 converts an interface{} to a uint32 and returns an error if types don't match.
func (r *Result) Uint32() (uint32, error) {
	const fn = "Uint32"

	switch r.object.(type) {
	case int:
		return uint32(r.object.(int)), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Uint64 converts an interface{} to a uint64 and returns an error if types don't match.
func (r *Result) Uint64() (uint64, error) {
	const fn = "Uint64"

	switch r.object.(type) {
	case int:
		return uint64(r.object.(int)), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int8 converts an interface{} to a int8 and returns an error if types don't match.
func (r *Result) Int8() (int8, error) {
	const fn = "Int8"

	switch r.object.(type) {
	case int:
		return int8(r.object.(int)), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int16 converts an interface{} to a int16 and returns an error if types don't match.
func (r *Result) Int16() (int16, error) {
	const fn = "Int16"

	switch r.object.(type) {
	case int:
		return int16(r.object.(int)), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int32 converts an interface{} to a int32 and returns an error if types don't match.
func (r *Result) Int32() (int32, error) {
	const fn = "Int32"

	switch r.object.(type) {
	case int:
		return int32(r.object.(int)), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int64 converts an interface{} to a int64 and returns an error if types don't match.
func (r *Result) Int64() (int64, error) {
	const fn = "Int64"

	switch r.object.(type) {
	case int:
		return int64(r.object.(int)), nil
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
	case float64:
		return float32(r.object.(float64)), nil
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
	case complex128:
		return complex64(r.object.(complex128)), nil
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

// String converts an interface{} to a string and returns an error if types don't match.
func (r *Result) String() (string, error) {
	const fn = "String"

	switch r.object.(type) {
	case string:
		return r.object.(string), nil
	default:
		return "", &ResultError{fn, r.object, ErrorInvalidSyntax}
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
