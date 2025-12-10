# krtequals

A Go static analysis tool (linter) that checks `Equals()` method implementations for KRT-style semantic equality issues.

## Overview

krtequals is designed to catch common bugs in `Equals()` method implementations, particularly for types used with [KRT (Kubernetes Declarative Controller Runtime)](https://github.com/istio/istio/tree/master/pkg/kube/krt). KRT relies on correct `Equals()` implementations to determine when resources have changed and need reprocessing.

## Installation

```bash
go install github.com/kgateway-dev/krtequals@latest
```

## Usage with golangci-lint

krtequals is designed to be used as a [golangci-lint](https://golangci-lint.run/) custom plugin. Add it to your `.golangci.yaml`:

```yaml
linters-settings:
  custom:
    krtequals:
      type: "module"
      description: "Checks Equals() implementations for KRT-style semantic equality issues"
      settings:
        deepEqual: true  # Enable reflect.DeepEqual detection (optional)
```

## What It Checks

### Missing Field Comparisons

The analyzer ensures that all exported fields of a struct are used in its `Equals()` method. If a field is not compared, it could lead to incorrect equality results.

```go
type MyStruct struct {
    Name    string
    Value   int
    Missing int  // This field is not compared - analyzer will flag it
}

func (m MyStruct) Equals(other MyStruct) bool {
    return m.Name == other.Name && m.Value == other.Value
    // Error: field "Missing" in type "MyStruct" is not used in Equals
}
```

### reflect.DeepEqual Usage (Optional)

When enabled via the `deepEqual` config option, the analyzer flags usage of `reflect.DeepEqual` inside `Equals()` methods. `reflect.DeepEqual` is slow and ignores the `+noKrtEquals` markers.

```go
func (m MyStruct) Equals(other MyStruct) bool {
    return reflect.DeepEqual(m, other)  // Flagged when deepEqual: true
}
```

## Markers

### `+noKrtEquals`

Use this marker to exclude a field from the equality check. Add it as a comment on the field:

```go
type MyStruct struct {
    Name  string
    // +noKrtEquals reason: internal bookkeeping
    cache map[string]string
}
```

### `+krtEqualsTodo`

Use this marker to temporarily exclude a field while indicating it needs to be addressed:

```go
type MyStruct struct {
    Name  string
    // +krtEqualsTodo reconcile waypoint usage
    Pending int
}
```

## Configuration

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `deepEqual` | bool | `false` | Enable detection of `reflect.DeepEqual` usage in `Equals()` methods |

## Development

### Running Tests

```bash
go test -v ./...
```

### Building

```bash
go build ./...
```

## License

Apache License 2.0
