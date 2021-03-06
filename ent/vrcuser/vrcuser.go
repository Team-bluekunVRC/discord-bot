// Code generated by entc, DO NOT EDIT.

package vrcuser

const (
	// Label holds the string label denoting the vrcuser type in the database.
	Label = "vrc_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeDiscord holds the string denoting the discord edge name in mutations.
	EdgeDiscord = "discord"
	// Table holds the table name of the vrcuser in the database.
	Table = "vrc_users"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "vrc_users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "vrc_user_user"
	// DiscordTable is the table that holds the discord relation/edge.
	DiscordTable = "vrc_users"
	// DiscordInverseTable is the table name for the DiscordUser entity.
	// It exists in this package in order to avoid circular dependency with the "discorduser" package.
	DiscordInverseTable = "discord_users"
	// DiscordColumn is the table column denoting the discord relation/edge.
	DiscordColumn = "discord_user_vrc"
)

// Columns holds all SQL columns for vrcuser fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "vrc_users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"discord_user_vrc",
	"vrc_user_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
