package gson

import (
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	json := `
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
		`

	g, err := NewGsonFromByte([]byte(json))
	if err != nil {
		t.Errorf("NewGsonFromByte(json) is error: %v", err)
	}

	if g == nil {
		t.Error("NewGsonFromByte(json) g is nil")
	}

	result, err := g.GetByKeys("friends")
	if err != nil {
		t.Errorf("GetByKeys(keys) is error: %v", err)
	}

	if result == nil {
		t.Errorf("GetByKeys(keys) result is nil")
	}

	slice, err := result.Slice()
	if err != nil {
		t.Errorf("Slice() is error: %v", err)
	}

	if slice == nil {
		t.Error("Slice() is nil")
	}

	for _, value := range slice {
		m, err := value.Map()
		if err != nil {
			t.Errorf("%v.Map() is error: %v", value, err)
		}

		_, err = m["name"].String()
		if err != nil {
			t.Errorf("%v.String() is error: %v", m["name"], err)
		}
	}

}

func TestNewGsonFromByte(t *testing.T) {
	tests := []struct {
		json     string
		expected *Gson
		isError  bool
	}{
		{
			json: `1`,
			expected: &Gson{
				jsonObject: float64(1),
			},
			isError: false,
		},
		{
			json: `"2"`,
			expected: &Gson{
				jsonObject: "2",
			},
			isError: false,
		},
		{
			json: `{"picture": "http://hogehoge"}`,
			expected: &Gson{
				jsonObject: map[string]interface{}{
					"picture": "http://hogehoge",
				},
			},
			isError: false,
		},
		{
			json:     `{afsf: adfaasf`,
			expected: nil,
			isError:  true,
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
			expected: &Gson{
				jsonObject: map[string]interface{}{
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
			isError: false,
		},
		{
			json:     `[{"name": "litt]`,
			expected: nil,
			isError:  true,
		},
	}

	for i, test := range tests {
		g, err := NewGsonFromByte([]byte(test.json))

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("i = %d NewGsonFromString(json) expected isError: %v, got: %v", i, test.isError, isError)
		}

		if !reflect.DeepEqual(g, test.expected) {
			t.Errorf("i = %d NewGsonFromString(json) expected: %v, got: %v", i, test.expected, g)
		}
	}
}

func TestGetByKeys(t *testing.T) {
	tests := []struct {
		json     string
		keys     []string
		expected *Result
		isError  bool
	}{
		{
			json:     `{"name": "hlts2"}`,
			keys:     []string{"name"},
			expected: &Result{"hlts2"},
			isError:  false,
		},
		{
			json:     `[{"name": "hlts2"}]`,
			keys:     []string{"0"},
			expected: &Result{map[string]interface{}{"name": "hlts2"}},
			isError:  false,
		},
		{
			json:     `[{"name": "hlts2"}]`,
			keys:     []string{"10"},
			expected: nil,
			isError:  true,
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
			expected: &Result{[]interface{}{
				map[string]interface{}{"id": "0", "name": "hiro"},
				map[string]interface{}{"id": "1", "name": "hlts2"}},
			},
			isError: false,
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
			expected: nil,
			isError:  true,
		},
	}

	for _, test := range tests {
		g, err := NewGsonFromByte([]byte(test.json))
		if err != nil {
			t.Errorf("NewGsonFromString(json) is error: %v", err)
		}

		if g == nil {
			t.Error("NewGsonFromString(json) g is nil")
		}

		result, err := g.GetByKeys(test.keys...)

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("GetByKeys(keys) expected isError: %v, got: %v", test.isError, isError)
		}

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("GetByKeys(keys) expected: %v, got: %v", test.expected, result)
		}
	}
}

func TestGetByPath(t *testing.T) {
	tests := []struct {
		json     string
		path     string
		expected *Result
		isError  bool
	}{
		{
			json:     `{"name": "hlts2"}`,
			path:     "name",
			expected: &Result{"hlts2"},
			isError:  false,
		},
		{
			json:     `[{"name": "hlts2"}]`,
			path:     "10",
			expected: nil,
			isError:  true,
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
			path: "friends",
			expected: &Result{[]interface{}{
				map[string]interface{}{"id": "0", "name": "hiro"},
				map[string]interface{}{"id": "1", "name": "hlts2"}},
			},
			isError: false,
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
			path:     "friends.100.name",
			expected: nil,
			isError:  true,
		},
	}

	for i, test := range tests {
		g, err := NewGsonFromByte([]byte(test.json))
		if err != nil {
			t.Errorf("i = %d NewGsonFromString(json) is error: %v", i, err)
		}

		if g == nil {
			t.Errorf("i = %d NewGsonFromString(json) g is nil", i)
		}

		result, err := g.GetByPath(test.path)

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("i = %d GetByPath(path) expected isError: %v, got: %v", i, test.isError, isError)
		}

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("i = %d GetByPath(path) expected: %v, got: %v", i, test.expected, result)
		}
	}
}

