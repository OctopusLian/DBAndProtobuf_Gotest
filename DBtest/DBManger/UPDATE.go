package DBManger

import "fmt"

func DBupdate(userid int){
	//更新数据
	fmt.Println("请输入要查询的用户ID：")
	fmt.Scan(&Userid)
	fmt.Println("请输入要更新的内容：a-名字；b-部门；c-进部门的时间：")
	var s string
	fmt.Scan(&s)
	switch s {
	case "a":
		fmt.Println("请输入要更改用户ID为%d的名字",Userid)
		fmt.Scan(&Username)
		fmt.Println(Username)
		fmt.Println(Userid)
		stmt,err := db.Prepare("UPDATE userif SET username=? where userid=?")
		CheckErr(err)

		res,err := stmt.Exec(Username,Userid)
		affect,err := res.RowsAffected()
		fmt.Println(affect)
	case "b":
		fmt.Println("请输入要更改用户ID为%d的部门",Userid)
		fmt.Scan(&Departname)
		stmt,err := db.Prepare("UPDATE userif SET departname=? where userid=?")
		CheckErr(err)
		res,err := stmt.Exec(Departname,Userid)
		affect,err := res.RowsAffected()
		fmt.Println(affect)
	case "c":
		fmt.Println("请输入要更改用户ID为%d的进部门的时间",Userid)
		fmt.Scan(&Created)
		stmt,err := db.Prepare("UPDATE userif SET created=? where userid=?")
		CheckErr(err)
		res,err := stmt.Exec(Created,Userid)
		affect,err := res.RowsAffected()
		fmt.Println(affect)
	default:

	}









}
