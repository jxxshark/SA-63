// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/jxxshark/app/ent/patient"
	"github.com/jxxshark/app/ent/specializedappoint"
	"github.com/jxxshark/app/ent/specializeddiag"
	"github.com/jxxshark/app/ent/user"
)

// Specializedappoint is the model entity for the Specializedappoint schema.
type Specializedappoint struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Date holds the value of the "Date" field.
	Date time.Time `json:"Date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpecializedappointQuery when eager-loading is set.
	Edges        SpecializedappointEdges `json:"edges"`
	patientid    *int
	diacnosticID *int
	userid       *int
}

// SpecializedappointEdges holds the relations/edges for other nodes in the graph.
type SpecializedappointEdges struct {
	// User holds the value of the user edge.
	User *User
	// Patient holds the value of the patient edge.
	Patient *Patient
	// Specializeddiag holds the value of the specializeddiag edge.
	Specializeddiag *Specializeddiag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SpecializedappointEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SpecializedappointEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[1] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// SpecializeddiagOrErr returns the Specializeddiag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SpecializedappointEdges) SpecializeddiagOrErr() (*Specializeddiag, error) {
	if e.loadedTypes[2] {
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
func (*Specializedappoint) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // Date
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Specializedappoint) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // patientid
		&sql.NullInt64{}, // diacnosticID
		&sql.NullInt64{}, // userid
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Specializedappoint fields.
func (s *Specializedappoint) assignValues(values ...interface{}) error {
	if m, n := len(values), len(specializedappoint.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field Date", values[0])
	} else if value.Valid {
		s.Date = value.Time
	}
	values = values[1:]
	if len(values) == len(specializedappoint.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field patientid", value)
		} else if value.Valid {
			s.patientid = new(int)
			*s.patientid = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field diacnosticID", value)
		} else if value.Valid {
			s.diacnosticID = new(int)
			*s.diacnosticID = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field userid", value)
		} else if value.Valid {
			s.userid = new(int)
			*s.userid = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the Specializedappoint.
func (s *Specializedappoint) QueryUser() *UserQuery {
	return (&SpecializedappointClient{config: s.config}).QueryUser(s)
}

// QueryPatient queries the patient edge of the Specializedappoint.
func (s *Specializedappoint) QueryPatient() *PatientQuery {
	return (&SpecializedappointClient{config: s.config}).QueryPatient(s)
}

// QuerySpecializeddiag queries the specializeddiag edge of the Specializedappoint.
func (s *Specializedappoint) QuerySpecializeddiag() *SpecializeddiagQuery {
	return (&SpecializedappointClient{config: s.config}).QuerySpecializeddiag(s)
}

// Update returns a builder for updating this Specializedappoint.
// Note that, you need to call Specializedappoint.Unwrap() before calling this method, if this Specializedappoint
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Specializedappoint) Update() *SpecializedappointUpdateOne {
	return (&SpecializedappointClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Specializedappoint) Unwrap() *Specializedappoint {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Specializedappoint is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Specializedappoint) String() string {
	var builder strings.Builder
	builder.WriteString("Specializedappoint(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Date=")
	builder.WriteString(s.Date.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Specializedappoints is a parsable slice of Specializedappoint.
type Specializedappoints []*Specializedappoint

func (s Specializedappoints) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
