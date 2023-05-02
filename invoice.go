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

type InvoiceService struct {
	Options   []option.RequestOption
	LineItems *InvoiceLineItemService
}

func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	r.LineItems = NewInvoiceLineItemService(opts...)
	return
}

// create invoice
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get invoice
func (r *InvoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update invoice
func (r *InvoiceService) Update(ctx context.Context, id string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list invoices
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *shared.Page[Invoice], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/invoices"
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

// list invoices
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *shared.PageAutoPager[Invoice] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Invoice struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails []InvoiceContactDetails `json:"contact_details,required"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID string `json:"counterparty_id,required"`
	// The counterparty's billing address.
	CounterpartyBillingAddress InvoiceCounterpartyBillingAddress `json:"counterparty_billing_address,required,nullable"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress InvoiceCounterpartyShippingAddress `json:"counterparty_shipping_address,required,nullable"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency shared.Currency `json:"currency,required,nullable"`
	// A free-form description of the invoice.
	Description string `json:"description,required"`
	// A future date by when the invoice needs to be paid.
	DueDate time.Time `json:"due_date,required" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress InvoiceInvoicerAddress `json:"invoicer_address,required,nullable"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID string `json:"originating_account_id,required"`
	// The URL of the hosted web UI where the invoice can be viewed.
	HostedURL string `json:"hosted_url,required"`
	// A unique record number assigned to each invoice that is issued.
	Number string `json:"number,required"`
	// The payment orders created for paying the invoice through the invoice payment
	// UI.
	PaymentOrders []PaymentOrder `json:"payment_orders,required"`
	// The URL where the invoice PDF can be downloaded.
	PdfURL string `json:"pdf_url,required,nullable"`
	// The status of the invoice.
	Status InvoiceStatus `json:"status,required"`
	// Total amount due in specified currency's smallest unit, e.g., $10 USD would be
	// represented as 1000.
	TotalAmount int64 `json:"total_amount,required"`
	JSON        InvoiceJSON
}

type InvoiceJSON struct {
	ID                          apijson.Metadata
	Object                      apijson.Metadata
	LiveMode                    apijson.Metadata
	CreatedAt                   apijson.Metadata
	UpdatedAt                   apijson.Metadata
	ContactDetails              apijson.Metadata
	CounterpartyID              apijson.Metadata
	CounterpartyBillingAddress  apijson.Metadata
	CounterpartyShippingAddress apijson.Metadata
	Currency                    apijson.Metadata
	Description                 apijson.Metadata
	DueDate                     apijson.Metadata
	InvoicerAddress             apijson.Metadata
	OriginatingAccountID        apijson.Metadata
	HostedURL                   apijson.Metadata
	Number                      apijson.Metadata
	PaymentOrders               apijson.Metadata
	PdfURL                      apijson.Metadata
	Status                      apijson.Metadata
	TotalAmount                 apijson.Metadata
	raw                         string
	Extras                      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Invoice using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceContactDetails struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              bool                                       `json:"live_mode,required"`
	CreatedAt             time.Time                                  `json:"created_at,required" format:"date-time"`
	UpdatedAt             time.Time                                  `json:"updated_at,required" format:"date-time"`
	DiscardedAt           time.Time                                  `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     string                                     `json:"contact_identifier,required"`
	ContactIdentifierType InvoiceContactDetailsContactIdentifierType `json:"contact_identifier_type,required"`
	JSON                  InvoiceContactDetailsJSON
}

type InvoiceContactDetailsJSON struct {
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

// UnmarshalJSON deserializes the provided bytes into InvoiceContactDetails using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InvoiceContactDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceContactDetailsContactIdentifierType string

const (
	InvoiceContactDetailsContactIdentifierTypeEmail       InvoiceContactDetailsContactIdentifierType = "email"
	InvoiceContactDetailsContactIdentifierTypePhoneNumber InvoiceContactDetailsContactIdentifierType = "phone_number"
	InvoiceContactDetailsContactIdentifierTypeWebsite     InvoiceContactDetailsContactIdentifierType = "website"
)

type InvoiceCounterpartyBillingAddress struct {
	Line1 string `json:"line1,required"`
	Line2 string `json:"line2"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// Region or State.
	Region string `json:"region,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	JSON    InvoiceCounterpartyBillingAddressJSON
}

type InvoiceCounterpartyBillingAddressJSON struct {
	Line1      apijson.Metadata
	Line2      apijson.Metadata
	Locality   apijson.Metadata
	Region     apijson.Metadata
	PostalCode apijson.Metadata
	Country    apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InvoiceCounterpartyBillingAddress using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *InvoiceCounterpartyBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCounterpartyShippingAddress struct {
	Line1 string `json:"line1,required"`
	Line2 string `json:"line2"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// Region or State.
	Region string `json:"region,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	JSON    InvoiceCounterpartyShippingAddressJSON
}

type InvoiceCounterpartyShippingAddressJSON struct {
	Line1      apijson.Metadata
	Line2      apijson.Metadata
	Locality   apijson.Metadata
	Region     apijson.Metadata
	PostalCode apijson.Metadata
	Country    apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InvoiceCounterpartyShippingAddress using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *InvoiceCounterpartyShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceInvoicerAddress struct {
	Line1 string `json:"line1,required"`
	Line2 string `json:"line2"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// Region or State.
	Region string `json:"region,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	JSON    InvoiceInvoicerAddressJSON
}

type InvoiceInvoicerAddressJSON struct {
	Line1      apijson.Metadata
	Line2      apijson.Metadata
	Locality   apijson.Metadata
	Region     apijson.Metadata
	PostalCode apijson.Metadata
	Country    apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InvoiceInvoicerAddress using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InvoiceInvoicerAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceStatus string

const (
	InvoiceStatusDraft          InvoiceStatus = "draft"
	InvoiceStatusPaid           InvoiceStatus = "paid"
	InvoiceStatusPaymentPending InvoiceStatus = "payment_pending"
	InvoiceStatusUnpaid         InvoiceStatus = "unpaid"
	InvoiceStatusVoided         InvoiceStatus = "voided"
)

type InvoiceNewParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails field.Field[[]InvoiceNewParamsContactDetails] `json:"contact_details"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID field.Field[string] `json:"counterparty_id,required"`
	// The counterparty's billing address.
	CounterpartyBillingAddress field.Field[InvoiceNewParamsCounterpartyBillingAddress] `json:"counterparty_billing_address,nullable"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress field.Field[InvoiceNewParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address,nullable"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency field.Field[shared.Currency] `json:"currency,nullable"`
	// A free-form description of the invoice.
	Description field.Field[string] `json:"description"`
	// A future date by when the invoice needs to be paid.
	DueDate field.Field[time.Time] `json:"due_date,required" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress field.Field[InvoiceNewParamsInvoicerAddress] `json:"invoicer_address,nullable"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID field.Field[string] `json:"originating_account_id,required"`
}

// MarshalJSON serializes InvoiceNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsContactDetails struct {
	ID     field.Field[string] `json:"id,required" format:"uuid"`
	Object field.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              field.Field[bool]                                                `json:"live_mode,required"`
	CreatedAt             field.Field[time.Time]                                           `json:"created_at,required" format:"date-time"`
	UpdatedAt             field.Field[time.Time]                                           `json:"updated_at,required" format:"date-time"`
	DiscardedAt           field.Field[time.Time]                                           `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     field.Field[string]                                              `json:"contact_identifier,required"`
	ContactIdentifierType field.Field[InvoiceNewParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
}

type InvoiceNewParamsContactDetailsContactIdentifierType string

const (
	InvoiceNewParamsContactDetailsContactIdentifierTypeEmail       InvoiceNewParamsContactDetailsContactIdentifierType = "email"
	InvoiceNewParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceNewParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceNewParamsContactDetailsContactIdentifierTypeWebsite     InvoiceNewParamsContactDetailsContactIdentifierType = "website"
)

type InvoiceNewParamsCounterpartyBillingAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceNewParamsCounterpartyShippingAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceNewParamsInvoicerAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceUpdateParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails field.Field[[]InvoiceUpdateParamsContactDetails] `json:"contact_details"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID field.Field[string] `json:"counterparty_id"`
	// The counterparty's billing address.
	CounterpartyBillingAddress field.Field[InvoiceUpdateParamsCounterpartyBillingAddress] `json:"counterparty_billing_address,nullable"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress field.Field[InvoiceUpdateParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address,nullable"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency field.Field[shared.Currency] `json:"currency,nullable"`
	// A free-form description of the invoice.
	Description field.Field[string] `json:"description"`
	// A future date by when the invoice needs to be paid.
	DueDate field.Field[time.Time] `json:"due_date" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress field.Field[InvoiceUpdateParamsInvoicerAddress] `json:"invoicer_address,nullable"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID field.Field[string] `json:"originating_account_id"`
	// When opening an invoice, whether to show the embedded payment UI with the
	// invoice. Default true.
	IncludePaymentUi field.Field[bool] `json:"include_payment_ui"`
	// Invoice status must be updated in a `PATCH` request that does not modify any
	// other invoice attributes. Valid state transitions are `draft` to `unpaid` and
	// `draft` or `unpaid` to `voided`.
	Status field.Field[string] `json:"status"`
}

// MarshalJSON serializes InvoiceUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceUpdateParamsContactDetails struct {
	ID     field.Field[string] `json:"id,required" format:"uuid"`
	Object field.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              field.Field[bool]                                                   `json:"live_mode,required"`
	CreatedAt             field.Field[time.Time]                                              `json:"created_at,required" format:"date-time"`
	UpdatedAt             field.Field[time.Time]                                              `json:"updated_at,required" format:"date-time"`
	DiscardedAt           field.Field[time.Time]                                              `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     field.Field[string]                                                 `json:"contact_identifier,required"`
	ContactIdentifierType field.Field[InvoiceUpdateParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
}

type InvoiceUpdateParamsContactDetailsContactIdentifierType string

const (
	InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail       InvoiceUpdateParamsContactDetailsContactIdentifierType = "email"
	InvoiceUpdateParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceUpdateParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceUpdateParamsContactDetailsContactIdentifierTypeWebsite     InvoiceUpdateParamsContactDetailsContactIdentifierType = "website"
)

type InvoiceUpdateParamsCounterpartyBillingAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceUpdateParamsCounterpartyShippingAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceUpdateParamsInvoicerAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes InvoiceListParams into a url.Values of the query parameters
// associated with this value
func (r InvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}