package sets

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("Just", t, func() {
		sets := New[string]()
		sets.Add("tx")
		sets.Adds("ali", "toutiao", "tx")

		slice := sets.ToSlice()
		t.Logf("slice:%+v", slice)

		So(sets.Len(), ShouldAlmostEqual, 3)
		So(sets.Contains("tx"), ShouldBeTrue)
	})

}

func TestJust(t *testing.T) {
	Convey("Just", t, func() {
		sets := Just(1, 2, 3)
		So(sets.Contains(1), ShouldBeTrue)
		So(sets.Contains(4), ShouldBeFalse)
	})
}

func Test_Set_Remove(t *testing.T) {
	Convey("Remove", t, func() {
		sets := Just(1, 2, 3)
		after := sets.Remove(1)
		So(after.Len(), ShouldEqual, 2)
		So(after.Contains(1), ShouldBeFalse)

		after = sets.Remove(1)
		So(after.Len(), ShouldEqual, 2)
		So(after.Contains(2), ShouldBeTrue)
		So(after.Contains(3), ShouldBeTrue)
	})
}

func Test_Set_Removes(t *testing.T) {
	Convey("Removes", t, func() {
		sets := Just("tx", "ali", "baidu")
		after := sets.Removes("baidu", "toutiao")
		So(after.Len(), ShouldEqual, 2)
		So(after.Contains("baidu"), ShouldBeFalse)

		after = sets.Removes("baidu", "ali")
		So(after.Len(), ShouldEqual, 1)
		So(after.Contains("tx"), ShouldBeTrue)
		So(after.Contains("ali"), ShouldBeFalse)
	})
}
