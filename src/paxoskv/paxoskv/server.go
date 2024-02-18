package paxoskv

import (
	"context"
	"errors"
	"sync"
)

var (
	NotEnoughQuorom  = errors.New("Not enough quorom")
	AcceptorBasePort = int64(3333)
)

type Version struct {
	mu       sync.Mutex
	acceptor Acceptor
}

type Versions map[int64]*Version

type KVServer struct {
	UnimplementedPaxosKVServer
	mu      sync.Mutex
	Storage map[string]Versions
}

func (s *KVServer) getLockedVersion(id *PaxosInstanceId) *Version {
	s.mu.Lock()
	defer s.mu.Unlock()

	versions, found := s.Storage[id.Key]

	if !found {
		versions = Versions{}
		s.Storage[id.Key] = versions
	}

	v, found := versions[id.Ver]

	if !found {
		v = &Version{
			acceptor: Acceptor{
				LastBal: &BalloutNum{},
				VBal:    &BalloutNum{},
				Val:     &Value{},
			},
		}
		versions[id.Ver] = v
	}

	v.mu.Lock()

	return v
}

func (s *KVServer) Prepare(ctx context.Context, r *Proposer) (*Acceptor, error) {
	v := s.getLockedVersion(r.Id)
}
