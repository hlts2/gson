package gson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pquerna/ffjson/ffjson"
)

var (
	// ErrorIndexOutOfRange represents index out of range error
	ErrorIndexOutOfRange = errors.New("index out of range")

	// ErrorNotSlice represents error that target object is not slice
	ErrorNotSlice = errors.New("not slice")

	// ErrorNotMap represents error that target object is not map
	ErrorNotMap = errors.New("not map")

	// ErrorInvalidJSONKey represents error that json key dose not exist
	ErrorInvalidJSONKey = errors.New("invalid json Key")

	// ErrorInvalidSyntax represents invaild syntax error
	ErrorInvalidSyntax = errors.New("invalid syntax")

	// ErrorInvalidNumber represents invalid number
	ErrorInvalidNumber = errors.New("invalid number")

	// ErrorInvalidObject represents invalid object
	ErrorInvalidObject = errors.New("invalid object")
)

// Result represents a json value that is returned from GetByPath() and GetByKeys().
type Result struct {
	object interface{}
}

// Gson is gson base struct
type Gson struct {
	jsonObject interface{}
}

// NewGsonFromByte returns Gson instance created from byte slice
func NewGsonFromByte(data []byte) (*Gson, error) {
	g := new(Gson)

	if err := ffjson.Unmarshal(data, &g.jsonObject); err != nil {
		return nil, err
	}

	return g, nil
}

// NewGsonFromReader returns Gson instance created from io.Reader
func NewGsonFromReader(reader io.Reader) (*Gson, error) {
	g := new(Gson)

	if err := ffjson.NewDecoder().DecodeReader(reader, &g.jsonObject); err != nil {
		return nil, err
	}

	return g, nil
}

func isJSON(object interface{}) bool {
	if _, err := ffjson.Marshal(object); err != nil {
		return false
	}
	return true
}

// Indent converts json object to json string
func (g *Gson) Indent(prefix, indent string) (string, error) {
	return indentJSONString(g.jsonObject, prefix, indent)
}

func indentJSONString(object interface{}, prefix, indent string) (string, error) {
	data, err := ffjson.Marshal(object)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, data, prefix, indent); err != nil {
		return "", err
	}

	return buf.String(), nil
}

/*
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
*/

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
		return nil, ErrorNotSlice
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

// Interface returns json object of Result
func (r *Result) Interface() interface{} {
	return r.object
}

