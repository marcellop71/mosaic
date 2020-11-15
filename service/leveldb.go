package service

import (
	"strings"

	levigo "github.com/jmhodges/levigo"

	"gitthub.com/marcellop71/mosaic/abe/log"
)

type StorageLeveldb struct{}

var db *levigo.DB

// init
func (abeStorage StorageLeveldb) Init(config Config) {
  var err error
	conf := config.Storage.Leveldb[config.Active.Label]
  opts := levigo.NewOptions()
  opts.SetCache(levigo.NewLRUCache(3<<30))
  opts.SetCreateIfMissing(true)
	err = levigo.DestroyDatabase(conf.Name, opts)
	if err != nil {
    log.Info("%s", err)
  }
  db, err = levigo.Open(conf.Name, opts)
  if err != nil {
    log.Error("%s", err)
  }
}

// set
func (abeStorage StorageLeveldb) Set(key string, value string) {
	var err error
	wo := levigo.NewWriteOptions()
	err = db.Put(wo, []byte(key), []byte(value))
	if err != nil {
		log.Error("%s", err)
	}
	wo.Close()
}

// get
func (abeStorage StorageLeveldb) Get(key string) string {
	var err error
	ro := levigo.NewReadOptions()
	value, err := db.Get(ro, []byte(key))
	if err != nil {
		log.Error("%s", err)
	}
	ro.Close()
  return string(value)
}

// scan
func (abeStorage StorageLeveldb) Scan(key string) []string {
	var err error
	var keys []string
	ro := levigo.NewReadOptions()
	ro.SetFillCache(false)
	it := db.NewIterator(ro)
	defer it.Close()
	for it.Seek([]byte(key)); it.Valid(); it.Next() {
		key_ := string(it.Key())
		if strings.HasPrefix(key_, key) {
				keys = append(keys, key_)
		} else {
			break
		}
	}
	if err = it.GetError(); err != nil {
		log.Error("%s", err)
	}
	return keys
}
