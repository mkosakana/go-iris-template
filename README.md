# 🦸🏼‍♂️ go-iris-sample


[Iris](https://github.com/kataras/iris) （Go言語フレームワーク）の使い方の参考用サンプルプロジェクト

<br>

📢 プロジェクト内で使用されている、

 - 変数
 - メソッド
 - ディレクトリ構造
 - ファイル
 - DB構造
 - その他諸々...

について、ご自由に変更してお使いください。

<br>

## 🎟 インストール

 1. [Go](https://go.dev) をインストールする
 2. 本プロジェクトをクローンする

```shell
$ cd $WORK_DIR
$ git clone https://github.com/mkosakana/go-iris-sample.git
$ cd go-iris-sample
```

<br>

## 🐲 使用方法

### ・`/_example-basic-api`

```shell
    go-iris-sample $ cd _example-basic-api
_example-basic-api $ go mod tidy -compat=1.17 && go mod download
_example-basic-api $ go run main.go
```

### ・`/_example-basic-view`

```shell
     go-iris-sample $ cd _example-basic-views
_example-basic-views $ go mod tidy -compat=1.17 && go mod download
_example-basic-views $ go run main.go
```

### ・`/_example-mvc-api`

```shell
  go-iris-sample $ _example-mvc-api
_example-mvc-api $ make build
_example-mvc-api $ make up
```

<br>

> ⚠️ Dockerが立ち上がったら、`docker-compose.yml` や `.env` などの内容を参考に、既存のエンドポイントに必要な「users」テーブルを作成してください。 
>    カラムは、「id（primary_key）」「name」「age」があれば問題ありません。