// Uint8 converts an interface{} to a uint8 and returns an error if types don't match.
func (r *Result) Uint8() (uint8, error) {
	const fn = "Uint8"

	switch v := r.object.(type) {
	case int:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint8(v), nil
	case int8:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint8(v), nil
	case int16:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint8(v), nil
	case int32:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint8(v), nil
	case int64:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint8(v), nil
	case uint:
		return uint8(v), nil
	case uint8:
		return v, nil
	case uint16:
		return uint8(v), nil
	case uint32:
		return uint8(v), nil
	case uint64:
		return uint8(v), nil
	case float32:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint8(v), nil
	case float64:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint8(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		u, err := strconv.ParseUint(v, 0, 8)
		if err == nil {
			return uint8(u), nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Uint16 converts an interface{} to a uint16 and returns an error if types don't match.
func (r *Result) Uint16() (uint16, error) {
	const fn = "Uint16"

	switch v := r.object.(type) {
	case int:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint16(v), nil
	case int8:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint16(v), nil
	case int16:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint16(v), nil
	case int32:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint16(v), nil
	case int64:
		if v < 0 {
			return 0, ErrorInvalidNumber
		}
		return uint16(v), nil
	case uint:
		return uint16(v), nil
	case uint8:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint32:
		return uint16(v), nil
	case uint64:
		return uint16(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		s, err := strconv.ParseUint(v, 0, 16)
		if err == nil {
			return uint16(s), nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Uint32 converts an interface{} to a uint32 and returns an error if types don't match.
func (r *Result) Uint32() (uint32, error) {
	const fn = "Uint32"

	switch v := r.object.(type) {
	case int:
		if v < 0 {
			return uint32(v), nil
		}
		return 0, ErrorInvalidNumber
	case int8:
		if v < 0 {
			return uint32(v), nil
		}
		return 0, ErrorInvalidNumber
	case int16:
		if v < 0 {
			return uint32(v), nil
		}
		return 0, ErrorInvalidNumber
	case int32:
		if v < 0 {
			return uint32(v), nil
		}
		return 0, ErrorInvalidNumber
	case int64:
		if v < 0 {
			return uint32(v), nil
		}
		return 0, ErrorInvalidNumber
	case uint:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint64:
		return uint32(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		u, err := strconv.ParseUint(v, 0, 32)
		if err == nil {
			return uint32(u), nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Uint64 converts an interface{} to a uint64 and returns an error if types don't match.
func (r *Result) Uint64() (uint64, error) {
	const fn = "Uint64"

	switch v := r.object.(type) {
	case int:
		if v < 0 {
			return uint64(v), nil
		}
		return 0, ErrorInvalidNumber
	case int8:
		if v < 0 {
			return uint64(v), nil
		}
		return 0, ErrorInvalidNumber
	case int16:
		if v < 0 {
			return uint64(v), nil
		}
		return 0, ErrorInvalidNumber
	case int32:
		if v < 0 {
			return uint64(v), nil
		}
		return 0, ErrorInvalidNumber
	case int64:
		if v < 0 {
			return uint64(v), nil
		}
		return 0, ErrorInvalidNumber
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		u, err := strconv.ParseUint(v, 0, 64)
		if err == nil {
			return u, nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Int8 converts an interface{} to a int8 and returns an error if types don't match.
func (r *Result) Int8() (int8, error) {
	const fn = "Int8"

	switch v := r.object.(type) {
	case int:
		return int8(v), nil
	case int8:
		return int8(v), nil
	case int16:
		return int8(v), nil
	case int32:
		return int8(v), nil
	case int64:
		return int8(v), nil
	case uint:
		return int8(v), nil
	case uint8:
		return int8(v), nil
	case uint16:
		return int8(v), nil
	case uint32:
		return int8(v), nil
	case uint64:
		return int8(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseInt(v, 0, 8)
		if err == nil {
			return int8(i), nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Int16 converts an interface{} to a int16 and returns an error if types don't match.
func (r *Result) Int16() (int16, error) {
	const fn = "Int16"

	switch v := r.object.(type) {
	case int:
		return int16(v), nil
	case int8:
		return int16(v), nil
	case int16:
		return int16(v), nil
	case int32:
		return int16(v), nil
	case int64:
		return int16(v), nil
	case uint:
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case uint16:
		return int16(v), nil
	case uint32:
		return int16(v), nil
	case uint64:
		return int16(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseInt(v, 0, 16)
		if err == nil {
			return int16(i), nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Int32 converts an interface{} to a int32 and returns an error if types don't match.
func (r *Result) Int32() (int32, error) {
	const fn = "Int32"

	switch v := r.object.(type) {
	case int:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return int32(v), nil
	case int64:
		return int32(v), nil
	case uint:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		return int32(v), nil
	case uint64:
		return int32(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseInt(v, 0, 32)
		if err == nil {
			return int32(i), nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Int64 converts an interface{} to a int64 and returns an error if types don't match.
func (r *Result) Int64() (int64, error) {
	const fn = "Int64"

	switch v := r.object.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseInt(v, 0, 64)
		if err == nil {
			return i, nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Int converts an interface{} to a int and returns an error if types don't match.
func (r *Result) Int() (int, error) {
	const fn = "Int"

	switch v := r.object.(type) {
	case int:
		return int(v), nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseInt(v, 0, 0)
		if err == nil {
			return int(i), nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Float32 converts an interface{} to a float32 and returns an error if types don't match.
func (r *Result) Float32() (float32, error) {
	const fn = "Float32"

	switch v := r.object.(type) {
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseFloat(v, 32)
		if err == nil {
			return float32(i), nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// Float64 converts an interface{} to a float64 and returns an error if types don't match.
func (r *Result) Float64() (float64, error) {
	const fn = "Float64"

	switch v := r.object.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return i, nil
		}
		return 0, err
	case nil:
		return 0, nil
	default:
		return 0, ErrorInvalidObject
	}
}

// String converts an interface{} to a string and returns an error if types don't match.
func (r *Result) String() (string, error) {
	const fn = "String"

	switch v := r.object.(type) {
	case int:
		return strconv.Itoa(v), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case bool:
		if v {
			return "true", nil
		}
		return "false", nil
	case string:
		return v, nil
	default:
		return "", ErrorInvalidObject
	}
}

// Bool converts an interface{} to a bool and returns an error if types don't match.
func (r *Result) Bool() (bool, error) {
	const fn = "Bool"

	switch v := r.object.(type) {
	case bool:
		return v, nil
	case nil:
		return false, nil
	case string:
		b, err := strconv.ParseBool(v)
		if err == nil {
			return b, nil
		}
		return false, err
	default:
		return false, ErrorInvalidObject
	}
}

// Slice converts an Result pointer slice and returns an error if types don't match.
func (r *Result) Slice() ([]*Result, error) {
	const fn = "Slice"

	switch slice := r.object.(type) {
	case []interface{}:

		results := make([]*Result, 0, len(slice))

		for _, val := range slice {
			results = append(results, &Result{object: val})
		}

		return results, nil
	default:
		return nil, ErrorNotSlice
	}
}

// Map converts an Result pointer slice and returns an error if types don't match.
func (r *Result) Map() (map[string]*Result, error) {
	const fn = "Map"

	switch m := r.object.(type) {
	case map[string]interface{}:
		rMap := make(map[string]*Result, len(m))

		for key, val := range m {
			rMap[key] = &Result{object: val}
		}

		return rMap, nil
	default:
		return nil, ErrorNotMap
	}
}
