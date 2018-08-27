package DBManger

import "fmt"

func DBselect(userid int){
	//查询数据
	rows,err := db.Query("SELECT * FROM userif")
	CheckErr(err)

	for rows.Next(){
		var userid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&userid,&username,&departname,&created)
		CheckErr(err)
		fmt.Println(userid)
		fmt.Println(username)
		fmt.Println(departname)
		fmt.Println(created)
	}
}
