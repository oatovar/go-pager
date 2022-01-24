// Package pager allows for easy cursor based pagination. The package is
// intended to assist in providing a cursor based approach similar to that
// of the GraphQL pagination specification.
package pager

import "fmt"

// Pager follows the GraphQL cursor connections specification.
// It captures the first/after and last/before arguments and
// determines if the service should paginate forwards or backwards.
type Pager struct {
	First           *uint   `json:"first,omitempty"`
	After           *string `json:"after,omitempty"`
	Last            *uint   `json:"last,omitempty"`
	Before          *string `json:"before,omitempty"`
	defaultPageSize uint    `json:"-"`
	maxPageSize     uint    `json:"-"`
}

var DefaultConfig = Config{
	DefaultPageSize: 10,
	MaxPageSize:     100,
}

type Config struct {
	// DefaultPageSize dictates the default limit when "first" or "last" is
	// not provided.
	DefaultPageSize uint
	// MaxPageSize dictates the max limit that can be supplied in a query.
	MaxPageSize uint
}

func New(configs ...Config) (*Pager, error) {
	if len(configs) == 0 {
		return &Pager{
			defaultPageSize: DefaultConfig.DefaultPageSize,
			maxPageSize:     DefaultConfig.MaxPageSize,
		}, nil
	}

	config := configs[0]

	if config.DefaultPageSize > config.MaxPageSize {
		return nil, fmt.Errorf("default page size greater than max page size")
	}
	return &Pager{
		defaultPageSize: config.DefaultPageSize,
		maxPageSize:     config.MaxPageSize,
	}, nil
}

// Process evaluates the arguments passed in and produces a cursor, limit and
// decision on whether the caller should forward paginate. The following conditions
// are taken into account in the listed order.
//
// If both an "after" and "before" cursor are supplied then an empty cursor is produced
// with the default page size as the limit and the caller should forward paginate.
//
// If an "after" cursor is defined then it's returned as the parsed cursor. The "first"
// argument is returned as the limit unless it exceeds the defined max page size, in which
// case the max page size will be used. If "first" is not provided then the default page
// size will be used as the limit. The caller should forward paginate.
//
// If a "before" cursor is defined then it will be returned as the parsed cursor. The
// "last" argument is returned as the limit unless it exceeds the defined max page size, in
// which case the max page size will be used. If "last" is not provided then the default page
// size will be used as the limit. The caller should not forward paginate.
//
// If no arguments are provided an empty cursor will be returned with the default page size
// and the caller should forward paginate.
func (p Pager) Process() (cursor string, limit uint, isForwardPagination bool) {
	if p.After != nil && p.Before != nil {
		return "", p.defaultPageSize, true
	}

	if p.After != nil && p.First != nil {
		return *p.After, min(p.maxPageSize, *p.First), true
	}

	if p.After != nil {
		return *p.After, p.defaultPageSize, true
	}

	if p.Before != nil && p.Last != nil {
		return *p.Before, min(p.maxPageSize, *p.Last), false
	}

	if p.Before != nil {
		return *p.Before, p.defaultPageSize, false
	}

	return "", p.defaultPageSize, true
}

func min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}
