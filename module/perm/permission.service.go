package perm

func GetUserPermissions(userId int) (permissions []string, err error) {
	permissions = []string{"userList", "userInfo"}
	return
}
