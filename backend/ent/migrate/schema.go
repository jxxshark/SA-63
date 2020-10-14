// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "patientname", Type: field.TypeString},
		{Name: "patientage", Type: field.TypeInt},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:        "patients",
		Columns:     PatientsColumns,
		PrimaryKey:  []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SpecializedappointsColumns holds the columns for the "specializedappoints" table.
	SpecializedappointsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "patientid", Type: field.TypeInt, Nullable: true},
		{Name: "diacnosticID", Type: field.TypeInt, Nullable: true},
		{Name: "userid", Type: field.TypeInt, Nullable: true},
	}
	// SpecializedappointsTable holds the schema information for the "specializedappoints" table.
	SpecializedappointsTable = &schema.Table{
		Name:       "specializedappoints",
		Columns:    SpecializedappointsColumns,
		PrimaryKey: []*schema.Column{SpecializedappointsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "specializedappoints_patients_appointment",
				Columns: []*schema.Column{SpecializedappointsColumns[2]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "specializedappoints_specializeddiags_appointment",
				Columns: []*schema.Column{SpecializedappointsColumns[3]},

				RefColumns: []*schema.Column{SpecializeddiagsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "specializedappoints_users_appointment",
				Columns: []*schema.Column{SpecializedappointsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SpecializeddiagsColumns holds the columns for the "specializeddiags" table.
	SpecializeddiagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "specializeddiacnostictype", Type: field.TypeString},
		{Name: "specialistname", Type: field.TypeInt, Nullable: true},
	}
	// SpecializeddiagsTable holds the schema information for the "specializeddiags" table.
	SpecializeddiagsTable = &schema.Table{
		Name:       "specializeddiags",
		Columns:    SpecializeddiagsColumns,
		PrimaryKey: []*schema.Column{SpecializeddiagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "specializeddiags_users_specializeddiag",
				Columns: []*schema.Column{SpecializeddiagsColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "useremail", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PatientsTable,
		SpecializedappointsTable,
		SpecializeddiagsTable,
		UsersTable,
	}
)

func init() {
	SpecializedappointsTable.ForeignKeys[0].RefTable = PatientsTable
	SpecializedappointsTable.ForeignKeys[1].RefTable = SpecializeddiagsTable
	SpecializedappointsTable.ForeignKeys[2].RefTable = UsersTable
	SpecializeddiagsTable.ForeignKeys[0].RefTable = UsersTable
}
