package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/snowflakedb/gosnowflake"
)

var (
	elog                = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	ilog                = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	snowflake_user      = flag.String("snowflake_user", os.Getenv("SNOWFLAKE_USER"), "Snowflake user")
	snowflake_pass      = flag.String("snowflake_pass", os.Getenv("SNOWFLAKE_PASS"), "Snowflake password")
	snowflake_accountid = flag.String("snowflake_accountid", os.Getenv("SNOWFLAKE_ACCOUNTID"), "Snowflake accountid")
	snowflake_region    = flag.String("snowflake_region", os.Getenv("SNOWFLAKE_REGION"), "Snowflake region")
	proxy_host          = flag.String("proxy_host", "127.0.01", "Proxy host")
	proxy_port          = flag.String("proxy_port", "8080", "Proxy port")
)

func main() {
	flag.Parse()
	connString := fmt.Sprintf("%s:%s@%s.%s.snowflakecomputing.com?useProxy=true&proxyHost=%s&proxyPort=%s", *snowflake_user, *snowflake_pass, *snowflake_accountid, *snowflake_region, *proxy_host, *proxy_port)
	db, err := sql.Open("snowflake", connString)
	if err != nil {
		elog.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT 1 AS One")
	if err != nil {
		elog.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		elog.Fatal(err)
	}

	header := join(columns, ",")
	ilog.Printf("%s", header)
	ilog.Printf("%s", strings.Repeat("-", len(header)))

	for rows.Next() {
		var value int
		err = rows.Scan(&value)
		if err != nil {
			elog.Fatal(err)
		}
		ilog.Printf("%d", value)
	}
}

func join[T any](s []T, sep string) string {
	var ret string
	for i, v := range s {
		if i == 0 {
			ret = fmt.Sprintf("%v", v)
		} else {
			ret = fmt.Sprintf("%v%v%v", ret, sep, v)
		}
	}
	return ret
}
