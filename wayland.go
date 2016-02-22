package wlc

/*
#cgo LDFLAGS: -lwlc
#include <stdlib.h>
#include <wlc/wlc-wayland.h>
*/
import "C"

import "unsafe"

type Resource C.wlc_resource

// GetWLDisplay returns wayland display.
func GetWLDisplay() *C.struct_wl_display {
	return C.wlc_get_wl_display()
}

// HandleFromWLSurface returns view handle from wl_surface resource.
func HandleFromWLSurface(resource *C.struct_wl_resource) View {
	return View(C.wlc_handle_from_wl_surface_resource(resource))
}

// HandleFromWLOutputResource returns output handle from wl_output resource.
func HandleFromWLOutputResource(resource *C.struct_wl_resource) Output {
	return Output(C.wlc_handle_from_wl_output_resource(resource))
}

// HandleFromWLSurfaceResource returns internal wlc surface from wl_surface
// resource.
func HandleFromWLSurfaceResource(resource *C.struct_wl_resource) Resource {
	return Resource(C.wlc_handle_from_wl_surface_resource(resource))
}

// SurfaceGetSize gets surface size.
func SurfaceGetSize(surface Resource) *Size {
	csize := C.wlc_surface_get_size(C.wlc_resource(surface))
	return sizeCtoGo(csize)
}

// SurfaceGetWLResource returns wl_surface resource from internal wlc surface.
func SurfaceGetWLResource(surface Resource) *C.struct_wl_resource {
	return C.wlc_surface_get_wl_resource(C.wlc_resource(surface))
}

// ViewFromSurface turns wl_surface into a wlc view. Returns 0 on failure.
// This will also trigger view.created callback as any view would.
func ViewFromSurface(surface Resource, client *C.struct_wl_client, interf *C.struct_wl_interface, implementation unsafe.Pointer, version, id uint32, userdata unsafe.Pointer) View {
	return View(
		C.wlc_view_from_surface(
			C.wlc_resource(surface),
			client,
			interf,
			implementation,
			C.uint32_t(version),
			C.uint32_t(id),
			userdata,
		))
}

// GetSurface returns internal wlc surface from view handle.
func (v View) GetSurface() Resource {
	return Resource(C.wlc_view_get_surface(C.wlc_handle(v)))
}

// GetRole returns surface role resource from view handle. Return value
// will be nil if the view was not assigned role or created with
// ViewCreateFromSurface().
func (v View) GetRole() *C.struct_wl_resource {
	return C.wlc_view_get_role(C.wlc_handle(v))
}
