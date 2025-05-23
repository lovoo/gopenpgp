package crypto

import (
	"bytes"
	"testing"

	"github.com/lovoo/gopenpgp/v3/internal"
)

func TestDetachedSignaturesWithUnknownPackets(t *testing.T) {
	const ricarda_key_armor = `-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: 2ADE 0F8A A059 6BC9 4E50  D2AD 9162 53AB 652E F195

xsDNBGJvrDABDACrZUY7VU1uZ/uzntlAKEHVF4mb3eYSdW1rE3hVke0HoDvtQbzq
KQ9qgfPaNwdkxRexgrNGSeKkQJcgR7gMWxFxM/FwddQIXfVL43nRlN+iGvFDYR+9
dn5gSOD9EvZUYLN9p6yR3Umyglt4NpdjYM0J+Rn2DVyfGHCtS+z1fdym1h1zdImo
rArBpWMEdvNN/6dR8BN67WSBs5pVsvnDPdjbeU+GPJVoRH4CWe/LdJnJDICPmlva
gAyeJeK+KitkxD8IIc9d18x5dV1Z/LL2o1B0Psort8+az2Z+NbkP2cUv0DDRyZIM
Ww1A6KMSSNuvewXGCU+gEFlGIA83zfq7XE3WNp9EPy9xCs610+KxSS9cftKdrMmU
cl37Gb1lZyLFIBUCwPdIyTvhs6r6rMgu4GQ46sPl8qeLbtaoXnPWj9V1Fn9RoMgz
Kq8ZItFCJfCbf+gamx/0S/0mE47giuAtymw8Hf53e3jsSANK28GpcsEUNqWT4djb
YgcWl2iXR7n0i9cAEQEAAcLBSQQfAQoAfQWCYm+sMAMLCQcJEJFiU6tlLvGVRxQA
AAAAAB4AIHNhbHRAbm90YXRpb25zLnNlcXVvaWEtcGdwLm9yZ6O6vv9v4H1aAR9Z
aEksrL/2rQGm4Omv2SCubCC1pk/DAxUKCAKbAQIeARYhBCreD4qgWWvJTlDSrZFi
U6tlLvGVAACp1wv+MEkIs8x2Wk6e+i4rbclnkBjR2tSwv3N66jcz09QeDt5xwX1Q
jq1TP+XGFJfFA+LUugbLWNzWqwh7CypSc6IAB82+Ha+lCMjY3SZLfTe6crUgDWOZ
wDprVND6z2g9UlgJ6rfFf329LEUziknjKFBbreONH12tOzvCxrxipp84J0DZEqO3
0VB2b2B7uO1X5V1aQMFj4JEafQ6Or6uXtR3Fwcs/JALwhqydy3LI4Y0s/y91qdPP
69DgugVDbG/Kkp5W2CNtGyd51cnEoFEbh9nV0hHR+rERU2ZcyoNfZQCyPuHMIT1L
PndcWLsT6Ga3TvXLUEhOOl6CPpox0aV7bsGpip0u54yGkhgR5YYLchcTVt15q5+/
M5Z03et1LcSZKk1ijSFXNsOWdD2ojneN3vhlKbJoDosT10WyCsXKL8JNP4F4q4nU
22+/N1sttbgzSatwPxyiC7kKymQmDtg8NCyyCtpDoRkEIzaMv+4DqoMVkvLRcF1K
vX+rNEg1OxneqMSEzS1SaWNhcmRhIFMuIMOBbHZhcmV6IDxyaWNhcmRhQG9wZW5w
Z3AuZXhhbXBsZT7CwUwEEwEKAIAFgmJvrDADCwkHCRCRYlOrZS7xlUcUAAAAAAAe
ACBzYWx0QG5vdGF0aW9ucy5zZXF1b2lhLXBncC5vcmfo22n7d9n+Nv3PxzMP93f7
rejDV4BFsobH6T5cZGvj8AMVCggCmQECmwECHgEWIQQq3g+KoFlryU5Q0q2RYlOr
ZS7xlQAAu6EL/iJ91h1dwcQMZpbKYnBbH/1U/ncbIlhkQf8YzjrpzVprU0sr4cKu
mx42EylcteODVoDN8ITGROUC8Prw5tu+NS2mjkKWUlj56eLkpS3A3D6NqPIFEzU3
KYztp+Y0hq2ZmsGsPgL1AbVPc2+ngUSL4pXgQ5hmVoD3B6VGqBReyKXhZNzH42S5
fMuOsY8+/xgQ0WDwSCVeUW7Otql24iFji9U4eL036AzOO1hXkbOesrw4E7GmU2mR
kkX/aDZ9bGYT3yVZ/CJkU4wrUHJWW1IJl/bl//becL8Vnqr92vQkYQnUh19zIi2n
6ualROjcoITjjRemvkBfM9zXTd6kVFKD+ySDnGtNq62Ukz0/COg81tAnnnXeNhlP
1M8yA6zfrY17tAFYLmALUrPYjVy4ZJuaHScnIcM5lHIKYn1ijetGQSjI4sT/mSi4
8fwNHFRufR7ta9XWv5b4+m8SAyrJ/FqnRoOwtoph0ZxSItjpv+qLPX2N4M8/JOqh
3QhI5ZDd+ScGJc7AzQRib6wwAQwArMv2IdO1f8a/brBc/+LeCqyqR+qLKzZRLAvF
+c6+qG7D9OzZCJfbPltnQ5BgRzIiSKgzDugDDt0m3fdWBxOQG5Ojo53Xu1ZTgRUD
0KWI4kH5Cs0gp3Kl62TdNAxNUlHdWqJ6G6WYVMlMhLmBVPjo65pT1+OON2v/O/qA
8ZJRlK3RyKym4Kcr8JbtX1roSRtVPRpi++Y1CGuhnK3UM68putqfgZOnTZxln2m2
BUrL8fWgme2rTJyrrN8fSM61dZZ+5E0SWDMBxEp5bimMrRjJc7V0JMUx2HiCDe/l
panzleAugfEVUeAYsEe9C+7k3p2r00roL0dWnrMgfhGnJy5ccmPyzEQpoA6ARq1a
EBk2Syr/QNO93GfDFcK9pTcDcQdUgq93x5s9LShLQ7fuvVXXxkl4QkdKJn+Vb2yP
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
=i+jK
-----END PGP PUBLIC KEY BLOCK-----
`
	const sig = `-----BEGIN PGP SIGNATURE-----

wsE7BAABCgBvBYJkit2rCRD7/MgqAV5zMEcUAAAAAAAeACBzYWx0QG5vdGF0aW9u
cy5zZXF1b2lhLXBncC5vcmelgsnbOPsc1mt9YRUNBmoAWxgPfqyzTjfUu2WQ1NQd
MhYhBNGmbhojsYLJmA94jPv8yCoBXnMwAAC0TAv8CsSeavcLIjJvUKBMV+I09fmA
D/HCf9ZtV7Y+JPOaOMBgWnNaoVINRenDbQLz69Yq8Xf8zyOOdX1GUpiO/0vdHnw2
TdJc6tkWB2yP63DfwGQD5xxSzltfvO1GV7BUYVIUXZqATp39wp7nmn6BJ15DtlLT
mUwwQ0eDfwtAGcyZBfDkz1V7p+AYuWWDCTRXk0LVl8mjsXheX1BX4uzlru6LWE88
JL1uNDcRfvg1WDBCV0zTQwWlmmsWF42KCft6p/c+ak5XstzEKalnDtASoU/o07bF
QXVr2ncHexfjyyZRo9ZH0s4jn71J83Djou4O6O9zgo19kL1aErEDC3J9n041qvvV
xi5ZVvmVVeTrGrJ71UyoMCRxyFu+vtBnvoRluLcw5NxKgniASJCRtXvUYYxG048I
9m6pSkC41QlEc+4r325i/Ny4p6VLwZDK+APN1MQRd85MHjKpGNAiVHYLtdAzdABr
VpAFLmdML2VLIiCG9R6Zw8U2Hww5DSy5Gq0UM1tJwsE7BAABCgBvBYJkit2rCRCz
+fFdlxFkrkcUAAAAAAAeACBzYWx0QG5vdGF0aW9ucy5zZXF1b2lhLXBncC5vcmdI
0DufwhTqCzYcRYQp4VAE+Qxp2Ruj+srVlIBkA20qZRYhBN9WS6uCrGY2n74Fz7P5
8V2XEWSuAADSCwwAhfAVzNHeLEgrXHJAgrpP31ubBebEVmOl3jUPg4+cko2LJBGE
egeONp4kVrxMmk9AKURfQ4UaVFXh6Cyhjq5zFZQJnZ2fmHJm/L1A7r+XUwTQ2dZy
LGVSZCPNbKEuJxcZzYO9fDtbJIoY5OeIBFmasvbGle4laotKhleMopfqzJLuZJIa
zxDAsJbdcm/sVbHLeElK4cjcEtiZZKNizzZnTE827WlZVll1FZ+a3Gi8SQXV+asy
aRVUb1hw1VzFj3ZmtcRIPctEy5uYZq8Oaa0U8Z+J10ertSGh/2I1Au7G7SXiYVzB
65LXj4daJ3JZTCjpqc82AJdnpVpREPpCfaopUchA+NYNaIaolx/SmgkeLxyWXff+
5GHfuVDQpD8Ma7hOx6aP+Ih/xI8OrMgx9sZ79L9mdbcSCaypqYd/oez1pwrfvcN8
c6rmTo+LPGN1CKceOyPdcKEZBnwY/QrNf3GDgZSnm69Iol16OvqstTrFU26z1WT6
sWC04HotaaBXdIHf
=nNOz
-----END PGP SIGNATURE-----`
	message := "Hello World :)"

	key, err := NewKeyFromArmored(ricarda_key_armor)
	if err != nil {
		t.Error("error: ", err)
	}
	pgp := PGP()
	verifier, err := pgp.Verify().VerificationKey(key).New()
	if err != nil {
		t.Error("error: ", err)
	}
	res, err := verifier.VerifyDetached([]byte(message), []byte(sig), Armor)
	if err != nil {
		t.Error("error: ", err)
	}
	if err = res.SignatureError(); err != nil {
		t.Error("Should have no signature error, got:", err)
	}
}

