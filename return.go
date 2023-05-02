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

type ReturnService struct {
	Options []option.RequestOption
}

func NewReturnService(opts ...option.RequestOption) (r *ReturnService) {
	r = &ReturnService{}
	r.Options = opts
	return
}

// Create a return.
func (r *ReturnService) New(ctx context.Context, body ReturnNewParams, opts ...option.RequestOption) (res *ReturnObject, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/returns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single return.
func (r *ReturnService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ReturnObject, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/returns/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of returns.
func (r *ReturnService) List(ctx context.Context, query ReturnListParams, opts ...option.RequestOption) (res *shared.Page[ReturnObject], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/returns"
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

// Get a list of returns.
func (r *ReturnService) ListAutoPaging(ctx context.Context, query ReturnListParams, opts ...option.RequestOption) *shared.PageAutoPager[ReturnObject] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type ReturnObject struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the object being returned or `null`.
	ReturnableID string `json:"returnable_id,required,nullable" format:"uuid"`
	// The type of object being returned or `null`.
	ReturnableType ReturnObjectReturnableType `json:"returnable_type,required,nullable"`
	// The return code. For ACH returns, this is the required ACH return code.
	Code ReturnObjectCode `json:"code,required,nullable"`
	// Often the bank will provide an explanation for the return, which is a short
	// human readable string.
	Reason string `json:"reason,required,nullable"`
	// If the return code is `R14` or `R15` this is the date the deceased counterparty
	// passed away.
	DateOfDeath time.Time `json:"date_of_death,required,nullable" format:"date"`
	// Some returns may include additional information from the bank. In these cases,
	// this string will be present.
	AdditionalInformation string `json:"additional_information,required,nullable"`
	// The current status of the return.
	Status ReturnObjectStatus `json:"status,required"`
	// The ID of the relevant Transaction Line Item or `null`.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// The ID of the relevant Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The ID of the relevant Internal Account.
	InternalAccountID string `json:"internal_account_id,required,nullable" format:"uuid"`
	// The type of return. Can be one of: `ach`, `ach_noc`, `au_becs`, `bacs`, `eft`,
	// `interac`, `manual`, `paper_item`, `wire`.
	Type ReturnObjectType `json:"type,required"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// Currency that this transaction is denominated in.
	Currency shared.Currency `json:"currency,required,nullable"`
	// If an originating return failed to be processed by the bank, a description of
	// the failure reason will be available.
	FailureReason string `json:"failure_reason,required,nullable"`
	// The role of the return, can be `originating` or `receiving`.
	Role ReturnObjectRole `json:"role,required"`
	// An array of Payment Reference objects.
	ReferenceNumbers []ReturnObjectReferenceNumbers `json:"reference_numbers,required"`
	// The ID of the ledger transaction linked to the return.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	JSON                ReturnObjectJSON
}

type ReturnObjectJSON struct {
	ID                    apijson.Metadata
	Object                apijson.Metadata
	LiveMode              apijson.Metadata
	CreatedAt             apijson.Metadata
	UpdatedAt             apijson.Metadata
	ReturnableID          apijson.Metadata
	ReturnableType        apijson.Metadata
	Code                  apijson.Metadata
	Reason                apijson.Metadata
	DateOfDeath           apijson.Metadata
	AdditionalInformation apijson.Metadata
	Status                apijson.Metadata
	TransactionLineItemID apijson.Metadata
	TransactionID         apijson.Metadata
	InternalAccountID     apijson.Metadata
	Type                  apijson.Metadata
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	FailureReason         apijson.Metadata
	Role                  apijson.Metadata
	ReferenceNumbers      apijson.Metadata
	LedgerTransactionID   apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ReturnObject using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ReturnObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ReturnObjectReturnableType string

const (
	ReturnObjectReturnableTypeIncomingPaymentDetail ReturnObjectReturnableType = "incoming_payment_detail"
	ReturnObjectReturnableTypePaperItem             ReturnObjectReturnableType = "paper_item"
	ReturnObjectReturnableTypePaymentOrder          ReturnObjectReturnableType = "payment_order"
	ReturnObjectReturnableTypeReturn                ReturnObjectReturnableType = "return"
	ReturnObjectReturnableTypeReversal              ReturnObjectReturnableType = "reversal"
)

type ReturnObjectCode string

const (
	ReturnObjectCode901           ReturnObjectCode = "901"
	ReturnObjectCode902           ReturnObjectCode = "902"
	ReturnObjectCode903           ReturnObjectCode = "903"
	ReturnObjectCode904           ReturnObjectCode = "904"
	ReturnObjectCode905           ReturnObjectCode = "905"
	ReturnObjectCode907           ReturnObjectCode = "907"
	ReturnObjectCode908           ReturnObjectCode = "908"
	ReturnObjectCode909           ReturnObjectCode = "909"
	ReturnObjectCode910           ReturnObjectCode = "910"
	ReturnObjectCode911           ReturnObjectCode = "911"
	ReturnObjectCode912           ReturnObjectCode = "912"
	ReturnObjectCode914           ReturnObjectCode = "914"
	ReturnObjectCodeR01           ReturnObjectCode = "R01"
	ReturnObjectCodeR02           ReturnObjectCode = "R02"
	ReturnObjectCodeR03           ReturnObjectCode = "R03"
	ReturnObjectCodeR04           ReturnObjectCode = "R04"
	ReturnObjectCodeR05           ReturnObjectCode = "R05"
	ReturnObjectCodeR06           ReturnObjectCode = "R06"
	ReturnObjectCodeR07           ReturnObjectCode = "R07"
	ReturnObjectCodeR08           ReturnObjectCode = "R08"
	ReturnObjectCodeR09           ReturnObjectCode = "R09"
	ReturnObjectCodeR10           ReturnObjectCode = "R10"
	ReturnObjectCodeR11           ReturnObjectCode = "R11"
	ReturnObjectCodeR12           ReturnObjectCode = "R12"
	ReturnObjectCodeR14           ReturnObjectCode = "R14"
	ReturnObjectCodeR15           ReturnObjectCode = "R15"
	ReturnObjectCodeR16           ReturnObjectCode = "R16"
	ReturnObjectCodeR17           ReturnObjectCode = "R17"
	ReturnObjectCodeR20           ReturnObjectCode = "R20"
	ReturnObjectCodeR21           ReturnObjectCode = "R21"
	ReturnObjectCodeR22           ReturnObjectCode = "R22"
	ReturnObjectCodeR23           ReturnObjectCode = "R23"
	ReturnObjectCodeR24           ReturnObjectCode = "R24"
	ReturnObjectCodeR29           ReturnObjectCode = "R29"
	ReturnObjectCodeR31           ReturnObjectCode = "R31"
	ReturnObjectCodeR33           ReturnObjectCode = "R33"
	ReturnObjectCodeR37           ReturnObjectCode = "R37"
	ReturnObjectCodeR38           ReturnObjectCode = "R38"
	ReturnObjectCodeR39           ReturnObjectCode = "R39"
	ReturnObjectCodeR51           ReturnObjectCode = "R51"
	ReturnObjectCodeR52           ReturnObjectCode = "R52"
	ReturnObjectCodeR53           ReturnObjectCode = "R53"
	ReturnObjectCodeCurrencycloud ReturnObjectCode = "currencycloud"
)

type ReturnObjectStatus string

const (
	ReturnObjectStatusCompleted  ReturnObjectStatus = "completed"
	ReturnObjectStatusFailed     ReturnObjectStatus = "failed"
	ReturnObjectStatusPending    ReturnObjectStatus = "pending"
	ReturnObjectStatusProcessing ReturnObjectStatus = "processing"
	ReturnObjectStatusReturned   ReturnObjectStatus = "returned"
	ReturnObjectStatusSent       ReturnObjectStatus = "sent"
)

type ReturnObjectType string

const (
	ReturnObjectTypeACH       ReturnObjectType = "ach"
	ReturnObjectTypeACHNoc    ReturnObjectType = "ach_noc"
	ReturnObjectTypeAuBecs    ReturnObjectType = "au_becs"
	ReturnObjectTypeBacs      ReturnObjectType = "bacs"
	ReturnObjectTypeBook      ReturnObjectType = "book"
	ReturnObjectTypeEft       ReturnObjectType = "eft"
	ReturnObjectTypeInterac   ReturnObjectType = "interac"
	ReturnObjectTypeManual    ReturnObjectType = "manual"
	ReturnObjectTypePaperItem ReturnObjectType = "paper_item"
	ReturnObjectTypeSepa      ReturnObjectType = "sepa"
	ReturnObjectTypeWire      ReturnObjectType = "wire"
)

type ReturnObjectRole string

const (
	ReturnObjectRoleOriginating ReturnObjectRole = "originating"
	ReturnObjectRoleReceiving   ReturnObjectRole = "receiving"
)

type ReturnObjectReferenceNumbers struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The vendor reference number.
	ReferenceNumber string `json:"reference_number,required"`
	// The type of the reference number. Referring to the vendor payment id.
	ReferenceNumberType ReturnObjectReferenceNumbersReferenceNumberType `json:"reference_number_type,required"`
	JSON                ReturnObjectReferenceNumbersJSON
}

type ReturnObjectReferenceNumbersJSON struct {
	ID                  apijson.Metadata
	Object              apijson.Metadata
	LiveMode            apijson.Metadata
	CreatedAt           apijson.Metadata
	UpdatedAt           apijson.Metadata
	ReferenceNumber     apijson.Metadata
	ReferenceNumberType apijson.Metadata
	raw                 string
	Extras              map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ReturnObjectReferenceNumbers
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ReturnObjectReferenceNumbers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ReturnObjectReferenceNumbersReferenceNumberType string

const (
	ReturnObjectReferenceNumbersReferenceNumberTypeACHOriginalTraceNumber          ReturnObjectReferenceNumbersReferenceNumberType = "ach_original_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeACHTraceNumber                  ReturnObjectReferenceNumbersReferenceNumberType = "ach_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeBankprovPaymentActivityDate     ReturnObjectReferenceNumbersReferenceNumberType = "bankprov_payment_activity_date"
	ReturnObjectReferenceNumbersReferenceNumberTypeBankprovPaymentID               ReturnObjectReferenceNumbersReferenceNumberType = "bankprov_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID         ReturnObjectReferenceNumbersReferenceNumberType = "bnk_dev_prenotification_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBnkDevTransferID                ReturnObjectReferenceNumbersReferenceNumberType = "bnk_dev_transfer_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBofaEndToEndID                  ReturnObjectReferenceNumbersReferenceNumberType = "bofa_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBofaTransactionID               ReturnObjectReferenceNumbersReferenceNumberType = "bofa_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCheckNumber                     ReturnObjectReferenceNumbersReferenceNumberType = "check_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeColumnFxQuoteID                 ReturnObjectReferenceNumbersReferenceNumberType = "column_fx_quote_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeColumnTransferID                ReturnObjectReferenceNumbersReferenceNumberType = "column_transfer_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverPaymentID             ReturnObjectReferenceNumbersReferenceNumberType = "cross_river_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverTransactionID         ReturnObjectReferenceNumbersReferenceNumberType = "cross_river_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCurrencycloudConversionID       ReturnObjectReferenceNumbersReferenceNumberType = "currencycloud_conversion_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCurrencycloudPaymentID          ReturnObjectReferenceNumbersReferenceNumberType = "currencycloud_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeDcBankTransactionID             ReturnObjectReferenceNumbersReferenceNumberType = "dc_bank_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeDwollaTransactionID             ReturnObjectReferenceNumbersReferenceNumberType = "dwolla_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeEftTraceNumber                  ReturnObjectReferenceNumbersReferenceNumberType = "eft_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeFedwireImad                     ReturnObjectReferenceNumbersReferenceNumberType = "fedwire_imad"
	ReturnObjectReferenceNumbersReferenceNumberTypeFedwireOmad                     ReturnObjectReferenceNumbersReferenceNumberType = "fedwire_omad"
	ReturnObjectReferenceNumbersReferenceNumberTypeFirstRepublicInternalID         ReturnObjectReferenceNumbersReferenceNumberType = "first_republic_internal_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsCollectionRequestID ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_collection_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsEndToEndID          ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsPaymentRequestID    ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_payment_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsRequestID           ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsUniquePaymentID     ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_unique_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeInteracMessageID                ReturnObjectReferenceNumbersReferenceNumberType = "interac_message_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcCcn                         ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_ccn"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcCustomerReferenceID         ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_customer_reference_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcEndToEndID                  ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcFirmRootID                  ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_firm_root_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcP3ID                        ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_p3_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentBatchID              ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_payment_batch_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentInformationID        ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_payment_information_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeLobCheckID                      ReturnObjectReferenceNumbersReferenceNumberType = "lob_check_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeOther                           ReturnObjectReferenceNumbersReferenceNumberType = "other"
	ReturnObjectReferenceNumbersReferenceNumberTypePartialSwiftMir                 ReturnObjectReferenceNumbersReferenceNumberType = "partial_swift_mir"
	ReturnObjectReferenceNumbersReferenceNumberTypePncClearingReference            ReturnObjectReferenceNumbersReferenceNumberType = "pnc_clearing_reference"
	ReturnObjectReferenceNumbersReferenceNumberTypePncInstructionID                ReturnObjectReferenceNumbersReferenceNumberType = "pnc_instruction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypePncMultipaymentID               ReturnObjectReferenceNumbersReferenceNumberType = "pnc_multipayment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypePncPaymentTraceID               ReturnObjectReferenceNumbersReferenceNumberType = "pnc_payment_trace_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeRtpInstructionID                ReturnObjectReferenceNumbersReferenceNumberType = "rtp_instruction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetAPIReferenceID            ReturnObjectReferenceNumbersReferenceNumberType = "signet_api_reference_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetConfirmationID            ReturnObjectReferenceNumbersReferenceNumberType = "signet_confirmation_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetRequestID                 ReturnObjectReferenceNumbersReferenceNumberType = "signet_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSilvergatePaymentID             ReturnObjectReferenceNumbersReferenceNumberType = "silvergate_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSwiftMir                        ReturnObjectReferenceNumbersReferenceNumberType = "swift_mir"
	ReturnObjectReferenceNumbersReferenceNumberTypeSwiftUetr                       ReturnObjectReferenceNumbersReferenceNumberType = "swift_uetr"
	ReturnObjectReferenceNumbersReferenceNumberTypeUsbankPaymentID                 ReturnObjectReferenceNumbersReferenceNumberType = "usbank_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeWellsFargoPaymentID             ReturnObjectReferenceNumbersReferenceNumberType = "wells_fargo_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeWellsFargoTraceNumber           ReturnObjectReferenceNumbersReferenceNumberType = "wells_fargo_trace_number"
)

type ReturnNewParams struct {
	// The ID of the object being returned or `null`.
	ReturnableID field.Field[string] `json:"returnable_id,required,nullable" format:"uuid"`
	// The return code. For ACH returns, this is the required ACH return code.
	Code field.Field[ReturnNewParamsCode] `json:"code,nullable"`
	// An optional description of the reason for the return. This is for internal usage
	// and will not be transmitted to the bank.”
	Reason field.Field[string] `json:"reason,nullable"`
	// If the return code is `R14` or `R15` this is the date the deceased counterparty
	// passed away.
	DateOfDeath field.Field[time.Time] `json:"date_of_death,nullable" format:"date"`
	// Some returns may include additional information from the bank. In these cases,
	// this string will be present.
	AdditionalInformation field.Field[string] `json:"additional_information,nullable"`
	// The type of object being returned. Currently, this may only be
	// incoming_payment_detail.
	ReturnableType field.Field[ReturnNewParamsReturnableType] `json:"returnable_type,required"`
}

// MarshalJSON serializes ReturnNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r ReturnNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReturnNewParamsCode string

const (
	ReturnNewParamsCode901           ReturnNewParamsCode = "901"
	ReturnNewParamsCode902           ReturnNewParamsCode = "902"
	ReturnNewParamsCode903           ReturnNewParamsCode = "903"
	ReturnNewParamsCode904           ReturnNewParamsCode = "904"
	ReturnNewParamsCode905           ReturnNewParamsCode = "905"
	ReturnNewParamsCode907           ReturnNewParamsCode = "907"
	ReturnNewParamsCode908           ReturnNewParamsCode = "908"
	ReturnNewParamsCode909           ReturnNewParamsCode = "909"
	ReturnNewParamsCode910           ReturnNewParamsCode = "910"
	ReturnNewParamsCode911           ReturnNewParamsCode = "911"
	ReturnNewParamsCode912           ReturnNewParamsCode = "912"
	ReturnNewParamsCode914           ReturnNewParamsCode = "914"
	ReturnNewParamsCodeR01           ReturnNewParamsCode = "R01"
	ReturnNewParamsCodeR02           ReturnNewParamsCode = "R02"
	ReturnNewParamsCodeR03           ReturnNewParamsCode = "R03"
	ReturnNewParamsCodeR04           ReturnNewParamsCode = "R04"
	ReturnNewParamsCodeR05           ReturnNewParamsCode = "R05"
	ReturnNewParamsCodeR06           ReturnNewParamsCode = "R06"
	ReturnNewParamsCodeR07           ReturnNewParamsCode = "R07"
	ReturnNewParamsCodeR08           ReturnNewParamsCode = "R08"
	ReturnNewParamsCodeR09           ReturnNewParamsCode = "R09"
	ReturnNewParamsCodeR10           ReturnNewParamsCode = "R10"
	ReturnNewParamsCodeR11           ReturnNewParamsCode = "R11"
	ReturnNewParamsCodeR12           ReturnNewParamsCode = "R12"
	ReturnNewParamsCodeR14           ReturnNewParamsCode = "R14"
	ReturnNewParamsCodeR15           ReturnNewParamsCode = "R15"
	ReturnNewParamsCodeR16           ReturnNewParamsCode = "R16"
	ReturnNewParamsCodeR17           ReturnNewParamsCode = "R17"
	ReturnNewParamsCodeR20           ReturnNewParamsCode = "R20"
	ReturnNewParamsCodeR21           ReturnNewParamsCode = "R21"
	ReturnNewParamsCodeR22           ReturnNewParamsCode = "R22"
	ReturnNewParamsCodeR23           ReturnNewParamsCode = "R23"
	ReturnNewParamsCodeR24           ReturnNewParamsCode = "R24"
	ReturnNewParamsCodeR29           ReturnNewParamsCode = "R29"
	ReturnNewParamsCodeR31           ReturnNewParamsCode = "R31"
	ReturnNewParamsCodeR33           ReturnNewParamsCode = "R33"
	ReturnNewParamsCodeR37           ReturnNewParamsCode = "R37"
	ReturnNewParamsCodeR38           ReturnNewParamsCode = "R38"
	ReturnNewParamsCodeR39           ReturnNewParamsCode = "R39"
	ReturnNewParamsCodeR51           ReturnNewParamsCode = "R51"
	ReturnNewParamsCodeR52           ReturnNewParamsCode = "R52"
	ReturnNewParamsCodeR53           ReturnNewParamsCode = "R53"
	ReturnNewParamsCodeCurrencycloud ReturnNewParamsCode = "currencycloud"
)

type ReturnNewParamsReturnableType string

const (
	ReturnNewParamsReturnableTypeIncomingPaymentDetail ReturnNewParamsReturnableType = "incoming_payment_detail"
)

type ReturnListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// Specify `internal_account_id` if you wish to see returns to/from a specific
	// account.
	InternalAccountID field.Field[string] `query:"internal_account_id"`
	// Specify `counterparty_id` if you wish to see returns that occurred with a
	// specific counterparty.
	CounterpartyID field.Field[string] `query:"counterparty_id"`
	// The ID of a valid returnable. Must be accompanied by `returnable_type`.
	ReturnableID field.Field[string] `query:"returnable_id"`
	// One of `payment_order`, `paper_item`, `reversal`, or `incoming_payment_detail`.
	// Must be accompanied by `returnable_id`.
	ReturnableType field.Field[ReturnListParamsReturnableType] `query:"returnable_type"`
}

// URLQuery serializes ReturnListParams into a url.Values of the query parameters
// associated with this value
func (r ReturnListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ReturnListParamsReturnableType string

const (
	ReturnListParamsReturnableTypeIncomingPaymentDetail ReturnListParamsReturnableType = "incoming_payment_detail"
	ReturnListParamsReturnableTypePaperItem             ReturnListParamsReturnableType = "paper_item"
	ReturnListParamsReturnableTypePaymentOrder          ReturnListParamsReturnableType = "payment_order"
	ReturnListParamsReturnableTypeReturn                ReturnListParamsReturnableType = "return"
	ReturnListParamsReturnableTypeReversal              ReturnListParamsReturnableType = "reversal"
)