package _type

type SongCreateRequest struct {
	Title      string `json:"title"`
	Overview   string `json:"overview"`
	AuthorAddr string `json:"author_addr"`
}
