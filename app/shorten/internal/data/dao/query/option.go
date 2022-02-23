package query

type Option func(o *Condition)

type Condition struct {
	Where  Where
	Paging Paging
}

type Where struct {
	Id   int64
	Code string
}

type Paging struct {
	Offset int
	Limit  int
}

func NewCondition() *Condition {
	return &Condition{
		Paging: Paging{
			Limit: 50,
		},
	}
}
