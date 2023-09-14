package models

import (
	"errors"
	"strings"

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
	Personality string
	Tone        string
	Profile     string
}

func (c *Character) getGenderText() string {
	gender, _ := ChoiceGender(c.Gender)

	switch gender {
	case "male":
		return "男"
	case "female":
		return "女"
	case "other":
	default:
		return ""
	}

	return ""
}

func (c *Character) BuildPrompt() string {
	var lines []string

	lines = append(lines, c.Name)

	nickname := strings.TrimSpace(c.Nickname)
	if len(nickname) > 0 {
		lines = append(lines, "名前: "+nickname)
	}

	gender := c.getGenderText()
	if len(gender) > 0 {
		lines = append(lines, gender)
	}

	outfit := strings.TrimSpace(c.Outfit)
	if len(outfit) > 0 {
		lines = append(lines, "容姿: "+outfit)
	}

	personality := strings.TrimSpace(c.Personality)
	if len(personality) > 0 {
		lines = append(lines, "性格: "+personality)
	}

	tone := strings.TrimSpace(c.Tone)
	if len(tone) > 0 {
		lines = append(lines, "口調: "+tone)
	}
	profile := strings.TrimSpace(c.Profile)
	if len(profile) > 0 {
		lines = append(lines, profile)
	}

	return strings.Join(lines, "\n")
}
