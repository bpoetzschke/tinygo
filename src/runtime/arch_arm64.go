package runtime

const GOARCH = "arm64"

// The bitness of the CPU (e.g. 8, 32, 64).
const TargetBits = 64

const deferExtraRegs = 0

// Align on word boundary.
func align(ptr uintptr) uintptr {
	return (ptr + 15) &^ 15
}

func getCurrentStackPointer() uintptr {
	return uintptr(stacksave())
}
