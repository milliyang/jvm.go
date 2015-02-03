package reflect

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _reflection(getCallerClass, "getCallerClass", "(I)Ljava/lang/Class;")
}

func _reflection(method Any, name, desc string) {
    rtc.RegisterNativeMethod("sun/reflect/Reflection", name, desc, method)
}

// public static native Class<?> getCallerClass(int i);
// (I)Ljava/lang/Class;
func getCallerClass(frame *rtda.Frame) {
    stack := frame.OperandStack()
    i := uint(stack.PopInt())
    callerClass := frame.Thread().TopNFrame(i).Method().Class().JClass()
    stack.PushRef(callerClass)
}