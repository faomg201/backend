// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/faomg201/app/ent/source"
)

// Source is the model entity for the Source schema.
type Source struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SOURCENAME holds the value of the "SOURCE_NAME" field.
	SOURCENAME string `json:"SOURCE_NAME,omitempty"`
	// SOURCEADDRESS holds the value of the "SOURCE_ADDRESS" field.
	SOURCEADDRESS string `json:"SOURCE_ADDRESS,omitempty"`
	// SOURCETLE holds the value of the "SOURCE_TLE" field.
	SOURCETLE string `json:"SOURCE_TLE,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SourceQuery when eager-loading is set.
	Edges SourceEdges `json:"edges"`
}

// SourceEdges holds the relations/edges for other nodes in the graph.
type SourceEdges struct {
	// SOURCERECORD holds the value of the SOURCE_RECORD edge.
	SOURCERECORD []*Recordfood
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SOURCERECORDOrErr returns the SOURCERECORD value or an error if the edge
// was not loaded in eager-loading.
func (e SourceEdges) SOURCERECORDOrErr() ([]*Recordfood, error) {
	if e.loadedTypes[0] {
		return e.SOURCERECORD, nil
	}
	return nil, &NotLoadedError{edge: "SOURCE_RECORD"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Source) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // SOURCE_NAME
		&sql.NullString{}, // SOURCE_ADDRESS
		&sql.NullString{}, // SOURCE_TLE
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Source fields.
func (s *Source) assignValues(values ...interface{}) error {
	if m, n := len(values), len(source.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field SOURCE_NAME", values[0])
	} else if value.Valid {
		s.SOURCENAME = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field SOURCE_ADDRESS", values[1])
	} else if value.Valid {
		s.SOURCEADDRESS = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field SOURCE_TLE", values[2])
	} else if value.Valid {
		s.SOURCETLE = value.String
	}
	return nil
}

// QuerySOURCERECORD queries the SOURCE_RECORD edge of the Source.
func (s *Source) QuerySOURCERECORD() *RecordfoodQuery {
	return (&SourceClient{config: s.config}).QuerySOURCERECORD(s)
}

// Update returns a builder for updating this Source.
// Note that, you need to call Source.Unwrap() before calling this method, if this Source
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Source) Update() *SourceUpdateOne {
	return (&SourceClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Source) Unwrap() *Source {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Source is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Source) String() string {
	var builder strings.Builder
	builder.WriteString("Source(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", SOURCE_NAME=")
	builder.WriteString(s.SOURCENAME)
	builder.WriteString(", SOURCE_ADDRESS=")
	builder.WriteString(s.SOURCEADDRESS)
	builder.WriteString(", SOURCE_TLE=")
	builder.WriteString(s.SOURCETLE)
	builder.WriteByte(')')
	return builder.String()
}

// Sources is a parsable slice of Source.
type Sources []*Source

func (s Sources) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}