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
	ErrDeleteFail   = core.ErrorCode{Code: 1000, Msg: "删除失败"}
	ErrCreateFail   = core.ErrorCode{Code: 1001, Msg: "创建失败"}
	ErrUpdateFail   = core.ErrorCode{Code: 1002, Msg: "更新失败"}
	ErrAddUserRoles = core.ErrorCode{Code: 1003, Msg: "添加用户角色失败"}

	// 登录注册相关
	ErrLoginFail     = core.ErrorCode{Code: 2000, Msg: "登录失败"}
	ErrLoginInfoFail = core.ErrorCode{Code: 2001, Msg: "获取登录信息失败"}

	// 配置菜单相关
	ErrGetMenus   = core.ErrorCode{Code: 2101, Msg: "获取菜单列表失败"}
	ErrDeleteMenu = core.ErrorCode{Code: 2102, Msg: "删除菜单失败"}
	ErrCreateMenu = core.ErrorCode{Code: 2103, Msg: "创建菜单失败"}
	ErrUpdateMenu = core.ErrorCode{Code: 2104, Msg: "更新菜单失败"}

	//配置页面相关
	ErrGetPage    = core.ErrorCode{Code: 2201, Msg: "获取页面配置失败"}
	ErrGetPages   = core.ErrorCode{Code: 2202, Msg: "获取页面列表失败"}
	ErrCreatePage = core.ErrorCode{Code: 2203, Msg: "创建页面失败"}
	ErrDeletePage = core.ErrorCode{Code: 2204, Msg: "删除页面失败"}
	ErrUpdatePage = core.ErrorCode{Code: 2205, Msg: "更新页面失败"}

	//配置角色相关
	ErrGetRoles     = core.ErrorCode{Code: 2301, Msg: "获取角色列表失败"}
	ErrCreateRole   = core.ErrorCode{Code: 2302, Msg: "创建角色失败"}
	ErrDeleteRole   = core.ErrorCode{Code: 2303, Msg: "删除角色失败"}
	ErrUpdateRole   = core.ErrorCode{Code: 2304, Msg: "更新角色失败"}
	ErrGetRole      = core.ErrorCode{Code: 2305, Msg: "获取角色失败(无此角色)"}
	ErrAddPerm      = core.ErrorCode{Code: 2306, Msg: "添加角色权限失败"}
	ErrGetRolePerms = core.ErrorCode{Code: 2307, Msg: "获取角色权限列表失败"}
	ErrDelPerms     = core.ErrorCode{Code: 2308, Msg: "删除权限失败"}

	//配置权限相关
	ErrGetPerms   = core.ErrorCode{Code: 2401, Msg: "获取权限列表失败"}
	ErrCreatePerm = core.ErrorCode{Code: 2402, Msg: "创建权限失败"}
	ErrUpdatePerm = core.ErrorCode{Code: 2403, Msg: "更新权限失败"}
)
