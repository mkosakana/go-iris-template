# 🦸🏼‍♂️ go-iris-sample
[Iris](https://github.com/kataras/iris) （Goフレームワーク）の使うための参考用テンプレートリポジトリ  


### 📢 announce
あくまでIrisの使い方を"参考"するためのリポジトリであり，ディレクトリ内で使用されている

- 変数
- メソッド
- ディレクトリ構造
- ファイル
- その他諸々

などについて，ご自由に変更してお使いください．


## 🎟 Install

### Route1: Install with `Docker`

```shell
// docker-composeからイメージのビルド
1. make build

// コンテナーをバックグラウンドで立ち上げ
2. make up

// mod ファイルのダウンロード
3. go mod download
```


### Route2: Install with `go.mod`

1. Download [golang](https://go.dev).  

2. Clone this repository on your working directory.
```shell
$ cd $WORK_DIR
$ git clone https://github.com/mkosakana/go-iris-sample.git
$ go mod download
```


## 🦉 Get Start

```shell
// ホームディレクトリから
go-iris-sample $ go run main.go
```

## 👾 TODO
 - DBとのやりとり

