package common

import (
	"fmt"
	"strings"

	"github.com/Postcord/objects"
)

const (
	CDNBase                 = "https://cdn.discordapp.com/"
	AvatarURLFmt            = CDNBase + "avatars/%d/%s.%s"
	GuildMemberAvatarURLFmt = CDNBase + "guilds/%d/users/%d/avatars/%s.%s"
	DefaultAvatarURLFmt     = CDNBase + "embed/avatars/%s.png"
)

func getAvatarFormat(hash string) string {
	if strings.HasPrefix(hash, "a_") {
		return "gif"
	}
	return "png"
}

// Get a users avatar in the following order: guild specific, account wide, default.  Member can be nil, guild ID can be 0 if member is nil.
func GetUsersAvatar(user *objects.User, member *objects.GuildMember, guildID objects.Snowflake, size ...int) string {
	url := ""
	if member != nil && member.Avatar != "" {
		url = fmt.Sprintf(GuildMemberAvatarURLFmt, guildID, user.ID, member.Avatar, getAvatarFormat(member.Avatar))
	} else if user.Avatar != "" {
		url = fmt.Sprintf(AvatarURLFmt, user.ID, user.Avatar, getAvatarFormat(user.Avatar))
	} else {
		url = fmt.Sprintf(DefaultAvatarURLFmt, user.Discriminator)
	}
	if len(size) > 0 {
		url = fmt.Sprintf("%s?size=%d", url, size[0])
	}
	return url
}
