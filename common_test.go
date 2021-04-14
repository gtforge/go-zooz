package zooz_test

import (
	"encoding/json"
	"testing"

	"github.com/gtforge/go-zooz"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDecodeJSON__UnmarshalJSON(t *testing.T) {
	type testCase struct {
		name         string
		incomingData []byte
		expectedErr  error
		expectedRes  zooz.DecodedJSON
	}

	testCases := []testCase{
		{
			name: "positive",
			incomingData: []byte(`
				"{\"a\":\"{\\\"b\\\":1}\"}"
			`),
			expectedRes: zooz.DecodedJSON{
				"a": zooz.DecodedJSON{
					"b": json.RawMessage("1"),
				},
			},
		},
		{
			name:         "json is not correct",
			incomingData: []byte("asdasd"),
			expectedErr:  errors.New("invalid character 'a' looking for beginning of value"),
		},
		{
			name: "more complicated json",
			incomingData: []byte(`
			"{\"c\":\"{\\\"b\\\":\\\"{\\\\\\\"a\\\\\\\":1}\\\"}\"}"
		`),
			expectedRes: zooz.DecodedJSON{
				"c": zooz.DecodedJSON{
					"b": zooz.DecodedJSON{
						"a": json.RawMessage("1"),
					},
				},
			},
		},
		{
			name: "more more complicated json",
			incomingData: []byte(`
			"{\"c\":{\"b\":{\"c\":{\"e\":{\"a\":1}}}},\"d\":{\"a\":{\"1\":\"\\\"\\\\\\\"{\\\\\\\\\\\\\\\"b\\\\\\\\\\\\\\\":{\\\\\\\\\\\\\\\"c\\\\\\\\\\\\\\\":{\\\\\\\\\\\\\\\"e\\\\\\\\\\\\\\\":{\\\\\\\\\\\\\\\"a\\\\\\\\\\\\\\\":1}}}}\\\\\\\"\\\"\"}}}"
		`),
			expectedRes: zooz.DecodedJSON{
				"c": zooz.DecodedJSON{
					"b": zooz.DecodedJSON{
						"c": zooz.DecodedJSON{
							"e": zooz.DecodedJSON{
								"a": json.RawMessage("1"),
							},
						},
					},
				},
				"d": zooz.DecodedJSON{
					"a": zooz.DecodedJSON{
						"1": zooz.DecodedJSON{
							"b": zooz.DecodedJSON{
								"c": zooz.DecodedJSON{
									"e": zooz.DecodedJSON{
										"a": json.RawMessage("1"),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "do not need to unquote string",
			incomingData: []byte(`
			{"b":{"c":{"e":{"a":1}}}}
			`),
			expectedRes: zooz.DecodedJSON{
				"b": zooz.DecodedJSON{
					"c": zooz.DecodedJSON{
						"e": zooz.DecodedJSON{
							"a": json.RawMessage("1"),
						},
					},
				},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			res := zooz.DecodedJSON{}
			err := json.Unmarshal(tC.incomingData, &res)
			if tC.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tC.expectedErr.Error())
				return
			}
			assert.Equal(t, tC.expectedRes, res)
		})
	}
}
