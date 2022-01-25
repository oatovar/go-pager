# go-pager

![tests workflow](https://github.com/oatovar/go-pager/actions/workflows/tests.yml/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/e31fd7a44a74a6dbdac1/maintainability)](https://codeclimate.com/github/oatovar/go-pager/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e31fd7a44a74a6dbdac1/test_coverage)](https://codeclimate.com/github/oatovar/go-pager/test_coverage)

Cursor based pagination made easy.

## Installation

```
go get -u github.com/oatovar/go-pager
```

## Example

The following example uses `gorilla/schema` to parse the
the query args from the URL query values.

```Golang

// Initialize pager and reuse. Ideally this would be provided
// via dependency injection.
var pager = pager.New()

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// Handle errors gracefully
	}

	var queryArgs pager.QueryArgs
	if err := schema.NewDecoder().Decode(queryArgs, r.Form); err != nil {
		// Handle errors gracefully
	}

    result := pager.Process(&queryArgs)
	// Use collected args further on...
}


```

## Disclaimer

This is a work in progress package. Once it has been finalized, it will
be released as `v1.0.0`. If you have any recommendations or would like
a specific feature, please create an issue with your use case and any
solution recommendations.
