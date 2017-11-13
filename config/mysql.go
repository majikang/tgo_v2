package config


type Mysql struct{
	Mysql []MysqlConf
}
type MysqlConf struct{
	Db string
	Conn MysqlConn
}

type MysqlConn struct {
	Write MysqlBase
	Reads []MysqlBase
	Pool	MysqlPool
}
type MysqlBase struct {
	Address  string
	Port     int
	User     string
	Password string
}

type MysqlPool struct {
	Max int
	IdleMax int
	LifeTimeSeconds int
}

var (
	mysqlConfig map[string]MysqlConf
)
func init(){
	if FeatureMysql() {
		config:= &Mysql{}

		defaultMysqlConfig := configMysqlGetDefault()

		configGet("mysql", config, defaultMysqlConfig)

		if len(config.Mysql) ==0{
			panic("mysql config is empty")
		}

		mysqlConfig = make(map[string]MysqlConf)

		for _,c:= range config.Mysql{
			mysqlConfig[c.Db] = c
		}
	}
}



func configMysqlGetDefault() *Mysql {
	return &Mysql{Mysql:[]MysqlConf{MysqlConf{
		Db:"tgo",
		Conn:MysqlConn{Write:  MysqlBase{"ip", 33062, "user", "password"},
			Reads: []MysqlBase{MysqlBase{"ip", 3306, "user", "password"}},
			Pool: MysqlPool{Max:16,IdleMax:5,LifeTimeSeconds:0},}}}}
}

func MysqlGet(dbName string) MysqlConf {

	return mysqlConfig[dbName]
}

func MysqlGetAll()map[string]MysqlConf{
	return mysqlConfig
}
