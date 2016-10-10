package mysql

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Context wraps a database connection to make certain queries easier.
type Context struct {
	db *sql.DB
}

func GetContext(username, password, url, database string) (*Context, error) {
	if url != "" {
		url = fmt.Sprintf("tcp(%s)", url)
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", username, password, url, database))
	if err != nil {
		return nil, err
	}
	return &Context{
		db,
	}, nil
}

func (c *Context) GetPrimaryKeyColumns(tableName string) (map[string]bool, error) {
	rows, err := c.db.Query("describe " + tableName)
	if err != nil {
		return nil, err
	}

	columns := make(map[string]bool, 0)
	for rows.Next() {
		var Field sql.NullString
		var Key sql.NullString
		var Ignore sql.NullString
		err = rows.Scan(&Field, &Ignore, &Ignore, &Key, &Ignore, &Ignore)
		if err != nil {
			panic(err)
		}
		if Key.Valid && Key.String == "PRI" && Field.Valid {
			columns[Field.String] = true
		}
	}

	return columns, rows.Err()
}

func (c *Context) MapTableTypes(tableName string) (map[string]string, error) {
	rows, err := c.db.Query("describe " + tableName)
	if err != nil {
		return nil, err
	}

	types := make(map[string]string, 0)
	for rows.Next() {
		var Field sql.NullString
		var Type sql.NullString
		var Ignore sql.NullString
		err = rows.Scan(&Field, &Type, &Ignore, &Ignore, &Ignore, &Ignore)
		if err != nil {
			panic(err)
		}
		if Type.Valid && Field.Valid {
			types[Field.String] = getGoTypeFromSqlType(Type.String)
		}
	}

	return types, rows.Err()
}

func getGoTypeFromSqlType(sqlType string) string {
	matchesStringType := regexp.MustCompile(`(varchar\(\d+\)|text)`)
	matchesIntType := regexp.MustCompile(`(int|tinyint|bigint)\(\d+\)`)
	sqlType = strings.ToLower(sqlType)

	switch {
	case matchesStringType.MatchString(sqlType):
		return "string"
	case matchesIntType.MatchString(sqlType):
		return "int"
	}

	return ""
}
