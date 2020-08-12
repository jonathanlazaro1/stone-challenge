package controller

import (
	"net/url"
	"strconv"
)

// ParseParamToInt gets an URL param and tries to convert it to an integer
func ParseParamToInt(query url.Values, key string, defaultValue int) (int, error) {
	sParam := query.Get(key)
	if sParam == "" {
		return defaultValue, nil
	}
	param, err := strconv.Atoi(sParam)
	if err != nil {
		return -1, err
	}

	return param, nil
}
