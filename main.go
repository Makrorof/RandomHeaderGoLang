package main

import (
	"github.com/Makrorof/RandomHTTPHeader/Utils"
)

//https://github.com/mdn/content/tree/main/files/en-us/web/http/headers

func main() {
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
