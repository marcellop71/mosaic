package main

import (
	"crypto/sha256"
	"path/filepath"

	"github.com/unicredit/mosaic/abe"
	"github.com/unicredit/mosaic/abe/log"
	"github.com/unicredit/mosaic/service"
)

func main() {
	log.Init("Info")

	filename, err := filepath.Abs("./examples/config.yaml")
	if err != nil {
		log.Panic("%s", err)
	}
	config := service.ReadConfig(filename).Config
	service.InitAbeService(config, nil)

	lib := "pbc"
	curve := config.Arithmetic.Curve
	service.GetPbcCurve(curve)

	org := "org0"
	service.SetupOrg(org, lib, curve)
	log.Info("set up organization %s on curve %s", org, curve)

	auths := []string{"auth0", "auth1"}
	for _, auth := range auths {
		service.SetupAuth(auth, org)
	}
	log.Info("set up authorities %s on organization %s", auths, org)

	user := "marcello.paris@gmail.com"
	attrs := []string{"A@auth0", "B@auth0"}
	for _, attr := range attrs {
		service.SetupUserkey(user, attr)
	}
	log.Info("user %s asking for keys for attributes %s", user, attrs)

	policies := []string{
		"A@auth0",
		"A@auth0 /\\ B@auth0",
		"A@auth0 /\\ A@auth0",
		"A@auth0 /\\ (B@auth0 /\\ (C@auth0 \\/ D@auth0))",
		"A@auth0 /\\ ((D@auth0 \\/ (B@auth0 /\\ C@auth0)) \\/ A@auth0)",
		"(A@auth0 \\/ C@auth0) /\\ (D@auth0 \\/ (B@auth0 /\\ C@auth0))",
		"(/\\ A@auth0 (\\/ A@auth0 D@auth0 (/\\ B@auth0 C@auth0)))",
	}

	for _, policy := range policies {
		log.Info("----------------")
		log.Info("policy: %s", policy)
		if (abe.CheckPolicy(policy, "") == "sat") {
			// ecnrypting
			seed := ""
			secret := service.NewRandomSecret(org, seed)
			secret_hash := sha256.Sum256([]byte(secret))
			log.Info("secret hash: %s", abe.Encode(secret_hash[:]))
			policy = abe.RewritePolicy(policy)
			auths := abe.ExtractAuthsFromPolicy(policy)
			auths = service.FetchAuthPubs(auths)
			secret_enc := abe.Encrypt(secret, policy, auths)

			// decrypting
			policy = abe.PolicyOfCiphertextJson(secret_enc)
			userattrs := service.FetchUserAttrs(user)
			if (abe.CheckPolicy(policy, userattrs) == "sat") {
				userattrs = abe.SelectUserAttrs(user, policy, userattrs)
				userattrs = service.FetchUserkeys(userattrs)
				secret_dec := abe.Decrypt(secret_enc, userattrs)
				secret_dec_hash := sha256.Sum256([]byte(secret_dec))

				if (abe.Encode(secret_dec_hash[:]) == abe.Encode(secret_hash[:])) {
					log.Info("secret correctly reconstructed")
				} else {
					log.Info("secret not correctly reconstructed")
				}
			} else {
				log.Info("user cannot satisfy the policy")
			}
		} else {
			log.Info("policy unsatisfiable")
		}
	}
}
