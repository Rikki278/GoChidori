package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type UserRole string

const (
	RoleUser UserRole = "ROLE_USER"
	RoleAdmin UserRole = "ROLE_ADMIN"	
	RoleModerator UserRole = "ROLE_MODERATOR"
)

func (r UserRole) String() string {
	return string(r)
}

func (r UserRole) isValid() bool{
	switch r {
	case RoleUser, RoleAdmin, RoleModerator:
		return true
	default:
		return false
	}
}

// value realise interface driver.Valuer
func (r UserRole) Value() (driver.Value, error){
	if !r.isValid() {
		return nil, errors.New("invalid role")
	}
	return string(r), nil
}

func (r *UserRole) Scan(value interface{}) error {
	if value == nil {
		*r = RoleUser // default value
		return nil
	}

	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan UserRole: invalid type %T", value)
	}

	*r = UserRole(strValue)
	if !(*r).isValid() {
		return fmt.Errorf("failed to scan UserRole: invalid value %s", strValue)
	}
	return nil
	
}

func (r UserRole) isAdmin() bool {
	return r == RoleAdmin
}

func (r UserRole) isModerator() bool {
	return r == RoleModerator
}

func (r UserRole) isUser() bool {
	return r == RoleUser
}

func (r UserRole) HasPermission(requiredRole UserRole) bool {
	if r == RoleAdmin {
		return true
	}

	if r == RoleModerator && (requiredRole == RoleUser || requiredRole == RoleModerator) {
		return true
	}

	if r == RoleUser && requiredRole == RoleUser {
		return true
	}

	return false
}
