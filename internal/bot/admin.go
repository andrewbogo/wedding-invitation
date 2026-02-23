package bot

var admins = map[int64]bool{}

func SetAdmins(a map[int64]bool) {
	admins = a
}

func IsAdmin(tgID int64) bool {
	return admins[tgID]
}
