package service

import (
  "io/ioutil"

  yaml "gopkg.in/yaml.v3"

  log "github.com/unicredit/abe/log"
)

type Redis struct {
  Name string `yaml:"name"`
  Addr string `yaml:"addr"`
  Password string `yaml:"password"`
}

type Storage struct {
  Redis map[string]Redis `yaml:"redis"`
}

type Active struct {
  Redis string `yaml:"redis"`
}

type Config struct {
  Storage Storage `yaml:"storage"`
  Active Active `yaml:"active"`
}

type Ext struct {
  Config Config `yaml:"config"`
}

func ReadConfig(filename string) Ext {
  log.Debug("reading config file")
  yamlFile, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Panic("no config file")
    panic(err)
  }
  var e Ext
  err = yaml.Unmarshal(yamlFile, &e)
  if err != nil {
    log.Panic("problem parsing config file")
    panic(err)
  }
  return e
}
