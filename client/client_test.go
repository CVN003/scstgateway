package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/CVN003/scstgateway/core"
	"github.com/CVN003/scstgateway/scst"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_GetLiveConfig(t *testing.T) {
	conn, err := grpc.NewClient("localhost:55101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := scst.NewSCSTGatewayClient(conn)
	resp, err := client.GetLiveConfig(context.Background(), &scst.GetLiveConfigReq{})
	if err != nil {
		panic(err)
	}
	config := core.SCSTMapping{}
	json.Unmarshal([]byte(resp.Data), &config)
	fmt.Printf("getLiveConfig: %v\n", config)
	os.WriteFile("/root/grpc-scst.json", []byte(resp.Data), 0644)
}

func TestSVD_ADD(t *testing.T) {
	conn, err := grpc.NewClient("localhost:55101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := scst.NewSCSTGatewayClient(conn)
	resp, err := client.AddSVD(context.Background(), &scst.AddSVDReq{
		LunName:           "lun_grpc_001",
		LunDeviceFullPath: "/dev/prod/v-yshtk2u2",
		VID:               "AAA",
		PID:               "BBB",
		USN:               "AABBCCDD",
		HandlerType:       "vdisk_blockio",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("add: %v\n", resp)
}

func Test_AddGroup(t *testing.T) {
	conn, err := grpc.NewClient("localhost:55101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := scst.NewSCSTGatewayClient(conn)
	resp, err := client.AddGroup(context.Background(), &scst.AddGroupReq{
		GroupName:  "group_grpc_001",
		TargetName: "cdm.iqn.2023-03.com.oe2203-12-69:tgt1",
		TargetType: "iscsi",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("addGroup: %v\n", resp)
}

func Test_Add2LunGroup(t *testing.T) {
	conn, err := grpc.NewClient("localhost:55101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := scst.NewSCSTGatewayClient(conn)
	resp, err := client.AddLun2Group(context.Background(), &scst.AddLun2GroupReq{
		LunName:    "lun_grpc_001",
		GroupName:  "group_grpc_001",
		TargetName: "cdm.iqn.2023-03.com.oe2203-12-69:tgt1",
		TargetType: "iscsi",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("addLun2Group: %v\n", resp)
}

func Test_AddIni2Group(t *testing.T) {
	conn, err := grpc.NewClient("localhost:55101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := scst.NewSCSTGatewayClient(conn)
	resp, err := client.AddIni2Group(context.Background(), &scst.AddIni2GroupReq{
		Ini:        "client1.oe2203-12-69:tgt1",
		GroupName:  "group_grpc_001",
		TargetName: "cdm.iqn.2023-03.com.oe2203-12-69:tgt1",
		TargetType: "iscsi",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("addIni2Group: %v\n", resp)
}

func Test_RemIni2Group(t *testing.T) {
	conn, err := grpc.NewClient("localhost:55101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := scst.NewSCSTGatewayClient(conn)
	resp, err := client.RemIni2Group(context.Background(), &scst.RemIni2GroupReq{
		Ini:        "client1.oe2203-12-69:tgt1",
		GroupName:  "group_grpc_001",
		TargetName: "cdm.iqn.2023-03.com.oe2203-12-69:tgt1",
		TargetType: "iscsi",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("remIni2Group: %v\n", resp)
}

func TestSVD_REMOVE(t *testing.T) {
	conn, err := grpc.NewClient("localhost:55101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := scst.NewSCSTGatewayClient(conn)
	resp, err := client.RemoveSVD(context.Background(), &scst.RemoveSVDReq{
		LunName:     "lun_grpc_001",
		HandlerType: "vdisk_blockio",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("remove: %v\n", resp)
}

func Test_SaveConfig(t *testing.T) {
	conn, err := grpc.NewClient("localhost:55101", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := scst.NewSCSTGatewayClient(conn)
	resp, err := client.SaveConfig(context.Background(), &scst.SaveConfigReq{
		Version: "r000001",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("saveConfig: %v\n", resp)
}
