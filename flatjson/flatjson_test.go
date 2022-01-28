package flatjson

import (
	"bytes"
	"github.com/nsf/jsondiff"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		inputJson      string
		want           string
		wantErrContain string
	}{
		{
			name: "Valid JSON with number, boolean, and string, depth 2 must be parsed successfully",
			inputJson: `
				{
					"a": 1,
					"b": true,
					"c": {
						"d": 3,
						"e": "test"
					}
				}`,
			want: `
				{
					"a": 1,
					"b": true,
					"c.d": 3,
					"c.e": "test"
				}`,
		},
		{
			name: "Valid JSON with depth 4 must be parsed successfully",
			inputJson: `
				{
					"aa": {
						"bb": {
							"cc": {
								"dd": 3.1415
							}
						}
					}
				}`,
			want: `
				{
					"aa.bb.cc.dd": 3.1415
				}`,
		},
		{
			name: "Invalid input JSON (extra comma) must throw error",
			inputJson: `
				{
					"aa": false,
				}`,
			wantErrContain: "invalid character",
		},
		{
			name: "Input JSON containing unsupported array must throw error",
			inputJson: `
				{
					"aa": [1, 2, 3]
				}`,
			wantErrContain: "unexpected type []interface",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			gotErr := FlatJSON(&Params{
				Reader: strings.NewReader(test.inputJson),
				Writer: &buf,
			})
			if (test.wantErrContain != "" && gotErr == nil) ||
				(test.wantErrContain == "" && gotErr != nil) ||
				(gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErrContain)) {
				t.Fatalf("want FlatJSON return error %q, got %q", test.wantErrContain, gotErr.Error())
			}
			if gotErr == nil {
				got := buf.String()
				diffOpts := jsondiff.DefaultConsoleOptions()
				if res, diff := jsondiff.Compare([]byte(test.want), []byte(got), &diffOpts); res != jsondiff.FullMatch {
					t.Errorf("FlatJSON(%q) returned diff\n%v", test.inputJson, diff)
				}
			}
		})
	}
}
