package gson

import (
	"reflect"
	"testing"
)

func TestNewGsonFromString(t *testing.T) {
	tests := []struct {
		json    string
		want    *Gson
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
			t.Errorf("NewGsonFromString(%s) isExpectedError: %v got: %v", test.json, test.isError, g)
		}
	}
}

func TestPath(t *testing.T) {
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
		g, _ := NewGsonFromString(test.json)

		result, err := g.Search(test.keys...)

		isError := !(err == nil)

		if test.isError != isError && reflect.DeepEqual(test.expected, result) {
			t.Errorf("Search(%v) isExpectedError: %v, expected: %v, got: %v", test.keys, test.isError, test.expected, result)
		}
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		json     string
		path     string
		expected *Result
		isError  bool
	}{
		{
			json:     `{"name": "hlts2"}`,
			path:     "/name",
			expected: &Result{"hlts2"},
			isError:  false,
		},
		{
			json:     `[{"name": "hlts2"}]`,
			path:     "/0",
			expected: &Result{map[string]interface{}{"name": "hlts2"}},
			isError:  false,
		},
		{
			json:     `[{"name": "hlts2"}]`,
			path:     "/10",
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
			path: "/friends",
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
			path:     "/friends/100/name",
			expected: nil,
			isError:  true,
		},
	}

	for _, test := range tests {
		g, _ := NewGsonFromString(test.json)

		result, err := g.Path(test.path)

		isError := !(err == nil)

		if test.isError != isError && reflect.DeepEqual(test.expected, result) {
			t.Errorf("Search(%v) isExpectedError: %v, expected: %v, got: %v", test.path, test.isError, test.expected, result)
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
			has:  true,
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
			path: "/friends/2/id",
			has:  true,
		},
		{
			json: `
				{"name": "hlts2"}
			`,
			path: "/nameeeee",
			has:  false,
		},
	}

	for _, test := range tests {
		g, _ := NewGsonFromString(test.json)

		has := g.HasWithPath(test.path)

		if test.has != has {
			t.Errorf("HasWithPath(%v) Expected: %v, got: %v", test.path, test.has, has)
		}
	}
}
