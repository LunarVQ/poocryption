// Command poocrypt is a thin CLI over the poocryption library.
//
//	poocrypt keygen [-bits N] -out NAME            # writes NAME.pub and NAME.key as emoji strings
//	poocrypt enc    -pub  FILE (-s STRING | -in FILE) [-out FILE]
//	poocrypt dec    -priv FILE (-s POOP   | -in FILE) [-out FILE]
//	poocrypt walk   -pub  FILE -x N -t STEPS        # demo: step x around the cycle t times
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "keygen":
		err = runKeygen(os.Args[2:])
	case "enc":
		err = runEnc(os.Args[2:])
	case "dec":
		err = runDec(os.Args[2:])
	case "walk":
		err = runWalk(os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "poocrypt:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `usage:
  poocrypt keygen [-bits N] -out NAME
  poocrypt enc    -pub  FILE (-s STRING | -in FILE) [-out FILE]
  poocrypt dec    -priv FILE (-s POOP   | -in FILE) [-out FILE]
  poocrypt walk   -pub  FILE -x N -t STEPS`)
}

func runKeygen(args []string) error {
	// TODO: poocryption.GenerateKey(bits); write emojiset.Encode(priv.MarshalPrivate())
	// to NAME.key (0600) and emojiset.Encode(pub.MarshalPublic()) to NAME.pub
	return fmt.Errorf("keygen not implemented")
}

func runEnc(args []string) error {
	// TODO: load key file, emojiset.Decode -> UnmarshalPublic; read -s or -in;
	// poocryption.Encrypt; write output
	return fmt.Errorf("enc not implemented")
}

func runDec(args []string) error {
	// TODO: load key file, emojiset.Decode -> UnmarshalPrivate; read input;
	// poocryption.Decrypt; write output
	return fmt.Errorf("dec not implemented")
}

func runWalk(args []string) error {
	// TODO: load public key; print pub.Step applied t times from x, one value per line,
	// so the cycle (…and its return to x) can be seen on small demo keys.
	return fmt.Errorf("walk not implemented")
}
