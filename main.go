package main

import (
	"ChaincodeInfo/sdkInit"
	"ChaincodeInfo/service"
	"ChaincodeInfo/web"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	cc_name    = "simplecc"
	cc_version = "1.0.0"
)

func main() {
	// init orgs information
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: "fixtures/channel-artifacts/Org1MSPanchors.tx",
		},
	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    "fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    "chaincode",
		ChaincodeVersion: cc_version,
	}

	// sdk setup
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and join
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")

	edu := service.ChaincodeInfo{
		Name:           "test",
		SignDate:         "test",
		Nation:         "test",
		CopyRightID:       "test",
		Creator:          "test",
		Holder:       "test",
		EnrollDate:     "test",
		NowState: "test",
		DoneDate:     "test",
		Major:          "test",
		QuaType:        "test",
		Length:         "test",
		Mode:           "test",
		Level:          "test",
		CopyRightLevel:     "test",
		CertNo:         "test",
		Photo:          "test.jpg",
	}

	serviceSetup, err := service.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk)
	if err != nil {
		fmt.Println()
		os.Exit(-1)
	}
	msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	result, err := serviceSetup.FindEduInfoByCopyRightID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.ChaincodeInfo
		json.Unmarshal(result, &edu)
		fmt.Println("根据证书编号查询信息成功：")
		fmt.Println(edu)
	}

	server := web.NewServer(serviceSetup)
	if err = server.Run(); err != nil {
		log.Fatalf("服务启动失败, 错误：%v", err)
	}
}
