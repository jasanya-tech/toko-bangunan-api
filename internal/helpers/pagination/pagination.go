package pagination

import (
	"encoding/base64"
	"encoding/json"

	paginateentity "toko-bangunan/internal/helpers/pagination/entities"
)

func GetPaginationOperator(pointNext bool, sortOrder string) (string, string) {
	if pointNext && sortOrder == "asc" {
		return ">", ""
	}
	if pointNext && sortOrder == "desc" {
		return "<", ""
	}
	if !pointNext && sortOrder == "asc" {
		return "<", "desc"
	}
	if !pointNext && sortOrder == "desc" {
		return ">", "asc"
	}
	return "", ""
}

func CalculatePagination(isFirstPage bool, hasPagination bool, limit int, data []paginateentity.PaginationCalculate, pointNext bool) paginateentity.PaginationInfo {
	pagination := paginateentity.PaginationInfo{}
	nextCursor := paginateentity.Cursor{}
	prevCursor := paginateentity.Cursor{}

	if isFirstPage {
		if hasPagination {
			nextCursor := createCursor(data[limit-1].ID, data[limit-1].CreatedAt, true)
			pagination = generatePage(nextCursor, nil)
		}
	} else {
		if pointNext {
			// if pointing next, it always has prev but it might not have next
			if hasPagination {
				nextCursor = createCursor(data[limit-1].ID, data[limit-1].CreatedAt, true)
			}

			prevCursor = createCursor(data[0].ID, data[0].CreatedAt, false)
			pagination = generatePage(nextCursor, prevCursor)

		} else {
			// this is case of prev, there will always be nest, but prev needs to be calculated
			nextCursor = createCursor(data[limit-1].ID, data[limit-1].CreatedAt, true)
			if hasPagination {
				prevCursor = createCursor(data[0].ID, data[0].CreatedAt, false)
			}
			pagination = generatePage(nextCursor, prevCursor)
		}
	}
	return pagination
}

func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func DecodedCursor(cursor string) (paginateentity.Cursor, error) {
	DecodedCursor, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, err
	}

	var cur paginateentity.Cursor
	if err := json.Unmarshal(DecodedCursor, &cur); err != nil {
		return nil, err
	}
	return cur, nil
}

func createCursor(id string, createdAt int64, pointNext bool) paginateentity.Cursor {
	return paginateentity.Cursor{
		"id":         id,
		"created_at": createdAt,
		"point_next": pointNext,
	}
}

func generatePage(next paginateentity.Cursor, prev paginateentity.Cursor) paginateentity.PaginationInfo {
	return paginateentity.PaginationInfo{
		NextCursor: encodeCursor(next),
		PrevCursor: encodeCursor(prev),
	}
}

func encodeCursor(cursor paginateentity.Cursor) string {
	if len(cursor) == 0 {
		return ""
	}
	serialized, err := json.Marshal(cursor)
	if err != nil {
		return ""
	}
	encodeCursor := base64.StdEncoding.EncodeToString(serialized)
	return encodeCursor
}
