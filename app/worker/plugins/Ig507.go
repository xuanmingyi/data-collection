package plugins

import (
	"fmt"
)

type IG507 struct {
}

func (i *IG507) Name() string {
	return "ig507"
}

func (i *IG507) Get(uri string) (response string) {
	//BASEURL := "http://ig507.com"
	//LICENSE := Config.License
	//http.Get()
	return "sss"
}

func (i *IG507) API() []string {
	return []string{
		"gplist", // 股票列表

		"info",  // 公司简介
		"index", // 所属指数
		"gg",    // 历届高管成员
		"ds",    // 历届董事会成员
		"js",    // 历届监事会成员
		"fh",    // 近年分红
		"zf",    // 近年增发
		"jjxs",  // 解禁限售
		"fs",    // 财务摘要
		"pf",    // 近一年各季度利润
		"cf",    // 近一年各季度现金流
		"ep",    // 近年业绩预告
		"fi",    // 财务指标
	}
}

func (i *IG507) Call(api string) (response string) {
	switch api {
	case "gplist":
		fmt.Printf("gplist\n")
	}
	return "sssssssssss"
}
