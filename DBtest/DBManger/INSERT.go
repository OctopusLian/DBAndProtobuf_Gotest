package DBManger

import(
	"fmt"
)

func DBinsert(){

	fmt.Println("请输入用户ID：")
	fmt.Scan(&Userid)
	fmt.Println("请输入名字：")
	fmt.Scan(&Username)
	fmt.Println("请输入部门：")
	fmt.Scan(&Departname)
	fmt.Println("请输入进部门时间")
	fmt.Scan(&Created)


	//插入数据
	stmt,err := db.Prepare("INSERT userif SET userid=?,username=?,departname=?,created=?")
	CheckErr(err)

	res,err := stmt.Exec(Userid,Username,Departname,Created)
	CheckErr(err)

	id,err := res.LastInsertId()

	fmt.Println(id)
}

