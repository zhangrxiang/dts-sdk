package main

import (
	"fmt"
	"github.com/Atian-OE/DTSSDK_Golang/dtssdk"
	"github.com/Atian-OE/DTSSDK_Golang/dtssdk/model"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Println("start")

	client:= dtssdk.NewDTSClient("127.0.0.1")

	client.CallConnected(func(addr string) {
		fmt.Println("连接成功:%s!",addr)
	})
	client.CallDisconnected(func(addr string) {
		fmt.Println("断开连接:%s!",addr)
	})
	time.Sleep(time.Second*2)

	rep1,err:= client.GetDefenceZone(1,"")
	fmt.Println(err)
	fmt.Println(rep1)


	rep,err:= client.CallZoneTempNotify(func(notify *model.ZoneTempNotify, e error) {
		fmt.Println("CallZoneTempNotify"+notify.DeviceID)
	})
	fmt.Println(err)
	fmt.Println(rep)


	rep,err= client.CallTempSignalNotify(func(notify *model.TempSignalNotify, e error) {
		fmt.Println("CallTempSignalNotify"+notify.DeviceID)
	})
	fmt.Println(err)
	fmt.Println(rep)

	rep,err= client.CallDeviceEventNotify(func(notify *model.DeviceEventNotify, e error) {
		fmt.Println("CallDeviceEventNotify"+notify.DeviceID)
	})
	fmt.Println(err)
	fmt.Println(rep)

	rep,err= client.CallZoneAlarmNotify(func(notify *model.ZoneAlarmNotify, e error) {
		fmt.Println("CallZoneAlarmNotify"+notify.DeviceID)
	})
	fmt.Println(err)
	fmt.Println(rep)



	<-ch
	fmt.Println("quit")
}