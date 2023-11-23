// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

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

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	UploadVideo(ctx context.Context, opts ...grpc.CallOption) (VideoService_UploadVideoClient, error)
	DownloadVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (VideoService_DownloadVideoClient, error)
	ForeignVideoExists(ctx context.Context, in *ForeignVideoCheck, opts ...grpc.CallOption) (*VideoExistenceResponse, error)
	GetVideoList(ctx context.Context, in *VideoQueryConfig, opts ...grpc.CallOption) (*VideoList, error)
	GetVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (*VideoMetadata, error)
	RateVideo(ctx context.Context, in *VideoRating, opts ...grpc.CallOption) (*Nothing, error)
	ViewVideo(ctx context.Context, in *VideoViewing, opts ...grpc.CallOption) (*Nothing, error)
	MakeComment(ctx context.Context, in *VideoComment, opts ...grpc.CallOption) (*Nothing, error)
	MakeCommentUpvote(ctx context.Context, in *CommentUpvote, opts ...grpc.CallOption) (*Nothing, error)
	GetCommentsForVideo(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
	GetVideoRecommendations(ctx context.Context, in *RecReq, opts ...grpc.CallOption) (*RecResp, error)
	ApproveVideo(ctx context.Context, in *VideoApproval, opts ...grpc.CallOption) (*Nothing, error)
	DeleteVideo(ctx context.Context, in *VideoDeletionReq, opts ...grpc.CallOption) (*Nothing, error)
	DeleteComment(ctx context.Context, in *CommentDeletionReq, opts ...grpc.CallOption) (*Nothing, error)
	GetFollowFeed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*VideoList, error)
	GetDanmaku(ctx context.Context, in *DanmakuQueryReq, opts ...grpc.CallOption) (*DanmakuList, error)
	AddDanmaku(ctx context.Context, in *Danmaku, opts ...grpc.CallOption) (*Nothing, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) UploadVideo(ctx context.Context, opts ...grpc.CallOption) (VideoService_UploadVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoService_ServiceDesc.Streams[0], "/proto.VideoService/uploadVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoServiceUploadVideoClient{stream}
	return x, nil
}

