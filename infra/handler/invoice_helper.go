package handler

import (
	"net/url"
	"strconv"
	"strings"
)

// isValidInvoiceParam checks if a string is a valid Invoice filter/sort param
func isValidInvoiceParam(str string) bool {
	params := [3]string{"reference_year", "reference_month", "document"}
	for _, p := range params {
		if p == str {
			return true
		}
	}
	return false
}

// ParseParamToInt gets an URL param and tries to convert it to an integer.
// If the URL param is empty, it returns the default value.
// If the converted value is below 1, it retuns 1.
// If the conversion to an integer does not succeed, it returns an error.
func parseParamToInt(query url.Values, key string, defaultValue int) (int, error) {
	sParam := query.Get(key)
	if sParam == "" {
		return defaultValue, nil
	}
	param, err := strconv.Atoi(sParam)
	if err != nil {
		return -1, err
	}
	if param < 1 {
		param = 1
	}

	return param, nil
}

// parseFilterByToMap gets URL param "filter" and tries to convert it to a list of valid Invoice filter params
func parseFilterByToMap(query url.Values) map[string]string {
	ret := map[string]string{}
	filterQuery := query.Get("filter")
	if filterQuery == "" {
		return ret
	}

	filters := strings.Split(filterQuery, ",")

	for _, filter := range filters {
		q := strings.Split(filter, ":")
		if len(q) < 2 {
			continue
		}
		if !isValidInvoiceParam(q[0]) {
			continue
		}
		ret[q[0]] = q[1]
	}
	return ret
}

// parseSortByToMap gets URL param "sort" and tries to convert it to a list of valid Invoice sort params
func parseSortByToMap(query url.Values) map[string]bool {
	ret := map[string]bool{}

	sortQuery := query.Get("sort")
	if sortQuery == "" {
		return ret
	}

	sorts := strings.Split(sortQuery, ",")

	for _, sort := range sorts {
		q := strings.Split(sort, ":")

		if !isValidInvoiceParam(q[0]) {
			continue
		}
		descending := false
		if len(q) > 1 && q[1] == "desc" {
			descending = true
		}
		ret[q[0]] = descending

	}
	return ret
}
