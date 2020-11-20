# Mosaic

*Mosaic* is a library doing Attribute Based Encryption (ABE). It is meant to be used as a example of a cryptographic core to be embedded in any solution willing to rely on an ABE scheme of this kind
(a Multi-Authority CipherText-Policy scheme, see below).

### Please remark

The code is for demonstration purposes only and it is not suitable
for embedding in any reliable cryptographic solutions.
A quantity of structure which should be needed in any reliable cryptographic solution is missing in the code base
(about randomness, hashing functions used, selection of curves and pairings and underlying cryptographic library, hash-to-groups primitives,
secure communication, in-memory management of cryptographic calculations and many other issues).
Moreover, the code is not fully tested (nor formally verified),
nor yet implementing any reliable error management procedure and
should not be used in production environments.

The library is under the copyleft license [AGPL-3.0](https://choosealicense.com/licenses/agpl-3.0/).

__library is under development (APIs are subject to change) please check for updates__

- [Attribute Based Encryption](#attribute-based-encryption)
- [The scheme implemented](#the-scheme-implemented)
- [APIs](#apis)
- [Examples](#examples)
- [Curves and pairings](#curves-and-pairings)

## Attribute Based Encryption

Attribute based encryption is about sharing a secret according to attributes and policies.
A policy is an expression (boolean or integer-valued with linear constraints)
whose variables are called attributes.
The cryptographic scheme allows the encryption of a secret
according to a given policy, splitting the secret among the attributes mentioned in the policy.
The secret could be recovered only by a set of keys corresponding to attributes
which satisfy the policy (i.e., make the policy expression evaluate to true).

Those keys are attached to attributes and users (both strings or, in general, hashable objects).
Users can obtain these attribute-keys asking some authority for them
(so, the secret could only be recovered by those users having enough attributes as to satisfy the policy).

In a sense, attribute based encryption generalizes asymmetric cryptography:
the communication is no more 1:1 and the sender (encrypting) could ignore
the users allowed to receive (decrypt) her messages.

### An example: document sharing

Assume a company wants to manage the access to certain documents among its employees.
So, the company instantiate an authority __company__, so that users can ask this authority
for attributes like __logistics@company__ or __it@company__ or __topsecret@company__.
It is up to the authority to decide whether to give a given attribute-key to a user or not.

A user that needs to encrypt some documents can create a policy like
__logistics@company \\/ topsecret@company__ (meaning that the policy is meant to grant
access rights to those user having keys for __logistics@company__ or for __topsecret@company__),
and use this policy to encrypt the documents.
Now, the authorization layer to the documents is embedded in the mathematics,
meaning that only those users having the appropriate keys will be able
to decrypt the documents.

Policy can also include linear constraints on non-negative integer-valued attributes,
so, users can ask for attributes like __level@company__ with some value from 0 onwards,
to be used in a policy like __topsecret@company \\/ (level@company > 0 )__

## The scheme implemented

The code implements the Multi-Authority CipherText-Policy scheme as described in
[*Efficient Statically-Secure Large-Universe Multi-Authority
Attribute-Based Encryption*](https://eprint.iacr.org/2015/016.pdf) by Yannis Rouselakis and Brent Waters.

Specifically, the implementation
- is based on elliptic curves pairings, powered by [__miracl__](https://github.com/miracl/core) (recommended) (code included in this repo) or [__pbc__](https://crypto.stanford.edu/pbc/) (including parts of the Go wrapper https://github.com/Nik-U/pbc) [rf. to config file later in this doc]
- uses linear secret sharing schemes as explained in section G of [*Decentralizing Attribute-Based Encryption*](https://eprint.iacr.org/2010/351.pdf) by Allison Lewko and Brent Waters
- allows numerical expressions in policies (see paragraph below) using
a bag-of-bits representation of attributes as in [*Ciphertext-Policy Attribute-Based Encryption*](https://www.cs.utexas.edu/~bwaters/publications/papers/cp-abe.pdf) by John Bethencourt, Amit Sahai, Brent Waters
- uses [__z3__](https://github.com/Z3Prover/z3) to verify, simplify and rewrite policies (it could use __z3__ also to solve and get hints from policies)
- collects the authorities into organizations: policy are allowed to mention only attributes from authorities within a given organization
- uses an elementary hash-to-group function to map attribute and user strings to the group

Syntax for attributes and policies:

- boolean attributes are in the form __attribute@authority__ (cfr. notation already used in [__Charm__](http://charm-crypto.io/))
- integer-valued attributes are written __attribute=value@authority__
- policies are arbitrary expressions in attributes with operators __/\\__ (AND) and __\\/__ (OR)
- policies allow __== < <= > >=__ between an attribute (on the left) and a numerical value (on the right)
- supporting also S-expressions for policies

## APIs

To use the library in a Go environment

```go
import "github.com/marcellop71/mosaic/abe"
```

Core __Go__ APIs are wrapped by JSON APIs so that every function
has only __string__ type parameters as inputs and outputs
(__string__ could be a plain or a JSON string depending on the function).

Everything binary is encoded to [base32](https://en.wikipedia.org/wiki/Base32)
(both JSON strings and field values in JSON objects).

Given the JSON APIs, the library could also be exported as a shared object (resulting in *libmosaic.h* and *libabe.so* in *./lib* usable by many languages). The wrapping is trivial because there are only __char*__ in each signature.

```bash
go build -tags=z3,miracl -buildmode c-shared -o lib/libmosaic.so lib/mosaic.go
```

(or *lib/mosaic_pbc.go*)

| GO API | GO API (JSON) | C API |
| ---- | ---- | ---- |
| __NewRandomOrg__ (lib string, curve Curve) Org | __NewRandomOrgJson__(string, string) string | __newRandomOrg__ |
| __NewRandomAuth__(org Org) AuthKeys | __NewRandomAuthJson__(string) string | __newRandomAuth__ |
| __NewRandomUserkey__(user string, attr string, authprv AuthPrv) Userkey | __NewRandomUserkeyJson__(string, string, string) string | __newRandomUserkey__ |
| __NewRandomSecret__(org Org, seed string) Point | __NewRandomSecretJson__(string, string) string | __newRandomSecret__ |
| __AuthPubsOfPolicy__(policy string) AuthPubs | __AuthPubsOfPolicyJson__(string) string | __authpubsOfPolicy__ |
| __PolicyOfCiphertext__(ct Ciphertext) string | __PolicyOfCiphertextJson__(string) string | __policyOfCiphertextJson__ |
| __SelectUserAttrs__(user string, policy string, userattrs Userattrs) Userattrs | __SelectUserAttrsJson__(string, string, string) string | __selectUserAttrs__ |
| __Encrypt__(secret Point, policy string, authpubs AuthPubs) string | __EncryptJson__(string, string, string) string | __encrypt__ |
| __Decrypt__(ct Ciphertext, userattrs Userattrs) Point | __DecryptJson__(string, string) string | __decrypt__ |
| __RewritePolicy__(policy string) string | __RewritePolicy__(string) string | __rewritePolicy__ |
| __CheckPolicy__(policy string, userattrs Userattrs) string | __CheckPolicyJson__(string, string) string | __checkPolicy__ |
| __GetPbcCurve__(curve string) Curve | __GetPbcCurve__(string) string | __getPbcCurve__ |

__RewritePolicy__ and __CheckPolicy__ needs __z3__ support, while __GetPbcCurve__ needs __pbc__.
Please remark that the support by __z3__ and __miracl__ (or __pbc__) is managed by the build tags "z3", "miracl" and "pbc". For example, run the example with

```bash
go run -tags=z3,miracl examples/ex_JsonAPI.go
```

Here, the signature is described by a suffix **Str** (plain string)
or **Json** (json string, for the objects serialized as json strings).

| export | signature | json schema for output |
| ---- | ---- | ---- |
| __newRandomOrg__ | __libStr, CurveJson (empty if not pbc) -> orgJson__ | json to be used as-is |
| __newRandomAuth__ | __orgJson -> authkeysJson__ | `{ "authpub": "authority public key json string", "authprv": "authority private key json string"}` |
| __newRandomUserkey__ | __userStr, attrStr, authprvJson -> userattrsJson__ | `{"user": "user string", "coeff": {"attribute0": [], ...}, "userkey": {"attribute0": "userkey json string", ...}}` |
| __newRandomSecret__ | __orgJson, seedStr -> secretStr__ | plain string |
| __rewritePolicy__ | __policyStr -> policyStr__ | plain string |
| __checkPolicy__ | __policyStr, userattrsJson (possibly empty) -> string__ {"sat", "unsat"} | plain string |
| __extractAuthsFromPolicy__ | __policyStr -> authpubsJson__ | `{"authpub": {"authority0": "authority public key", "authority1": "authority public key json string", ...}}` |
| __extractPolicyFromCiphertext__ | __ciphertextJson -> policyStr__ | plain string |
| __selectUserAttrs__ | __userStr, policyStr, userattrsJson -> userattrsJson__ | `{"user": "user string", "coeff": {"attribute0": [], ...}, "userkey": {"attribute0": "userkey json string", ...}}` |
| __encrypt__ | __secretJson, policyStr, authpubsJson -> ciphertextJson__ | json to be used as-is |
| __decrypt__ | __ciphertextJson, userattrsJson -> secretStr__ | plain string |
| __getPbcCurve__ (only for pbc) | __curveStr -> CurveJson__ | json to be used as-is |

The json strings returned by __getCurve__, __newRandomOrg__ and __encrypt__ are meant to be used as-is: stored somewhere or re-fed into the relevant functions.

An application needs to interact only with __authkeysJson__, __authpubsJson__ and __userattrsJson__.
__authkeysJson__ holds the public and private keys for a new authority.
__authpubsJson__ is a json holding in the key 'authpub' the association between an authority (name)
and its corresponding public key, while __userattrsJson__ holds the associations
between an attribute and its use (coefficent) in the key "coeff" and
between an attribute and its corresponding user key in the key "key".

```go

type Org struct { ... } // organization
type AuthPub struct { ... } // authority public params
type AuthPrv struct { ... } // authority private params
type Userkey struct { ... } // user key
type Ciphertext struct { ... } // ciphertext encrypting the msg

type AuthKeys struct { // authority public and private keys
	AuthPub string `json:"authpub"` // type AuthPub
	AuthPrv string `json:"authprv"` // type AuthPrv
}

type AuthPubs struct { // map of authorities public keys
	AuthPub map[string]string `json:"authpub"` // type AuthPub
}

type UserAttrs struct { // map of user attributes
	User string `json:"user"` // user
	Coeff map[string][]int `json:"coeff"` // attr -> its coefficients
	Userkey map[string]string `json:"userkey"` // Userkey
}
```

### Dependencies

The only external dependency that needs to be available is [__z3__](https://github.com/Z3Prover/z3) (optionally, also [__pbc__](https://crypto.stanford.edu/pbc/)).

Internally, the Go packages used are: [__logrus__](https://github.com/sirupsen/logrus) for logging and [__antlr4__](https://github.com/antlr/antlr4) for parsing.

The examples (not the library, see below) use
- [__Leveldb__](https://github.com/google/leveldb) (via the Go wrapper https://github.com/jmhodges/levigo)
- or a [__Redis__](https://redis.io/) server (via the Go wrapper https://github.com/go-redis/redis)

Dependencies or database support could be easily installed:
- mac os: *brew install z3* and *brew install leveldb redis* (in case also, *brew install pbc*)
- linux: see __./docker/Dockerfile__

### Docker

You can build a docker images with everything:

```bash
./docker/build_image.sh
```

Source is made up of 2 Go packages:
- *abe*: core package defining basic types and the cryptographic scheme
- *service*: service package defining interactions with databases

The file structure of the core lib is the following:

| file | description |
| -------- | ----------- |
| __arith_miracl.go__ | __miracl__ interface |
| __arith_pbc.go__ | __pbc__ interface |
| __crypto.go__ | ABE scheme algorithms and JSON APIs |
| __policy_z3.go__ | policy management using __z3__|
| __policy.go__ | parsing and policy management |
| __sss.go__ | linear secret sharing scheme |
| __types.go__ | structs and their serialization |
| *Policy.g4* | __antlr4__ parser grammar |
| *PolicyLexer.g4* | __antlr4__ lexer grammar |
| *log/* | logging |
| *miracl/* | __miracl__ core library (__Go__ version) |
| *parser/* | policy parser (generated by [__antlr4__](https://www.antlr.org/) using the *./abe/policy.g4* grammar file) |
| *pbc/* | __pbc__ library wrapper |
| *z3/* | __z3__ library wrapper |

## Examples

An __./examples/ex_GoAPI_noservice_noz3.go__ is there to show how the APIs could be used.
The aim is to encrypt a secret according to a given policy (iterating
over a small set of policies) and recover it (if feasible),
according to a selection of attributes given to a user.

```bash
go run -tags=z3,miracl examples/ex_GoAPI_noservice_noz3.go
```

or

```bash
docker run -it mosaic:latest go run -tags=z3,miracl examples/ex_GoAPI_noservice_noz3.go
```

### Storage

__./examples/ex_GoAPI_noservice_noz3.go__ is an application using the __github.com/mosaic/abe__ package
and a __./mosaic/service__ package to handle some storage for the keys,
the attributes and the users.
The service package implemented can use [__Leveldb__](https://github.com/google/leveldb)
or a [__Redis__](https://redis.io/) server.
Details in the configuration file __./examples/config.yaml__

```yaml
config:
  arithmetic:
    curve: BN254
    library: miracl
  storage:
    redis:
      local0:
        addr: "127.0.0.1:6379"
        password: ""
    leveldb:
      local0:
        name: "mosaic.db"
  active:
    type: "leveldb"
    label: "local0"
```

### Workflow

The workflow implemented in the example is basic.
Using attribute based encryption for a document sharing application
could mean that an encryptor:

- creates a random __secret__
- uses this __secret__ to encrypt (symmetric) a document
- creates a __policy__ string describing access to the document
- creates an encrypted version __secret_enc__ of the __secret__ according to the __policy__
- shares the encrypted document and the __secret_enc__ + an hashed version of the secret __secret_hash__

now if a user needs to access the document:

- user asks an __authority__ for its __keys__ corresponding to a given attribute
- user decrypts the __secret_enc__ into a clear __secret__
- user checks if she can recover the __secret_hash__
- user uses this __secret__ to decrypt the document

### Setup curve, organization and authorities

```go
service.SetupOrg("org0", "miracl", "BN254")   // setting up an organization (hosting a set of authorities) onto a given curve
service.SetupAuth("auth0", "org0")  // setting up an authority into a given organization
```

### Using boolean policies

```go
policy := "A@auth0 /\\ (B@auth0 /\\ (C@auth0 \\/ D@auth0))" // a policy as a standard boolean expression

// encrypting
secret := service.NewRandomSecret("org0", seed)  // create a new secret
policy = abe.RewritePolicy(policy)               // simplify policy
auths := abe.AuthPubsOfPolicyJson(policy)      // extracts authorities mentioned in the policy
auths = service.FetchAuthPubs(auths)             // collects public keys of those authorities
secret_enc := abe.EncryptJson(secret, policy, auths) // encrypts the secret into a ciphertext

// user asking for keys
user := "marcello.paris@gmail.com"
service.SetupUserkey(user, "A@auth0") // for the user creates a key corresponding to the given attribute
service.SetupUserkey(user, "B@auth0") // for the user creates a key corresponding to the given attribute
service.SetupUserkey(user, "C@auth0") // for the user creates a key corresponding to the given attribute
service.SetupUserkey(user, "D@auth0") // for the user creates a key corresponding to the given attribute

// decrypting
policy = abe.PolicyOfCiphertextJson(secret_enc)     // extract the policy embedded in the ciphertext
userattrs := service.FetchUserAttrs(user)                // collects the available user attributes
userattrs = abe.SelectUserAttrsJson(user, policy, userattrs) // select which user attributes (if any) are useful for the given policy
userattrs = service.FetchUserkeys(userattrs)             // collects user keys corresponding to the useful attributes
secret_dec := abe.DecryptJson(secret_enc, userattrs)         // decrypts the ciphertext into the secret plaintext
```

### Integer-valued attributes

```go
policy := "(A@auth0 > 1) /\\ B@auth0"
...
service.SetupUserkey(user, "A=5@auth0")
```

## Curves and pairings

Use "BN254" or "BLS12381" or "BN462" if using [__miracl__](https://github.com/miracl/core) (recommended)
or use "SS512" or "BN254" if using [__pbc__](https://crypto.stanford.edu/pbc/).

Let the pairing be G1 x G2 -> GT

The scheme encrypts a point in GT as:
- a point on GT (the point to be encrypted perturbed by a secret)
- a bags of 4-points tuples (i.e., bag of points on GT x G1 x G1 x G2) [one tuple for each attribute-leave in the policy]
