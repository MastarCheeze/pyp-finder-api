package main

import (
	"errors"
	"regexp"
	"strings"
)

type Type string

const (
	TypeInsert      Type = "INSERT"
	TypePre         Type = "PRE"
	TypeUnspecified Type = "" // QP and MS
)

type Season string

const (
	SeasonFM Season = "F/M"
	SeasonMJ Season = "M/J"
	SeasonON Season = "O/N"
)

// Paper is the representation of a question paper, mark scheme, or other types of material.
// Only type and season are validated, subject code, component code and year are not
type Paper struct {
	Subject   string // 4-digit subject code
	Component string // 2-digit component/variant
	Type      Type   // type of paper
	Season    Season // season
	Year      string // 2-digit year
}

var re = regexp.MustCompile(`^(?<subject>\d{4})\/(?<component>\d{2})\/((?<type>INSERT|PRE)\/)?(?<season>(F\/M|M\/J|O\/N))\/(?<year>\d{2})$`)

func parseCode(code string) (*Paper, error) {
	match := re.FindStringSubmatch(strings.ToUpper(code))
	if match == nil {
		return nil, errors.New("Invalid paper code.")
	}

	// format search result into nice map
	results := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			results[name] = match[i]
		}
	}

	// validate type and season
	var type_ Type
	switch results["type"] {
	case "INSERT":
		type_ = TypeInsert
	case "PRE":
		type_ = TypePre
	default:
		type_ = TypeUnspecified
	}

	var season Season
	switch results["season"] {
	case "F/M":
		season = SeasonFM
	case "M/J":
		season = SeasonMJ
	case "O/N":
		season = SeasonON
	}

	return &Paper{
			Subject:   results["subject"],
			Component: results["component"],
			Type:      type_,
			Season:    season,
			Year:      results["year"],
		},
		nil
}
