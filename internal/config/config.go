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

// APP_PORT
var port string

// APP_REGION
var region string

// APP_VERSION
var version string

// AWS_PROFILE
var awsProfile string

// Config fields correspond to config file keys less the prefix
type Config struct {
	account         string // APP_ACCOUNT
	api             string // APP_API
	apiBase         string // APP_API_BASE
	apiDomain       string // APP_API_DOMAIN
	apiPath         string // APP_API_PATH
	apiProxy        string // APP_API_PROXY
	apiRoot         string // APP_API_ROOT
	apiStageName    string // APP_API_STAGE_NAME
	apiSubdomain    string // APP_API_SUBDOMAIN
	booksApi        string // APP_BOOKS_API
	booksBase       string // APP_BOOKS_BASE
	booksBasePath   string // APP_BOOKS_BASE_PATH
	booksStageName  string // APP_BOOKS_STAGE_NAME
	certArn         string // APP_CERT_ARN
	debug           string // APP_DEBUG
	dir             string // APP_DIR
	dnsHostedZone   string // APP_DNS_HOSTED_ZONE
	lambdaArn       string // APP_LAMBDA_ARN
	lambdaBase      string // APP_LAMBDA_BASE
	lambdaHandler   string // APP_LAMBDA_HANDLER
	lambdaName      string // APP_LAMBDA_NAME
	lambdaPerm      string // APP_LAMBDA_PERM
	lambdaPolicyArn string // APP_LAMBDA_POLICY_ARN
	lambdaRoleArn   string // APP_LAMBDA_ROLE_ARN
	port            string // APP_PORT
	region          string // APP_REGION
	version         string // APP_VERSION
	awsProfile      string // AWS_PROFILE
}

// Account is APP_ACCOUNT
func (c *Config) Account() string {
	return c.account
}

// Api is APP_API
func (c *Config) Api() string {
	return c.api
}

// ApiBase is APP_API_BASE
func (c *Config) ApiBase() string {
	return c.apiBase
}

// ApiDomain is APP_API_DOMAIN
func (c *Config) ApiDomain() string {
	return c.apiDomain
}

// ApiPath is APP_API_PATH
func (c *Config) ApiPath() string {
	return c.apiPath
}

// ApiProxy is APP_API_PROXY
func (c *Config) ApiProxy() string {
	return c.apiProxy
}

// ApiRoot is APP_API_ROOT
func (c *Config) ApiRoot() string {
	return c.apiRoot
}

// ApiStageName is APP_API_STAGE_NAME
func (c *Config) ApiStageName() string {
	return c.apiStageName
}

// ApiSubdomain is APP_API_SUBDOMAIN
func (c *Config) ApiSubdomain() string {
	return c.apiSubdomain
}

// BooksApi is APP_BOOKS_API
func (c *Config) BooksApi() string {
	return c.booksApi
}

// BooksBase is APP_BOOKS_BASE
func (c *Config) BooksBase() string {
	return c.booksBase
}

// BooksBasePath is APP_BOOKS_BASE_PATH
func (c *Config) BooksBasePath() string {
	return c.booksBasePath
}

// BooksStageName is APP_BOOKS_STAGE_NAME
func (c *Config) BooksStageName() string {
	return c.booksStageName
}

// CertArn is APP_CERT_ARN
func (c *Config) CertArn() string {
	return c.certArn
}

// Debug is APP_DEBUG
func (c *Config) Debug() string {
	return c.debug
}

// Dir is APP_DIR
func (c *Config) Dir() string {
	return c.dir
}

// DnsHostedZone is APP_DNS_HOSTED_ZONE
func (c *Config) DnsHostedZone() string {
	return c.dnsHostedZone
}

// LambdaArn is APP_LAMBDA_ARN
func (c *Config) LambdaArn() string {
	return c.lambdaArn
}

// LambdaBase is APP_LAMBDA_BASE
func (c *Config) LambdaBase() string {
	return c.lambdaBase
}

// LambdaHandler is APP_LAMBDA_HANDLER
func (c *Config) LambdaHandler() string {
	return c.lambdaHandler
}

// LambdaName is APP_LAMBDA_NAME
func (c *Config) LambdaName() string {
	return c.lambdaName
}

// LambdaPerm is APP_LAMBDA_PERM
func (c *Config) LambdaPerm() string {
	return c.lambdaPerm
}

// LambdaPolicyArn is APP_LAMBDA_POLICY_ARN
func (c *Config) LambdaPolicyArn() string {
	return c.lambdaPolicyArn
}

// LambdaRoleArn is APP_LAMBDA_ROLE_ARN
func (c *Config) LambdaRoleArn() string {
	return c.lambdaRoleArn
}

// Port is APP_PORT
func (c *Config) Port() string {
	return c.port
}

// Region is APP_REGION
func (c *Config) Region() string {
	return c.region
}

// Version is APP_VERSION
func (c *Config) Version() string {
	return c.version
}

