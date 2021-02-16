/**
 * @Author Kokutas
 * @Description 网卡信息
 * @Date 2021/2/17 0:08
 **/
package model

import "net"

//NIC: Network Interface Card 网卡
type NIC struct {
	// 网卡、网络适配器序号
	Index int `json:"Index,omitempty"`
	// 网卡/网络适配器名称
	Name string `json:"Name,omitempty"`
	// 网卡/网络适配器最大传输单元MTU
	MTU  int `json:"MTU,omitempty"`
	// 网卡/网络适配器MAC地址
	MAC string `json:"MAC,omitempty"`
	// 网卡/网络适配器Flags标记
	Flags string `json:"Flags,omitempty"`
	// 网卡/网络适配器相关的IPv4信息
	IPv4 IPInfo `json:"IPv4,omitempty"`
	// 网卡/网络适配器相关的IPv6信息
	IPv6 IPInfo
}
// IPInfo 网卡相关信息
type IPInfo struct {
	// IP 地址
	Addr net.IP `json:"Addr,omitempty"`
	// IP所在网络/掩码位192.168.1.0/24
	Net string `json:"Net,omitempty"`
	// 网卡/网络适配器掩码位/子网掩码
	SubMask int `json:"SubMask,omitempty"`
	//网卡/网络适配器IP掩码总位数
	MaskBits int `json:"MaskBits,omitempty"`
	//网卡/网络适配器IP的可用地址个数
	Available uint `json:"Available,omitempty"`
	// 网卡/网络适配器IP的掩码地址
	Mask net.IP `json:"Mask,omitempty"`
	// 网卡/网络适配器IP的第一个可用地址
	First net.IP `json:"First,omitempty"`
	// 网卡/网络适配器IP的最后一个可用地址
	Last net.IP `json:"Last,omitempty"`
	// 网卡/网络适配器加入的组播地址
	Multicast []net.IP `json:"Multicast,omitempty"`
}

func (nic *NIC) String() string {

}