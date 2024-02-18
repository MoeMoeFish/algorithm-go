package paxoskv

import (
	"errors"
)

var (
	NotEnoughQuorom  = errors.New("Not enough quorom")
	AcceptorBasePort = int64(3333)
)

type KVServer struct {
	UnimplementedPaxosKVServer
}
