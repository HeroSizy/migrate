# pkger
```go
package main

import (
	"errors"
	"log"

	"code.in.spdigital.sg/sp-digital/migrate/v4"
	"github.com/markbates/pkger"

	_ "code.in.spdigital.sg/sp-digital/migrate/v4/database/postgres"
	_ "code.in.spdigital.sg/sp-digital/migrate/v4/source/pkger"
	_ "github.com/lib/pq"
)

func main() {
	pkger.Include("/module/path/to/migrations")
	m, err := migrate.New("pkger:///module/path/to/migrations", "postgres://postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	if err := m.Up(); errors.Is(err, migrate.ErrNoChange) {
		log.Println(err)
	} else if err != nil {
		log.Fatalln(err)
	}
}
```
