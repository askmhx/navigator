package model

import "time"

type AccountInf struct {
	InstitutionNo string
	AccountNo     string
	RelationNo    string
	AccountType   string
	AccountName   string
	Available     int
	Freeze        int
	Stock         int
	Abnormal      int
	Status        int
	Memo          string
	CreatedAt     time.Time
	CreatedBy     string
	UpdatedAt     time.Time
	UpdatedBy     string
}
type AccountLog struct {
	TradeSsn      string
	TradeDate     time.Time
	TradeTime     time.Time
	TradeStatus   int
	TradeAmount   int
	SrcAccountNo  string
	SrcType       int
	DestAccountNo string
	DestType      int
	BookStatus    int
	Memo          string
	CreatedAt     time.Time
	CreatedBy     string
	UpdatedAt     time.Time
	UpdatedBy     string
}
type ChannelLog struct {
	MerchantId        string
	SrcOrderDate      time.Time
	SrcOrderTime      time.Time
	SrcOrderSsn       string
	SrcOrderStatus    string
	ChannelId         int
	ChannelStatus     string
	ChannelDate       time.Time
	ChannelTime       time.Time
	ChannelSsn        string
	ChannelSettleDate time.Time
	ChannelFee        float64
	OrderAmount       float64
	OrderInf          string
	Memo              string
	CreatedAt         time.Time
	CreatedBy         string
	UpdatedAt         time.Time
	UpdatedBy         string
}
type EmailLog struct {
	MerchantId string
	SrcModule  string
	SendDate   time.Time
	SendTime   time.Time
	Email      string
	Content    string
	Status     int
	Memo       string
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}
type SmsLog struct {
	MerchantId string
	SrcModule  string
	SendDate   time.Time
	SendTime   time.Time
	Mobile     string
	Content    string
	Status     int
	Channel    string
	Memo       string
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}
type TradeLog struct {
	MerchantId      string
	SubMerchantId   string
	StoreId         string
	SrcOrderDate    time.Time
	SrcOrderTime    time.Time
	SrcOrderSsn     string
	SrcOrderStatus  int
	OrderDate       time.Time
	OrderTime       time.Time
	OrderSsn        string
	OrderStatus     int
	OrderSettleDate time.Time
	OrderType       string
	OrderAmount     float64
	OrderFee        float64
	OrderRealReward float64
	OrderCalcReward float64
	OrderEvent      string
	OrderInf        string
	DeviceId        string
	DeviceIp        string
	DeviceLongitude string
	DeviceLatitude  string
	UserId          string
	AttachInf       string
	ReturnUrl       string
	NotifyUrl       string
	NotifyStatus    int
	NotifyTimes     int
	NotifyLastTime  time.Time
	NotifyMessage   string
	Memo            string
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       time.Time
	UpdatedBy       string
}
type ChannelInf struct {
	Id             string
	ChannelName    string
	ChannelAccount string
	TradeType      string
	Status         int
	Memo           string
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      time.Time
	UpdatedBy      string
}
type MerchantAccount struct {
	MerchantId  string
	AccountType string
	AccountNo   string
	Memo        string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}
type MerchantAttach struct {
	MerchantId   string
	AttachName   string
	AttachType   string
	AttachPath   string
	AttachExpire time.Time
	Memo         string
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
}
type MerchantCard struct {
	MerchantId       string
	CardType         string
	CardNo           string
	CardHolderName   string
	CardHolderId     string
	CardHolderIdType string
	CardHolderMobile string
	Province         string
	City             string
	District         string
	BankNo           string
	BankBranchNo     string
	Memo             string
	CreatedAt        time.Time
	CreatedBy        string
	UpdatedAt        time.Time
	UpdatedBy        string
}
type MerchantInf struct {
	MerchantId      string
	OutMerchantId   string
	MerchantName    string
	MerchantType    string
	AgentId         string
	ShortName       string
	AliasName       string
	LicenseNo       string
	LegalPerson     string
	LegalIdType     string
	LegalId         string
	LegalMobile     string
	ProvinceCode    string
	CityCode        string
	DistrictCode    string
	Address         string
	ContactPerson   string
	ContactIdType   string
	ContactId       string
	ContactMobile   string
	CustomerManager string
	Status          int
	Memo            string
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       time.Time
	UpdatedBy       string
}
type MerchantIp struct {
	MerchantId string
	AllowIp    string
	AllowHost  string
	AllowQps   string
	Status     int
	Memo       string
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}
type MerchantKey struct {
	MerchantId string
	KeyIndex   int
	KeyType    string
	KeyData    string
	KeyExpire  time.Time
	Status     int
	Memo       string
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}
type MerchantOperator struct {
	MerchantId             string
	OperatorId             string
	OperatorMobile         string
	OperatorEmail          string
	OperatorWechat         string
	OperatorName           string
	OperatorIdCard         string
	OperatorIdCardType     string
	OperatorPassword       string
	OperatorPasswordExpire time.Time
	OperatorBindIp         string
	OperatorRole           string
	Status                 int
	Memo                   string
	CreatedAt              time.Time
	CreatedBy              string
	UpdatedAt              time.Time
	UpdatedBy              string
}
type MerchantRegion struct {
	Id         int
	ParentId   int
	Level      int
	Name       string
	MergerName string
	ShortName  string
	Pinyin     string
	ZipCode    string
	AreaCode   string
	CityCode   string
	Latitude   string
	Longitude  string
	Memo       string
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}
type MerchantRiskControl struct {
	MerchantId       string
	TradeType        string
	LimitCardType    string
	LimitTotalAmount float64
	LimitEachAmount  float64
	LimitLatitude    string
	LimitLongitude   string
	Status           int
	Memo             string
	CreatedAt        time.Time
	CreatedBy        string
	UpdatedAt        time.Time
	UpdatedBy        string
}
type MerchantStore struct {
	StoreId      string
	OutStoreId   string
	StoreName    string
	AreaCode     string
	CityCode     string
	DistrictCode string
	Address      string
	ZipCode      string
	Latitude     string
	Longitude    string
	Status       int
	Memo         string
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
}
type MerchantTrade struct {
	MerchantId  string
	TradeId     string
	TradeExpire time.Time
	TradeFeeId  string
	Status      int
	Memo        string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}
type TradeFee struct {
	FeeId       string
	CalcObjType string
	CalcObjMin  float64
	CalcObjMax  float64
	FeeType     string
	Fee         float64
	FeeMax      float64
	Memo        string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}
type TradeInf struct {
	TradeId   string
	TradeName string
	TradeType string
	Status    int
	Memo      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
type TradeRouter struct {
	Id          int
	TradeType   string
	MerchantId  string
	CalcObjType string
	CalcObjMin  float64
	CalcObjMax  float64
	MinAmount   float64
	MaxAmount   float64
	ChannelId   string
	Status      int
	Memo        string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}
