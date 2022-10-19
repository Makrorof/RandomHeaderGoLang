package Headers

import "github.com/Makrorof/RandomHTTPHeader/Utils"

var connectionList []*RandomHeader = []*RandomHeader{
	{optionalVals: []string{"keep-alive", "close"}, singleOptional: true},
}

func GetConnection() string {
	return connectionList[Utils.RandomInt(0, len(connectionList))].GetRandom()
}
