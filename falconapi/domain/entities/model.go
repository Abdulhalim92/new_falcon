package entities

import (
	"encoding/xml"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Product struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Name      string
	Price     float32
}

type ErrorModel struct {
	Err     error
	Message string
	Code    int
}

type SumAmount struct {
	Sum      float32 `gorm:"column:sum"`
	Count    int     `gorm:"column:count"`
	Currency string  `gorm:"column:curr"`
}

type paymentStatus struct {
	ID     int    `json:"ID"`
	Status string `json:"Status"`
}

// PingHumoXML struct for testing ping connection.
type PingHumoXML struct {
	XMLName xml.Name `xml:"response"`
	CodeTag string   `xml:"status"`
}

type HumoExtPing struct {
	XMLName   xml.Name `xml:"response"`
	CodeTag   string   `xml:"status"`
	Timestamp string   `xml:"timestamp"`
}

// User struct. Contain user info
type User struct {
	ID        int64      `gorm:"column:user_id;primary_key"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	Email     string     `gorm:"column:email"`
	Desc      string     `gorm:"column:user_desc"`
	FullName  string     `gorm:"column:fullname"`
	Phone     string     `gorm:"column:phone"`
	Salt      string     `gorm:"column:salt"`
	Role      string     `gorm:"column:role"`
	Disabled  bool       `gorm:"column:disabled"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	LoginAt   *time.Time `gorm:"column:login_at"`
}
type TUser struct {
	ID        int64      `gorm:"column:user_id;primary_key"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	Email     string     `gorm:"column:email"`
	FullName  string     `gorm:"column:fullname"`
	Desc      string     `gorm:"column:user_desc"`
	Phone     string     `gorm:"column:phone"`
	Salt      string     `gorm:"column:salt"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	LoginAt   *time.Time `gorm:"column:login_at"`
}

// TableName corresponding for User struct
func (User) TableName() string {
	return "tusers"
}

// TableName corresponding for User struct
func (TUser) TableName() string {
	return "tusers"
}

// Agent struct. Contain agent info.
type Agent struct {
	ID           int64   `gorm:"column:agent_id;primary_key"`
	ParentID     int64   `gorm:"column:parent_id"`
	Remain       float64 `gorm:"column:remain"`
	AgentDesc    string  `gorm:"column:agent_desc"`
	AgentName    string  `gorm:"column:agent_name"`
	Overdraft    float64 `gorm:"column:overdraft"`
	AccountFi    string  `gorm:"column:account_fi"`
	AgentAddress string  `gorm:"column:agent_address"`
	AgentITN     string  `gorm:"column:agent_itn"`
	Currency     string  `gorm:"column:currency"`
}

// TableName corresponding for Agent struct
func (Agent) TableName() string {
	return "tagents"
}

type NotifyRule struct {
	ID           int64      `gorm:"column:id;primary_key"`
	AgentID      int64      `gorm:"column:agent_id"`
	EndpointID   int64      `gorm:"column:endpoint_id"`
	EndpointType int64      `gorm:"column:endpoint_type"`
	ServiceID    int64      `gorm:"column:service_id"`
	GatewayUniq  string     `gorm:"column:gateway_unique"`
	VendorID     int64      `gorm:"column:vendor_id"`
	NotifyStatus string     `gorm:"column:notify_status"`
	Template     string     `gorm:"column:template"`
	Disabled     bool       `gorm:"column:disabled"`
	CreatedAT    *time.Time `gorm:"column:created_at"`
	CreatedBy    int        `gorm:"column:created_by" json:"-"`
	ModifiedBy   int        `gorm:"column:modified_by" json:"-"`
	ModifiedAT   *time.Time `gorm:"column:modified_at" json:"-"`
}

// TableName corresponding for Agent struct
func (NotifyRule) TableName() string {
	return "tnotifyrules"
}

// Currency struct. Contain agent info.
type Currency struct {
	ID        int64   `gorm:"column:id;primary_key"`
	CreatedAt string  `gorm:"column:created_at"`
	DeletedAt float64 `gorm:"column:deleted_at"`
	UpdatedAt string  `gorm:"column:updated_at"`
	Currency  string  `gorm:"column:currency"`
	Buy       float64 `gorm:"column:buy"`
	Sell      float64 `gorm:"column:sell"`
}

// TableName corresponding for Currency struct
func (Currency) TableName() string {
	return "tchanges"
}

// Exchange struct
type Exchange struct {
	Base      string  `gorm:"column:base;primary_key"`
	Valuation string  `gorm:"column:valuation;primary_key"`
	CreatedAt string  `gorm:"column:created_at"`
	UpdatedAt string  `gorm:"column:updated_at"`
	Rate      float64 `gorm:"column:rate"`
}

// TableName corresponding for Exchange struct
func (Exchange) TableName() string {
	return "texchange"
}

