package biz

// ShortenIn Shorten 入参
type ShortenIn struct {
	URL string
}

// ShortenIn Shorten 出参
type ShortenOut struct {
	Code string
	URL  string
}

type DecodeIn struct {
	Code string
}

type DecodeOut struct {
	URL string
}
