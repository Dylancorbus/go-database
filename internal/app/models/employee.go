package models

import (
	"errors"
	"strconv"
	"strings"
)

type Employee struct {
	Name     string
	Salary   int
	Vacation bool
}

func New(str string) (*Employee, error) {
	if str == "" {
		return nil, errors.New("empty string")
	}
	strArr := strings.Split(str, ";")
	b, err := strconv.ParseBool(strArr[2])
	i, err := strconv.ParseInt(strArr[1], 0, 36)
	if err != nil {
		return nil, err
	}
	emp := Employee{strArr[0], int(i), b}
	return &emp, nil
}