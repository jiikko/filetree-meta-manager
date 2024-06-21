# filetree-meta-manager

このツールは、あらかじめアップロードしたファイルツリーを WEB から閲覧するために作成されました。ファイルの実体には関与せず、主にバックアップ用 HDD の管理を容易にすることを目的としています。

ツールは CLI ツールと WEB アプリケーションの 2 つのコンポーネントから構成されています。CLI ツールはファイルツリーの収集および WEB アプリケーションへの送信を行います。WEB アプリケーションは、収集されたファイルツリーの閲覧を提供します。

![image](https://github.com/jiikko/filetree-meta-manager/assets/1664497/f98cd076-ed14-4da9-80ae-d1b7d4b07444)

## 使っている技術

- Ruby
  - `cat .ruby-version`
- MySQL 8.X
- GCP Cloud Run

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

## 開発者向け情報

### サーバのセットアップ

- brew install mysql rbenv
- rbenv install `cat .ruby-version`
- bin/rails db:drop db:create db:migrate

### CLI ツールのセットアップ

- brew install go
- go run cmd/dump-filetree/main.go --version

### テストの実行

- `bin/rspec`

### デプロイメント

#### WEB アプリケーション

- リポジトリを fork して自分でホスティングしてください
- 本番環境でユーザを新規作成するには、環境変数 SIGNUP_ENABLED に 1 をセットしてデプロイしてください
- TODO

#### CLI ツール

- `.app_version`を変更して、 github actions workflow から `Release CLI` を手動実行してください

### デバッグ TIPS

- 開発環境の CLI ツールを実行する方法
  - cd client; go run cmd/dump-filetree/main.go tmp
