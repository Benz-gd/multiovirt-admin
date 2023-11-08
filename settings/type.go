package settings

var Conf = new(AppConfig)

type AppConfig struct {
	Name             string            `mapstructure:"name"`
	Mode             string            `mapstructure:"start_mode"`
	Version          string            `mapstructure:"version"`
	Port             int               `mapstructure:"port"`
	Locale           *Locale           `mapstructure:"locale"`
	SnowFlake        *SnowFlake        `mapstructure:"snowflake"`
	LogConfig        *LogConfig        `mapstructure:"log"`
	MySQLBase        *MySQLBase        `mapstructure:"mysqlbase"`
	RedisConfig      *RedisConfig      `mapstructure:"redis"`
	AuthConfig       *Auth             `mapstructure:"auth"`
	PostgreSQLConfig *PostgreSQLConfig `mapstructure:"postgresql"`
	MySQLCMDB        *MySQLCMDB        `mapstructure:"mysqlcmdb"`
	ZabbixConfig     *ZabbixConfig     `mapstructure:"zabbix"`
}

type Auth struct {
	Access_Token_Expire  int `mapstruct:"access_token_expire"`
	Refresh_Token_Expire int `mapstruct:"refresh_token_expire"`
}

type Locale struct {
	Locale string `mapstructure:"locale"`
}

type SnowFlake struct {
	Location  string `mapstructure:"location"`
	StartTime string `mapstructure:"starttime"`
	CenterId  int64  `mapstructure:"centerId"`
	WorkerId  int64  `mapstructure:"workerId"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLBase struct {
	Host                 string `mapstructure:"host"`
	Port                 int    `mapstructure:"port"`
	User                 string `mapstructure:"user"`
	Password             string `mapstructure:"password"`
	DBName               string `mapstructure:"dbname"`
	MysqlQuery           string `mapstructure:"mysqlquery"`
	MysqlCharset         string `mapstructure:"mysqlcharset"`
	MysqlCollation       string `mapstructure:"mysqlcollation"`
	MysqlMaxIdelConns    int    `mapstructure:"mysqlmaxidleconns"`
	MysqlMaxOpenConns    int    `mapstructure:"mysqlmaxopenconns"`
	MysqlConnMaxLifetime int    `mapstructure:"mysqlconnmaxlifetime"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type PostgreSQLConfig struct {
	Host              string `mapstructure:"host"`
	Port              int    `mapstructure:"port"`
	User              string `mapstructure:"user"`
	Password          string `mapstructure:"password"`
	DBName            string `mapstructure:"dbname"`
	TimeZone          string `mapstructure:"timezone"`
	PGMaxIdelConns    int    `mapstructure:"pgmaxidleconns"`
	PGMaxOpenConns    int    `mapstructure:"pgmaxopenconns"`
	PGConnMaxLifetime int    `mapstructure:"pgconnmaxlifetime"`
	PGPreStatement    string `mapstructure:"pgprestatement"`
}

type ZabbixConfig struct {
	Url      string `mapstructure:"url"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type MySQLCMDB struct {
	Host                 string `mapstructure:"host"`
	Port                 int    `mapstructure:"port"`
	User                 string `mapstructure:"user"`
	Password             string `mapstructure:"password"`
	DBName               string `mapstructure:"dbname"`
	MysqlQuery           string `mapstructure:"mysqlquery"`
	MysqlCharset         string `mapstructure:"mysqlcharset"`
	MysqlCollation       string `mapstructure:"mysqlcollation"`
	MysqlMaxIdelConns    int    `mapstructure:"mysqlmaxidleconns"`
	MysqlMaxOpenConns    int    `mapstructure:"mysqlmaxopenconns"`
	MysqlConnMaxLifetime int    `mapstructure:"mysqlconnmaxlifetime"`
}
