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
				DefaultPageSize: 10,
				MaxPageSize:     100,
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
		pager  *Pager
		in     *QueryArgs
		expect *Result
	}{
		"nil query args": {
			pager: &Pager{
				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: nil,
			expect: &Result{
				Cursor:              "",
				Limit:               DefaultConfig.DefaultPageSize,
				IsForwardPagination: true,
			},
		},
		"empty args": {
			pager: &Pager{
				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: &QueryArgs{},
			expect: &Result{
				Cursor:              "",
				Limit:               DefaultConfig.DefaultPageSize,
				IsForwardPagination: true,
			},
		},
		"after and before are both defined": {
			pager: &Pager{
				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: &QueryArgs{
				After:  PtrFromStr("testCursor"),
				Before: PtrFromStr("testCursor"),
			},
			expect: &Result{
				Cursor:              "",
				Limit:               DefaultConfig.DefaultPageSize,
				IsForwardPagination: true,
			},
		},
		"after and first are both defined": {
			pager: &Pager{

				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: &QueryArgs{
				After: PtrFromStr("testCursor"),
				First: PtrFromUint(10),
			},
			expect: &Result{
				Cursor:              "testCursor",
				Limit:               10,
				IsForwardPagination: true,
			},
		},
		"after defined and first greater than max page size": {
			pager: &Pager{
				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: &QueryArgs{
				After: PtrFromStr("testCursor"),
				First: PtrFromUint(125),
			},
			expect: &Result{
				Cursor:              "testCursor",
				Limit:               DefaultConfig.MaxPageSize,
				IsForwardPagination: true,
			},
		},
		"after is defined without first": {
			pager: &Pager{
				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: &QueryArgs{
				After: PtrFromStr("testCursor"),
			},
			expect: &Result{
				Cursor:              "testCursor",
				Limit:               DefaultConfig.DefaultPageSize,
				IsForwardPagination: true,
			},
		},
		"before and last are both defined": {
			pager: &Pager{
				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: &QueryArgs{
				Before: PtrFromStr("testCursor"),
				Last:   PtrFromUint(30),
			},
			expect: &Result{
				Cursor:              "testCursor",
				Limit:               30,
				IsForwardPagination: false,
			},
		},
		"before defined and last greater than max page size": {
			pager: &Pager{
				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: &QueryArgs{
				Before: PtrFromStr("testCursor"),
				Last:   PtrFromUint(125),
			},
			expect: &Result{
				Cursor:              "testCursor",
				Limit:               DefaultConfig.MaxPageSize,
				IsForwardPagination: false,
			},
		},
		"both is defined without last": {
			pager: &Pager{
				DefaultPageSize: DefaultConfig.DefaultPageSize,
				MaxPageSize:     DefaultConfig.MaxPageSize,
			},
			in: &QueryArgs{
				Before: PtrFromStr("testCursor"),
			},
			expect: &Result{
				Cursor:              "testCursor",
				Limit:               DefaultConfig.DefaultPageSize,
				IsForwardPagination: false,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.pager.Process(tc.in)
			assert.Equal(t, got, tc.expect, "results should be equal")
		})
	}
}
