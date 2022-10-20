package RandomHeaderGoLang

import (
	"log"
	"testing"

	"github.com/Makrorof/RandomHeaderGoLang/Headers"
	"github.com/Makrorof/RandomHeaderGoLang/Utils"
)

func TestRandomHeader(t *testing.T) {

}

func TestRandomAccept(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetAccept())
	}
}

func TestRandomAcceptEncoding(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetAcceptEncoding())
	}
}

func TestRandomAcceptLanguage(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetAcceptLanguage())
	}
}

func TestRandomCacheControl(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetCacheControl())
	}
}

func TestRandomConnection(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetConnection())
	}
}

func TestRandomDeviceMemory(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetDeviceMemory())
	}
}

func TestRandomDownlink(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetDownlink())
	}
}

func TestRandomRTT(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetRTT())
	}
}

func TestRandomViewportWidth(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println(Headers.GetViewportWidth())
	}
}

func TestGenerateHeader(t *testing.T) {
	Utils.Init()

	for i := 0; i < 10000; i++ {
		log.Println("----------------------")
		log.Println(GenerateHeader())
		log.Println("----------------------")
	}
}
