package json_test

import (
	"testing"

	"github.com/jslyzt/go-json"

	"github.com/smartystreets/goconvey/convey"
)

type int64Test struct {
	Astr    string `json:"astr,omitempty"`
	Aint64  int64  `json:"aint64,omitempty"`
	Bstr    string `json:"bstr,omitempty"`
	Buint64 uint64 `json:"buint64,omitempty"`
}

var (
	testint64dt = &int64Test{
		Astr:    "999999999999999999",
		Aint64:  999999999999999999,
		Bstr:    "8888888888888888888",
		Buint64: 8888888888888888888,
	}
)

func TestInt64DEncode(t *testing.T) {
	var (
		err  error
		data []byte

		odt = &int64Test{}
	)

	convey.Convey("Marshal", t, func() {
		data, err = json.Marshal(odt)

		convey.So(err, convey.ShouldBeNil)
		convey.So(data, convey.ShouldNotBeNil)
	})

	convey.Convey("Marshal2", t, func() {
		data, err = json.Marshal(testint64dt)

		convey.So(err, convey.ShouldBeNil)
		convey.So(data, convey.ShouldNotBeNil)
	})

	convey.Convey("Unmarshal", t, func() {
		err = json.Unmarshal(data, odt)

		convey.So(err, convey.ShouldBeNil)

		convey.ShouldEqual(testint64dt.Astr, odt.Astr)
		convey.ShouldEqual(testint64dt.Aint64, odt.Aint64)
		convey.ShouldEqual(testint64dt.Bstr, odt.Bstr)
		convey.ShouldEqual(testint64dt.Buint64, odt.Buint64)
	})

}

func TestInt64DEncodeOption(t *testing.T) {
	var (
		err  error
		data []byte

		odt = &int64Test{}
	)

	convey.Convey("Marshal", t, func() {
		data, err = json.MarshalWithOption(odt, json.NoOmitEmpty())

		convey.So(err, convey.ShouldBeNil)
		convey.So(data, convey.ShouldNotBeNil)
	})

	convey.Convey("Marshal2", t, func() {
		data, err = json.MarshalWithOption(testint64dt, json.Int64ToString())

		convey.So(err, convey.ShouldBeNil)
		convey.So(data, convey.ShouldNotBeNil)
	})

	convey.Convey("Unmarshal", t, func() {
		err = json.UnmarshalWithOption(data, odt, json.DecodeInt64ToString())

		convey.So(err, convey.ShouldBeNil)

		convey.ShouldEqual(testint64dt.Astr, odt.Astr)
		convey.ShouldEqual(testint64dt.Aint64, odt.Aint64)
		convey.ShouldEqual(testint64dt.Bstr, odt.Bstr)
		convey.ShouldEqual(testint64dt.Buint64, odt.Buint64)
	})
}
