package models

import (
	"errors"

	"github.com/jinzhu/gorm"
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
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Gender      string `json:"gender"`
	Outfit      string `json:"outfit"`
	Hairstyle   string `json:"hairstyle"`
	Personality string `json:"personality"`
	Tone        string `json:"tone"`
	Profile     string `json:"profile"`
}
