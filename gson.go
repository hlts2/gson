package gson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/spf13/cast"
)

// Represents an error when the search fails for the value of JSON.
var (
	ErrorIndexOutOfRange = errors.New("index out of range")
	ErrorInvalidJSONKey  = errors.New("invalid json Key")
)

type (
	// Gson is gson base structor. Gson wrapps JSON object.
	Gson struct {
		object interface{}
	}

	// Result represents a JSON value.
	Result struct {
		object interface{}
	}
)

// CreateWithBytes creates a Gson object with []byte.
func CreateWithBytes(data []byte) (*Gson, error) {
	g := new(Gson)

	err := ffjson.Unmarshal(data, &g.object)
	if err != nil {
		return nil, errors.Wrap(err, "faild to unmarshal")
	}

	return g, nil
}

// CreateWithReader creates Gson object with io.Reader.
func CreateWithReader(reader io.Reader) (*Gson, error) {
	g := new(Gson)

	err := ffjson.NewDecoder().DecodeReader(reader, &g.object)
	if err != nil {
		return nil, errors.Wrap(err, "faild to decode")
	}

	return g, nil
}

// Object returns JSON object wrapped by Gson.
func (g *Gson) Object() interface{} {
	return g.object
}

// Indent appends indent to dst an indented form of the JSON object wrapped by Gson.
func (g *Gson) Indent(dst *bytes.Buffer, prefix, indent string) error {
	return errors.Wrap(indentJSON(dst, g.object, prefix, indent), "faild to create indent")
}

func indentJSON(dst *bytes.Buffer, object interface{}, prefix, indent string) error {
	var src bytes.Buffer
	err := ffjson.NewEncoder(&src).Encode(object)
	if err != nil {
		return errors.Wrap(err, "faild to encode")
	}

	err = json.Indent(dst, src.Bytes(), prefix, indent)
	if err != nil {
		return errors.Wrap(err, "faild to appends indent")
	}
	return nil
}

// GetByKeys returns JSON value corresponding to keys. keys represents key of hierarchy of json.
// e.g)
// {
//     "key-1": {
//         "key-2": {
//             "key-3": [
//                 "hello",
//                 "world",
//             ]
//         }
//     }
// }
// key-1, key-2, key-3, 0 => "hello"
// key-1, key-2, key-3, # => "hello", "world"
func (g *Gson) GetByKeys(keys ...string) (*Result, error) {
	r, err := getByKeys(keys, g.object)
	if err != nil {
		return nil, errors.Wrap(err, "faild to get JSON value")
	}
	return r, nil
}

// GetByPath returns JSON value corresponding to the path.
// e.g)
// {
//     "key-1": {
//         "key-2": {
//             "key-3": [
//                 "hello",
//                 "world",
//             ]
//         }
//     }
// }
// key-1.key-2.key-3.0 => "hello"
// key-1, key-2, key-3, # => "hello", "world"
func (g *Gson) GetByPath(path string) (*Result, error) {
	r, err := getByKeys(strings.Split(path, "."), g.object)
	if err != nil {
		return nil, errors.Wrap(err, "faild to get JSON value")
	}
	return r, nil
}

func getByKeys(keys []string, object interface{}) (*Result, error) {
	for i, key := range keys {
		if m, ok := object.(map[string]interface{}); ok {
			if object, ok = m[key]; !ok {
				return nil, ErrorInvalidJSONKey
			}
		} else if s, ok := object.([]interface{}); ok {
			if key == "#" {
				r, err := getByKeysFromSlice(keys[i+1:], s)
				if err != nil {
					return nil, errors.Wrap(err, "faild to get JSON value from slice")
				}
				return r, nil
			}

			idx, err := strconv.Atoi(key)
			if err != nil {
				return nil, ErrorInvalidJSONKey
			}

			if idx < 0 || idx >= len(s) {
				return nil, ErrorIndexOutOfRange
			}
			object = s[idx]
		}
	}

	return &Result{object}, nil
}

func getByKeysFromSlice(keys []string, s []interface{}) (*Result, error) {
	if len(keys) == 0 {
		return &Result{s}, nil
	}

	objects := make([]interface{}, 0, len(s))
	for _, v := range s {
		r, err := getByKeys(keys, v)
		if err != nil {
			return nil, errors.Wrap(err, "faild to get JSON value")
		}

		objects = append(objects, r.Interface())
	}

	return &Result{objects}, nil
}

// Result cast object wrapped by Result to object wrapped by Gson.
func (g *Gson) Result() *Result {
	return &Result{g.object}
}

// Indent appends indent to dst an indented form of the JSON object wrapped by Result.
func (r *Result) Indent(buf *bytes.Buffer, prefix, indent string) error {
	err := indentJSON(buf, r.object, prefix, indent)
	if err != nil {
		return errors.Wrap(err, "faild to create indent")
	}
	return nil
}

// Interface returns JSON object wrapped by Result.
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
	v, err := r.StringE()
	if err != nil {
		return fmt.Sprintf("%v", r.object)
	}
	return v
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
func (r *Result) SliceE() ([]*Result, error) {
	v, err := cast.ToSliceE(r.object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return toSlice(v), nil
}

// Slice casts an interface to a []Result type.
func (r *Result) Slice() []*Result {
	v := cast.ToSlice(r.object)
	return toSlice(v)
}

func toSlice(s []interface{}) []*Result {
	results := make([]*Result, len(s))

	for i, v := range s {
		results[i] = &Result{v}
	}
	return results
}

// MapE casts an map[string]Result type and returns an error if types don't match.
func (r *Result) MapE() (map[string]*Result, error) {
	v, err := cast.ToStringMapE(r.object)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return toMap(v), nil
}

// Map casts an interface to a map[string]Resul type.
func (r *Result) Map() map[string]*Result {
	v := cast.ToStringMap(r.object)
	return toMap(v)
}

func toMap(m map[string]interface{}) map[string]*Result {
	results := make(map[string]*Result, len(m))
	for k, v := range m {
		results[k] = &Result{v}
	}
	return results
}

// Gson cast object wrapped by Gson to object wrapped by Result.
func (r *Result) Gson() *Gson {
	return &Gson{r.object}
}
