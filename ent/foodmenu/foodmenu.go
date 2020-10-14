// Code generated by entc, DO NOT EDIT.

package foodmenu

const (
	// Label holds the string label denoting the foodmenu type in the database.
	Label = "foodmenu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFOODMENUNAME holds the string denoting the foodmenu_name field in the database.
	FieldFOODMENUNAME = "foodmenu_name"

	// EdgeFOODMENURECORD holds the string denoting the foodmenu_record edge name in mutations.
	EdgeFOODMENURECORD = "FOODMENU_RECORD"

	// Table holds the table name of the foodmenu in the database.
	Table = "foodmen_us"
	// FOODMENURECORDTable is the table the holds the FOODMENU_RECORD relation/edge.
	FOODMENURECORDTable = "recordfoods"
	// FOODMENURECORDInverseTable is the table name for the Recordfood entity.
	// It exists in this package in order to avoid circular dependency with the "recordfood" package.
	FOODMENURECORDInverseTable = "recordfoods"
	// FOODMENURECORDColumn is the table column denoting the FOODMENU_RECORD relation/edge.
	FOODMENURECORDColumn = "foodmenu_foodmenu_record"
)

// Columns holds all SQL columns for foodmenu fields.
var Columns = []string{
	FieldID,
	FieldFOODMENUNAME,
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