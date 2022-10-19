package Headers

import "github.com/Makrorof/RandomHeaderGoLang/Utils"

var rttList []*RandomHeader = []*RandomHeader{
	{optionalVals: []string{"50", "125"}, singleOptional: true},
}

func GetRTT() string {
	return rttList[Utils.RandomInt(0, len(rttList))].GetRandom()
}
