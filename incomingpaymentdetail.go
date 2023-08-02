// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// IncomingPaymentDetailService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewIncomingPaymentDetailService] method instead.
type IncomingPaymentDetailService struct {
	Options []option.RequestOption
}

// NewIncomingPaymentDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIncomingPaymentDetailService(opts ...option.RequestOption) (r *IncomingPaymentDetailService) {
	r = &IncomingPaymentDetailService{}
	r.Options = opts
	return
}

// Get an existing Incoming Payment Detail.
func (r *IncomingPaymentDetailService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *IncomingPaymentDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/incoming_payment_details/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing Incoming Payment Detail.
func (r *IncomingPaymentDetailService) Update(ctx context.Context, id string, body IncomingPaymentDetailUpdateParams, opts ...option.RequestOption) (res *IncomingPaymentDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/incoming_payment_details/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of Incoming Payment Details.
func (r *IncomingPaymentDetailService) List(ctx context.Context, query IncomingPaymentDetailListParams, opts ...option.RequestOption) (res *shared.Page[IncomingPaymentDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/incoming_payment_details"
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

// Get a list of Incoming Payment Details.
func (r *IncomingPaymentDetailService) ListAutoPaging(ctx context.Context, query IncomingPaymentDetailListParams, opts ...option.RequestOption) *shared.PageAutoPager[IncomingPaymentDetail] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Simulate Incoming Payment Detail
func (r *IncomingPaymentDetailService) NewAsync(ctx context.Context, params IncomingPaymentDetailNewAsyncParams, opts ...option.RequestOption) (res *shared.AsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/simulations/incoming_payment_details/create_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type IncomingPaymentDetail struct {
	ID string `json:"id,required" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// The date on which the corresponding transaction will occur.
	AsOfDate  time.Time `json:"as_of_date,required" format:"date"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the incoming payment detail.
	Currency shared.Currency `json:"currency,required,nullable"`
	// The raw data from the payment pre-notification file that we get from the bank.
	Data map[string]interface{} `json:"data,required"`
	// One of `credit` or `debit`.
	Direction IncomingPaymentDetailDirection `json:"direction,required"`
	// The ID of the Internal Account for the incoming payment detail. This is always
	// present.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// The ID of the ledger transaction linked to the incoming payment detail or
	// `null`.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The last 4 digits of the originating account_number for the incoming payment
	// detail.
	OriginatingAccountNumberSafe string `json:"originating_account_number_safe,required,nullable"`
	// The type of the originating account number for the incoming payment detail.
	OriginatingAccountNumberType IncomingPaymentDetailOriginatingAccountNumberType `json:"originating_account_number_type,required,nullable"`
	// The routing number of the originating account for the incoming payment detail.
	OriginatingRoutingNumber string `json:"originating_routing_number,required,nullable"`
	// The type of the originating routing number for the incoming payment detail.
	OriginatingRoutingNumberType IncomingPaymentDetailOriginatingRoutingNumberType `json:"originating_routing_number_type,required,nullable"`
	// The current status of the incoming payment order. One of `pending`, `completed`,
	// or `returned`.
	Status IncomingPaymentDetailStatus `json:"status,required"`
	// The ID of the reconciled Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The ID of the reconciled Transaction Line Item or `null`.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// One of: `ach`, `book`, `check`, `eft`, `interac`, `rtp`, `sepa`, `signet`, or
	// `wire`.
	Type      IncomingPaymentDetailType `json:"type,required"`
	UpdatedAt time.Time                 `json:"updated_at,required" format:"date-time"`
	// The identifier of the vendor bank.
	VendorID string `json:"vendor_id,required,nullable" format:"uuid"`
	// If the incoming payment detail is in a virtual account, the serialized virtual
	// account object.
	VirtualAccount VirtualAccount `json:"virtual_account,required,nullable"`
	// If the incoming payment detail is in a virtual account, the ID of the Virtual
	// Account.
	VirtualAccountID string `json:"virtual_account_id,required,nullable" format:"uuid"`
	// The account number of the originating account for the incoming payment detail.
	OriginatingAccountNumber string `json:"originating_account_number,nullable"`
	JSON                     incomingPaymentDetailJSON
}

// incomingPaymentDetailJSON contains the JSON metadata for the struct
// [IncomingPaymentDetail]
type incomingPaymentDetailJSON struct {
	ID                           apijson.Field
	Amount                       apijson.Field
	AsOfDate                     apijson.Field
	CreatedAt                    apijson.Field
	Currency                     apijson.Field
	Data                         apijson.Field
	Direction                    apijson.Field
	InternalAccountID            apijson.Field
	LedgerTransactionID          apijson.Field
	LiveMode                     apijson.Field
	Metadata                     apijson.Field
	Object                       apijson.Field
	OriginatingAccountNumberSafe apijson.Field
	OriginatingAccountNumberType apijson.Field
	OriginatingRoutingNumber     apijson.Field
	OriginatingRoutingNumberType apijson.Field
	Status                       apijson.Field
	TransactionID                apijson.Field
	TransactionLineItemID        apijson.Field
	Type                         apijson.Field
	UpdatedAt                    apijson.Field
	VendorID                     apijson.Field
	VirtualAccount               apijson.Field
	VirtualAccountID             apijson.Field
	OriginatingAccountNumber     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *IncomingPaymentDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// One of `credit` or `debit`.
type IncomingPaymentDetailDirection string

const (
	IncomingPaymentDetailDirectionCredit IncomingPaymentDetailDirection = "credit"
	IncomingPaymentDetailDirectionDebit  IncomingPaymentDetailDirection = "debit"
)

// The type of the originating account number for the incoming payment detail.
type IncomingPaymentDetailOriginatingAccountNumberType string

const (
	IncomingPaymentDetailOriginatingAccountNumberTypeClabe         IncomingPaymentDetailOriginatingAccountNumberType = "clabe"
	IncomingPaymentDetailOriginatingAccountNumberTypeIban          IncomingPaymentDetailOriginatingAccountNumberType = "iban"
	IncomingPaymentDetailOriginatingAccountNumberTypeOther         IncomingPaymentDetailOriginatingAccountNumberType = "other"
	IncomingPaymentDetailOriginatingAccountNumberTypePan           IncomingPaymentDetailOriginatingAccountNumberType = "pan"
	IncomingPaymentDetailOriginatingAccountNumberTypeWalletAddress IncomingPaymentDetailOriginatingAccountNumberType = "wallet_address"
)

// The type of the originating routing number for the incoming payment detail.
type IncomingPaymentDetailOriginatingRoutingNumberType string

const (
	IncomingPaymentDetailOriginatingRoutingNumberTypeAba          IncomingPaymentDetailOriginatingRoutingNumberType = "aba"
	IncomingPaymentDetailOriginatingRoutingNumberTypeAuBsb        IncomingPaymentDetailOriginatingRoutingNumberType = "au_bsb"
	IncomingPaymentDetailOriginatingRoutingNumberTypeBrCodigo     IncomingPaymentDetailOriginatingRoutingNumberType = "br_codigo"
	IncomingPaymentDetailOriginatingRoutingNumberTypeCaCpa        IncomingPaymentDetailOriginatingRoutingNumberType = "ca_cpa"
	IncomingPaymentDetailOriginatingRoutingNumberTypeChips        IncomingPaymentDetailOriginatingRoutingNumberType = "chips"
	IncomingPaymentDetailOriginatingRoutingNumberTypeCnaps        IncomingPaymentDetailOriginatingRoutingNumberType = "cnaps"
	IncomingPaymentDetailOriginatingRoutingNumberTypeGBSortCode   IncomingPaymentDetailOriginatingRoutingNumberType = "gb_sort_code"
	IncomingPaymentDetailOriginatingRoutingNumberTypeInIfsc       IncomingPaymentDetailOriginatingRoutingNumberType = "in_ifsc"
	IncomingPaymentDetailOriginatingRoutingNumberTypeMyBranchCode IncomingPaymentDetailOriginatingRoutingNumberType = "my_branch_code"
	IncomingPaymentDetailOriginatingRoutingNumberTypeSwift        IncomingPaymentDetailOriginatingRoutingNumberType = "swift"
)

// The current status of the incoming payment order. One of `pending`, `completed`,
// or `returned`.
type IncomingPaymentDetailStatus string

const (
	IncomingPaymentDetailStatusCompleted IncomingPaymentDetailStatus = "completed"
	IncomingPaymentDetailStatusPending   IncomingPaymentDetailStatus = "pending"
	IncomingPaymentDetailStatusReturned  IncomingPaymentDetailStatus = "returned"
)

// One of: `ach`, `book`, `check`, `eft`, `interac`, `rtp`, `sepa`, `signet`, or
// `wire`.
type IncomingPaymentDetailType string

const (
	IncomingPaymentDetailTypeACH     IncomingPaymentDetailType = "ach"
	IncomingPaymentDetailTypeBook    IncomingPaymentDetailType = "book"
	IncomingPaymentDetailTypeCheck   IncomingPaymentDetailType = "check"
	IncomingPaymentDetailTypeEft     IncomingPaymentDetailType = "eft"
	IncomingPaymentDetailTypeInterac IncomingPaymentDetailType = "interac"
	IncomingPaymentDetailTypeRtp     IncomingPaymentDetailType = "rtp"
	IncomingPaymentDetailTypeSepa    IncomingPaymentDetailType = "sepa"
	IncomingPaymentDetailTypeSignet  IncomingPaymentDetailType = "signet"
	IncomingPaymentDetailTypeWire    IncomingPaymentDetailType = "wire"
)

type IncomingPaymentDetailUpdateParams struct {
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r IncomingPaymentDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IncomingPaymentDetailListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Filters incoming payment details with an as_of_date starting on or before the
	// specified date (YYYY-MM-DD).
	AsOfDateEnd param.Field[time.Time] `query:"as_of_date_end" format:"date"`
	// Filters incoming payment details with an as_of_date starting on or after the
	// specified date (YYYY-MM-DD).
	AsOfDateStart param.Field[time.Time] `query:"as_of_date_start" format:"date"`
	// One of `credit` or `debit`.
	Direction param.Field[IncomingPaymentDetailListParamsDirection] `query:"direction"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
	// The current status of the incoming payment order. One of `pending`, `completed`,
	// or `returned`.
	Status param.Field[IncomingPaymentDetailListParamsStatus] `query:"status"`
	// One of: `ach`, `book`, `check`, `eft`, `interac`, `rtp`, `sepa`, `signet`, or
	// `wire`.
	Type param.Field[IncomingPaymentDetailListParamsType] `query:"type"`
	// If the incoming payment detail is in a virtual account, the ID of the Virtual
	// Account.
	VirtualAccountID param.Field[string] `query:"virtual_account_id"`
}

// URLQuery serializes [IncomingPaymentDetailListParams]'s query parameters as
// `url.Values`.
func (r IncomingPaymentDetailListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// One of `credit` or `debit`.
type IncomingPaymentDetailListParamsDirection string

const (
	IncomingPaymentDetailListParamsDirectionCredit IncomingPaymentDetailListParamsDirection = "credit"
	IncomingPaymentDetailListParamsDirectionDebit  IncomingPaymentDetailListParamsDirection = "debit"
)

// The current status of the incoming payment order. One of `pending`, `completed`,
// or `returned`.
type IncomingPaymentDetailListParamsStatus string

const (
	IncomingPaymentDetailListParamsStatusCompleted IncomingPaymentDetailListParamsStatus = "completed"
	IncomingPaymentDetailListParamsStatusPending   IncomingPaymentDetailListParamsStatus = "pending"
	IncomingPaymentDetailListParamsStatusReturned  IncomingPaymentDetailListParamsStatus = "returned"
)

// One of: `ach`, `book`, `check`, `eft`, `interac`, `rtp`, `sepa`, `signet`, or
// `wire`.
type IncomingPaymentDetailListParamsType string

const (
	IncomingPaymentDetailListParamsTypeACH     IncomingPaymentDetailListParamsType = "ach"
	IncomingPaymentDetailListParamsTypeBook    IncomingPaymentDetailListParamsType = "book"
	IncomingPaymentDetailListParamsTypeCheck   IncomingPaymentDetailListParamsType = "check"
	IncomingPaymentDetailListParamsTypeEft     IncomingPaymentDetailListParamsType = "eft"
	IncomingPaymentDetailListParamsTypeInterac IncomingPaymentDetailListParamsType = "interac"
	IncomingPaymentDetailListParamsTypeRtp     IncomingPaymentDetailListParamsType = "rtp"
	IncomingPaymentDetailListParamsTypeSepa    IncomingPaymentDetailListParamsType = "sepa"
	IncomingPaymentDetailListParamsTypeSignet  IncomingPaymentDetailListParamsType = "signet"
	IncomingPaymentDetailListParamsTypeWire    IncomingPaymentDetailListParamsType = "wire"
)

type IncomingPaymentDetailNewAsyncParams struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount"`
	// Defaults to today.
	AsOfDate param.Field[time.Time] `json:"as_of_date" format:"date"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// Defaults to a random description.
	Description param.Field[string] `json:"description"`
	// One of `credit`, `debit`.
	Direction param.Field[IncomingPaymentDetailNewAsyncParamsDirection] `json:"direction"`
	// The ID of one of your internal accounts.
	InternalAccountID param.Field[string] `json:"internal_account_id" format:"uuid"`
	// One of `ach`, `wire`, `check`.
	Type param.Field[IncomingPaymentDetailNewAsyncParamsType] `json:"type"`
	// An optional parameter to associate the incoming payment detail to a virtual
	// account.
	VirtualAccountID param.Field[string] `json:"virtual_account_id" format:"uuid"`
	IdempotencyKey   param.Field[string] `header:"Idempotency-Key"`
}

func (r IncomingPaymentDetailNewAsyncParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of `credit`, `debit`.
type IncomingPaymentDetailNewAsyncParamsDirection string

const (
	IncomingPaymentDetailNewAsyncParamsDirectionCredit IncomingPaymentDetailNewAsyncParamsDirection = "credit"
	IncomingPaymentDetailNewAsyncParamsDirectionDebit  IncomingPaymentDetailNewAsyncParamsDirection = "debit"
)

// One of `ach`, `wire`, `check`.
type IncomingPaymentDetailNewAsyncParamsType string

const (
	IncomingPaymentDetailNewAsyncParamsTypeACH     IncomingPaymentDetailNewAsyncParamsType = "ach"
	IncomingPaymentDetailNewAsyncParamsTypeBook    IncomingPaymentDetailNewAsyncParamsType = "book"
	IncomingPaymentDetailNewAsyncParamsTypeCheck   IncomingPaymentDetailNewAsyncParamsType = "check"
	IncomingPaymentDetailNewAsyncParamsTypeEft     IncomingPaymentDetailNewAsyncParamsType = "eft"
	IncomingPaymentDetailNewAsyncParamsTypeInterac IncomingPaymentDetailNewAsyncParamsType = "interac"
	IncomingPaymentDetailNewAsyncParamsTypeRtp     IncomingPaymentDetailNewAsyncParamsType = "rtp"
	IncomingPaymentDetailNewAsyncParamsTypeSepa    IncomingPaymentDetailNewAsyncParamsType = "sepa"
	IncomingPaymentDetailNewAsyncParamsTypeSignet  IncomingPaymentDetailNewAsyncParamsType = "signet"
	IncomingPaymentDetailNewAsyncParamsTypeWire    IncomingPaymentDetailNewAsyncParamsType = "wire"
)
