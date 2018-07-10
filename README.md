FORMAT: 1A
# ケーススタディ HAL大阪　API一覧

# Group BOCCO x 学習ドリル
## ユーザー登録 [/signup]
ユーザー情報の登録、およびサインインするためのAPI

### サインアップ [POST]
ユーザー情報の登録を行います。

+ Request (applicaition/json)

    + Attribute
        + name: sample
        + pass: password

+ Response 200 (application/json)

    + Attribute

        + success: ユーザー登録を行いました。
+ Response 400 (application/json)

    + Attribute

        + error: 登録済みのユーザー名です。
 
## サインイン [/signin]

### サインイン [POST]
登録されているユーザー情報を元にサインインを行います。

+ Request (applicaition/json)

    + Attribute
        + name: sample
        + pass: password

+ Response 200 (application/json)

    + Attribute

        + token: sample

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## ユーザー削除 [/user]
### ユーザー削除[DELETE]
登録されているユーザー情報を削除します。

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: ユーザー情報を削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## 子供情報設定 [/child/{child_id}]
### 子供情報追加[POST]
子供の誕生日と性別、ニックネームを設定します。

+ Request(application/json)
    + Headers

            Authorization: token

    + Attributes
        + nickname: sample
        + birthday : `2000-01-01`
        + sex : 0 (number) - 0:男、1:女

+ Response 200 (application/json)

    + Attribute

        + child_id: 1 (number)

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### 子供情報取得[GET]
子供ID,誕生日,ニックネーム,性別の一覧を取得します。

+ Request(application/json)
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes

        + children (array)
            + (object)
                + child_id: 1 (number)
                + birthday: `2016-10-01T09:00:00+09:00`
                + nickname: sample
                + sex: 0 (number) - 0:男、1:女

            + (object)
                + child_id: 2 (number)
                + birthday: `2017-03-19T09:00:00+09:00`
                + nickname: index
                + sex: 1 (number) - 0:男、1:女

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### 子供情報削除[DELETE]
登録されている子どもIDの情報を削除します。

+ Parameters
    + child_id: 1

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: 削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## デバイス [/device/{device_id}]

### デバイスID発行[POST]
新規登録するデバイスIDの発行を行います。

+ Request(application/json)
    + Headers

            Authorization: token

    + Attribute
        + child_id: 1 (number)

+ Response 200 (application/json)

    + Attribute

        + pin: 0000

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー
     

### デバイス一覧取得[GET]
現在登録されているデバイスIDの一覧を取得します。

+ Request(application/json)
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes
        
        + devices (array)
            + (object)
                + child_id :1 (number)
                + nickname: sample
                + child_devices (array)
                    + sample,
                    + index
            + (object)
                + child_id: 2 (number)
                + nickname: test
                + child_devices (array)
                    + test,
                    + buf
            
+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### デバイスID削除[DELETE]
登録されているデバイスIDを削除します。

+ Parameters
    + device_id: sample

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: ボタンIDを削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## BOCCOAPI [/bocco]

### BOCCOAPI設定[POST]
BOCCOAPIに登録したメールアドレスと、パスワードの入力

+ Request
    + Headers

            Authorization: token

    + Attribute
        + email: sample@gmail.com
        + key : sample - APIkey
        + pass: abc123

+ Response 200 (application/json)

    + Attribute

        + success: メールアドレスとパスワードを登録しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### BOCCOAPI設定の取得[GET]
BOCCOAPIに登録したメールアドレスの取得

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + email: sample@gmail.com

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### BOCCOAPI削除[DELETE]
BOCCOAPIに登録したメールアドレスと、パスワードの削除

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: メールアドレスとパスワードを削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー


## グラフ用のデータ取得 [/work/record/{child_id}{?filter}]

### グラフ用の解答データの取得[GET]
指定された子どもの記録情報を取得

+ Parameters
    + child_id: 1
    + filter: date - もしくはgenre

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes
        + records (array)
            + (object)
                + date: `2018-06-21T13:35:08+09:00` - 回答日時
                + num_ans: 10(number) - 回答数
                + num_corr: 5(number)- 正答数

            + (object)
                + date: `2018-06-22T13:35:08+09:00`
                + num_ans: 7(number) - 回答数
                + num_corr: 6(number)- 正答数

    + Attributes
        + records (array)
            + (object)
                + num_probs: 50 - ジャンルの総問題数
                + genre:  算数 - 回答ジャンル
                + num_ans: 10(number) - 回答数
                + num_corr: 5(number)- 正答数

            + (object)
                + num_probs: 30 - ジャンルの総問題数
                + genre:  社会 - 回答ジャンル
                + num_ans: 8(number) - 回答数
                + num_corr: 8(number)- 正答数

   

+ Response 400 (application/json)

    + Attribute

        + error: 回答情報が見つかりませんでした。

## 詳細データの取得 [/work/record/detail/{child_id}{?date,genre}]