type VideoService_UploadVideoClient interface {
	Send(*InputVideoChunk) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type videoServiceUploadVideoClient struct {
	grpc.ClientStream
}

func (x *videoServiceUploadVideoClient) Send(m *InputVideoChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoServiceUploadVideoClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoServiceClient) DownloadVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (VideoService_DownloadVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoService_ServiceDesc.Streams[1], "/proto.VideoService/downloadVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoServiceDownloadVideoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VideoService_DownloadVideoClient interface {
	Recv() (*ResponseVideoChunk, error)
	grpc.ClientStream
}

type videoServiceDownloadVideoClient struct {
	grpc.ClientStream
}

func (x *videoServiceDownloadVideoClient) Recv() (*ResponseVideoChunk, error) {
	m := new(ResponseVideoChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoServiceClient) ForeignVideoExists(ctx context.Context, in *ForeignVideoCheck, opts ...grpc.CallOption) (*VideoExistenceResponse, error) {
	out := new(VideoExistenceResponse)
	err := c.cc.Invoke(ctx, "/proto.VideoService/foreignVideoExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetVideoList(ctx context.Context, in *VideoQueryConfig, opts ...grpc.CallOption) (*VideoList, error) {
	out := new(VideoList)
	err := c.cc.Invoke(ctx, "/proto.VideoService/getVideoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (*VideoMetadata, error) {
	out := new(VideoMetadata)
	err := c.cc.Invoke(ctx, "/proto.VideoService/getVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) RateVideo(ctx context.Context, in *VideoRating, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.VideoService/rateVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) ViewVideo(ctx context.Context, in *VideoViewing, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.VideoService/viewVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) MakeComment(ctx context.Context, in *VideoComment, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.VideoService/MakeComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) MakeCommentUpvote(ctx context.Context, in *CommentUpvote, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.VideoService/MakeCommentUpvote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetCommentsForVideo(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	out := new(CommentListResponse)
	err := c.cc.Invoke(ctx, "/proto.VideoService/GetCommentsForVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetVideoRecommendations(ctx context.Context, in *RecReq, opts ...grpc.CallOption) (*RecResp, error) {
	out := new(RecResp)
	err := c.cc.Invoke(ctx, "/proto.VideoService/GetVideoRecommendations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) ApproveVideo(ctx context.Context, in *VideoApproval, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.VideoService/ApproveVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) DeleteVideo(ctx context.Context, in *VideoDeletionReq, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.VideoService/DeleteVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) DeleteComment(ctx context.Context, in *CommentDeletionReq, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.VideoService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetFollowFeed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*VideoList, error) {
	out := new(VideoList)
	err := c.cc.Invoke(ctx, "/proto.VideoService/GetFollowFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetDanmaku(ctx context.Context, in *DanmakuQueryReq, opts ...grpc.CallOption) (*DanmakuList, error) {
	out := new(DanmakuList)
	err := c.cc.Invoke(ctx, "/proto.VideoService/GetDanmaku", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) AddDanmaku(ctx context.Context, in *Danmaku, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.VideoService/addDanmaku", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	UploadVideo(VideoService_UploadVideoServer) error
	DownloadVideo(*VideoRequest, VideoService_DownloadVideoServer) error
	ForeignVideoExists(context.Context, *ForeignVideoCheck) (*VideoExistenceResponse, error)
	GetVideoList(context.Context, *VideoQueryConfig) (*VideoList, error)
	GetVideo(context.Context, *VideoRequest) (*VideoMetadata, error)
	RateVideo(context.Context, *VideoRating) (*Nothing, error)
	ViewVideo(context.Context, *VideoViewing) (*Nothing, error)
	MakeComment(context.Context, *VideoComment) (*Nothing, error)
	MakeCommentUpvote(context.Context, *CommentUpvote) (*Nothing, error)
	GetCommentsForVideo(context.Context, *CommentRequest) (*CommentListResponse, error)
	GetVideoRecommendations(context.Context, *RecReq) (*RecResp, error)
	ApproveVideo(context.Context, *VideoApproval) (*Nothing, error)
	DeleteVideo(context.Context, *VideoDeletionReq) (*Nothing, error)
	DeleteComment(context.Context, *CommentDeletionReq) (*Nothing, error)
	GetFollowFeed(context.Context, *FeedReq) (*VideoList, error)
	GetDanmaku(context.Context, *DanmakuQueryReq) (*DanmakuList, error)
	AddDanmaku(context.Context, *Danmaku) (*Nothing, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) UploadVideo(VideoService_UploadVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedVideoServiceServer) DownloadVideo(*VideoRequest, VideoService_DownloadVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadVideo not implemented")
}
func (UnimplementedVideoServiceServer) ForeignVideoExists(context.Context, *ForeignVideoCheck) (*VideoExistenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForeignVideoExists not implemented")
}
func (UnimplementedVideoServiceServer) GetVideoList(context.Context, *VideoQueryConfig) (*VideoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoList not implemented")
}
func (UnimplementedVideoServiceServer) GetVideo(context.Context, *VideoRequest) (*VideoMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedVideoServiceServer) RateVideo(context.Context, *VideoRating) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateVideo not implemented")
}
func (UnimplementedVideoServiceServer) ViewVideo(context.Context, *VideoViewing) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewVideo not implemented")
}
func (UnimplementedVideoServiceServer) MakeComment(context.Context, *VideoComment) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeComment not implemented")
}
func (UnimplementedVideoServiceServer) MakeCommentUpvote(context.Context, *CommentUpvote) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeCommentUpvote not implemented")
}
func (UnimplementedVideoServiceServer) GetCommentsForVideo(context.Context, *CommentRequest) (*CommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsForVideo not implemented")
}
func (UnimplementedVideoServiceServer) GetVideoRecommendations(context.Context, *RecReq) (*RecResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoRecommendations not implemented")
}
func (UnimplementedVideoServiceServer) ApproveVideo(context.Context, *VideoApproval) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveVideo not implemented")
}
func (UnimplementedVideoServiceServer) DeleteVideo(context.Context, *VideoDeletionReq) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVideo not implemented")
}
func (UnimplementedVideoServiceServer) DeleteComment(context.Context, *CommentDeletionReq) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedVideoServiceServer) GetFollowFeed(context.Context, *FeedReq) (*VideoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowFeed not implemented")
}
func (UnimplementedVideoServiceServer) GetDanmaku(context.Context, *DanmakuQueryReq) (*DanmakuList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDanmaku not implemented")
}
func (UnimplementedVideoServiceServer) AddDanmaku(context.Context, *Danmaku) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDanmaku not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_UploadVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoServiceServer).UploadVideo(&videoServiceUploadVideoServer{stream})
}

type VideoService_UploadVideoServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*InputVideoChunk, error)
	grpc.ServerStream
}

