// Code generated by entc, DO NOT EDIT.

package specializedappoint

const (
	// Label holds the string label denoting the specializedappoint type in the database.
	Label = "specializedappoint"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeSpecializeddiag holds the string denoting the specializeddiag edge name in mutations.
	EdgeSpecializeddiag = "specializeddiag"

	// Table holds the table name of the specializedappoint in the database.
	Table = "specializedappoints"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "specializedappoints"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "userid"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "specializedappoints"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patientid"
	// SpecializeddiagTable is the table the holds the specializeddiag relation/edge.
	SpecializeddiagTable = "specializedappoints"
	// SpecializeddiagInverseTable is the table name for the Specializeddiag entity.
	// It exists in this package in order to avoid circular dependency with the "specializeddiag" package.
	SpecializeddiagInverseTable = "specializeddiags"
	// SpecializeddiagColumn is the table column denoting the specializeddiag relation/edge.
	SpecializeddiagColumn = "diacnosticID"
)

// Columns holds all SQL columns for specializedappoint fields.
var Columns = []string{
	FieldID,
	FieldDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Specializedappoint type.
var ForeignKeys = []string{
	"patientid",
	"diacnosticID",
	"userid",
}
