package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

const (
	DB_HOSTNAME string = "hostname"
	DB_PORT     int    = 3306
	DB_USERNAME string = "db-user"
	DB_NAME     string = "test-db"

	GET_COUNT_MARIADB_ACTIVE_CONNECTIONS_SQL string = `show status where variable_name = 'threads_connected'`
)

// GetMetricMariadbActiveConnections returns number of active conections
func GetMetricMariadbActiveConnections() int {

	db, err := sql.Open("mariadb", getMariaDBInfo())

	if err != nil {
		log.Error(err, "Failed to establish connection with maria database")
	}

	defer db.Close()

	row := db.QueryRow(GET_COUNT_MARIADB_ACTIVE_CONNECTIONS_SQL)

	numConenctionsActive := 0

	err = row.Scan(&numConenctionsActive)

	if err != nil {
		log.Error(err, "Error while parsing  data")
	}

	return numConenctionsActive

}

func getMariaDBInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", DB_HOSTNAME, DB_PORT, DB_USERNAME, DB_NAME)
}
