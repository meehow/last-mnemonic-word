# Last Mnemonic Word (AKA unsafe gift-wallet generator)

This CLI program implements my weird idea of generating gift cryptocurrency wallet from hand picked words.

Words are BIP-39 compatible. First 23 words can be selected from list of 2048 words
(in english, chinese, french, italian, japanese, korean or spanish),
but last word contains a checksum, so program will ask you to choose it from short list of 8 words.

## Warning

This wallet will be way less secure than wallets generated from random values.
Please don't store large sums on it and/or explain gifted person that funds should be transferred to a wallet with randomly generated seeds.

## Demo

[![asciicast](https://asciinema.org/a/367340.svg)](https://asciinema.org/a/367340)
