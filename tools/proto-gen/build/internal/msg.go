// This file is generated automatically, DO NOT modify it manually!!! 
package msg
import (
    "github.com/jeckbjy/fairy"
)

const (
    MsgId_HandlerDictReq = 1
    MsgId_HandlerDictRsp = 2
    
)

func init() {
    fairy.RegisterMessage(&HandlerDictReq{}, MsgId_HandlerDictReq)
    fairy.RegisterMessage(&HandlerDictRsp{}, MsgId_HandlerDictRsp)
    
}
