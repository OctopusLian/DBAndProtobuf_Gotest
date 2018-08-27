package DBManger

import "fmt"

func DBselect(){
	//查询数据
	rows,err := db.Query("SELECT * FROM userif")
	CheckErr(err)

	for rows.Next(){
		/*
		var userid int
		var username string
		var departname string
		var created string
		*/
		err = rows.Scan(&Id,&Userid,&Created,&Departname,&Username)
		CheckErr(err)
		fmt.Println(Userid)
		fmt.Println(Username)
		fmt.Println(Departname)
		fmt.Println(Created)
	}
}