### 詳細な解答データの取得[GET]
指定された子どもの記録情報を取得

+ Parameters
    + child_id: 1
    + date: `2018-07-04`
    + genre: 1 - genre_id

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes
        + records (array)
            + (object)
                + date: `2018-06-21T13:35:08+09:00` - 回答日時
                + genre_name: 算数
                + detail(array)
                    + (object)
                        + sentence: 1 + 1は？
                        + user_ans: 2
                        + correct: 2
                        + result: true (boolean) - 正解:true,不正解:false
                    
                    + (object)
                        + sentence: 3 - 2は？
                        + user_ans: 2
                        + correct: 1
                        + result: false (boolean)

            + (object)
                + date: `2018-06-22T13:35:08+09:00`
                + genre_name: 社会
                + detail(array)
                    + (object)
                        + sentence: 兵庫県の県庁所在地は？
                        + user_answer: 兵庫市
                        + correct: 神戸市
                        + result: false (boolean)

+ Response 400 (application/json)

    + Attribute

        + error: 回答情報が見つかりませんでした。

## メッセージ [/work/message/{child_id}{?condtion}]

### メッセージ登録[POST]
オリジナルメッセージの登録を行います。

+ Request (application/json)
    + Headers

            Authorization: token

    + Attribute
        + child_id: 1 (number)
        + message_call : 3 (number) - (1: 正解,2:不正解,3: 連続正解時)
        + condition : 10 (number) - 3の時
        + message: practice

+ Response 200 (application/json)

    + Attribute

        + success: メッセージを編集しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

### メッセージ取得[GET]
登録されているメッセージとメッセージ出力条件を取得します。

+ Request
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attributes

        + messages(array)
            + (object)
                + child_id: 1 (number)
                + nickname: sample
                + child_messages(array)
                    + (object)
                        + message_call: 2 (number)
                        + message: practice
                    + (object)
                        + message_call: 3 (number)
                        + condtion: 5 (number)
                        + message: sample
            + (object)
                + child_id: 2 (number)
                + nickname: index
                + child_messages(array)
                    + (object)
                        + message_call: 3 (number)
                        + condition: 10 (number)
                        + message: sample
                    + (object)
                        + message_call: 1 (number)
                        + message: hoge

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

## メッセージ削除 [/work/message/{message_id}]
### メッセージ削除[DELETE]
オリジナルメッセージの削除を行います。

+ Parameters
    + message_id: sample

+ Request (application/json)
    + Headers

            Authorization: token

+ Response 200 (application/json)

    + Attribute

        + success: メッセージを削除しました。

+ Response 400 (application/json)

    + Attribute

        + error: ログインエラー

# Group ICリーダーAPI

## デバイス [/thing/registration]

### デバイス登録[POST]
デバイスIDと各デバイスとの紐付けを行います。


+ Request (applicaition/json)
 
    + Attribute
        + pin: 0000
        
+ Response 200 (application/json)

 + Attribute

      + device_id: sample

+ Response 400 (application/json)

    + Attribute

        + error: pinが見つかりません。

## ICリーダー [/thing/reader]
### 回答データを送信[POST]
デバイス情報と読み取ったタグの情報を送信。

+ Request(application/json)

    + Attribute
        + device_id: sample
        + uuid: 1234
        + old_uuid: 5678
        
+ Response 200 (application/json)

    + Attribute

        + success: true (boolean)

+ Response 418 (application/json)

    + Attribute

        + error: データベースエラー

# Group 問題作成API
## 問題作成 [/question/create]
### 問題データの作成[POST]
問題データをDBに登録

+ Request(application/json)
	
	+ Attribute
		+ book_id: 1(number)
		+ question_no: 1(number)
		+ sentence(array)
			+ (object)
				+ tag_id: sample
				+ text: 回文はどれ？
		+ answer(array)
			+ (object)
				+ tag_id: index
				+ text: 絵本
			+ (object)
				+ tag_id: buf
				+ text: 新聞紙
			+ (object)
				+ tag_id: hoge
				+ text: 漫画
		+ correct: buf
		+ genre: 1

+ Response 200 (application/json)

    + Attribute

        + success: true (boolean)

+ Response 400 (application/json)

    + Attribute

        + error: 登録に失敗しました。

## 分野 [/question/genre]
### 分野追加[POST]
問題分野を追加

+ Request(application/json)
	+ Attribute
		+ genre_name: 英語

+ Response 200 (application/json)

    + Attribute

        + genre_id: 1 (number)

+ Response 400 (application/json)

    + Attribute

        + error: 登録に失敗しました。

### 分野取得[GET]
登録されている分野を取得

+ Response 200 (application/json)

    + Attribute

        + genre(array)
        	+ 算数
        	+ 社会
        	+ 英語

+ Response 400 (application/json)

    + Attribute

        + error: 分野が登録されていませんでした。



