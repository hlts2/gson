package gson

import (
	"reflect"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		json     string
		want     *Gson
		hasError bool
	}{
		{
			json: `1`,
			want: &Gson{
				object: float64(1),
			},
			hasError: false,
		},
		{
			json: `"2"`,
			want: &Gson{
				object: "2",
			},
			hasError: false,
		},
		{
			json: `{"picture": "http://hogehoge"}`,
			want: &Gson{
				object: map[string]interface{}{
					"picture": "http://hogehoge",
				},
			},
			hasError: false,
		},
		{
			json:     `{afsf: adfaasf`,
			want:     nil,
			hasError: true,
		},
		{
			json: `
				{"friends": [
     					{
        					"id": 0,
							"name": "hiro"
						},
						{
							"id": 1,
							"name": "hiroto"
						},
						{
							"id": 2,
							"name": "hlts2"
						}
					]}
	  			`,
			want: &Gson{
				object: map[string]interface{}{
					"friends": []interface{}{
						map[string]interface{}{
							"id":   float64(0),
							"name": "hiro",
						},
						map[string]interface{}{
							"id":   float64(1),
							"name": "hiroto",
						},
						map[string]interface{}{
							"id":   float64(2),
							"name": "hlts2",
						},
					},
				},
			},
			hasError: false,
		},
		{
			json:     `[{"name": "litt]`,
			want:     nil,
			hasError: true,
		},
	}

	for i, test := range tests {
		t.Run("CreateWithBytes", func(t *testing.T) {
			g, err := CreateWithBytes([]byte(test.json))

			hasError := !(err == nil)

			if test.hasError != hasError {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.hasError, hasError)
			}

			if !reflect.DeepEqual(g, test.want) {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.want, g)
			}
		})

		t.Run("CreateWithReader", func(t *testing.T) {
			g, err := CreateWithReader(strings.NewReader(test.json))

			hasError := !(err == nil)

			if test.hasError != hasError {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.hasError, hasError)
			}

			if !reflect.DeepEqual(g, test.want) {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.want, g)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		json     string
		keys     []string
		want     *Result
		hasError bool
	}{
		{
			json:     `{"name": "hlts2"}`,
			keys:     []string{"name"},
			want:     &Result{"hlts2"},
			hasError: false,
		},
		{
			json:     `[{"name": "hlts2"}]`,
			keys:     []string{"0"},
			want:     &Result{map[string]interface{}{"name": "hlts2"}},
			hasError: false,
		},
		{
			json:     `[{"name": "hlts2"}]`,
			keys:     []string{"10"},
			want:     nil,
			hasError: true,
		},
		{
			json: `
				{"friends": [
      				{
						"id": "0",
						"name": "hiro"
					},
      				{
						"id": "1",
						"name": "hlts2"
					}
				]}
			`,
			keys: []string{"friends"},
			want: &Result{[]interface{}{
				map[string]interface{}{"id": "0", "name": "hiro"},
				map[string]interface{}{"id": "1", "name": "hlts2"}},
			},
			hasError: false,
		},
		{
			json: `
				{"friends": [
      				{
        				"id": 0,
						"name": "hiro"
					},
					{
						"id": 1,
						"name": "hiroto"
					},
					{
						"id": 2,
						"name": "hlts2"
					}
				]}
			`,
			keys:     []string{"friends", "100", "name"},
			want:     nil,
			hasError: true,
		},
	}

	for i, test := range tests {
		g, err := CreateWithBytes([]byte(test.json))
		if err != nil {
			t.Errorf("CreateWithBytes returned error: %v", err)
		}

		if g == nil {
			t.Error("CreateWithBytes returned nil")
		}

		t.Run("GetByKeys", func(t *testing.T) {
			r, err := g.GetByKeys(test.keys...)

			hasError := !(err == nil)

			if test.hasError != hasError {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.hasError, hasError)
			}

			if !reflect.DeepEqual(test.want, r) {
				t.Errorf("tests[%d] - want: %v, got: %v", i, test.want, r)
			}
		})

		t.Run("GetByPath", func(t *testing.T) {
			r, err := g.GetByPath(strings.Join(test.keys, "."))

			hasError := !(err == nil)

			if test.hasError != hasError {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.hasError, hasError)
			}

			if !reflect.DeepEqual(test.want, r) {
				t.Errorf("tests[%d] - want: %v, got: %v", i, test.want, r)
			}
		})
	}
}

