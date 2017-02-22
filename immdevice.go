package main

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMMDevice struct {
	ole.IUnknown
}

type IMMDeviceVtbl struct {
	ole.IUnknownVtbl
	Activate          uintptr
	OpenPropertyStore uintptr
	GetId             uintptr
	GetState          uintptr
}

func (v *IMMDevice) VTable() *IMMDeviceVtbl {
	return (*IMMDeviceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMMDevice) Activate(refIID *ole.GUID, ctx uint32, param, obj interface{}) (err error) {
	err = activate(v, refIID, ctx, param, obj)
	return
}

func (v *IMMDevice) GetId(strId *uint16) (err error) {
	err = getId(v, strId)
	return
}

func (v *IMMDevice) GetState(state *uint32) (err error) {
	err = getState(v, state)
	return
}
