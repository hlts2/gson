package gson

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/spf13/cast"
)

// Represents errors for searching json's value
var (
	ErrorIndexOutOfRange = errors.New("index out of range")
	ErrorInvalidJSONKey  = errors.New("invalid json Key")
	ErrorInvalidObject   = errors.New("invalid object")
)

// Result represents a json value that is returned from GetByPath() and GetByKeys().
type Result struct {
	object interface{}
}

// Gson is gson base structor
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

// NewGsonFromInterface returns Gson instance created from interface
func NewGsonFromInterface(object interface{}) (*Gson, error) {
	g := new(Gson)

	if !isJSONObject(object) {
		return nil, ErrorInvalidObject
	}

	g.jsonObject = object
	return g, nil
}

// Interface returns json object
func (g *Gson) Interface() interface{} {
	return g.jsonObject
}

// Indent converts json object to json string
func (g *Gson) Indent(dist *bytes.Buffer, prefix, indent string) error {
	return indentJSON(dist, g.jsonObject, prefix, indent)
}

func isJSONObject(object interface{}) bool {
	if _, err := ffjson.Marshal(object); err == nil {
		return true
	}
	return false
}

func indentJSON(dist *bytes.Buffer, object interface{}, prefix, indent string) error {
	var src bytes.Buffer
	err := ffjson.NewEncoder(&src).Encode(object)
	if err != nil {
		return err
	}

	err = json.Indent(dist, src.Bytes(), prefix, indent)
	if err != nil {
		return err
	}
	return nil
}

// GetByKeys returns json value corresponding to keys. keys represents key of hierarchy of json
func (g *Gson) GetByKeys(keys ...string) (*Result, error) {
	return g.getByKeys(keys)
}

// GetByPath returns json value corresponding to path.
func (g *Gson) GetByPath(path string) (*Result, error) {
	return g.getByKeys(strings.Split(path, "."))
}

func (g *Gson) getByKeys(keys []string) (*Result, error) {
	jsonObject := g.jsonObject

	for _, key := range keys {
		if mmap, ok := jsonObject.(map[string]interface{}); ok {
			if val, ok := mmap[key]; ok {
				jsonObject = val
				continue
			}
		} else if marray, ok := jsonObject.([]interface{}); ok {
			idx64, err := strconv.ParseInt(key, 10, 0)
			idx := int(idx64)
			if err == nil {
				if idx >= 0 && idx < len(marray) {
					jsonObject = marray[idx]
					continue
				} else {
					return nil, ErrorIndexOutOfRange
				}
			}
		}
		return nil, ErrorInvalidJSONKey
	}

	return &Result{jsonObject}, nil
}

func (g *Gson) Result() *Result {
	return &Result{object: g.jsonObject}
}

// Indent converts json object to json buffer
func (r *Result) Indent(buf *bytes.Buffer, prefix, indent string) error {
	err := indentJSON(buf, r.object, prefix, indent)
	if err != nil {
		return err
	}
	return nil
}

// Interface returns json object of Result
func (r *Result) Interface() interface{} {
	return r.object
}

