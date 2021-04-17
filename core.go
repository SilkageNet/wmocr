package wanmei

import (
	"errors"
	"sync"
	"syscall"
	"unsafe"
)

var (
	core   *wmCore
	locker sync.Mutex
)

var (
	ErrWMAlreadySetup  = errors.New("wm.already.setup")
	ErrWMLoadDatFile   = errors.New("wm.load.dat.file.error")
	ErrWMRecognizeFile = errors.New("wm.recognize.file.error")
)

type wmCore struct {
	dll *syscall.LazyDLL
}

func newWmCore(dllPath string) (*wmCore, error) {
	var dll = syscall.NewLazyDLL(dllPath)
	var err = dll.Load()
	if err != nil {
		return nil, err
	}
	return &wmCore{dll: dll}, err
}

func (c *wmCore) LoadWmFromFile(datPath, password string) (uintptr, error) {
	var fn = c.dll.NewProc("LoadWmFromFileEx")
	var pathBuf = append([]byte(datPath), 0)
	var passBuf = append([]byte(password), 0)
	var id, _, err = fn.Call(uintptr(unsafe.Pointer(&pathBuf[0])), uintptr(unsafe.Pointer(&passBuf[0])))
	if err != nil {
		return 0, err
	}
	if id == -1 {
		return 0, ErrWMLoadDatFile
	}
	return id, nil
}

func (c *wmCore) RecognizeFile(id uintptr, imgPath string) (string, error) {
	var fn = c.dll.NewProc("GetImageFromFileEx")
	var pathBuf = append([]byte(imgPath), 0)
	var retBuf = make([]byte, 5000, 5000)
	var ret, _, err = fn.Call(id, uintptr(unsafe.Pointer(&pathBuf[0])), uintptr(unsafe.Pointer(&retBuf[0])))
	if err != nil {
		return "", err
	}
	if ret != 0 {
		return "", ErrWMRecognizeFile
	}
	var l int
	for i, v := range retBuf {
		if v == 0 {
			l = i + 1
			break
		}
	}
	retBuf = retBuf[:l]
	return string(retBuf), nil
}

func Setup(dllPath string) error {
	locker.Lock()
	defer locker.Unlock()
	if core != nil {
		return ErrWMAlreadySetup
	}
	var err error
	core, err = newWmCore(dllPath)
	return err
}
