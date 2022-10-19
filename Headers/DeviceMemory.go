package Headers

import "github.com/Makrorof/RandomHTTPHeader/Utils"

var deviceMemoryList []*RandomHeader = []*RandomHeader{
	{optionalVals: []string{"0.25", "0.5", "1", "2", "4", "8"}, singleOptional: true},
}

func GetDeviceMemory() string {
	return deviceMemoryList[Utils.RandomInt(0, len(deviceMemoryList))].GetRandom()
}
