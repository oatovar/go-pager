package pager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		in     *Config
		expect *Pager
	}{
		"default config": {
			in: &DefaultConfig,
			expect: &Pager{
				defaultPageSize: 10,
				maxPageSize:     100,
			},
		},
		"default page size > max page size": {
			in: &Config{
				DefaultPageSize: 10,
				MaxPageSize:     5,
			},
			expect: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := New(*tc.in)
			assert.Equal(t, tc.expect, got)
		})
	}
}

func TestProcess(t *testing.T) {
	tests := map[string]struct {
		in                        *Pager
		expectCursor              string
		expectLimit               uint
		expectIsForwardPagination bool
	}{
		"empty arguments": {
			in: &Pager{
				defaultPageSize: DefaultConfig.DefaultPageSize,
				maxPageSize:     DefaultConfig.MaxPageSize,
			},
			expectCursor:              "",
			expectLimit:               DefaultConfig.DefaultPageSize,
			expectIsForwardPagination: true,
		},
		"after and before are both defined": {
			in: &Pager{
				After:           ptrFromStr("testCursor"),
				Before:          ptrFromStr("testCursor"),
				defaultPageSize: DefaultConfig.DefaultPageSize,
				maxPageSize:     DefaultConfig.MaxPageSize,
			},
			expectCursor:              "",
			expectLimit:               DefaultConfig.DefaultPageSize,
			expectIsForwardPagination: true,
		},
		"after and first are both defined": {
			in: &Pager{
				After:           ptrFromStr("testCursor"),
				First:           ptrFromUint(10),
				defaultPageSize: DefaultConfig.DefaultPageSize,
				maxPageSize:     DefaultConfig.MaxPageSize,
			},
			expectCursor:              "testCursor",
			expectLimit:               10,
			expectIsForwardPagination: true,
		},
		"after defined and first greater than max page size": {
			in: &Pager{
				After:           ptrFromStr("testCursor"),
				First:           ptrFromUint(125),
				defaultPageSize: DefaultConfig.DefaultPageSize,
				maxPageSize:     DefaultConfig.MaxPageSize,
			},
			expectCursor:              "testCursor",
			expectLimit:               DefaultConfig.MaxPageSize,
			expectIsForwardPagination: true,
		},
		"after is defined without first": {
			in: &Pager{
				After:           ptrFromStr("testCursor"),
				defaultPageSize: DefaultConfig.DefaultPageSize,
				maxPageSize:     DefaultConfig.MaxPageSize,
			},
			expectCursor:              "testCursor",
			expectLimit:               DefaultConfig.DefaultPageSize,
			expectIsForwardPagination: true,
		},
		"before and last are both defined": {
			in: &Pager{
				Before:          ptrFromStr("testCursor"),
				Last:            ptrFromUint(30),
				defaultPageSize: DefaultConfig.DefaultPageSize,
				maxPageSize:     DefaultConfig.MaxPageSize,
			},
			expectCursor:              "testCursor",
			expectLimit:               30,
			expectIsForwardPagination: false,
		},
		"before defined and last greater than max page size": {
			in: &Pager{
				Before:          ptrFromStr("testCursor"),
				Last:            ptrFromUint(125),
				defaultPageSize: DefaultConfig.DefaultPageSize,
				maxPageSize:     DefaultConfig.MaxPageSize,
			},
			expectCursor:              "testCursor",
			expectLimit:               DefaultConfig.MaxPageSize,
			expectIsForwardPagination: false,
		},
		"both is defined without last": {
			in: &Pager{
				Before:          ptrFromStr("testCursor"),
				defaultPageSize: DefaultConfig.DefaultPageSize,
				maxPageSize:     DefaultConfig.MaxPageSize,
			},
			expectCursor:              "testCursor",
			expectLimit:               DefaultConfig.DefaultPageSize,
			expectIsForwardPagination: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotCursor, gotLimit, gotIsForwardPagination := tc.in.Process()
			assert.Equal(t, gotCursor, tc.expectCursor, "cursors should be equal")
			assert.Equal(t, gotLimit, tc.expectLimit, "limits should be equal")
			assert.Equal(t, gotIsForwardPagination, tc.expectIsForwardPagination, "cursors should be equal")
		})
	}
}

func ptrFromStr(in string) *string {
	return &in
}

func ptrFromUint(in uint) *uint {
	return &in
}
