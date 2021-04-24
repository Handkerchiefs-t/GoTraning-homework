package week_2

import (
	"database/sql"
	"github.com/pkg/errors"
)


//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
//我们需要根据具体的业务来确定是否要将 sql.ErrNoRows 向上抛出
//	如果这个查询应该有结果而数据库中没有数据，这应该被看做是一个错误，应该向上抛出
//	如果这只是一个普通的查询结果，没有数据是正常现象，我们直接可以处理掉这个 error

func mockSelect() ([]string, error) {
	return []string{}, sql.ErrNoRows
}

//应该有结果，将错误上抛
func handleIfNecessaryResult() ([]string, error) {
	res, err := mockSelect()
	if errors.Is(err, sql.ErrNoRows) {
		err = errors.Wrap(err, "err: NoRows")
		return res, err
	} else if err != nil {
		err = errors.Wrap(err, "err: other error")
		return res, err
	}

	return res, nil
}

//如果不一定要有结果，
func handleIfNotNecessaryResult() ([]string, error) {
	res, err := mockSelect()
	if errors.Is(err, sql.ErrNoRows) == false {
		//log.Info("some info what you want to talk")
	} else if err != nil {
		err = errors.Wrap(err, "err: other error")
		return res, err
	}

	return res, nil
}