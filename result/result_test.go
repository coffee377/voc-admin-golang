package result

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 根据长宽获取面积
func TestResult(t *testing.T) {
	r := Result{Code: "401", Message: "未授权", Data: ""}
	Convey("test result", t, func() {
		Convey("401", func() {
			So(r.HttpStatus(), ShouldEqual, 401)
		})
		Convey("data is not null", func() {
			So(r.Data, ShouldNotBeNil)
		})
	})
}
