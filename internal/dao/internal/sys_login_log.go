// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLoginLogDao is the data access object for table sys_login_log.
type SysLoginLogDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SysLoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysLoginLogColumns defines and stores column names for table sys_login_log.
type SysLoginLogColumns struct {
	InfoId        string // 访问ID
	LoginName     string // 登录账号
	Ipaddr        string // 登录IP地址
	LoginLocation string // 登录地点
	Browser       string // 浏览器类型
	Os            string // 操作系统
	Status        string // 登录状态（0成功 1失败）
	Msg           string // 提示消息
	LoginTime     string // 登录时间
	Module        string // 登录模块
}

// sysLoginLogColumns holds the columns for table sys_login_log.
var sysLoginLogColumns = SysLoginLogColumns{
	InfoId:        "info_id",
	LoginName:     "login_name",
	Ipaddr:        "ipaddr",
	LoginLocation: "login_location",
	Browser:       "browser",
	Os:            "os",
	Status:        "status",
	Msg:           "msg",
	LoginTime:     "login_time",
	Module:        "module",
}

// NewSysLoginLogDao creates and returns a new DAO object for table data access.
func NewSysLoginLogDao() *SysLoginLogDao {
	return &SysLoginLogDao{
		group:   "default",
		table:   "sys_login_log",
		columns: sysLoginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysLoginLogDao) Columns() SysLoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
