// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/faomg201/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// USEREMAIL applies equality check predicate on the "USER_EMAIL" field. It's identical to USEREMAILEQ.
func USEREMAIL(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUSEREMAIL), v))
	})
}

// USERNAME applies equality check predicate on the "USER_NAME" field. It's identical to USERNAMEEQ.
func USERNAME(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUSERNAME), v))
	})
}

// USEREMAILEQ applies the EQ predicate on the "USER_EMAIL" field.
func USEREMAILEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILNEQ applies the NEQ predicate on the "USER_EMAIL" field.
func USEREMAILNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILIn applies the In predicate on the "USER_EMAIL" field.
func USEREMAILIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUSEREMAIL), v...))
	})
}

// USEREMAILNotIn applies the NotIn predicate on the "USER_EMAIL" field.
func USEREMAILNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUSEREMAIL), v...))
	})
}

// USEREMAILGT applies the GT predicate on the "USER_EMAIL" field.
func USEREMAILGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILGTE applies the GTE predicate on the "USER_EMAIL" field.
func USEREMAILGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILLT applies the LT predicate on the "USER_EMAIL" field.
func USEREMAILLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILLTE applies the LTE predicate on the "USER_EMAIL" field.
func USEREMAILLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILContains applies the Contains predicate on the "USER_EMAIL" field.
func USEREMAILContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILHasPrefix applies the HasPrefix predicate on the "USER_EMAIL" field.
func USEREMAILHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILHasSuffix applies the HasSuffix predicate on the "USER_EMAIL" field.
func USEREMAILHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILEqualFold applies the EqualFold predicate on the "USER_EMAIL" field.
func USEREMAILEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUSEREMAIL), v))
	})
}

// USEREMAILContainsFold applies the ContainsFold predicate on the "USER_EMAIL" field.
func USEREMAILContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUSEREMAIL), v))
	})
}

// USERNAMEEQ applies the EQ predicate on the "USER_NAME" field.
func USERNAMEEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUSERNAME), v))
	})
}

// USERNAMENEQ applies the NEQ predicate on the "USER_NAME" field.
func USERNAMENEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUSERNAME), v))
	})
}

// USERNAMEIn applies the In predicate on the "USER_NAME" field.
func USERNAMEIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUSERNAME), v...))
	})
}

// USERNAMENotIn applies the NotIn predicate on the "USER_NAME" field.
func USERNAMENotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUSERNAME), v...))
	})
}

// USERNAMEGT applies the GT predicate on the "USER_NAME" field.
func USERNAMEGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUSERNAME), v))
	})
}

// USERNAMEGTE applies the GTE predicate on the "USER_NAME" field.
func USERNAMEGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUSERNAME), v))
	})
}

// USERNAMELT applies the LT predicate on the "USER_NAME" field.
func USERNAMELT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUSERNAME), v))
	})
}

// USERNAMELTE applies the LTE predicate on the "USER_NAME" field.
func USERNAMELTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUSERNAME), v))
	})
}

// USERNAMEContains applies the Contains predicate on the "USER_NAME" field.
func USERNAMEContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUSERNAME), v))
	})
}

// USERNAMEHasPrefix applies the HasPrefix predicate on the "USER_NAME" field.
func USERNAMEHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUSERNAME), v))
	})
}

// USERNAMEHasSuffix applies the HasSuffix predicate on the "USER_NAME" field.
func USERNAMEHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUSERNAME), v))
	})
}

// USERNAMEEqualFold applies the EqualFold predicate on the "USER_NAME" field.
func USERNAMEEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUSERNAME), v))
	})
}

// USERNAMEContainsFold applies the ContainsFold predicate on the "USER_NAME" field.
func USERNAMEContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUSERNAME), v))
	})
}

// HasUSERRECORD applies the HasEdge predicate on the "USER_RECORD" edge.
func HasUSERRECORD() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(USERRECORDTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, USERRECORDTable, USERRECORDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUSERRECORDWith applies the HasEdge predicate on the "USER_RECORD" edge with a given conditions (other predicates).
func HasUSERRECORDWith(preds ...predicate.Recordfood) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(USERRECORDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, USERRECORDTable, USERRECORDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
