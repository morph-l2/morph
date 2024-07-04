//go:build !musl
// +build !musl

package zstd

/*
#cgo LDFLAGS: ${SRCDIR}/libscroll_zstd_linux_arm64.a
*/
import "C"
