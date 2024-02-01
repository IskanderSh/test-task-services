package storage

const GetUserQuery = `SELECT uuid, name, surname, email, role FROM users WHERE uuid=$1`

const UpdateUserQuery = `UPDATE users SET %s`

const DeleteUserQuery = `DELETE FROM users WHERE uuid=$1`
