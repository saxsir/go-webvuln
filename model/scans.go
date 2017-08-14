// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package model

import "database/sql"

func ScanUser(r *sql.Row) (User, error) {
	var s User
	if err := r.Scan(
		&s.Id,
		&s.Username,
	); err != nil {
		return User{}, err
	}
	return s, nil
}

func ScanUsers(rs *sql.Rows) ([]User, error) {
	structs := make([]User, 0, 16)
	var err error
	for rs.Next() {
		var s User
		if err = rs.Scan(
			&s.Id,
			&s.Username,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanSession(r *sql.Row) (Session, error) {
	var s Session
	if err := r.Scan(
		&s.SessionId,
		&s.UserId,
	); err != nil {
		return Session{}, err
	}
	return s, nil
}

func ScanSessions(rs *sql.Rows) ([]Session, error) {
	structs := make([]Session, 0, 16)
	var err error
	for rs.Next() {
		var s Session
		if err = rs.Scan(
			&s.SessionId,
			&s.UserId,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanActivity(r *sql.Row) (Activity, error) {
	var s Activity
	if err := r.Scan(
		&s.Id,
		&s.UserId,
		&s.Username,
		&s.Body,
		&s.CreatedAt,
	); err != nil {
		return Activity{}, err
	}
	return s, nil
}

func ScanActivitys(rs *sql.Rows) ([]Activity, error) {
	structs := make([]Activity, 0, 16)
	var err error
	for rs.Next() {
		var s Activity
		if err = rs.Scan(
			&s.Id,
			&s.UserId,
			&s.Username,
			&s.Body,
			&s.CreatedAt,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}
