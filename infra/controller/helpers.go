package controller

import (
	"net/url"
	"strconv"
)

// ParsePageNumber parses a page number p from a map
func ParsePageNumber(query url.Values) (int, error) {
	sPage := query.Get("p")
	if sPage == "" {
		return 1, nil
	}
	page, err := strconv.Atoi(sPage)
	if err != nil {
		return -1, err
	}
	return page, nil
}
