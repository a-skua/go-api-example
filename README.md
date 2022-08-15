# Golang API Example
Go言語によるAPI実装例になります。
API実装時の参考にして貰えればと思います。

This is an API implementation example in Golang.
I hope you can use it as a reference when implementing the API.

## 要件
会社の人事情報管理サービスAPIを作ることになりました。
会社には「営業部、開発部」といった組織階層が存在します。
組織階層には深さの制限はありません。
組織階層はツリー構造になります。
親階層は一意に定りますが、子階層は任意の数だけ発生します。
会社には所属する従業員が存在します。
一部の従業員をシステムの管理者とすることができます。
管理者は複数人設定することができます。
将来、システムの管理者には細かく操作権限が付与される可能性があります。
また会社には「部長、課長」と言った肩書きを用意することができます。
各階層には従業員を配置することができます。
組織階層、肩書き、従業員はそれぞれ紐づけることができます。
紐付けには「社員Aは営業部長と開発部長を兼任する」と言ったパターンも存在します。
近年、社会的な副業の推進により、従業員が複数の企業に所属することも一般的になってきました。

## 簡易API設計
今回はREST APIを想定しています。
会社はユーザーによって作成することができ、
会社を作成したユーザーは管理者として会社の従業員になります。
管理者であるユーザーは会社の全てを操作できますが、
管理者でない従業員は、会社に所属する自身の情報の確認しか許されていません。

- ユーザー情報を扱うエンドポイント
  `/user`
  - 登録
    `POST /user`
    - 条件
      - `user.name`
        - 1文字以上、255文字以下
      - `user.password`
        - 8文字以上、255文字以下
        - Responseの `user.password`は伏せ字(`*****`)とする
    - Request Body
      ```json
      {
        "user": {
          "name": "Bob",
          "password": "password"
        }
      }
      ```
    - Response Body
      ```json
      {
        "user": {
          "id": 1,
          "name": "Bob",
          "password": "*****",
          "updated_at": "2006-01-02T15:04:05Z07:00"
        }
      }
      ```
  - 取得
    `GET /user/{user_id}`
    - 条件
      - `user.password`
        - 表示内容は伏字(`*****`)固定
    - Response Body
      ```json
      {
        "user": {
          "id": 1,
          "name": "Bob",
          "password": "*****",
          "updated_at": "2006-01-02T15:04:05Z07:00"
        }
      }
      ```
  - 更新
    `PUT /user/{user_id}`
    - 条件
      (登録時と同じ)
      - `user.name`
        - 1文字以上、255文字以下
      - `user.password`
        - 8文字以上、255文字以下
        - Responseの `user.password`は伏せ字(`*****`)とする
    - Request Body
      ```json
      {
        "user": {
          "name": "Bob",
          "password": "password"
        }
      }
      ```
    - Response Body
      ```json
      {
        "user": {
          "id": 1,
          "name": "Bob",
          "password": "*****",
          "updated_at": "2006-01-02T15:04:05Z07:00"
        }
      }
      ```
  - 削除
    `DELETE /user/{user_id}`
    - Response Body
      ```json
      {
        "user": {}
      }
      ```

- 会社情報を扱うエンドポイント
  `/company`
  - 登録 `POST /company`
    - 条件
      - `company.name`
        - 1文字以上255文字以下
      - `company.owner_id`
        - 実在するユーザーID
    - Request Body
      ```json
      {
        "company": {
          "name": "GREATE COMPANY",
          "owner_id": 1
        }
      }
      ```
    - Response Body
      ```json
      {
        "company": {
          "id": 1,
          "name": "GREATE COMPANY",
          "owner_id": 1,
          "updated_at": "2006-01-02T15:04:05Z07:00"
        }
      }
      ```
  - 取得
    `GET /company/{company_id}`
    - Response Body
      ```json
      {
        "company": {
          "id": 1,
          "name": "GREATE COMPANY",
          "owner_id": 1,
          "updated_at": "2006-01-02T15:04:05Z07:00"
        }
      }
      ```
  - 更新
    `PUT /company/{company_id}`
    - 条件
      (登録時と同じ)
      - `company.name`
        - 1文字以上255文字以下
      - `company.owner_id`
        - 実在するユーザーID
    - Request Body
      ```json
      {
        "company": {
          "name": "GREATE COMPANY",
          "owner_id": 1
        }
      }
      ```
    - Response Body
      ```json
      {
        "company": {
          "id": 1,
          "name": "GREATE COMPANY",
          "owner_id": 1,
          "updated_at": "2006-01-02T15:04:05Z07:00"
        }
      }
      ```
  - 削除
    `DELETE /company/{company_id}`
    - Response Body
      ```json
      {
        "company": {}
      }
      ```

## このリポジトリの使い方
開発によく使うコマンドは `Makefile` にまとめています。
`make up` で API を実行できます。
次に、 `make migrate` で DBの初期化を行うことで、実際に API として利用することができます。

### Docker
開発環境として、`docker comkpose` を利用しています。
`docker compose` の各サービスの役割は次のとおりです。
- `api`
  - API本体
- `db`
  - データベース本体
  - `mysql` を利用しています
- `migrate`
  - データベースのマイグレーションを行うためのコンテナ
  - マイグレーションツールとして、 `Rails` の `ActiveRecord` を利用しています
- `gopher`
  - `go test` や `go fmt` など、 `go` の実行環境用のコンテナ

### Dirctory Structure
```
.
├── _e2e            # E2Eテスト
├── _img            # Docker Images
├── _migrate        # データベース Migration
├── cover           # Coverage出力
└── src
    ├── cmd         # package main
    ├── env         #
    ├── http-handle # HTTPハンドラ
    ├── pkg         # メインプログラム
    └── repository  # データベース
```

### 実装済みエンドポイント
- [x] `/user`
- [ ] `/company`
