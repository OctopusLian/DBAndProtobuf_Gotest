package DBManger

import(
	"fmt"
)

func DBinsert(){
	var username,departname,created string
	var userid int
	fmt.Println("请输入用户ID：")
	fmt.Scan(&userid)
	fmt.Println("请输入名字：")
	fmt.Scan(&username)
	fmt.Println("请输入部门：")
	fmt.Scan(&departname)
	fmt.Println("请输入进部门时间")
	fmt.Scan(&created)


	//插入数据
	stmt,err := db.Prepare("INSERT userif SET userid=?,username=?,departname=?,created=?")
	CheckErr(err)

	res,err := stmt.Exec(userid,username,departname,created)
	CheckErr(err)

	id,err := res.LastInsertId()

	fmt.Println(id)
}

