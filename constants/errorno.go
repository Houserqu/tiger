package constants

import "houserqu.com/tiger/core"

var (
	// 通用错误 < 1000
	ErrNotLogin     = core.ErrorCode{Code: 402, Msg: "未登录"}
	ErrNoPermission = core.ErrorCode{Code: 403, Msg: "无权限"}
	ErrParam        = core.ErrorCode{Code: 400, Msg: "参数错误"}
	ErrNotFound     = core.ErrorCode{Code: 404, Msg: "查找失败"}
	ErrDB           = core.ErrorCode{Code: 600, Msg: "数据库错误"}
	// 具体业务错误 > 1000
	ErrDeleteFail = core.ErrorCode{Code: 1000, Msg: "删除失败"}
	ErrCreateFail = core.ErrorCode{Code: 1001, Msg: "创建失败"}
	ErrUpdateFail = core.ErrorCode{Code: 1002, Msg: "更新失败"}

	// 登录注册相关
	ErrLoginFail     = core.ErrorCode{Code: 2000, Msg: "登录失败"}
	ErrLoginInfoFail = core.ErrorCode{Code: 2001, Msg: "获取登录信息失败"}

	// 配置菜单相关
	ErrGetMenus   = core.ErrorCode{Code: 2002, Msg: "获取菜单列表失败"}
	ErrDeleteMenu = core.ErrorCode{Code: 2008, Msg: "删除菜单失败"}
	ErrCreateMenu = core.ErrorCode{Code: 2009, Msg: "创建菜单失败"}
	ErrUpdateMenu = core.ErrorCode{Code: 2010, Msg: "更新菜单失败"}

	//配置页面相关
	ErrGetPage    = core.ErrorCode{Code: 2003, Msg: "获取页面配置失败"}
	ErrGetPages   = core.ErrorCode{Code: 2004, Msg: "获取页面列表失败"}
	ErrCreatePage = core.ErrorCode{Code: 2005, Msg: "创建页面失败"}
	ErrDeletePage = core.ErrorCode{Code: 2006, Msg: "删除页面失败"}
	ErrUpdatePage = core.ErrorCode{Code: 2007, Msg: "更新页面失败"}
)
