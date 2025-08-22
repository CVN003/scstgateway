package scst

import (
	context "context"
	"encoding/json"

	"github.com/CVN003/scstgateway/core"
)

type Gateway struct {
	UnimplementedSCSTGatewayServer
}

func (g *Gateway) AddSVD(ctx context.Context, req *AddSVDReq) (*SCSTResp, error) {

	resp := &SCSTResp{
		Code: 0,
		Msg:  "success",
	}
	svd := &core.SVD{
		USN:               req.USN,
		VID:               req.VID,
		PID:               req.PID,
		LunName:           req.LunName,
		LunDeviceFullPath: req.LunDeviceFullPath,
		Handler:           req.HandlerType,
	}
	if err := svd.Add(); err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, err
	}
	return resp, nil
}

func (g *Gateway) RemoveSVD(ctx context.Context, req *RemoveSVDReq) (*SCSTResp, error) {
	resp := &SCSTResp{
		Code: 0,
		Msg:  "success",
	}
	svd := &core.SVD{
		LunName: req.LunName,
		Handler: req.HandlerType,
	}
	if err := svd.Remove(); err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, err
	}
	return resp, nil
}

func (g *Gateway) AddGroup(ctx context.Context, req *AddGroupReq) (*SCSTResp, error) {
	resp := &SCSTResp{
		Code: 0,
		Msg:  "success",
	}
	group := &core.Target{
		Name:       req.TargetName,
		TargetType: req.TargetType,
	}
	if err := group.AddGroup(req.GroupName); err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, err
	}
	return resp, nil
}

func (g *Gateway) AddLun2Group(ctx context.Context, req *AddLun2GroupReq) (*SCSTResp, error) {
	resp := &SCSTResp{
		Code: 0,
		Msg:  "success",
	}
	if err := core.AddLun2Grp(req.LunName, req.TargetName, req.TargetType, req.GroupName); err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, err
	}
	return resp, nil
}

func (g *Gateway) AddIni2Group(ctx context.Context, req *AddIni2GroupReq) (*SCSTResp, error) {
	resp := &SCSTResp{
		Code: 0,
		Msg:  "success",
	}
	tar := &core.Target{
		Name:       req.TargetName,
		TargetType: req.TargetType,
	}
	if err := tar.AddIni2Group(req.Ini, req.GroupName); err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, err
	}
	return resp, nil
}

func (g *Gateway) RemIni2Group(ctx context.Context, req *RemIni2GroupReq) (*SCSTResp, error) {
	resp := &SCSTResp{
		Code: 0,
		Msg:  "success",
	}
	tar := &core.Target{
		Name:       req.TargetName,
		TargetType: req.TargetType,
	}
	if err := tar.RemIni2Group(req.Ini, req.GroupName); err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, err
	}
	return resp, nil
}

func (g *Gateway) GetLiveConfig(ctx context.Context, req *GetLiveConfigReq) (*SCSTResp, error) {
	resp := &SCSTResp{
		Code: 0,
		Msg:  "success",
	}
	config, err := core.GetLiveConfig()
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp, err
	}
	if jsonStr, err := json.Marshal(config); err != nil {
		return resp, err
	} else {
		resp.Data = string(jsonStr)
		return resp, nil
	}

}
