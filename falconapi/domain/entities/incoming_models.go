package entities

import (
	"encoding/xml"
	"time"
)

type IncomingUserRole struct {
	userName string // userName
	role     string // roleName
	userID   int    // id
}

type PasswordChange struct {
	OldPassword string `json:"oldpassword" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}

type ReceivedPayment struct {
	DateFrom       string   `json:"dateFrom" binding:"required"`
	DateTo         string   `json:"dateTo" binding:"required"`
	TimeFrom       *string  `json:"TimeFrom"`
	TimeTo         *string  `json:"TimeTo"`
	PhoneNumber    *string  `json:"phoneNumber"`
	AmountOriginal *float64 `json:"AmountOriginal"`
	TerminalID     *int     `json:"TerminalId"`
	AgentID        *int     `json:"agentId"`
	TransactionID  *int     `json:"transactionId"`
	Provider       *int     `json:"providerId"`
	Status         *int     `json:"statusId"`
	Page           *int     `json:"page"`
	Gateway        *int     `json:"gateway"`
	ItemOnPage     *int     `json:"itemOnPage"`
}

type ReceivedChecks struct {
	DateFrom   string  `json:"dateFrom" binding:"required"`
	DateTo     string  `json:"dateTo" binding:"required"`
	TimeFrom   *string `json:"TimeFrom"`
	PaymentID  *int    `json:"paymentID"`
	Account    *string `json:"account"`
	TimeTo     *string `json:"TimeTo"`
	Page       *int    `json:"page"`
	ItemOnPage *int    `json:"itemOnPage"`
}

type TcellParam struct {
	DateFrom string  `json:"dateFrom" binding:"required"`
	DateTo   string  `json:"dateTo" binding:"required"`
	TimeFrom *string `json:"TimeFrom"`
	TimeTo   *string `json:"TimeTo"`
}

type UserCreate struct {
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	FullName    string  `json:"fullName"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	IsAgent     bool    `json:"isAgent"`
	ParentAgent *int64  `json:"parentAgent"`
	Address     string  `json:"address"`
	Overdraft   float64 `json:"overdraft"`
	Remain      float64 `json:"remain"`
}

type TerminalCreate struct {
	Type              int64   `json:"terminalType"`
	AgentID           int64   `json:"agent"`
	Address           string  `json:"address"`
	RegionID          int64   `json:"region_id"`
	Description       string  `json:"description"`
	Threshold         int64   `json:"threshold"`
	Disabled          bool    `json:"activation"`
	Phone             string  `json:"phone"`
	Apikey            *string `json:"apiKey"`
	AuthParam1        string  `json:"authParam1"`
	Num               int64   `json:"endpointNum"`
	SecretKey         *string `json:"secretKey"`
	Password          *string `json:"password"`
	CustomServiceList bool    `json:"customServiceList"`
}

type XlsPayment struct {
	DateFrom      string  `json:"dateFrom" binding:"required"`
	DateTo        string  `json:"dateTo" binding:"required"`
	TimeFrom      *string `json:"TimeFrom"`
	TimeTo        *string `json:"TimeTo"`
	PhoneNumber   *string `json:"phoneNumber"`
	TerminalID    *int    `json:"TerminalId"`
	AgentID       *int    `json:"agentId"`
	TransactionID *int    `json:"transactionId"`
	Provider      *int    `json:"providerId"`
	Status        *int    `json:"statusId"`
}

type NewsJSON struct {
	NewsID  int32     `json:"NewsId"`
	Subject string    `json:"Subject"`
	Date    time.Time `json:"Date"`
	Title   string    `json:"Title"`
	Body    string    `json:"Body"`
}

type CollectionJSON struct {
	DateFrom   string `json:"dateFrom"`
	DateTo     string `json:"dateTo"`
	TerminalID int64  `json:"terminalId"`
	RegionID   *int64 `json:"region_id"`
	Date       string `json:"date"`
}

type AgentJSON struct {
	Agent          string  `json:"agent"`
	AgentAddr      string  `json:"agentAddr"`
	AgentBalance   float64 `json:"agentBalance"`
	AgentDesc      string  `json:"agentDesc"`
	AgentINN       string  `json:"agentINN"`
	AgentOverdraft float64 `json:"agentOverdraft"`
	AgentParent    int64   `json:"agentParent"`
	Currency       int64   `json:"currency"`
}

