package main

import (
	"fmt"
	"os/exec"
)

func CreateMySqlDump(cfg Config, file string) {
	cmd := exec.Command(
		"mariadb-dump",
		fmt.Sprintf("--host=%s", cfg.DBHost),
		fmt.Sprintf("--port=%s", cfg.DBPort),
		fmt.Sprintf("--user=%s", cfg.DBUser),
		fmt.Sprintf("--password=%s", cfg.DBPass),
		"--ssl=false",
		"--single-transaction",
		"--routines",
		"--events",
		"--triggers",
		"--all-databases",
		// ignore system databases
		"--ignore-database=information_schema",
		"--ignore-database=mysql",
		"--ignore-database=performance_schema",
		"--ignore-database=sys",
		// ignore system databases
		fmt.Sprintf("--result-file=%s", file),
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		panic(fmt.Sprintf("error: %v\noutput: %s", err, output))
	}
}
