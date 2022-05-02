# 🦸🏼‍♂️ go-iris-sample

[Iris](https://github.com/kataras/iris) （Goフレームワーク）の使い方参考用テンプレートリポジトリ  


### 📢 announce

あくまでIrisの使い方を"参考"するためのリポジトリであり，ディレクトリ内で使用されている

- 変数
- メソッド
- ディレクトリ構造
- ファイル
- その他諸々...

について，ご自由に変更してお使いください．


## 🎟 Install

1. Download [golang](https://go.dev).

2. Clone this repository on your working directory.

```shell
$ cd $WORK_DIR
$ git clone https://github.com/mkosakana/go-iris-sample.git
$ cd go-iris-sample
$ make install
```


## 🐲 Get Start

### Route1: Without Docker
only be able to use `/_example-basic-api` or `/_example-basic-view`, because they have NO Data Base connection

```shell
go-iris-sample $ cd _example-basic-api/ or _example-basic-view/
$ go run main.go
```


### Route2: With Docker (Data Base connection ✨）

1. `.air.toml` ファイル内 `cmd = "go build -o ./tmp/main ./_example-mvc-api/main.go"` の，
`"./_example-mvc-api/main.go"` で指定されたディレクトリがコンテナとして立ち上がる（デフォルトでは「_example-mvc-api」）ため，**「_example-mvc-api」** の箇所を Docker を使って作業したいディレクトリ名に置換する．

4. start Docker.

```shell
$ cd go-iris-sample

// 1. docker-composeからイメージのビルド
$ make build

// 2. コンテナー・DBをバックグラウンドで立ち上げ
$ make up
```
