# GORM SingleStore Driver

## Installation

Make sure you already have [gorm](https://github.com/go-gorm/gorm) installed. Then add the SingleStore driver as a dependency to your Go project.

```
go get github.com/singlestore-labs/gorm-singlestore
```

This uses [go-sql-driver/singlestore](https://github.com/go-sql-driver/singlestore) under the hood.

## Quick Start

```go
import (
  "github.com/singlestore-labs/gorm-singlestore"
  "gorm.io/gorm"
)

dsn := "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
db, err := gorm.Open(singlestore.Open(dsn), &gorm.Config{})
```

Checkout [https://gorm.io](https://gorm.io) for more details.