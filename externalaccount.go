package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type ExternalAccountService struct {
	Options []option.RequestOption
}

func NewExternalAccountService(opts ...option.RequestOption) (r *ExternalAccountService) {
	r = &ExternalAccountService{}
	r.Options = opts
	return
}

// create external account
func (r *ExternalAccountService) New(ctx context.Context, body ExternalAccountNewParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/external_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// show external account
func (r *ExternalAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update external account
func (r *ExternalAccountService) Update(ctx context.Context, id string, body ExternalAccountUpdateParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list external accounts
func (r *ExternalAccountService) List(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) (res *shared.Page[ExternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/external_accounts"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// list external accounts
func (r *ExternalAccountService) ListAutoPaging(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[ExternalAccount] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete external account
func (r *ExternalAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// complete verification of external account
func (r *ExternalAccountService) CompleteVerification(ctx context.Context, id string, body ExternalAccountCompleteVerificationParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s/complete_verification", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// verify external account
func (r *ExternalAccountService) Verify(ctx context.Context, id string, body ExternalAccountVerifyParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s/verify", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExternalAccount struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Can be `checking`, `savings` or `other`.
	AccountType ExternalAccountType `json:"account_type,required"`
	// Either `individual` or `business`.
	PartyType ExternalAccountPartyType `json:"party_type,required,nullable"`
	// The address associated with the owner or `null`.
	PartyAddress ExternalAccountPartyAddress `json:"party_address,required,nullable"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           string          `json:"name,required,nullable"`
	CounterpartyID string          `json:"counterparty_id,required,nullable" format:"uuid"`
	AccountDetails []AccountDetail `json:"account_details,required"`
	RoutingDetails []RoutingDetail `json:"routing_details,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The legal name of the entity which owns the account.
	PartyName      string                          `json:"party_name,required"`
	ContactDetails []ExternalAccountContactDetails `json:"contact_details,required"`
	// If the external account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID    string                            `json:"ledger_account_id,required,nullable" format:"uuid"`
	VerificationStatus ExternalAccountVerificationStatus `json:"verification_status,required"`
	JSON               ExternalAccountJSON
}

type ExternalAccountJSON struct {
	ID                 apijson.Metadata
	Object             apijson.Metadata
	LiveMode           apijson.Metadata
	CreatedAt          apijson.Metadata
	UpdatedAt          apijson.Metadata
	DiscardedAt        apijson.Metadata
	AccountType        apijson.Metadata
	PartyType          apijson.Metadata
	PartyAddress       apijson.Metadata
	Name               apijson.Metadata
	CounterpartyID     apijson.Metadata
	AccountDetails     apijson.Metadata
	RoutingDetails     apijson.Metadata
	Metadata           apijson.Metadata
	PartyName          apijson.Metadata
	ContactDetails     apijson.Metadata
	LedgerAccountID    apijson.Metadata
	VerificationStatus apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccount using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalAccountType string

const (
	ExternalAccountTypeCash        ExternalAccountType = "cash"
	ExternalAccountTypeChecking    ExternalAccountType = "checking"
	ExternalAccountTypeLoan        ExternalAccountType = "loan"
	ExternalAccountTypeNonResident ExternalAccountType = "non_resident"
	ExternalAccountTypeOther       ExternalAccountType = "other"
	ExternalAccountTypeOverdraft   ExternalAccountType = "overdraft"
	ExternalAccountTypeSavings     ExternalAccountType = "savings"
)

type ExternalAccountPartyType string

const (
	ExternalAccountPartyTypeBusiness   ExternalAccountPartyType = "business"
	ExternalAccountPartyTypeIndividual ExternalAccountPartyType = "individual"
)

type ExternalAccountPartyAddress struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	Line1     string    `json:"line1,required,nullable"`
	Line2     string    `json:"line2,required,nullable"`
	// Locality or City.
	Locality string `json:"locality,required,nullable"`
	// Region or State.
	Region string `json:"region,required,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required,nullable"`
	JSON    ExternalAccountPartyAddressJSON
}

type ExternalAccountPartyAddressJSON struct {
	ID         apijson.Metadata
	Object     apijson.Metadata
	LiveMode   apijson.Metadata
	CreatedAt  apijson.Metadata
	UpdatedAt  apijson.Metadata
	Line1      apijson.Metadata
	Line2      apijson.Metadata
	Locality   apijson.Metadata
	Region     apijson.Metadata
	PostalCode apijson.Metadata
	Country    apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccountPartyAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ExternalAccountPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalAccountContactDetails struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              bool                                               `json:"live_mode,required"`
	CreatedAt             time.Time                                          `json:"created_at,required" format:"date-time"`
	UpdatedAt             time.Time                                          `json:"updated_at,required" format:"date-time"`
	DiscardedAt           time.Time                                          `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     string                                             `json:"contact_identifier,required"`
	ContactIdentifierType ExternalAccountContactDetailsContactIdentifierType `json:"contact_identifier_type,required"`
	JSON                  ExternalAccountContactDetailsJSON
}

type ExternalAccountContactDetailsJSON struct {
	ID                    apijson.Metadata
	Object                apijson.Metadata
	LiveMode              apijson.Metadata
	CreatedAt             apijson.Metadata
	UpdatedAt             apijson.Metadata
	DiscardedAt           apijson.Metadata
	ContactIdentifier     apijson.Metadata
	ContactIdentifierType apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccountContactDetails
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ExternalAccountContactDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalAccountContactDetailsContactIdentifierType string

const (
	ExternalAccountContactDetailsContactIdentifierTypeEmail       ExternalAccountContactDetailsContactIdentifierType = "email"
	ExternalAccountContactDetailsContactIdentifierTypePhoneNumber ExternalAccountContactDetailsContactIdentifierType = "phone_number"
	ExternalAccountContactDetailsContactIdentifierTypeWebsite     ExternalAccountContactDetailsContactIdentifierType = "website"
)

type ExternalAccountVerificationStatus string

const (
	ExternalAccountVerificationStatusPendingVerification ExternalAccountVerificationStatus = "pending_verification"
	ExternalAccountVerificationStatusUnverified          ExternalAccountVerificationStatus = "unverified"
	ExternalAccountVerificationStatusVerified            ExternalAccountVerificationStatus = "verified"
)

type ExternalAccountNewParams struct {
	// Can be `checking`, `savings` or `other`.
	AccountType field.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType field.Field[ExternalAccountNewParamsPartyType] `json:"party_type,nullable"`
	// Required if receiving wire payments.
	PartyAddress field.Field[ExternalAccountNewParamsPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           field.Field[string]                                   `json:"name,nullable"`
	CounterpartyID field.Field[string]                                   `json:"counterparty_id,required,nullable" format:"uuid"`
	AccountDetails field.Field[[]ExternalAccountNewParamsAccountDetails] `json:"account_details"`
	RoutingDetails field.Field[[]ExternalAccountNewParamsRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       field.Field[string] `json:"party_name"`
	PartyIdentifier field.Field[string] `json:"party_identifier"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken field.Field[string]                                   `json:"plaid_processor_token"`
	ContactDetails      field.Field[[]ExternalAccountNewParamsContactDetails] `json:"contact_details"`
}

// MarshalJSON serializes ExternalAccountNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ExternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountNewParamsPartyType string

const (
	ExternalAccountNewParamsPartyTypeBusiness   ExternalAccountNewParamsPartyType = "business"
	ExternalAccountNewParamsPartyTypeIndividual ExternalAccountNewParamsPartyType = "individual"
)

type ExternalAccountNewParamsPartyAddress struct {
	Line1 field.Field[string] `json:"line1,nullable"`
	Line2 field.Field[string] `json:"line2,nullable"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,nullable"`
	// Region or State.
	Region field.Field[string] `json:"region,nullable"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,nullable"`
}

type ExternalAccountNewParamsAccountDetails struct {
	AccountNumber     field.Field[string]                                                  `json:"account_number,required"`
	AccountNumberType field.Field[ExternalAccountNewParamsAccountDetailsAccountNumberType] `json:"account_number_type"`
}

type ExternalAccountNewParamsAccountDetailsAccountNumberType string

const (
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban          ExternalAccountNewParamsAccountDetailsAccountNumberType = "iban"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeClabe         ExternalAccountNewParamsAccountDetailsAccountNumberType = "clabe"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeWalletAddress ExternalAccountNewParamsAccountDetailsAccountNumberType = "wallet_address"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypePan           ExternalAccountNewParamsAccountDetailsAccountNumberType = "pan"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeOther         ExternalAccountNewParamsAccountDetailsAccountNumberType = "other"
)

type ExternalAccountNewParamsRoutingDetails struct {
	RoutingNumber     field.Field[string]                                                  `json:"routing_number,required"`
	RoutingNumberType field.Field[ExternalAccountNewParamsRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       field.Field[ExternalAccountNewParamsRoutingDetailsPaymentType]       `json:"payment_type"`
}

type ExternalAccountNewParamsRoutingDetailsRoutingNumberType string

const (
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba          ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "aba"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAuBsb        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "au_bsb"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeBrCodigo     ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "br_codigo"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCaCpa        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "ca_cpa"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeChips        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "chips"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCnaps        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "cnaps"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeGBSortCode   ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "gb_sort_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeInIfsc       ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "in_ifsc"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeMyBranchCode ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "my_branch_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSwift        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "swift"
)

type ExternalAccountNewParamsRoutingDetailsPaymentType string

const (
	ExternalAccountNewParamsRoutingDetailsPaymentTypeACH         ExternalAccountNewParamsRoutingDetailsPaymentType = "ach"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeAuBecs      ExternalAccountNewParamsRoutingDetailsPaymentType = "au_becs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBacs        ExternalAccountNewParamsRoutingDetailsPaymentType = "bacs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBook        ExternalAccountNewParamsRoutingDetailsPaymentType = "book"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCard        ExternalAccountNewParamsRoutingDetailsPaymentType = "card"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCheck       ExternalAccountNewParamsRoutingDetailsPaymentType = "check"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeEft         ExternalAccountNewParamsRoutingDetailsPaymentType = "eft"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCrossBorder ExternalAccountNewParamsRoutingDetailsPaymentType = "cross_border"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeInterac     ExternalAccountNewParamsRoutingDetailsPaymentType = "interac"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeMasav       ExternalAccountNewParamsRoutingDetailsPaymentType = "masav"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeNeft        ExternalAccountNewParamsRoutingDetailsPaymentType = "neft"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeProvxchange ExternalAccountNewParamsRoutingDetailsPaymentType = "provxchange"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeRtp         ExternalAccountNewParamsRoutingDetailsPaymentType = "rtp"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSen         ExternalAccountNewParamsRoutingDetailsPaymentType = "sen"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSepa        ExternalAccountNewParamsRoutingDetailsPaymentType = "sepa"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSignet      ExternalAccountNewParamsRoutingDetailsPaymentType = "signet"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeWire        ExternalAccountNewParamsRoutingDetailsPaymentType = "wire"
)

type ExternalAccountNewParamsContactDetails struct {
	ContactIdentifier     field.Field[string]                                                      `json:"contact_identifier"`
	ContactIdentifierType field.Field[ExternalAccountNewParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

type ExternalAccountNewParamsContactDetailsContactIdentifierType string

const (
	ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail       ExternalAccountNewParamsContactDetailsContactIdentifierType = "email"
	ExternalAccountNewParamsContactDetailsContactIdentifierTypePhoneNumber ExternalAccountNewParamsContactDetailsContactIdentifierType = "phone_number"
	ExternalAccountNewParamsContactDetailsContactIdentifierTypeWebsite     ExternalAccountNewParamsContactDetailsContactIdentifierType = "website"
)

type ExternalAccountUpdateParams struct {
	// Either `individual` or `business`.
	PartyType field.Field[ExternalAccountUpdateParamsPartyType] `json:"party_type,nullable"`
	// Can be `checking`, `savings` or `other`.
	AccountType    field.Field[ExternalAccountType] `json:"account_type"`
	CounterpartyID field.Field[string]              `json:"counterparty_id,nullable" format:"uuid"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name field.Field[string] `json:"name,nullable"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName    field.Field[string]                                  `json:"party_name"`
	PartyAddress field.Field[ExternalAccountUpdateParamsPartyAddress] `json:"party_address"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes ExternalAccountUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ExternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountUpdateParamsPartyType string

const (
	ExternalAccountUpdateParamsPartyTypeBusiness   ExternalAccountUpdateParamsPartyType = "business"
	ExternalAccountUpdateParamsPartyTypeIndividual ExternalAccountUpdateParamsPartyType = "individual"
)

type ExternalAccountUpdateParamsPartyAddress struct {
	Line1 field.Field[string] `json:"line1,nullable"`
	Line2 field.Field[string] `json:"line2,nullable"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,nullable"`
	// Region or State.
	Region field.Field[string] `json:"region,nullable"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,nullable"`
}

type ExternalAccountListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// Searches the ExternalAccount's party_name AND the Counterparty's party_name
	PartyName      field.Field[string] `query:"party_name"`
	CounterpartyID field.Field[string] `query:"counterparty_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes ExternalAccountListParams into a url.Values of the query
// parameters associated with this value
func (r ExternalAccountListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ExternalAccountCompleteVerificationParams struct {
	Amounts field.Field[[]int64] `json:"amounts"`
}

// MarshalJSON serializes ExternalAccountCompleteVerificationParams into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r ExternalAccountCompleteVerificationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountVerifyParams struct {
	// The ID of the internal account where the micro-deposits originate from. Both
	// credit and debit capabilities must be enabled.
	OriginatingAccountID field.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// Both ach and eft are supported payment types.
	PaymentType field.Field[ExternalAccountVerifyParamsPaymentType] `json:"payment_type,required"`
	// Defaults to the currency of the originating account.
	Currency field.Field[shared.Currency] `json:"currency,nullable"`
}

// MarshalJSON serializes ExternalAccountVerifyParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ExternalAccountVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountVerifyParamsPaymentType string

const (
	ExternalAccountVerifyParamsPaymentTypeACH         ExternalAccountVerifyParamsPaymentType = "ach"
	ExternalAccountVerifyParamsPaymentTypeAuBecs      ExternalAccountVerifyParamsPaymentType = "au_becs"
	ExternalAccountVerifyParamsPaymentTypeBacs        ExternalAccountVerifyParamsPaymentType = "bacs"
	ExternalAccountVerifyParamsPaymentTypeBook        ExternalAccountVerifyParamsPaymentType = "book"
	ExternalAccountVerifyParamsPaymentTypeCard        ExternalAccountVerifyParamsPaymentType = "card"
	ExternalAccountVerifyParamsPaymentTypeCheck       ExternalAccountVerifyParamsPaymentType = "check"
	ExternalAccountVerifyParamsPaymentTypeCrossBorder ExternalAccountVerifyParamsPaymentType = "cross_border"
	ExternalAccountVerifyParamsPaymentTypeEft         ExternalAccountVerifyParamsPaymentType = "eft"
	ExternalAccountVerifyParamsPaymentTypeInterac     ExternalAccountVerifyParamsPaymentType = "interac"
	ExternalAccountVerifyParamsPaymentTypeMasav       ExternalAccountVerifyParamsPaymentType = "masav"
	ExternalAccountVerifyParamsPaymentTypeNeft        ExternalAccountVerifyParamsPaymentType = "neft"
	ExternalAccountVerifyParamsPaymentTypeProvxchange ExternalAccountVerifyParamsPaymentType = "provxchange"
	ExternalAccountVerifyParamsPaymentTypeRtp         ExternalAccountVerifyParamsPaymentType = "rtp"
	ExternalAccountVerifyParamsPaymentTypeSen         ExternalAccountVerifyParamsPaymentType = "sen"
	ExternalAccountVerifyParamsPaymentTypeSepa        ExternalAccountVerifyParamsPaymentType = "sepa"
	ExternalAccountVerifyParamsPaymentTypeSignet      ExternalAccountVerifyParamsPaymentType = "signet"
	ExternalAccountVerifyParamsPaymentTypeWire        ExternalAccountVerifyParamsPaymentType = "wire"
)