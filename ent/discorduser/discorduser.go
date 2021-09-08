// Code generated by entc, DO NOT EDIT.

package discorduser

const (
	// Label holds the string label denoting the discorduser type in the database.
	Label = "discord_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeVrc holds the string denoting the vrc edge name in mutations.
	EdgeVrc = "vrc"
	// Table holds the table name of the discorduser in the database.
	Table = "discord_users"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "discord_user_user"
	// VrcTable is the table that holds the vrc relation/edge.
	VrcTable = "vrc_users"
	// VrcInverseTable is the table name for the VRCUser entity.
	// It exists in this package in order to avoid circular dependency with the "vrcuser" package.
	VrcInverseTable = "vrc_users"
	// VrcColumn is the table column denoting the vrc relation/edge.
	VrcColumn = "discord_user_vrc"
)

// Columns holds all SQL columns for discorduser fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)