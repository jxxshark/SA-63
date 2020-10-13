// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldUseremail holds the string denoting the useremail field in the database.
	FieldUseremail = "useremail"

	// EdgeAppointment holds the string denoting the appointment edge name in mutations.
	EdgeAppointment = "appointment"
	// EdgeSpecializeddiag holds the string denoting the specializeddiag edge name in mutations.
	EdgeSpecializeddiag = "specializeddiag"

	// Table holds the table name of the user in the database.
	Table = "users"
	// AppointmentTable is the table the holds the appointment relation/edge.
	AppointmentTable = "specializedappoints"
	// AppointmentInverseTable is the table name for the Specializedappoint entity.
	// It exists in this package in order to avoid circular dependency with the "specializedappoint" package.
	AppointmentInverseTable = "specializedappoints"
	// AppointmentColumn is the table column denoting the appointment relation/edge.
	AppointmentColumn = "userid"
	// SpecializeddiagTable is the table the holds the specializeddiag relation/edge.
	SpecializeddiagTable = "specializeddiags"
	// SpecializeddiagInverseTable is the table name for the Specializeddiag entity.
	// It exists in this package in order to avoid circular dependency with the "specializeddiag" package.
	SpecializeddiagInverseTable = "specializeddiags"
	// SpecializeddiagColumn is the table column denoting the specializeddiag relation/edge.
	SpecializeddiagColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldUseremail,
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// UseremailValidator is a validator for the "useremail" field. It is called by the builders before save.
	UseremailValidator func(string) error
)
