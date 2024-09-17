package enums

type QueryOperator string

const (
	And  QueryOperator = "AND"
	Or   QueryOperator = "OR"
	Like QueryOperator = "LIKE"
	In   QueryOperator = "IN"
	Not  QueryOperator = "NOT"
	Eq   QueryOperator = "="
	Ne   QueryOperator = "!="
	Gt   QueryOperator = ">"
	Lt   QueryOperator = "<"
	GtEq QueryOperator = ">="
	LtEq QueryOperator = "<="
	Is   QueryOperator = "IS"
)
