package Headers

import "github.com/Makrorof/RandomHeaderGoLang/Utils"

var downlinkList []*RandomHeader = []*RandomHeader{
	{optionalVals: []string{"1.5", "1.7", "2", "5", "7", "10", "7.95"}, singleOptional: true},
}

func GetDownlink() string {
	return downlinkList[Utils.RandomInt(0, len(downlinkList))].GetRandom()
}
