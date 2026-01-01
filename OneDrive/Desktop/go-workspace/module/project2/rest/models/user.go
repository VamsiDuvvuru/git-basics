package models

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() int {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	result, err := DB.Exec(query, u.Email, u.Password)
	if err != nil {
		panic(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return int(lastInsertID)
}

func (user *User) GetUsersData() ([]User, error) {
	rows, err := DB.Query("SELECT id , email, password FROM users where email = ?", user.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var userData User
		err := rows.Scan(&userData.ID, &userData.Email, &userData.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, userData)
	}
	return users, nil
}
