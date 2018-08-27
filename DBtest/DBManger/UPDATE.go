package DBManger

import "fmt"

func DBupdate(userid int){
	//更新数据
	stmt,err := db.Prepare("UPDATE userif SET username=? where userid=?")
	CheckErr(err)

	res,err := stmt.Exec("happy",20180620)
	affect,err := res.RowsAffected()

	fmt.Println(affect)
}
