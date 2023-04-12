# WalikaのアルゴリズムをGoで実装してみる
以下のサイトを参考にしました。
- https://qiita.com/MasashiHamaguchi/items/0348082984b8c94ca581

## 全体の流れ
1. 誰がいくら払ったかを入力する。
2. 入力された金額の合計を求め、人数で割り、１人あたりの支払額を求める。
3. 2で求めた金額と、実際に払った金額の差分を求め、メンバーと差分のmapを作成する。
4. 支払いが最も多い人から順に並び替える。
5. 支払いが最も多い人(金額をFとする)に、最も少ない人(金額をLとする)が支払いを行う。  
このとき、送金額は min(F, |L|)  になる。
6. 送金額が0になるまで4, 5を繰り返す。

### 入力形式
- map を使用。  
候補1
```
sampleData = {name1: payAmount, name2: payAmount, ...}
```
候補2
```
sampleData = [
{name: A, payAmount: 100},
{name: B, payAmount: 200},
{name: C, payAmount: 0},
        .
	.
	.
	,
] 
```
# walicago
