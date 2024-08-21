# GORM SingleStore Driver

**Attention**: The code in this repository is intended for experimental use only and is not fully tested, documented, or supported by SingleStore. Visit the [SingleStore Forums](https://www.singlestore.com/forum/) to ask questions about this repository.

## Installation

Make sure you already have [gorm](https://github.com/go-gorm/gorm) installed. Then add the SingleStore driver as a dependency to your Go project.

```
go get github.com/singlestore-labs/gorm-singlestore
```

This uses [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) under the hood.

## Quick Start

```go
import (
  "github.com/singlestore-labs/gorm-singlestore"
  "gorm.io/gorm"
)

dsn := "root:password@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
db, err := gorm.Open(singlestore.Open(dsn), &gorm.Config{})
```

Checkout [https://gorm.io](https://gorm.io) for more details.

## Resources

* [Documentation](https://docs.singlestore.com)
* [Twitter](https://twitter.com/SingleStoreDevs)
* [SingleStore Forums](https://www.singlestore.com/forum)
