package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Type string

const (
	TypeQP     Type = "QP"
	TypeMS     Type = "MS"
	TypeINSERT Type = "INSERT"
	TypePRE    Type = "PRE"
)

type Season string

const (
	SeasonFM Season = "F/M"
	SeasonMJ Season = "M/J"
	SeasonON Season = "O/N"
)

// Paper is the representation of a question paper, mark scheme, or other types of material.
// Only type and season are validated
// Subject code, component code and year are not validated
type Paper struct {
	SubjectCode string // 4-digit subject code
	Component   string // 2-digit component & variant
	Type        Type   // type of paper
	Season      Season // season
	Year        string // 2-digit year
}

func (p Paper) String() string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", p.SubjectCode, p.Component, p.Type, p.Season, p.Year)
}

// case insensitive regex search
var re = regexp.MustCompile(`(?i)^(?<subject>\d{4})\/(?<component>\d{2})\/((?<type>INSERT|PRE|)\/)?(?<season>\w\/\w)\/(?<year>\d{2})$`)

func ParseCode(code string, type_ string) (*Paper, error) {
	if code == "" {
		return nil, errors.New("Paper code not provided.")
	}

	// match regex
	// example match: ["9231/33/INSERT/M/J/21" "9231" "33" "INSERT/" "INSERT" "M/J" "21"]
	match := re.FindStringSubmatch(code)
	if match == nil {
		return nil, errors.New("Invalid paper code.")
	}

	// season and type
	var season Season
	switch strings.ToUpper(match[5]) {
	case "F/M":
		season = SeasonFM
	case "M/J":
		season = SeasonMJ
	case "O/N":
		season = SeasonON
	default:
		return nil, errors.New("Invalid season code.")
	}

	var paperType Type
	switch strings.ToUpper(match[4]) {
	case "INSERT":
		paperType = TypeINSERT
	case "PRE":
		paperType = TypePRE
	case "":
		switch strings.ToUpper(type_) {
		case "QP":
			paperType = TypeQP
		case "MS":
			paperType = TypeMS
		case "INSERT":
			paperType = TypeINSERT
		case "PRE":
			paperType = TypePRE
		case "":
			return nil, errors.New("Paper type not provided.")
		default:
			return nil, errors.New("Invalid paper type.")
		}
	default:
		return nil, errors.New("Invalid paper type.")
	}

	return &Paper{
			SubjectCode: match[1],
			Component:   match[2],
			Type:        paperType,
			Season:      season,
			Year:        match[6],
		},
		nil
}

// TODO write tests
