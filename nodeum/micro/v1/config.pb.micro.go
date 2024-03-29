// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: nodeum/micro/v1/config.proto

package microv1

import (
	fmt "fmt"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ConfigService service

func NewConfigServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "ConfigService.ShowConfig",
			Path:    []string{"/config"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "ConfigService.UpdateConfig",
			Path:    []string{"/config"},
			Method:  []string{"PUT"},
			Handler: "rpc",
		},
	}
}

// Client API for ConfigService service

type ConfigService interface {
	// ShowConfig shows the current config
	ShowConfig(ctx context.Context, in *ShowConfigRequest, opts ...client.CallOption) (*ShowConfigResponse, error)
	// UpdateConfig updates the current config
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...client.CallOption) (*UpdateConfigResponse, error)
}

type configService struct {
	c    client.Client
	name string
}

func NewConfigService(name string, c client.Client) ConfigService {
	return &configService{
		c:    c,
		name: name,
	}
}

func (c *configService) ShowConfig(ctx context.Context, in *ShowConfigRequest, opts ...client.CallOption) (*ShowConfigResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigService.ShowConfig", in)
	out := new(ShowConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...client.CallOption) (*UpdateConfigResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigService.UpdateConfig", in)
	out := new(UpdateConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigService service

type ConfigServiceHandler interface {
	// ShowConfig shows the current config
	ShowConfig(context.Context, *ShowConfigRequest, *ShowConfigResponse) error
	// UpdateConfig updates the current config
	UpdateConfig(context.Context, *UpdateConfigRequest, *UpdateConfigResponse) error
}

func RegisterConfigServiceHandler(s server.Server, hdlr ConfigServiceHandler, opts ...server.HandlerOption) error {
	type configService interface {
		ShowConfig(ctx context.Context, in *ShowConfigRequest, out *ShowConfigResponse) error
		UpdateConfig(ctx context.Context, in *UpdateConfigRequest, out *UpdateConfigResponse) error
	}
	type ConfigService struct {
		configService
	}
	h := &configServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ConfigService.ShowConfig",
		Path:    []string{"/config"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ConfigService.UpdateConfig",
		Path:    []string{"/config"},
		Method:  []string{"PUT"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&ConfigService{h}, opts...))
}

type configServiceHandler struct {
	ConfigServiceHandler
}

func (h *configServiceHandler) ShowConfig(ctx context.Context, in *ShowConfigRequest, out *ShowConfigResponse) error {
	return h.ConfigServiceHandler.ShowConfig(ctx, in, out)
}

func (h *configServiceHandler) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, out *UpdateConfigResponse) error {
	return h.ConfigServiceHandler.UpdateConfig(ctx, in, out)
}
