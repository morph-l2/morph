//go:build !musl
// +build !musl

package zstd

/*
#cgo LDFLAGS: ${SRCDIR}/libscroll_zstd_linux_amd64.a
*/
import "C"
