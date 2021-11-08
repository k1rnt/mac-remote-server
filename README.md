# mac-remote-server
[![Actions Status](https://github.com/k1rnt/mac-remote-server/workflows/linter/badge.svg)](https://github.com/k1rnt/mac-remote-server/actions)
[![Actions Status](https://github.com/k1rnt/mac-remote-server/workflows/test/badge.svg)](https://github.com/k1rnt/mac-remote-server/actions)

## 概要
- 「Hey Siri macのロック解除して」
- 「Hey Siri macをスリープして」
- 「Hey Siri macの〇〇を教えて」

ができる奴

## 選定技術
- go
  - [echo](github.com/labstack/echo/v4)
  - [mack](github.com/andybrewer/mack)
- AppleScript

## TODO
第二フェーズとしてWebクライアントと管理画面の画面を作りたい

## 使い方
1. サーバー立ち上げ
```cmd
$go build -o mac-remote-server main.go
$./mac-remote-server
```
2. iOSのショートカットアプリで上記サーバーで立ち上げた  
エンドポイント(同一LAN内のmacのIPアドレス)に対応するリクエストを送る物を設定する。
4. Siriに2で作成したショートカットを登録する  
※ショートカット名を"macをスリープ"など好きに設定する
6. 「Hey Siri macをスリープ！」
