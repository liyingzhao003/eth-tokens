package main

import (
	"flag"
	"fmt"
	"github.com/eager7/eth_tokens/script/built"
	"github.com/eager7/eth_tokens/script/coin_gecko"
	"github.com/eager7/eth_tokens/script/ether_scan"
)

func main() {
	var g = flag.Bool("g", false, "")
	var e = flag.Bool("e", false, "")
	flag.Parse()
	if *g {
		tokenList, err := built.TokenListFromGit(built.URLTokenList)
		if err != nil {
			panic(err)
		}
		fmt.Println("success get token list:", len(tokenList))
		if err := built.InitializeTokens(`../../tokens`, tokenList); err != nil {
			panic(err)
		}
	}
	if *e {
		spider, err := ether_scan.Initialize("http://47.52.157.31:8585")
		if err != nil {
			panic(err)
		}
		tokenList, err := spider.BuiltTokensFromEtherScan()
		if err != nil {
			panic(err)
		}
		for index := range tokenList {
			coin_gecko.ReplaceTokenLogoFromCoinGecko(&tokenList[index])
		}
		fmt.Println("success get token list:", len(tokenList))
		for _, token := range tokenList {
			coin_gecko.ReplaceTokenLogoFromCoinGecko(token)
		}
		tokenList = append(tokenList, &built.TokenInfo{
			Symbol:   "ETH",
			Name:     "Ethereum",
			Type:     "Main Coin",
			Address:  "0x0000000000000000000000000000000000000000",
			Decimals: 18,
			Logo: built.Logo{
				Src: "https://www.cryptocompare.com/media/20646/eth_logo.png?width=200",
			},
		})
		if err := built.InitializeTokens(`../../tokens`, tokenList); err != nil {
			panic(err)
		}
	}

	tokens, err := built.CollectTokens(`../../tokens`)
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
	for i, t := range tokens {
		if t.Contract == "0x0000000000000000000000000000000000000000" {
			tokens = append(tokens[:i], tokens[i+1:]...)
		}
	}
	fmt.Println("len tokens:", len(tokens))
=======

>>>>>>> a9adde2b2ccc50d30799e6486a54168657ba98a4
	eth := built.Token{
		Name:     "Ethereum",
		Symbol:   "ETH",
		Contract: "0x0000000000000000000000000000000000000000",
		Decimals: 18,
		Logo:     "https://www.cryptocompare.com/media/20646/eth_logo.png?width=200",
		Invalid:  true,
	}
	tokens = append([]*built.Token{&eth}, tokens...)
	if err := built.BuildDist(`../../dist`, tokens); err != nil {
		panic(err)
	}
	if err := built.BuildReadme(`../../tokens.md`, tokens); err != nil {
		panic(err)
	}
}
