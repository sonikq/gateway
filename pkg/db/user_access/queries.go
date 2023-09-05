package user_access

const (
	addUserAccess    = `call test_abac.add_user_access($1, $2, $3, $4, $5);`
	checkUserAccess  = `select test_abac.check_user_access($1, $2, $3, $4, $5);`
	deleteUserAccess = `call test_abac.delete_user_access($1);`
	updateUserAccess = `call test_abac.update_user_access($1, $2, $3, $4, $5);`
)
