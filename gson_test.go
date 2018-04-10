package gson

import (
	"reflect"
	"testing"
)

func TestNewGsonFromString(t *testing.T) {
	tests := []struct {
		json    string
		isError bool
	}{
		{
			json:    `1`,
			isError: false,
		},
		{
			json:    `"2"`,
			isError: false,
		},
		{
			json:    `{"picture": "http://hogehoge"}`,
			isError: false,
		},
		{
			json:    `{afsf: adfaasf`,
			isError: true,
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
			isError: false,
		},
		{
			json:    `[{"name": "little"}, {"name": "tiny"}]`,
			isError: false,
		},
		{
			json:    `[{"name": "litt]`,
			isError: true,
		},
	}

	for _, test := range tests {
		g, err := NewGsonFromString(test.json)

		isError := !(err == nil)

		if test.isError != isError {
			t.Error(isError)
			t.Errorf("NewGsonFromString isExpectedError: %v got: %v", test.isError, g)
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
		g, err := NewGsonFromString(test.json)
		if err != nil {
			t.Errorf("NewGsonFromString is error: %v", err)
		}

		result, err := g.GetByKeys(test.keys...)

		isError := !(err == nil)

		if test.isError != isError && reflect.DeepEqual(test.expected, result) {
			t.Errorf("Search isExpectedError: %v, expected object: %v, got: %v", test.isError, test.expected, result)
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

	for _, test := range tests {
		g, err := NewGsonFromString(test.json)
		if err != nil {
			t.Errorf("NewGsonFromString is error: %v", err)
		}

		result, err := g.GetByPath(test.path)

		isError := !(err == nil)

		if test.isError != isError && reflect.DeepEqual(test.expected, result) {
			t.Errorf("Search isExpectedError: %v, expected object: %v, got: %v", test.isError, test.expected, result)
		}
	}
}

func TestHasWithPath(t *testing.T) {
	tests := []struct {
		json string
		path string
		has  bool
	}{
		{
			json: `1`,
			path: "1",
			has:  false,
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
			path: "friends.2.id",
			has:  true,
		},
		{
			json: `
				{"name": "hlts2"}
			`,
			path: "nameeeee",
			has:  false,
		},
	}

	for _, test := range tests {
		g, err := NewGsonFromString(test.json)
		if err != nil {
			t.Errorf("NewGsonFromString is error: %v", err)
		}

		has := g.HasWithPath(test.path)

		if test.has != has {
			t.Errorf("HasWithPath expected: %v, got: %v", test.has, has)
		}
	}
}

func TestHasWithKeys(t *testing.T) {
	tests := []struct {
		json string
		keys []string
		has  bool
	}{
		{
			json: `1`,
			keys: []string{"1"},
			has:  false,
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
			keys: []string{"friends", "2", "id"},
			has:  true,
		},
		{
			json: `
				{"name": "hlts2"}
			`,
			keys: []string{"nameeeee"},
			has:  false,
		},
	}

	for _, test := range tests {
		g, err := NewGsonFromString(test.json)
		if err != nil {
			t.Errorf("NewGsonFromString is error: %v", err)
		}

		has := g.HasWithKeys(test.keys...)

		if test.has != has {
			t.Errorf("HasWithKeys expected: %v, got: %v", test.has, has)
		}
	}
}

func TestIndent(t *testing.T) {
	jsonString := `{"friends": [{"name": "hlts2"}, {"name": "hiroto"}]}`
	expectedJSONString := `{
 "friends": [
  {
   "name": "hlts2"
  },
  {
   "name": "hiroto"
  }
 ]
}`
	g, _ := NewGsonFromString(jsonString)

	gotJSONString, _ := g.Indent("", " ")
	if expectedJSONString != gotJSONString {
		t.Errorf("expected: %v, got: %v", expectedJSONString, gotJSONString)
	}
}

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

	for _, test := range tests {
		g, err := NewGsonFromString(test.json)
		if err != nil {
			t.Errorf("NewGsonFromString is error: %v", err)
		}

		result, err := g.GetByKeys("ID")
		if err != nil {
			t.Errorf("GetByKeys is error: %v", err)
		}

		got, err := result.Uint8()

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("expected isError: %v, got: %v", test.isError, isError)
		}

		if test.expected != got {
			t.Errorf("expected: %v, got: %v", test.expected, got)
		}
	}
}
