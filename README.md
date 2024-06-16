# filetree-meta-manager

このツールは、ファイルツリーをWEBから閲覧するためのものです。ファイルの実体には関与せず、主にバックアップ用HDDの管理を容易にするために制作されました。

このツールは、CLIツールとWEBアプリケーションの2つのコンポーネントから構成されています。CLIツールは、ファイルツリーの収集およびWEBアプリケーションへの送信を行います。WEBアプリケーションは、収集されたファイルツリーの閲覧を提供します。

![image](https://github.com/jiikko/filetree-meta-manager/assets/1664497/f98cd076-ed14-4da9-80ae-d1b7d4b07444)

## 使っている技術

- Ruby
  - `cat .ruby-version`
- MySQL 8.X

## チュートリアル

- WEB アプリケーションでアカウントを作成
  - WEB アプリケーションが公開されていない場合は、自分でホスティングしてください
- API キーを取得
  - WEB アプリケーションから API キーを取得してください。あとで使用します。
- CLI ツールをダウンロード
  - https://github.com/jiikko/filetree-meta-manager/releases から、実行環境に合う最新のバイナリをダウンロードしてください
- 設定ファイルの雛形作成
  - `./filetree_dumper --init-config target_dir`
  - NOTE: ディレクトリのルートには、設定ファイルが必要です
- 設定ファイルの編集
  - `target_dir/.filetree_meta_manager.yml` をテキストエディタで開いて、url, api_key, device を設定してください
- ファイルツリーのダンプと送信
  - `./filetree_dumper target_dir`を実行して、WEB アプリケーションに送信する
- WEB アプリケーションでファイルツリーを閲覧

## 開発関係

### サーバのセットアップ

- brew install mysql rbenv
- rbenv install `cat .ruby-version`
- bin/rails db:drop db:create db:migrate

### テストの実行

- `bin/rspec`

### デプロイメント

- WEB アプリケーション
  - TODO
- CLI ツール
  - `.app_version`を変更して、 github actions workflow から `Release CLI` を手動実行してください

### デバッグ TIPS

- 開発環境の cli ツールを実行する方法
  - cd client; go run cmd/dump-filetree/main.go tmp
