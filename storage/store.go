package storage

import (
	"fmt"
	"io/ioutil"
)

type Store struct {
	Path string
}

func NewStore(path string) *Store {
	return &Store{
		Path: path,
	}
}

func (s *Store) Save(appId, name, bundled string) (error, string) {
	fileName := fmt.Sprintf("%s/%s-%s-.json", s.Path, name, appId)
	err := ioutil.WriteFile(fileName, []byte(bundled), 0644)
	if err != nil {
		return err, ""
	}

	return nil, appId
}
