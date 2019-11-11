# Hack Bowl2019_backend
## 概要
<p>
HackBowl2019 で利用するAPIベース実装<br>
API仕様書は SwaggerEditor に定義ファイルの内容を入力して参照してください。
</p>

SwaggerEditor: <https://editor.swagger.io> <br>
定義ファイル: `./api-document.yaml`

## 事前準備

### MySQLへDDL/DMLを反映する
以下のファイルを実行する

1. [./db/ddl.sql](./db/ddl.sql)

※ 実行時に DefaultCharacterSetをutf8にして実行すること。

(DDL(DataDefinitionLanguage)とはデータベースの構造や構成を定義するためのSQL文)<br>
(DML(DataManipulationLanguage)とはデータの管理・操作を定義するためのSQL文)

### API用のデータベースの接続情報を設定する
環境変数にデータベースの接続情報を設定します。<br>
ターミナルのセッション毎に設定したり、.bash_profileで設定を行います。

Macの場合
```
$ export MYSQL_USER=[MySQLの接続ユーザ名] \
    MYSQL_PASSWORD=[MySQLの接続パスワード] \
    MYSQL_HOST=[サーバのグローバルIP] \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=hack_bowl
```

Windowsの場合
```
$ SET MYSQL_USER=[MySQLの接続ユーザ名]
$ SET MYSQL_PASSWORD=[MySQLの接続パスワード]
$ SET MYSQL_HOST=[サーバのグローバルIP]
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=hack_bowl
```

## ローカル起動方法
```
$ go run ./cmd/main.go
```

### デプロイ方法
今回はLinuxサーバ向けにローカルでビルドを行い、<br>
ビルドされたバイナリファイルをGCEインスタンス上に配置して起動することでデプロイを行います。
#### ローカルビルド
Macの場合
```
$ GOOS=linux GOARCH=amd64 go build -o hack-bowl ./cmd/main.go
```

Windowsの場合
```
$ SET GOOS=linux
$ SET GOARCH=amd64
$ go build -o hack-bowl ./cmd/main.go
```

このコマンドの実行で `hack-bowl` という成果物を起動するバイナリファイルが生成されます。<br>
GOOS,GOARCHで「Linux用のビルド」を指定しています。
#### デプロイ
1. 秘密鍵の権限を変更する
    ```
    // 所有者のみ読み書きできる権限へ変更
    $ sudo chmod 600 [秘密鍵のパス]
    ```
2. バイナリファイルをscpでサーバ上へコピー

    Macの場合
    ```
    $ scp -i [秘密鍵のパス] [ビルドで生成されたバイナリファイルのパス] centos@[サーバのIPアドレス]:~/
    ```
    
    Windowsでteratermの場合
    
    -> ファイルをコンソール画面にドラッグアンドドロップ
    
3. コピーしたファイルに実行権限を追加する
    ```
    $ chmod +x hack-bowl 
    ```

4. バイナリを実行

    サーバにSSHログインをして以下のコマンドを実行
    ```
    // 実行するバイナリファイルに実行権限を付加する
    $ chmod +x hack-bowl
    ```
    ```
    // 実行(APIサーバの起動)
    $ MYSQL_USER=[MySQLの接続ユーザ名] \
        MYSQL_PASSWORD=[MySQLの接続パスワード] \
        MYSQL_HOST=localhost \
        MYSQL_PORT=3306 \
        MYSQL_DATABASE=hack_bowl \
        nohup ./dojo-api &
    ```