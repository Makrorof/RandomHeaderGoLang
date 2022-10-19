package Headers

import (
	"github.com/Makrorof/RandomHeaderGoLang/Utils"
)

var acceptList []*RandomHeader = []*RandomHeader{
	{staticVals: []string{"*/*"}},
	{staticVals: []string{"text/html", "*/*"}},
	{staticVals: []string{"text/html", "*/*;q=0.%d"}, optionalVals: []string{"application/xhtml+xml", "application/xml;q=0.%d", "image/webp", "image/avif", "image/apng", "application/signed-exchange;v=b3;q=0.%d"}},
}

//"*/*",
//"text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8",
//"text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, */*;q=0.8",

func GetAccept() string {
	return acceptList[Utils.RandomInt(0, len(acceptList))].GetRandom()
}
