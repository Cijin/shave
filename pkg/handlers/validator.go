package handlers

import (
	"context"
	"net/http"

	"shave/pkg/data"
)

/*
Validator only checks for simple things
without the any calls to the database
that would happen in the handler itself
Example it can check to make sure:
* Required fields are not empty
* Strings with a specific format (like email) are correct
* Numbers are within an acceptable range
*/

type validator interface {
	// Valid checks the object and returns any
	// problems. If len(problems) == 0 then
	// the object is valid.
	Valid(ctx context.Context) data.Problems
}

type Decoder interface {
	Decode(dst interface{}, src map[string][]string) error
}

func decodeValid[T validator](r *http.Request, decoder Decoder) (T, data.Problems, error) {
	var v T
	if err := r.ParseForm(); err != nil {
		return v, nil, err
	}

	if err := decoder.Decode(&v, r.PostForm); err != nil {
		return v, nil, err
	}

	problems := v.Valid(r.Context())
	if problems.Any() {
		return v, problems, nil
	}

	return v, nil, nil
}
