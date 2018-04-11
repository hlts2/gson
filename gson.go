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

// Result represents a json value that is returned from GetByPath() and GetByKeys().
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
	if err := json.NewDecoder(reader).Decode(object); err != nil {
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

// Uint8 converts an interface{} of Result object to a uint8
func (r *Result) Uint8() uint8 {
	switch v := r.object.(type) {
	case int:
		if v < 0 {
			return 0
		}
		return uint8(v)
	case int8:
		if v < 0 {
			return 0
		}
		return uint8(v)
	case int16:
		if v < 0 {
			return 0
		}
		return uint8(v)
	case int32:
		if v < 0 {
			return 0
		}
		return uint8(v)
	case int64:
		if v < 0 {
			return 0
		}
		return uint8(v)
	case uint:
		return uint8(v)
	case uint8:
		return v
	case uint16:
		return uint8(v)
	case uint32:
		return uint8(v)
	case uint64:
		return uint8(v)
	case float32:
		if v < 0 {
			return 0
		}
		return uint8(v)
	case float64:
		if v < 0 {
			return 0
		}
		return uint8(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		u, err := strconv.ParseUint(v, 0, 8)
		if err == nil {
			return uint8(u)
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Uint16 converts an interface{} of Result object to a uint16
func (r *Result) Uint16() uint16 {
	switch v := r.object.(type) {
	case int:
		if v < 0 {
			return 0
		}
		return uint16(v)
	case int8:
		if v < 0 {
			return 0
		}
		return uint16(v)
	case int16:
		if v < 0 {
			return 0
		}
		return uint16(v)
	case int32:
		if v < 0 {
			return 0
		}
		return uint16(v)
	case int64:
		if v < 0 {
			return 0
		}
		return uint16(v)
	case uint:
		return uint16(v)
	case uint8:
		return uint16(v)
	case uint16:
		return v
	case uint32:
		return uint16(v)
	case uint64:
		return uint16(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		s, err := strconv.ParseUint(v, 0, 16)
		if err == nil {
			return uint16(s)
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Uint32 converts an interface{} of Result object to a uint32
func (r *Result) Uint32() uint32 {
	switch v := r.object.(type) {
	case int:
		if v < 0 {
			return uint32(v)
		}
		return 0
	case int8:
		if v < 0 {
			return uint32(v)
		}
		return 0
	case int16:
		if v < 0 {
			return uint32(v)
		}
		return 0
	case int32:
		if v < 0 {
			return uint32(v)
		}
		return 0
	case int64:
		if v < 0 {
			return uint32(v)
		}
		return 0
	case uint:
		return uint32(v)
	case uint8:
		return uint32(v)
	case uint16:
		return uint32(v)
	case uint32:
		return v
	case uint64:
		return uint32(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		u, err := strconv.ParseUint(v, 0, 32)
		if err == nil {
			return uint32(u)
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Uint64 converts an interface{} of Result object to a uint64
func (r *Result) Uint64() uint64 {
	switch v := r.object.(type) {
	case int:
		if v < 0 {
			return uint64(v)
		}
		return 0
	case int8:
		if v < 0 {
			return uint64(v)
		}
		return 0
	case int16:
		if v < 0 {
			return uint64(v)
		}
		return 0
	case int32:
		if v < 0 {
			return uint64(v)
		}
		return 0
	case int64:
		if v < 0 {
			return uint64(v)
		}
		return 0
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		u, err := strconv.ParseUint(v, 0, 64)
		if err == nil {
			return u
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Int8 converts an interface{} of Result object to a int8
func (r *Result) Int8() int8 {
	switch v := r.object.(type) {
	case int:
		return int8(v)
	case int8:
		return int8(v)
	case int16:
		return int8(v)
	case int32:
		return int8(v)
	case int64:
		return int8(v)
	case uint:
		return int8(v)
	case uint8:
		return int8(v)
	case uint16:
		return int8(v)
	case uint32:
		return int8(v)
	case uint64:
		return int8(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := strconv.ParseInt(v, 0, 8)
		if err == nil {
			return int8(i)
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Int16 converts an interface{} of Result object to a int16
func (r *Result) Int16() int16 {
	switch v := r.object.(type) {
	case int:
		return int16(v)
	case int8:
		return int16(v)
	case int16:
		return int16(v)
	case int32:
		return int16(v)
	case int64:
		return int16(v)
	case uint:
		return int16(v)
	case uint8:
		return int16(v)
	case uint16:
		return int16(v)
	case uint32:
		return int16(v)
	case uint64:
		return int16(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := strconv.ParseInt(v, 0, 16)
		if err == nil {
			return int16(i)
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Int32 converts an interface{} of Result object to a int32
func (r *Result) Int32() int32 {
	switch v := r.object.(type) {
	case int:
		return int32(v)
	case int8:
		return int32(v)
	case int16:
		return int32(v)
	case int32:
		return int32(v)
	case int64:
		return int32(v)
	case uint:
		return int32(v)
	case uint8:
		return int32(v)
	case uint16:
		return int32(v)
	case uint32:
		return int32(v)
	case uint64:
		return int32(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := strconv.ParseInt(v, 0, 32)
		if err == nil {
			return int32(i)
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Int64 converts an interface{} of Result object to a int64
func (r *Result) Int64() int64 {
	switch v := r.object.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return int64(v)
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := strconv.ParseInt(v, 0, 64)
		if err == nil {
			return i
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Int converts an interface{} of Result object to a int
func (r *Result) Int() int {
	switch v := r.object.(type) {
	case int:
		return int(v)
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := strconv.ParseInt(v, 0, 0)
		if err == nil {
			return int(i)
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Float32 converts an interface{} of Result object to a float32.
func (r *Result) Float32() float32 {
	switch v := r.object.(type) {
	case int:
		return float32(v)
	case int8:
		return float32(v)
	case int16:
		return float32(v)
	case int32:
		return float32(v)
	case int64:
		return float32(v)
	case uint:
		return float32(v)
	case uint8:
		return float32(v)
	case uint16:
		return float32(v)
	case uint32:
		return float32(v)
	case uint64:
		return float32(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := strconv.ParseFloat(v, 32)
		if err == nil {
			return float32(i)
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// Float64 converts an interface{} of Result object to a float64.
func (r *Result) Float64() float64 {
	switch v := r.object.(type) {
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return i
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// String converts an interface{} of Result object to a string.
func (r *Result) String() string {
	switch v := r.object.(type) {
	case int:
		return strconv.Itoa(v)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case bool:
		if v {
			return "true"
		}
		return "false"
	case string:
		return v
	default:
		return ""
	}
}

// Bool converts an interface{} of Result object to a bool
func (r *Result) Bool() bool {
	switch v := r.object.(type) {
	case bool:
		return v
	case nil:
		return false
	case string:
		b, err := strconv.ParseBool(v)
		if err == nil {
			return b
		}
		return false
	default:
		return false
	}
}

// MapInterface converts an interface{} of Result object to a map[string]interface{}.
func (r *Result) MapInterface() map[string]interface{} {
	switch r.object.(type) {
	case map[string]interface{}:
		return r.object.(map[string]interface{})
	default:
		return map[string]interface{}{}
	}
}

// MapInterfaceSlice converts an interface{} of Result object to a []map[string]interface{}.
func (r *Result) MapInterfaceSlice() []map[string]interface{} {
	switch r.object.(type) {
	case []map[string]interface{}:
		return r.object.([]map[string]interface{})
	default:
		return []map[string]interface{}{}
	}
}
