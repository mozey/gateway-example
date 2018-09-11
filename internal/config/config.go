// Code generated with https://github.com/mozey/config DO NOT EDIT

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// APP_ACCOUNT
var account string

// APP_API
var api string

// APP_API_BASE
var apiBase string

// APP_API_DOMAIN
var apiDomain string

// APP_API_PATH
var apiPath string

// APP_API_PROXY
var apiProxy string

// APP_API_ROOT
var apiRoot string

// APP_API_STAGE_NAME
var apiStageName string

// APP_API_SUBDOMAIN
var apiSubdomain string

// APP_BOOKS_API
var booksApi string

// APP_BOOKS_BASE
var booksBase string

// APP_BOOKS_BASE_PATH
var booksBasePath string

// APP_BOOKS_STAGE_NAME
var booksStageName string

// APP_CERT_ARN
var certArn string

// APP_DEBUG
var debug string

// APP_DIR
var dir string

// APP_DNS_HOSTED_ZONE
var dnsHostedZone string

// APP_LAMBDA_ARN
var lambdaArn string

// APP_LAMBDA_BASE
var lambdaBase string

// APP_LAMBDA_HANDLER
var lambdaHandler string

// APP_LAMBDA_NAME
var lambdaName string

// APP_LAMBDA_PERM
var lambdaPerm string

// APP_LAMBDA_POLICY_ARN
var lambdaPolicyArn string

// APP_LAMBDA_ROLE_ARN
var lambdaRoleArn string

// APP_REGION
var region string

// AWS_PROFILE
var awsProfile string

// Config fields correspond to config file keys less the prefix
type Config struct {
	Account         string // APP_ACCOUNT
	Api             string // APP_API
	ApiBase         string // APP_API_BASE
	ApiDomain       string // APP_API_DOMAIN
	ApiPath         string // APP_API_PATH
	ApiProxy        string // APP_API_PROXY
	ApiRoot         string // APP_API_ROOT
	ApiStageName    string // APP_API_STAGE_NAME
	ApiSubdomain    string // APP_API_SUBDOMAIN
	BooksApi        string // APP_BOOKS_API
	BooksBase       string // APP_BOOKS_BASE
	BooksBasePath   string // APP_BOOKS_BASE_PATH
	BooksStageName  string // APP_BOOKS_STAGE_NAME
	CertArn         string // APP_CERT_ARN
	Debug           string // APP_DEBUG
	Dir             string // APP_DIR
	DnsHostedZone   string // APP_DNS_HOSTED_ZONE
	LambdaArn       string // APP_LAMBDA_ARN
	LambdaBase      string // APP_LAMBDA_BASE
	LambdaHandler   string // APP_LAMBDA_HANDLER
	LambdaName      string // APP_LAMBDA_NAME
	LambdaPerm      string // APP_LAMBDA_PERM
	LambdaPolicyArn string // APP_LAMBDA_POLICY_ARN
	LambdaRoleArn   string // APP_LAMBDA_ROLE_ARN
	Region          string // APP_REGION
	AwsProfile      string // AWS_PROFILE
}

// New creates an instance of Config.
// Build with ldflags to set the package vars.
// Env overrides package vars.
// Fields correspond to the config file keys less the prefix.
// The config file must have a flat structure
func New() *Config {
	conf := &Config{}
	SetVars(conf)
	SetEnv(conf)
	return conf
}

// SetVars sets non-empty package vars on Config
func SetVars(conf *Config) {

	if account != "" {
		conf.Account = account
	}

	if api != "" {
		conf.Api = api
	}

	if apiBase != "" {
		conf.ApiBase = apiBase
	}

	if apiDomain != "" {
		conf.ApiDomain = apiDomain
	}

	if apiPath != "" {
		conf.ApiPath = apiPath
	}

	if apiProxy != "" {
		conf.ApiProxy = apiProxy
	}

	if apiRoot != "" {
		conf.ApiRoot = apiRoot
	}

	if apiStageName != "" {
		conf.ApiStageName = apiStageName
	}

	if apiSubdomain != "" {
		conf.ApiSubdomain = apiSubdomain
	}

	if booksApi != "" {
		conf.BooksApi = booksApi
	}

	if booksBase != "" {
		conf.BooksBase = booksBase
	}

	if booksBasePath != "" {
		conf.BooksBasePath = booksBasePath
	}

	if booksStageName != "" {
		conf.BooksStageName = booksStageName
	}

	if certArn != "" {
		conf.CertArn = certArn
	}

	if debug != "" {
		conf.Debug = debug
	}

	if dir != "" {
		conf.Dir = dir
	}

	if dnsHostedZone != "" {
		conf.DnsHostedZone = dnsHostedZone
	}

	if lambdaArn != "" {
		conf.LambdaArn = lambdaArn
	}

	if lambdaBase != "" {
		conf.LambdaBase = lambdaBase
	}

	if lambdaHandler != "" {
		conf.LambdaHandler = lambdaHandler
	}

	if lambdaName != "" {
		conf.LambdaName = lambdaName
	}

	if lambdaPerm != "" {
		conf.LambdaPerm = lambdaPerm
	}

	if lambdaPolicyArn != "" {
		conf.LambdaPolicyArn = lambdaPolicyArn
	}

	if lambdaRoleArn != "" {
		conf.LambdaRoleArn = lambdaRoleArn
	}

	if region != "" {
		conf.Region = region
	}

	if awsProfile != "" {
		conf.AwsProfile = awsProfile
	}

}

