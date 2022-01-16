package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

// Init
// startTime 从哪一年开始
// machineId 分布式系统的机器ID
func Init(startTime string, machineId int64) (err error) {
	start, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	snowflake.Epoch = start.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return
}

func GenerateId() int64 {
	return node.Generate().Int64()
}
