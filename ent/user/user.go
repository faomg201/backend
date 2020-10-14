// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUSEREMAIL holds the string denoting the user_email field in the database.
	FieldUSEREMAIL = "user_email"
	// FieldUSERNAME holds the string denoting the user_name field in the database.
	FieldUSERNAME = "user_name"

	// EdgeUSERRECORD holds the string denoting the user_record edge name in mutations.
	EdgeUSERRECORD = "USER_RECORD"

	// Table holds the table name of the user in the database.
	Table = "users"
	// USERRECORDTable is the table the holds the USER_RECORD relation/edge.
	USERRECORDTable = "recordfoods"
	// USERRECORDInverseTable is the table name for the Recordfood entity.
	// It exists in this package in order to avoid circular dependency with the "recordfood" package.
	USERRECORDInverseTable = "recordfoods"
	// USERRECORDColumn is the table column denoting the USER_RECORD relation/edge.
	USERRECORDColumn = "user_user_record"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUSEREMAIL,
	FieldUSERNAME,
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
