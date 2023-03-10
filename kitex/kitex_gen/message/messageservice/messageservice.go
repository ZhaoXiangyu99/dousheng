// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	"fmt"
	message "github.com/bytedance-youthcamp-jbzx/tiktok/kitex/kitex_gen/message"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return messageServiceServiceInfo
}

var messageServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MessageService"
	handlerType := (*message.MessageService)(nil)
	methods := map[string]kitex.MethodInfo{
		"MessageChat":   kitex.NewMethodInfo(messageChatHandler, newMessageChatArgs, newMessageChatResult, false),
		"MessageAction": kitex.NewMethodInfo(messageActionHandler, newMessageActionArgs, newMessageActionResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "message",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func messageChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(message.MessageChatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(message.MessageService).MessageChat(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MessageChatArgs:
		success, err := handler.(message.MessageService).MessageChat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MessageChatResult)
		realResult.Success = success
	}
	return nil
}
func newMessageChatArgs() interface{} {
	return &MessageChatArgs{}
}

func newMessageChatResult() interface{} {
	return &MessageChatResult{}
}

type MessageChatArgs struct {
	Req *message.MessageChatRequest
}

func (p *MessageChatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(message.MessageChatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MessageChatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MessageChatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MessageChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessageChatArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MessageChatArgs) Unmarshal(in []byte) error {
	msg := new(message.MessageChatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessageChatArgs_Req_DEFAULT *message.MessageChatRequest

func (p *MessageChatArgs) GetReq() *message.MessageChatRequest {
	if !p.IsSetReq() {
		return MessageChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessageChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessageChatResult struct {
	Success *message.MessageChatResponse
}

var MessageChatResult_Success_DEFAULT *message.MessageChatResponse

func (p *MessageChatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(message.MessageChatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MessageChatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MessageChatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MessageChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessageChatResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MessageChatResult) Unmarshal(in []byte) error {
	msg := new(message.MessageChatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessageChatResult) GetSuccess() *message.MessageChatResponse {
	if !p.IsSetSuccess() {
		return MessageChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessageChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.MessageChatResponse)
}

func (p *MessageChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func messageActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(message.MessageActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(message.MessageService).MessageAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MessageActionArgs:
		success, err := handler.(message.MessageService).MessageAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MessageActionResult)
		realResult.Success = success
	}
	return nil
}
func newMessageActionArgs() interface{} {
	return &MessageActionArgs{}
}

func newMessageActionResult() interface{} {
	return &MessageActionResult{}
}

type MessageActionArgs struct {
	Req *message.MessageActionRequest
}

func (p *MessageActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(message.MessageActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MessageActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MessageActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MessageActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MessageActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MessageActionArgs) Unmarshal(in []byte) error {
	msg := new(message.MessageActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessageActionArgs_Req_DEFAULT *message.MessageActionRequest

func (p *MessageActionArgs) GetReq() *message.MessageActionRequest {
	if !p.IsSetReq() {
		return MessageActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessageActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type MessageActionResult struct {
	Success *message.MessageActionResponse
}

var MessageActionResult_Success_DEFAULT *message.MessageActionResponse

func (p *MessageActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(message.MessageActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MessageActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MessageActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MessageActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MessageActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MessageActionResult) Unmarshal(in []byte) error {
	msg := new(message.MessageActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessageActionResult) GetSuccess() *message.MessageActionResponse {
	if !p.IsSetSuccess() {
		return MessageActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessageActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.MessageActionResponse)
}

func (p *MessageActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) MessageChat(ctx context.Context, Req *message.MessageChatRequest) (r *message.MessageChatResponse, err error) {
	var _args MessageChatArgs
	_args.Req = Req
	var _result MessageChatResult
	if err = p.c.Call(ctx, "MessageChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessageAction(ctx context.Context, Req *message.MessageActionRequest) (r *message.MessageActionResponse, err error) {
	var _args MessageActionArgs
	_args.Req = Req
	var _result MessageActionResult
	if err = p.c.Call(ctx, "MessageAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
