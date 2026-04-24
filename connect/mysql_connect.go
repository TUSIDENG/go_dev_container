package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(host.docker.internal:3306)/default"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("无法 ping 数据库: %v", err)
	}

	fmt.Println("✅ 连接成功！数据库列表：")

	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatalf("查询数据库列表失败: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var database string
		if err := rows.Scan(&database); err != nil {
			log.Fatalf("读取结果失败: %v", err)
		}
		fmt.Println(database)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("遍历结果失败: %v", err)
	}
}
