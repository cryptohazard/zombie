#!/bin/bash

# Main code from  https://www.reddit.com/r/Monero/comments/6e6k53/brute_forcing_a_monero_wallet/

./zombie -print -f $1 | for word in $(</dev/stdin)  ;do echo "Trying: $word"; echo -e "$2\n$word" | ./monero-wallet-cli | egrep "^Opened wallet" && break; done
