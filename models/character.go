package models

import (
	"errors"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   = Gender("male")
	Female = Gender("female")
	Other  = Gender("other")
)

func (g *Gender) toString() {
	// 実装
}

func ChoiceGender(value string) (Gender, error) {
	switch value {
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	case "other":
		return Other, nil
	default:
		return "", errors.New("対応するGenderが見つかりません")
	}
}

type Character struct {
	gorm.Model
	UserID      uint
	Name        string
	Nickname    string
	Gender      string
	Outfit      string
	Hairstyle   string
	Personality string
	Tone        string
	Profile     string
}
