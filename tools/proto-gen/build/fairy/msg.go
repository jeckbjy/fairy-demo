// This file is generated automatically, DO NOT modify it manually!!! 
package msg
import (
    "github.com/jeckbjy/fairy"
)

const (
    MsgId_AuthReq = 1
    MsgId_AuthRsp = 2
    MsgId_LoginReq = 3
    MsgId_LoginRsp = 4
    
)

func init() {
    fairy.RegisterMessage(&AuthReq{}, MsgId_AuthReq)
    fairy.RegisterMessage(&AuthRsp{}, MsgId_AuthRsp)
    fairy.RegisterMessage(&LoginReq{}, MsgId_LoginReq)
    fairy.RegisterMessage(&LoginRsp{}, MsgId_LoginRsp)
    
}
