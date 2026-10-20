package snowflake

import "github.com/bwmarrin/snowflake"

var node, _ = snowflake.NewNode(1)

func Generate() int64 {
	return node.Generate().Int64()
}
