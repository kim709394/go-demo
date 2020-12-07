package hello

import (
	"fmt"
	"github.com/kim709394/go-demo/hello"
	"testing"
)

func TestInterConvert(t *testing.T) {
	hello.InterConvert()
}

func TestNullInter(t *testing.T) {
	hello.NullInterface()
}

func TestSwitInter(t *testing.T) {
	var a hello.MyInterface
	a = new(hello.Str)

	hello.SwitInter(a)
}

func TestErr(t *testing.T) {
	fmt.Println(hello.Err(nil))
	fmt.Println(hello.Err(1))
}
