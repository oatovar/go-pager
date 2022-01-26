# go-pager

![tests workflow](https://github.com/oatovar/go-pager/actions/workflows/tests.yml/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/e31fd7a44a74a6dbdac1/maintainability)](https://codeclimate.com/github/oatovar/go-pager/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e31fd7a44a74a6dbdac1/test_coverage)](https://codeclimate.com/github/oatovar/go-pager/test_coverage)

Cursor based pagination made easy.

## Installation

```
go get -u github.com/oatovar/go-pager
```

## Examples

You can use the provider binder to bind the values from
`url.URL` into a `QueryArgs` struct.

```Golang

func Handler(w http.ResponseWriter, r *http.Request) {
	var queryArgs pager.QueryArgs
	err := pager.BindRequest(r, &queryArgs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := pager.New()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
    result := p.Process(&queryArgs)
	// Use collected args further on...
}
```

### Gorilla Scheme Usage

You can utilize the `github.com/gorilla/schema` package to parse the
the query args from the URL query values.

```Golang

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var queryArgs pager.QueryArgs
	err := schema.NewDecoder().Decode(queryArgs, r.Form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := pager.New()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    result := p.Process(&queryArgs)
	// Use collected args further on...
}
```

### Echo and Fiber

Additionally, you can use this with any framework that utilizes
the `query` struct tags to bind query params to a struct. Notably,
this includes the [Fiber](https://github.com/gofiber/fiber) and
[Echo](https://github.com/labstack/echo) http frameworks. An example,
is provided below for each respectively.

```Golang
// Echo Usage Example
func Handler(c echo.Context) error {
	var queryArgs pager.QueryArgs
	err := (&echo.DefaultBinder{}).BindQueryParams(c, &queryArgs)
	if err != nil {
		return fmt.Errorf("oops we encountered an error processing your request")
	}

	return c.JSON(http.StatusOK, queryArgs)
}
```

```Golang
// Fiber Usage Example
func Handler(c *fiber.Ctx) error {
	var queryArgs pager.QueryArgs
	err := c.QueryParser(&queryArgs)
	if err != nil {
		return fmt.Errorf("oops we encountered an error processing your request")
	}

	return c.JSON(queryArgs)
}
```

## F.A.Q.

**Does this need to be used with GraphQL?** No, you can use it as it best fits
in your application.

**Do you need to use it in query params?** No, you can potentially utilize it to parse
the request body if you're building a GraphQL API.

**Why not use offset and limit?** As applications begin to grow, cursor based pagination
is a lot more efficient. If you use a relational database, the offset/limit approach will
force the server to parse a lot of records potentially. You can read more about how
[Slack evaluated the usage of cursor based pagination](https://slack.engineering/evolving-api-pagination-at-slack/)
to get the tradeoffs that should be considered.

## Disclaimer

This is a work in progress package. Once it has been finalized, it will
be released as `v1.0.0`. If you have any recommendations or would like
a specific feature, please create an issue with your use case and any
solution recommendations.