/*
func TestHasWithPath(t *testing.T) {
	tests := []struct {
		json     string
		path     string
		expected bool
	}{
		{
			json:     `1`,
			path:     "1",
			expected: false,
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
			path:     "friends.2.id",
			expected: true,
		},
		{
			json: `
				{"name": "hlts2"}
			`,
			path:     "nameeeee",
			expected: false,
		},
	}

	for i, test := range tests {
		g, err := NewGsonFromByte([]byte(test.json))
		if err != nil {
			t.Errorf("i = %d NewGsonFromString(json) is error: %v", i, err)
		}

		if g == nil {
			t.Errorf("i = %d NewGsonFromString(json) g is nil", i)
		}

		got := g.HasWithPath(test.path)

		if test.expected != got {
			t.Errorf("i = %d HasWithPath(path) expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestHasWithKeys(t *testing.T) {
	tests := []struct {
		json     string
		keys     []string
		expected bool
	}{
		{
			json:     `1`,
			keys:     []string{"1"},
			expected: false,
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
			keys:     []string{"friends", "2", "id"},
			expected: true,
		},
		{
			json: `
				{"name": "hlts2"}
			`,
			keys:     []string{"nameeeee"},
			expected: false,
		},
	}

	for i, test := range tests {
		g, err := NewGsonFromByte([]byte(test.json))
		if err != nil {
			t.Errorf("i = %d NewGsonFromString(json) is error: %v", i, err)
		}

		if g == nil {
			t.Errorf("i = %d NewGsonFromString(json) g is nil", i)
		}

		got := g.HasWithKeys(test.keys...)

		if test.expected != got {
			t.Errorf("i = %d HasWithKeys(keys) expected: %v, got: %v", i, test.expected, got)
		}
	}
}
*/

func TestUint8(t *testing.T) {
	tests := []struct {
		json     string
		expected uint8
		isError  bool
	}{
		{
			json:     `{"ID": 123}`,
			expected: uint8(123),
			isError:  false,
		},
	}

	for i, test := range tests {
		g, err := NewGsonFromByte([]byte(test.json))
		if err != nil {
			t.Errorf("i = %d NewGsonFromString(json) is error: %v", i, err)
		}

		if g == nil {
			t.Errorf("i = %d NewGsonFromString(json) g is nil", i)
		}

		result, err := g.GetByKeys("ID")
		if err != nil {
			t.Errorf("i = %d GetByKeys(keys) is error: %v", i, err)
		}

		got, err := result.Uint8()

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("i = %d Uint8() expected isError: %v, got: %v", i, test.isError, isError)
		}

		if test.expected != got {
			t.Errorf("i = %d GetByKeys(keys) expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		json     string
		expected []*Result
		isError  bool
	}{
		{
			json: `{"IDs": [{"ID": "1111"}, {"ID": "2222"}]}`,
			expected: []*Result{
				{
					object: map[string]interface{}{
						"ID": "1111",
					},
				},
				{
					object: map[string]interface{}{
						"ID": "2222",
					},
				},
			},
			isError: false,
		},
	}

	for ti, test := range tests {
		g, err := NewGsonFromByte([]byte(test.json))
		if err != nil {
			t.Errorf("i = %d NewGsonFromString(json) is error: %v", ti, err)
		}

		if g == nil {
			t.Errorf("i = %d NewGsonFromString(json) g is nil", ti)
		}

		result, err := g.GetByKeys("IDs")
		if err != nil {
			t.Errorf("i = %d GetByKeys(keys) is error: %v", ti, err)
		}

		slice, err := result.Slice()

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("i = %d Slice() expected isError: %v, got: %v", ti, test.isError, isError)
		}

		for si := range slice {
			if !reflect.DeepEqual(test.expected[si].object, slice[si].object) {
				t.Errorf("i = %d Slice() expected Result: %v, got: %v", ti, test.expected[si].object, slice[si].object)
			}
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		json     string
		expected map[string]*Result
		isError  bool
	}{
		{
			json: `{"Accounts": [{"ID": "1111", "Name": "hlts2"}]}`,
			expected: map[string]*Result{
				"ID": &Result{
					object: "1111",
				},
				"Name": &Result{
					object: "hlts2",
				},
			},
			isError: false,
		},
	}

	for i, test := range tests {
		g, err := NewGsonFromByte([]byte(test.json))
		if err != nil {
			t.Errorf("i = %d NewGsonFromString(json) is error: %v", i, err)
		}

		if g == nil {
			t.Errorf("i = %d NewGsonFromString(json) g is nil", i)
		}

		result, err := g.GetByKeys("Accounts", "0")

		if err != nil {
			t.Errorf("i = %d GetByKeys(keys) is error: %v", i, err)
		}

		m, err := result.Map()

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("i = %d Map() expected isError: %v, got: %v", i, test.isError, isError)
		}

		if !reflect.DeepEqual(test.expected, m) {
			t.Errorf("i = %d Map() expected: %v, got: %v", i, test.expected, m)
		}
	}
}
