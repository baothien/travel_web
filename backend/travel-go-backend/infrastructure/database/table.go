package pg

import "fmt"

const (
	DefaultSchema = "public"
)

func TableWithSchema(table string) string {
	return fmt.Sprintf("%s.%s", DefaultSchema, table)
}
