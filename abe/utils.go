package abe

import (
  "encoding/base32"
  "encoding/json"
)

func Init() {
}

func encode(x []byte) string {
	y := base32.StdEncoding.EncodeToString(x)
	return y
}

func decode(x string) []byte {
	y, _ := base32.StdEncoding.DecodeString(x)
	return y
}

func JsonObj2Str(x interface{}) string {
  x_, _ := json.Marshal(x)
  return string(x_)
}

func unique(s []string) []string {
  keys := make(map[string]bool)
  s_ := []string{}
  for _, tmp := range s {
    if _, value := keys[tmp]; !value {
      keys[tmp] = true
      s_ = append(s_, tmp)
    }
  }
  return s_
}
