package math

import (
	"testing"
)

/**
go的测试：将路径切换到测试文件所在目录，运行go test命令，go会自动测试所有的测试用例
测试用例的特点是函数名以Test开始，而且具有唯一参数t *testing.T
*/
func TestAdd(t *testing.T) {
	var a = 100
	var b = 200
	var val = Add(a, b)
	if val != a+b {
		t.Error("Test Case [", "TestAdd", "] Failed!")
	}
}
func TestSubtract(t *testing.T) {
	var a = 100
	var b = 200
	var val = Subtract(a, b)
	if val != a-b {
		t.Error("Test Case [", "TestSubtract", "] Failed!")
	}
}
func TestMultiply(t *testing.T) {
	var a = 100
	var b = 200
	var val = Multiply(a, b)
	if val != a*b {
		t.Error("Test Case [", "TestMultiply", "] Failed!")
	}
}
func TestDivideNormal(t *testing.T) {
	var a = 100
	var b = 200
	var val = Divide(a, b)
	if val != a/b {
		t.Error("Test Case [", "TestDivideNormal", "] Failed!")
	}
}
