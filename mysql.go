package hood

import (
	"fmt"
	"time"
)

type Mysql struct {
	Base
}


func NewMysql() Dialect {
	d := &Mysql{}
	d.Base.Dialect = d
	return d
}

func (d *Mysql) NextMarker(pos *int) string {
	return "?"
}

func (d *Mysql) Quote(s string) string {
	return fmt.Sprintf("`%s`", s)
}

func (d *Mysql) SqlType(f interface{}, size int) string {
	switch f.(type) {
	case Id:
		return "bigint"
	case VarChar:
		if size < 1 {
			size = 255
		}
		return fmt.Sprintf("varchar(%d)", size)
	case time.Time, Created, Updated:
		return "timestamp"
	case bool:
		return "boolean"
	case int, int8, int16, int32, uint, uint8, uint16, uint32:
		return "int"
	case int64, uint64:
		return "bigint"
	case float32, float64:
		return "double"
	case []byte:
		return "longblob"
	case string:
		return "longtext"
	}
	panic("invalid sql type")
}

func (d *Mysql) KeywordAutoIncrement() string {
	return "AUTO_INCREMENT"
}
