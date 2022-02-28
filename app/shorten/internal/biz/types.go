package biz

// ShortenURL 短链
type ShortenURL struct {
	ID    int64
	URL   string
	Host  string
	URI   string
	Query string
	Code  string
}

type CreateIn struct {
	URL string
}

type CreateOut struct {
	ID   int64
	URL  string
	Code string
}

type GetIn struct {
	ID   int64
	Code string
}

type GetOut struct {
	ID   int64
	URL  string
	Code string
}

type DeleteIn struct {
	ID   int64
	Code string
}
