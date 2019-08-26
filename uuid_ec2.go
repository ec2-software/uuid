package uuid

import (
	"errors"
)

// ImplementsGraphQLType implements part of the inteface to be a GraphQL input type.
func (UUID) ImplementsGraphQLType(name string) bool {
	return name == "ID"
}

// UnmarshalGraphQL implements part of the inteface to be a GraphQL input type.
func (id *UUID) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		var parsedID UUID
		parsedID, err = FromString(input)
		*id = UUID(parsedID)
	default:
		err = errors.New("wrong type")
	}
	return err
}

// Null converts a UUID into a Valid NullUUID
func (id UUID) Null() NullUUID {
	return NullUUID{UUID: id, Valid: true}
}

// Ptr returns a pointer to the UUID if valid, and nil if not.
func (id NullUUID) Ptr() *UUID {
	if !id.Valid {
		return nil
	}
	return &id.UUID
}
