/**
 * @Author Kokutas
 * @Description IP信息
 * @Date 2021/2/17 0:09
 **/
package model

import (
	"fmt"
	"net"
	"strings"
)

// IPAddress ip address 信息
type IPAddress struct {
	// 查询结果状态
	Status string `json:"status,omitempty"`
	// 国家
	Country string `json:"country,omitempty"`
	// 国家代码
	CountryCode string `json:"countrycode,omitempty"`
	// 区域名称
	RegionName string `json:"regionName,omitempty"`
	// 区域代码
	Region string `json:"region,omitempty"`
	// 城市
	City string `json:"city,omitempty"`
	// 邮政编码
	Zip string `json:"zip,omitempty"`
	// 纬度
	Lat float64 `json:"lat,omitempty"`
	// 经度
	Lon float64 `json:"lon,omitempty"`
	// 时区
	Timezone string `json:"timezone,omitempty"`
	// 服务商
	Isp string `json:"isp,omitempty"`
	// 机构
	Org string `json:"org,omitempty"`
	// 运营商（AS Whois）
	As string `json:"as,omitempty"`
	// 查询IP
	IP net.IP `json:"query,omitempty"`
}
func (address *IPAddress)String()string{
	infoStr :=make([]string,0)
	if len(address.Country)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","国家",address.Country))
	}
	if len(address.CountryCode)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","国家代码",address.CountryCode))
	}
	if len(address.RegionName)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","区域",address.RegionName))
	}
	if len(address.Region)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","区域代码",address.Region))
	}
	if len(address.City)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","城市",address.City))
	}
	if len(address.Zip)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","邮政编码",address.Zip))
	}
	if address.Lat!=0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","纬度",address.Lat))
	}
	if  address.Lon!=0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","经度",address.Lon))
	}

	if len(address.Isp)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","服务商",address.Isp))
	}
	if len(address.Org)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","机构",address.Org))
	}
	if len(address.As)>0{
		infoStr = append(infoStr,fmt.Sprintf("%v : %v\r\n","运营商",address.As))
	}

	if len(address.IP)>0&& (len(address.Country)>0||len(address.Region)>0){
		infoStr = append(infoStr, fmt.Sprintf("%v : %v\r\n","公网IP",address.IP))
	}else{
		infoStr = append(infoStr, fmt.Sprintf("%v : %v\r\n","内网IP",address.IP))
	}
	str:=""
	for _, info := range infoStr {
		if len(strings.TrimSpace(strings.ReplaceAll(info,"\r\n","")))>0{
			str+=info
		}
	}
	return str
}