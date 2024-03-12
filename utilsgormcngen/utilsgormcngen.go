package utilsgormcngen

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func AssertDone(err error) {
	if err != nil {
		panic(errors.WithMessage(err, "wrong"))
	}
}

func AssertEquals[T comparable](a, b T) {
	if a != b {
		panic(errors.New("not equals"))
	}
}

func SoftNeatString(v interface{}) string {
	data, err := NeatBytes(v)
	if err != nil {
		return "" //when the result is empty string, means wrong
	}
	return string(data)
}

func NeatBytes(v interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, errors.WithMessage(err, "marshal object is wrong")
	}
	return data, nil
}
