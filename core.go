package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc.h>

// internal wlc_interface reference.
extern struct wlc_interface interface_wlc;
extern void init_interface(uint32_t mask);
*/
import "C"

import "unsafe"

// TODO set log handler

// Init
func Init(i *Interface) bool {
	wlcInterface = i
	var enableMask uint32 = 0

	// output
	if i.Output.Created != nil {
		enableMask |= 1 << 0
	}

	if i.Output.Destroyed != nil {
		enableMask |= 1 << 1
	}

	if i.Output.Focus != nil {
		enableMask |= 1 << 2
	}

	if i.Output.Resolution != nil {
		enableMask |= 1 << 3
	}

	if i.Output.Render.Pre != nil {
		enableMask |= 1 << 4
	}

	if i.Output.Render.Post != nil {
		enableMask |= 1 << 5
	}

	// view
	if i.View.Created != nil {
		enableMask |= 1 << 6
	}

	if i.View.Destroyed != nil {
		enableMask |= 1 << 7
	}

	if i.View.Focus != nil {
		enableMask |= 1 << 8
	}

	if i.View.MoveToOutput != nil {
		enableMask |= 1 << 9
	}

	if i.View.Request.Geometry != nil {
		enableMask |= 1 << 10
	}

	if i.View.Request.State != nil {
		enableMask |= 1 << 11
	}

	if i.View.Request.Move != nil {
		enableMask |= 1 << 12
	}

	if i.View.Request.Resize != nil {
		enableMask |= 1 << 13
	}

	if i.View.Render.Pre != nil {
		enableMask |= 1 << 14
	}

	if i.View.Render.Post != nil {
		enableMask |= 1 << 15
	}

	// keyboard
	if i.Keyboard.Key != nil {
		enableMask |= 1 << 16
	}

	// pointer
	if i.Pointer.Button != nil {
		enableMask |= 1 << 17
	}

	if i.Pointer.Scroll != nil {
		enableMask |= 1 << 18
	}

	if i.Pointer.Motion != nil {
		enableMask |= 1 << 19
	}

	// touch
	if i.Touch.Touch != nil {
		enableMask |= 1 << 20
	}

	// compositor
	if i.Compositor.Ready != nil {
		enableMask |= 1 << 21
	}

	// input
	if i.Input.Created != nil {
		enableMask |= 1 << 22
	}

	if i.Input.Destroyed != nil {
		enableMask |= 1 << 23
	}

	// init wlc_interface struct
	C.init_interface(C.uint32_t(enableMask))

	return bool(C.wlc_init(&C.interface_wlc, 0, nil))
}

// Terminate wlc.
func Terminate() {
	C.wlc_terminate()
}

// GetBackendType query the backend wlc is using.
func GetBackendType() BackendType {
	return BackendType(C.wlc_get_backend_type())
}

// Exec program.
func Exec(bin string, args []string) {
	cbin := C.CString(bin)
	defer C.free(unsafe.Pointer(cbin))
	cargs := strSlicetoCArray(args)
	defer freeCStrArray(cargs)
	C.wlc_exec(cbin, cargs)
}

// Run even loop.
func Run() {
	C.wlc_run()
}

// TODO make more go friendly

// HandleSetUserData Link custom data to handle.
// Client must allocate and handle the data as some C type.
func HandleSetUserData(handle Handle, userdata unsafe.Pointer) {
	C.wlc_handle_set_user_data(C.wlc_handle(handle), userdata)
}

// HandleGetUserData Get custom linked user data from handle.
func HandleGetUserData(handle Handle) unsafe.Pointer {
	return C.wlc_handle_get_user_data(C.wlc_handle(handle))
}

// TODO wlc_event_loop*

// EventSourceTimerUpdate Update timer to trigger after delay.
// Returns true on success.
func EventSourceTimerUpdate(source EventSource, ms_delay int32) bool {
	return bool(C.wlc_event_source_timer_update(
		source,
		C.int32_t(ms_delay),
	))
}

// EventSourceRemove Remove event source from event loop.
func EventSourceRemove(source EventSource) {
	C.wlc_event_source_remove(source)
}
