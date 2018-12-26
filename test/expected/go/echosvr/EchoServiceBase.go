// Code generated by protoapi:go; DO NOT EDIT.

package echosvr

import (
	"github.com/labstack/echo"
	"github.com/yoozoo/protoapi/protoapigo"
)

// EchoService is the interface contains all the controllers
type EchoService interface {
	Echo(c echo.Context, req *Msg) (resp *Msg, err error)
}

func _echo_Handler(srv EchoService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(Msg)

		if err = c.Bind(req); err != nil {

			return c.JSON(500, err)

		}
		/*

		 */
		resp, err := srv.Echo(c, req)
		if err != nil {

			return c.String(500, err.Error())
		}

		return c.JSON(200, resp)
	}
}

// RegisterEchoService is used to bind routers
func RegisterEchoService(e *echo.Echo, srv EchoService) {
	RegisterEchoServiceWithPrefix(e, srv, "")
}

// RegisterEchoServiceWithPrefix is used to bind routers with custom prefix
func RegisterEchoServiceWithPrefix(e *echo.Echo, srv EchoService, prefix string) {
	// switch to strict JSONAPIBinder, if using echo's DefaultBinder
	if _, ok := e.Binder.(*echo.DefaultBinder); ok {
		e.Binder = new(protoapigo.JSONAPIBinder)
	}
	e.POST(prefix+"/EchoService.echo", _echo_Handler(srv))
}
