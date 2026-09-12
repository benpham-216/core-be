# Contributing

Use a focused working branch and pull request. Describe lifecycle, policy, API, migration and event-contract changes explicitly. Event payload changes must be backward compatible or versioned. Do not mutate published event history.

Every pull request must include relevant unit tests and the output of `go test ./...` and `go vet ./...`.
