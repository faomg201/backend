// Code generated by entc, DO NOT EDIT.

package recordfood

const (
	// Label holds the string label denoting the recordfood type in the database.
	Label = "recordfood"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"

	// EdgeRECORDUSER holds the string denoting the record_user edge name in mutations.
	EdgeRECORDUSER = "RECORD_USER"
	// EdgeRECORDFOODMENU holds the string denoting the record_foodmenu edge name in mutations.
	EdgeRECORDFOODMENU = "RECORD_FOODMENU"
	// EdgeRECORDINGREDIENT holds the string denoting the record_ingredient edge name in mutations.
	EdgeRECORDINGREDIENT = "RECORD_INGREDIENT"
	// EdgeRECORDSOURCE holds the string denoting the record_source edge name in mutations.
	EdgeRECORDSOURCE = "RECORD_SOURCE"

	// Table holds the table name of the recordfood in the database.
	Table = "recordfoods"
	// RECORDUSERTable is the table the holds the RECORD_USER relation/edge.
	RECORDUSERTable = "recordfoods"
	// RECORDUSERInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RECORDUSERInverseTable = "users"
	// RECORDUSERColumn is the table column denoting the RECORD_USER relation/edge.
	RECORDUSERColumn = "user_user_record"
	// RECORDFOODMENUTable is the table the holds the RECORD_FOODMENU relation/edge.
	RECORDFOODMENUTable = "recordfoods"
	// RECORDFOODMENUInverseTable is the table name for the FOODMENU entity.
	// It exists in this package in order to avoid circular dependency with the "foodmenu" package.
	RECORDFOODMENUInverseTable = "foodmen_us"
	// RECORDFOODMENUColumn is the table column denoting the RECORD_FOODMENU relation/edge.
	RECORDFOODMENUColumn = "foodmenu_foodmenu_record"
	// RECORDINGREDIENTTable is the table the holds the RECORD_INGREDIENT relation/edge.
	RECORDINGREDIENTTable = "recordfoods"
	// RECORDINGREDIENTInverseTable is the table name for the Mainingre entity.
	// It exists in this package in order to avoid circular dependency with the "mainingre" package.
	RECORDINGREDIENTInverseTable = "mainingres"
	// RECORDINGREDIENTColumn is the table column denoting the RECORD_INGREDIENT relation/edge.
	RECORDINGREDIENTColumn = "mainingre_mainingre_record"
	// RECORDSOURCETable is the table the holds the RECORD_SOURCE relation/edge.
	RECORDSOURCETable = "recordfoods"
	// RECORDSOURCEInverseTable is the table name for the Source entity.
	// It exists in this package in order to avoid circular dependency with the "source" package.
	RECORDSOURCEInverseTable = "sources"
	// RECORDSOURCEColumn is the table column denoting the RECORD_SOURCE relation/edge.
	RECORDSOURCEColumn = "source_source_record"
)

// Columns holds all SQL columns for recordfood fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Recordfood type.
var ForeignKeys = []string{
	"foodmenu_foodmenu_record",
	"mainingre_mainingre_record",
	"source_source_record",
	"user_user_record",
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
