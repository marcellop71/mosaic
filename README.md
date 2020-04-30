# abe

*abe* is a library doing Attribute Based Encryption.
The code implements is a Multi-Authority CipherText-Policy scheme from
[*Efficient Statically-Secure Large-Universe Multi-Authority
Attribute-Based Encryption*](https://eprint.iacr.org/2015/016.pdf) by Yannis Rouselakis and Brent Waters.

Converting boolean formulas for policies in LSSS schemes as explained in section G of [*Decentralizing Attribute-Based Encryption*](https://eprint.iacr.org/2010/351.pdf) by Allison Lewko and Brent Waters.
Using [*z3*](https://github.com/Z3Prover/z3) to check, simplify and rewrite policies.

## policies

1. attributes are in the form **attribute@authority** (cfr. notation already used in [*Charm*](http://charm-crypto.io/))
2. policies are arbitrary expressions in attributes with operators **/\\** (AND) and **\\/** (OR)
3. the library supporting extended syntax for int-valued attributes
4. policies allow **== < <= > >=** between an attribute and a numerical value
5. supporting also S-expressions for policies

examples (an authority **company** has been created)

1. can create attributes like **logistics@company** or **it@company** or **topsecret@company**
2. can create policy **logistics@company \\/ topsecret@company**
3. can create attributes like **level@company** with value 0 onwards
4. can create policy **topsecret@company \\/ (level@company > 0 )**
5. can write the policy as **(\\/ topsecret@company (> level@company 0))**

## Quick start

Dependencies:

- [*z3*](https://github.com/Z3Prover/z3) (*brew install z3* on Mac Os)
- [*pbc*](https://crypto.stanford.edu/pbc/) (*brew install pbc* on Mac Os)

Storage:

- *redis*: the examples use a Redis server available

## Examples

```bash
go run script/example_00.go
```

### setup curve, organization and authorities

```go
service.SetupCurve("SS512")       // setting up a curve
service.SetupOrg("org0", "SS512") // setting up an organization (hostng a set of authorities) onto a given curve
service.SetupAuth("rnd", "org0")  // setting up an authority into a given organization
```

### using boolean policies

```go
policy := "A@rnd /\\ (B@rnd /\\ (C@rnd \\/ D@rnd))"       // a policy as a standard boolean expression

// encrypting
passphrase := service.CreateRandomMsg("org0")             // -> create a random passphrase
policy = abe.CheckPolicy(policy)                          // -> check, simplify and rewrite policy
auths := abe.ExtractAuthsFromPolicy(policy)               // -> extracts authorities mentioned in the policy
auths = service.FetchAuthPubs(auths)                      // -> collects public keys of those authorities
passphrase_enc := abe.Encrypt(passphrase, policy, auths)  // -> encrypts the passphrase into a ciphertext

// user asking for keys
user := "marcello.paris@gmail.com"
service.SetupUserkey(user, "A@rnd") // -> for the user creates a key corresponding to the given attribute
service.SetupUserkey(user, "B@rnd") // -> for the user creates a key corresponding to the given attribute
service.SetupUserkey(user, "C@rnd") // -> for the user creates a key corresponding to the given attribute
service.SetupUserkey(user, "D@rnd") // -> for the user creates a key corresponding to the given attribute

// decrypting
policy = abe.ExtractPolicyFromCiphertext(passphrase_enc)            // -> extract the policy embedded in the ciphertext
userattrs := service.FetchUserAttrs(user)                           // -> collects the available user attributes
userattrs = abe.SelectUserAttrs(user, policy, userattrs)            // -> select which user attributes (if any) are useful for the given policy
userkeys := service.FetchUserkeysPrv(user, userattrs)               // -> collects user keys corresponding to the useful attributes
passphrase_dec := abe.Decrypt(passphrase_enc, userattrs, userkeys)  // -> decrypts the ciphertext into the passphrase plaintext
```

### integer-valued attributes

```go
policy := "(A@rnd > 1) /\\ B@rnd"
...
service.SetupUserkeyWithValue(user, "A@rnd", 5)
```

### curves

- https://crypto.stanford.edu/pbc/manual/ch08s08.html

## Use cases

### document sharing

- create a random message **msg** (function *CreateRandomMsg* outputs a base32 text)
- use this **msg** to encrypt (symmetric) a document
- create a **policy** string
- create an encrypted version **msg_enc** of the message **msg** according to the **policy**
- share the encrypted document and the **msg_enc**

now if a user needs to access the document:

- user asks its **keys**
- user decrypts the **msg_enc** into a clear **msg**
- user uses this **msg** to decrypt the document

## Core library

### design

Library could be exported as a shared object with *./script/build_lib.sh*
(resulting in *libabe.h* and *libabe.so* in *./lib*)

- every external function using only string I/O (plain or JSON)
- everything binary to [base32](https://en.wikipedia.org/wiki/Base32) (url and file names ok)

| function | description |
| -------- | ----------- |
| **NewRandomOrg** | new organization |
| **NewRandomAuth** | new authority |
| **NewRandomUserkey** | new user key |
| **NewRandomMsg** | new msg |
| **Encrypt** | encrypts a message |
| **Decrypt** | decrypts a message |
| **CheckPolicy** | verify, simplify and rewrite policy as a boolean S-expression |
| **ExtractPolicyFromCiphertext** | extract policy embedded in ciphertext |
| **ExtractAuthsFromPolicy** | service routine extracting authorities mentioned in a given policy |
| **SelectUserAttrs** | selects user attributes for decrypting |

### code

made up of 4 Go packages:
- *abe*: core package defining basic types and the cryptographic scheme
- *service*: service package defining interactions with databases
- *parser*: policy parsing generated by [Antlr](https://www.antlr.org/) using the *./grammar/policy.g4* grammar file
- *log* (logging)

file structure:

| file | description |
| -------- | ----------- |
| *./abe/arith_pbc.go* | *pbc* library wrapper (to be refactored using interfaces) |
| *./abe/crypto.go* | ABE algorithms |
| *./abe/encode.go* | encode/decode types <-> json |
| *./abe/policy.go* | parsing and policy management |
| *./abe/sss.go* | linear secret sharing scheme |
| *./abe/types.go* | structs |
| *./abe/utils.go* | generic functions |
