package main

import (
	"encoding/hex"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
)





func sample1() {
	fmt.Printf("sample1\n")
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	msg := []byte("abc")
	pub := sec.GetPublicKey()
	sig := sec.SignByte(msg)
	fmt.Printf("verify=%v\n", sig.VerifyByte(pub, msg))
}

func sample2() {
	fmt.Printf("sample2\n")
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	fmt.Printf("sec:%s\n", sec.SerializeToHexStr())
	pub := sec.GetPublicKey()
	fmt.Printf("1.pub:%s\n", pub.SerializeToHexStr())
	fmt.Printf("1.pub x=%x\n", pub)
	var P *bls.G1 = bls.CastFromPublicKey(pub)
	bls.G1Normalize(P, P)
	fmt.Printf("2.pub:%s\n", pub.SerializeToHexStr())
	fmt.Printf("2.pub x=%x\n", pub)
	fmt.Printf("P.X=%x\n", P.X.Serialize())
	fmt.Printf("P.Y=%x\n", P.Y.Serialize())
	fmt.Printf("P.Z=%x\n", P.Z.Serialize())
}

func sample3() {
	fmt.Printf("sample3\n")
	var sec bls.SecretKey
	b := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 0xff
	}
	err := sec.SetLittleEndianMod(b)
	if err != nil {
		fmt.Printf("err")
		return
	}
	fmt.Printf("sec=%x\n", sec.Serialize())
}

func sample4() {
	fmt.Printf("sample4\n")
	var sec bls.SecretKey
	secByte, _ := hex.DecodeString("4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d")
	sec.Deserialize(secByte)
	msg := []byte("message to be signed.")
	fmt.Printf("sec:%x\n", sec.Serialize())
	pub := sec.GetPublicKey()
	fmt.Printf("pub:%x\n", pub.Serialize())
	sig := sec.SignByte(msg)
	fmt.Printf("sig:%x\n", sig.Serialize())
}

func sample5() {
	fmt.Println("sample5\n")
	var sec bls.SecretKey
	var sec2 bls.SecretKey
	var sec3 bls.SecretKey
	var sec4 bls.SecretKey
	// todo
	//secByte, _ := hex.DecodeString("4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d")
	secByte, _ := hex.DecodeString("4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d")
	secByte2, _ := hex.DecodeString("4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d")
	secByte3, _ := hex.DecodeString("4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d")
	secByte4, _ := hex.DecodeString("4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d")

	sec.Deserialize(secByte)
	sec2.Deserialize(secByte2)
	sec3.Deserialize(secByte3)
	sec4.Deserialize(secByte4)

	// todo
	msg := []byte("hello!!!")
	msg2 := []byte("hello2!!!")
	msg3 := []byte("hello3!!!")
	msg4 := []byte("hello4!!!")

	pub := sec.GetPublicKey()
	pub2 := sec2.GetPublicKey()
	pub3 := sec3.GetPublicKey()
	pub4 := sec4.GetPublicKey()

	

	// sign
	sig := sec.SignByte(msg)
	sig2 := sec.SignByte(msg2)
	sig3 := sec.SignByte(msg3)
	sig4 := sec.SignByte(msg4)
	fmt.Printf("sig:%x\n", sig.Serialize())

	// verify
	fmt.Println(sig.VerifyByte(pub, msg))
	fmt.Println(sig2.VerifyByte(pub2, msg2))
	fmt.Println(sig3.VerifyByte(pub3, msg3))
	fmt.Println(sig4.VerifyByte(pub4, msg4))

	// create aggregate signature 
	sigs := []bls.Sign{*sig, *sig2, *sig3, *sig4}
	var aggSig bls.Sign
	aggSig.Aggregate(sigs)
	fmt.Println(aggSig)

	// verify aggregate signature
	// todo
	// pubkeys := []bls.PublicKey{*pub, *pub2, *pub3, *pub4}
	// msgs := {secByte, secByte2,secByte3, secByte4}
	// fmt.Println(aggSig.AggregateVerify(pubkeys, msgs))

	
	
	
}

func main() {
	bls.Init(bls.BLS12_381)
	bls.SetETHmode(bls.EthModeDraft07)
	sample1()
	sample2()
	sample3()
	sample4()

	sample5()
}
