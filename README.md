# database-go

A simple key-value database for simple projects.

## Install

```sh
$ go get github.com/davipatricio/database-go
```

## Usage
```go
package main

import (
    "fmt"

    "github.com/davipatricio/database-go"
)

func main() {
    db := database.New("mydatabase.json")
    db.Load()

    db.Set("key", "value")
    db.Save()

    value := db.Get("key")
    fmt.Println(db.Get("key"))

    db.Delete("key")
    db.Save()
}
```

## License
MIT