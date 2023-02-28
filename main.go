package RandomHeaderGoLang

import (
	"github.com/Makrorof/RandomHeaderGoLang/Utils"
)

//https://github.com/mdn/content/tree/main/files/en-us/web/http/headers

func init() {
	//Set seed
	Utils.Init()

	//....
}

//Accept => '*/*', 'text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8', 'text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, */*;q=0.8'
//Accept-Encoding => https://en.wikipedia.org/wiki/HTTP_compression
//Accept-Language '*','en-US,en;q=0.9',	'tr-TR,tr;q=0.9,en-US;q=0.8,en;q=0.7', 'en-US;q=0.8,en;q=0.7', 'en-US,en;q=0.5'
//Cache-Control asagida
//Connection keep-alive, close
/////////////////////////////////////

//HOST
//Referer

//Access-Control-Allow-Origin *

//device-memory Possible values are: 0.25, 0.5, 1, 2, 4, 8.
//sec-ch-device-memory

//downlink en fazla 10, 1.7 falan => 1.5, 1.7, 2, 5, 7, 10, 7.95
//ect 4g
//rtt 50, 125

//sec-ch-viewport-width
//viewport-width => 1920, 960, 2560, 1280

/////////
//Host: 127.0.0.1:65432
//Connection: keep-alive
//Cache-Control: max-age=0
//sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"
//sec-ch-ua-mobile: ?0
//sec-ch-ua-platform: "macOS"
//Upgrade-Insecure-Requests: 1
//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36
//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
//Sec-Fetch-Site: none
//Sec-Fetch-Mode: navigate
//Sec-Fetch-User: ?1
//Sec-Fetch-Dest: document
//Accept-Encoding: gzip, deflate, br
//Accept-Language: en-GB,en-US;q=0.9,en;q=0.8
