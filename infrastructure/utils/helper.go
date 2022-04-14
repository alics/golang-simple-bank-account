package util

import (
	"github.com/bwmarrin/snowflake"
	"reflect"
)

func IsNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func NewId() int64 {
	node, _ := snowflake.NewNode(0)
	id := node.Generate()
	return int64(id)
}
