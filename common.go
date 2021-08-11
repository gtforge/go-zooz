package zooz

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// Maxium level of citiation
const (
	maxCitationLevel = 1000
)

// Result represents status and category of some methods response.
type Result struct {
	Status      string `json:"status"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Description string `json:"description"`
}

// ClientInfo represents optional request params for some methods.
type ClientInfo struct {
	IPAddress string
	UserAgent string
}

// Address is a set of fields describing customer address.
type Address struct {
	Country   string `json:"country,omitempty"`
	State     string `json:"state,omitempty"`
	City      string `json:"city,omitempty"`
	Line1     string `json:"line1,omitempty"`
	Line2     string `json:"line2,omitempty"`
	ZipCode   string `json:"zip_code,omitempty"`
	Title     string `json:"title,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
}

// AdditionalDetails is a set of any custom key-value info.
type AdditionalDetails map[string]string

// ThreeDSecureAttributes is a set of attributes for 3D-Secure.
type ThreeDSecureAttributes struct {
	Encoding string `json:"encoding"`
	XID      string `json:"xid"`
	CAVV     string `json:"cavv"`
	EciFlag  string `json:"eci_flag"`
}

// ThreeDSecure13Attributes is a set of attributes for 3D-Secure 1.3.
type ThreeDSecure13Attributes struct {
	External      ExternalThreeDSecure `json:"external"`
	Internal      InternalThreeDSecure `json:"internal"`
	ScaExemptions ScaExemptions        `json:"sca_exemptions"`
}

// ExternalThreeDSecure is a set of fields describing Extra params of 3D-Secure 2
type ExternalThreeDSecure struct {
	ThreeDSecureVersion              string `json:"three_d_secure_version"`
	ThreeDSecureAuthenticationStatus string `json:"three_d_secure_authentication_status"`
	XID                              string `json:"xid"`
	DsXID                            string `json:"ds_xid"`
	Encoding                         string `json:"encoding"`
	CAVV                             string `json:"cavv"`
	EciFlag                          string `json:"eci_flag"`
	AuthenticationId                 string `json:"authentication_id"`
}

// InternalThreeDSecure is a set of fields describing Internal params of 3D-Secure 2
type InternalThreeDSecure struct {
	ThreeDSecureServerTransactionId    string `json:"three_d_secure_server_transaction_id"`
	DataCollectionCompleteInd 	       string `json:"data_collection_complete_ind"`
	DeviceChannel                      string `json:"device_channel"`
	WorkPhone                          string `json:"work_phone"`
	MobilePhone                        string `json:"mobile_phone"`
	HomePhone                          string `json:"home_phone"`
	MobilePhoneCountry                 string `json:"mobile_phone_country"`
	HomePhoneCountry                   string `json:"home_phone_country"`
	WorkPhoneCountry                   string `json:"work_phone_country"`
	AddressMatch                       string `json:"address_match"`
	ProductCode                        string `json:"product_code"`
	ShippingMethodIndicator            string `json:"shipping_method_indicator"`
	DeliveryTimeFrame                  string `json:"delivery_time_frame"`
	ReorderIndicator                   string `json:"reorder_indicator"`
	PreOrderIndicator                  string `json:"pre_order_indicator"`
	PreOrderDate                       string `json:"pre_order_date"`
	AccountAgeIndicator                string `json:"account_age_indicator"`
	AccountCreateDate                  string `json:"account_create_date"`
	AccountChangeIndicator             string `json:"account_change_indicator"`
	AccountChangeDate                  string `json:"account_change_date"`
	AccountPwdChangeIndicator          string `json:"account_pwd_change_indicator"`
	AccountPwdChangeDate               string `json:"account_pwd_change_date"`
	AccountAdditionalInformation       string `json:"account_additional_information"`
	ShippingAddressUsageIndicator      string `json:"shipping_address_usage_indicator"`
	ShippingAddressUsageDate           string `json:"shipping_address_usage_date"`
	TransactionCountDay                string `json:"transaction_count_day"`
	TransactionCountYear               string `json:"transaction_count_year"`
	AddCardAttemptsDay                 string `json:"add_card_attempts_day"`
	AccountPurchasesSixMonths          string `json:"account_purchases_six_months"`
	FraudActivity                      string `json:"fraud_activity"`
	ShippingNameIndicator              string `json:"shipping_name_indicator"`
	PaymentAccountIndicator            string `json:"payment_account_indicator"`
	PaymentAccountAge                  string `json:"payment_account_age"`
	RequestorAuthenticationMethod      string `json:"requestor_authentication_method"`
	RequestorAuthenticationTimestamp   string `json:"requestor_authentication_timestamp"`
	PriorAuthenticationData            string `json:"prior_authentication_data"`
	PriorAuthenticationMethod          string `json:"prior_authentication_method"`
	PriorAuthenticationTimestamp       string `json:"prior_authentication_timestamp"`
	PriorAuthenticationRef             string `json:"prior_authentication_ref"`
	PurchaseDateTime                   string `json:"purchase_date_time"`
	RecurringEndDate                   string `json:"recurring_end_date"`
	RecurringFrequency                 string `json:"recurring_frequency"`
	BrowserHeader                      string `json:"browser_header"`
	BrowserJavaEnabled                 string `json:"browser_java_enabled"`
	BrowserLanguage                    string `json:"browser_language"`
	BrowserColorDepth                  string `json:"browser_color_depth"`
	BrowserScreenHeight                string `json:"browser_screen_height"`
	BrowserScreenWidth                 string `json:"browser_screen_width"`
	BrowserTimeZone                    string `json:"browser_time_zone"`
	ChallengeIndicator                 string `json:"challenge_indicator"`
	ChallengeWindowSize                string `json:"challenge_window_size"`
	RequestorAuthenticationData        string `json:"requestor_authentication_data"`
	SdkAppId                           string `json:"sdk_app_id"`
	SdkEncryptedData                   string `json:"sdk_encrypted_data"`
	SdkMaxTimeout                      string `json:"sdk_max_timeout"`
	SdkReferenceNumber                 string `json:"sdk_reference_number"`
	SdkTransactionId                   string `json:"sdk_transaction_id"`
	SdkInterface                       string `json:"sdk_interface"`
	SdkUiType                          string `json:"sdk_ui_type"`
	SdkEphemeralPublicKey              string `json:"sdk_ephemeral_public_key"`
}

