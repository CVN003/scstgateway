package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/CVN003/scstgateway/scst"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
