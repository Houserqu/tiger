package PERMISSION

// TODO: 写一个脚本自动扫描所有权限并写 DB，避免需要手动创建
var (
	// 超级管理员权限
	ADMIN = "ADMIN"

	// 用户相关
	USER_ALL       = "USER_ALL"       // 所有
	USER_LIST      = "USER_LIST"      // 列表
	USER_CREATE    = "USER_CREATE"    // 创建
	USER_DELETE    = "USER_DELETE"    // 删除
	USER_UPDATE    = "USER_UPDATE"    // 更新
	USER_ADD_ROLES = "USER_ADD_ROLES" //添加用户角色
	USER_DEL_ROLES = "USER_DEL_ROLES" //删除用户角色
	// 权限管理
	AUTH_ALL = "AUTH_ALL" // 所有

	// 菜单管理
	MENU_ALL = "MENU_ALL" // 所有
)
