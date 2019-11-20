package result

import (
	errors2 "github.com/coffee377/voc-admin/internal/app/errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestResult(t *testing.T) {
	r := Of(nil, errors2.Unauthorized)
	Convey("test result", t, func() {
		Convey("401", func() {
			So(r.Status, ShouldEqual, 401)
		})
		Convey("data is not null", func() {
			So(r.Data, ShouldBeNil)
		})
	})
}
