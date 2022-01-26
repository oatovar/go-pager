// Package pager allows for easy cursor based pagination. The package is
// intended to assist in providing a cursor based approach similar to that
// of the GraphQL pagination specification.
package pager

import "fmt"

// Pager follows the GraphQL cursor based pagination specification.
// It processes QueryArgs and outputs a Result that contains the cursor,
// limit, and if the service should paginate forwards or backwards.
type Pager struct {
	DefaultPageSize uint `json:"-"`
	MaxPageSize     uint `json:"-"`
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
			DefaultPageSize: DefaultConfig.DefaultPageSize,
			MaxPageSize:     DefaultConfig.MaxPageSize,
		}, nil
	}

	config := configs[0]

	if config.DefaultPageSize > config.MaxPageSize {
		return nil, fmt.Errorf("default page size greater than max page size")
	}
	return &Pager{
		DefaultPageSize: config.DefaultPageSize,
		MaxPageSize:     config.MaxPageSize,
	}, nil
}

// QueryArgs captures all arguments required for GraphQL cursor based
// pagination.
type QueryArgs struct {
	First  *uint   `json:"first" schema:"first" query:"first"`
	After  *string `json:"after" schema:"after" query:"after"`
	Last   *uint   `json:"last" schema:"last" query:"last"`
	Before *string `json:"before" schema:"before" query:"before"`
}

// Result represents the outcome of processing the supplied QueryArgs.
type Result struct {
	// Cursor defines where to begin pagination.
	Cursor string
	// Limit defines how many records to return.
	Limit uint
	// IsForwardPagination defines if the query should look for records
	// that are less than or greater than the comparable cursor.
	IsForwardPagination bool
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
func (p Pager) Process(args *QueryArgs) *Result {
	if args == nil {
		return &Result{
			Cursor:              "",
			Limit:               p.DefaultPageSize,
			IsForwardPagination: true,
		}
	}

	if args.After != nil && args.Before != nil {
		return &Result{
			Cursor:              "",
			Limit:               p.DefaultPageSize,
			IsForwardPagination: true,
		}
	}

	if args.After != nil && args.First != nil {
		return &Result{
			Cursor:              *args.After,
			Limit:               min(p.MaxPageSize, *args.First),
			IsForwardPagination: true,
		}
	}

	if args.After != nil {
		return &Result{
			Cursor:              *args.After,
			Limit:               p.DefaultPageSize,
			IsForwardPagination: true,
		}
	}

	if args.Before != nil && args.Last != nil {
		return &Result{
			Cursor:              *args.Before,
			Limit:               min(p.MaxPageSize, *args.Last),
			IsForwardPagination: false,
		}
	}

	if args.Before != nil {
		return &Result{
			Cursor:              *args.Before,
			Limit:               p.DefaultPageSize,
			IsForwardPagination: false,
		}
	}

	return &Result{
		Cursor:              "",
		Limit:               p.DefaultPageSize,
		IsForwardPagination: true,
	}
}

func min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}
