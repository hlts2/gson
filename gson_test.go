package gson

import (
	"reflect"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		json  string
		want  *Gson
		iserr bool
	}{
		{
			json: `1`,
			want: &Gson{
				object: float64(1),
			},
			iserr: false,
		},
		{
			json: `"2"`,
			want: &Gson{
				object: "2",
			},
			iserr: false,
		},
		{
			json: `{"picture": "http://hogehoge"}`,
			want: &Gson{
				object: map[string]interface{}{
					"picture": "http://hogehoge",
				},
			},
			iserr: false,
		},
		{
			json:  `{afsf: adfaasf`,
			want:  nil,
			iserr: true,
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
			iserr: false,
		},
		{
			json:  `[{"name": "litt]`,
			want:  nil,
			iserr: true,
		},
	}

	for i, test := range tests {
		t.Run("CreateWithBytes", func(t *testing.T) {
			g, err := CreateWithBytes([]byte(test.json))

			iserr := !(err == nil)

			if test.iserr != iserr {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.iserr, iserr)
			}

			if !reflect.DeepEqual(g, test.want) {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.want, g)
			}
		})

		t.Run("CreateWithReader", func(t *testing.T) {
			g, err := CreateWithReader(strings.NewReader(test.json))

			if want, got := test.iserr, !(err == nil); want != got {
				t.Errorf("test[%d] - want: %v, but got: %v", i, want, got)
			}

			if !reflect.DeepEqual(g, test.want) {
				t.Errorf("tests[%d] - want: %v, but got: %v", i, test.want, g)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		json  string
		keys  []string
		want  *Result
		iserr bool
	}{
		{
			json:  `{"name": "hlts2"}`,
			keys:  []string{"name"},
			want:  &Result{"hlts2"},
			iserr: false,
		},
		{
			json:  `[{"name": "hlts2"}]`,
			keys:  []string{"0"},
			want:  &Result{map[string]interface{}{"name": "hlts2"}},
			iserr: false,
		},
		{
			json:  `[{"name": "hlts2"}]`,
			keys:  []string{"10"},
			want:  nil,
			iserr: true,
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
			iserr: false,
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
			keys:  []string{"friends", "100", "name"},
			want:  nil,
			iserr: true,
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

			if want, got := test.iserr, !(err == nil); want != got {
				t.Errorf("test[%d] - want: %v, but got: %v", i, want, got)
			}

			if !reflect.DeepEqual(test.want, r) {
				t.Errorf("tests[%d] - want: %v, got: %v", i, test.want, r)
			}
		})

		t.Run("GetByPath", func(t *testing.T) {
			r, err := g.GetByPath(strings.Join(test.keys, "."))

			if want, got := test.iserr, !(err == nil); want != got {
				t.Errorf("test[%d] - want: %v, but got: %v", i, want, got)
			}

			if !reflect.DeepEqual(test.want, r) {
				t.Errorf("tests[%d] - want: %v, got: %v", i, test.want, r)
			}
		})
	}
}

func TestSliceE(t *testing.T) {
	tests := []struct {
		json  string
		want  []*Result
		iserr bool
	}{
		{
			json: `{"Users": [{"ID": "1111"}, {"ID": "2222"}]}`,
			want: []*Result{
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
			iserr: false,
		},
	}

	for i, test := range tests {
		g, err := CreateWithBytes([]byte(test.json))
		if err != nil {
			t.Errorf("tests[%d] - CreateWithBytes returned error: %v", i, err)
		}

		result, err := g.GetByKeys("Users")
		if err != nil {
			t.Errorf("tests[%d] - GetByKeys returned error: %v", i, err)
		}

		s, err := result.SliceE()

		if want, got := test.iserr, !(err == nil); want != got {
			t.Errorf("tests[%d] - want: %v, got: %v", i, want, got)
		}

		if want, got := test.want, s; !reflect.DeepEqual(want, got) {
			t.Errorf("tests[%d] - want: %v, got: %v", i, want, got)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		json  string
		want  map[string]*Result
		iserr bool
	}{
		{
			json: `{"Accounts": [{"ID": "1111", "Name": "hlts2"}]}`,
			want: map[string]*Result{
				"ID": {
					object: "1111",
				},
				"Name": {
					object: "hlts2",
				},
			},
			iserr: false,
		},
	}

	for i, test := range tests {
		g, err := CreateWithBytes([]byte(test.json))
		if err != nil {
			t.Errorf("tests[%d] - CreateWithBytes returned error: %v", i, err)
		}

		if g == nil {
			t.Errorf("tests[%d] - CreateWithBytes returned nil", i)
		}

		result, err := g.GetByKeys("Accounts", "0")

		if err != nil {
			t.Errorf("tests[%d] - GetByKeys returned error: %v", i, err)
		}

		m, err := result.MapE()

		if want, got := test.iserr, !(err == nil); want != got {
			t.Errorf("tests[%d] - want: %v, got: %v", i, want, got)
		}

		if want, got := test.want, m; !reflect.DeepEqual(want, got) {
			t.Errorf("tests[%d] - want: %v, got: %v", i, want, got)
		}
	}
}
