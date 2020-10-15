// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/jxxshark/app/ent/specializeddiag"
	"github.com/jxxshark/app/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Useremail holds the value of the "useremail" field.
	Useremail string `json:"useremail,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Appointment holds the value of the appointment edge.
	Appointment []*Specializedappoint
	// Specializeddiag holds the value of the specializeddiag edge.
	Specializeddiag *Specializeddiag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AppointmentOrErr returns the Appointment value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AppointmentOrErr() ([]*Specializedappoint, error) {
	if e.loadedTypes[0] {
		return e.Appointment, nil
	}
	return nil, &NotLoadedError{edge: "appointment"}
}

// SpecializeddiagOrErr returns the Specializeddiag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SpecializeddiagOrErr() (*Specializeddiag, error) {
	if e.loadedTypes[1] {
		if e.Specializeddiag == nil {
			// The edge specializeddiag was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: specializeddiag.Label}
		}
		return e.Specializeddiag, nil
	}
	return nil, &NotLoadedError{edge: "specializeddiag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // username
		&sql.NullString{}, // useremail
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field username", values[0])
	} else if value.Valid {
		u.Username = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field useremail", values[1])
	} else if value.Valid {
		u.Useremail = value.String
	}
	return nil
}

// QueryAppointment queries the appointment edge of the User.
func (u *User) QueryAppointment() *SpecializedappointQuery {
	return (&UserClient{config: u.config}).QueryAppointment(u)
}

// QuerySpecializeddiag queries the specializeddiag edge of the User.
func (u *User) QuerySpecializeddiag() *SpecializeddiagQuery {
	return (&UserClient{config: u.config}).QuerySpecializeddiag(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", useremail=")
	builder.WriteString(u.Useremail)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
