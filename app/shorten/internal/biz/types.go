package biz

// OriginURL .
type OriginURL struct {
	Url string
}

// ShortenURL 短链
type ShortenURL struct {
	ID       int64
	URLFull  string
	URLHost  string
	URLUri   string
	URLQuery string
	URLCode  string
}
