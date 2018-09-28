// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

import (
	"regexp"

	"github.com/labstack/echo"
)

const (
	Email string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

var (
	rxEmail = regexp.MustCompile(Email)
)

// AppService is the interface contains all the controllers
type AppService interface {
	GetEnv(echo.Context, *EnvListRequest) (*EnvListResponse, *Error)
	RegisterService(echo.Context, *RegisterServiceRequest) (*RegisterServiceResponse, *Error)
	UpdateService(echo.Context, *UpdateServiceRequest) (*UpdateServiceResponse, *Error)
	UploadProtoFile(echo.Context, *UploadProtoFileRequest) (*UploadProtoFileResponse, *Error)
	GetTags(echo.Context, *TagListRequest) (*TagListResponse, *Error)
	GetProducts(echo.Context, *ProductListRequest) (*ProductListResponse, *Error)
	GetServices(echo.Context, *ServiceListRequest) (*ServiceListResponse, *Error)
	SearchServices(echo.Context, *ServiceSearchRequest) (*ServiceListResponse, *Error)
	GetKeyList(echo.Context, *KeyListRequest) (*KeyListResponse, *Error)
	GetKeyValueList(echo.Context, *KeyValueListRequest) (*KeyValueListResponse, *Error)
	SearchKeyValueList(echo.Context, *SearchKeyValueListRequest) (*KeyValueListResponse, *Error)
	UpdateKeyValue(echo.Context, *KeyValueRequest) (*KeyValueResponse, *Error)
	FetchKeyHistory(echo.Context, *KVHistoryRequest) (*KVHistoryResponse, *Error)
}

func _getEnv_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(EnvListRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.GetEnv(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _registerService_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(RegisterServiceRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.RegisterService(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _updateService_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(UpdateServiceRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.UpdateService(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _uploadProtoFile_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(UploadProtoFileRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.UploadProtoFile(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _getTags_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(TagListRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.GetTags(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _getProducts_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(ProductListRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.GetProducts(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _getServices_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(ServiceListRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.GetServices(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _searchServices_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(ServiceSearchRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.SearchServices(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _getKeyList_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(KeyListRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.GetKeyList(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _getKeyValueList_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(KeyValueListRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.GetKeyValueList(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _searchKeyValueList_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(SearchKeyValueListRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.SearchKeyValueList(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _updateKeyValue_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(KeyValueRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.UpdateKeyValue(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}
func _fetchKeyHistory_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		in := new(KVHistoryRequest)

		if err = c.Bind(in); err != nil {
			resp := CommonError{BindError: &BindError{Message: err.Error()}}
			return c.JSON(420, resp)
		}

		if valErr := in.Validate(); valErr != nil {
			resp := CommonError{ValidateError: valErr}
			return c.JSON(420, resp)
		}

		out, error := srv.FetchKeyHistory(c, in)
		if error != nil {
			return c.JSON(400, error)
		}

		return c.JSON(200, out)
	}
}

// RegisterAppService is used to bind routers
func RegisterAppService(e *echo.Echo, srv AppService) {
	e.POST("/AppService.getEnv", _getEnv_Handler(srv))
	e.POST("/AppService.registerService", _registerService_Handler(srv))
	e.POST("/AppService.updateService", _updateService_Handler(srv))
	e.POST("/AppService.uploadProtoFile", _uploadProtoFile_Handler(srv))
	e.POST("/AppService.getTags", _getTags_Handler(srv))
	e.POST("/AppService.getProducts", _getProducts_Handler(srv))
	e.POST("/AppService.getServices", _getServices_Handler(srv))
	e.POST("/AppService.searchServices", _searchServices_Handler(srv))
	e.POST("/AppService.getKeyList", _getKeyList_Handler(srv))
	e.POST("/AppService.getKeyValueList", _getKeyValueList_Handler(srv))
	e.POST("/AppService.searchKeyValueList", _searchKeyValueList_Handler(srv))
	e.POST("/AppService.updateKeyValue", _updateKeyValue_Handler(srv))
	e.POST("/AppService.fetchKeyHistory", _fetchKeyHistory_Handler(srv))
}
