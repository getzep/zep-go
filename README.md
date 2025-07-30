# Zep Go Library

<p align="center">
  <a href="https://www.getzep.com/">
    <img src="https://raw.githubusercontent.com/getzep/zep/main/assets/zep-logo-icon-gradient-rgb.svg" width="150" alt="Zep Logo">
  </a>
</p>

<h1 align="center">
Zep: Long-Term Memory for ‍AI Assistants.
</h1>
<h2 align="center">Recall, understand, and extract data from chat histories. Power personalized AI experiences.</h2>
<br />
<p align="center">
  <a href="https://discord.gg/W8Kw6bsgXQ"><img
    src="https://img.shields.io/badge/Discord-%235865F2.svg?&logo=discord&logoColor=white"
    alt="Chat on Discord"
  /></a>
  <a href="https://twitter.com/intent/follow?screen_name=zep_ai" target="_new"><img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/zep_ai"></a>
 <a href="https://pkg.go.dev/github.com/getzep/zep-go/v2"><img src="https://pkg.go.dev/badge/github.com/getzep/zep-go/v2.svg" alt="Go Reference"></a> 
  <a href="https://goreportcard.com/report/github.com/getzep/zep-go/v2"><img src="https://goreportcard.com/badge/github.com/getzep/zep-go/v2" alt="Go Report Card"></a> 
  <img
  src="https://github.com/getzep/zep-go/actions/workflows/ci.yml/badge.svg"
  alt="CI"
  />
<a href="https://github.com/fern-api/fern">
    <img
      src="https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen"
      alt="CI"
      />
</a>
</p>
<p align="center">
<a href="https://help.getzep.com/">Documentation</a> | 
<a href="https://help.getzep.com/langchain/">LangChain</a> | 
<a href="https://discord.gg/W8Kw6bsgXQ">Discord</a><br />
<a href="https://www.getzep.com">www.getzep.com</a>
</p>



The Zep Go library provides convenient access to the Zep Cloud API from Go.

## Requirements

This module requires Go version >= 1.13.

# Installation

Run the following command to use the Zep Go library in your module:

```sh
go get github.com/getzep/zep-go/v3
```

## Initialize Client

```go
import (
  "github.com/getzep/zep-go/v3"
  zepclient "github.com/getzep/zep-go/v3/client"
  "github.com/getzep/zep-go/v3/option"
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

A variety of request options are included to adapt the behavior of the library, which includes
configuring authorization tokens, or providing your own instrumented `*http.Client`. Both of
these options are shown below:

```go
client := zepclient.NewClient(
  option.WithAPIKey("<YOUR_API_KEY>"),
  option.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

These request options can either be specified on the client so that they're applied on _every_
request (shown above), or for an individual request like so:

```go
_, _ = client.Thread.GetUserContext(ctx, "thread_id", nil, option.WithAPIKey("<YOUR_API_KEY>"))
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

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

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
_, err := client.Thread.GetUserContext(ctx, "thread_id", nil)
if err != nil {
  if badRequestErr, ok := err.(*zep.BadRequestError);
    // Do something with the bad request ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
_, err := client.Thread.GetUserContext(ctx, "thread_id", nil)
if err != nil {
  var badRequestErr *zep.BadRequestError
  if errors.As(err, badRequestErr) {
    // Do something with the bad request ...
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability
to access the type with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
_, err := client.Thread.GetUserContext(ctx, "thread_id", nil)
if err != nil {
  return fmt.Errorf("failed to get memory: %w", err)
}
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the `README.md` are always very welcome!
