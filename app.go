package main

import "github.com/Tnze/CoolQ-Golang-SDK/cqp"

//go:generate cqcfg -c .
// cqp: 名称: Rondom-Pic
// cqp: 版本: 1.0.0:1
// cqp: 作者: Liumik
// cqp: 简介: 一个使用Golang编写的随机发送图片的CoolQ插件
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "top.atori.ronpic" // TODO: 修改为这个插件的ID
	cqp.PrivateMsg = onPrivateMsg
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return 0
}