//
// func TestUint8(t *testing.T) {
// 	tests := []struct {
// 		json     string
// 		expected uint8
// 		isError  bool
// 	}{
// 		{
// 			json:     `{"ID": 123}`,
// 			expected: uint8(123),
// 			isError:  false,
// 		},
// 	}
//
// 	for i, test := range tests {
// 		g, err := NewGsonFromByte([]byte(test.json))
// 		if err != nil {
// 			t.Errorf("i = %d NewGsonFromString(json) is error: %v", i, err)
// 		}
//
// 		if g == nil {
// 			t.Errorf("i = %d NewGsonFromString(json) g is nil", i)
// 		}
//
// 		result, err := g.GetByKeys("ID")
// 		if err != nil {
// 			t.Errorf("i = %d GetByKeys(keys) is error: %v", i, err)
// 		}
//
// 		got, err := result.Uint8()
//
// 		isError := !(err == nil)
//
// 		if test.isError != isError {
// 			t.Errorf("i = %d Uint8() expected isError: %v, got: %v", i, test.isError, isError)
// 		}
//
// 		if test.expected != got {
// 			t.Errorf("i = %d GetByKeys(keys) expected: %v, got: %v", i, test.expected, got)
// 		}
// 	}
// }
//
// func TestSlice(t *testing.T) {
// 	tests := []struct {
// 		json     string
// 		expected []*Result
// 		isError  bool
// 	}{
// 		{
// 			json: `{"IDs": [{"ID": "1111"}, {"ID": "2222"}]}`,
// 			expected: []*Result{
// 				{
// 					object: map[string]interface{}{
// 						"ID": "1111",
// 					},
// 				},
// 				{
// 					object: map[string]interface{}{
// 						"ID": "2222",
// 					},
// 				},
// 			},
// 			isError: false,
// 		},
// 	}
//
// 	for ti, test := range tests {
// 		g, err := NewGsonFromByte([]byte(test.json))
// 		if err != nil {
// 			t.Errorf("i = %d NewGsonFromString(json) is error: %v", ti, err)
// 		}
//
// 		if g == nil {
// 			t.Errorf("i = %d NewGsonFromString(json) g is nil", ti)
// 		}
//
// 		result, err := g.GetByKeys("IDs")
// 		if err != nil {
// 			t.Errorf("i = %d GetByKeys(keys) is error: %v", ti, err)
// 		}
//
// 		slice, err := result.Slice()
//
// 		isError := !(err == nil)
//
// 		if test.isError != isError {
// 			t.Errorf("i = %d Slice() expected isError: %v, got: %v", ti, test.isError, isError)
// 		}
//
// 		for si := range slice {
// 			if !reflect.DeepEqual(test.expected[si].object, slice[si].object) {
// 				t.Errorf("i = %d Slice() expected Result: %v, got: %v", ti, test.expected[si].object, slice[si].object)
// 			}
// 		}
// 	}
// }
//
// func TestMap(t *testing.T) {
// 	tests := []struct {
// 		json     string
// 		expected map[string]*Result
// 		isError  bool
// 	}{
// 		{
// 			json: `{"Accounts": [{"ID": "1111", "Name": "hlts2"}]}`,
// 			expected: map[string]*Result{
// 				"ID": {
// 					object: "1111",
// 				},
// 				"Name": {
// 					object: "hlts2",
// 				},
// 			},
// 			isError: false,
// 		},
// 	}
//
// 	for i, test := range tests {
// 		g, err := NewGsonFromByte([]byte(test.json))
// 		if err != nil {
// 			t.Errorf("i = %d NewGsonFromString(json) is error: %v", i, err)
// 		}
//
// 		if g == nil {
// 			t.Errorf("i = %d NewGsonFromString(json) g is nil", i)
// 		}
//
// 		result, err := g.GetByKeys("Accounts", "0")
//
// 		if err != nil {
// 			t.Errorf("i = %d GetByKeys(keys) is error: %v", i, err)
// 		}
//
// 		m, err := result.Map()
//
// 		isError := !(err == nil)
//
// 		if test.isError != isError {
// 			t.Errorf("i = %d Map() expected isError: %v, got: %v", i, test.isError, isError)
// 		}
//
// 		if !reflect.DeepEqual(test.expected, m) {
// 			t.Errorf("i = %d Map() expected: %v, got: %v", i, test.expected, m)
// 		}
// 	}
// }
