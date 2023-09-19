package tool

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"gopkg.in/yaml.v3"
)

var (
	 Conf Config
	 cosClient *cos.Client
)


func init() {

	file, e := os.ReadFile("./etc/config.yaml")
	if e != nil {
		log.Fatal(e)
		return
	}

	e = yaml.Unmarshal(file, &Conf)
	if e != nil {
		log.Fatal(e)
	}

	if Conf.COS.ChunkSize < 4 {
		log.Fatal("chunk size is too small")
	}

	if d, err := time.ParseDuration(Conf.Server.Token.ExpireDuration); err != nil {
		log.Fatal(err)
	} else {
		Conf.Server.Token.expireDuration = d
	}

	u, _ := url.Parse(Conf.COS.COSBucketAddr)
	b := &cos.BaseURL{BucketURL: u}
	cosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  Conf.COS.SecretID,
			SecretKey: Conf.COS.SecretKey,
		},
	})
}


type Config struct {
	Server     Server     `yaml:"server"`
	COS        COS        `yaml:"cos"`
	DataSource DataSource `yaml:"dataSource"`
	Email Email `yaml:"email"`
}

type Server struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
	Token Token `yaml:"token"`
}

func (s Server) String() string {
	return fmt.Sprintf("%s:%s", s.Ip, s.Port)
}

type Token struct {
	ExpireDuration string `yaml:"expireDuration"`
	expireDuration time.Duration
	Issuer string `yaml:"issuer"`
}

type COS struct {
	COSBucketAddr string `yaml:"cosBucketAddr"`
	SecretID      string `yaml:"secretId"`
	SecretKey     string `yaml:"secretKey"`
	InnerPath string `yaml:"innerPath"`
	// MB
	ChunkSize int `yaml:"chunkSize"`
}

type Email struct {
	Username string `yaml:"username"`
	AuthCode string `yaml:"authCode"`
	SMTP SMTP `yaml:"smtp"`
}

type SMTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s SMTP) String() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type DataSource struct {
	SourceName string `yaml:"sourceName"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Address    string `yaml:"address"`
	Database   string `yaml:"database"`
}

func (d DataSource) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.Username, d.Password, d.Address, d.Database)
}