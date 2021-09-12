
// (sql.ErrNoRows)查询用户等操作属于业务层面的报错，需要Wrap
package main
import (
	"database/sql"
	"github.com/pkg/errors"
)
type User struct {
	Id   int
	Name string
}
var Db *sql.DB
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:talbe@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
	    panic(err)
	}
}
func QueryUserById(id string) (User, error) {
	var user User
	err := Db.QueryRow("select id ,name from user where id = ?" ,id ).Scan(&user.Id,&user.Name)
	if err != nil{
		return user,errors.Wrap(err,"data not found")
	}
	return user,nil
}


