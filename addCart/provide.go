package addCart

import (
	"flashsale/util"
)

// Configure DB as per configuration
func ConfigureDB(db string, dbconfig util.DBconfig) (Store, error) {
	if db == "sql" {
		return NewSQLdb(dbconfig), nil
	} else if db == "inMem" {
		return NewInMemdb(), nil
	}
	return nil, nil
}
