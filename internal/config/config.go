// Code generated with https://github.com/mozey/config DO NOT EDIT

package config

import (
	"os"
)

// APP_TIMESTAMP
var timestamp string

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

type Config struct {
	Timestamp string // APP_TIMESTAMP

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

var conf *Config

// New creates an instance of Config,
// fields are set from private package vars or OS env.
// For dev the config is read from env.
// The prod build must be compiled with ldflags to set the package vars.
// OS env vars will override ldflags if set.
// Config fields correspond to the config file keys less the prefix.
// Use https://github.com/mozey/config to manage the JSON config file
func New() *Config {
	var v string

	v = os.Getenv("APP_TIMESTAMP")
	if v != "" {
		timestamp = v
	}

	v = os.Getenv("APP_ACCOUNT")
	if v != "" {
		account = v
	}

	v = os.Getenv("APP_API")
	if v != "" {
		api = v
	}

	v = os.Getenv("APP_API_BASE")
	if v != "" {
		apiBase = v
	}

	v = os.Getenv("APP_API_DOMAIN")
	if v != "" {
		apiDomain = v
	}

	v = os.Getenv("APP_API_PATH")
	if v != "" {
		apiPath = v
	}

	v = os.Getenv("APP_API_PROXY")
	if v != "" {
		apiProxy = v
	}

	v = os.Getenv("APP_API_ROOT")
	if v != "" {
		apiRoot = v
	}

	v = os.Getenv("APP_API_STAGE_NAME")
	if v != "" {
		apiStageName = v
	}

	v = os.Getenv("APP_API_SUBDOMAIN")
	if v != "" {
		apiSubdomain = v
	}

	v = os.Getenv("APP_BOOKS_API")
	if v != "" {
		booksApi = v
	}

	v = os.Getenv("APP_BOOKS_BASE")
	if v != "" {
		booksBase = v
	}

	v = os.Getenv("APP_BOOKS_BASE_PATH")
	if v != "" {
		booksBasePath = v
	}

	v = os.Getenv("APP_BOOKS_STAGE_NAME")
	if v != "" {
		booksStageName = v
	}

	v = os.Getenv("APP_CERT_ARN")
	if v != "" {
		certArn = v
	}

	v = os.Getenv("APP_DEBUG")
	if v != "" {
		debug = v
	}

	v = os.Getenv("APP_DIR")
	if v != "" {
		dir = v
	}

	v = os.Getenv("APP_DNS_HOSTED_ZONE")
	if v != "" {
		dnsHostedZone = v
	}

	v = os.Getenv("APP_LAMBDA_ARN")
	if v != "" {
		lambdaArn = v
	}

	v = os.Getenv("APP_LAMBDA_BASE")
	if v != "" {
		lambdaBase = v
	}

	v = os.Getenv("APP_LAMBDA_HANDLER")
	if v != "" {
		lambdaHandler = v
	}

	v = os.Getenv("APP_LAMBDA_NAME")
	if v != "" {
		lambdaName = v
	}

	v = os.Getenv("APP_LAMBDA_PERM")
	if v != "" {
		lambdaPerm = v
	}

	v = os.Getenv("APP_LAMBDA_POLICY_ARN")
	if v != "" {
		lambdaPolicyArn = v
	}

	v = os.Getenv("APP_LAMBDA_ROLE_ARN")
	if v != "" {
		lambdaRoleArn = v
	}

	v = os.Getenv("APP_REGION")
	if v != "" {
		region = v
	}

	v = os.Getenv("AWS_PROFILE")
	if v != "" {
		awsProfile = v
	}

	conf = &Config{
		Timestamp: timestamp,

		Account:         account,
		Api:             api,
		ApiBase:         apiBase,
		ApiDomain:       apiDomain,
		ApiPath:         apiPath,
		ApiProxy:        apiProxy,
		ApiRoot:         apiRoot,
		ApiStageName:    apiStageName,
		ApiSubdomain:    apiSubdomain,
		BooksApi:        booksApi,
		BooksBase:       booksBase,
		BooksBasePath:   booksBasePath,
		BooksStageName:  booksStageName,
		CertArn:         certArn,
		Debug:           debug,
		Dir:             dir,
		DnsHostedZone:   dnsHostedZone,
		LambdaArn:       lambdaArn,
		LambdaBase:      lambdaBase,
		LambdaHandler:   lambdaHandler,
		LambdaName:      lambdaName,
		LambdaPerm:      lambdaPerm,
		LambdaPolicyArn: lambdaPolicyArn,
		LambdaRoleArn:   lambdaRoleArn,
		Region:          region,
		AwsProfile:      awsProfile,
	}

	return conf
}

// Refresh returns a new Config if the Timestamp has changed
func Refresh() *Config {
	if conf == nil {
		// conf not initialised
		return New()
	}

	timestamp = os.Getenv("APP_TIMESTAMP")
	if conf.Timestamp != timestamp {
		// Timestamp changed, reload config
		return New()
	}

	// No change
	return conf
}
