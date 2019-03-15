// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package content

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"strings"
	"wing/logger"
	"wing/utils"
)

// OrmUtil : provides a named orm object to query or exec sql commend
type OrmUtil struct {
	Ormer   orm.Ormer // ormer object
	ormName string    // ormer name
}

// NewOrmUtil create a OrmUtil object and than using the given name ormer
func NewOrmUtil(name string) (*OrmUtil, error) {
	ormer := orm.NewOrm()
	if err := ormer.Using(name); err != nil {
		logger.E("Using", name, "orm err:", err)
		return nil, err
	}

	u := &OrmUtil{ormName: name, Ormer: ormer}
	logger.I("Created a orm util as name:", name)
	return u, nil
}

// Query execute query sql data and fill into the given container
func (u *OrmUtil) Query(sqlstr string, container interface{}, args ...interface{}) error {
	if u.Ormer == nil {
		logger.E("Ormer", "["+u.ormName+"]", "not using!")
		return utils.ErrOrmNotUsing
	}

	if err := u.Ormer.Raw(sqlstr, args).QueryRow(container); err != nil {
		if strings.Index(err.Error(), "no row found") != -1 {
			logger.I("No row found!")
			return utils.ErrNoRowFound
		}
		logger.E("Query sql:["+sqlstr+"] err:", err)
		return err
	}
	logger.I("Query ["+sqlstr+"] args:{", args, "} retulst:{", container, "}")
	return nil
}

// Exec handle insert|update|delete sql data
func (u *OrmUtil) Exec(sqlstr string, args ...interface{}) (sql.Result, error) {
	if u.Ormer == nil {
		logger.E("Ormer", "["+u.ormName+"]", "not using!")
		return nil, utils.ErrOrmNotUsing
	}

	result, err := u.Ormer.Raw(sqlstr, args).Exec()
	if err != nil {
		logger.E("Exec sql:["+sqlstr+"] err:", err)
		return nil, err
	}
	logger.I("Executed sql:[" + sqlstr + "]")
	return result, nil
}
