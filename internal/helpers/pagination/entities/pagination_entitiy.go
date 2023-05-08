package entities

type Pagination struct {
	Page   int
	Order  string
	Cursor string
}

type PaginationInfo struct {
	NextCursor string
	PrevCursor string
}

type PaginationCalculate struct {
	ID        string
	CreatedAt int64
}

type Cursor map[string]interface{}
