# Modern Treasury Go API Library

<a href="https://pkg.go.dev/github.com/Modern-Treasury/modern-treasury-go"><img src="https://pkg.go.dev/badge/github.com/Modern-Treasury/modern-treasury-go.svg" alt="Go Reference"></a>

The Modern Treasury Go library provides convenient access to the Modern Treasury REST
API from applications written in Go.

## Installation

Within a Go module, you can just import this package and let the Go compiler
resolve this module.

```go
import (
	"github.com/Modern-Treasury/modern-treasury-go" // imported as moderntreasury
)
```

Or, explicitly import this package with

```
go get -u 'github.com/Modern-Treasury/modern-treasury-go'
```

## Documentation

The API documentation can be found [here](https://docs.moderntreasury.com).

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
)

func main() {
	modernTreasuryClient := moderntreasury.NewModernTreasury(
		option.WithOrganizationID("my-organization-ID"), // defaults to os.LookupEnv("MODERN_TREASURY_ORGANIZATION_ID")
		option.WithAPIKey("my api key"),                 // defaults to os.LookupEnv("MODERN_TREASURY_API_KEY")
	)
	externalAccount, err := client.ExternalAccounts.New(context.TODO(), &requests.ExternalAccountNewParams{
		CounterpartyID: moderntreasury.F("9eba513a-53fd-4d6d-ad52-ccce122ab92a"),
		Name:           moderntreasury.F("my bank"),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", externalAccount)
}

```

### Request Fields

Types for requests look like the following:

```go
type FooParams struct {
	ID     fields.Field[string] `json:"id,required"`
	Number fields.Field[int64]  `json:"number,required"`
	Name   fields.Field[string] `json:"name"`
	Other  fields.Field[Bar]    `json:"other"`
}

type Bar struct {
	Number fields.Field[int64]  `json:"number"`
	Name   fields.Field[string] `json:"name"`
}
```

For each field, you can either supply a value field with
`moderntreasury.F(...)`, a `null` value with `moderntreasury.NullField()`, or
some raw JSON value with `moderntreasury.RawField(...)` that you specify as a
byte slice. We also provide convenient helpers `moderntreasury.Int(...)` and
`moderntreasury.Str(...)`. If you do not supply a value, then we do not
populate the field. An example request may look like

```go
params := &FooParams{
	// Normally populates this field as `"id": "food_id"`
	ID: moderntreasury.F("foo_id"),

	// Integer helper casts integer values and literals to fields.Field[int64]
	Number: moderntreasury.Int(12),

	// Explicitly sends this field as null, e.g., `"name": null`
	Name: moderntreasury.NullField[string](),

	// Overrides this field as `"other": "ovveride_this_field"`
	Other: moderntreasury.RawField[Bar]("override_this_field")
}
```

If you want to add or override a field in the JSON body, then you can use the
`option.WithJSONSet(key string, value interface{})` RequestOption, which you
can read more about [here](#requestoptions). Internally, this uses
'github.com/tidwall/sjson' library, so you can compose complex access as seen
[here](https://github.com/tidwall/sjson).

### Response Objects

Response objects in this SDK have value type members. Accessing properties on
response objects is as simple as:

```go
res, err := client.Service.Foo(context.TODO())
res.Name // is just some string value
```

If null, not present, or invalid, all fields will simply be their empty values.

If you want to be able to tell that the value is either `null`, not present, or
invalid, we provide metadata in the `JSON` property of every response object.

```go
// This is true if `name` is _either_ not present or explicitly null
res.JSON.Name.IsNull()

// This is true if `name` is not present
res.JSON.Name.IsMissing()

// This is true if `name` is present, but not coercable
res.JSON.Name.IsMissing()

// If needed, you can access the Raw JSON value of the field by accessing
res.JSON.Name.Raw()
```

You can also access the JSON value of the entire object with `res.JSON.Raw`.

There may be instances where we provide experimental or private API features
for some customers. In those cases, the related features will not be exposed to
the SDK as typed fields, and are instead deserialized to an internal map. We
provide methods to get and set these json fields in API objects.

```go
// Access the JSON value as
body := res.JSON.Extras["extra_field"].Raw()

// You can `Unmarshal` the JSON into a struct as needed
custom := struct{A string, B int64}{}
json.Unmarshal(body, &custom)
```

### RequestOptions

This library uses the functional options pattern. `RequestOptions` are closures
that mutate a `RequestConfig`. These options can be supplied to the client or
at individual requests, and they will cascade appropriately.

At each request, these closures will be run in the order that they are
supplied, after the defaults for that particular request.

For example:

```go
client := moderntreasury.NewModernTreasury(
	// Adds header to every request made by client
	option.WithHeader("X-Some-Header", "custom_header_info"),
	// Adds query to every request made by client
	option.WithQuery("test_token", "my_test_token"),
)

client.ExternalAccounts.New(
	context.TODO(),
	...,
	// These options override the client options
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	option.WithQuery("test_token", "some_other_test_token"),
)

client.ExternalAccounts.New(
	context.TODO(),
	...,
	// WithHeaderDel removes the header set in the client
	// from this request
	option.WithHeaderDel("X-Some-Header"),
	// WithQueryDel removes the query set in the client
	// from this request
	option.WithQueryDel("test_token"),
)
```

### Pagination

This library provides some conveniences for working with paginated list endpoints.

You can use `.ListAutoPager()` methods to iterate through items across all pages:

```go
iter := client.ExternalAccounts.ListAutoPager(context.TODO(), &requests.ExternalAccountListParams{})
// Automatically fetches more pages as needed.
for iter.Next() {
	externalAccount := iter.Current()
	fmt.Printf("%+v\n", externalAccount)
}
if err := iter.Err(); err != nil {
	panic(err.Error())
}
```

Or you can use simple `.List()` methods to fetch a single page and receive a standard response object
with additional helper methods like `.GetNextPage()`, e.g.:

```go
page, err := client.ExternalAccounts.List(context.TODO(), &requests.ExternalAccountListParams{})
for page != nil {
	for _, externalAccount := range page.Items {
		fmt.Printf("%+v\n", externalAccount)
	}
	page, err = page.GetNextPage()
}
if err != nil {
	panic(err.Error())
}
```

### Errors

Errors are still WIP.

## Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
Connection errors (for example, due to a network connectivity problem), 409 Conflict, 429 Rate Limit,
and >=500 Internal errors will all be retried by default.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
modernTreasuryClient := moderntreasury.NewModernTreasury(
	option.WithOrganizationID("my-organization-ID"), // defaults to os.LookupEnv("MODERN_TREASURY_ORGANIZATION_ID")
	option.WithMaxRetries(0),                        // default is 2
)

// Override per-request:
client.ExternalAccounts.List(
	context.TODO(),
	&requests.ExternalAccountListParams{},
	option.WithMaxRetries(5),
)
```

### Middleware

You may apply any middleware you wish by overriding the `http.Client` with
`option.WithClient(client)`. An example of a basic logging middleware is given
below:

```go
TODO
```

## Status

This package is in beta. Its internals and interfaces are not stable and
subject to change without a major semver bump; please reach out if you rely on
any undocumented behavior.

We are keen for your feedback; please email us at
[sdk-feedback@moderntreasury.com](mailto:sdk-feedback@moderntreasury.com) or open an issue with questions, bugs, or
suggestions.

## Requirements

This library requires Go 1.18+.