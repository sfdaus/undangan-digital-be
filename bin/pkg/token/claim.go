package token

type Claim struct {
	ProfileCode  string   `json:"profile_code"`
	UserID       string   `json:"userId"`
	RefreshToken string   `json:"refreshToken"`
	Key          string   `json:"key"`
	Scope        string   `json:"scope"`
	Roles        []string `json:"roles"`
	AppCode      string   `json:"appCode"`
}

type ClaimRegister struct {
	Email  string `json:"Email"`
	UserID string `json:"userId"`
	Key    string `json:"key"`
}

type ClaimRefreshToken struct {
	UserID      string `json:"userId"`
	ExpiredDate string `json:"expiredDate"`
	Key         string `json:"key"`
}

type OTP struct {
	UserID       string `json:"userId"`
	Key          string `json:"key"`
	Bracket      string `json:"bracket"`
	OTPNumber    string `json:"otpNumber"`
	Counter      int    `json:"counter"`
	Blocked      bool   `json:"blocked"`
	InitTime     string `json:"initTime"`
	FailCounter  int    `json:"failCounter"`
	PasswordTemp string `json:"passwordTemp"`
}

type ChangeEmailOTP struct {
	UserID      string `json:"userId"`
	FullName    string `json:"fullName"`
	NewEmail    string `json:"newEmail"`
	Key         string `json:"key"`
	Bracket     string `json:"bracket"`
	OTPNumber   string `json:"otpNumber"`
	Counter     int    `json:"counter"`
	Blocked     bool   `json:"blocked"`
	InitTime    string `json:"initTime"`
	FailCounter int    `json:"failCounter"`
}

type Link struct {
	Email           string `json:"email"`
	UserID          string `json:"user_id"`
	VerifyLink      string `json:"VerifyLink"`
	Counter         int    `json:"counter"`
	InitTime        string `json:"initTime"`
	PartnershipMail string `json:"partnership_mail"`
}

type ClaimResend struct {
	UserID      string `json:"userId"`
	Key         string `json:"key"`
	Scope       string `json:"scope"`
	Counter     int    `json:"counter"`
	AvailableAt int    `json:"availableAt"`
}
