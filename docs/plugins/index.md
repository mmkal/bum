# Plugins

When Chaos Bum runs inside of Netflix, it integrates with a number of
proprietary systems and contains some Netflix-specific business logic. For example:

* Terminations are logged with an internal event tracking system 
* Metrics are logged to an internal metrics system.
* Credentials are decrypted using an internal secrets system.
* Dynamic configuration properties are retrieved from an internal
  configuration system.
* Some custom rules that prevent certain termination combinations from
  occurring.

In order to support  release Chaos Bum as open source, these proprietary
integrations are implemented as *plugins* that aren't released. Chaos Bum
ships with no-op implementations of these plugins.


## Building Chaos Bum with custom plugins

As an example, let's say you wished to implement a custom
[constrainer](Constrainer) for your organization.

This doc assumes that you will put the code in
`$GOPATH/example.com/chaosbum`. You should substitute "example.com" with
something relevant to your organization.

### 1. Grab the open source Chaos Bum source

If you haven't done this already, ensure the open source code is on your local
machine. You can use `go get` for this:

    go get github.com/netflix/chaosbum/cmd/chaosbum

### 2. Create a file with the custom constrainer implementation.

File: `$GOPATH/src/example.com/chaosbum/constrainer.go`

See the [Constrainer](Constrainer) page for an example implementation.


### 3. Create the file that loads the plugins

File:  `$GOPATH/src/example.com/chasbum/cmd/chaosbum/main.go`

It looks like this:

```go
package main


import (
    "github.com/Netflix/chaosbum/command"

    _ "example.com/chaosbum/constrainer"
)

func main() {
    command.Execute()
}
```

### 4. Build the custom Chaos Bum binary

```
go build example.com/chaosbum/cmd/chaosbum
```
