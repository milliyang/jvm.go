package lang

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"math"
)

func init() {
	_double(doubleToRawLongBits, "doubleToRawLongBits", "(D)J")
	_double(longBitsToDouble, "longBitsToDouble", "(J)D")
}

func _double(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/Double", name, desc, method)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *rtda.Frame) {
	vars := frame.LocalVars()
	value := vars.GetDouble(0)

	// todo
	bits := math.Float64bits(value)
	stack := frame.OperandStack()
	stack.PushLong(int64(bits))
}

// public static native double longBitsToDouble(long bits);
// (J)D
func longBitsToDouble(frame *rtda.Frame) {
	vars := frame.LocalVars()
	bits := vars.GetLong(0)

	// todo
	value := math.Float64frombits(uint64(bits))
	stack := frame.OperandStack()
	stack.PushDouble(value)
}
