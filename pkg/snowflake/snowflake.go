package snowflake

import (
	"time"

	"github.com/SoonDubu923/go-forum/config"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// Init initializes the snowflake node. Snowflake is a distributed algorithm for generating unique IDs, but we only use one node here for simplicity.
//
// Params:
// - startTime: the start time of the snowflake node
// - machineID: the ID of the machine
// Returns:
// - error: if any error occurs
func Init(cfg *config.SnowflakeConfig) (err error) {
    var st time.Time
    st, err = time.Parse("2006-01-02", cfg.StartTime)
    if err != nil {
        return
    }
    snowflake.Epoch = st.UnixNano() / 1000000
    node, err = snowflake.NewNode(cfg.MachineID)
    return
}

// GenID generates a unique ID.
//
// Returns:
// - int64: the generated ID
func GenID() int64 {
    return node.Generate().Int64()
}