// UserRole struct. There are supervisor, admin, operator and etc.
type UserRole struct {
	ID       int64  `gorm:"column:role_id;primary_key"`
	RoleName string `gorm:"column:role_name"`
	RoleDesc string `gorm:"column:role_desc"`
}

// TableName corresponding for UserRole struct
func (UserRole) TableName() string {
	return "troles"
}

// User2Role struct
type User2Role struct {
	ID     int64 `gorm:"column:id;primary_key"`
	UserID int64 `gorm:"column:user_id"`
	RoleID int64 `gorm:"column:role_id"`
}

// TableName corresponding for User2Role struct
func (User2Role) TableName() string {
	return "tuser2role"
}

// Payment struct
type Payment struct {
	ID               int64      `gorm:"column:payment_id;primary_key"`
	ReferenceID      string     `gorm:"column:reference_id"`
	ReceiptID        string     `gorm:"column:receipt_id"`
	TrnxID           string     `gorm:"column:trnx_id"`
	ServiceID        int64      `gorm:"column:service_id"`
	Account          string     `gorm:"column:account"`
	Amount           float64    `gorm:"column:amount"`
	Amount2Credit    float64    `gorm:"column:amount2credit"`
	DetailedStatus   string     `gorm:"column:detailed_status"`
	RemainBefore     float64    `gorm:"column:remain_before"`
	RemainAfter      float64    `gorm:"column:remain_after"`
	ExtStatus        string     `gorm:"column:ext_status"`
	ExtParam         string     `gorm:"column:ext_param"`
	CreatedAt        *time.Time `gorm:"column:created_at"`
	UpdateAt         *time.Time `gorm:"column:updated_at"`
	ProccessedAt     *time.Time `gorm:"column:proccessed_at;timestamp"`
	SubmittedAt      *time.Time `gorm:"column:submitted_at;timestamp"`
	Status           int        `gorm:"column:status"`
	AgentID          int        `gorm:"column:agent_id"`
	AgentTermID      int        `gorm:"column:agent_term_id"`
	NotifyRoute      string     `gorm:"column:notify_route"`
	NotifyAt         *time.Time `gorm:"column:notify_at"`
	Retries          int        `gorm:"column:retries"`
	AgentTermAddr    string     `gorm:"column:agent_term_addr"`
	PaymentStatus    string     `gorm:"-"`
	GatewayUnique    string     `gorm:"column:gateway_unique"`
	AgentTermReceipt string     `gorm:"column:agent_term_receipt"`
	BillInfo         string     `gorm:"column:bill_info"`
	RevokeID         *string    `gorm:"column:revoke_id"`
	Currency         string     `gorm:"column:currency"`
	Currency2GW      string     `gorm:"column:currency2gw"`
	Amount2Credit2GW float64    `gorm:"column:amount2credit2gw"`
	Rate             string     `gorm:"column:rate"`
	AcceptedAt       *time.Time `gorm:"column:accepted_at"`
	CheckForReestr   bool       `gorm:"-"`
}

type PaymentXls struct {
	ID               int64      `gorm:"column:payment_id;primary_key"`
	ReferenceID      string     `gorm:"column:reference_id"`
	ReceiptID        string     `gorm:"column:receipt_id"`
	TrnxID           string     `gorm:"column:trnx_id"`
	ServiceID        int64      `gorm:"column:service_id"`
	Account          string     `gorm:"column:account"`
	Amount           float64    `gorm:"column:amount"`
	Amount2Credit    float64    `gorm:"column:amount2credit"`
	DetailedStatus   string     `gorm:"column:detailed_status"`
	RemainBefore     float64    `gorm:"column:remain_before"`
	RemainAfter      float64    `gorm:"column:remain_after"`
	ExtStatus        string     `gorm:"column:ext_status"`
	ExtParam         string     `gorm:"column:ext_param"`
	CreatedAt        *time.Time `gorm:"column:created_at"`
	UpdateAt         *time.Time `gorm:"column:updated_at"`
	ProccessedAt     *time.Time `gorm:"column:proccessed_at;timestamp"`
	SubmittedAt      *time.Time `gorm:"column:submitted_at;timestamp"`
	Status           int        `gorm:"column:status"`
	AgentID          int        `gorm:"column:agent_id"`
	AgentName        string     `gorm:"column:agent_name"`
	AgentTermID      int        `gorm:"column:agent_term_id"`
	NotifyRoute      string     `gorm:"column:notify_route"`
	NotifyAt         *time.Time `gorm:"column:notify_at"`
	Retries          int        `gorm:"column:retries"`
	AgentTermAddr    string     `gorm:"column:agent_term_addr"`
	PaymentStatus    string     `gorm:"-"`
	GatewayUnique    string     `gorm:"column:gateway_unique"`
	AgentTermReceipt string     `gorm:"column:agent_term_receipt"`
	BillInfo         string     `gorm:"column:bill_info"`
	RevokeID         *string    `gorm:"column:revoke_id"`
	Currency         string     `gorm:"column:currency"`
	Currency2GW      string     `gorm:"column:currency2gw"`
	Amount2Credit2GW float64    `gorm:"column:amount2credit2gw"`
	Rate             string     `gorm:"column:rate"`
	AcceptedAt       *time.Time `gorm:"column:accepted_at"`
	CheckForReestr   bool       `gorm:"-"`
	CategoryID       int        `gorm:"category_id"`
	ServiceDesc      string     `gorm:"service_desc"`
	CategoryDesc     string     `gorm:"column:category_desc"`
	VendorName       string     `gorm:"column:vendor_name"`
}

