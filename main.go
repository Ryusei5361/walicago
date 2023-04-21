package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type liquidation struct {
	creditor string // 債権者、お金をもらう人
	debtor   string // 債務者、お金を払う人
	amount   int    // 払う額
}

func main() {
	// 精算を記録
	var adjustment []liquidation
	// 差額を記録
	difference := make(map[string]int)

	member, originalPayment, actualPayment := readFile("sample")

	// 平均との差額を計算
	for _, name := range member {
		difference[name] = actualPayment[name] - originalPayment[name]
	}

	// 再帰処理
	difference, adjustment = calculation(difference, adjustment)

	// 残債
	remainDebt := 0
	for _, last := range difference {
		remainDebt += last
	}

	fmt.Println(difference)
	fmt.Println(adjustment)
	fmt.Println(remainDebt)
}

func calculation(payment map[string]int, adjustment []liquidation) (map[string]int, []liquidation) {

	// 現在の最大債務者と最大債権者を取得
	creditor, priceCreditor := maxOfInts(payment)
	debtor, priceDebtor := minOfInts(payment)

	// 最大債権者と最大債務者の差額
	amount := min(abs(priceDebtor), priceCreditor)

	if amount == 0 {
		return payment, adjustment
	}

	payment[creditor] -= amount
	payment[debtor] += amount
	adjustment = append(adjustment, liquidation{creditor: creditor, debtor: debtor, amount: amount})

	return calculation(payment, adjustment)
}

const MaxInt = int(^uint(0) >> 1)

var sc = bufio.NewScanner(os.Stdin)

func readFile(fileName string) ([]string, map[string]int, map[string]int) {
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
	// 各メンバーの本来払うべき金額
	originalPayment := make(map[string]int)
	// 各メンバーの実際に払っている金額
	actualPayment := make(map[string]int)

	text1 := strings.Split(scanner.Text(), " ")
	for _, s := range text1 {
		member = append(member, s)
		originalPayment[s] = 0 // 初期値0
		actualPayment[s] = 0   // 初期値0
	}

	// 支払いを計算する
	for scanner.Scan() {
		text2 := strings.Split(scanner.Text(), ":")
		creditor := text2[0]
		debtors := strings.Split(text2[1], " ")
		amount := s2i(text2[2])
		amountPerMember := amount / len(debtors)
		for _, name := range debtors {
			originalPayment[name] += amountPerMember
		}
		actualPayment[creditor] += amount
	}

	return member, originalPayment, actualPayment
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

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
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

func maxOfInts(a map[string]int) (string, int) {
	maxMember := ""
	maxPrice := -MaxInt
	for name, price := range a {
		if price > maxPrice {
			maxPrice = price
			maxMember = name
		}
	}
	return maxMember, maxPrice
}

func minOfInts(a map[string]int) (string, int) {
	minMember := ""
	minPrice := MaxInt
	for name, price := range a {
		if price < minPrice {
			minPrice = price
			minMember = name
		}
	}
	return minMember, minPrice
}

// String -> Int
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Failed : " + s + " can't convert to int")
	}
	return v
}
