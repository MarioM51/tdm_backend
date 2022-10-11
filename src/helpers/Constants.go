package helpers

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Constants struct {
	Domain                 string
	SiteName               string
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
	SMTP                   SMTPCredentials
	PaymentInfo            PaymentInfomation
}

type SMTPCredentials struct {
	Host     string
	Port     int
	Username string
	Password string
}

type DBCredentials struct {
	host     string
	user     string
	pass     string
	dbname   string
	port     string
	timezone string
}

type PaymentInfomation struct {
	Clabe    string
	Owner    string
	BankName string
	Concept  string
}

func (c *Constants) LoadConstants() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	c.Env = os.Getenv("ENV")

	c.loadRoutesInfoInfo()
	c.loadDBCredentials()
	c.loadStaticInfo()
	c.loadPaymantInfo()
	c.loadSMTPInfo()

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

func (c *Constants) loadPaymantInfo() {
	c.PaymentInfo.Clabe = os.Getenv("PAYMENT_INFO_CLABE")
	c.PaymentInfo.Owner = os.Getenv("PAYMENT_INFO_OWNER")
	c.PaymentInfo.BankName = os.Getenv("PAYMENT_INFO_BANK_NAME")
	c.PaymentInfo.Concept = os.Getenv("PAYMENT_INFO_CONCEPT")
}

func (c *Constants) loadDBCredentials() {
	c.DatabaseCredentials.host = os.Getenv("DB_HOST")
	c.DatabaseCredentials.user = os.Getenv("USER")
	c.DatabaseCredentials.pass = os.Getenv("PASS")
	c.DatabaseCredentials.port = os.Getenv("DB_PORT")
	c.DatabaseCredentials.dbname = os.Getenv("DB_NAME")
	c.DatabaseCredentials.timezone = os.Getenv("TIMEZONE")
}

func (c *Constants) loadStaticInfo() {
	c.StaticResourcesVersion = os.Getenv("STATIC_FILES_VERSION")
	c.StaticFolder = os.Getenv("STATIC_FOLDER")
	c.WebComponentsFolder = os.Getenv("WEB_COMPONENTS_FOLDER")
}

func (c *Constants) loadSMTPInfo() {
	c.SMTP.Host = os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic(err)
	}
	c.SMTP.Port = smtpPort
	c.SMTP.Username = os.Getenv("SMTP_USERNAME")
	c.SMTP.Password = os.Getenv("SMTP_PASSWORD")
}

func (c *Constants) loadRoutesInfoInfo() {
	c.Domain = os.Getenv("DOMAIN")
	c.Port = os.Getenv("PORT")

	c.ApiPrefix = "/api"
	c.BackUrl = "http://" + c.Domain + ":" + c.Port
	c.UrlBlogImage = c.ApiPrefix + "/blogs/%v/image?updatedAt=%v"
	c.UrlProductImage = c.ApiPrefix + "/products/image/%v?updatedAt=%v"
}
