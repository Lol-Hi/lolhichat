package helpers

import (
	"github.com/sqids/sqids-go"
	"slices"
	"errors"
)

type DecodedUrl struct {
	id				int
	pageType	string
}

var squid *sqids.Sqids
var validTypes = []string{"thread", "comment"}

func InitSquid() error {
	s, err := sqids.New(sqids.Options{
		MinLength: 6,
		Alphabet: "Zh3p1JdsjbQUFVMCu8mk9ElYr5XvqLTPzOAxye0aiDtcwG4RofWgN7B6HKSI2n",
	})
	if err != nil {
		return err
	}
	squid = s
	return nil
}

func EncodeUrl(id int, pageType string) (string, error) {
	if !slices.Contains(validTypes, pageType) {
		return "", errors.New("Invalid page type")
	}

	numSlice := make([]uint64, len(pageType)+1)
	numSlice[0] = uint64(id)
	for i, ch := range pageType {
		numSlice[i+1] = uint64(ch)
	}

	urlCode, err := squid.Encode(numSlice)
	if err != nil {
		return "", err
	}
	return urlCode, nil
}

func DecodeUrl(urlCode string) (*DecodedUrl, error) {
	numSlice := squid.Decode(urlCode)

	runeSlice := make([]rune, len(numSlice)-1)
	for i, val := range numSlice {
		if i > 0 {
			runeSlice[i-1] = rune(val)
		}
	}

	did := int(numSlice[0])
	dpt := string(runeSlice)

	if !slices.Contains(validTypes, dpt) {
		return nil, errors.New("Invalid url code")
	}

	durl := &DecodedUrl{
		id: did,
		pageType: dpt,
	}

	return durl, nil
}