func TestSignedAndEncryptedMessages(t *testing.T) {
	const bob_key_armor = `-----BEGIN PGP PRIVATE KEY BLOCK-----
Comment: Bob's OpenPGP Transferable Secret Key

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
=miES
-----END PGP PRIVATE KEY BLOCK-----
`
	msg := `-----BEGIN PGP MESSAGE-----
Comment: Plaintext is "Hello\r\nWorld!\n"

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
Ik4fzHKnqkcmTwUng+v5sDcU2R2x4Rel6YuR
=knHG
-----END PGP MESSAGE-----
`
	expected := "Hello\r\nWorld!\n"

	key, err := NewKeyFromArmored(bob_key_armor)
	if err != nil {
		t.Error("error: ", err)
	}
	pgp := PGP()
	decryptor, err := pgp.Decryption().DecryptionKey(key).VerificationKey(key).New()
	if err != nil {
		t.Error("error: ", err)
	}
	res, err := decryptor.Decrypt([]byte(msg), Armor)
	if err != nil {
		t.Error("error: ", err)
	}
	if err = res.SignatureError(); err != nil {
		t.Error("Should have no signature error, got:", err)
	}
	if !bytes.Equal(res.Bytes(), []byte(expected)) {
		t.Errorf("Should be equal, got: %s", res.Bytes())
	}
}

