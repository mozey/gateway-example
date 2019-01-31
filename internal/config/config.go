
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
// APP_CERT_ARN
var certArn string
// APP_DEBUG
var debug string
// APP_DIR
var dir string
// APP_DNS_HOSTED_ZONE
var dnsHostedZone string
// APP_GW_BASE_API
var gwBaseApi string
// APP_GW_DOMAIN
var gwDomain string
// APP_GW_ID_API
var gwIdApi string
// APP_GW_PATH_API
var gwPathApi string
// APP_GW_PROXY_API
var gwProxyApi string
// APP_GW_ROOT_API
var gwRootApi string
// APP_GW_STAGE_NAME_API
var gwStageNameApi string
// APP_GW_SUBDOMAIN
var gwSubdomain string
// APP_LAMBDA_ARN_API
var lambdaArnApi string
// APP_LAMBDA_BASE_API
var lambdaBaseApi string
// APP_LAMBDA_HANDLER_API
var lambdaHandlerApi string
// APP_LAMBDA_NAME_API
var lambdaNameApi string
// APP_LAMBDA_PERM_API
var lambdaPermApi string
// APP_LAMBDA_POLICY_ARN
var lambdaPolicyArn string
// APP_LAMBDA_ROLE_ARN_API
var lambdaRoleArnApi string
// APP_PORT_API
var portApi string
// APP_PORT_CONSOLE
var portConsole string
// APP_REGION
var region string
// APP_VERSION_API
var versionApi string
// AWS_PROFILE
var awsProfile string


// Config fields correspond to config file keys less the prefix
type Config struct {
	
	account string // APP_ACCOUNT
	certArn string // APP_CERT_ARN
	debug string // APP_DEBUG
	dir string // APP_DIR
	dnsHostedZone string // APP_DNS_HOSTED_ZONE
	gwBaseApi string // APP_GW_BASE_API
	gwDomain string // APP_GW_DOMAIN
	gwIdApi string // APP_GW_ID_API
	gwPathApi string // APP_GW_PATH_API
	gwProxyApi string // APP_GW_PROXY_API
	gwRootApi string // APP_GW_ROOT_API
	gwStageNameApi string // APP_GW_STAGE_NAME_API
	gwSubdomain string // APP_GW_SUBDOMAIN
	lambdaArnApi string // APP_LAMBDA_ARN_API
	lambdaBaseApi string // APP_LAMBDA_BASE_API
	lambdaHandlerApi string // APP_LAMBDA_HANDLER_API
	lambdaNameApi string // APP_LAMBDA_NAME_API
	lambdaPermApi string // APP_LAMBDA_PERM_API
	lambdaPolicyArn string // APP_LAMBDA_POLICY_ARN
	lambdaRoleArnApi string // APP_LAMBDA_ROLE_ARN_API
	portApi string // APP_PORT_API
	portConsole string // APP_PORT_CONSOLE
	region string // APP_REGION
	versionApi string // APP_VERSION_API
	awsProfile string // AWS_PROFILE
}


