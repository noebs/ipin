package ipin

import (
	"flag"
	"fmt"

	// "github.com/adonese/crypto"
	"github.com/google/uuid"
)

var uid = uuid.New().String()

func main() {
	key := flag.String("key", "", "public key from ebs")
	ipin := flag.String("ipin", "0000", "ipin you want to create its pin block")
	uid1 := flag.String("uuid", "", "uuid for transaction")
	flag.Parse()

	if uid1 != nil {
		fmt.Print(Encrypt(*key, *ipin, *uid1))
	} else {
		fmt.Print(Encrypt(*key, *ipin, uid))
	}

}
