package storage

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const bundledPath = "bundled"

type Store struct {
	Path string
}

type MiniProgram struct {
	AppId string
	Name  string
}

func NewStore(path string) *Store {
	return &Store{
		Path: path,
	}
}

func (s *Store) Save(appId, name, bundled string) (error, string) {
	fileName := fmt.Sprintf("%s/%s-%s.json", s.Path, name, appId)
	err := ioutil.WriteFile(fileName, []byte(bundled), 0644)
	if err != nil {
		return err, ""
	}

	return nil, appId
}

func (s *Store) Scan() (error, []MiniProgram) {
	programs := make([]MiniProgram, 0)
	files, err := ioutil.ReadDir(bundledPath)
	if err != nil {
		return err, programs
	}

	for _, f := range files {
		name := f.Name()
		index := strings.Index(name, "-")
		n := name[0:index]
		appId := name[index+1:len(name)-5]
		p := MiniProgram{
			AppId: n,
			Name:  appId,
		}
		programs = append(programs, p)
	}

	return nil, programs
}
