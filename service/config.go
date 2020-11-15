package service

import (
  "io/ioutil"

  yaml "gopkg.in/yaml.v3"

  "github.com/marcellop71/mosaic/abe/log"
)

type ArithmeticConf struct {
  Curve string `yaml:"curve"`
  Library string `yaml:"library"`
}

type RedisConf struct {
  Addr string `yaml:"addr"`
  Password string `yaml:"password"`
}

type LeveldbConf struct {
  Name string `yaml:"name"`
}

type StorageConf struct {
  Redis map[string]RedisConf `yaml:"redis"`
  Leveldb map[string]LeveldbConf `yaml:"leveldb"`
}

type ActiveConf struct {
  Type string `yaml:"type"`
  Label string `yaml:"label"`
}

type Config struct {
  Arithmetic ArithmeticConf `yaml:"arithmetic"`
  Storage StorageConf `yaml:"storage"`
  Active ActiveConf `yaml:"active"`
}

type Ext struct {
  Config Config `yaml:"config"`
}

func ReadConfig(filename string) Ext {
  yamlFile, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Panic("no config file")
  }
  var e Ext
  err = yaml.Unmarshal(yamlFile, &e)
  if err != nil {
    log.Panic("problem parsing config file")
  }
  return e
}
