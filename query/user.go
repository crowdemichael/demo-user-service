package query

const (
	QueryGetUserProfile = `
		SELECT
			name,
			email
		FROM users
		WHERE id = ?;
	`
)
