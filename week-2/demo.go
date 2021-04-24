package week_2

import (
	"database/sql"
	errors2 "errors"
)


//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


func mockSelect() (sql.Rows, error) {
	return sql.Rows{}, sql.ErrNoRows
}

//解答
func handleIfNeedResult() (sql.Rows, error){
	res, err := mockSelect()
	if errors2.Is(err, sql.ErrNoRows) {

	} else {

	}

	return res, nil
}