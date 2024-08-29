// STEP06: レア度ごとに出る確率を変えてみよう

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	// 0から99までの間で乱数を生成する
	num := rand.Intn(100)

	// TODO: 変数numが0〜79のときは"ノーマル"、
	if 0 <= num && num <= 79 {
		fmt.Println(num)
		fmt.Println("N")
	} else if 80 <= num && num <= 94 {
		fmt.Println(num)
		fmt.Println("R")
	} else if 95 <= num && num <= 98 {
		fmt.Println(num)
		fmt.Println("SR")
	} else {
		fmt.Println(num)
		fmt.Println("XR")
	}

	// 80〜94のときは"R"、95〜98のときは"SR"、
	// それ以外のときは"XR"と表示する
}
