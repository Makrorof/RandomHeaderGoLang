package Headers

import "github.com/Makrorof/RandomHTTPHeader/Utils"

var cacheControlList []*RandomHeader = []*RandomHeader{
	{optionalVals: []string{"max-age=604800", "no-cache", "max-age=604800, must-revalidate", "no-store", "private", "public"}, singleOptional: true},
}

func GetCacheControl() string {
	return cacheControlList[Utils.RandomInt(0, len(cacheControlList))].GetRandom()
}