// TableName corresponding for Payment struct
func (Payment) TableName() string {
	return "tpayments"
}

type ChecksStruct struct {
	PaymentID int64     `gorm:"column:payment_id;primary_key"`
	Status    string    `gorm:"column:status"`
	Notify    string    `gorm:"column:notify_text"`
	CreatedAT time.Time `gorm:"column:created_at"`
	Account   string    `gorm:"account"`
}

func (ChecksStruct) TableName() string {
	return "treceipt"
}

// Service struct
type Service struct {
	ID              int64    `gorm:"column:service_id"`
	ServiceDesc     string   `gorm:"column:service_desc"`
	CategoryID      int64    `gorm:"column:category_id"`
	Category        Category `gorm:"foreignkey:category_id"`
	Disabled        bool     `gorm:"column:disabled"`
	TermDisabled    bool     `gorm:"column:term_disabled"`
	XMLGateDisabled bool     `gorm:"column:xmlgate_disabled"`
	TermIndex       int64    `gorm:"column:term_index"`
	GatewayUnique   string   `gorm:"column:gateway_unique"`
	Currency        string   `gorm:"column:currency"`
	ModifiedAt      *time.Time
	ModifiedBy      *int
	SubCategoryID   int `gorm:"sub_category_id"`
	SubMccID        int `gorm:"column:sub_mcc_id"`
}

// TableName corresponding for Payment struct
func (Service) TableName() string {
	return "tservices"
}

type Category struct {
	CategoryID   int64  `gorm:"column:category_id;primary_key"`
	CategoryDesc string `gorm:"column:category_desc"`
}

func (Category) TableName() string {
	return "tservicecategories"
}

type TRegion struct {
	ID   int64  `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (TRegion) TableName() string {
	return "tregion"
}

// TEndpoint struct
type TEndpoint struct {
	ID                int64      `gorm:"column:endpoint_id;primary_key;"`
	Type              int64      `gorm:"column:type"`
	Agent             int64      `gorm:"column:agent_id"`
	RegionID          int64      `gorm:"column:region_id"`
	Address           string     `gorm:"column:address"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at"`
	CreatedBy         int64      `gorm:"column:created_by"`
	UpdatedBy         *int64     `gorm:"column:updated_by"`
	APIKey            string     `gorm:"column:apikey"`
	APIKeyExp         *time.Time `gorm:"column:apikey_expire"`
	SecretKey         string     `gorm:"column:secretkey"`
	EndpointDesc      string     `gorm:"column:endpoint_desc"`
	AuthParam1        string     `gorm:"column:auth_param1"`
	AuthParam2        string     `gorm:"column:auth_param2"`
	EndpointThreshold int64      `gorm:"column:endpoint_threshold"`
	EndpointName      string     `gorm:"column:endpoint_name"`
	Phone             string     `gorm:"column:phone"`
	EndpointDisabled  bool       `gorm:"column:endpoint_disabled"`
	EndpointNum       int64      `gorm:"column:endpoint_num"`
	CustomServiceList bool       `gorm:"column:custom_servicelist"`
}

// TableName corresponding for TEndpoint struct
func (TEndpoint) TableName() string {
	return "tendpoints"
}

// TOnlineCancel struct
type TOnlineCancel struct {
	CancelID          int32      `gorm:"column:cancel_id;primary_key"`
	UserId            int        `gorm:"column:user_id"`
	PaymentId         int        `gorm:"column:payment_id"`
	PaymentProviderId int        `gorm:"column:service_id"`
	AmountOriginal    float32    `gorm:"column:amount"`
	PhoneNumber       string     `gorm:"column:account"`
	CancelStatus      *int       `gorm:"column:cancel_status"`
	GTStatus          int        `gorm:"-"`
	ReasonDesc        string     `gorm:"column:reason_desc"`
	DateSubmitted     *time.Time `gorm:"column:date_submitted"`
}

// TableName corresponding for TOnlineCancel struct
func (TOnlineCancel) TableName() string {
	return "tcancelonlines"
}

