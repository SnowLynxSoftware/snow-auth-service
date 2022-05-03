package migrations

import (
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/models"
)

func Get20220427() models.DBMigrationData {
	return models.DBMigrationData{
		Name: "2022-04-27 - Initial",
		File: "migration-20220427.go",
		SQL:  sql20220427,
	}
}

var sql20220427 = `
CREATE TABLE users(
	id INT NOT NULL AUTO_INCREMENT,
	email VARCHAR(255) NOT NULL,
	hash VARCHAR(1024),
	banned BOOL NOT NULL DEFAULT 0,
	verified BOOL NOT NULL DEFAULT 0,
	archived DATETIME,
	created DATETIME DEFAULT CURRENT_TIMESTAMP,
	lastModified DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY ( id )
);
`
