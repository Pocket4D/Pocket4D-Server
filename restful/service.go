package restful

import (
	"github.com/FireStack-Lab/pocket4d-server/storage"
	"github.com/google/uuid"
)

type Service struct {
	Store *storage.Store
}

func NewService(store *storage.Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) Save(name, bundled string) (error, string) {
	appId := uuid.New().String()
	return s.Store.Save(appId, name, bundled)
}

func (s *Service) Scan() (error, []storage.MiniProgram) {
	return s.Store.Scan()
}