// PaymentCancel struct
type PaymentCancel struct {
	PaymentID string `json:"paymentId"`
}

type GTUserBalance struct {
	Balance    float64 `gorm:"column:Balance"`
	TrustLevel float64 `gorm:"column:TrustLevel"`
}

type GTWUserBalance struct {
	XMLName   xml.Name `xml:"response"`
	Balance   float64  `xml:"balance"`
	Overdraft float64  `xml:"overdraft"`
}

// News struct
type News struct {
	ID         int64      `gorm:"column:news_id;primary_key"`
	Title      string     `gorm:"column:title"`
	Shortcut   string     `gorm:"column:shortcut"`
	Text       string     `gorm:"column:text"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
	CreatedBy  int64      `gorm:"column:created_by"`
	ModifiedBy int64      `gorm:"column:modified_by"`
}

// TableName corresponding for News struct
func (News) TableName() string {
	return "tnews"
}

// TEncashment struct
type TEncashment struct {
	ID          int64     `gorm:"column:collection_id;primary_key" ru:"ID"`
	Date        time.Time `gorm:"column:collection_date" ru:"Дата инкассации"`
	PrevDate    time.Time `gorm:"column:prev_date"`
	EndpointID  int64     `gorm:"column:endpoint_id" ru:"ID терминала"`
	Address     string    `gorm:"column:address" ru:"Адрес терминала"`
	MoneyPairs  string    `gorm:"column:pairs_arr" ru:"Купюры"`
	TotalAmount float64   `gorm:"column:total_amount" ru:"Сумма"`
	Desc        string    `gorm:"column:collection_desc" ru:"Описание"`
	CollectorID int64     `gorm:"column:collector_id" ru:"ID коллектора"`
	EndpointNum string    `gorm:"column:endpoint_num" ru:"Номер терминала"`
	RegionID    int64     `gorm:"column:region_id" ru:"Регион"`
}

type TEncashmentXLS struct {
	ID          int64     `ru:"ID"`
	Date        time.Time `ru:"Дата инкассации"`
	EndpointID  int64     `ru:"ID терминала"`
	Address     string    `ru:"Адрес терминала"`
	MoneyPairs  string    `ru:"Купюры"`
	TotalAmount float64   `ru:"Сумма"`
	Desc        string    `ru:"Описание"`
	CollectorID int64     `ru:"ID коллектора"`
	EndpointNum string    `ru:"Номер терминала"`
	Region      string    `ru:"Регион"`
}

// // TableName correspond for TEncashment
// func (TEncashment) TableName() string {
// 	return "tcollections"
// }

// AddServiceStruct
type AddServiceStruct struct {
	ID            int64  `gorm:"column:service_id;primary_key"`
	ServiceDesc   string `gorm:"column:service_desc" json:"provider"`
	Disabled      bool   `gorm:"column:disabled"`
	TermDisabled  bool   `gorm:"column:term_disabled" json:"termDisabled"`
	XMLDisabled   bool   `gorm:"column:xmlgate_disabled" json:"xmlDisabled"`
	TermIndex     int64  `gorm:"column:term_index" json:"termIndex"`
	GatewayUnique string `gorm:"column:gateway_unique" json:"gwUnique"`
	MatchingID    string `gorm:"-" json:"matchingID"`
	CreatedAt     *time.Time
	ModifiedAt    *time.Time
	CreatedBy     *int
	ModifiedBy    *int
	SubCategoryID int `gorm:"column:sub_category_id" json:"subCategory"`
	SubMccID      int `gorm:"column:sub_mcc_id" json:"subCategoryMCC"`
}

func (AddServiceStruct) TableName() string {
	return "tservices"
}

type MatchingService struct {
	ServiceID     int64  `gorm:"column:service_id"`
	MatchingID    string `gorm:"column:matching_id"`
	GatewayUnique string `gorm:"column:gateway_unique"`
}

func (MatchingService) TableName() string {
	return "tservice2gwmatch"
}

type Vendor struct {
	VendorID   int64      `gorm:"column:vendor_id;primary_key" json:"id"`
	VendorName string     `gorm:"column:vendor_name" json:"vendorName"`
	VendorDesc string     `gorm:"column:vendor_desc" json:"vendorDesc"`
	VendorFi   string     `gorm:"column:vendor_fi" json:"vendorFi"`
	Remain     float64    `gorm:"column:remain" json:"balance"`
	Overdraft  *float64   `gorm:"column:overdraft" json:"overdraft"`
	Currency   string     `gorm:"column:currency" json:"currency"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"-"`
	CreatedBy  int64      `gorm:"column:created_by" json:"-"`
	ModifiedBy *int64     `gorm:"column:modified_by" json:"-"`
}
type Vendor1 struct {
	VendorID   int64      `gorm:"column:vendor_id;primary_key" json:"id"`
	VendorName string     `gorm:"column:vendor_name" json:"vendorName"`
	VendorDesc string     `gorm:"column:vendor_desc" json:"vendorDesc"`
	VendorFi   string     `gorm:"column:vendor_fi" json:"vendorFi"`
	Currency   string     `gorm:"column:currency" json:"currency"`
	Remain     float64    `gorm:"column:remain" json:"-"`
	Overdraft  *float64   `gorm:"column:overdraft" json:"overdraft"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"-"`
	CreatedBy  int64      `gorm:"column:created_by" json:"-"`
	ModifiedBy *int64     `gorm:"column:modified_by" json:"-"`
}

func (Vendor) TableName() string {
	return "tvendors"
}

func (Vendor1) TableName() string {
	return "tvendors"
}

type GwMatching struct {
	ID       int64  `gorm:"column:service_id"`
	Desc     string `gorm:"column:service_desc"`
	Matching string `gorm:"column:matching_id"`
	GwUnique string `gorm:"column:gateway_unique"`
}

type UpdateGW struct {
	Old GwMatching
	New GwMatching
}
type GWMTCH struct {
	ID       int64  `gorm:"column:service_id"`
	Matching string `gorm:"column:matching_id"`
	GwUnique string `gorm:"column:gateway_unique"`
}

func (GwMatching) TableName() string {
	return "tservice2gwmatch"
}
func (GWMTCH) TableName() string {
	return "tservice2gwmatch"
}

type Gateway struct {
	ID     int64  `gorm:"column:gateway_id;primary_key"`
	Desc   string `gorm:"column:gateway_desc"`
	Unique string `gorm:"column:gateway_unique"`
	Param1 string `gorm:"column:gateway_param1"`
	Param2 string `gorm:"column:gateway_param2"`
	Param3 string `gorm:"column:gateway_param3"`
	Vendor int64  `gorm:"column:gateway_vendor_id"`
}

func (Gateway) TableName() string {
	return "tgateways"
}

type paymentInterchangeIncoming struct {
	DateFrom   string  `gorm:"column:dateFrom"`
	DateTo     string  `gorm:"column:dateTo"`
	Status     *string `gorm:"column:status"`
	Agent      *int64  `gorm:"column:agent_id"`
	TerminalID *int    `gorm:"agent_term_id" json:"terminalId"`
	Gateway    string
	Page       *int `json:"page"`
}

type PaymentInterchange struct {
	ID           int64      `gorm:"column:payment_id;primary_key"`
	Status       string     `gorm:"column:interchange_status"`
	CreatedAt    *time.Time `gorm:"column:created_at;timestamp"`
	ProccessedAt *time.Time `gorm:"column:proccessed_at;timestamp"`
}

func (PaymentInterchange) TableName() string {
	return "tpayment2interchange"
}

type PaymentInterchange2Cft struct {
	ID                 int64         `gorm:"column:tran_id;primary_key"`
	Status             string        `gorm:"column:tran_status" ru:"Статус"`
	CreatedAt          *time.Time    `gorm:"column:created_at;timestamp" ru:"Дата создания"`
	ProccessedAt       *time.Time    `gorm:"column:proccessed_at;timestamp" `
	TranSum            float64       `gorm:"column:tran_sum" ru:"Сумма"`
	GWUnique           string        `gorm:"column:gateway_unique" ru:"Шлюз"`
	ServiceID          int64         `gorm:"column:service_id" ru:"Сервис"`
	VendorID           *int64        `gorm:"column:vendor_id" ru:"Вендор"`
	TranCount          int64         `gorm:"column:tran_count" ru:"Кол-во транзакций"`
	PaymentList        pq.Int64Array `gorm:"not null;type:BigInt[];column:paymentid_list" ru:""`
	TransSumGW         float64       `gorm:"column:tran_sum_2gw" ru:"Сумма отправки"`
	Currency           string        `gorm:"column:currency" ru:"Валюта"`
	CurrencyGw         string        `gorm:"column:currency2gw" ru:"Валюта"`
	Rate               float64       `gorm:"column:rate" ru:"Курс"`
	Agent              int64         `gorm:"column:agent_id" ru:"Агент"`
	TranSumWithComm    float64       `gorm:"column:tran_sum_with_comm`
	TranSumWithComm2Gw float64       `gorm:"column:tran_sum_with_comm_2gw"`
	AgentTermID        int           `gorm:"column:agent_term_id"`
	TranCreatedAt      *time.Time    `gorm:"column:tran_created_at"`
	TrnxID             string        `gorm:"column:trnx_id"`
	Account            string        `gorm:"column:account"`
}

