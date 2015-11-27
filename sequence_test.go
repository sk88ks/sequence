package sequence

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSet(t *testing.T) {
	Convey("Given element", t, func() {
		e := Element{}

		Convey("When set a value", func() {
			e.Set("set_test", "This is a test")

			Convey("Then internal map should have the value", func() {
				val, _ := e.items["set_test"]
				str, ok := val.(string)
				So(ok, ShouldBeTrue)
				So(str, ShouldEqual, "This is a test")

			})
		})
	})
}

func TestGet(t *testing.T) {
	Convey("Given element without map", t, func() {
		e := Element{}

		Convey("When get a value", func() {
			_, ok := e.Get("get_test")

			Convey("Then result should be nil", func() {
				So(ok, ShouldBeFalse)

			})
		})
	})

	Convey("Given element with map", t, func() {
		e := Element{}
		e.Set("get_test", "This is a test")

		Convey("When get a value", func() {
			val, ok := e.Get("get_test")

			Convey("Then result should be nil", func() {
				So(ok, ShouldBeTrue)
				str, ok := val.(string)
				So(ok, ShouldBeTrue)
				So(str, ShouldEqual, "This is a test")

			})
		})
	})
}

func TestGetFloat64(t *testing.T) {
	Convey("Given mapped element", t, func() {
		e := Element{}
		var val float64
		val = 1.23456
		e.Set("float_test", val)

		Convey("When get with not exist key", func() {
			res := e.GetFloat64("not_exist")

			Convey("Then zero should be returned", func() {
				So(res, ShouldEqual, 0)

			})
		})

		Convey("When get with valid key", func() {
			res := e.GetFloat64("float_test")

			Convey("Then resut should equal to expected", func() {
				So(res, ShouldEqual, val)

			})
		})
	})
}

func TestGetString(t *testing.T) {
	Convey("Given mapped element", t, func() {
		e := Element{}
		val := "This is a test"
		e.Set("str_test", val)

		Convey("When get with not exist key", func() {
			res := e.GetString("not_exist")

			Convey("Then empty should be returned", func() {
				So(res, ShouldEqual, "")

			})
		})

		Convey("When get with valid key", func() {
			res := e.GetString("str_test")

			Convey("Then resut should equal to expected", func() {
				So(res, ShouldEqual, val)

			})
		})
	})
}

func TestSortByFloat64Desc(t *testing.T) {
	Convey("Given elements", t, func() {
		elements := Elements{}
		for i := 0; i < 5; i++ {
			e := Element{}
			e.Set("score", float64(i*10))
			e.Set("id", "test00"+strconv.Itoa(i))
			elements = append(elements, e)
		}

		Convey("When sort by score desc", func() {
			elements.SortByFloat64Desc("score")

			Convey("Then elements should be desc sorted by its score", func() {
				So(len(elements), ShouldEqual, 5)
				So(elements[0].GetString("id"), ShouldEqual, "test004")
				So(elements[1].GetString("id"), ShouldEqual, "test003")
				So(elements[2].GetString("id"), ShouldEqual, "test002")
				So(elements[3].GetString("id"), ShouldEqual, "test001")
				So(elements[4].GetString("id"), ShouldEqual, "test000")

			})
		})
	})
}

func TestFilter(t *testing.T) {
	Convey("Given elements and filter function", t, func() {
		elements := Elements{}
		for i := 0; i < 5; i++ {
			e := Element{}
			e.Set("score", float64(i*10))
			e.Set("id", "test00"+strconv.Itoa(i))
			elements = append(elements, e)
		}

		f := func(e Element) bool {
			if e.GetFloat64("score") > 10 {
				return true
			}
			return false
		}

		Convey("When filtering", func() {
			res := elements.Filter(f)
			res.SortByFloat64Desc("score")

			Convey("Then result should have filterd elements", func() {
				So(len(res), ShouldEqual, 3)
				So(res[0].GetString("id"), ShouldEqual, "test004")
				So(res[1].GetString("id"), ShouldEqual, "test003")
				So(res[2].GetString("id"), ShouldEqual, "test002")

			})
		})
	})
}

func TestMap(t *testing.T) {
	Convey("Given elements and filter function", t, func() {
		elements := Elements{}
		for i := 0; i < 5; i++ {
			e := Element{}
			e.Set("score", float64(i*10))
			e.Set("id", "test00"+strconv.Itoa(i))
			elements = append(elements, e)
		}

		f := func(e Element) Element {
			score := e.GetFloat64("score")
			e.Set("score", score*2)
			return e
		}

		Convey("When filtering", func() {
			res := elements.Map(f)
			res.SortByFloat64Desc("score")

			Convey("Then result should have filterd elements", func() {
				So(len(res), ShouldEqual, 5)
				So(res[0].GetString("id"), ShouldEqual, "test004")
				So(res[1].GetString("id"), ShouldEqual, "test003")
				So(res[2].GetString("id"), ShouldEqual, "test002")
				So(res[3].GetString("id"), ShouldEqual, "test001")
				So(res[4].GetString("id"), ShouldEqual, "test000")
				So(res[0].GetFloat64("score"), ShouldEqual, 80)
				So(res[1].GetFloat64("score"), ShouldEqual, 60)
				So(res[2].GetFloat64("score"), ShouldEqual, 40)
				So(res[3].GetFloat64("score"), ShouldEqual, 20)
				So(res[4].GetFloat64("score"), ShouldEqual, 0)

			})
		})
	})
}
