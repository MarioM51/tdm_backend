package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Constants struct {
	Domain                 string
	Port                   string
	ApiPrefix              string
	Env                    string
	BackUrl                string
	UrlBlogImage           string
	UrlProductImage        string
	DatabaseCredentials    DBCredentials
	StaticResourcesVersion string
	StaticFolder           string
	WebComponentsFolder    string
}

type DBCredentials struct {
	host     string
	user     string
	pass     string
	dbname   string
	port     string
	timezone string
}

func (c *Constants) LoadConstants() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	c.Env = os.Getenv("ENV")

	c.Domain = os.Getenv("DOMAIN")
	c.Port = os.Getenv("PORT")

	c.ApiPrefix = "/api"
	c.BackUrl = "http://" + c.Domain + ":" + c.Port
	c.UrlBlogImage = c.ApiPrefix + "/blogs/%v/image?updatedAt=%v"
	c.UrlProductImage = c.ApiPrefix + "/products/image/%v?updatedAt=%v"

	c.DatabaseCredentials.host = os.Getenv("DB_HOST")
	c.DatabaseCredentials.user = os.Getenv("USER")
	c.DatabaseCredentials.pass = os.Getenv("PASS")
	c.DatabaseCredentials.port = os.Getenv("DB_PORT")
	c.DatabaseCredentials.dbname = os.Getenv("DB_NAME")
	c.DatabaseCredentials.timezone = os.Getenv("TIMEZONE")

	c.StaticResourcesVersion = os.Getenv("STATIC_FILES_VERSION")
	c.StaticFolder = os.Getenv("STATIC_FOLDER")
	c.WebComponentsFolder = os.Getenv("WEB_COMPONENTS_FOLDER")
}

func (c *Constants) IsProduction() bool {
	return c.Env == "PRODUCTION"
}

func (c *Constants) IsLocal() bool {
	return c.Env == "local"
}

func (*Constants) AvalaibleEnviroments() string {
	return "local, PRODUCTION"
}