func TestProblematicCleartextSignature(t *testing.T) {
	const bob_key_armor = `-----BEGIN PGP PRIVATE KEY BLOCK-----
Comment: Bob's OpenPGP Transferable Secret Key

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
=miES
-----END PGP PRIVATE KEY BLOCK-----
`
	const message = `From the grocery store we need:

- tofu
- vegetables
- noodles

`
	key, err := NewKeyFromArmored(bob_key_armor)
	if err != nil {
		t.Error("error: ", err)
	}
	pgp := PGP()
	signer, err := pgp.Sign().SigningKey(key).New()
	if err != nil {
		t.Error("error: ", err)
	}
	msg, err := signer.SignCleartext([]byte(message))
	if err != nil {
		t.Error("error: ", err)
	}
	verifier, err := pgp.Verify().VerificationKey(key).New()
	if err != nil {
		t.Error("error: ", err)
	}
	res, err := verifier.VerifyCleartext(msg)
	if err != nil {
		t.Error("error: ", err)
	}
	if err = res.SignatureError(); err != nil {
		t.Error("Should have no signature error, got:", err)
	}
	if !bytes.Equal(res.Cleartext(), []byte(message)) {
		t.Errorf("Should be equal, got: %s", res.Cleartext())
	}
}

func TestDetachedSignaturesLowRSABits(t *testing.T) {
	const ricarda_key_armor = `-----BEGIN PGP PRIVATE KEY BLOCK-----

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
KNs/Xw4PLUnp6wQO2d/6oZay
=ffeX
-----END PGP PRIVATE KEY BLOCK-----	
`
	const sig = `-----BEGIN PGP SIGNATURE-----

wsE7BAABCgBvBYJkpnJFCRB8L6pN+Tw3skcUAAAAAAAeACBzYWx0QG5vdGF0aW9u
cy5zZXF1b2lhLXBncC5vcme7t8QmfP04d0zI8u8HUPmZcNbExpgA23j+/ICq8B3N
EhYhBB3c4V8JIXzuLzs3YHwvqk35PDeyAAA9fwwAtaSuGkMxdU3CyO1n8x7mBaZS
F2QNzvUKlhei1vPf5RaXTlNmlDWl/wF7aqRlrbdSHnqvi5lbhVVNP5LU62tS1J1N
WudlfoVhIexYOtoYOvOnsR3nBT95Mevk64UTVPs2rAfNDSD/0TKtK/gL+B9UBI2K
8UJwxcjtO116qwsO6gJq5XJjn0y34nNdIJH8Rr5I1jtl6h3Xlh2et58Yp9LmubNw
mrqvl+ES8bh8Q++KpesnswQnSZR9e4lp6FY6w7X6UDyxtJF9S0SBZnsduJtsoxUl
U8J1SUKwl0ATk9a+4+bwkwK9UJXFzwqEJvLoMe4U5QJKYazosr7Mxw197/iYOLXD
LlSSmcS1v5Oo37CVG+cZZnyiB1vra7x1fnpR+EA/bbsKkXXQjAepMhb61IFq44dL
j/6t75dByS9g29r2Fen+KFHh4QZrai9jEabepP9nKgNza7hIEeZtVdNyJfENUfQd
l6qOAQGHaxD0whADX2J53HqPUpcbidzw3jEKwWgI
=+Dhj
-----END PGP SIGNATURE-----
`
	message := "Hello World :)"

	key, err := NewKeyFromArmored(ricarda_key_armor)
	if err != nil {
		t.Error("error: ", err)
	}
	pgp := PGP()
	verifier, err := pgp.Verify().VerificationKey(key).New()
	if err != nil {
		t.Error("error: ", err)
	}
	res, err := verifier.VerifyDetached([]byte(message), []byte(sig), Armor)
	if err != nil {
		t.Error("error: ", err)
	}
	if err = res.SignatureError(); err == nil {
		t.Error("Should have signature error", err)
	}
}

