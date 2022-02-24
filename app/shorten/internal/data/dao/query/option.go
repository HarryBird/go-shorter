package query

const (
	DefaultPagingLimit = 20
	MaxPagingLimit     = 100
)

type Option func(o *Condition)

type Condition struct {
	Where  Where
	Paging Paging
}

type Where struct {
	Id        int64
	Code      string
	IsDeleted bool
}

type Paging struct {
	Offset int
	Limit  int
}

func NewCondition() *Condition {
	return &Condition{
		Where: Where{
			IsDeleted: false,
		},
		Paging: Paging{
			Limit: DefaultPagingLimit,
		},
	}
}
