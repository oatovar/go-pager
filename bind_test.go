package pager

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindRequest(t *testing.T) {
	tests := map[string]struct {
		in     *http.Request
		expect *QueryArgs
	}{
		"after and first defined": {
			in: &http.Request{
				URL: &url.URL{
					RawQuery: "after=abcdef&first=10&before=&last=",
				},
			},
			expect: &QueryArgs{
				After: PtrFromStr("abcdef"),
				First: PtrFromUint(10),
			},
		},
		"param defined multiple times": {
			in: &http.Request{
				URL: &url.URL{
					RawQuery: "after=abcdef&after=abcdefg&first=10&before=&last=",
				},
			},
			expect: &QueryArgs{
				After: PtrFromStr("abcdef"),
				First: PtrFromUint(10),
			},
		},
		"all empty": {
			in: &http.Request{
				URL: &url.URL{
					RawQuery: "after=&first=&before=&last=",
				},
			},
			expect: &QueryArgs{},
		},
		"before and last defined": {
			in: &http.Request{
				URL: &url.URL{
					RawQuery: "after=&first=&before=abcdef&last=10",
				},
			},
			expect: &QueryArgs{
				Before: PtrFromStr("abcdef"),
				Last:   PtrFromUint(10),
			},
		},
		"negative last value": {
			in: &http.Request{
				URL: &url.URL{
					RawQuery: "after=opaqueCursor&first=-10&before=&last",
				},
			},
			expect: &QueryArgs{
				After: PtrFromStr("opaqueCursor"),
				First: PtrFromUint(10),
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var got QueryArgs
			err := BindRequest(tc.in, &got)
			assert.Nil(t, err, "no error should be returned")
			assert.Equal(t, *tc.expect, got)
		})
	}
}
