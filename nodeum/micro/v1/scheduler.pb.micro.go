// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: nodeum/micro/v1/scheduler.proto

package microv1

import (
	fmt "fmt"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Api Endpoints for SchedulerService service

func NewSchedulerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "SchedulerService.ReadSchedule",
			Path:    []string{"/schedules/{task_id}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "SchedulerService.ListSchedules",
			Path:    []string{"/schedules"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "SchedulerService.WriteSchedule",
			Path:    []string{"/schedules"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "SchedulerService.DeleteSchedule",
			Path:    []string{"/schedules/{task_id}"},
			Method:  []string{"DELETE"},
			Handler: "rpc",
		},
	}
}

// Client API for SchedulerService service

type SchedulerService interface {
	// ReadSchedule returns information about a specific schedule, including date
	// of next occurrence.
	ReadSchedule(ctx context.Context, in *ReadScheduleRequest, opts ...client.CallOption) (*ReadScheduleResponse, error)
	// ListSchedules returns all actives schedules.
	ListSchedules(ctx context.Context, in *ListSchedulesRequest, opts ...client.CallOption) (*ListSchedulesResponse, error)
	// WriteSchedule creates or update a schedule for a task.
	// If one schedule already exists for the task, it will be overwritten.
	WriteSchedule(ctx context.Context, in *WriteScheduleRequest, opts ...client.CallOption) (*WriteScheduleResponse, error)
	// DeleteSchedule cancels and deletes a schedule for a task.
	DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...client.CallOption) (*DeleteScheduleResponse, error)
}

type schedulerService struct {
	c    client.Client
	name string
}

func NewSchedulerService(name string, c client.Client) SchedulerService {
	return &schedulerService{
		c:    c,
		name: name,
	}
}

func (c *schedulerService) ReadSchedule(ctx context.Context, in *ReadScheduleRequest, opts ...client.CallOption) (*ReadScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "SchedulerService.ReadSchedule", in)
	out := new(ReadScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerService) ListSchedules(ctx context.Context, in *ListSchedulesRequest, opts ...client.CallOption) (*ListSchedulesResponse, error) {
	req := c.c.NewRequest(c.name, "SchedulerService.ListSchedules", in)
	out := new(ListSchedulesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerService) WriteSchedule(ctx context.Context, in *WriteScheduleRequest, opts ...client.CallOption) (*WriteScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "SchedulerService.WriteSchedule", in)
	out := new(WriteScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerService) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...client.CallOption) (*DeleteScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "SchedulerService.DeleteSchedule", in)
	out := new(DeleteScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SchedulerService service

type SchedulerServiceHandler interface {
	// ReadSchedule returns information about a specific schedule, including date
	// of next occurrence.
	ReadSchedule(context.Context, *ReadScheduleRequest, *ReadScheduleResponse) error
	// ListSchedules returns all actives schedules.
	ListSchedules(context.Context, *ListSchedulesRequest, *ListSchedulesResponse) error
	// WriteSchedule creates or update a schedule for a task.
	// If one schedule already exists for the task, it will be overwritten.
	WriteSchedule(context.Context, *WriteScheduleRequest, *WriteScheduleResponse) error
	// DeleteSchedule cancels and deletes a schedule for a task.
	DeleteSchedule(context.Context, *DeleteScheduleRequest, *DeleteScheduleResponse) error
}

func RegisterSchedulerServiceHandler(s server.Server, hdlr SchedulerServiceHandler, opts ...server.HandlerOption) error {
	type schedulerService interface {
		ReadSchedule(ctx context.Context, in *ReadScheduleRequest, out *ReadScheduleResponse) error
		ListSchedules(ctx context.Context, in *ListSchedulesRequest, out *ListSchedulesResponse) error
		WriteSchedule(ctx context.Context, in *WriteScheduleRequest, out *WriteScheduleResponse) error
		DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, out *DeleteScheduleResponse) error
	}
	type SchedulerService struct {
		schedulerService
	}
	h := &schedulerServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SchedulerService.ReadSchedule",
		Path:    []string{"/schedules/{task_id}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SchedulerService.ListSchedules",
		Path:    []string{"/schedules"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SchedulerService.WriteSchedule",
		Path:    []string{"/schedules"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SchedulerService.DeleteSchedule",
		Path:    []string{"/schedules/{task_id}"},
		Method:  []string{"DELETE"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&SchedulerService{h}, opts...))
}

type schedulerServiceHandler struct {
	SchedulerServiceHandler
}

func (h *schedulerServiceHandler) ReadSchedule(ctx context.Context, in *ReadScheduleRequest, out *ReadScheduleResponse) error {
	return h.SchedulerServiceHandler.ReadSchedule(ctx, in, out)
}

func (h *schedulerServiceHandler) ListSchedules(ctx context.Context, in *ListSchedulesRequest, out *ListSchedulesResponse) error {
	return h.SchedulerServiceHandler.ListSchedules(ctx, in, out)
}

func (h *schedulerServiceHandler) WriteSchedule(ctx context.Context, in *WriteScheduleRequest, out *WriteScheduleResponse) error {
	return h.SchedulerServiceHandler.WriteSchedule(ctx, in, out)
}

func (h *schedulerServiceHandler) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, out *DeleteScheduleResponse) error {
	return h.SchedulerServiceHandler.DeleteSchedule(ctx, in, out)
}
