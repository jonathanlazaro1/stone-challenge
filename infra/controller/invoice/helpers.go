package invoice

import (
	"net/url"
	"strconv"
)

// ParseParamToInt gets an URL param and tries to convert it to an integer
func parseParamToInt(query url.Values, key string, defaultValue int) (int, error) {
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

// ParseFilterByToMap gets an URL param and tries to convert it to a list of valid query filters
func parseFilterByToMap(query url.Values) (map[string]string, error) {
	filterByMap := make(map[string]string)
	filterBy := query.Get("filter")
	if filterBy == "" {
		return filterByMap, nil
	}

	return filterByMap, nil
}
