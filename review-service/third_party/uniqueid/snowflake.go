package uniqueid

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func InitID(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return err
}

func GetID() int64 {
	return node.Generate().Int64()
}
