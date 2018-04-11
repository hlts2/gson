package gson

import (
	"reflect"
	"testing"
)

func TestNewGsonFromString(t *testing.T) {
	tests := []struct {
		json     string
		expected *Gson
		isError  bool
	}{
		{
			json: `1`,
			expected: &Gson{
				jsonObject: 1,
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
				jsonObject: map[string][]interface{}{
					"friends": []interface{}{
						map[string]interface{}{
							"id":   0,
							"name": "hiro",
						},
						map[string]interface{}{
							"id":   1,
							"name": "hiroto",
						},
						map[string]interface{}{
							"id":   2,
							"name": "hlts2",
						},
					},
				},
			},
			isError: false,
		},
		{
			json: `[{"name": "litt]`,
			expected: &Gson{
				jsonObject: []map[string]interface{}{
					map[string]interface{}{
						"name": "litt",
					},
				},
			},
			isError: true,
		},
	}

	for i, test := range tests {
		_, err := NewGsonFromString(test.json)

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("i = %d NewGsonFromString(json) expected isError: %v, got: %v", i, test.isError, isError)
		}

		// if !reflect.DeepEqual(g, test.expected) {
		// 	t.Errorf("i = %d NewGsonFromString(json) expected: %v, got: %v", i, test.expected.jsonObject, g.jsonObject)
		// }
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
		g, err := NewGsonFromString(test.json)
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
		g, err := NewGsonFromString(test.json)
		if err != nil {
			t.Errorf("i = %d NewGsonFromString() is error: %v", i, err)
		}

		if g == nil {
			t.Errorf("i = %d NewGsonFromString() g is nil", i)
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
		g, err := NewGsonFromString(test.json)
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
		g, err := NewGsonFromString(test.json)
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

func TestUint8(t *testing.T) {
	tests := []struct {
		json     string
		expected uint8
	}{
		{
			json:     `{"ID": 123}`,
			expected: uint8(123),
		},
	}

	for i, test := range tests {
		g, err := NewGsonFromString(test.json)
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

		got := result.Uint8()

		if test.expected != got {
			t.Errorf("i = %d Uint8() expected: %v, got: %v", i, test.expected, got)
		}
	}
}
