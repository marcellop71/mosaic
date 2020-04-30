package service

import (
	"encoding/json"
	"fmt"
	"os"

	redis "github.com/go-redis/redis/v7"

	abe "github.com/unicredit/abe/abe"

	log "github.com/unicredit/abe/log"
)

type AbeServiceStorageRedis struct{}

var rc *redis.Client

// init storage
func (abeStorage AbeServiceStorageRedis) InitStorage(config Config) {
	log.Debug("init")

	redis := config.Storage.Redis[config.Active.Redis]
	initRedis(redis.Addr, redis.Password)
	FlushRedis()
}

func initRedis(addr string, password string) {
	log.Debug("init Redis client")
	rc = redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: 0})
	if rc == nil {
		log.Error("Redis client error")
		os.Exit(1)
	}
}

func GetClient() *redis.Client {
	return rc
}

func FlushRedis() {
	log.Debug("flush Redis")
	rc.FlushAll()
}

func SetStringDict(rkey string, s []string) {
	rc.HMSet(rkey, s)
}

func (abeStorage AbeServiceStorageRedis) StoreCurve(curveName string, curve_ *abe.AbeCurve_) {
	log.Debug("store curve")
	key := fmt.Sprintf("pub:curve:%s", curveName)
	SetStringDict(key, []string{
		"params", curve_.Params,
		"maxrndexp", curve_.Maxrndexp,
	})
}

func (abeStorage AbeServiceStorageRedis) StoreOrg(orgName string, org_ *abe.AbeOrg_) {
	log.Debug("store organization")
	key := fmt.Sprintf("pub:org:%s", orgName)
	SetStringDict(key, []string{
		"curve", org_.Curve,
		"g1", org_.G1,
		"g2", org_.G2,
		"e", org_.E,
	})
}

func (abeStorage AbeServiceStorageRedis) StoreAuthPub(authName string, authpub_ *abe.AbeAuthPub_) {
	log.Debug("store authority pub")
	key := fmt.Sprintf("pub:auth:%s", authName)
	SetStringDict(key, []string{
		"org", authpub_.Org,
		"ealpha", authpub_.Ealpha,
		"g1y", authpub_.G1y,
	})
}

func (abeStorage AbeServiceStorageRedis) StoreAuthPrv(authName string, authprv_ *abe.AbeAuthPrv_) {
	log.Debug("store authority prv")
	key := fmt.Sprintf("prv:auth:%s", authName)
	SetStringDict(key, []string{
		"org", authprv_.Org,
		"alpha", authprv_.Alpha,
		"y", authprv_.Y,
	})
}

func (abeStorage AbeServiceStorageRedis) StoreUserkeyPrv(user string, attr string, userkey_ *abe.AbeUserkey_) {
	log.Debug("store user key")
	key := fmt.Sprintf("prv:userkey:%s:%s", user, attr)
	SetStringDict(key, []string{
		"org", userkey_.Org,
		"K", userkey_.K,
		"KP", userkey_.KP,
	})
}

func (abeStorage AbeServiceStorageRedis) FetchCurve(curveName string) string {
	log.Debug("fetch curve")
	rc := GetClient()
	key := fmt.Sprintf("pub:curve:%s", curveName)
	x, _ := rc.HMGet(key, "params", "maxrndexp").Result()
	curve_ := &abe.AbeCurve_{x[0].(string), x[1].(string)}
	return abe.JsonObj2Str(curve_)
}

//FetchOrgPub call from external
func (abeStorage AbeServiceStorageRedis) FetchOrg(orgName string) string {
	log.Debug("fetching organization %s params", orgName)
	rc := GetClient()
	key := fmt.Sprintf("pub:org:%s", orgName)
	x, _ := rc.HMGet(key, "curve", "g1", "g2", "e").Result()
	org_ := &abe.AbeOrg_{x[0].(string), x[1].(string), x[2].(string), x[3].(string)}
	return abe.JsonObj2Str(org_)
}

//FetchAuthPub called from external
func (abeStorage AbeServiceStorageRedis) FetchAuthPub(authName string) string {
	log.Debug("fetching authority %s public params", authName)
	rc := GetClient()
	key := fmt.Sprintf("pub:auth:%s", authName)
	x, _ := rc.HMGet(key, "org", "ealpha", "g1y").Result()
	authpub_ := &abe.AbeAuthPub_{x[0].(string), x[1].(string), x[2].(string)}
	return abe.JsonObj2Str(authpub_)
}

//FetchAuthPrv called from external
func (abeStorage AbeServiceStorageRedis) FetchAuthPrv(authName string) string {
	log.Debug("fetching authority %s private params", authName)
	rc := GetClient()
	key := fmt.Sprintf("prv:auth:%s", authName)
	x, _ := rc.HMGet(key, "org", "alpha", "y").Result()
	authprv_ := &abe.AbeAuthPrv_{x[0].(string), x[1].(string), x[2].(string)}
	return abe.JsonObj2Str(authprv_)
}

//FetchUserkeyPrv called from external
func (abeStorage AbeServiceStorageRedis) FetchUserkeyPrv(user string, attr string) string {
	log.Debug("fetching key %s for %s", attr, user)
	rc := GetClient()
	key := fmt.Sprintf("prv:userkey:%s:%s", user, attr)
	x, _ := rc.HMGet(key, "org", "K", "KP").Result()
	userkey_ := &abe.AbeUserkey_{x[0].(string), x[1].(string), x[2].(string)}
	return abe.JsonObj2Str(userkey_)
}

func (abeStorage AbeServiceStorageRedis) FetchUserAttrs(user string) string {
	log.Debug("fetching user attributes %s", user)
	rc := GetClient()
	key := fmt.Sprintf("reg:user:%s", user)

	userattrs := make(map[string][]int)
	iter := rc.SScan(key, 0, "", 0).Iterator()
	for iter.Next() {
		userattrs[iter.Val()] = []int{}
	}
	return abe.JsonObj2Str(&abe.AbeUserAttrs_{user, userattrs})
}

func (abeStorage AbeServiceStorageRedis) FetchUserkeysPrv(user string, userattrs_json string) string {
	var userattrs abe.AbeUserAttrs_
	json.Unmarshal([]byte(userattrs_json), &userattrs)
	userkeys := make(map[string]string)
	for attr, _ := range userattrs.Coeff {
		userkeys[attr] = abeStorage.FetchUserkeyPrv(user, attr)
	}
	return abe.JsonObj2Str(&abe.AbeUserKeys_{user, userkeys})
}

func (abeStorage AbeServiceStorageRedis) AddOrgToCurveOrgSet(curveName string, orgName string) {
	key := fmt.Sprintf("reg:curve:%s", curveName)
	rc := GetClient()
	rc.SAdd(key, orgName)
}

func (abeStorage AbeServiceStorageRedis) AddAuthToOrgAuthSet(orgName string, authName string) {
	key := fmt.Sprintf("reg:org:%s", orgName)
	rc := GetClient()
	rc.SAdd(key, authName)
}

func (abeStorage AbeServiceStorageRedis) AddAttrToUserAttrSet(user string, attr string) {
	key := fmt.Sprintf("reg:user:%s", user)
	rc := GetClient()
	rc.SAdd(key, attr)
}
