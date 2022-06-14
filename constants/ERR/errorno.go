package ERR

import "houserqu.com/tiger/core"

var (
	// 通用错误 < 1000
	NotLogin     = core.ErrorCode{Code: 402, Msg: "未登录"}
	NoPermission = core.ErrorCode{Code: 403, Msg: "无权限"}
	Param        = core.ErrorCode{Code: 400, Msg: "参数错误"}
	NotFound     = core.ErrorCode{Code: 404, Msg: "查找失败"}
	DB           = core.ErrorCode{Code: 600, Msg: "数据库错误"}
	// 具体业务错误 > 1000
	DeleteFail   = core.ErrorCode{Code: 1000, Msg: "删除失败"}
	CreateFail   = core.ErrorCode{Code: 1001, Msg: "创建失败"}
	UpdateFail   = core.ErrorCode{Code: 1002, Msg: "更新失败"}
	AddUserRoles = core.ErrorCode{Code: 1003, Msg: "添加用户角色失败"}
	DelUserRoles = core.ErrorCode{Code: 1004, Msg: "删除用户角色失败"}

	// 登录注册相关
	LoginFail     = core.ErrorCode{Code: 2000, Msg: "登录失败"}
	LoginInfoFail = core.ErrorCode{Code: 2001, Msg: "获取登录信息失败"}

	// 配置菜单相关
	GetMenus   = core.ErrorCode{Code: 2101, Msg: "获取菜单列表失败"}
	DeleteMenu = core.ErrorCode{Code: 2102, Msg: "删除菜单失败"}
	CreateMenu = core.ErrorCode{Code: 2103, Msg: "创建菜单失败"}
	UpdateMenu = core.ErrorCode{Code: 2104, Msg: "更新菜单失败"}

	//配置页面相关
	GetPage    = core.ErrorCode{Code: 2201, Msg: "获取页面配置失败"}
	GetPages   = core.ErrorCode{Code: 2202, Msg: "获取页面列表失败"}
	CreatePage = core.ErrorCode{Code: 2203, Msg: "创建页面失败"}
	DeletePage = core.ErrorCode{Code: 2204, Msg: "删除页面失败"}
	UpdatePage = core.ErrorCode{Code: 2205, Msg: "更新页面失败"}

	//配置角色相关
	GetRoles     = core.ErrorCode{Code: 2301, Msg: "获取角色列表失败"}
	CreateRole   = core.ErrorCode{Code: 2302, Msg: "创建角色失败"}
	DeleteRole   = core.ErrorCode{Code: 2303, Msg: "删除角色失败"}
	UpdateRole   = core.ErrorCode{Code: 2304, Msg: "更新角色失败"}
	GetRole      = core.ErrorCode{Code: 2305, Msg: "获取角色失败(无此角色)"}
	AddPerm      = core.ErrorCode{Code: 2306, Msg: "添加角色权限失败"}
	GetRolePerms = core.ErrorCode{Code: 2307, Msg: "获取角色权限列表失败"}
	DelPerms     = core.ErrorCode{Code: 2308, Msg: "删除权限失败"}

	//配置权限相关
	GetPerms   = core.ErrorCode{Code: 2401, Msg: "获取权限列表失败"}
	CreatePerm = core.ErrorCode{Code: 2402, Msg: "创建权限失败"}
	UpdatePerm = core.ErrorCode{Code: 2403, Msg: "更新权限失败"}
)
