package wanmei

import (
	"errors"
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

func NewWmOcr(fns ...OptionFn) (*WmOcr, error) {
	var opt = &option{DLLPath: DefaultDLLPath}
	for _, fn := range fns {
		fn(opt)
	}
	var dll = syscall.NewLazyDLL(opt.DLLPath)
	var err = dll.Load()
	if err != nil {
		return nil, err
	}
	var ocr = &WmOcr{dll: dll}
	ocr.id, err = ocr.setupDatFile(opt.DatPath, opt.DatPassword)
	if err != nil {
		return nil, err
	}
	return ocr, ocr.setupOption(opt)
}

func (w *WmOcr) setupDatFile(datPath, password string) (uintptr, error) {
	var fn = w.dll.NewProc("LoadWmFromFileEx")
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
		return "", ErrWMRecognize
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

func (w *WmOcr) Recognize(buff []byte) (string, error) {
	var fn = w.dll.NewProc("GetImageFromBufferEx")
	var retBuf = make([]byte, 5000, 5000)
	var ret, _, err = fn.Call(w.id, uintptr(unsafe.Pointer(&buff[0])), uintptr(len(buff)),
		uintptr(unsafe.Pointer(&retBuf[0])))
	if err != nil {
		return "", err
	}
	if ret != 0 {
		return "", ErrWMRecognize
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
