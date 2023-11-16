# Golang 備忘録

## nil ってなに？

_A.型定義を保存できる null._

null は何もかも未定義の状態.

それに対して nil は値は未定義であるが型は宣言されている状態.

## func funcName(foo) bar{...における foo と bar は何か？

_A. foo -> 引数、 bar -> 返り値_

## make(map[int]int)とは何か？

_辞書型配列の作成を行う_

map[foo]bar

foo によって要素の型を定義し、bar によって index の型を宣言する.

```golang
language := make(map[string]string)

langeage["japan"] = "japanese"
language["america"] = "english"

fmt.Println(language)//map["japan":"japanese", "america":"english"]
fmt.Println(language["japan"])//"japanese"

```

## for range とは何か？

_A. index の数だけ for 文を回すことができる_

以下の例では`i`に index の値,`v`には配列の要素を代入する.

```golang
for i, v = range [int]{1,2,3} {
fmt.Println(i, v)
//0, 1
//1, 2
//2, 3
}
```

また注意点として map 配列の場合実行される順序がランダムになってしまうという点です.

```Golang
package main

import "fmt"

func main() {
	m := map[int]string{
		1: "value1",
		2: "value2",
		3: "value3",
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
```

この時出力される値はランダムになってしまいます。

> 古い言語仕様では、マップの反復順序が定義しておらず、ハードウェアプラットフォーム間で実行順が異なっていた。反復順序に依存すると移植性がなくなるため、ランダム性を持たせている。というのがこの挙動の理由みたいです。

[公式 Document](https://go.dev/doc/go1#iteration)

順番に実行する際には以下のようなコードで実装を行う.

```Golang
package main

import (
	"fmt"
	"sort"
)

func main() {

  //map配列の宣言
	m := map[int]string{
		1: "value1",
		2: "value2",
		3: "value3",
	}

  //mのインデックス番号用の配列
	var keys []int

  //mのインデックス番号を格納（この際、格納される番号はランダム）
	for k := range m {
		keys = append(keys, k)
	}

  //格納されたランダムなindex番号をソート
	sort.Ints(keys)

  //map配列ではないkeysによってfor文を回す
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
```
