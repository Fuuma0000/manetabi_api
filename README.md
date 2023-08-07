# リレーション

---

```mermaid
erDiagram
    users ||--o{ plans: "投稿する"
    plans ||--o{ blocks: "持つ"
    blocks ||--o{ photos: "持つ"
    plans ||--o{ plans_locations: "登録する"
    plans_locations||--o{ locations: "参照する"
    users ||--o{ plans_bookmarks: "ブックマークする"
    users ||--o{ blocks_bookmarks: "ブックマークする"
    plans ||--o{ plans_bookmarks: "ブックマークされる"
    users ||--o{ groups: "所属する"
    plans_categories }o--|| categories: "参照する"
    plans ||--o{ groups : "所属される"
    plans ||--o{ plans_categories: "登録する"
    plans ||--o{ plans_likes: "いいねされる"
    users ||--o{ plans_likes: "いいねする"
    blocks ||--o{ blocks_bookmarks: "ブックマークされる"

    users {
        user_id INT(255) PK
        user_name VARCHAR(255)
        email VARCHAR(255)
        password VARCHAR(255)
        profile_image_path VARCHAR(255)
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    plans {
        plan_id INT(255) PK
        user_id INT(255) FK
        title VARCHAR(100)
        description TEXT
        thumbnail_path VARCHAR(255)
        cost INT(255)
        start_date TIMESTAMP
        end_date TIMESTAMP
        is_public BOOLEAN
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    blocks {
        block_id INT(255) PK
        plan_id INT(255) FK
        block_name VARCHAR(255)
        start_date TIMESTAMP
        end_date TIMESTAMP
        memo TEXT
        cost INT(255)
        address VARCHAR(255)
        details TEXT
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    photos {
        photo_id INT(255) PK
        block_id INT(255) FK
        photo_path VARCHAR(255)
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    locations {
        location_id INT(255) PK
        location_name VARCHAR(255)
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    categories {
        category_id INT(255) PK
        category_name VARCHAR(255)
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    plans_bookmarks {
        user_id INT(255) FK
        plan_id INT(255) FK
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    blocks_bookmarks {
        user_id INT(255) FK
        block_id INT(255) FK
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    plans_likes {
        user_id INT(255) FK
        plan_id INT(255) FK
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    plans_locations {
        plan_id INT(255) FK
        location_id INT(255) FK
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    plans_categories {
        plan_id INT(255) FK
        category_id INT(255) FK
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }
    groups {
        plan_id INT(255) FK
        user_id INT(255) FK
        created_at TIMESTAMP
        updated_at TIMESTAMP
    }

```

# データ設計

---

## **user**

| カラム名           | 説明                     | 型           | Key         | AUTOINCREMENT | Unique | Nullable |
| ------------------ | ------------------------ | ------------ | ----------- | ------------- | ------ | -------- |
| user_id            | ユーザ ID                | INT(255)     | PRIMARY KEY | YES           |        | NO       |
| user_name          | ユーザー名               | VARCHAR(255) |             |               |        | NO       |
| email              | メールアドレス           | VARCHAR(255) |             |               | YES    | NO       |
| password           | ハッシュ化したパスワード | VARCHAR(255) |             |               |        | NO       |
| profile_image_path | プロフィール画像パス     | VARCHAR(255) |             |               |        | YES      |
| created_at         | 作成した時間             | TIMESTAMP    |             |               |        | NO       |
| updated_at         | 更新した時間             | TIMESTAMP    |             |               |        | NO       |

## **plans**

| カラム名       | 説明               | 型           | Default | Key         | AUTOINCREMENT | Nullable |
| -------------- | ------------------ | ------------ | ------- | ----------- | ------------- | -------- |
| plan_id        | プラン ID          | INT(255)     |         | PRIMARY KEY | YES           | NO       |
| user_id        | ユーザ ID          | INT(255)     |         | FOREIGN KEY |               | NO       |
| title          | タイトル           | VARCHAR(100) |         |             |               | NO       |
| description    | 説明文             | TEXT         |         |             |               | YES      |
| thumbnail_path | サムネイル画像パス | VARCHAR(255) |         |             |               | YES      |
| cost           | 費用               | INT(255)     |         |             |               | YES      |
| start_date     | 開始日時           | TIMESTAMP    |         |             |               | YES      |
| end_date       | 終了日時           | TIMESTAMP    |         |             |               | YES      |
| is_public      | 公開フラグ         | BOOLEAN      | FALSE   |             |               | NO       |
| created_at     | 作成した時間       | TIMESTAMP    |         |             |               | NO       |
| updated_at     | 更新した時間       | TIMESTAMP    |         |             |               | NO       |

## **blocks**

| カラム名   | 説明         | 型           | Key         | AUTOINCREMENT | Nullable |
| ---------- | ------------ | ------------ | ----------- | ------------- | -------- |
| block_id   | ブロック ID  | INT(255)     | PRIMARY KEY | YES           | NO       |
| plan_id    | プラン ID    | INT(255)     | FOREIGN KEY |               | NO       |
| block_name | ブロック名   | VARCHAR(255) |             |               | NO       |
| start_date | 開始日時     | TIMESTAMP    |             |               | NO       |
| end_date   | 終了日時     | TIMESTAMP    |             |               | NO       |
| memo       | メモ         | TEXT         |             |               | YES      |
| cost       | 費用         | INT          |             |               | YES      |
| address    | 住所         | VARCHAR(255) |             |               | YES      |
| details    | 詳細         | TEXT         |             |               | YES      |
| created_at | 作成した時間 | TIMESTAMP    |             |               | NO       |
| updated_at | 更新した時間 | TIMESTAMP    |             |               | NO       |

