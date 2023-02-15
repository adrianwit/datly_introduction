package config

type (
	Connection struct {
		Name   string
		Driver string
		DSN    string
	}

	Config struct {
		DemoDb *Connection
		AclDb  *Connection
		BqDev  *Connection
	}
)

func (c *Config) InitTest() {
	c.DemoDb = &Connection{Name: "demo", Driver: "mysql", DSN: "demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true"}
	c.AclDb = &Connection{Name: "acl", Driver: "dynamodb", DSN: "dynamodb://localhost:8000/us-west-1?key=dummy&secret=dummy"}
	c.BqDev = &Connection{Name: "bqdev", Driver: "bigquery", DSN: "bigquery://viant-e2e/bqdev"}
}
