// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/service/v1/i_dict.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-monolithic-demo/gen/api/go/system/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDictServiceCreateDict = "/admin.service.v1.DictService/CreateDict"
const OperationDictServiceDeleteDict = "/admin.service.v1.DictService/DeleteDict"
const OperationDictServiceGetDict = "/admin.service.v1.DictService/GetDict"
const OperationDictServiceListDict = "/admin.service.v1.DictService/ListDict"
const OperationDictServiceUpdateDict = "/admin.service.v1.DictService/UpdateDict"

type DictServiceHTTPServer interface {
	// CreateDict 创建字典
	CreateDict(context.Context, *v11.CreateDictRequest) (*emptypb.Empty, error)
	// DeleteDict 删除字典
	DeleteDict(context.Context, *v11.DeleteDictRequest) (*emptypb.Empty, error)
	// GetDict 查询字典
	GetDict(context.Context, *v11.GetDictRequest) (*v11.Dict, error)
	// ListDict 查询字典列表
	ListDict(context.Context, *v1.PagingRequest) (*v11.ListDictResponse, error)
	// UpdateDict 更新字典
	UpdateDict(context.Context, *v11.UpdateDictRequest) (*emptypb.Empty, error)
}

func RegisterDictServiceHTTPServer(s *http.Server, srv DictServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/dicts", _DictService_ListDict0_HTTP_Handler(srv))
	r.GET("/admin/v1/dicts/{id}", _DictService_GetDict0_HTTP_Handler(srv))
	r.POST("/admin/v1/dicts", _DictService_CreateDict0_HTTP_Handler(srv))
	r.PUT("/admin/v1/dicts/{dict.id}", _DictService_UpdateDict0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/dicts/{id}", _DictService_DeleteDict0_HTTP_Handler(srv))
}

func _DictService_ListDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceListDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDict(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListDictResponse)
		return ctx.Result(200, reply)
	}
}

func _DictService_GetDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetDictRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceGetDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDict(ctx, req.(*v11.GetDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Dict)
		return ctx.Result(200, reply)
	}
}

func _DictService_CreateDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateDictRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceCreateDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDict(ctx, req.(*v11.CreateDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _DictService_UpdateDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateDictRequest
		if err := ctx.Bind(&in.Dict); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceUpdateDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDict(ctx, req.(*v11.UpdateDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _DictService_DeleteDict0_HTTP_Handler(srv DictServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteDictRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictServiceDeleteDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDict(ctx, req.(*v11.DeleteDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type DictServiceHTTPClient interface {
	CreateDict(ctx context.Context, req *v11.CreateDictRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteDict(ctx context.Context, req *v11.DeleteDictRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetDict(ctx context.Context, req *v11.GetDictRequest, opts ...http.CallOption) (rsp *v11.Dict, err error)
	ListDict(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListDictResponse, err error)
	UpdateDict(ctx context.Context, req *v11.UpdateDictRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type DictServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewDictServiceHTTPClient(client *http.Client) DictServiceHTTPClient {
	return &DictServiceHTTPClientImpl{client}
}

func (c *DictServiceHTTPClientImpl) CreateDict(ctx context.Context, in *v11.CreateDictRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/dicts"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictServiceCreateDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) DeleteDict(ctx context.Context, in *v11.DeleteDictRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/dicts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceDeleteDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) GetDict(ctx context.Context, in *v11.GetDictRequest, opts ...http.CallOption) (*v11.Dict, error) {
	var out v11.Dict
	pattern := "/admin/v1/dicts/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceGetDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) ListDict(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListDictResponse, error) {
	var out v11.ListDictResponse
	pattern := "/admin/v1/dicts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictServiceListDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictServiceHTTPClientImpl) UpdateDict(ctx context.Context, in *v11.UpdateDictRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/dicts/{dict.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictServiceUpdateDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Dict, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