// Uint8E casts an interface to a uint8 type and returns an error if types don't match.
func (r *Result) Uint8E() (uint8, error) {
	v, err := cast.ToUint8E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Uint8 casts an interface to a uint8 type.
func (r *Result) Uint8() uint8 {
	return cast.ToUint8(r.object)
}

// Uint16E casts an interface to a uint16 type and returns an error if types don't match.
func (r *Result) Uint16E() (uint16, error) {
	v, err := cast.ToUint16E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Uint16 casts an interface to a uint16 type.
func (r *Result) Uint16() uint16 {
	return cast.ToUint16(r.object)
}

// Uint32E casts an interface{} to a uint32 type and returns an error if types don't match.
func (r *Result) Uint32E() (uint32, error) {
	v, err := cast.ToUint32E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Uint32 casts an interface to a uint32 type.
func (r *Result) Uint32() uint32 {
	return cast.ToUint32(r.object)
}

// Uint64E casts an interface to a uint64 type and returns an error if types don't match.
func (r *Result) Uint64E() (uint64, error) {
	v, err := cast.ToUint64E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Uint64 casts an interface to a uint64 type.
func (r *Result) Uint64() uint64 {
	return cast.ToUint64(r.object)
}

// Int8E casts an interface to a int8 type and returns an error if types don't match.
func (r *Result) Int8E() (int8, error) {
	v, err := cast.ToInt8E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Int8 casts an interface to a int8 type.
func (r *Result) Int8() int8 {
	return cast.ToInt8(r.object)
}

// Int16E casts an interface to a int16 type and returns an error if types don't match.
func (r *Result) Int16E() (int16, error) {
	v, err := cast.ToInt16E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Int16 casts an interface to a int16 type.
func (r *Result) Int16() int16 {
	return cast.ToInt16(r.object)
}

// Int32E casts an interface to a int32 type and returns an error if types don't match.
func (r *Result) Int32E() (int32, error) {
	v, err := cast.ToInt32E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Int32 casts an interface to a int32 type.
func (r *Result) Int32() int32 {
	return cast.ToInt32(r.object)
}

// Int64E casts an interface to a int64 type and returns an error if types don't match.
func (r *Result) Int64E() (int64, error) {
	v, err := cast.ToInt64E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Int64 casts an interface to a int64 type.
func (r *Result) Int64() int64 {
	return cast.ToInt64(r.object)
}

// IntE casts an interface to a int type and returns an error if types don't match.
func (r *Result) IntE() (int, error) {
	v, err := cast.ToIntE(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Int casts an interface to a int type.
func (r *Result) Int() int {
	return cast.ToInt(r.object)
}

// Float32E casts an interface to a float32 type and returns an error if types don't match.
func (r *Result) Float32E() (float32, error) {
	v, err := cast.ToFloat32E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Float32 casts an interface to a float32 type.
func (r *Result) Float32() float32 {
	return cast.ToFloat32(r.object)
}

// Float64E casts an interface to a float64 type and returns an error if types don't match.
func (r *Result) Float64E() (float64, error) {
	v, err := cast.ToFloat64E(r.object)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return v, nil
}

// Float64 casts an interface to a float64 type.
func (r *Result) Float64() float64 {
	return cast.ToFloat64(r.object)
}

// StringE casts an interface to a string type and returns an error if types don't match.
func (r *Result) StringE() (string, error) {
	v, err := cast.ToStringE(r.object)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return v, nil
}

// String casts an interface to a string type.
func (r *Result) String() string {
	return cast.ToString(r.object)
}

// BoolE casts an interface to a bool type and returns an error if types don't match.
func (r *Result) BoolE() (bool, error) {
	v, err := cast.ToBoolE(r.object)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return v, nil
}

// Bool casts an interface to a bool type.
func (r *Result) Bool() bool {
	return cast.ToBool(r.object)
}

// SliceE casts an []Result type and returns an error if types don't match.
func (r *Result) SliceE() ([]Result, error) {
	v, err := cast.ToSliceE(r.object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return toSlice(v), nil
}

// Slice casts an interface to a []Result type.
func (r *Result) Slice() []Result {
	v := cast.ToSlice(r.object)
	return toSlice(v)
}

func toSlice(s []interface{}) []Result {
	results := make([]Result, len(s))

	for i, v := range s {
		results[i] = Result{
			object: v,
		}
	}
	return results
}

// MapE casts an map[string]Result type and returns an error if types don't match.
func (r *Result) MapE() (map[string]Result, error) {
	v, err := cast.ToStringMapE(r.object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return toMap(v), nil
}

// Map casts an interface to a map[string]Resul type.
func (r *Result) Map() map[string]Result {
	v := cast.ToStringMap(r.object)
	return toMap(v)
}

func toMap(m map[string]interface{}) map[string]Result {
	results := make(map[string]Result, len(m))
	for k, v := range m {
		results[k] = Result{
			object: v,
		}
	}
	return results
}

// Gson casts an interface to *Gson type.
func (r *Result) Gson() *Gson {
	return &Gson{jsonObject: r.object}
}
