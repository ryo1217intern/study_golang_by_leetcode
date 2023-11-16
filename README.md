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
