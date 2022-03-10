// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: protobuf/scheduler.micro.proto

package schedulerpb

import (
	fmt "fmt"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Api Endpoints for Scheduler service

func NewSchedulerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "Scheduler.ReadSchedule",
			Path:    []string{"/schedules/{task_id}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "Scheduler.ListSchedules",
			Path:    []string{"/schedules"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "Scheduler.WriteSchedule",
			Path:    []string{"/schedules"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "Scheduler.DeleteSchedule",
			Path:    []string{"/schedules/{task_id}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
	}
}

// Client API for Scheduler service

type SchedulerService interface {
	// ReadSchedule returns information about a specific schedule, including date
	// of next occurrence.
	ReadSchedule(ctx context.Context, in *ScheduleIdRequest, opts ...client.CallOption) (*ReadScheduleResponse, error)
	// ListSchedules returns all actives schedules.
	ListSchedules(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListSchedulesResponse, error)
	// WriteSchedule creates or update a schedule for a task.
	// If one schedule already exists for the task, it will be overwritten.
	WriteSchedule(ctx context.Context, in *WriteScheduleRequest, opts ...client.CallOption) (*ScheduleResponse, error)
	DeleteSchedule(ctx context.Context, in *ScheduleIdRequest, opts ...client.CallOption) (*ScheduleResponse, error)
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

func (c *schedulerService) ReadSchedule(ctx context.Context, in *ScheduleIdRequest, opts ...client.CallOption) (*ReadScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "Scheduler.ReadSchedule", in)
	out := new(ReadScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerService) ListSchedules(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListSchedulesResponse, error) {
	req := c.c.NewRequest(c.name, "Scheduler.ListSchedules", in)
	out := new(ListSchedulesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerService) WriteSchedule(ctx context.Context, in *WriteScheduleRequest, opts ...client.CallOption) (*ScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "Scheduler.WriteSchedule", in)
	out := new(ScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerService) DeleteSchedule(ctx context.Context, in *ScheduleIdRequest, opts ...client.CallOption) (*ScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "Scheduler.DeleteSchedule", in)
	out := new(ScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerHandler interface {
	// ReadSchedule returns information about a specific schedule, including date
	// of next occurrence.
	ReadSchedule(context.Context, *ScheduleIdRequest, *ReadScheduleResponse) error
	// ListSchedules returns all actives schedules.
	ListSchedules(context.Context, *emptypb.Empty, *ListSchedulesResponse) error
	// WriteSchedule creates or update a schedule for a task.
	// If one schedule already exists for the task, it will be overwritten.
	WriteSchedule(context.Context, *WriteScheduleRequest, *ScheduleResponse) error
	DeleteSchedule(context.Context, *ScheduleIdRequest, *ScheduleResponse) error
}

func RegisterSchedulerHandler(s server.Server, hdlr SchedulerHandler, opts ...server.HandlerOption) error {
	type scheduler interface {
		ReadSchedule(ctx context.Context, in *ScheduleIdRequest, out *ReadScheduleResponse) error
		ListSchedules(ctx context.Context, in *emptypb.Empty, out *ListSchedulesResponse) error
		WriteSchedule(ctx context.Context, in *WriteScheduleRequest, out *ScheduleResponse) error
		DeleteSchedule(ctx context.Context, in *ScheduleIdRequest, out *ScheduleResponse) error
	}
	type Scheduler struct {
		scheduler
	}
	h := &schedulerHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Scheduler.ReadSchedule",
		Path:    []string{"/schedules/{task_id}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Scheduler.ListSchedules",
		Path:    []string{"/schedules"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Scheduler.WriteSchedule",
		Path:    []string{"/schedules"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Scheduler.DeleteSchedule",
		Path:    []string{"/schedules/{task_id}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Scheduler{h}, opts...))
}

type schedulerHandler struct {
	SchedulerHandler
}

func (h *schedulerHandler) ReadSchedule(ctx context.Context, in *ScheduleIdRequest, out *ReadScheduleResponse) error {
	return h.SchedulerHandler.ReadSchedule(ctx, in, out)
}

func (h *schedulerHandler) ListSchedules(ctx context.Context, in *emptypb.Empty, out *ListSchedulesResponse) error {
	return h.SchedulerHandler.ListSchedules(ctx, in, out)
}

func (h *schedulerHandler) WriteSchedule(ctx context.Context, in *WriteScheduleRequest, out *ScheduleResponse) error {
	return h.SchedulerHandler.WriteSchedule(ctx, in, out)
}

func (h *schedulerHandler) DeleteSchedule(ctx context.Context, in *ScheduleIdRequest, out *ScheduleResponse) error {
	return h.SchedulerHandler.DeleteSchedule(ctx, in, out)
}
