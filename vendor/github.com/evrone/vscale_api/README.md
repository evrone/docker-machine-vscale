# Vscale API

[![Vexor status](https://ci.vexor.io/projects/4089aaeb-e6f6-4400-a8ec-0d00c6db8c9f/status.svg)](https://ci.vexor.io/ui/projects/4089aaeb-e6f6-4400-a8ec-0d00c6db8c9f/builds)

## Installation and documentation

To install `Vscale API`, simply run:

```
$ go get github.com/evrone/vscale_api
```

## Getting Started

``` go
package main
import(
  "fmt"
  vscale "github.com/evrone/vscale_api"
)

func GetAccountInfo() {
  client := vscale.New("API_SECRET_TOKEN")
  account, _, err := client.Account.Get()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Account info: %v", account)
}
```

## Contribution Guidelines

01. Fork
02. Change
03. PR
