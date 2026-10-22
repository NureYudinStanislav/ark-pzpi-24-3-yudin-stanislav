package model

type UserStatus int

const (
	StatusPending UserStatus = iota
	StatusActive
	StatusSuspended
)

const DayInSeconds = 24 * 60 * 60

func CheckAccess(status UserStatus) int {
	if status == StatusSuspended {
		return DayInSeconds
	}
	return 0
}
