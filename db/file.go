package db

import (
	"fmt"
	mydb "goRedisDemo/db/mysql"
)

/**
INSERT IGNORE 与INSERT INTO的区别就是INSERT IGNORE会忽略数据库中已经存在 的数据，
如果数据库没有数据，就插入新的数据，如果有数据的话就跳过这条数据。
这样就可以保留数据库中已经存在数据，达到在间隙中插入数据的目的。
*/

/**
mysql_num_rows()返回结果集中行的数目.此命令公对select语句有效.
要取得insert , update , 或者delete查询所影响到的行的数目,用mysql_affected_rows().
*/

// OnFileUploadFinished : 文件上传完成，保存meta
func OnFileUploadFinished(filehash string, filename string,
	filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file (`file_sha1`,`file_name`,`file_size`," +
			"`file_addr`,`status`) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been uploaded before", filehash)
		}
		return true
	}
	return false
}
