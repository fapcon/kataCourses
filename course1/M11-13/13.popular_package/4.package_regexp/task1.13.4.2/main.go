package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Ad struct {
	Title    string
	Описание string
}

func main() {
	ads := []Ad{
		{
			Title:    "Куплю велосипед MeRiDa",
			Описание: "Куплю велосипед meriDA в хорошем состоянии.",
		},
		{
			Title:    "Продам ВаЗ 2101",
			Описание: "Продам ваз 2101 в хорошем состоянии.",
		},
		{
			Title:    "Продам БМВ",
			Описание: "Продам бМв в хорошем состоянии.",
		},
		{
			Title:    "Продам macBook pro",
			Описание: "Продам macBook PRO в хорошем состоянии.",
		},
	}

	ads = censorAds(ads, map[string]string{
		"велосипед merida": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macbook pro":      "Macbook Pro",
	})

	for _, ad := range ads {
		fmt.Println(ad.Title)
		fmt.Println(ad.Описание)
		fmt.Println()
	}
}

func censorAds(ads []Ad, censor map[string]string) []Ad {
	var s string

	keysCensor := make([]string, len(censor))
	for k := range censor {
		for i := 0; i < len(ads); i++ {
			if strings.Contains(strings.ToLower(ads[i].Title), strings.ToLower(k)) {
				keysCensor[i] = k
			}
		}
	}

	for i := 0; i < len(ads); i++ {
		if strings.Contains(strings.ToLower(ads[i].Title), strings.ToLower(keysCensor[0])) {
			var re = regexp.MustCompile(`(?i)(велосипед) (?i)(merida)`)
			if re.MatchString(ads[i].Title) {
				s = keysCensor[i]
				ads[i].Title = re.ReplaceAllString(ads[i].Title, censor[s])
			}
			if re.MatchString(ads[i].Описание) {
				s = keysCensor[i]
				ads[i].Описание = re.ReplaceAllString(ads[i].Описание, censor[s])
			}
		}
		if strings.Contains(strings.ToLower(ads[i].Title), strings.ToLower(keysCensor[1])) {
			var re = regexp.MustCompile(`(?i)(ваз)`)
			if re.MatchString(ads[i].Title) {
				s = keysCensor[i]
				ads[i].Title = re.ReplaceAllString(ads[i].Title, censor[s])
			}
			if re.MatchString(ads[i].Описание) {
				s = keysCensor[i]
				ads[i].Описание = re.ReplaceAllString(ads[i].Описание, censor[s])
			}
		}
		if strings.Contains(strings.ToLower(ads[i].Title), strings.ToLower(keysCensor[2])) {
			var re = regexp.MustCompile(`(?i)(бмв)`)
			if re.MatchString(ads[i].Title) {
				s = keysCensor[i]
				ads[i].Title = re.ReplaceAllString(ads[i].Title, censor[s])
			}
			if re.MatchString(ads[i].Описание) {
				s = keysCensor[i]
				ads[i].Описание = re.ReplaceAllString(ads[i].Описание, censor[s])
			}
		}
		if strings.Contains(strings.ToLower(ads[i].Title), strings.ToLower(keysCensor[3])) {
			var re = regexp.MustCompile(`(?i)(macbook) (?i)(pro)`)
			if re.MatchString(ads[i].Title) {
				s = keysCensor[i]
				ads[i].Title = re.ReplaceAllString(ads[i].Title, censor[s])
			}
			if re.MatchString(ads[i].Описание) {
				s = keysCensor[i]
				ads[i].Описание = re.ReplaceAllString(ads[i].Описание, censor[s])
			}
		}

	}

	return ads
}
