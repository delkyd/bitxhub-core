package contracts

import (
	"github.com/meshplus/bitxhub-core/contracts/boltvm"
)

type Store struct {
	boltvm.Stub
}

func (s *Store) Set(key string, value string) *boltvm.Response {
	s.SetObject(key, value)

	return boltvm.Success(nil)
}

func (s *Store) Get(key string) *boltvm.Response {
	var v string
	ok := s.Stub.GetObject(key, &v)
	if !ok {
		return boltvm.Error("there is not exist key")
	}

	return boltvm.Success([]byte(v))
}
