package main

import (
	"database/sql"
	"time"
)

type migration struct {
	VersionId int64
	IsApplied bool
	Tstamp    time.Time
}

func getLatestMigrationInfo(db *sql.DB) (migration, error) {
	var currentMigrationInfo migration

	query := "SELECT version_id, is_applied, tstamp FROM goose_db_version ORDER BY version_id DESC LIMIT 1"
	row := db.QueryRow(query)
	err := row.Scan(&currentMigrationInfo.VersionId, &currentMigrationInfo.IsApplied, &currentMigrationInfo.Tstamp)
	if err != nil {
		return currentMigrationInfo, err
	}

	return currentMigrationInfo, nil
}
