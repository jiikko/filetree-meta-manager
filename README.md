# README

- ファイルツリーを WEB から閲覧するためのツールです. ファイルの実体には関与しません
- バックアップ用 HDD の管理を楽にするために制作しました

## System dependencies

- Ruby
  - `cat .ruby-version`
- MySQL8.X

## チュートリアル

- WEB アプリケーションでアカウントを作成してください
  - WEB アプリケーションが公開されていない場合は、自分でホスティングしてください
- WEB アプリケーション API キーを取得してください
  - API キーは後で使います
- https://github.com/jiikko/filetree-meta-manager/releases から最新のバイナリをダウンロードしてください
- 設定ファイルの雛形を作成してください
  - `./filetree_dumper --init-config target_dir`
  - NOTE: ディレクトリの root には、設定ファイルが必要です
- `target_dir/.filetree_meta_manager.yml` をテキストエディタで開いて、url, api_key, device を設定してください
- `filetree_dumper target_dir`を実行して、WEB アプリケーションに送信する
- WEB を確認する

## 開発

### サーバのセットアップ

- brew install mysql
- bin/rails db:drop db:create db:migrate

### テストを実行する

- `bin/rspec`

### Deployment

- WEB アプリケーション
  - TODO
- CLI ツール
  - `.app_version`を変更して、 github actions workflow から `Release CLI` を手動実行してください

### デバッグ TIPS

- 開発環境の cli ツールを実行する方法
  - cd client; go run cmd/dump-filetree/main.go tmp
