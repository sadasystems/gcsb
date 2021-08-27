package generator

import (
	"math/rand"
	"time"
	"unsafe"
)

// Assert that StringGenerator implements Generator
var _ Generator = (*StringGenerator)(nil)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type (
	// StringGenerator returns randomly generated strings of a fixed length
	StringGenerator struct {
		len int
		src rand.Source
	}

	// StringGeneratorConfig contains config values for the generator
	StringGeneratorConfig struct {
		Length int
		Source rand.Source
	}
)

func NewStringGenerator(cfg StringGeneratorConfig) (StringGenerator, error) {
	ret := StringGenerator{
		len: cfg.Length,
		src: cfg.Source,
	}

	if ret.src == nil {
		ret.src = rand.NewSource(time.Now().UnixNano())
	}

	return ret, nil
}

/*
 * Next returns the next randomly generated value
 *
 * The random string generation method was borrowed from icza
 * See: https://stackoverflow.com/a/31832326/145479
 */
func (s *StringGenerator) Next() string {
	b := make([]byte, s.len)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := s.len-1, s.src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = s.src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
