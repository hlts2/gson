package gson

import (
	"reflect"
	"strings"
	"testing"
)

type TestData struct {
	json       string
	isGoson    bool
	path       string
	keys       []string
	result     *Result
	objectType reflect.Kind
	isErr      bool
}

type TestDatas []TestData

func (n TestData) CheckError(err error) bool {
	if n.isErr && err != nil {
		return true
	}

	if !n.isErr && err == nil {
		return true
	}

	return false
}

func (n TestData) CheckGosonInstance(g *Goson) bool {
	if n.isGoson && g != nil {
		return true
	}

	if !n.isGoson && g == nil {
		return true
	}
	return false
}

func (n TestData) CheckResultObject(result *Result) bool {
	if result == nil && n.result == nil {
		return true
	}

	if reflect.DeepEqual(n.result.object, result.object) {
		return true
	}
	return false
}

func (n TestData) CheckResultConvert() error {
	var err error

	switch n.objectType {
	case reflect.Uint8:
		_, err = n.result.Uint8()
	case reflect.Uint16:
		_, err = n.result.Uint16()
	case reflect.Uint32:
		_, err = n.result.Uint32()
	case reflect.Uint64:
		_, err = n.result.Uint64()
	case reflect.Int8:
		_, err = n.result.Int8()
	case reflect.Int16:
		_, err = n.result.Int16()
	case reflect.Int32:
		_, err = n.result.Int32()
	case reflect.Int64:
		_, err = n.result.Int64()
	case reflect.Int:
		_, err = n.result.Int()
	case reflect.Float32:
		_, err = n.result.Float32()
	case reflect.Float64:
		_, err = n.result.Float64()
	case reflect.Complex64:
		_, err = n.result.Complex64()
	case reflect.Complex128:
		_, err = n.result.Complex128()
	case reflect.String:
		_, err = n.result.String()
	}
	return err
}

var newTests = TestDatas{
	TestData{
		json:    `1`,
		isGoson: true,
		isErr:   false,
	},
	TestData{
		json:    `"2"`,
		isGoson: true,
		isErr:   false,
	},
	TestData{
		json:    `{"picture": "http://hogehoge"}`,
		isGoson: true,
		isErr:   false,
	},
	TestData{
		json:    `{afsf: adfaasf`,
		isGoson: false,
		isErr:   true,
	},
	TestData{
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
		isGoson: true,
		isErr:   false,
	},
	TestData{
		json:    `[{"name": "little"}, {"name": "tiny"}]`,
		isGoson: true,
		isErr:   false,
	},
	TestData{
		json:    `[{"name": "litt]`,
		isGoson: false,
		isErr:   true,
	},
}

func TestNewGosonFromString(t *testing.T) {
	for _, test := range newTests {
		g, err := NewGosonFromString(test.json)

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckGosonInstance(g) {
			t.Error("")
		}
	}
}

func TestNewGosonFromByte(t *testing.T) {
	for _, test := range newTests {
		g, err := NewGosonFromByte([]byte(test.json))

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckGosonInstance(g) {
			t.Error("")
		}
	}
}

func TestNewGosonFromReader(t *testing.T) {
	for _, test := range newTests {
		g, err := NewGosonFromReader(strings.NewReader(test.json))

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckGosonInstance(g) {
			t.Error("")
		}
	}
}

var searchTests = TestDatas{
	TestData{
		json:   `{"name": "hlts2"}`,
		keys:   []string{"name"},
		path:   "/name",
		result: &Result{"hlts2"},
		isErr:  false,
	},
	TestData{
		json:   `[{"name": "hlts2"}]`,
		keys:   []string{"0"},
		path:   "/0",
		result: &Result{map[string]interface{}{"name": "hlts2"}},
		isErr:  false,
	},
	TestData{
		json:   `[{"name": "hlts2"}]`,
		keys:   []string{"10"},
		path:   "/10",
		result: nil,
		isErr:  true,
	},
	TestData{
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
		keys:   []string{"friends"},
		path:   "/friends",
		result: &Result{[]interface{}{map[string]interface{}{"id": "0", "name": "hiro"}, map[string]interface{}{"id": "1", "name": "hlts2"}}},
		isErr:  false,
	},
	TestData{
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
		keys:   []string{"friends", "100", "name"},
		path:   "/friends/100/name",
		result: nil,
		isErr:  true,
	},
}

func TestSearch(t *testing.T) {
	for _, test := range searchTests {
		g, _ := NewGosonFromString(test.json)

		result, err := g.Search(test.keys...)

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckResultObject(result) {

		}
	}
}

func TestPath(t *testing.T) {
	for _, test := range searchTests {
		g, _ := NewGosonFromString(test.json)

		result, err := g.Path(test.path)

		if !test.CheckError(err) {
			t.Error("")
		}

		if !test.CheckResultObject(result) {
			t.Error("")
		}
	}
}

var resultTests = TestDatas{
	TestData{
		result:     &Result{10},
		objectType: reflect.Uint8,
		isErr:      false,
	},
	TestData{
		result:     &Result{10},
		objectType: reflect.Uint16,
		isErr:      false,
	},
	TestData{
		result:     &Result{10},
		objectType: reflect.Uint32,
		isErr:      false,
	},
	TestData{
		result:     &Result{10},
		objectType: reflect.Uint64,
		isErr:      false,
	},
	TestData{
		result:     &Result{10},
		objectType: reflect.Int8,
		isErr:      false,
	},
	TestData{
		result:     &Result{10},
		objectType: reflect.Int16,
		isErr:      false,
	},
	TestData{
		result:     &Result{10},
		objectType: reflect.Int32,
		isErr:      false,
	},
	TestData{
		result:     &Result{10},
		objectType: reflect.Int64,
		isErr:      false,
	},
	TestData{
		result:     &Result{10},
		objectType: reflect.Int,
		isErr:      false,
	},
	TestData{
		result:     &Result{1.0},
		objectType: reflect.Float32,
		isErr:      false,
	},
	TestData{
		result:     &Result{1.0},
		objectType: reflect.Float64,
		isErr:      false,
	},
	TestData{
		result:     &Result{1 + 0i},
		objectType: reflect.Complex64,
		isErr:      false,
	},
	TestData{
		result:     &Result{1 + 0i},
		objectType: reflect.Complex128,
		isErr:      false,
	},
	TestData{
		result:     &Result{"hlts2"},
		objectType: reflect.String,
		isErr:      false,
	},
}

func TestResult(t *testing.T) {

	/*
		var err error

		for _, test := range resultTests {
			err = test.CheckResultConvert()

			if !test.CheckError(err) {
				fmt.Println(test.result)
				fmt.Println(err)
				t.Error("")
			}
		}
	*/
}
