package service

import (
	"encoding/json"
	"strings"

	abe "github.com/unicredit/abe/abe"
	log "github.com/unicredit/abe/log"
)

type AbeServiceStorage interface {
	InitStorage(config Config)
	StoreCurve(curveName string, curve_ *abe.AbeCurve_)
	StoreOrg(orgName string, org_ *abe.AbeOrg_)
	StoreAuthPub(authName string, authpub_ *abe.AbeAuthPub_)
	StoreAuthPrv(authName string, authprv_ *abe.AbeAuthPrv_)
	StoreUserkeyPrv(user string, attr string, userkey_ *abe.AbeUserkey_)
	FetchCurve(curveName string) string
	FetchOrg(orgName string) string
	FetchAuthPub(authName string) string
	FetchAuthPrv(authName string) string
	FetchUserkeyPrv(user string, attr string) string
	FetchUserAttrs(user string) string
	FetchUserkeysPrv(user string, userattrs_json string) string
	AddOrgToCurveOrgSet(curveName string, orgName string)
	AddAuthToOrgAuthSet(orgName string, authName string)
	AddAttrToUserAttrSet(user string, attr string)
}

var storage AbeServiceStorage

func InitAbeService(config Config, abeServiceStorage AbeServiceStorage) {
	if abeServiceStorage != nil {
		storage = abeServiceStorage
	} else {
		storage = new(AbeServiceStorageRedis)
	}
	storage.InitStorage(config)
}

// generate new curve and stores it
func SetupCurve(curveName string) string {
	log.Debug("setup curve %s", curveName)

	curve_json := abe.NewRandomCurve(curveName)

	var curve_ abe.AbeCurve_
	json.Unmarshal([]byte(curve_json), &curve_)
	storage.StoreCurve(curveName, &curve_)

	return curve_json
}

// generate new organization and stores it
func SetupOrg(orgName string, curveName string) string {
	log.Debug("setup organization %s on %s", orgName, curveName)

	curve_json := storage.FetchCurve(curveName)
	org_json := abe.NewRandomOrg(curve_json)

	var org_ abe.AbeOrg_
	json.Unmarshal([]byte(org_json), &org_)
	storage.StoreOrg(orgName, &org_)
	storage.AddOrgToCurveOrgSet(curveName, orgName)

	return org_json
}

// generate new authority (private and public params) and stores it
func SetupAuth(authName string, orgName string) string {
	log.Debug("setup authority %s in organization %s", authName, orgName)

	org_json := storage.FetchOrg(orgName)
	authpub_json, authprv_json := abe.NewRandomAuth(org_json)

	var authpub_ abe.AbeAuthPub_
	var authprv_ abe.AbeAuthPrv_
	json.Unmarshal([]byte(authpub_json), &authpub_)
	json.Unmarshal([]byte(authprv_json), &authprv_)
	storage.StoreAuthPrv(authName, &authprv_)
	storage.StoreAuthPub(authName, &authpub_)
	storage.AddAuthToOrgAuthSet(orgName, authName)

	return authpub_json
}

// generate new user key and stores it
func SetupUserkey(user string, attr string) string {
	log.Debug("generate key %s for %s", attr, user)

	authprv_json := storage.FetchAuthPrv(strings.SplitN(attr, "@", 2)[1])
	userkey_json := abe.NewRandomUserkey(user, attr, authprv_json)

	var userkey_ abe.AbeUserkey_
	json.Unmarshal([]byte(userkey_json), &userkey_)
	storage.StoreUserkeyPrv(user, attr, &userkey_)
	storage.AddAttrToUserAttrSet(user, attr)

	return userkey_json
}

func SetupUserkeyWithValue(user string, attr string, value int) []string {
	var userkeys_json []string
	attrs := abe.GetBagOfBits(attr, value)
	for _, attr := range attrs {
		userkey_json := SetupUserkey(user, attr)
		userkeys_json = append(userkeys_json, userkey_json)
	}
	return userkeys_json
}

// generate new random msg from organization name
func CreateRandomMsg(orgName string) string {
	org_json := storage.FetchOrg(orgName)
	return abe.NewRandomMsg(org_json)
}

func FetchUserAttrs(user string) string {
	return storage.FetchUserAttrs(user)
}

func FetchUserkeysPrv(user string, userattrs_json string) string {
	return storage.FetchUserkeysPrv(user, userattrs_json)
}

func FetchAuthPubs(auths_json string) string {
	log.Debug("fetching authority public keys")
	var auths abe.AbeAuthPubs_
	json.Unmarshal([]byte(auths_json), &auths)
	for attr, _ := range auths.AuthPub {
		auths.AuthPub[attr] = storage.FetchAuthPub(attr)
	}
	return abe.JsonObj2Str(&abe.AbeAuthPubs_{
		auths.AuthPub})
}