func (PaymentInterchange2Cft) TableName() string {
	return "tinterchange2cft"
}

type Actlog struct {
	ID              int64      `gorm:"column:acct_id"`
	RemainBefore    float64    `gorm:"column:remain_before"`
	RemainAfter     float64    `gorm:"column:remain_after"`
	OperationAmount float64    `gorm:"column:operation_amount"`
	SourceAgent     int64      `gorm:"column:source_agentid"`
	TargetAgent     int64      `gorm:"column:target_agentid"`
	Desc            string     `gorm:"column:reason_desc"`
	CreatedBy       int64      `gorm:"column:created_by"`
	ActType         string     `gorm:"column:acct_type"`
	Vendor          *int64     `gorm:"column:vendor_id"`
	CreatedAt       *time.Time `gorm:"column:created_at;timestamp"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;timestamp"`
}

func (Actlog) TableName() string {
	return "tacctlog"
}

type TerminalStatus struct {
	EndpointID        int64      `gorm:"column:endpoint_id"`
	EndpointNum       int64      `gorm:"column:endpoint_num"`
	Phone             string     `gorm:"column:phone"`
	Address           string     `gorm:"column:address"`
	LastPayment       *time.Time `gorm:"column:last_created_payment"`
	LastPing          *time.Time `gorm:"column:lastping"`
	Region_id         int64      `gorm:"column:region_id"`
	RegionName        string     `gorm:"column:name"`
	Status            string     `gorm:"-"`
	LastPaymentDetail string     `gorm:"-"`
	DetailStatus      string     `gorm:"column:status"`
	EndpointDisabled  bool       `gorm:"column:endpoint_disabled"`
}

