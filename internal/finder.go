package internal

import (
	"errors"
	"strings"
)

const urlTemplate = "https://bestexamhelp.com/exam/{exam}/{subjectName}-{subjectCode}/20{year}/{subjectCode}_{seasonLetter}{year}_{type}_{component}.pdf"

var subjectMap = map[string]map[string]string{
	"cambridge-o-level": {
		"7707": "accounting",
		"5090": "biology",
		"7115": "business-studies",
		"3204": "bengali",
		"7094": "bangladesh-studies",
		"5070": "chemistry",
		"7100": "commerce",
		"2210": "computer-science",
		"7010": "computer-studies",
		"2281": "economics",
		"1123": "english-language",
		"4024": "mathematics-d",
		"4037": "mathematics-additional",
		"5054": "physics",
		"7110": "principles-of-accounts",
		"2059": "pakistan-studies",
		"4040": "statistics",
	},
	"cambridge-igcse": {
		"0452": "accounting",
		"0508": "arabic-first-language",
		"0400": "art-and-design",
		"0600": "agriculture",
		"0610": "biology",
		"0450": "business-studies",
		"0620": "chemistry",
		"0478": "computer-science",
		"0509": "chinese-first-language",
		"0411": "drama",
		"0455": "economics",
		"0500": "english-first-language",
		"0524": "english-first-language-us",
		"0475": "english-literature",
		"0680": "environmental-management",
		"0454": "enterprise",
		"0648": "food-and-nutrition",
		"0501": "french-first-language",
		"0460": "geography",
		"0457": "global-perspectives",
		"0470": "history",
		"0493": "islamiyat",
		"0580": "mathematics",
		"0606": "mathematics-additional",
		"0607": "mathematics-international",
		"0410": "music",
		"0696": "malay-first-language",
		"0625": "physics",
		"0413": "physical-education",
		"0490": "religious-studies",
		"0653": "science-combined",
		"0654": "sciences-co-ordinated",
		"0495": "sociology",
		"0502": "spanish-first-language",
		"0471": "travel-and-tourism",
	},
	"cambridge-igcse-9-1": {
		"0985": "accounting",
		"7184": "arabic-first-language",
		"0989": "art-and-design",
		"0970": "biology",
		"0986": "business-studies",
		"0971": "chemistry",
		"0984": "computer-science",
		"0994": "drama",
		"0987": "economics",
		"0990": "english-first-language",
		"0992": "english-literature",
		"0976": "geography",
		"0977": "history",
		"0980": "mathematics",
		//"0978": "music", // link broken on website
		"0972": "physics",
		"0995": "physical-education",
		"0973": "sciences-co-ordinated",
	},
	"cambridge-international-a-level": {
		"9706": "accounting",
		"9713": "applied-information-and-communication-technology",
		"9700": "biology",
		"9609": "business",
		"9707": "business-studies",
		"9701": "chemistry",
		"9618": "computer-science",
		"9608": "computer-science",
		"9691": "computing",
		"9708": "economics",
		"9489": "history",
		"9488": "islamic-studies",
		"9084": "law",
		"9709": "mathematics",
		"9231": "mathematics-further",
		"9607": "media-studies",
		"9702": "physics",
		"9990": "psychology",
		"9698": "psychology",
		"9699": "sociology",
	},
}

var seasonMap = map[Season]string{
	SeasonFM: "m",
	SeasonMJ: "s",
	SeasonON: "w",
}

var typeMap = map[Type]string{
	TypeQP:     "qp",
	TypeMS:     "ms",
	TypeINSERT: "in",
	TypePRE:    "pm",
}

func GetPaperUrl(p *Paper) (string, error) {
	var subject, exam string
	var ok bool
	var innerMap map[string]string
	for exam, innerMap = range subjectMap {
		subject, ok = innerMap[p.SubjectCode]
		if ok {
			break
		}
	}

	if !ok {
		return "", errors.New("Cannot find subject code.")
	}

	url := urlTemplate
	url = strings.ReplaceAll(url, "{exam}", exam)
	url = strings.ReplaceAll(url, "{subjectName}", subject)
	url = strings.ReplaceAll(url, "{subjectCode}", p.SubjectCode)
	url = strings.ReplaceAll(url, "{year}", p.Year)
	url = strings.ReplaceAll(url, "{seasonLetter}", seasonMap[p.Season])
	url = strings.ReplaceAll(url, "{type}", typeMap[p.Type])
	url = strings.ReplaceAll(url, "{component}", p.Component)

	if !PageExists(url) {
		return "", errors.New("Could not find paper.")
	}

	return url, nil
}
