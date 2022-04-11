# 🦸🏼‍♂️ go-iris-sample
[Iris](https://github.com/kataras/iris) （Goフレームワーク）の使うための参考用テンプレートリポジトリ  
（※ 本来，APIやWebアプリケーションの開発にIrisを使用する際は，DBと接続してデータを取得・更新...などを行う機能がつくと思いますが，今回は割愛しています）


### 📢 announce
あくまでIrisの使い方を"参考"するためのリポジトリであり，
ディレクトリ内で使用されている
- 変数
- メソッド
- ディレクトリ構造
- ファイル
- その他諸々

などについて，ご自由に変えてお使いください．


## 🎟 Install

```go
// docker-composeからイメージのビルド
1. make build

// コンテナーをバックグラウンドで立ち上げ
2. make up

// mod ファイルのダウンロード
3. go mod download
```

## 🦉 Get Start

```go
// 使うディレクトリまで移動
cd _example-...

go run main.go
```
