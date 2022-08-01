// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: nodeum/micro/v1/dispatcher.proto

package microv1

import (
	fmt "fmt"
	_ "github.com/nodeum-io/nodeum-plugins/nodeum/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
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

// Api Endpoints for DispatcherService service

func NewDispatcherServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "DispatcherService.Start",
			Path:    []string{"/task/start"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "DispatcherService.Pause",
			Path:    []string{"/task/pause/{id}"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "DispatcherService.Resume",
			Path:    []string{"/task/resume/{id}"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "DispatcherService.Stop",
			Path:    []string{"/task/stop/{id}"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
	}
}

// Client API for DispatcherService service

type DispatcherService interface {
	Start(ctx context.Context, in *DispatcherServiceStartRequest, opts ...client.CallOption) (*DispatcherServiceStartResponse, error)
	Pause(ctx context.Context, in *DispatcherServicePauseRequest, opts ...client.CallOption) (*DispatcherServicePauseResponse, error)
	Resume(ctx context.Context, in *DispatcherServiceResumeRequest, opts ...client.CallOption) (*DispatcherServiceResumeResponse, error)
	Stop(ctx context.Context, in *DispatcherServiceStopRequest, opts ...client.CallOption) (*DispatcherServiceStopResponse, error)
}

type dispatcherService struct {
	c    client.Client
	name string
}

func NewDispatcherService(name string, c client.Client) DispatcherService {
	return &dispatcherService{
		c:    c,
		name: name,
	}
}

func (c *dispatcherService) Start(ctx context.Context, in *DispatcherServiceStartRequest, opts ...client.CallOption) (*DispatcherServiceStartResponse, error) {
	req := c.c.NewRequest(c.name, "DispatcherService.Start", in)
	out := new(DispatcherServiceStartResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherService) Pause(ctx context.Context, in *DispatcherServicePauseRequest, opts ...client.CallOption) (*DispatcherServicePauseResponse, error) {
	req := c.c.NewRequest(c.name, "DispatcherService.Pause", in)
	out := new(DispatcherServicePauseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherService) Resume(ctx context.Context, in *DispatcherServiceResumeRequest, opts ...client.CallOption) (*DispatcherServiceResumeResponse, error) {
	req := c.c.NewRequest(c.name, "DispatcherService.Resume", in)
	out := new(DispatcherServiceResumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherService) Stop(ctx context.Context, in *DispatcherServiceStopRequest, opts ...client.CallOption) (*DispatcherServiceStopResponse, error) {
	req := c.c.NewRequest(c.name, "DispatcherService.Stop", in)
	out := new(DispatcherServiceStopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DispatcherService service

type DispatcherServiceHandler interface {
	Start(context.Context, *DispatcherServiceStartRequest, *DispatcherServiceStartResponse) error
	Pause(context.Context, *DispatcherServicePauseRequest, *DispatcherServicePauseResponse) error
	Resume(context.Context, *DispatcherServiceResumeRequest, *DispatcherServiceResumeResponse) error
	Stop(context.Context, *DispatcherServiceStopRequest, *DispatcherServiceStopResponse) error
}

func RegisterDispatcherServiceHandler(s server.Server, hdlr DispatcherServiceHandler, opts ...server.HandlerOption) error {
	type dispatcherService interface {
		Start(ctx context.Context, in *DispatcherServiceStartRequest, out *DispatcherServiceStartResponse) error
		Pause(ctx context.Context, in *DispatcherServicePauseRequest, out *DispatcherServicePauseResponse) error
		Resume(ctx context.Context, in *DispatcherServiceResumeRequest, out *DispatcherServiceResumeResponse) error
		Stop(ctx context.Context, in *DispatcherServiceStopRequest, out *DispatcherServiceStopResponse) error
	}
	type DispatcherService struct {
		dispatcherService
	}
	h := &dispatcherServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DispatcherService.Start",
		Path:    []string{"/task/start"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DispatcherService.Pause",
		Path:    []string{"/task/pause/{id}"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DispatcherService.Resume",
		Path:    []string{"/task/resume/{id}"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DispatcherService.Stop",
		Path:    []string{"/task/stop/{id}"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&DispatcherService{h}, opts...))
}

type dispatcherServiceHandler struct {
	DispatcherServiceHandler
}

func (h *dispatcherServiceHandler) Start(ctx context.Context, in *DispatcherServiceStartRequest, out *DispatcherServiceStartResponse) error {
	return h.DispatcherServiceHandler.Start(ctx, in, out)
}

func (h *dispatcherServiceHandler) Pause(ctx context.Context, in *DispatcherServicePauseRequest, out *DispatcherServicePauseResponse) error {
	return h.DispatcherServiceHandler.Pause(ctx, in, out)
}

func (h *dispatcherServiceHandler) Resume(ctx context.Context, in *DispatcherServiceResumeRequest, out *DispatcherServiceResumeResponse) error {
	return h.DispatcherServiceHandler.Resume(ctx, in, out)
}

func (h *dispatcherServiceHandler) Stop(ctx context.Context, in *DispatcherServiceStopRequest, out *DispatcherServiceStopResponse) error {
	return h.DispatcherServiceHandler.Stop(ctx, in, out)
}
