package Headers

func GetSecChrome() map[string]string {
	headers := make(map[string]string)
	headers["sec-ch-device-memory"] = ""
	headers["sec-ch-dpr"] = ""
	headers["sec-ch-ua"] = ""
	headers["sec-ch-ua-mobile"] = ""
	headers["sec-ch-ua-platform"] = ""
	headers["sec-ch-ua-platform-version"] = ""
	headers["sec-ch-viewport-width"] = ""

	return headers
}
