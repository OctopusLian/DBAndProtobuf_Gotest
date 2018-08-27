package DBManger

import "fmt"

func DBdelete(userid int){
	//删除数据
	stmt,err := db.Prepare("DELETE from userif where userid=?")
	CheckErr(err)

	res,err := stmt.Exec(20180651)
	CheckErr(err)

	affect,err := res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)
}
