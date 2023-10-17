package dotenv

import (
	"fmt"
)

// APPENV get app env
func APPENV() string {
	return GetString("APP_ENV", "local")
}

// ISNFT nft checker
func ISNFT() bool {
	return GetBool("IS_NFT", false)
}

// IsAppEnvProduction app env checker
func IsAppEnvProduction() bool {
	return APPENV() == "production"
}

// ISUSEHTTPS check use https
func ISUSEHTTPS() bool {
	return GetBool("IS_USE_HTTPS", false)
}

// APPPORT get app port
func APPPORT() string {
	return GetString("APP_PORT", ":8080")
}

// APPTLSCERTFILENAME get app tls cert filename
func APPTLSCERTFILENAME() string {
	return GetString("APP_TLS_CERT_FILENAME", "")
}

// APPTLSKEYFILENAME get app tls key filename
func APPTLSKEYFILENAME() string {
	return GetString("APP_TLS_KEY_FILENAME", "")
}

// PREFIX get endpoint prefix
func PREFIX() string {
	return GetString("PREFIX", "v1")
}

// APPTIMEZONE get app timezone
func APPTIMEZONE() string {
	return GetString("APP_TIMEZONE", "Asia/Jakarta")
}

// MYSQLCONFIG get mysql connection config
func MYSQLCONFIG() string {
	dbHost := GetString("MYSQL_HOST", "localhost")
	dbPort := GetString("MYSQL_PORT", "3306")
	dbUser := GetString("MYSQL_NAME", "username")
	dbPass := GetString("MYSQL_PASS", "password")
	dbName := GetString("MYSQL_DB_NAME", "db_name")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}

// REDISURL get redis url
func REDISURL() string {
	return GetString("REDIS_URL", "")
}

// REDISPASS get redis pass
func REDISPASS() string {
	return GetString("REDIS_PASS", "")
}

// ISREDISTLS check redis tls or not
func ISREDISTLS() bool {
	return GetBool("IS_REDIS_TLS", false)
}

// REDISDB get redis db
func REDISDB() int {
	return GetInt("REDIS_DB", 0)
}

func ELASTICPATHCRT() string {
	return GetString("ELASTIC_PATH_CRT", "")
}

func ELASTICADDRESS() string {
	return GetString("ELASTIC_ADDRESS", "https://localhost:9200")
}

func ELASTICUSERNAME() string {
	return GetString("ELASTIC_USERNAME", "elastic")
}

func ELASTICPASSWORD() string {
	return GetString("ELASTIC_PASSWORD", "")
}
