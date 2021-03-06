// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Team-bluekunVRC/discord-bot/ent/discorduser"
)

// DiscordUser is the model entity for the DiscordUser schema.
type DiscordUser struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiscordUserQuery when eager-loading is set.
	Edges DiscordUserEdges `json:"edges"`
}

// DiscordUserEdges holds the relations/edges for other nodes in the graph.
type DiscordUserEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Vrc holds the value of the vrc edge.
	Vrc []*VRCUser `json:"vrc,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordUserEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// VrcOrErr returns the Vrc value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordUserEdges) VrcOrErr() ([]*VRCUser, error) {
	if e.loadedTypes[1] {
		return e.Vrc, nil
	}
	return nil, &NotLoadedError{edge: "vrc"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DiscordUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case discorduser.FieldID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DiscordUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DiscordUser fields.
func (du *DiscordUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case discorduser.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				du.ID = value.String
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the DiscordUser entity.
func (du *DiscordUser) QueryUser() *UserQuery {
	return (&DiscordUserClient{config: du.config}).QueryUser(du)
}

// QueryVrc queries the "vrc" edge of the DiscordUser entity.
func (du *DiscordUser) QueryVrc() *VRCUserQuery {
	return (&DiscordUserClient{config: du.config}).QueryVrc(du)
}

// Update returns a builder for updating this DiscordUser.
// Note that you need to call DiscordUser.Unwrap() before calling this method if this DiscordUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (du *DiscordUser) Update() *DiscordUserUpdateOne {
	return (&DiscordUserClient{config: du.config}).UpdateOne(du)
}

// Unwrap unwraps the DiscordUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (du *DiscordUser) Unwrap() *DiscordUser {
	tx, ok := du.config.driver.(*txDriver)
	if !ok {
		panic("ent: DiscordUser is not a transactional entity")
	}
	du.config.driver = tx.drv
	return du
}

// String implements the fmt.Stringer.
func (du *DiscordUser) String() string {
	var builder strings.Builder
	builder.WriteString("DiscordUser(")
	builder.WriteString(fmt.Sprintf("id=%v", du.ID))
	builder.WriteByte(')')
	return builder.String()
}

// DiscordUsers is a parsable slice of DiscordUser.
type DiscordUsers []*DiscordUser

func (du DiscordUsers) config(cfg config) {
	for _i := range du {
		du[_i].config = cfg
	}
}