// AwsProfile is AWS_PROFILE
func (c *Config) AwsProfile() string {
	return c.awsProfile
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
		conf.account = account
	}

	if api != "" {
		conf.api = api
	}

	if apiBase != "" {
		conf.apiBase = apiBase
	}

	if apiDomain != "" {
		conf.apiDomain = apiDomain
	}

	if apiPath != "" {
		conf.apiPath = apiPath
	}

	if apiProxy != "" {
		conf.apiProxy = apiProxy
	}

	if apiRoot != "" {
		conf.apiRoot = apiRoot
	}

	if apiStageName != "" {
		conf.apiStageName = apiStageName
	}

	if apiSubdomain != "" {
		conf.apiSubdomain = apiSubdomain
	}

	if booksApi != "" {
		conf.booksApi = booksApi
	}

	if booksBase != "" {
		conf.booksBase = booksBase
	}

	if booksBasePath != "" {
		conf.booksBasePath = booksBasePath
	}

	if booksStageName != "" {
		conf.booksStageName = booksStageName
	}

	if certArn != "" {
		conf.certArn = certArn
	}

	if debug != "" {
		conf.debug = debug
	}

	if dir != "" {
		conf.dir = dir
	}

	if dnsHostedZone != "" {
		conf.dnsHostedZone = dnsHostedZone
	}

	if lambdaArn != "" {
		conf.lambdaArn = lambdaArn
	}

	if lambdaBase != "" {
		conf.lambdaBase = lambdaBase
	}

	if lambdaHandler != "" {
		conf.lambdaHandler = lambdaHandler
	}

	if lambdaName != "" {
		conf.lambdaName = lambdaName
	}

	if lambdaPerm != "" {
		conf.lambdaPerm = lambdaPerm
	}

	if lambdaPolicyArn != "" {
		conf.lambdaPolicyArn = lambdaPolicyArn
	}

	if lambdaRoleArn != "" {
		conf.lambdaRoleArn = lambdaRoleArn
	}

	if port != "" {
		conf.port = port
	}

	if region != "" {
		conf.region = region
	}

	if version != "" {
		conf.version = version
	}

	if awsProfile != "" {
		conf.awsProfile = awsProfile
	}

}

// SetEnv sets non-empty env vars on Config
func SetEnv(conf *Config) {
	var v string

	v = os.Getenv("APP_ACCOUNT")
	if v != "" {
		conf.account = v
	}

	v = os.Getenv("APP_API")
	if v != "" {
		conf.api = v
	}

	v = os.Getenv("APP_API_BASE")
	if v != "" {
		conf.apiBase = v
	}

	v = os.Getenv("APP_API_DOMAIN")
	if v != "" {
		conf.apiDomain = v
	}

	v = os.Getenv("APP_API_PATH")
	if v != "" {
		conf.apiPath = v
	}

	v = os.Getenv("APP_API_PROXY")
	if v != "" {
		conf.apiProxy = v
	}

	v = os.Getenv("APP_API_ROOT")
	if v != "" {
		conf.apiRoot = v
	}

	v = os.Getenv("APP_API_STAGE_NAME")
	if v != "" {
		conf.apiStageName = v
	}

	v = os.Getenv("APP_API_SUBDOMAIN")
	if v != "" {
		conf.apiSubdomain = v
	}

	v = os.Getenv("APP_BOOKS_API")
	if v != "" {
		conf.booksApi = v
	}

	v = os.Getenv("APP_BOOKS_BASE")
	if v != "" {
		conf.booksBase = v
	}

	v = os.Getenv("APP_BOOKS_BASE_PATH")
	if v != "" {
		conf.booksBasePath = v
	}

	v = os.Getenv("APP_BOOKS_STAGE_NAME")
	if v != "" {
		conf.booksStageName = v
	}

	v = os.Getenv("APP_CERT_ARN")
	if v != "" {
		conf.certArn = v
	}

	v = os.Getenv("APP_DEBUG")
	if v != "" {
		conf.debug = v
	}

	v = os.Getenv("APP_DIR")
	if v != "" {
		conf.dir = v
	}

	v = os.Getenv("APP_DNS_HOSTED_ZONE")
	if v != "" {
		conf.dnsHostedZone = v
	}

	v = os.Getenv("APP_LAMBDA_ARN")
	if v != "" {
		conf.lambdaArn = v
	}

	v = os.Getenv("APP_LAMBDA_BASE")
	if v != "" {
		conf.lambdaBase = v
	}

	v = os.Getenv("APP_LAMBDA_HANDLER")
	if v != "" {
		conf.lambdaHandler = v
	}

	v = os.Getenv("APP_LAMBDA_NAME")
	if v != "" {
		conf.lambdaName = v
	}

	v = os.Getenv("APP_LAMBDA_PERM")
	if v != "" {
		conf.lambdaPerm = v
	}

	v = os.Getenv("APP_LAMBDA_POLICY_ARN")
	if v != "" {
		conf.lambdaPolicyArn = v
	}

	v = os.Getenv("APP_LAMBDA_ROLE_ARN")
	if v != "" {
		conf.lambdaRoleArn = v
	}

	v = os.Getenv("APP_PORT")
	if v != "" {
		conf.port = v
	}

	v = os.Getenv("APP_REGION")
	if v != "" {
		conf.region = v
	}

	v = os.Getenv("APP_VERSION")
	if v != "" {
		conf.version = v
	}

	v = os.Getenv("AWS_PROFILE")
	if v != "" {
		conf.awsProfile = v
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
