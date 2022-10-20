package RandomHeaderGoLang

import "github.com/Makrorof/RandomHeaderGoLang/Headers"

func GenerateHeaderHR(host string, referer string) map[string][]string {
	header := make(map[string][]string)

	//HOST
	//Referer

	if host != "" {
		header["HOST"] = []string{host}
	}

	if referer != "" {
		header["Referer"] = []string{referer}
	}

	header["Accept"] = []string{Headers.GetAccept()}
	header["Accept-Encoding"] = []string{Headers.GetAcceptEncoding()}
	header["Accept-Language"] = []string{Headers.GetAcceptLanguage()}
	header["Cache-Control"] = []string{Headers.GetCacheControl()}
	header["Connection"] = []string{Headers.GetConnection()}
	header["downlink"] = []string{Headers.GetDownlink()}
	header["ect"] = []string{"4g"}
	header["rtt"] = []string{Headers.GetRTT()}
	header["Access-Control-Allow-Origin"] = []string{"*"}

	memory := Headers.GetDeviceMemory()
	header["device-memory"] = []string{memory}
	header["sec-ch-device-memory"] = []string{memory}

	viewportWidth := Headers.GetViewportWidth()
	header["sec-ch-viewport-width"] = []string{viewportWidth}
	header["viewport-width"] = []string{viewportWidth}
	//
	return header
}

func GenerateHeader() map[string][]string {
	return GenerateHeaderHR("", "")
}

func GenerateHeaderR(referer string) map[string][]string {
	return GenerateHeaderHR("", referer)
}

func GenerateHeaderH(host string) map[string][]string {
	return GenerateHeaderHR(host, "")
}
