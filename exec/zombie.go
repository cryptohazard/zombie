package main

import (
	"flag"
	"zombie"
)

func main() {
	file := flag.String("f", "guesses", "format file")
	//wordlist := flag.String("w", "wordlist", "wordlist file")
	//alphabet:= flag.String("a","","alphabet of symbols to consider. Ex: base58)
	print := flag.Bool("print", false, "Print the generated canditates to standard output")
	guessWif := flag.Bool("wif", false, "WIF cracker. Generate valid wif from the format or wordlist file")

	flag.Parse()

	out, wg := zombie.Parse(*file)

	if *print {
		zombie.Print(out, wg)
		return
	}
	if *guessWif {
		zombie.CrackWif(10, out, wg)
		return
	}
	zombie.Parse(*file)
}