//ScaExemptions is a set of fields describing Exemption of 3D-Secure 2
type ScaExemptions struct {
    ExemptionAction       string `json:"exemption_action"`
    RequestExemptionStage string `json:"request_exemption_stage"`
    ExemptionReason       string `json:"exemption_reason"`
    TlaScore              string `json:"tra_score"`
}

// Installments is a set of options of installments.
type Installments struct {
	NumberOfInstallments    int64 `json:"number_of_installments"`
	FirstPaymentAmount      int64 `json:"first_payment_amount"`
	RemainingPaymentsAmount int64 `json:"remaining_payments_amount"`
}

// ProviderData is a set of params describing payment provider.
type ProviderData struct {
	ProviderName          string             `json:"provider_name"`
	ResponseCode          string             `json:"response_code"`
	Description           string             `json:"description"`
	RawResponse           DecodedJSON        `json:"raw_response"`
	AvsCode               string             `json:"avs_code"`
	AuthorizationCode     string             `json:"authorization_code"`
	TransactionID         string             `json:"transaction_id"`
	ExternalID            string             `json:"external_id"`
	Documents             []ProviderDocument `json:"documents"`
	AdditionalInformation map[string]string  `json:"additional_information"`
	NetworkTransactionID  string             `json:"network_transaction_id"`
    ThreeDSecureResult    ThreeDSecureResult `json:"three_d_secure_result"`
}

type ThreeDSecureResult struct {
    ThreeDSecureServerTransactionId  string `json:"three_d_secure_server_transaction_id"`
    XID                              string `json:"xid"`
    ThreeDSecureAuthenticationStatus string `json:"three_d_secure_authentication_status"`
    ThreeDSecureVersion              string `json:"three_d_secure_version"`
    DsXID                            string `json:"ds_xid"`
    EciFlag                          string `json:"eci_flag"`
    CAVV                             string `json:"cavv"`
}

