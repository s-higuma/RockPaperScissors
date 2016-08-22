package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Rock = iota
	Scissors
	Paper
)

const (
	Draw = iota
	Lose
	Win
)

func main() {
	var playerCount, rivalCount, drawCount int

	fmt.Println("じゃんけんを行います。")

	for playerCount < 3 && rivalCount < 3 {
		fmt.Printf("「ぐう」は%d、「ちょき」は%d、「ぱー」は%dを入力してください。\n", 0, 1, 2)

		var playerResult int
		fmt.Scan(&playerResult)

		rand.Seed(time.Now().UnixNano())
		rivalResult := rand.Intn(2)

		result := Judge(playerResult, rivalResult)

		switch result {
		case Win:
			fmt.Printf("じゃんけんぽん！ 結果：勝ち　")
			playerCount++
		case Lose:
			fmt.Printf("じゃんけんぽん！ 結果：負け　")
			rivalCount++
		case Draw:
			fmt.Printf("じゃんけんぽん！ 結果：あいこ　")
			drawCount++
		}
		fmt.Printf("自分：%s　相手：%s\n", referHandByNum(playerResult), referHandByNum(rivalResult))
	}
	fmt.Printf("結果：　%d勝%d敗%d分\n", playerCount, rivalCount, drawCount)
}

func Judge(playerResult, rivalResult int) (result int) {
	result = (playerResult - rivalResult + 3) % 3
	return
}

func referHandByNum(n int) (hand string) {
	switch n {
	case Rock:
		hand = "ぐう"
	case Scissors:
		hand = "ちょき"
	case Paper:
		hand = "ぱー"
	}
	return
}
