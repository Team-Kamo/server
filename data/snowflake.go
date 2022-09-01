package data

import (
	"OctaneServer/config"
	"errors"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/rs/zerolog/log"
)

var (
	SnowFlakeNode *snowflake.Node
	nodeNum       int
	ErrNodeFull   = errors.New("snowflake node is full")
)

func newSnowFlakeNode() error {
	//ToDo: Implement correct algorithm
	nodeNum = os.Getpid() % 1000
	done := false
	for i := nodeNum + 1; i < 1024; i++ {
		a := strconv.Itoa(i)
		_, err := db.Get(TableNodes, a)
		if err == ErrNoEntry {
			db.Insert(TableNodes, a, os.Getpid())
			done = true
			break
		} else if err != nil {
			return err
		}
	}
	if !done {
		return ErrNodeFull
	}

	node, err := snowflake.NewNode(int64(nodeNum))
	if err != nil {
		return err
	}
	log.Info().Int("node", nodeNum).Msg(config.Msg[config.CurrentConfig.Lang].Console.SnowFlakeNode)
	SnowFlakeNode = node
	return nil
}

func closeSnowFlakeNode() error {
	return db.Delete(TableNodes, strconv.Itoa(nodeNum))
}

func SnowFlake() int64 {
	return SnowFlakeNode.Generate().Int64()
}
