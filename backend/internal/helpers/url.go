// Package helpers contains the helper functions for other services used by the api.
// This file contains the functions related to the encoding and decoding of the url codes for threads and comments.
package helpers

import (
	"errors"
	"slices"

	"github.com/sqids/sqids-go"
)

// DecodedUrl is the format of the object that will be encoded in the url code.
type DecodedUrl struct {
	ID       int
	PageType string
}

// Initialised sqids object
var squid *sqids.Sqids
var validTypes = []string{"thread", "comment"}

// InitSquid initialises the sqid object for to allow for the processing of url codes.
// It returns an error on failure.
func InitSquid() error {
	s, err := sqids.New(sqids.Options{
		MinLength: 6,
		Alphabet:  "Zh3p1JdsjbQUFVMCu8mk9ElYr5XvqLTPzOAxye0aiDtcwG4RofWgN7B6HKSI2n",
	})
	if err != nil {
		return err
	}
	squid = s
	return nil
}

// EncodeUrl takes in the id of the thread/comment and the type (thread/comment) of the page being encoded.
// It returns the encoded url code string on success, and an error on failure.
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

// DecodeUrl takes in the url code string and decodes the code.
// It returns the contents encoded in the url code on success, and an error on failure.
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
		ID:       did,
		PageType: dpt,
	}

	return durl, nil
}
