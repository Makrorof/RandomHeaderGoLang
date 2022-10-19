package Headers

import "github.com/Makrorof/RandomHTTPHeader/Utils"

var acceptLanguageList []*RandomHeader = []*RandomHeader{
	{staticVals: []string{"*"}},
	{staticVals: []string{}, optionalVals: []string{"en-US,en;q=0.%d", "en-US;q=1.0,en;q=0.%d"}, optionalMin: 7, optionalMax: 10, singleOptional: true},
	{staticVals: []string{"en-US,en;q=1.0"}, optionalVals: []string{"tr-TR,tr;q=0.%d", "da-DK;q=0.%d", "de-DE;q=0.%d", "fr-FR;q=0.%d", "it-IT;q=0.%d", "ru-RU;q=0.%d"}, right: true, singleOptional: true},
	{staticVals: []string{"en-US;q=1.0,en;q=0.9"}, optionalVals: []string{"tr-TR,tr;q=0.%d", "da-DK;q=0.%d", "de-DE;q=0.%d", "fr-FR;q=0.%d", "it-IT;q=0.%d", "ru-RU;q=0.%d"}, right: true, singleOptional: true},
	//{staticVals: []string{""}, optionalVals: []string{}},
}

func GetAcceptLanguage() string {
	return acceptLanguageList[Utils.RandomInt(0, len(acceptLanguageList))].GetRandom()
}
