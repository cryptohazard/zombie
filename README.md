Zombie
=============

This is a collection of tools to explore security on Blockchains. We focus specifically on cryptography, notably random number generation, private key, wallet cracking, brainwallet etc.

Disclaimer
---------------
Just because you can steal someone else money does not mean you should do it.

Inspiration(great and funny talk at Defcon): https://github.com/ryancdotorg/brainflayer

## Installation
You need to have a working go(lang) environment in version >=1.9 and clone this repository. I can provide executable if there are requests.

You first need to get the dependency:
```
$ go get github.com/btcsuite/btcutil/base58
```

Now you can go in the ```exec``` folder and build the executable:
```
$ cd exec/
$ go build -zombie.go
$ ./zombie
```
##  Usage
```
$ ./zombie -h
Usage of ./zombie:
  -f string
    	format file (default "guesses")
  -print
    	Print the generated canditates to standard output
  -wif
    	WIF cracker. Generate valid wif from the format or wordlist file
```

## Wordlist generation

### Goal
The goal is to generate password/keys when you know part of the target. This is not suitable if you have a high number of candidates for some parts.
Hopefully next versions will support candidates from file generated with the tools like [crunch](https://tools.kali.org/password-attacks/crunch) or [john the ripper](https://tools.kali.org/password-attacks/john).

### Format
First you need to fill the format file(see ```exec/format_example``` for an example):
```
// read the format line and put it in an array of candidates
// one time symbol between parts candidates
// first symbol is a delimiter follow by a part
// repeat the delimiter before each part
// ex: aEa3  	=> [E 3]
// ex: !g!d!e!p => [g d e p]
// ex: %OUI%NON%YES%NO => [OUI NON YES NO]
```
A useful feature that is missing is being able to set all the characters of the alphabet you are consider, using ```?```. The alphabets considered for addition are [base58](https://en.bitcoin.it/wiki/Base58Check_encoding) and [BIP39 wordlists](https://github.com/bitcoin/bips/blob/master/bip-0039/bip-0039-wordlists.md) mainly, and maybe hexadecimal/binary in case we want to play directly on bytes/bits level.

### Usage
The philosophy of this function is to use it with the cracking functions available (not much for now) or to print the resulting candidates, with the option ```-print```.
### Example: Monero Wallet Cracker
Let us assume you forgot *exactly* your Monero wallet password but you know the parts in it. You can use ```zombie``` to generate the password candidates and then pipe it to the wallet. I made a Bash script to use where your wallet and the Monero cli are located:
```
./monero_cracker.sh format_file brute_forcing_a_monero_wallet
```

## Wif Cracker: guess a private keys
This function is useful to take on  [contest like this one](http://jangodfrey.com/illustration/guess-my-bitcoin/guess-my-bitcoin.php):
![img](http://jangodfrey.com/illustration/guess-my-bitcoin/GUESS_MY_BITCOIN_tiled.png)
It will generate the candidates from the format, check if they are valid [wif](http://learnmeabitcoin.com/glossary/wif) and print the valid wif. You need to then derive them and see if they hit your target. Unfortunately, we did not get the 1/2 BTC :-( because we had a wrong assumption and did not consider the case color.

*Remember*:You need to have partial knowledge of the key, somehow, and hope that you can bruteforce the remaining space.

The solution to the contest: [5JKPapJwgyEij3sxYRAEnixyiFgxqkVhgZXv9bWWknBexegx6tM](https://twitter.com/guessmybitcoin/status/781887409394974720)
You can test it with the provided ```guesses``` file.

## Next features
* Number of candidates generated and size
* Accept data generated from crunch/John
* Add Public key derivation for BTC, ETH, ...
* Brainwallet cracker
* Mnemonic phrases support(BIP39,...)
* Try very crappy random number generator: date, hour, low/high Hamming weight (00..00, 11..11)
* Advance ECC Discrete logarithm attacks
