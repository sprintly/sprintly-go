# Sprint.ly Go API

This is a Sprint.ly API client in Go. It aims to have full API
coverage, but will be implemented as need. Patches are welcome!

[![Build Status](https://drone.io/github.com/sprintly/UserVoice-Import/status.png)](https://drone.io/github.com/sprintly/UserVoice-Import/latest)
[![Coverage Status](https://coveralls.io/repos/sprintly/sprintly-go/badge.png?branch=master)](https://coveralls.io/r/sprintly/sprintly-go?branch=master)

## How to install

```bash
go get github.com/sprintly/sprintly-go/
```

## How to use it

```go
import (
       "github.com/sprintly/sprintly-go/sprintly"
)

s_client := sprintly.NewSprintlyClient("email@example.org", "sprintly_api_key", 123)
// use s_client
```

## Making Changes

To make changes, submit a pull request! Your commit should contain
tests in the style of those already in the repository. If you have any
questions, please contact us and let us know!