// ProviderDocument represents provider document.
type ProviderDocument struct {
	Descriptor  string `json:"descriptor"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
	Href        string `json:"href"`
}

// PaymentMethodDetails represents payment method details for POST requests.
type PaymentMethodDetails struct {
	Type              string            `json:"type"`
	Token             string            `json:"token,omitempty"`
	CreditCardCvv     string            `json:"credit_card_cvv,omitempty"`
	SourceType        string            `json:"source_type,omitempty"`
	Vendor            string            `json:"vendor,omitempty"`
	AdditionalDetails AdditionalDetails `json:"additional_details,omitempty"`
}

// This object represents the configuration of the provider that handled the transaction,
// as defined in your PaymentsOS Control Center account.
// Note that the object does not include your provider authentication credentials.
type ProviderConfiguration struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
	ProviderID  string `json:"provider_id"`
	Type        string `json:"type"`
	AccountID   string `json:"account_id"`
	Href        string `json:"href"`
}

// The line items of the order.
type OrderLineItem struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Quantity  int    `json:"quantity"`
	UnitPrice int64  `json:"unit_price"`
}

// Details of the order. Note that order fields required for level 2 and 3 data, must be passed separately
// in a Create Capture request within a level_2_3 object (fields passed in the order object only are not recognized as level 2 and 3 data fields).
type Order struct {
	ID                string            `json:"id"`
	AdditionalDetails AdditionalDetails `json:"additional_details,omitempty"`
	LineItems         []OrderLineItem   `json:"line_items"`
	TaxAmount         int64             `json:"tax_amount"`
	TaxPercentage     decimal.Decimal   `json:"tax_percentage"`
}

type DecodedJSON map[string]interface{}

// UnmarshalJSON is called when the function json.Unmarshal is called.
func (d *DecodedJSON) UnmarshalJSON(data []byte) error {
	var (
		err           error
		unquotedData  string
		rawSet        = make(map[string]json.RawMessage)
		citationLevel = 0
	)

	unquotedData, err = strconv.Unquote(string(data))
	// Error can be returned if it's not quoted string
	if err != nil {
		// There are not quoted data, trying to unmarshal
		err = json.Unmarshal(data, &rawSet)
	} else {
		// Save iteraction result separatly because it can be nil
		var iterationUnquotedData string
		for err == nil && citationLevel < maxCitationLevel {
			iterationUnquotedData, err = strconv.Unquote(string(unquotedData))
			// Save the last successful result
			if err == nil {
				unquotedData = iterationUnquotedData
			}
			citationLevel++
		}
		err = json.Unmarshal([]byte(unquotedData), &rawSet)
	}
	if err != nil {
		return errors.Wrap(err, "can't unmarshal to map")
	}
	*d = make(map[string]interface{})
	for k, v := range rawSet {
		dValue := make(DecodedJSON)
		// Trying to unmarshal the value because it can contain another JSON object
		err = json.Unmarshal(v, &dValue)
		if err != nil {
			var unmarshaledValue interface{}
			// Trying to unmarshal value as is
			err = json.Unmarshal(v, &unmarshaledValue)
			// If unsuccessfull then we set the raw value
			if err != nil {
				(*d)[k] = v
			} else { // set unmarshaled data
				(*d)[k] = unmarshaledValue
			}
			continue
		}
		(*d)[k] = dValue

	}
	return nil
}

// Contains information about the decision flow executed by the Decision Engine and the rules that were evaluated as payments pass through the flow.
// Read more about the Decision Engine in the PaymentsOS developer guide.
type DecisionEngineExecution struct {
	ID            string         `json:"id"`
	Created       string         `json:"created"`
	FlowID        string         `json:"flow_id"`
	Status        string         `json:"status"`
	PolicyResults []PolicyResult `json:"policy_results"`
}

// Describes the results of the policy rules executed in the flow.
type PolicyResult struct {
	Type                  string `json:"type"`
	ProviderName          string `json:"provider_name"`
	ProviderConfiguration string `json:"provider_configuration"`
	Name                  string `json:"name"`
	ExecutionTime         string `json:"execution_time"`
	Transaction           string `json:"transaction"`
	Result                string `json:"result"`
}

// Level 2 and Level 3 card data provides more information for business, commercial, corporate, purchasing,
// and government cardholders. Credit card transactions submitted with Level 2 and Level 3 card data can obtain
// lower interchange rates and thus provide you with a lower processing cost.
type Level23 struct {
	OrderID             json.RawMessage `json:"id"`
	TaxMode             string          `json:"tax_mode"`
	TaxAmount           int64           `json:"tax_amount"`
	ShippingAmount      int64           `json:"shipping_amount"`
	FromShippingZipCode string          `json:"from_shipping_zip_code"`
	DutyAmount          int64           `json:"duty_amount"`
	DiscountAmount      int64           `json:"discount_amount"`
	LineItems           []LevelLineItem `json:"line_items"`
	ShippingAddress     *Address        `json:"shipping_address"`
}

// The line items of the order.
type LevelLineItem struct {
	Name               string          `json:"name"`
	ID                 string          `json:"id"`
	Quantity           int             `json:"quantity"`
	UnitPrice          int64           `json:"unit_price"`
	CommodityCode      string          `json:"commodity_code"`
	UnitOfMeasure      string          `json:"unit_of_measure"`
	TaxMode            string          `json:"tax_mode"`
	TaxAmount          int64           `json:"tax_amount"`
	TaxPercentage      decimal.Decimal `json:"tax_percentage"`
	DiscountAmount     int64           `json:"discount_amount"`
	DiscountPercentage decimal.Decimal `json:"discount_percentage"`
	SubTotal           int64           `json:"sub_total"`
}