// SetEnv sets non-empty env vars on Config
func SetEnv(conf *Config) {
	var v string

	v = os.Getenv("APP_ACCOUNT")
	if v != "" {
		conf.Account = v
	}

	v = os.Getenv("APP_API")
	if v != "" {
		conf.Api = v
	}

	v = os.Getenv("APP_API_BASE")
	if v != "" {
		conf.ApiBase = v
	}

	v = os.Getenv("APP_API_DOMAIN")
	if v != "" {
		conf.ApiDomain = v
	}

	v = os.Getenv("APP_API_PATH")
	if v != "" {
		conf.ApiPath = v
	}

	v = os.Getenv("APP_API_PROXY")
	if v != "" {
		conf.ApiProxy = v
	}

	v = os.Getenv("APP_API_ROOT")
	if v != "" {
		conf.ApiRoot = v
	}

	v = os.Getenv("APP_API_STAGE_NAME")
	if v != "" {
		conf.ApiStageName = v
	}

	v = os.Getenv("APP_API_SUBDOMAIN")
	if v != "" {
		conf.ApiSubdomain = v
	}

	v = os.Getenv("APP_BOOKS_API")
	if v != "" {
		conf.BooksApi = v
	}

	v = os.Getenv("APP_BOOKS_BASE")
	if v != "" {
		conf.BooksBase = v
	}

	v = os.Getenv("APP_BOOKS_BASE_PATH")
	if v != "" {
		conf.BooksBasePath = v
	}

	v = os.Getenv("APP_BOOKS_STAGE_NAME")
	if v != "" {
		conf.BooksStageName = v
	}

	v = os.Getenv("APP_CERT_ARN")
	if v != "" {
		conf.CertArn = v
	}

	v = os.Getenv("APP_DEBUG")
	if v != "" {
		conf.Debug = v
	}

	v = os.Getenv("APP_DIR")
	if v != "" {
		conf.Dir = v
	}

	v = os.Getenv("APP_DNS_HOSTED_ZONE")
	if v != "" {
		conf.DnsHostedZone = v
	}

	v = os.Getenv("APP_LAMBDA_ARN")
	if v != "" {
		conf.LambdaArn = v
	}

	v = os.Getenv("APP_LAMBDA_BASE")
	if v != "" {
		conf.LambdaBase = v
	}

	v = os.Getenv("APP_LAMBDA_HANDLER")
	if v != "" {
		conf.LambdaHandler = v
	}

	v = os.Getenv("APP_LAMBDA_NAME")
	if v != "" {
		conf.LambdaName = v
	}

	v = os.Getenv("APP_LAMBDA_PERM")
	if v != "" {
		conf.LambdaPerm = v
	}

	v = os.Getenv("APP_LAMBDA_POLICY_ARN")
	if v != "" {
		conf.LambdaPolicyArn = v
	}

	v = os.Getenv("APP_LAMBDA_ROLE_ARN")
	if v != "" {
		conf.LambdaRoleArn = v
	}

	v = os.Getenv("APP_REGION")
	if v != "" {
		conf.Region = v
	}

	v = os.Getenv("AWS_PROFILE")
	if v != "" {
		conf.AwsProfile = v
	}

}

// LoadFile sets the env from file and returns a new instance of Config
func LoadFile(mode string) (conf *Config, err error) {
	p := fmt.Sprintf(path.Join(os.Getenv("GOPATH"),
		"/src/github.com/mozey/gateway/config.%v.json"), mode)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	configMap := make(map[string]string)
	err = json.Unmarshal(b, &configMap)
	if err != nil {
		return nil, err
	}
	for key, val := range configMap {
		os.Setenv(key, val)
	}
	return New(), nil
}
