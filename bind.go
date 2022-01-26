package pager

import (
	"fmt"
	"net/http"
	"strconv"
)

func BindRequest(r *http.Request, out *QueryArgs) error {
	if r == nil {
		return fmt.Errorf("request must be non-nil reference")
	}

	if out == nil {
		return fmt.Errorf("out must be non-nil reference")
	}

	query := r.URL.Query()

	var (
		after  = query.Get("after")
		first  = query.Get("first")
		before = query.Get("before")
		last   = query.Get("last")
	)

	if len(after) > 0 {
		out.After = &after
	}

	if len(first) > 0 {
		val, err := strconv.Atoi(first)
		if err != nil {
			return err
		}

		firstUint := UintFromInt(val)
		out.First = &firstUint
	}

	if len(before) > 0 {
		out.Before = &before
	}

	if len(last) > 0 {
		val, err := strconv.Atoi(last)
		if err != nil {
			return err
		}

		lastUint := UintFromInt(val)
		out.Last = &lastUint
	}

	return nil
}
