package core

var (
	// 通用错误 < 1000
	ErrNotLogin     = ErrorCode{Code: 402, Msg: "未登录"}
	ErrNoPermission = ErrorCode{Code: 403, Msg: "无权限"}
	ErrParam        = ErrorCode{Code: 400, Msg: "参数错误"}
	ErrNotFound     = ErrorCode{Code: 404, Msg: "查找失败"}
	ErrDB           = ErrorCode{Code: 600, Msg: "数据库错误"}
	// 具体业务错误 > 1000
	ErrDeleteFail = ErrorCode{Code: 1000, Msg: "删除失败"}
	ErrCreateFail = ErrorCode{Code: 1001, Msg: "创建失败"}
	ErrUpdateFail = ErrorCode{Code: 1002, Msg: "更新失败"}

	// 登录注册相关
	ErrLoginFail     = ErrorCode{Code: 2000, Msg: "登录失败"}
	ErrLoginInfoFail = ErrorCode{Code: 2001, Msg: "获取登录信息失败"}

	// 配置相关
	ErrGetMenus   = ErrorCode{Code: 2002, Msg: "获取菜单列表失败"}
	ErrGetPage    = ErrorCode{Code: 2003, Msg: "获取页面配置失败"}
	ErrGetPages   = ErrorCode{Code: 2004, Msg: "获取页面列表失败"}
	ErrCreatePage = ErrorCode{Code: 2005, Msg: "创建页面失败"}
)
