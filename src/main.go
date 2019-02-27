package main

import (
	"fmt"
	"math/big"
	"syscall/js"

	abi "github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

func decodeHex(i []js.Value) {
	v := i[0].String()
	// n := new(big.Int)
	// var number []byte
	// inputNo, ok := n.SetString(v, 10)
	// if !ok {
	// 	fmt.Println("Unable to make string to uint256")
	// 	return
	// }
	number = abi.U256(v)

	fmt.Printf("%s", number)

}

// decodeTokenID
func decodeTokenID(i []js.Value) {
	v := i[0].String()
	n := new(big.Int)
	inputNo, ok := n.SetString(v, 10)
	if !ok {
		fmt.Println("Unable to make string to uint256")
		return
	}
	fmt.Println(inputNo)
	clearLowNo, ok := n.SetString("0xffffffffffffffffffffffffffffffff00000000000000000000000000000000", 10)
	if !ok {
		fmt.Println("Unable to make string to uint256")
		return
	}
	fmt.Println(clearLowNo)
	clearHighNo, ok := n.SetString("0x00000000000000000000000000000000ffffffffffffffffffffffffffffffff", 10)
	if !ok {
		fmt.Println("Unable to make string to uint256")
		return
	}
	fmt.Println(clearHighNo)
	// js.Global().Get("window").Call("fireEvent", no)
	return
}

// func expandNegative128BitCast(val) {

// }

// func add(i []js.Value) {
// 	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
// 	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
// }

func registerCallbacks() {
	js.Global().Set("decodeTokenId", js.NewCallback(decodeTokenID))
	js.Global().Set("decodeHex", js.NewCallback(decodeHex))
}

// UINT256 setup

// A Number represents a generic integer with a bounding function limiter. Limit is called after each operations
// to give "fake" bounded integers. New types of Number can be created through NewInitialiser returning a lambda
// with the new Initialiser.
type Number struct {
	num   *big.Int
	limit func(n *Number) *Number
}

// Uint256 Return a Number with a UNSIGNED limiter up to 256 bits
func Uint256(n int64) *Number {
	return &Number{big.NewInt(n), limitUnsigned256}
}

var tt256m1 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

func limitUnsigned256(x *Number) *Number {
	x.num.And(x.num, tt256m1)
	return x
}
