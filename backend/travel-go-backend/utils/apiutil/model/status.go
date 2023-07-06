package model

type Status int

const (
	// status user
	NOT_VERIFY Status = iota
	VERIFY
	PENDING
	BLOCKED
	DEACTIVE
	ACTIVE
	WAITING
	FAILED
	SUCCESS
	FINISH
	NOT_FINISH
	NOT_VERIFY_MAIL
	DRI_ACCEPT
)

func (r Status) Name() string {
	return [...]string{
		"NOT_VERIFY", "VERIFY",
		"PENDING", "BLOCKED",
		"DEACTIVE", "ACTIVE",
		"WAITING", "FAILED",
		"SUCCESS", "FINISH",
		"NOT_FINISH", "NOT_VERIFY_MAIL",
		"DRI_ACCEPT",
	}[r]
}
