//go:build wiondows
// +build wiondows

package wmocr

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	ErrWMLoadDatFile = errors.New("wm.load.dat.file.error")
	ErrWMRecognize   = errors.New("wm.recognize.error")
)

type WmOcr struct {
	dll *syscall.LazyDLL
	id  uintptr
}

func NewWmOcr(dllPath, datPath, datPassword string, fns ...OptionFn) (*WmOcr, error) {
	var opt = &option{}
	for _, fn := range fns {
		fn(opt)
	}
	var dll = syscall.NewLazyDLL(dllPath)
	var err = dll.Load()
	if err != nil {
		return nil, fmt.Errorf("load dll err: %v", err)
	}
	var ocr = &WmOcr{dll: dll}
	ocr.id, err = ocr.setupDatFile(datPath, datPassword)
	if err != nil {
		return nil, fmt.Errorf("setup dat file err: %v", err)
	}
	err = ocr.setupOption(opt)
	if err != nil {
		return nil, fmt.Errorf("setup option err: %v", err)
	}
	return ocr, err
}

func (w *WmOcr) setupDatFile(datPath, password string) (uintptr, error) {
	var fn = w.dll.NewProc("LoadWmFromFileEx")
	var pathBuf = append([]byte(datPath), 0)
	var passBuf = append([]byte(password), 0)
	var id, _, err = fn.Call(uintptr(unsafe.Pointer(&pathBuf[0])), uintptr(unsafe.Pointer(&passBuf[0])))
	if err != nil {
		//log.Println(err)
	}
	if int(id) == -1 {
		return 0, ErrWMLoadDatFile
	}
	return id, nil
}

func (w *WmOcr) setupOption(opt *option) error {
	var err error
	var fn = w.dll.NewProc("SetWmOptionEx")
	if opt.RetType != 0 {
		_, _, err = fn.Call(w.id, 1, uintptr(opt.RetType))
		if err != nil {
			return err
		}
	}
	if opt.SegmentationType != 0 {
		_, _, err = fn.Call(w.id, 2, uintptr(opt.SegmentationType))
		if err != nil {
			return err
		}
	}
	if opt.RecognizeType != 0 {
		_, _, err = fn.Call(w.id, 3, uintptr(opt.RecognizeType))
		if err != nil {
			return err
		}
	}
	if opt.AccelerationType != 0 {
		_, _, err = fn.Call(w.id, 4, uintptr(opt.AccelerationType))
		if err != nil {
			return err
		}
	}
	if opt.AccelerationRet != 0 {
		_, _, err = fn.Call(w.id, 5, uintptr(opt.AccelerationRet))
		if err != nil {
			return err
		}
	}
	if opt.MinSimilarity != 0 && opt.MinSimilarity != 90 {
		_, _, err = fn.Call(w.id, 6, uintptr(opt.MinSimilarity))
		if err != nil {
			return err
		}
	}
	if opt.CharSpace != 0 {
		_, _, err = fn.Call(w.id, 7, uintptr(opt.CharSpace))
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *WmOcr) RecognizeFile(imgPath string) (string, error) {
	var fn = w.dll.NewProc("GetImageFromFileEx")
	var pathBuf = append([]byte(imgPath), 0)
	var retBuf = make([]byte, 5000, 5000)
	var ret, _, err = fn.Call(w.id, uintptr(unsafe.Pointer(&pathBuf[0])), uintptr(unsafe.Pointer(&retBuf[0])))
	if err != nil {
		return "", err
	}
	if ret != 0 {
		//log.Println(err)
	}
	var l int
	for i, v := range retBuf {
		if v == 0 {
			l = i + 1
			break
		}
	}
	if l > 0 {
		retBuf = retBuf[:l-1]
	}
	return string(retBuf), nil
}

func (w *WmOcr) Recognize(buff []byte) (string, error) {
	var fn = w.dll.NewProc("GetImageFromBufferEx")
	var retBuf = make([]byte, 5000, 5000)
	var ret, _, err = fn.Call(w.id, uintptr(unsafe.Pointer(&buff[0])), uintptr(len(buff)),
		uintptr(unsafe.Pointer(&retBuf[0])))
	if err != nil {
		//log.Println(err)
	}
	if ret == 0 {
		return "", ErrWMRecognize
	}
	var l int
	for i, v := range retBuf {
		if v == 0 {
			l = i + 1
			break
		}
	}
	if l > 0 {
		retBuf = retBuf[:l-1]
	}
	return string(retBuf), nil
}
