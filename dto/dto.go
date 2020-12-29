package dto

import (
	"errors"
)

type CommonDtoFactory struct {
}

func NewDtoFactory() *CommonDtoFactory {
	return &CommonDtoFactory{}
}

func (f *CommonDtoFactory) CreateDto(name string) (Dto, error) {
	if name == "product" {
		return &ProductDto{}, nil
	}
	return nil, errors.New("dto not found")
}

type Dto interface {
	TableToInfo(interface{}) (interface{}, error)
}
