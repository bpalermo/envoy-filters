package filter

import (
	"regexp"
	"strconv"
)

const (
	routingRegexPattern = `\pL\d+-(\d+)`
	hashHeaderEvenValue = "0"
	hashHeaderOddValue  = "1"
)

type Filter struct {
	re *regexp.Regexp
}

func NewFilter() *Filter {
	return &Filter{
		re: regexp.MustCompile(routingRegexPattern),
	}
}

func (f *Filter) Hash(partitionID string) *string {
	matches := f.re.FindStringSubmatch(partitionID)
	if len(matches) < 2 {
		return nil
	}

	num, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return nil
	}

	var result string
	if num%2 == 0 {
		result = hashHeaderEvenValue
	} else {
		result = hashHeaderOddValue
	}

	return &result
}
