package queries

import (
	"fmt"
	"github.com/RandySteven/Library-GO/enums"
)

type (
	MigrationQuery string
	DropQuery      string
	GoQuery        string

	QueryParam struct {
		Key      string
		Value    interface{}
		Operator enums.QueryOperator
	}

	QueryCondition struct {
		Operator enums.QueryOperator
		Param    *QueryParam
	}
)

func (p *QueryParam) ToString() string {
	return fmt.Sprintf("%s %s %s", p.Key, p.Operator, p.Value)
}

func ToString(conditions []*QueryCondition) string {
	return ""
}

func (q MigrationQuery) ToString() string {
	return string(q)
}

func (q DropQuery) ToString() string {
	return string(q)
}

func (q GoQuery) ToString() string {
	return string(q)
}