// Account is APP_ACCOUNT
func (c *Config) Account() string {
	return c.account
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
// GwBaseApi is APP_GW_BASE_API
func (c *Config) GwBaseApi() string {
	return c.gwBaseApi
}
// GwDomain is APP_GW_DOMAIN
func (c *Config) GwDomain() string {
	return c.gwDomain
}
// GwIdApi is APP_GW_ID_API
func (c *Config) GwIdApi() string {
	return c.gwIdApi
}
// GwPathApi is APP_GW_PATH_API
func (c *Config) GwPathApi() string {
	return c.gwPathApi
}
// GwProxyApi is APP_GW_PROXY_API
func (c *Config) GwProxyApi() string {
	return c.gwProxyApi
}
// GwRootApi is APP_GW_ROOT_API
func (c *Config) GwRootApi() string {
	return c.gwRootApi
}
// GwStageNameApi is APP_GW_STAGE_NAME_API
func (c *Config) GwStageNameApi() string {
	return c.gwStageNameApi
}
// GwSubdomain is APP_GW_SUBDOMAIN
func (c *Config) GwSubdomain() string {
	return c.gwSubdomain
}
// LambdaArnApi is APP_LAMBDA_ARN_API
func (c *Config) LambdaArnApi() string {
	return c.lambdaArnApi
}
// LambdaBaseApi is APP_LAMBDA_BASE_API
func (c *Config) LambdaBaseApi() string {
	return c.lambdaBaseApi
}
// LambdaHandlerApi is APP_LAMBDA_HANDLER_API
func (c *Config) LambdaHandlerApi() string {
	return c.lambdaHandlerApi
}
// LambdaNameApi is APP_LAMBDA_NAME_API
func (c *Config) LambdaNameApi() string {
	return c.lambdaNameApi
}
// LambdaPermApi is APP_LAMBDA_PERM_API
func (c *Config) LambdaPermApi() string {
	return c.lambdaPermApi
}
// LambdaPolicyArn is APP_LAMBDA_POLICY_ARN
func (c *Config) LambdaPolicyArn() string {
	return c.lambdaPolicyArn
}
// LambdaRoleArnApi is APP_LAMBDA_ROLE_ARN_API
func (c *Config) LambdaRoleArnApi() string {
	return c.lambdaRoleArnApi
}
// PortApi is APP_PORT_API
func (c *Config) PortApi() string {
	return c.portApi
}
// PortConsole is APP_PORT_CONSOLE
func (c *Config) PortConsole() string {
	return c.portConsole
}
// Region is APP_REGION
func (c *Config) Region() string {
	return c.region
}
// VersionApi is APP_VERSION_API
func (c *Config) VersionApi() string {
	return c.versionApi
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
	
	if gwBaseApi != "" {
		conf.gwBaseApi = gwBaseApi
	}
	
	if gwDomain != "" {
		conf.gwDomain = gwDomain
	}
	
	if gwIdApi != "" {
		conf.gwIdApi = gwIdApi
	}
	
	if gwPathApi != "" {
		conf.gwPathApi = gwPathApi
	}
	
	if gwProxyApi != "" {
		conf.gwProxyApi = gwProxyApi
	}
	
	if gwRootApi != "" {
		conf.gwRootApi = gwRootApi
	}
	
	if gwStageNameApi != "" {
		conf.gwStageNameApi = gwStageNameApi
	}
	
	if gwSubdomain != "" {
		conf.gwSubdomain = gwSubdomain
	}
	
	if lambdaArnApi != "" {
		conf.lambdaArnApi = lambdaArnApi
	}
	
	if lambdaBaseApi != "" {
		conf.lambdaBaseApi = lambdaBaseApi
	}
	
	if lambdaHandlerApi != "" {
		conf.lambdaHandlerApi = lambdaHandlerApi
	}
	
	if lambdaNameApi != "" {
		conf.lambdaNameApi = lambdaNameApi
	}
	
	if lambdaPermApi != "" {
		conf.lambdaPermApi = lambdaPermApi
	}
	
	if lambdaPolicyArn != "" {
		conf.lambdaPolicyArn = lambdaPolicyArn
	}
	
	if lambdaRoleArnApi != "" {
		conf.lambdaRoleArnApi = lambdaRoleArnApi
	}
	
	if portApi != "" {
		conf.portApi = portApi
	}
	
	if portConsole != "" {
		conf.portConsole = portConsole
	}
	
	if region != "" {
		conf.region = region
	}
	
	if versionApi != "" {
		conf.versionApi = versionApi
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
	
	v = os.Getenv("APP_GW_BASE_API")
	if v != "" {
		conf.gwBaseApi = v
	}
	
	v = os.Getenv("APP_GW_DOMAIN")
	if v != "" {
		conf.gwDomain = v
	}
	
	v = os.Getenv("APP_GW_ID_API")
	if v != "" {
		conf.gwIdApi = v
	}
	
	v = os.Getenv("APP_GW_PATH_API")
	if v != "" {
		conf.gwPathApi = v
	}
	
	v = os.Getenv("APP_GW_PROXY_API")
	if v != "" {
		conf.gwProxyApi = v
	}
	
	v = os.Getenv("APP_GW_ROOT_API")
	if v != "" {
		conf.gwRootApi = v
	}
	
	v = os.Getenv("APP_GW_STAGE_NAME_API")
	if v != "" {
		conf.gwStageNameApi = v
	}
	
	v = os.Getenv("APP_GW_SUBDOMAIN")
	if v != "" {
		conf.gwSubdomain = v
	}
	
	v = os.Getenv("APP_LAMBDA_ARN_API")
	if v != "" {
		conf.lambdaArnApi = v
	}
	
	v = os.Getenv("APP_LAMBDA_BASE_API")
	if v != "" {
		conf.lambdaBaseApi = v
	}
	
	v = os.Getenv("APP_LAMBDA_HANDLER_API")
	if v != "" {
		conf.lambdaHandlerApi = v
	}
	
	v = os.Getenv("APP_LAMBDA_NAME_API")
	if v != "" {
		conf.lambdaNameApi = v
	}
	
	v = os.Getenv("APP_LAMBDA_PERM_API")
	if v != "" {
		conf.lambdaPermApi = v
	}
	
	v = os.Getenv("APP_LAMBDA_POLICY_ARN")
	if v != "" {
		conf.lambdaPolicyArn = v
	}
	
	v = os.Getenv("APP_LAMBDA_ROLE_ARN_API")
	if v != "" {
		conf.lambdaRoleArnApi = v
	}
	
	v = os.Getenv("APP_PORT_API")
	if v != "" {
		conf.portApi = v
	}
	
	v = os.Getenv("APP_PORT_CONSOLE")
	if v != "" {
		conf.portConsole = v
	}
	
	v = os.Getenv("APP_REGION")
	if v != "" {
		conf.region = v
	}
	
	v = os.Getenv("APP_VERSION_API")
	if v != "" {
		conf.versionApi = v
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
