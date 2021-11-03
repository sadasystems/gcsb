package schema

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"cloud.google.com/go/spanner/spansql"
)

const (
	maxString = 2621440  // https://cloud.google.com/spanner/docs/data-definition-language#string
	maxByte   = 10485760 // https://cloud.google.com/spanner/docs/data-definition-language#bytes
)

var lengthRegexp = regexp.MustCompile(`\(([0-9]+|MAX)\)$`)

// TODO: Return an error rather than panic
func ParseSpannerType(spannerType string) spansql.Type {
	ret := spansql.Type{}

	dt := spannerType

	// separate type and length from dt with length such as STRING(32) or BYTES(256)
	m := lengthRegexp.FindStringSubmatchIndex(dt)
	if m != nil {
		lengthStr := dt[m[2]:m[3]]
		if lengthStr == "MAX" {
			ret.Len = spansql.MaxLen
		} else {
			l, err := strconv.Atoi(lengthStr)
			if err != nil {
				panic("could not convert precision")
			}
			ret.Len = int64(l)
		}

		// trim length from dt
		dt = dt[:m[0]] + dt[m[1]:]
	}

	if strings.HasPrefix(dt, "ARRAY<") {
		ret.Array = true
		dt = strings.TrimSuffix(strings.TrimPrefix(dt, "ARRAY<"), ">")
	}

	ret.Base = parseType(dt)
	if ret.Len == spansql.MaxLen {
		switch ret.Base {
		case spansql.String:
			ret.Len = maxString
		case spansql.Bytes:
			ret.Len = maxByte
		default:
			ret.Len = 1024 // This should never happen
		}
	}

	return ret
}

func parseType(dt string) spansql.TypeBase {
	var ret spansql.TypeBase
	switch dt {
	case "BOOL":
		ret = spansql.Bool
	case "STRING":
		ret = spansql.String
	case "INT64":
		ret = spansql.Int64
	case "FLOAT64":
		ret = spansql.Float64
	case "BYTES":
		ret = spansql.Bytes
	case "TIMESTAMP":
		ret = spansql.Timestamp
	case "DATE":
		ret = spansql.Date
	case "NUMERIC":
		ret = spansql.Numeric
	default:
		panic(fmt.Sprintf("unknown spanner type '%s'", dt)) // TODO: return error. dont panic
	}

	return ret
}
