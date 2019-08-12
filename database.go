package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	defer db.Close()//关闭连接

	//验证是否可以成功连接
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

}
