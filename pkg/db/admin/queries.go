package admin

const (
	addUserMatrixByAdmin    = `call test_abac.add_user_data($1, $2, $3, $4, $5, $6, $7, $8);`
	deleteUserMatrixByAdmin = `call test_abac.delete_user_data($1);`
	updateUserMatrixByAdmin = `call test_abac.update_user_data($1, $2, $3, $4, $5, $6, $7, $8);`
)
