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
- `/user`: ユーザーを扱う。
  - `POST /user`: `name` は1文字以上、 `password` は8文字以上とする。
    ```json
    {
      "user": {
        "name": "Bob",
        "password": "password"
      }
    }
    ```
  - `GET /user/{user_id}`: `password`の表示は 伏字(`*****`)固定。
    ```json
    {
      "user": {
        "id": 1,
        "name": "Bob",
        "password": "*****",
        "companies": [
          {
            "id": 1,
            "name": "GREATE COMPANY"
          }
        ]
      }
    }
    ```
  - `PUT /user/{user_id}`: 更新時の条件は作成時と同じ。
    ```json
    {
      "user": {
        "name": "Bob",
        "password": "password"
      }
    }
    ```
  - `DELETE /user/{user_id}`
    ```json
    {
      "user": {}
    }
    ```

- `/company`: 会社情報を扱う。
  - `POST /company`: `name` は1文字以上とする。
    ```json
    {
      "company": {
        "name": "GREATE COMPANY"
      }
    }
    ```
  - `GET /company/{company_id}`
    ```json
    {
      "company": {
        "id": 1,
        "name": "GREATE COMPANY"
      }
    }
    ```
  - `PUT /company/{company_id}`: 更新時の条件は作成時と同じ。
    ```json
    {
      "company": {
        "name": "GREATE COMPANY"
      }
    }
    ```
  - `DELETE /company/{company_id}`

- `/company/{company_id}/groups`: 会社の組織階層を扱う。
  - `GET /company/{company_id}/groups`: Header `X-User-Id` 必須。
    ```json
    {
      "groups": [
        {
          "id": 1,
          "name": "開発部",
          "parent": null,
          "children": [
            {
              "id": 2,
              "name": "第1課",
              "children": []
            },
            {
              "id": 3,
              "name": "第2課",
              "children": []
            }
          ]
        }
      ]
    }
    ```
  - `POST /company/{company_id}/groups`: `name` は1文字以上、 `parent` は整数もしくは `null`, `children` は整数リストとする。 Header `X-User-Id` 必須。
    ```json
    {
      "group": {
        "name": "開発部",
        "parent": null,
        "children": [ 2, 3 ]
      }
    }
    ```
  - `GET /company/{company_id}/groups/{group_id}`: Header `X-User-Id` 必須。
    ```json
    {
      "group": {
        "id": 1,
        "name": "開発部",
        "parent": null,
        "children": [
          {
            "id": 2,
            "name": "第1課",
            "children": []
          },
          {
            "id": 3,
            "name": "第2課",
            "children": []
          }
        ]
      }
    }
    ```
  - `PUT /company/{company_id}/groups/{group_id}`: 更新時の条件は作成時と同じ。 Header `X-User-Id` 必須。
    ```json
    {
      "group": {
        "name": "開発部",
        "parent": null,
        "children": [ 2, 3 ]
      }
    }
    ```
  - `DELETE /company/{company_id}/groups/{group_id}`: Header `X-User-Id` 必須。

- `/company/{company_id}/titles`: 会社の肩書きを扱う。
  - `GET /company/{company_id}/titles`: Header `X-User-Id` 必須。
    ```json
    {
      "titles": [
        {
          "id": 1,
          "name": "社長"
        },
        {
          "id": 2,
          "name": "部長"
        }
      ]
    }
    ```
  - `POST /company/{company_id}/titles`: `name` は1文字以上とする。 Header `X-User-Id` 必須。
    ```json
    {
      "title": {
        "name": "社長"
      }
    }
    ```
  - `GET /company/{company_id}/titles/{title_id}`: Header `X-User-Id` 必須。
    ```json
    {
      "title": {
        "id": 1,
        "name": "社長"
      }
    }
    ```
  - `PUT /company/{company_id}/titles`: 更新時の条件は作成時と同じ。 Header `X-User-Id` 必須。
    ```json
    {
      "title": {
        "name": "社長"
      }
    }
    ```
  - `DELETE /company/{company_id}/titles`: Header `X-User-Id` 必須。

