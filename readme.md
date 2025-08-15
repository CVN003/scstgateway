#SCST Gateway

基于GRPC实现SCST内核层的所有操作

已经实现的：
- vdisk_blockio、vdisk_fileio
- lun group
- ini
- target(iscsi/qla2x00t)


服务端示例：

```go
import (
	"fmt"
	"net"

	"github.com/CVN003/scstgateway/scst"
	"google.golang.org/grpc"
)

func main() {
	var c chan int
	go func() {
		lis, err := net.Listen("tcp", ":55101")
		if err != nil {
			panic(err)
		}
		ser := grpc.NewServer()
		scst.RegisterSCSTGatewayServer(ser, &scst.Gateway{})
		if err := ser.Serve(lis); err != nil {
			panic(err)
		}
		c <- 1
	}()

	fmt.Printf("scst gateway server start at :55101")
	<-c
}
```

客户端示例:

```go
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
```
更多用例可以参考client/client_test.go
















商用合作联系:cvn009

