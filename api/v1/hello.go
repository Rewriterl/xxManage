package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// v1 指迭代版本
// 对外提供的输入输出结构定义

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type HelloRes struct {
	g.Meta `mime:"application/json"`
}