- `/company/{company_id}/roles`: サービスにおける役割を扱う。
  - `GET /company/{company_id}/roles`: Header `X-User-Id` 必須。
    ```json
    {
      "roles": [
        {
          "id": 1,
          "name": "管理者"
        }
      ]
    }
    ```
  - `GET /company/{company_id}/roles/{role_id}`: Header `X-User-Id` 必須。
    ```json
    {
      "role": {
        "id": 1,
        "name": "管理者"
      }
    }
    ```
- `/company/{company_id}/employees`: 会社の従業員情報を扱う。
  - `GET /company/{company_id}/users`: Header `X-User-Id` 必須。
    ```json
    {
      "employees": [
        {
          "id": 1,
          "name": "Bob",
          "password": "*****",
          "roles": [
            {
              "id": 1,
              "name": "管理者"
            }
          ],
          "information": [
            {
              "group": {
                "id": 1,
                "name": "開発部"
              },
              "title": {
                "id": 2,
                "name": "部長"
              }
            }
          ]
        }
      ]
    }
    ```
  - `POST /company/{company_id}/employees/{user_id}`: `user_id` は Path による指定。
    `roles` は整数リスト、`group` と `title` はそれぞれ整数。 Header `X-User-Id` 必須。
    ```json
    {
      "employee": {
        "roles": [ 1 ],
        "information": [
          {
            "group": 1,
            "title": 2
          }
        ]
      }
    }
    ```
  - `GET /company/{company_id}/employees/{user_id}`: Header `X-User-Id` 必須。
    ```json
    {
      "employee": {
        "id": 1,
        "name": "Bob",
        "password": "*****",
        "roles": [
          {
            "id": 1,
            "name": "管理者"
          }
        ],
        "information": [
          {
            "group": {
              "id": 1,
              "name": "開発部"
            },
            "title": {
              "id": 2,
              "name": "部長"
            }
          }
        ]
      }
    }
    ```
  - `PUT /company/{company_id}/employees/{user_id}`: 更新時の条件は作成時と同じ。 Header `X-User-Id` 必須。
    ```json
    {
      "employee": {
        "roles": [ 1 ],
        "information": [
          {
            "group": 1,
            "title": 2
          }
        ]
      }
    }
    ```
  - `DELETE /company/{company_id}/employees/{user_id}`: Header `X-User-Id` 必須。

## このリポジトリの使い方
開発によく使うコマンドは `Makefile` にまとめています。
`make up` で API を実行できます。
次に、 `make migrate` で DBの初期化を行うことで、実際に API として利用することができます。

### Docker
開発環境として、`docker comkpose` を利用しています。
`docker compose` の各サービスの役割は次のとおりです。
- `api`: API本体。
- `db`: データベース本体。ここでは `mysql` を使用。
- `migrate`: データベースのマイグレーションを行うためのコンテナ。
  ここではマイグレーションツールとして、 `Rails` の `ActiveRecord` を利用している。
- `gopher`: `go test` や `go fmt` など、 `go` の実行環境用のコンテナ。

### Dirctories
```
.
├── _e2e
├── _img
│   └── initdb.d
├── _migrate
│   └── db
│       └── migrate
├── cover
└── src
    ├── cmd
    ├── env
    ├── http-handle
    │   ├── request
    │   └── response
    ├── pkg
    │   ├── entity
    │   ├── repository
    │   └── service
    └── rdb-repository
        └── model
```
- `src`: API のソースファイル群はここにまとめてあります。
  - `cmd`: API の `main` 関数がここにあります。
    このファイルから、 `import` されている `packages` を辿ることで、
    ファイルの全体像を把握できると思います。
  - `pkg`: API のコア `package` 群です。
    `pkg` の内部に配置されているコードは HTTP や RDB について知ることはありません。
  - `http-handle`: HTTP について集約しています。
  - `rdb-repository`: RDB (今回は MySQL) について集約しています。
- `cover`: テストのカバレッジはここに出力されます(git の管理対象外です)。
- `_migrate`: マイグレーションファイルをまとめてあります。
- `_img`: `Dockerfile` などをまとめてあります。
- `_e2e`: `e2e` テストファイルをまとめてあります。
- `_migrate`: マイグレーションファイルをまとめてあります。
