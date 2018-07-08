package persist

import (
	"log"

	"github.com/gohah/go-learn/crawler/engine"
	"github.com/gohah/go-learn/crawler/persist"
	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("item:%s save error:%s", item, err)
	}

	return err
}
