package RandomHeaderGoLang

import (
	"fmt"
	"strings"

	"github.com/Makrorof/RandomHeaderGoLang/Headers"
)

func GenerateHeadersByNames(names []string) map[string]string {
	return GenerateHeadersByNamesX(names, false, 0)
}

func GenerateHeadersByNamesX(names []string, addSecChrome bool, chromeVersion int) map[string]string {
	headers := make(map[string]string)

	for _, headerName := range names {
		switch strings.ToLower(headerName) {
		case "accept":
			headers["accept"] = Headers.GetAccept()
		case "accept-encoding":
			headers["accept-encoding"] = Headers.GetAcceptEncoding()
		case "accept-language":
			headers["accept-language"] = Headers.GetAcceptLanguage()
		case "Connection":
			headers["Connection"] = Headers.GetConnection()
		case "device-memory":
			headers["device-memory"] = Headers.GetDeviceMemory()
		case "downlink":
			headers["downlink"] = Headers.GetDownlink()
		case "dpr":
			headers["dpr"] = "1"
		case "ect":
			headers["ect"] = "4g"
		case "rtt":
			headers["rtt"] = Headers.GetRTT()
		case "sec-fetch-dest":
			headers["sec-fetch-dest"] = "document"
		case "sec-fetch-mode":
			headers["sec-fetch-mode"] = "navigate"
		case "sec-fetch-site":
			headers["sec-fetch-site"] = "same-origin"
		case "sec-fetch-user":
			headers["sec-fetch-user"] = "?1"
		case "upgrade-insecure-requests":
			headers["upgrade-insecure-requests"] = "1"
		case "viewport-width":
			headers["viewport-width"] = Headers.GetViewportWidth()
		}
	}

	if addSecChrome {
		if chromeVersion == 0 {
			chromeVersion = 110
		}

		headers["sec-ch-ua"] = fmt.Sprintf(`"Chromium";v="%s", "Not A(Brand";v="24", "Google Chrome";v="%s"`, fmt.Sprint(chromeVersion), fmt.Sprint(chromeVersion))
		headers["sec-ch-ua-mobile"] = `?0`
		headers["sec-ch-ua-platform"] = `"Windows"`
		headers["sec-ch-ua-platform-version"] = `"10.0.0"`

		if val, ok := headers["device-memory"]; ok {
			headers["sec-ch-device-memory"] = val
		} else {
			headers["sec-ch-device-memory"] = Headers.GetDeviceMemory()
		}

		if val, ok := headers["dpr"]; ok {
			headers["sec-ch-dpr"] = val
		} else {
			headers["sec-ch-dpr"] = "1"
		}

		if val, ok := headers["viewport-width"]; ok {
			headers["sec-ch-viewport-width"] = val
		} else {
			headers["sec-ch-viewport-width"] = Headers.GetViewportWidth()
		}
	}

	return headers
}
