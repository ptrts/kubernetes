package utils

import (
	"k8s.io/klog/v2"
	"runtime"
)

func LogAllGoroutines() {
	// Делаем большой буфер
	// 1 MiB, при необходимости — больше
	buf := make([]byte, 1<<20)

	// true = дамп всех горутин
	n := runtime.Stack(buf, true)
	klog.ErrorS(nil, "all goroutines stack trace", "stack", string(buf[:n]))
}
