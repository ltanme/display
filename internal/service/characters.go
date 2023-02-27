package service

import (
	"context"
	"display/internal/model"
)

type ICharacters interface {
	GetList(ctx context.Context, in model.ContentGetListInput) (out *model.CharactersOutput, err error)
}

var characters ICharacters

func Characters() ICharacters {
	if characters == nil {
		panic("implement not found for interface ICharacters, forgot register?")
	}
	return characters
}

func RegisterCharacters(i ICharacters) {
	characters = i
}
