package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	member, payment, ave := useBufioScanner("sample")

	// 平均との差額を計算
	for name, price := range payment {
		payment[name] = price - ave
	}

	fmt.Println(member)
	fmt.Println(payment)
	fmt.Println(ave)
}

//func calculation(payment []string, liquidation []string) (payment, liquidation) {
//	// 現在の最大債務者と最大債権者を取得
//	creditor = payment[]
//	debtor = payment[-1]
//}

const MaxInt = int(^uint(0) >> 1)

var sc = bufio.NewScanner(os.Stdin)

func useBufioScanner(fileName string) ([]string, map[string]int, int) {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {
			panic("Failed : file can't open")
		}
	}(fp)

	scanner := bufio.NewScanner(fp)
	// 1行目を読み込む
	scanner.Scan()
	// メンバーの配列と、各メンバーがどれだけ支払っているかのmap
	member := make([]string, 0)
	payment := make(map[string]int)
	payment2 := make([]map[string]int, 0)
	sum := 0

	text1 := strings.Split(scanner.Text(), " ")
	for _, s := range text1 {
		member = append(member, s)
		payment[s] = 0 // 初期値0
	}

	// 2行目以降の支払いを計算する
	for scanner.Scan() {
		text2 := strings.Split(scanner.Text(), " ")
		payment[text2[0]] += s2i(text2[1])
		sum += s2i(text2[1])
	}
	for x, y := range payment {
		payment2 = append(payment2, map[string]int{x: y})
	}

	return member, payment, sum / len(member)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readSpaceStringList() []string {
	sc.Scan()
	stringList := make([]string, 0)
	text := strings.Split(sc.Text(), " ")
	for _, s := range text {
		stringList = append(stringList, s)
	}
	return stringList
}

func readInt1() int {
	sc.Scan()
	return s2i(sc.Text())
}

func maxOfInts(a []int) int {
	res := -MaxInt
	for _, v := range a {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minOfInts(a []int) int {
	res := MaxInt
	for _, v := range a {
		res = min(res, v)
	}
	return res
}

// String -> Int
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Failed : " + s + " can't convert to int")
	}
	return v
}
