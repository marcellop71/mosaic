package main

import (
	"path/filepath"

	abe "github.com/unicredit/abe/abe"
	service "github.com/unicredit/abe/service"

	log "github.com/unicredit/abe/log"
)

var config service.Config

func main() {
	log.Init("Info")
	log.Debug("starting")

	filename, _ := filepath.Abs("config/config.yaml")
	e := service.ReadConfig(filename)
	config = e.Config

	abe.Init()
	service.InitAbeService(config, nil)

	service.SetupCurve("SS512")
	service.SetupOrg("org0", "SS512")
	service.SetupAuth("rnd", "org0")

	//policy := "(A@rnd == 5) /\\ B@rnd"
	policy := "(A@rnd > 1) /\\ B@rnd"
	//policy := "(A@rnd < 10) /\\ B@rnd"
	//policy := "A@rnd /\\ A@rnd"
	//policy := "A@rnd /\\ B@rnd"
	//policy := "A@rnd /\\ (B@rnd /\\ (C@rnd \\/ D@rnd))"
	//policy := "A@rnd /\\ ((D@rnd \\/ (B@rnd /\\ C@rnd)) \\/ A@rnd)"
	//policy := "A@rnd /\\ (D@rnd \\/ (B@rnd /\\ C@rnd))"
	//policy := "(/\\ A@rnd (\\/ D@rnd (/\\ B@rnd C@rnd)))"
	//policy := "(A@rnd \\/ C@rnd) /\\ (D@rnd \\/ (B@rnd /\\ C@rnd))"
	//policy := "(/\\ A@rnd (\\/ A@rnd D@rnd (/\\ B@rnd C@rnd)))"

	// encrypting
	msg := service.CreateRandomMsg("org0")
	log.Info("msg: %s", msg)
	policy = abe.CheckPolicy(policy)
	auths := abe.ExtractAuthsFromPolicy(policy)
	auths = service.FetchAuthPubs(auths)
	msg_enc := abe.Encrypt(msg, policy, auths)

	// user asking for keys
	user := "marcello.paris@gmail.com"
	service.SetupUserkeyWithValue(user, "A@rnd", 5)
	//service.SetupUserkey(user, "A@rnd")
	service.SetupUserkey(user, "B@rnd")
	service.SetupUserkey(user, "C@rnd")
	service.SetupUserkey(user, "D@rnd")

	// decrypting
	policy = abe.ExtractPolicyFromCiphertext(msg_enc)
	userattrs := service.FetchUserAttrs(user)
	userattrs = abe.SelectUserAttrs(user, policy, userattrs)
	userkeys := service.FetchUserkeysPrv(user, userattrs)
	msg_dec := abe.Decrypt(msg_enc, userattrs, userkeys)
	log.Info("msg reconstructed: %s", msg_dec)

	log.Debug("exit")
}
