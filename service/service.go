package service

import (
	"fmt"
	"strings"

	"github.com/marcellop71/mosaic/abe"
	//"github.com/marcellop71/mosaic/abe/log"
)

type Storage interface {
	Init(config Config)
	Set(key string, value string)
	Get(key string) string
	Scan(key string) []string
}

var storage Storage

func InitAbeService(config Config, abeStorage Storage) {
	if abeStorage != nil {
		storage = abeStorage
	} else {
		switch config.Active.Type {
		case "Redis":
			storage = new(StorageRedis)
		case "Leveldb":
			storage = new(StorageLeveldb)
		default:
			storage = new(StorageLeveldb)
		}
	}
	storage.Init(config)
}

// generate new organization and stores it
func SetupOrg(orgStr string, lib string, curveStr string, seed string) string {
	curveJson := abe.Encode(fmt.Sprintf("{\"name\": \"%s\", \"seed\": \"%s\"}", curveStr, seed))
	if (lib == "pbc") {
		curveJson = FetchCurve(curveStr)
	}
	orgJson := abe.NewRandomOrgJson(curveJson)
	StoreOrg(orgStr, orgJson)
	return orgJson
}

// generate new authority (private and public params) and stores it
func SetupAuth(authStr string, orgStr string) string {
	orgJson := FetchOrg(orgStr)
	authkeysJson := abe.NewRandomAuthJson(orgJson)
	authkeys := abe.NewAuthKeysOfJsonStr(authkeysJson)
	StoreAuthPrv(authStr, authkeys.AuthPrv_)
	StoreAuthPub(authStr, authkeys.AuthPub_)
	return authkeysJson
}

// generate new user key and stores it
func SetupUserkey(user string, attr string) string {
	authStr := strings.Split(attr, "@")[1]
	authprvJson := FetchAuthPrv(authStr)
	userattrsJson := abe.NewRandomUserkeyJson(user, attr, authprvJson)
	userattrs := abe.NewUserAttrsOfJsonStr(userattrsJson)
	for attr, userkeyJson := range userattrs.Userkey_ {
		StoreUserkey(user, attr, userkeyJson)
	}
	return userattrsJson
}

// generate new random secret for an organization
func NewRandomSecret(orgStr string) string {
	orgJson := FetchOrg(orgStr)
	secretJson := abe.NewRandomSecretJson(orgJson)
	return secretJson
}

// store Curve json object in key pub:curve:<curveStr>
func StoreCurve(curveStr string, curveJson string) {
	storage.Set(fmt.Sprintf("pub:curve:%s", curveStr), curveJson)
}

// fetch Curve json object from key pub:curve:<curveStr>
func FetchCurve(curveStr string) string {
	return storage.Get(fmt.Sprintf("pub:curve:%s", curveStr))
}

// store Org json object in key pub:org:<orgStr>
func StoreOrg(orgStr string, orgJson string) {
	storage.Set(fmt.Sprintf("pub:org:%s", orgStr), orgJson)
}

// fetch Org json object from key pub:org:<orgStr>
func FetchOrg(orgStr string) string {
	return storage.Get(fmt.Sprintf("pub:org:%s", orgStr))
}

// store AuthPub json object in key pub:auth:<authStr>
func StoreAuthPub(authStr string, authpub_json string) {
	storage.Set(fmt.Sprintf("pub:auth:%s", authStr), authpub_json)
}

// fetch AuthPub json object from key pub:auth:<authStr>
func FetchAuthPub(authStr string) string {
	return storage.Get(fmt.Sprintf("pub:auth:%s", authStr))
}

// store AuthPrv json object in key prv:auth:<authStr>
func StoreAuthPrv(authStr string, authprvJson string) {
	storage.Set(fmt.Sprintf("prv:auth:%s", authStr), authprvJson)
}

// fetch AuthPrv json object from key prv:auth:<authStr>
func FetchAuthPrv(authStr string) string {
	return storage.Get(fmt.Sprintf("prv:auth:%s", authStr))
}

// store Userkey json object in key prv:userkey:<user>:<attr>
func StoreUserkey(user string, attr string, userkey_json string) {
	storage.Set(fmt.Sprintf("prv:userkey:%s:%s", user, attr), userkey_json)
}

// fetch Userkey json object from key prv:userkey:<user>:<attr>
func FetchUserkey(user string, attr string) string {
	return storage.Get(fmt.Sprintf("prv:userkey:%s:%s", user, attr))
}

// fetch authorities public keys
func FetchAuthPubs(authpubsJson string) string {
	authpubs := abe.NewAuthPubsOfJsonStr(authpubsJson)
	for attr, _ := range authpubs.AuthPub_ {
		authpubs.AuthPub_[attr] = FetchAuthPub(attr)
	}
	return abe.JsonObjToEncStr(&abe.AuthPubs{
		AuthPub_: authpubs.AuthPub_})
}

// fetch UserAttrs json object scanning for keys of the form prv:userkey:<user>:*
func FetchUserAttrs(user string) string {
	coeff := make(map[string][]int)
	keys := storage.Scan(fmt.Sprintf("prv:userkey:%s:", user))
	for _, tmp := range keys {
		attr := strings.Split(tmp, ":")[3]
		coeff[attr] = []int{}
	}
	return abe.JsonObjToEncStr(&abe.UserAttrs{
		User: user, Coeff: coeff})
}

// fetch Userkeys json object from keys of the form prv:userkey:<user>:<attr>
func FetchUserkeys(userattrsJson string) string {
	userattrs := abe.NewUserAttrsOfJsonStr(userattrsJson)
	userkeys := make(map[string]string)
	for attr, _ := range userattrs.Coeff {
		userkeys[attr] = FetchUserkey(userattrs.User, attr)
	}
	return abe.JsonObjToEncStr(&abe.UserAttrs{
		User: userattrs.User, Coeff: userattrs.Coeff, Userkey_: userkeys})
}
