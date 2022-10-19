package Headers

import "github.com/Makrorof/RandomHTTPHeader/Utils"

var viewportWidthList []*RandomHeader = []*RandomHeader{
	{optionalVals: []string{"1920", "960", "2560", "1280"}, singleOptional: true},
}

func GetViewportWidth() string {
	return viewportWidthList[Utils.RandomInt(0, len(viewportWidthList))].GetRandom()
}
