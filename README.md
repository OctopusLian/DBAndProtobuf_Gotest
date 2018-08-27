## 使用Go语言操作mysql数据库和对Protobuf的运用  
更新日志：  
- 2018.8.27：完成数据库的增删改查逻辑操作，插入有一处bug，查询更新和删除还处于测试阶段。  
- 2018.8.27晚：  
解决数据库的查询bug，原因在于mysql输出数据的列顺序与逻辑上的顺序不符，以mysql输出顺序为准；  
在更新、删除和关闭数据库的逻辑上遇到了同一类型的bug_panic: runtime error: invalid memory address or nil pointer dereference；
  
  