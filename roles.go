package common

import "github.com/Postcord/objects"

func HasRole(user *objects.GuildMember, role *objects.Role) bool {
	for _, memberRole := range user.Roles {
		if memberRole == role.ID {
			return true
		}
	}
	return false
}

func GetHighestRole(user *objects.GuildMember, roles []*objects.Role) *objects.Role {
	highestRole := &objects.Role{}
	for _, role := range roles {
		if HasRole(user, role) && role.Position > highestRole.Position {
			highestRole = role
		}
	}
	return highestRole
}