type videoServiceUploadVideoServer struct {
	grpc.ServerStream
}

func (x *videoServiceUploadVideoServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoServiceUploadVideoServer) Recv() (*InputVideoChunk, error) {
	m := new(InputVideoChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VideoService_DownloadVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VideoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoServiceServer).DownloadVideo(m, &videoServiceDownloadVideoServer{stream})
}

type VideoService_DownloadVideoServer interface {
	Send(*ResponseVideoChunk) error
	grpc.ServerStream
}

type videoServiceDownloadVideoServer struct {
	grpc.ServerStream
}

func (x *videoServiceDownloadVideoServer) Send(m *ResponseVideoChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _VideoService_ForeignVideoExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForeignVideoCheck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).ForeignVideoExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/foreignVideoExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).ForeignVideoExists(ctx, req.(*ForeignVideoCheck))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoQueryConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/getVideoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideoList(ctx, req.(*VideoQueryConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/getVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideo(ctx, req.(*VideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_RateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoRating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).RateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/rateVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).RateVideo(ctx, req.(*VideoRating))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_ViewVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoViewing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).ViewVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/viewVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).ViewVideo(ctx, req.(*VideoViewing))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_MakeComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).MakeComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/MakeComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).MakeComment(ctx, req.(*VideoComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_MakeCommentUpvote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentUpvote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).MakeCommentUpvote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/MakeCommentUpvote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).MakeCommentUpvote(ctx, req.(*CommentUpvote))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetCommentsForVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetCommentsForVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/GetCommentsForVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetCommentsForVideo(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetVideoRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideoRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/GetVideoRecommendations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideoRecommendations(ctx, req.(*RecReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_ApproveVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoApproval)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).ApproveVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/ApproveVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).ApproveVideo(ctx, req.(*VideoApproval))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_DeleteVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoDeletionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).DeleteVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/DeleteVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).DeleteVideo(ctx, req.(*VideoDeletionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentDeletionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).DeleteComment(ctx, req.(*CommentDeletionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetFollowFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetFollowFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/GetFollowFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetFollowFeed(ctx, req.(*FeedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetDanmaku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DanmakuQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetDanmaku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/GetDanmaku",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetDanmaku(ctx, req.(*DanmakuQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_AddDanmaku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Danmaku)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).AddDanmaku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoService/addDanmaku",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).AddDanmaku(ctx, req.(*Danmaku))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "foreignVideoExists",
			Handler:    _VideoService_ForeignVideoExists_Handler,
		},
		{
			MethodName: "getVideoList",
			Handler:    _VideoService_GetVideoList_Handler,
		},
		{
			MethodName: "getVideo",
			Handler:    _VideoService_GetVideo_Handler,
		},
		{
			MethodName: "rateVideo",
			Handler:    _VideoService_RateVideo_Handler,
		},
		{
			MethodName: "viewVideo",
			Handler:    _VideoService_ViewVideo_Handler,
		},
		{
			MethodName: "MakeComment",
			Handler:    _VideoService_MakeComment_Handler,
		},
		{
			MethodName: "MakeCommentUpvote",
			Handler:    _VideoService_MakeCommentUpvote_Handler,
		},
		{
			MethodName: "GetCommentsForVideo",
			Handler:    _VideoService_GetCommentsForVideo_Handler,
		},
		{
			MethodName: "GetVideoRecommendations",
			Handler:    _VideoService_GetVideoRecommendations_Handler,
		},
		{
			MethodName: "ApproveVideo",
			Handler:    _VideoService_ApproveVideo_Handler,
		},
		{
			MethodName: "DeleteVideo",
			Handler:    _VideoService_DeleteVideo_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _VideoService_DeleteComment_Handler,
		},
		{
			MethodName: "GetFollowFeed",
			Handler:    _VideoService_GetFollowFeed_Handler,
		},
		{
			MethodName: "GetDanmaku",
			Handler:    _VideoService_GetDanmaku_Handler,
		},
		{
			MethodName: "addDanmaku",
			Handler:    _VideoService_AddDanmaku_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "uploadVideo",
			Handler:       _VideoService_UploadVideo_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "downloadVideo",
			Handler:       _VideoService_DownloadVideo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "videoservice.proto",
}
