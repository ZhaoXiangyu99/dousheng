package utils

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"time"
)

// GetOnlyId 雪花算法生成唯一Id
func GetOnlyId() int64 {
	node, _ := snowflake.NewNode(1)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int64(node.Generate())/1e10 + rnd.Int63n(10000000)
}
