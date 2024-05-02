package main

import (
	// "database/sql"
	// "fmt"
	// "log"

	_ "github.com/lib/pq"
)

type Test struct {
	col1 int
	col2 int
	col3 int
}

/*
func main() {
	// PostgreSQL接続情報
	db, err := sql.Open("postgres", "host=postgres user=test password='password' dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データベースへの接続確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// テスト用SQL
	sql := "select * from testschema.sample;"
	pstatement, err := db.Prepare(sql)
    if err != nil {
        log.Fatal(err)
    }

	var test Test
	err = pstatement.QueryRow().Scan(&test.col1, &test.col2, &test.col3)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println(test)
}
*/

/*
備忘録

postgresログイン
psql -h localhost -U postgres

ユーザーとパスワードの設定
create role test with login password 'password';

スキーマをユーザーへ権限の付与
GRANT ALL PRIVILEGES ON SCHEMA testschema TO test;

データベース作成
create database testdb;

データベース選択
\c testdb;

テーブル作成
CREATE TABLE  testschema.sample (
  col1 VARCHAR(10),
  col2 VARCHAR(10),
  col3 VARCHAR(10),
  PRIMARY KEY (col1)
);

テーブルのselect権限付与
grant select on testschema.sample to test;

*/
