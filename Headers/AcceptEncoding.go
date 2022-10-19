package Headers

import "github.com/Makrorof/RandomHTTPHeader/Utils"

var acceptEncodingList []*RandomHeader = []*RandomHeader{
	{staticVals: []string{"*"}},
	{staticVals: []string{"gzip"}, optionalVals: []string{"compress", "deflate", "br", "identity"}},
	{staticVals: []string{"gzip;q=1.0"}, optionalVals: []string{"compress;q=0.%d", "deflate;q=0.%d", "br;q=0.%d", "identity;q=0.%d"}},
	//{staticVals: []string{""}, optionalVals: []string{}},
}

//Accept-Encoding: gzip
//Accept-Encoding: compress
//Accept-Encoding: deflate
//Accept-Encoding: br
//Accept-Encoding: identity
//Accept-Encoding: *

func GetAcceptEncoding() string {
	return acceptEncodingList[Utils.RandomInt(0, len(acceptEncodingList))].GetRandom()
}
