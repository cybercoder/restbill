package op

type QueryOperator string

const (
	Equal          QueryOperator = "="
	NotEqual       QueryOperator = "!="
	GreaterThan    QueryOperator = ">"
	GreaterOrEqual QueryOperator = ">="
	LessThan       QueryOperator = "<"
	LessOrEqual    QueryOperator = "<="
	Like           QueryOperator = "LIKE"
	In             QueryOperator = "IN"
	NotIn          QueryOperator = "NOT IN"
	IsNull         QueryOperator = "IS NULL"
	IsNotNull      QueryOperator = "IS NOT NULL"
)
