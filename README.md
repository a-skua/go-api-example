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
`user`情報の作成(CREATE)は誰でもできますが、
更新(UPDATE) 取得(GET) 削除(DELETE)はユーザー本人である必要があります。
ここでは Header `X-User-Id`と Path `user_id` が一致することによって
本人であることを確認します。
**※実運用では真似しないようお願いします。**
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
        "password": "*****"
      }
    }
    ```
  - `GET /user/{user_id}`: `password`の表示は 伏字(`*****`)固定。 Header `X-User-Id` 必須。
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
  - `PUT /user/{user_id}`: 更新時の条件は作成時と同じ。 Header `X-User-Id` 必須。
    ```json
    {
      "user": {
        "name": "Bob",
        "password": "*****"
      }
    }
    ```
  - `DELETE /user/{user_id}`: Header `X-User-Id` 必須。

- `/company`: 会社情報を扱う。
  - `POST /company`: `name` は1文字以上とする。 Header `X-User-Id` 必須。
    ```json
    {
      "company": {
        "name": "GREATE COMPANY"
      }
    }
    ```
  - `GET /company/{company_id}`: Header `X-User-Id` 必須。
    ```json
    {
      "company": {
        "id": 1,
        "name": "GREATE COMPANY"
      }
    }
    ```
  - `PUT /company/{company_id}`: 更新時の条件は作成時と同じ。 Header `X-User-Id` 必須。
    ```json
    {
      "company": {
        "name": "GREATE COMPANY"
      }
    }
    ```
  - `DELETE /company/{company_id}`: Header `X-User-Id` 必須。

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
          "管理者"
        }
      ]
    }
    ```
  - `GET /company/{company_id}/roles/{role_id}`: Header `X-User-Id` 必須。
    ```json
    {
      "role": {
        "id": 1,
        "管理者"
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
              "管理者"
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
            "管理者"
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
