package store

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"paris/annoyindex"
	"paris/config"
	"paris/logger"
	"paris/meta"
	"paris/model"
)

var (
	annoyIndex annoyindex.AnnoyIndexAngular
)

const (
	Feature = 100
)

func init() {
	meta, err := meta.NewMeta()
	if err != nil {
		log.Fatalf("Fatal error store : %v\n", err)
	}
	annoyIndex = annoyindex.NewAnnoyIndexAngular(Feature)
	if meta.Index != 0 {
		annoyIndex.Load(fmt.Sprintf("%s/%d", config.C.GetString("data.dir"), meta.Index))
	}
	for fid := range meta.Fids {
		fn := fmt.Sprintf("%s/%d", config.C.GetString("data.dir"), fid)
		f, err := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
		if err != nil {
			logger.L.Errorf("store read fn = %s, err = %v\n", fn, err)
			continue
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
					continue
			}
			item := &model.Item{}
			err = json.Unmarshal([]byte(line), item)
			if err != nil {
				continue
			}
			annoyIndex.AddItem(item.Id, item.Vector)
		}
		err = os.Remove(fn)

	}
}
