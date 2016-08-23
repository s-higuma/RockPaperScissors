package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 識別子iotaは、定数宣言文（const）内で使用される、連続する型未定の整数の定数を表します。
// この定数は予約語constが現れたときに0に初期化され、各定数定義の後に1ずつインクリメントされます。
const (
	Rock     = 1 + iota // Rock = 1
	Scissors            // Scissors = 2
	Paper               // Paper = 3
)

func referHand(n int) (hand string) {
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

const (
	Draw = iota // Draw = 0
	Lose        // Lose = 1
	Win         // Win = 2
)

func Judge(playerHand, rivalHand int) (result int) {
	// ex）ぐう(1)とぱー(3)の場合
	// ( 1 - 3 + 3) % 3 = 0 … 1 → Lose
	result = (playerHand - rivalHand + 3) % 3
	return
}

func main() {
	var playerWinCount, rivalWinCount, drawCount int

	fmt.Println("じゃんけんを行います。")

	for playerWinCount < 2 && rivalWinCount < 2 {
		fmt.Printf("「ぐう」は%d、「ちょき」は%d、「ぱー」は%dを入力してください。\n", 1, 2, 3)

		var playerHand int
		fmt.Scan(&playerHand)

		rand.Seed(time.Now().UnixNano())
		rivalHand := rand.Intn(3) + 1

		result := Judge(playerHand, rivalHand)

		switch result {
		case Win:
			fmt.Printf("じゃんけんぽん！ 結果：勝ち　")
			playerWinCount++
		case Lose:
			fmt.Printf("じゃんけんぽん！ 結果：負け　")
			rivalWinCount++
		case Draw:
			fmt.Printf("じゃんけんぽん！ 結果：あいこ　")
			drawCount++
		}
		fmt.Printf("自分：%s　相手：%s\n", referHand(playerHand), referHand(rivalHand))
	}
	fmt.Printf("結果：　%d勝%d敗%d分\n", playerWinCount, rivalWinCount, drawCount)
}
