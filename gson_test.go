package gson

import (
	"testing"
)

type NewTest struct {
	json  string
	isErr bool
}

type NewTests []NewTest

var newTests = NewTests{
	NewTest{
		json:  `1`,
		isErr: false,
	},
	NewTest{
		json:  `"2"`,
		isErr: false,
	},
	NewTest{
		json:  `{"picture": "http://hogehoge"}`,
		isErr: false,
	},
	NewTest{
		json:  `{afsf: adfaasf`,
		isErr: true,
	},
	NewTest{
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
		isErr: false,
	},
	NewTest{
		json:  `[{"name": "little"}, {"name": "tiny"}]`,
		isErr: false,
	},
	NewTest{
		json:  `[{"name": "litt]`,
		isErr: true,
	},
}

func CheckError(t *testing.T, isErr bool, err error) bool {
	if isErr && err == nil {
		return false
	}

	if !isErr && err != nil {
		return false
	}

	return true
}

func TestNewGosonFromString(t *testing.T) {
	for _, test := range newTests {
		_, err := NewGosonFromString(test.json)

		if CheckError(t, test.isErr, err) {

		}
	}
}