type UserJSON struct {
	Fname    string `json:"fname"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	UserDesc string `json:"userDesc"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserpassJSON struct {
	UserID  int    `json:"UserID"`
	NewPass string `json:"newPass"`
	RePass  string `json:"rePass"`
}

type IncomRequest struct {
	XMLName xml.Name `xml:"request" json:"-"`
	Command string   `xml:"cmd" json:"-"`
	UserID  int      `xml:"userid" json:"-"`
	AgentID int      `xml:"id" json:"agent"`
	Sum     float64  `xml:"amount" json:"sum"`
}

type ProcResp struct {
	XMLName xml.Name `xml:"response" json:"-"`
	Code    int      `xml:"code" json:"code"`
	Desc    string   `xml:"desc" json:"desc"`
	Comment string   `xml:"comment" json:"comment"`
}

type CancelPaymentReq struct {
	XMLName   xml.Name `xml:"request" json:"-"`
	Command   string   `xml:"cmd" json:"-"`
	PaymentID int      `xml:"id" json:"paymentID"`
	UserID    int      `xml:"userid" json:"-"`
	Status    int      `json:"status"`
}

type SendReceipt struct {
	XMLName   xml.Name `xml:"request" json:"-"`
	Command   string   `xml:"cmd" json:"-"`
	PaymentID int      `xml:"id" json:"tranid"`
	UserID    int      `xml:"userid" json:"-"`
	Phone     string   `xml:"phone" json:"phone"`
}

type ChangePayment struct {
	XMLName   xml.Name `xml:"request" json:"-"`
	Command   string   `xml:"cmd" json:"-"`
	PaymentID int      `xml:"id" json:"tranid"`
	UserID    int      `xml:"userid" json:"-"`
	Status    int      `xml:"status" json:"status"`
}
type RecastPaymentReq struct {
	XMLName   xml.Name `xml:"request" json:"-"`
	Command   string   `xml:"cmd" json:"-"`
	PaymentID int      `xml:"id" json:"paymentID"`
	UserID    int      `xml:"userid" json:"-"`
	Gateway   *string  `xml:"gateway" json:"gateway"`
	ServiceID *int     `xml:"serviceid" json:"serviceID"`
	Account   *string  `xml:"account" json:"account"`
}

type EditAgentJSON struct {
	ID        int64  `json:"agentID"`
	Desc      string `json:"agentDesc"`
	Address   string `json:"agentAddress"`
	ParentID  int64  `json:"parentID"`
	Overdraft string `json:"overdraft"`
	ITN       string `json:"agentItn"`
}

type SearchLog struct {
	ParamID int64 `json:"paramId"`
}

type SearchRecast struct {
	DateFrom        string
	DateTo          string
	DisputPaymentID *int64
}

type CommandSearch struct {
	DateFrom string `binding:"required"`
	DateTo   string `binding:"required"`
}

type ActLogRequest struct {
	DateFrom string `binding:"required"`
	DateTo   string `binding:"required"`
	Agent    *int
}

type CorrespondJSON struct {
	Value     []string `json:"value"`
	DateFrom  string   `json:"dateFrom"`
	DateTo    string   `json:"dateTo"`
	Status    *int     `json:"status"`
	Provider  *int     `json:"provider"`
	Payment   string   `json:"payment"`
	IsReverse bool     `json:"isReverse"`
	Gateway   *string
}

type HandyCorrResp struct {
	Left  []string `json:"left"`
	Right []string `json:"right"`
}

type ReestrListJSON struct {
	DateFrom   string  `json:"dateFrom" binding:"required"`
	DateTo     string  `json:"dateTo" binding:"required"`
	Status     *int    `json:"status"`
	Detail     *string `json:"detail"`
	AgentID    *int    `json:"agent_id"`
	VendorID   *int    `json:"vendor_id"`
	Page       int     `json:"page"`
	ItemOnPage int     `json:"itemOnPage"`
}

type DeviceJSON struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

type Endpoint2DeviceJSON struct {
	EndpointID int64 `json:"endpointID"`
	DeviceID   int64 `json:"deviceID"`
	Disabled   bool  `json:"disabled"`
}

type RegistList struct {
	ID             int64     `json:"id"`
	AgentID        int64     `json:"agent_id"`
	VendorID       int64     `json:"vendor_id"`
	PartnerID      int64     `json:"partner_id"`
	RegDate        time.Time `json:"reg_date"`
	DateStart      time.Time `json:"date_start"`
	DateStop       time.Time `json:"date_stop"`
	Status         int64     `json:"status"`
	HumoErr        int64     `json:"humo_errors"`
	PartnerErr     int64     `json:"partner_errors"`
	Detail         string    `json:"detail"`
	FileName       string    `json:"filename"`
	InputFilename  string    `json:"input_filename"`
	OutputFilename string    `json:"output_filename"`
	DetailStatus   string    `json:"detail_status"`
	Comment        string    `json:"comment"`
}

type Partner struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	IntervalStart  int64     `json:"interval_start"`
	IntervalStop   int64     `json:"interval_stop"`
	InputFilename  string    `json:"input_filename"`
	OutputFilename string    `json:"output_filename"`
	SaveFilename   string    `json:"save_filename"`
	DiffTime       int64     `json:"difference_in_time"`
	IntervalType   string    `json:"interval_type"`
	DateStart      time.Time `json:"date_start"`
	WhichTime      string    `json:"which_time"`
	AgentID        int64     `json:"agent_id"`
	AgentType      string    `json:"agent_type"`
	Active         bool      `json:"active"`
	Statuses       string    `json:"statuses"`
	VendorID       int64     `json:"vendor_id"`
	Gateways       []int     `json:"gateways"`
}

/* sub category mcc */
type SubCategoryOfSerivceMCCJson struct {
	Name        string `json:"name"`
	MccID       int    `json:"mccID"`
	Description string `json:"description"`
	Code        int    `json:"code"`
}
