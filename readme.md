[![Build Status](https://travis-ci.org/herumi/bls-eth-go-binary.png)](https://travis-ci.org/herumi/bls-eth-go-binary)
# bls for eth with compiled static library

This repository contains compiled static library of https://github.com/herumi/bls with `BLS_ETH=1`.

* SecretKey; Fr
* PublicKey; G1
* Sign; G2

# News
- 2020/Jul/03 remove old tests and use the latest hash function defined at [draft-07](https://www.ietf.org/id/draft-irtf-cfrg-hash-to-curve-07.txt) is set by default.
- 2020/May/22 `SignHashWithDomain`, `VerifyHashWithDomain`, `VerifyAggregateHashWithDomain` are removed.
- 2020/May/15 `EthModeDraft07` is added for [draft-07](https://www.ietf.org/id/draft-irtf-cfrg-hash-to-curve-07.txt).
- 2020/Apr/20 `EthModeDraft06` is default. Call `SetETHmode(EthModeDraft05)` to use older evrsion.
- 2020/Mar/26 The signature value in `SetETHmode(2)` has changed because of changing DST in hash-to-curve function.
- 2020/Mar/17 This library supports [eth2.0 functions](https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/beacon-chain.md#bls-signatures). But the spec of hash-to-curve function may be changed.

Init as the followings:

```
Init(BLS12_381)
```

then, you can use the following functions.

bls-eth-go-binary | eth2.0 spec name|
------|-----------------|
SecretKey::SignByte|Sign|
PublicKey::VerifyByte|Verify|
Sign::Aggregate|Aggregate|
Sign::FastAggregateVerify|FastAggregateVerify|
Sign::AggregateVerifyNoCheck|AggregateVerify|

The size of message must be 32 byte.

Check functions:
- VerifySignatureOrder ; make `deserialize` check the correctness of the order
- Sign::IsValidOrder ; check the correctness of the order
- VerifyPublicKeyOrder ; make `deserialize` check the correctness of the order
- PublicKey::IsValidOrder ; check the correctness of the order
- AreAllMsgDifferent ; check that all messages are different each other

# How to run `examples/sample.go`

```
go get github.com/herumi/bls-eth-go-binary/
go run examples/sample.go
```

# How to build the static library
The following steps are not necessary if you use compiled binary in this repository.

## Linux, Mac, Windows(mingw64)
```
mkdir work
cd work
git clone https://github.com/herumi/mcl
git clone https://github.com/herumi/bls
git clone https://github.com/herumi/bls-eth-go-binary
cd bls-eth-go-binary
make CXX=clang++
go test ./bls -bench "Pairing|Sign|Verify"
```

clang generates better binary than gcc.

## Android
```
make android
```

## iOS
```
make ios
```

## How to cross-compile for armeabi-v7a
Check `llc --version` shows to support the target armv7a.

```
make ../mcl/src/base32.ll
env CXX=clang++ BIT=32 ARCH=arm _OS=android _ARCH=armeabi-v7a make MCL_USE_GMP=0 UNIT=4 CFLAGS_USER="-target armv7a-linux-eabi -fPIC"
```

* Remark : clang++ must support `armv7a-linux-eabi`.
### How to build sample.go for armeabi-v7a

```
env CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARM=7 GOARCH=arm go build examples/sample.go
env QEMU_LD_PREFIX=/usr/arm-linux-gnueabi qemu-arm ./sample
```

## How to cross-compile for mipsel
Check `llc --version` shows to support the target mips.

```
make ../mcl/src/base32.ll
env CXX=clang++ BIT=32 ARCH=mipsel _OS=linux _ARCH=mipsle make MCL_USE_GMP=0 UNIT=4 CFLAGS_USER="-target mipsel-linux -fPIC"
```

* Remark : clang++ must support `mipsel-linux`.
### How to build sample.go for mipsel

```
env CC=mipsel-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build examples/sample.go
env QEMU_LD_PREFIX=/usr/mipsel-linux-gnu qemu-mipsel ./sample
```

# How to use the static library from C
```
#define BLS_ETH
#include <mcl/bn_c384_256.h>
#include <bls/bls.h>
```

# Author
MITSUNARI Shigeo(herumi@nifty.com)

# Sponsors welcome
[GitHub Sponsor](https://github.com/sponsors/herumi)
