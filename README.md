<div align="center">

# PuyoDeliver 📦

<!-- s;HidemaruOwO/Puyodeliver;User/Repository;g -->

[![Test CLI](https://github.com/HidemaruOwO/Puyodeliver/actions/workflows/test.yml/badge.svg)](https://github.com/HidemaruOwO/Puyodeliver/actions/workflows/test.yml)
![最終コミット](https://img.shields.io/github/last-commit/HidemaruOwO/Puyodeliver?style=flat-square)
![リポジトリのスター](https://img.shields.io/github/stars/HidemaruOwO/Puyodeliver?style=flat-square)
![問題](https://img.shields.io/github/issues/HidemaruOwO/Puyodeliver?style=flat-square)
![オープンな問題](https://img.shields.io/github/issues-raw/HidemaruOwO/Puyodeliver?style=flat-square)
![バグの問題](https://img.shields.io/github/issues/HidemaruOwO/Puyodeliver/bug?style=flat-square)

<!-- ![image](https://github.com/HidemaruOwO/Puyodeliver/assets/82384920/bf4ccddf-3eae-4fae-97f4-d2b59bec919f) -->

## なんだこれは？

ウェブ上のぷよっとしたかわいいファイルブラウザーです。

</div>

- Select Language

<table>
  <thead>
    <tr>
      <th style="text-align:center"><a href="README.md">🎌日本語</a></th>
      <th style="text-align:center"><a href="README.en.md">🤡English</a></th>
      <th style="text-align:center"><a href="README.zh-CN.md">🐉简体中文</a></th>
      <th style="text-align:center"><a href="README.zh-TW.md">🍜繁体中文</a></th>
      <th style="text-align:center"><a href="README.ko.md">🌸한국어</a></th>
    </tr>
  </thead>
</table>

## Usage 💨

GitHubの[Release](https://github.com/HidemaruOwO/Puyodeliver/releases)もしくは、Actionsの[Artifact](https://github.com/HidemaruOwO/Puyodeliver/actions/workflows/build.yml)から実行ファイルをダウンロードして実行してください。

```bash
mkdir public
PUBLIC_FOLDER=public/ ./puyodeliver
```

<!-- ## Install 😊 -->

<!-- このスクリプトを実行してください。 -->

<!-- ```bash -->
<!-- ./install.sh -->
<!-- ``` -->

## Build 🔨

```bash
git clone https://github.com/HidemaruOwO/Puyodeliver.git
cd Puyodeliver
go build src/main.go
```

## Dependencies 🪡

このアプリを使用するには、以下のコマンドをパスに登録してください。

- `git`

## Repository Tools 🔧

- [ ] Depandabotのセットアップ
- [ ] CodeQLのセットアップ

<details>
<summary>メモ</summary>

- Depandabotのセットアップ
  - `.github/dependabot.yml`の`package-ecosystem`に値を設定 (例: npm,yarn,pip)
- CodeQLのセットアップ
  - https://dev.classmethod.jp/articles/github-code-scanning/
  - [対応言語](https://codeql.github.com/docs/codeql-overview/supported-languages-and-frameworks/)

</details>

## For Contributor 🤝

本プロジェクトにコントリービュートする場合は[コントリービュートガイド](docs/README.md)をお読みください。

## Reference ✨

- [doremire/Awesome-README](https://github.com/doremire/Awesome-README)
