// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package waveout

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modwinmm = windows.NewLazySystemDLL("winmm.dll")

	procwaveOutOpen            = modwinmm.NewProc("waveOutOpen")
	procwaveOutClose           = modwinmm.NewProc("waveOutClose")
	procwaveOutGetVolume       = modwinmm.NewProc("waveOutGetVolume")
	procwaveOutSetVolume       = modwinmm.NewProc("waveOutSetVolume")
	procwaveOutPrepareHeader   = modwinmm.NewProc("waveOutPrepareHeader")
	procwaveOutUnprepareHeader = modwinmm.NewProc("waveOutUnprepareHeader")
	procwaveOutWrite           = modwinmm.NewProc("waveOutWrite")
	procwaveOutPause           = modwinmm.NewProc("waveOutPause")
	procwaveOutRestart         = modwinmm.NewProc("waveOutRestart")
	procwaveOutReset           = modwinmm.NewProc("waveOutReset")
	procwaveOutBreakLoop       = modwinmm.NewProc("waveOutBreakLoop")
)

func Open(handle *syscall.Handle, deviceID uint32, waveFormat *WaveFormatEx, callback uintptr, inst uint32, flag uint32) (result MMRESULT) {
	r0, _, _ := syscall.Syscall6(procwaveOutOpen.Addr(), 6, uintptr(unsafe.Pointer(handle)), uintptr(deviceID), uintptr(unsafe.Pointer(waveFormat)), uintptr(callback), uintptr(inst), uintptr(flag))
	result = MMRESULT(r0)
	return
}

func Close(handle syscall.Handle) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutClose.Addr(), 1, uintptr(handle), 0, 0)
	result = MMRESULT(r0)
	return
}

func GetVolume(handle syscall.Handle, volume *uint32) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutGetVolume.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(volume)), 0)
	result = MMRESULT(r0)
	return
}

func SetVolume(handle syscall.Handle, volume uint32) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutSetVolume.Addr(), 2, uintptr(handle), uintptr(volume), 0)
	result = MMRESULT(r0)
	return
}

func PrepareHeader(handle syscall.Handle, header *WaveHdr, size uint32) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutPrepareHeader.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(header)), uintptr(size))
	result = MMRESULT(r0)
	return
}

func UnprepareHeader(handle syscall.Handle, header *WaveHdr, size uint32) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutUnprepareHeader.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(header)), uintptr(size))
	result = MMRESULT(r0)
	return
}

func Write(handle syscall.Handle, header *WaveHdr, size uint32) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutWrite.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(header)), uintptr(size))
	result = MMRESULT(r0)
	return
}

func Pause(handle syscall.Handle) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutPause.Addr(), 1, uintptr(handle), 0, 0)
	result = MMRESULT(r0)
	return
}

func Restart(handle syscall.Handle) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutRestart.Addr(), 1, uintptr(handle), 0, 0)
	result = MMRESULT(r0)
	return
}

func Reset(handle syscall.Handle) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutReset.Addr(), 1, uintptr(handle), 0, 0)
	result = MMRESULT(r0)
	return
}

func BreakLoop(handle syscall.Handle) (result MMRESULT) {
	r0, _, _ := syscall.Syscall(procwaveOutBreakLoop.Addr(), 1, uintptr(handle), 0, 0)
	result = MMRESULT(r0)
	return
}
