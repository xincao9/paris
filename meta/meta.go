package meta

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"paris/config"
	"path/filepath"
)

type Meta struct {
	Fids  []int64 `json:"fids"`
	Index int64   `json:"index"`
}

const (
	metaFn     = "meta.json"
	DefaultDir = "/tmp/paris"
)

func NewMeta() (*Meta, error) {
	dir := config.C.GetString("data.dir")
	if dir == "" {
		dir = DefaultDir
	}
	fn := filepath.Join(dir, metaFn)
	ok, err := isExist(fn)
	if err != nil {
		return nil, err
	}
	if ok {
		b, err := ioutil.ReadFile(fn)
		if err != nil {
			return nil, err
		}
		m := &Meta{}
		err = json.Unmarshal(b, m)
		if err != nil {
			return nil, err
		}
		return m, nil
	}
	m := &Meta{}
	m.Save()
	return m, nil
}

func (m *Meta) Save() error {
	dir := config.C.GetString("data.dir")
	fn := filepath.Join(dir)
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	os.Mkdir(dir, 0777)
	return ioutil.WriteFile(fn, b, 0777)
}

func isExist(fn string) (bool, error) {
	_, err := os.Stat(fn)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
