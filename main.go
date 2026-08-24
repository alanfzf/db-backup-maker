package main

func main() {
	// recover from error
	defer AppRecovery()

	cfg := LoadConfig()
	// dump database to file
	file := CreateTmpFile()
	CreateMySqlDump(cfg, file)

	// compress and upload to s3
	gzFile := CompressFile(file)
	UploadDumpToS3(cfg, gzFile)

	// remove temp files
	DeleteFile(file)
	DeleteFile(gzFile)
}