## **photos**

| カラム名   | 説明         | 型           | Key         | AUTOINCREMENT | Nullable |
| ---------- | ------------ | ------------ | ----------- | ------------- | -------- |
| photo_id   | 画像 ID      | INT(255)     | PRIMARY KEY | YES           | NO       |
| block_id   | ブロック ID  | INT(255)     | FOREIGN KEY |               | NO       |
| photo_path | 画像パス     | VARCHAR(255) |             |               | NO       |
| created_at | 作成した時間 | TIMESTAMP    |             |               | NO       |
| updated_at | 更新した時間 | TIMESTAMP    |             |               | NO       |

## **locations**

| カラム名      | 説明         | 型           | Key         | AUTOINCREMENT | Unique | Nullable |
| ------------- | ------------ | ------------ | ----------- | ------------- | ------ | -------- |
| location_id   | 場所 ID      | INT(255)     | PRIMARY KEY | YES           |        | NO       |
| location_name | 場所名       | VARCHAR(255) |             |               | YES    | NO       |
| created_at    | 作成した時間 | TIMESTAMP    |             |               |        | NO       |
| updated_at    | 更新した時間 | TIMESTAMP    |             |               |        | NO       |

## **categories**

| カラム名      | 説明         | 型           | Key         | AUTOINCREMENT | Unique | Nullable |
| ------------- | ------------ | ------------ | ----------- | ------------- | ------ | -------- |
| category_id   | カテゴリ ID  | INT(255)     | PRIMARY KEY | YES           |        | NO       |
| category_name | カテゴリ名   | VARCHAR(255) |             |               | YES    | NO       |
| created_at    | 作成した時間 | TIMESTAMP    |             |               |        | NO       |
| updated_at    | 更新した時間 | TIMESTAMP    |             |               |        | NO       |

## **plans_bookmarks**

| カラム名   | 説明         | 型        | Key         | Unique | Nullable |
| ---------- | ------------ | --------- | ----------- | ------ | -------- |
| user_id    | ユーザ ID    | INT(255)  | FOREIGN KEY |        | NO       |
| plan_id    | ブロック ID  | INT(255)  | FOREIGN KEY |        | NO       |
| created_at | 作成した時間 | TIMESTAMP |             |        | NO       |
| updated_at | 更新した時間 | TIMESTAMP |             |        | NO       |

## **blocks_bookmarks**

| カラム名   | 説明         | 型        | Key         | Nullable |
| ---------- | ------------ | --------- | ----------- | -------- |
| user_id    | ユーザ ID    | INT(255)  | FOREIGN KEY | NO       |
| block_id   | ブロック ID  | INT(255)  | FOREIGN KEY | NO       |
| created_at | 作成した時間 | TIMESTAMP |             | NO       |
| updated_at | 更新した時間 | TIMESTAMP |             | NO       |

## **plans_likes**

| カラム名   | 説明         | 型        | Key         | Nullable |
| ---------- | ------------ | --------- | ----------- | -------- |
| user_id    | ユーザ ID    | INT(255)  | FOREIGN KEY | NO       |
| plan_id    | 投稿 ID      | INT(255)  | FOREIGN KEY | NO       |
| created_at | 作成した時間 | TIMESTAMP |             | NO       |
| updated_at | 更新した時間 | TIMESTAMP |             | NO       |

## **plans_locations**

| カラム名    | 説明         | 型        | Key         | Nullable |
| ----------- | ------------ | --------- | ----------- | -------- |
| plan_id     | 投稿 ID      | INT(255)  | FOREIGN KEY | NO       |
| location_id | 場所 ID      | INT(255)  | FOREIGN KEY | NO       |
| created_at  | 作成した時間 | TIMESTAMP |             | NO       |
| updated_at  | 更新した時間 | TIMESTAMP |             | NO       |

## **plans_categories プランのカテゴリーを保存する**

| カラム名    | 説明         | 型        | Key         | Nullable |
| ----------- | ------------ | --------- | ----------- | -------- |
| plan_id     | 投稿 ID      | INT(255)  | FOREIGN KEY | NO       |
| category_id | カテゴリ ID  | INT(255)  | FOREIGN KEY | NO       |
| created_at  | 作成した時間 | TIMESTAMP |             | NO       |
| updated_at  | 更新した時間 | TIMESTAMP |             | NO       |

## **groups プランにユーザが所属する**

| カラム名   | 説明         | 型        | Key         | Nullable |
| ---------- | ------------ | --------- | ----------- | -------- |
| plan_id    | プラン ID    | INT(255)  | FOREIGN KEY | NO       |
| user_id    | ユーザ ID    | INT(255)  | FOREIGN KEY | NO       |
| created_at | 作成した時間 | TIMESTAMP |             | NO       |
| updated_at | 更新した時間 | TIMESTAMP |             | NO       |
