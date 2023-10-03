package config

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// Config struct for reading data from json
type Config struct {
	Keycloak                    KeycloakParams     `mapstructure:"keycloak"`
	Database                    Database           `mapstructure:"db"`
	ProcAddr                    string             `mapstructure:"proc_addr"`
	TokenAuth                   string             `mapstructure:"token_auth"`
	AppParams                   Params             `mapstructure:"app"`
	TcellParams                 TcellParams        `mapstructure:"tcell"`
	FraudParams                 TcellParams        `mapstructure:"fraud"`
	BlackListParams             string             `mapstructure:"blacklist_params"`
	DisputeParams               DisputeParams      `mapstructure:"dispute"`
	BalanceParams               BalanceParams      `mapstructure:"balances"`
	PrecheckParams              PrecheckParams     `mapstructure:"pre_check"`
	TermParams                  TermParams         `mapstructure:"terminal"`
	OnlineCancel                OnlineCancelParams `mapstructure:"online_cancel"`
	ComlianceToken              string             `mapstructure:"compliance_token"`
	ComplianceAddr              string             `mapstructure:"compliance_addr"`
	ReestrAddr                  string             `mapstructure:"reestr_addr"`
	ReestrToken                 string             `mapstructure:"reestr_token"`
	JobCreateCommandForTerminal bool               `mapstructure:"create_command"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SslMode  string `mapstructure:"ssl_mode"`
}

type KeycloakParams struct {
	Realm               string `mapstructure:"realm"`
	BaseUrl             string `mapstructure:"base_url"`
	ClientID            string `mapstructure:"client_id"`
	ClientSecret        string `mapstructure:"client_secret"`
	RealmRS256PublicKey string `mapstructure:"realm_rs256_public_key"`
}

type TcellParams struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}

// Params contains params of server meta data
type Params struct {
	ServerName   string `mapstructure:"server_name"`
	PortRun      string `mapstructure:"port_run"`
	TokenTimeout int    `mapstructure:"token_timeout"`
	LogFile      string `mapstructure:"log_file"`
	Key          string `mapstructure:"key"`
}

type DisputeParams struct {
	APIKey string `mapstructure:"api_key"`
	ID     int    `mapstructure:"id"`
}

type BalanceParams struct {
	BabilonMBalance string `mapstructure:"babilon_m_balance"`
	MegafonBalance  string `mapstructure:"megafon_balance"`
	BeelineBalance  string `mapstructure:"beeline_balance"`
	TcellBalance    string `mapstructure:"tcell_balance"`
	FormulaBalance  string `mapstructure:"formula_balance"`
}

type PrecheckParams struct {
	Somontjpsk string `mapstructure:"somontj_psk"`
	Somontjpsi string `mapstructure:"somontj_psi"`
	SomontjUrl string `mapstructure:"somontj_url"`
	HumoCreUrl string `mapstructure:"humo_cre_url"`
}

type TermParams struct {
	Pin       string `mapstructure:"pin"`
	SecretAES string `mapstructure:"secret_aes"`
}

type OnlineCancelParams struct {
	Megafon string `mapstructure:"megafon"`
	Beeline string `mapstructure:"beeline"`
	Babilon string `mapstructure:"babilon"`
}

type GTUserBalance struct {
	Balance    float64 `gorm:"column:Balance"`
	TrustLevel float64 `gorm:"column:TrustLevel"`
}
type GatewayUserBalance struct {
	XMLName   xml.Name `xml:"response"`
	Balance   float64  `xml:"balance"`
	Overdraft float64  `xml:"overdraft"`
}

type TermChPw struct {
	XMLName xml.Name `xml:"response"`
	PWKey   string   `xml:"pwkey"`
}

type SampleXML1Value struct {
	XMLName xml.Name `xml:"response"`
	Value   string   `xml:"value"`
}

func ReadConfig() *Config {
	var appParams Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetEnvPrefix("falcon")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}
	log.Println("viper config initialized")

	err = viper.Unmarshal(&appParams)
	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}

	return &appParams
}
