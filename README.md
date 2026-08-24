# Zep Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=https%3A%2F%2Fgithub.com%2Fgetzep%2Fzep-go)

The Zep Go library provides convenient access to the Zep APIs from Go.

## Table of Contents

- [Requirements](#requirements)
- [Initialize Client](#initialize-client)
- [Add Messages to Thread](#add-messages-to-thread)
- [Get User Context](#get-user-context)
- [Optionals](#optionals)
- [Request Options](#request-options)
- [Automatic Retries](#automatic-retries)
- [Reference](#reference)
- [Usage](#usage)
- [Environments](#environments)
- [Pagination](#pagination)
- [Errors](#errors)
- [Advanced](#advanced)
  - [Response Headers](#response-headers)
  - [Retries](#retries)
  - [Timeouts](#timeouts)
  - [Explicit Null](#explicit-null)
- [Contributing](#contributing)

## Requirements

This module requires Go version >= 1.13.

# Installation

Run the following command to use the Zep Go library in your module:

```sh
go get github.com/getzep/zep-go/v4
```

## Initialize Client

```go
import (
  "github.com/getzep/zep-go/v4"
  zepclient "github.com/getzep/zep-go/v4/client"
  "github.com/getzep/zep-go/v4/option"
)

client := zepclient.NewClient(
  // this api key is `api_secret` line from zep.yaml of your local server or your Zep cloud api-key
  option.WithAPIKey("<YOUR_API_KEY>"),
)
```

## Add Messages to thread

```go
_, err = client.Thread.AddMessages(ctx, threadID, &zep.AddThreadMessagesRequest{
    Messages: []*zep.Message{
        {
            Name:     zep.String("customer"),
            Content:  "Hello, can I buy some shoes?",
            Role:     "user",
        },
    },
})
```

## Get User context

```go
threadUserContext, err := client.Thread.GetUserContext(
    ctx,
    threadID,
    nil,
)
```

## Optionals

This library models optional primitives and enum types as pointers. This is primarily meant to distinguish
default zero values from explicit values (e.g. `false` for `bool` and `""` for `string`). A collection of
helper functions are provided to easily map a primitive or enum to its pointer-equivalent (e.g. `zep.Int`).

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes configuring
authorization tokens, or providing your own instrumented `*http.Client`.

These request options can either be
specified on the client so that they're applied on every request, or for an individual request, like so:

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

```go
// Specify default options applied on every request.
client := client.NewClient(
    option.WithAPIKey("<YOUR_API_KEY>"),
    option.WithHTTPClient(
        &http.Client{
            Timeout: 5 * time.Second,
        },
    ),
)

// Specify options for an individual request.
response, err := client.Batch.Create(
    ...,
    option.WithAPIKey("<YOUR_API_KEY>"),
)
```

When credentials are not explicitly provided, the client reads them from the
following environment variables:

- `ZEP_API_KEY`

## Automatic Retries

The Zep Go client is instrumented with automatic retries with exponential backoff. A request will be
retried as long as the request is deemed retriable and the number of retry attempts has not grown larger
than the configured retry limit (default: 2).

A request is deemed retriable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [409](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409) (Conflict)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

You can use the `option.WithMaxAttempts` option to configure the maximum retry limit to
your liking. For example, if you want to disable retries for the client entirely, you can
set this value to 1 like so:

```go
client := zepclient.NewClient(
  option.WithMaxAttempts(1),
)
```

This can be done for an individual request, too:

```go
_, _ = client.Thread.GetUserContext(ctx, "thread_id", nil, option.WithMaxAttempts(1))
```

## Reference

A full reference for this library is available [here](https://github.com/getzep/zep-go/blob/HEAD/./reference.md).

## Usage

Instantiate and use the client with the following:

```go
package example

import (
    context "context"

    zep "github.com/getzep/zep-go/v4"
    client "github.com/getzep/zep-go/v4/client"
    option "github.com/getzep/zep-go/v4/option"
)

func do() {
    client := client.NewClient(
        option.WithAPIKey(
            "<value>",
        ),
    )
    request := &zep.CreateBatchRequest{}
    client.Batch.Create(
        context.TODO(),
        request,
    )
}
```

## Environments

You can choose between different environments by using the `option.WithBaseURL` option. You can configure any arbitrary base
URL, which is particularly useful in test environments.

```go
client := client.NewClient(
    option.WithBaseURL(zep.Environments.Default),
)
```

## Pagination

List endpoints are paginated. The SDK provides an iterator so that you can simply loop over the items.
You can also iterate page-by-page using the `GetNextPage` helper method.

The `Page.Results` attribute, which contains the relevant list of items returned by the call to the server,
is the only attribute you will need for most use cases. But if need be, several other attributes are available:

- `Page.Response` contains the full spec-defined response as returned by the server.
- `Page.StatusCode` and `Page.Header` returns HTTP metadata associated with the call to the server.
- `Page.RawResponse` returns the pagination object if you need to access its fields (like `Next`).

```go
// Loop over the items using the provided iterator.
ctx := context.TODO()
page, err := client.Batch.List(
    ctx,
    ...
)
if err != nil {
    return err
}
iter := page.Iterator()
for iter.Next(ctx) {
    item := iter.Current()
    fmt.Printf("Got item: %v", *item)
}
if err := iter.Err(); err != nil {
    return err
}

// Alternatively, iterate page-by-page.
for page != nil {
    for _, item := range page.Results {
        fmt.Printf("Got item: %v", *item)
    }
    page, err = page.GetNextPage(ctx)
    if errors.Is(err, core.ErrNoPages) {
        break
    }
    if err != nil {
        return err
    }
}

// Paginated endpoints return a Page with directly accessible headers, status code, and full response
page, err = client.Batch.List(
    ctx,
    ...
)
if err != nil {
    return err
}

// Access response metadata directly from the page
fmt.Printf("Got headers: %v", page.Header)
fmt.Printf("Got status code: %d", page.StatusCode)

// Access the full spec-defined response object
fullResponse := page.Response

// Access individual fields from the pagination object
nextCursor := page.RawResponse.Next
```

## Errors

Structured error types are returned from API calls that return non-success status codes. These errors are compatible
with the `errors.Is` and `errors.As` APIs, so you can access the error like so:

```go
response, err := client.Batch.Create(...)
if err != nil {
    var apiError *core.APIError
    if errors.As(err, &apiError) {
        // Do something with the API error ...
    }
    return err
}
```

## Advanced

### Response Headers

You can access the raw HTTP response data by using the `WithRawResponse` field on the client. This is useful
when you need to examine the response headers received from the API call. (When the endpoint is paginated,
the raw HTTP response data will be included automatically in the Page response object.)

```go
response, err := client.Batch.WithRawResponse.Create(...)
if err != nil {
    return err
}
fmt.Printf("Got response headers: %v", response.Header)
fmt.Printf("Got status code: %d", response.StatusCode)
```

### Retries

The SDK is instrumented with automatic retries with exponential backoff. A request will be retried as long
as the request is deemed retryable and the number of retry attempts has not grown larger than the configured
retry limit (default: 2).

Which status codes are retried depends on the `retryStatusCodes` generator configuration:

**`legacy`** (current default): retries on
- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#server_error_responses) (All server errors, including 500)

**`recommended`**: retries on
- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [502](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502) (Bad Gateway)
- [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503) (Service Unavailable)
- [504](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504) (Gateway Timeout)

If the `Retry-After` header is present in the response, the SDK will prioritize respecting its value exactly
over the default exponential backoff.

Use the `option.WithMaxAttempts` option to configure this behavior for the entire client or an individual request:

```go
client := client.NewClient(
    option.WithMaxAttempts(1),
)

response, err := client.Batch.Create(
    ...,
    option.WithMaxAttempts(1),
)
```

### Timeouts

Setting a timeout for each individual request is as simple as using the standard context library. Setting a one second timeout for an individual API call looks like the following:

```go
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()

response, err := client.Batch.Create(ctx, ...)
```

### Explicit Null

If you want to send the explicit `null` JSON value through an optional parameter, you can use the setters\
that come with every object. Calling a setter method for a property will flip a bit in the `explicitFields`
bitfield for that setter's object; during serialization, any property with a flipped bit will have its
omittable status stripped, so zero or `nil` values will be sent explicitly rather than omitted altogether:

```go
type ExampleRequest struct {
    // An optional string parameter.
    Name *string `json:"name,omitempty" url:"-"`

    // Private bitmask of fields set to an explicit value and therefore not to be omitted
    explicitFields *big.Int `json:"-" url:"-"`
}

request := &ExampleRequest{}
request.SetName(nil)

response, err := client.Batch.Create(ctx, request, ...)
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
