// This file is generated automatically, DO NOT modify it manually!!! 
package msg
import (
    "github.com/name5566/leaf/network/protobuf"
    "github.com/golang/protobuf/proto"
)

var (
    Processor = protobuf.NewProcessor()
)

func register(msgid int32, msgobj proto.Message) {
    Processor.RegisterEx(uint16(msgid), msgobj)
}

func init() {
    register(1, &AuthReq{})
    register(2, &AuthRsp{})
    register(3, &LoginReq{})
    register(4, &LoginRsp{})
    
}
