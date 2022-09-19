// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: notes/notes.proto

package notes

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotesClient is the client API for Notes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotesClient interface {
	// Saving a single note.
	Save(ctx context.Context, in *Note, opts ...grpc.CallOption) (*NoteSaveReply, error)
	// Retrieving a single note.
	Load(ctx context.Context, in *NoteSearch, opts ...grpc.CallOption) (*Note, error)
}

type notesClient struct {
	cc grpc.ClientConnInterface
}

func NewNotesClient(cc grpc.ClientConnInterface) NotesClient {
	return &notesClient{cc}
}

func (c *notesClient) Save(ctx context.Context, in *Note, opts ...grpc.CallOption) (*NoteSaveReply, error) {
	out := new(NoteSaveReply)
	err := c.cc.Invoke(ctx, "/notes.Notes/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) Load(ctx context.Context, in *NoteSearch, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotesServer is the server API for Notes service.
// All implementations must embed UnimplementedNotesServer
// for forward compatibility
type NotesServer interface {
	// Saving a single note.
	Save(context.Context, *Note) (*NoteSaveReply, error)
	// Retrieving a single note.
	Load(context.Context, *NoteSearch) (*Note, error)
	mustEmbedUnimplementedNotesServer()
}

// UnimplementedNotesServer must be embedded to have forward compatible implementations.
type UnimplementedNotesServer struct {
}

func (UnimplementedNotesServer) Save(context.Context, *Note) (*NoteSaveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedNotesServer) Load(context.Context, *NoteSearch) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedNotesServer) mustEmbedUnimplementedNotesServer() {}

// UnsafeNotesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotesServer will
// result in compilation errors.
type UnsafeNotesServer interface {
	mustEmbedUnimplementedNotesServer()
}

func RegisterNotesServer(s grpc.ServiceRegistrar, srv NotesServer) {
	s.RegisterService(&Notes_ServiceDesc, srv)
}

func _Notes_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).Save(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteSearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).Load(ctx, req.(*NoteSearch))
	}
	return interceptor(ctx, in, info, handler)
}

// Notes_ServiceDesc is the grpc.ServiceDesc for Notes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notes.Notes",
	HandlerType: (*NotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Notes_Save_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _Notes_Load_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notes/notes.proto",
}