func TestDetachedSignaturesAndroid(t *testing.T) {
	const ricarda_key_armor = `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GopenPGP 2.1.1
Comment: https://gopenpgp.org

xsBNBF1BfxUBCADUpiiG3AhQK08E2nBmQ50XeztOWArmknINQV41pqGFW5VQkfbQ
3FYsANhLGqbDBQ0XxmocjKL7W7W8Y4xmHCGgkCUy6gAqGbi+sXY9Sl8xqQNHuZDh
WVdqT8+Rtv+DRxp/XrGkzC1U8CBYUmmKS92ldy0/zZIvgQXT6t5Q+v+BeUSv4jCs
nY3BE0UBOljtrTXlOcXRZHQxORWG+kon0qgcJERdwwzhxY6eT8jEfAfJY0hzQaYg
+6bj6ZR0zkMtY2Psq2M05kzEw4On/dezZETAu1e9fSqfk1mp+H6BeLJ9RUyrFK/P
qIO48+pU8CmAvTdx5eIihyOM16CFg/3GgV85ABEBAAHNMWFkYW10c3RAcHJvdG9u
bWFpbC5ibHVlIDxhZGFtdHN0QHByb3Rvbm1haWwuYmx1ZT7CwGgEEwEIABwFAl1B
fxUJEBHDHo5eB/TQAhsDAhkBAgsJAhUIAADHoggACDYDZkyTMZX69k9uoAygAQ75
2kb52r0L3dSLge+hUelxJOiVUznbavzVhzjzF2FucXP0csOSJygHNejjS62BDtsX
iIoPiVDO1+Hr72X4gE277VeZ1b6VozJvKC0+H4fhg/EtkD07oVhHJRxOOVlgUXGK
nP2lz2ojny0McggvN+Nk8ILqt6ImlvEk6CnTs9XdmcmosMiQU+U+THKrKZ+5Yec8
4kzlHG8ee7Tim2yn9n/FuBStrYkTJUsDuAL/LOfF9DnzTzukK6kqpDB6kDfMeYQJ
Lq+Tu642n74P0lqOO0Wy7imI/hxM1W8yqcNdafS7PCuGHD99mecdKWVeYHCCY87A
TQRdQX8VAQgAyAIj6ytLB1CUinP6f6aVKlLSM9e9Figv1CAw53NHeQBbmwzU5bZn
tE6UERnvJy6ul7hJr9Jop+O1/jA6zaGanF5wv0nEvTHcoYRpJ4QiJgiQxvhOdItH
29+jBV1F44xOzlGnEzFAv7GbPecKHAsQgX9qYCj+5ydcttQ29gWQ6nN23G03R3Lb
KRS9H2uw1SIYGgif8FgKpJemwJjuSibyViXTf3JC8ZUtYbq+vIXqATFFtbrUHfKM
AKlHo0uLYGq1rRINGR6Dmhu6bGhZonuW0na4+5Wh86kg9c/YI7jSIIspRRkH+v7+
RXH51h8Rbc2Tiv64qy7cIJIH0Bk0lFAaIQARAQABwsBfBBgBCAATBQJdQX8VCRAR
wx6OXgf00AIbDAAAgvAIAGyLaHYTjiXG/ORIIAgdQhKBYOozrOS4EcOPNdMSPBwT
6P+BpNS/GD33pzANVKM9Mft8+NnePUFOR2f4QJrQ1VvSj9ko8P2sX7IAe7+rG614
LQfzjG5R16KlSVOMFW0K2L8ZxumDdYl/N0BhgtZmB1lg1xY2TPHoDetznMnHG8sL
6u6vyhGl5a6qcW2g1urlF0VF/CEqg1lwAKhFHIFiNR+X6jCjg0KJa9MjAW6oICOx
oX0jp195mWix6suRJSWVK14uieT6uL5yYC5tZMz+t9rs7YxCkHxFRT1H5ZLHUD/r
93liqW+pzUx+bVdz5qNMb0ZonHZRLe3/Fzb19x8UMPc=
=6gp8
-----END PGP PUBLIC KEY BLOCK-----
`
	const sig = `-----BEGIN PGP SIGNATURE-----
Version: GopenPGP 2.4.10
Comment: https://gopenpgp.org

wsBzBAABCgAnBQJjocOZCZARwx6OXgf00BYhBCDPNjtY7JnnIuU+xBHDHo5eB/TQ
AACGgwf7Bx6J7JLZ2G6RFvr/wtl0DENZxUVS4H3wZPEIuVTh3/Lzd5BHfWN/mD+q
Sz0BcjRNxAI+nDY2/J8HPIibNg1NDlUgrgxK0NPLS1DMWmtoW3JTF5sfFMyiVGxo
RH4oluOe/UQcfxYTbMr8/EX8Gc9kdx4U7MqQNEc9CM5VIuxrfMpSZ2hvn5zlwexQ
WdnWjVWePpbwpltX98wTlAtU93XARUgeIMrzkhEBc1sNSg6/ynECLENm8EMxWQmj
9lpaROb2Fw50G7S1YjSUlc7WK+e4+IIP3Fqw/b21Kd1BasHS92OuHZNalbxyJA0F
V6Zkmvzj3h9CucLSJw1Bo6ZJTDbkBQ==
=fVs7
-----END PGP SIGNATURE-----
`
	message := "This is a test\nWith trailing spaces:    \n  With leading spaces\nWith trailing tabs:\t\t\n\tWith leading tabs\nWith trailing carriage returns:\r\n\rWith leading carriage returns\n\t \r With a mix \t\r\n"
	message = internal.TrimEachLine(message)
	key, err := NewKeyFromArmored(ricarda_key_armor)
	if err != nil {
		t.Error("error: ", err)
	}
	pgp := PGP()
	verifier, err := pgp.Verify().VerificationKey(key).Utf8().New()
	if err != nil {
		t.Error("error: ", err)
	}
	res, err := verifier.VerifyDetached([]byte(message), []byte(sig), Armor)
	if err != nil {
		t.Error("error: ", err)
	}
	if err = res.SignatureError(); err != nil {
		t.Error("Should have signature error", err)
	}
}
