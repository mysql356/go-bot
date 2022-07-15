package pkg

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

func Pg() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertId int
	err = db.QueryRow("INSERT INTO dev.userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "manoj", "finance", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)

	fmt.Println("# Updating")
	stmt, err := db.Prepare("update dev.userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("ram", lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM dev.userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | department | created ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}

	/*
	   fmt.Println("# Deleting")
	   stmt, err = db.Prepare("delete from userinfos where uid=$1")
	   checkErr(err)

	   res, err = stmt.Exec(lastInsertId)
	   checkErr(err)

	   affect, err = res.RowsAffected()
	   checkErr(err)

	   fmt.Println(affect, "rows changed")
	*/
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
pq - A pure Go postgres driver for Go's database/sql package
go get github.com/lib/pq
https://github.com/lib/pq
 

create userinfo_uid_seq manually

CREATE TABLE dev.userinfo
(
  uid integer NOT NULL DEFAULT nextval('userinfo_uid_seq'::regclass),
  username character varying(100) NOT NULL,
  departname character varying(500) NOT NULL,
  created date,
  CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE dev.userinfo
  OWNER TO postgres;


*/
