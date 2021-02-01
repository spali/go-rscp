package auth

type AuthLevel uint8

// all auth levels as constant
//go:generate enumer -type=AuthLevel -json
//nolint: golint,stylecheck
const (
	NO_AUTH   AuthLevel = 0
	USER      AuthLevel = 10
	INSTALLER AuthLevel = 20
	SERVICE   AuthLevel = 30
	ADMIN     AuthLevel = 40
	E3DC      AuthLevel = 50
	E3DC_ROOT AuthLevel = 60
)
