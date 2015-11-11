package kexec

import (
	"os"
	"syscall"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommand(t *testing.T) {
	Convey("1 should equal 1", t, func() {
		So(1, ShouldEqual, 1)
	})

	Convey("kexec should work as normal os/exec", t, func() {
		cmd := Command("echo", "-n", "123")
		data, err := cmd.Output()
		So(err, ShouldBeNil)
		So(string(data), ShouldEqual, "123")
	})

	Convey("the terminate should kill proc", t, func() {
		cmd := CommandString("sleep 51")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		time.Sleep(time.Millisecond * 50)
		cmd.Terminate(syscall.SIGINT)
		err := cmd.Wait()
		So(err, ShouldNotBeNil)
		//So(err.Error(), ShouldEqual, "signal: interrupt")
	})
}