type Tendpoint2Service struct {
	EndpointID  int64      `gorm:"column:endpoint_id"`
	ServiceID   int64      `gorm:"column:service_id"`
	Disabled    bool       `gorm:"column:disabled"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	ModifiedtAt *time.Time `gorm:"column:modified_at"`
	CreatedBy   *int64     `gorm:"column:created_by"`
	ModifiedBy  *int64     `gorm:"column:modified_by"`
}

type TendpntServ struct {
	Old Tendpoint2Service
	New Tendpoint2Service
}

func (Tendpoint2Service) TableName() string {
	return "tendpoint2servicelist"
}

type PaymentLog struct {
	ID         int64     `gorm:"column:paymentlog_id;primary_key"`
	PaymentID  int64     `gorm:"column:payment_id"`
	LogDetails string    `gorm:"column:log_details"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	Status     int
}

func (PaymentLog) TableName() string {
	return "tpaymentlog"
}

type StatusTerminal struct {
	EndpointID int64      `gorm:"column:endpoint_id"`
	Status     string     `gorm:"column:status"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	Param1     string     `gorm:"column:param1"`
}

type Tping struct {
	ID         int64     `gorm:"column:ping_id"`
	EndpointID int64     `gorm:"column:endpoint_id"`
	Status     string    `gorm:"column:status"`
	Param1     string    `gorm:"column:param1"`
	Param2     string    `gorm:"column:param2"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (Tping) TableName() string {
	return "tendpointpings"
}

type Tpayment2Recast struct {
	ID              int64     `gorm:"column:recast_id;primary_key"`
	DisputPaymentID int64     `gorm:"column:disput_payment_id"`
	RecastPaymentID int64     `gorm:"column:recast_payment_id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	CreatedBy       int64     `gorm:"column:created_by"`
}

func (Tpayment2Recast) TableName() string {
	return "tpayment2recast"
}

type Commission struct {
	ID                 int64   `gorm:"column:commission_id;primary_key"`
	AgentID            *int64  `gorm:"column:agent_id"`
	ServiceID          *int64  `gorm:"column:service_id"`
	CommissionPercent  float64 `gorm:"column:commission_percent"`
	CommissionDisabled bool    `gorm:"column:commission_disabled"`
}

func (Commission) TableName() string {
	return "tcommissions"
}

type CommissionProfile struct {
	// EndpointID int64    `gorm:"column:endpoint_id"`
	// ServiceID  int64    `gorm:"column:service_id"`
	From      float64  `gorm:"column:f"`
	To        float64  `gorm:"column:t"`
	I         int64    `gorm:"column:i"`
	Value     *float64 `gorm:"column:v"`
	ProfileID int64    `gorm:"column:profile_id"`
}

func (CommissionProfile) TableName() string {
	return "tcommissionprofiles"
}

type CommProfChange struct {
	New CommissionProfile
	Old CommissionProfile
}

type EndpointCommand struct {
	ID          int64      `gorm:"column:cmd_id;primary_key"`
	EndpointID  int64      `gorm:"column:endpoint_id"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	CreatedBy   *int64     `gorm:"column:created_by"`
	ProcessedAt *time.Time `gorm:"column:processed_at"`
	Command     int64      `gorm:"column:cmd"`
	AgentID     int64      `gorm:"column:agent_id"`
	Proceeded   bool       `gorm:"column:proceeded"`
	Param1      *string    `gorm:"column:cmd_param1"`
	Param2      *string    `gorm:"column:cmd_param2"`
	Param3      *string    `gorm:"column:cmd_param3"`
	Url         *string    `gorm:"column:url"`
}

type EndpointId struct {
	EndpointId int64 `gorm:"column:endpoint_id"`
}

func (EndpointCommand) TableName() string {
	return "tendpointcmds"
}

type Correspond struct {
	PaymentID   int64  `gorm:"column:payment_id"`
	ReferenceID string `gorm:"column:reference_id"`
	ReceiptID   string `gorm:"column:receipt_id"`
	TrnxID      string `gorm:"column:trnx_id"`
}

func (Correspond) TableName() string {
	return "tpayments"
}

type IncomingTermReport struct {
	Datefrom string
	Dateto   string
	Terminal *string
	Provider *string
	Status   *string
	Agent    *string
}

type TermReportKiosk struct {
	Provider string
	Terminal int
	Count    int
	Sum      float64
}

type TermReport struct {
	EndpointNum int
	Address     string
	Sum         float64
}

type ProvidersReport struct {
	ServiceID   int
	ServiceDesc string
	Count       int
	Sum         float64
}

type GWReport struct {
	GatewayUnique string
	Count         int
	Sum           float64
}

type PingReport struct {
	EndpointID        int        `gorm:"endpoint_id" ru:"ID терминала"`
	EndpointNum       int        `ru:"Номер терминала"`
	Address           string     `ru:"Адрес"`
	CreatedAt         *time.Time `json:"-"`
	Duration          string     `ru:"Простой"`
	ActiveTime        float64    `gorm:"-" json:"-"`
	ActiveTimePercent string     `gorm:"-" ru:"% доступности"`
}

type ServiceArray struct {
	Key      int   `binding:"required"`
	Services []int `binding:"required"`
}

type TermMoneyReport struct {
	EndpointID  int
	Status      string
	EndpointNum string
	Address     string
	Sum         float64 `gorm:"-" json:"-"`
	RegionID    int64
}

type TermMoneyReportXLS struct {
	EndpointID  int
	Status      string
	EndpointNum string
	Address     string
	Sum         float64 `gorm:"-" json:"-"`
	Region      string
}

type TermReq struct {
	Date     string `json:"date" binding:"required"`
	RegionID *int   `json:"region_id"`
}

type Comprofile2Service struct {
	EndpointID int64 `gorm:"endpoint_id"`
	ServiceID  int64 `gorm:"service_id"`
	ProfileID  int   `gorm:"profile_id"`
}

func (Comprofile2Service) TableName() string {
	return "tendpoint2commissonprofile"
}

type Com2ServiceReq struct {
	Service    int
	Endpoint   int
	ProfileID  int
	Exceptions []int
}

type ReestrHead struct {
	STRNO        int
	MFORIC       string
	MFOKNA       string
	RSTDATE      time.Time
	QNTPAYS      int
	SUMPAYS      int
	NUMPLAT      int
	REC_ACCOUNT  string
	REM_ACCOUNT  string
	REMARKS_PLAT string
	REESTRBODY   *[]ReestrBody
}

type ReestrBody struct {
	STRNO    string
	BANKNO   string
	PAYTYPE  string
	DOCNO    string
	DOCTIME  time.Time
	AMOUNT   int
	REMARKS  string
	CLID     string
	REMNAME  string
	REMDATE  string
	REMPLACE string
	TAXREM   string
	PASPTREM string
	MFOREC   string
	ACCREC   string
	TAXREC   string
	RECNAME  string
	OKATOREC string
	KPPREC   string
	CBC      string
	AGRNO    string
	CARDNO   string
	CARDEXP  string
	SYSID    string
	SYSNO    string
	TERID    string
	TERNO    string
	TERLOC   string
	TERADR   string
	Check    bool
}
type ReestrAns struct {
	Head ReestrAnswHead
	Body *[]ReestrAnwsBody
}

type TempTable struct {
	// gorm.Model
	TrnxID string `gorm:"column:trnx_id" `
	Bankno string `gorm:"column:bankno"`
}

func (TempTable) TableName() string {
	return "temptable"
}

type Device struct {
	ID          int64  `gorm:"column:id;primary_key"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

func (Device) TableName() string {
	return "tdevices"
}

type Endpoint2Device struct {
	EndpointID  int64      `gorm:"column:endpoint_id"`
	EndpointNum int64      `gorm:"-"`
	DeviceID    int64      `gorm:"column:device_id"`
	Disabled    bool       `gorm:"column:disabled"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	ModifiedAt  *time.Time `gorm:"column:modified_at"`
	CreatedBy   int        `gorm:"column:created_by"`
	ModifiedBy  int        `gorm:"column:modified_by"`
}

func (Endpoint2Device) TableName() string {
	return "tendpoint2devices"
}

type ReestrAnswHead struct {
	STRNO        int
	MFOKNA       string
	MFORIC       string
	RSTDATE      string
	QNTPAYS      int
	SUMPAYS      int
	SUMAUTHPAYS  int
	NUMPLAT      int
	REC_ACCOUNT  string
	REM_ACCOUNT  string
	REMARKS_PLAT string
}

type ReestrAnwsBody struct {
	STRNO   int
	BANKNO  string
	PAYSTAT int
	Comment string
}

type Rules struct {
	ID               int64   `json:"id,omitempty"`
	Name             string  `json:"name,omitempty"`
	From             []int64 `json:"from,omitempty"`
	To               []int64 `json:"to,omitempty"`
	Account          string  `json:"account,omitempty"`
	Priority         int     `json:"priority,omitempty"`
	Active           *bool   `json:"active,omitempty"`
	Limit            float64 `json:"limit,omitempty"`
	CustomDay        int     `json:"custom_day,omitempty"`
	TransactionCount int64   `json:"transaction_count,omitempty"`
	Period           string  `json:"period,omitempty"`
	CheckField       string  `json:"check_field,omitempty"`
	ForAccount       bool    `json:"for_account,omitempty"`
}
type BlackList struct {
	ID      int64
	To      pq.Int64Array
	Account string
	Active  *bool
	Actor   string
}

type ActionHistory struct {
	ID          int64
	BlacklistID int64
	Action      string
	ActedBy     string
	CreatedAt   *time.Time
}

type ServiceListView struct {
	ServiceID       int    `gorm:"column:service_id" json:"serviceID" ru:"ServiceID"`
	ServiceDesc     string `gorm:"column:service_desc" json:"serviceDesc" ru:"Описание"`
	VendorDesc      string `gorm:"column:vendor_desc" json:"vendorDesc" ru:"Контрагент"`
	CategoryDesc    string `gorm:"column:name" json:"categoryDesc" ru:"Тип услуги"`
	SubCategoryDesc string `gorm:"column:subcategorydesc" json:"subCategoryDesc" ru:"Подтип услуги"`
	MssCategoryName string `gorm:"column:msscategoryname" json:"mssCategoryName" ru:"Тип услуги - MCC"`
	MssName         string `gorm:"column:mssname" json:"mssName" ru:"Подтип услуги MCC"`
	Description     string `gorm:"column:description" json:"description" ru:"MCC Description"`
	MccGrouping     string `gorm:"column:mccgrouping" json:"mccGrouping" ru:"MCC Grouping"`
	MccCode         int    `gorm:"column:mcccode" json:"mccCode" ru:"MCC Code"`
	GatewayUniq     string `gorm:"column:gateway_unique" json:"gateway_uniq" ru:"Шлюз"`
	MatchingID      string `gorm:"column:matching_id" json:"matching_id" ru:"ID контрагента"`
}

func (ServiceListView) TableName() string {
	return "service"
}

type CategoryOfSerivce struct {
	ID   int    `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (CategoryOfSerivce) TableName() string {
	return "tservice_categories"
}

type SubCategoryOfSerivce struct {
	ID         int    `gorm:"column:id;primary_key" json:"id" `
	Name       string `gorm:"column:name" json:"name"`
	CategoryID int    `gorm:"column:category_id" json:"categoryID" `
}

func (SubCategoryOfSerivce) TableName() string {
	return "tservice_sub_categories"
}

type CategoryOfSerivceMCC struct {
	ID    int    `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Group string `gorm:"column:group" json:"group"`
}

func (CategoryOfSerivceMCC) TableName() string {
	return "tservice_mcc"
}

type SubCategoryOfSerivceMCC struct {
	ID          int    `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	MccID       int    `gorm:"column:mcc_id" json:"mccID"`
	Description string `gorm:"column:description" json:"description"`
	Code        int    `gorm:"column:code" json:"code"`
}

func (SubCategoryOfSerivceMCC) TableName() string {
	return "tservice_sub_mcc"
}

type RespCompliance struct {
	Data []DataCompliance `json:"data"`
}

type DataCompliance struct {
	ID             int     `json:"ID"`
	CreatedAt      string  `json:"CreatedAt"`
	Account        string  `json:"Acount"`
	Amount         float64 `json:"Amount"`
	Currency       string  `json:"Currency"`
	IDSeriesNumber string  `json:"IDSeriesNumber"`
	SenderFio      string  `json:"SenderFio"`
	SenderBirthDay string  `json:"SenderBirthDay"`
	Check          bool    `json:"Check"`
	AgentID        int     `json:"AgentID"`
	OperID         string  `json:"OperID"`
	TrnxID         string  `json:"TrnxID"`
}
