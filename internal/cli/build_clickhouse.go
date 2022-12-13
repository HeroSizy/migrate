//go:build clickhouse
// +build clickhouse

package cli

import (
	_ "code.in.spdigital.sg/sp-digital/migrate/v4/database/clickhouse"
	_ "github.com/ClickHouse/clickhouse-go"
)
