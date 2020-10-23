package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

func exitError(message string, code int) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(code)
}

func ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}

func main() {
	lang := flag.String("lang", "english", "Select mnemonic language (english chinese_simplified chinese_traditional french italian japanese, korean, spanish)")
	list := flag.Bool("list", false, "List avaliable words")
	build := flag.Bool("build", false, "Build a mnemonic")
	flag.Parse()
	switch *lang {
	case "english":
	case "chinese_simplified":
		bip39.SetWordList(wordlists.ChineseSimplified)
	case "chinese_traditional":
		bip39.SetWordList(wordlists.ChineseTraditional)
	case "french":
		bip39.SetWordList(wordlists.French)
	case "italian":
		bip39.SetWordList(wordlists.Italian)
	case "japanese":
		bip39.SetWordList(wordlists.Japanese)
	case "korean":
		bip39.SetWordList(wordlists.Korean)
	case "spanish":
		bip39.SetWordList(wordlists.Spanish)
	default:
		exitError("supported languages: english chinese_simplified chinese_traditional french italian japanese korean spanish", 3)
	}
	if *list {
		fmt.Println(strings.Join(bip39.GetWordList(), " "))
	}
	if *build {
		doBuild()
	}
	if !*list && !*build {
		flag.Usage()
	}
}

func doBuild() {
	fmt.Println(
		`To build your mnemonic, you can hand pick first 23 words of the mnemonic.
List word can be chosen from shorter list, to match correct checksum of your mnemonic.`)
	wordlist := bip39.GetWordList()
	words := make([]string, 23)
	prompt := promptui.Select{
		Items: wordlist,
		Size:  30,
	}
	for i := range words {
		prompt.Label = fmt.Sprintf("Select %s word", ordinal(i+1))
		_, word, err := prompt.Run()
		if err != nil {
			exitError(err.Error(), 1)
		}
		words[i] = word
	}
	first23 := strings.Join(words, " ")
	items := make([]string, 0, 8)
	for _, word := range wordlist {
		_, err := bip39.MnemonicToByteArray(first23 + " " + word)
		if err == nil {
			items = append(items, word)
		}
	}
	prompt.Label = "Select last word"
	prompt.Items = items
	_, word, err := prompt.Run()
	if err != nil {
		exitError(err.Error(), 2)
	}
	words = append(words, word)
	fmt.Println("Your mnemonic:")
	fmt.Println(strings.Join(words, " "))
}
