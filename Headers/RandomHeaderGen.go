package Headers

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Makrorof/RandomHTTPHeader/Utils"
)

type RandomHeader struct {
	staticVals   []string
	optionalVals []string

	singleOptional bool
	right          bool

	optionalMin int
	optionalMax int
}

func (data *RandomHeader) GetRandom() string {
	static := ""
	optional := ""
	optionalCount := 0
	optionalVals := data.optionalVals

	if len(data.staticVals) > 0 {
		for _, accept := range data.staticVals {
			if strings.Contains(accept, "%d") {
				static += fmt.Sprintf(accept, Utils.RandomInt(6, 10)) + ","
			} else {
				static += accept + ","
			}
		}

		static = static[:len(static)-1]
	}

	if data.singleOptional && len(data.optionalVals) > 0 {
		optionalCount = 1
	} else if !data.singleOptional {
		optionalCount = Utils.RandomInt(0, len(data.optionalVals)+1)
	}

	if optionalCount == 0 {
		return static
	}

	if data.optionalMin == 0 {
		data.optionalMin = 1
		data.optionalMax = 9
	}

	rand.Shuffle(len(optionalVals), func(i, j int) { optionalVals[i], optionalVals[j] = optionalVals[j], optionalVals[i] })

	for i := 0; i < optionalCount; i++ {
		if strings.Contains(optionalVals[i], "%d") {
			optional += fmt.Sprintf(optionalVals[i], Utils.RandomInt(data.optionalMin, data.optionalMax)) + ","
		} else {
			optional += optionalVals[i] + ","
		}
	}

	optional = optional[:len(optional)-1]

	if data.right {
		if optional != "" {
			return optional + "," + static
		} else {
			return static
		}
	}

	if static != "" {
		return static + "," + optional
	} else {
		return optional
	}
}
