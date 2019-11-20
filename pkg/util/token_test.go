package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var TOKEN string = ""
var UID string = "coffee377"

func TestGenToken(t *testing.T) {
	TOKEN = GenToken(UID, time.Second)
	Convey("GEN TOKEN", t, func() {
		So(TOKEN, ShouldNotBeNil)
	})
}

func TestValidToken(t *testing.T) {
	//log.Printf("TOKEN：%s", TOKEN)
	for i := 0; i < 5; i++ {
		valid, uid := ValidToken(TOKEN)
		Convey("TOKEN VALID", t, func() {
			So(valid, ShouldEqual, i <= 1)
		})
		Convey("TOKEN UID:", t, func() {
			So(uid, ShouldEqual, UID)
		})
		time.Sleep(time.Second)
	}
}
