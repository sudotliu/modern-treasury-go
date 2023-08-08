// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apiform"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// DocumentService contains methods and other services that help with interacting
// with the Modern Treasury API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Create a document.
func (r *DocumentService) New(ctx context.Context, params DocumentNewParams, opts ...option.RequestOption) (res *Document, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get an existing document.
func (r *DocumentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of documents.
func (r *DocumentService) List(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) (res *shared.Page[Document], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/documents"
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

// Get a list of documents.
func (r *DocumentService) ListAutoPaging(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) *shared.PageAutoPager[Document] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Document struct {
	ID              string                   `json:"id,required" format:"uuid"`
	CreatedAt       time.Time                `json:"created_at,required" format:"date-time"`
	DiscardedAt     time.Time                `json:"discarded_at,required,nullable" format:"date-time"`
	DocumentDetails []DocumentDocumentDetail `json:"document_details,required"`
	// A category given to the document, can be `null`.
	DocumentType string `json:"document_type,required,nullable"`
	// The unique identifier for the associated object.
	DocumentableID string `json:"documentable_id,required" format:"uuid"`
	// The type of the associated object. Currently can be one of `payment_order`,
	// `transaction`, `paper_item`, `expected_payment`, `counterparty`, `organization`,
	// `case`, `internal_account`, `decision`, or `external_account`.
	DocumentableType DocumentDocumentableType `json:"documentable_type,required"`
	File             DocumentFile             `json:"file,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The source of the document. Can be `vendor`, `customer`, or `modern_treasury`.
	Source    string    `json:"source,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	JSON      documentJSON
}

// documentJSON contains the JSON metadata for the struct [Document]
type documentJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	DiscardedAt      apijson.Field
	DocumentDetails  apijson.Field
	DocumentType     apijson.Field
	DocumentableID   apijson.Field
	DocumentableType apijson.Field
	File             apijson.Field
	LiveMode         apijson.Field
	Object           apijson.Field
	Source           apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentDocumentDetail struct {
	ID                     string    `json:"id,required" format:"uuid"`
	CreatedAt              time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt            time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	DocumentIdentifier     string    `json:"document_identifier,required"`
	DocumentIdentifierType string    `json:"document_identifier_type,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	Object    string    `json:"object,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	JSON      documentDocumentDetailJSON
}

// documentDocumentDetailJSON contains the JSON metadata for the struct
// [DocumentDocumentDetail]
type documentDocumentDetailJSON struct {
	ID                     apijson.Field
	CreatedAt              apijson.Field
	DiscardedAt            apijson.Field
	DocumentIdentifier     apijson.Field
	DocumentIdentifierType apijson.Field
	LiveMode               apijson.Field
	Object                 apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DocumentDocumentDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the associated object. Currently can be one of `payment_order`,
// `transaction`, `paper_item`, `expected_payment`, `counterparty`, `organization`,
// `case`, `internal_account`, `decision`, or `external_account`.
type DocumentDocumentableType string

const (
	DocumentDocumentableTypeCase            DocumentDocumentableType = "case"
	DocumentDocumentableTypeCounterparty    DocumentDocumentableType = "counterparty"
	DocumentDocumentableTypeExpectedPayment DocumentDocumentableType = "expected_payment"
	DocumentDocumentableTypeExternalAccount DocumentDocumentableType = "external_account"
	DocumentDocumentableTypeInternalAccount DocumentDocumentableType = "internal_account"
	DocumentDocumentableTypeOrganization    DocumentDocumentableType = "organization"
	DocumentDocumentableTypePaperItem       DocumentDocumentableType = "paper_item"
	DocumentDocumentableTypePaymentOrder    DocumentDocumentableType = "payment_order"
	DocumentDocumentableTypeTransaction     DocumentDocumentableType = "transaction"
	DocumentDocumentableTypeDecision        DocumentDocumentableType = "decision"
	DocumentDocumentableTypeConnection      DocumentDocumentableType = "connection"
)

type DocumentFile struct {
	// The MIME content type of the document.
	ContentType string `json:"content_type"`
	// The original filename of the document.
	Filename string `json:"filename"`
	// The size of the document in bytes.
	Size int64 `json:"size"`
	JSON documentFileJSON
}

// documentFileJSON contains the JSON metadata for the struct [DocumentFile]
type documentFileJSON struct {
	ContentType apijson.Field
	Filename    apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DocumentFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentNewParams struct {
	// The unique identifier for the associated object.
	DocumentableID   param.Field[string]                            `json:"documentable_id,required"`
	DocumentableType param.Field[DocumentNewParamsDocumentableType] `json:"documentable_type,required"`
	File             param.Field[io.Reader]                         `json:"file,required" format:"binary"`
	// A category given to the document, can be `null`.
	DocumentType   param.Field[string] `json:"document_type"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r DocumentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type DocumentNewParamsDocumentableType string

const (
	DocumentNewParamsDocumentableTypeCases            DocumentNewParamsDocumentableType = "cases"
	DocumentNewParamsDocumentableTypeCounterparties   DocumentNewParamsDocumentableType = "counterparties"
	DocumentNewParamsDocumentableTypeExpectedPayments DocumentNewParamsDocumentableType = "expected_payments"
	DocumentNewParamsDocumentableTypeExternalAccounts DocumentNewParamsDocumentableType = "external_accounts"
	DocumentNewParamsDocumentableTypeInternalAccounts DocumentNewParamsDocumentableType = "internal_accounts"
	DocumentNewParamsDocumentableTypeOrganizations    DocumentNewParamsDocumentableType = "organizations"
	DocumentNewParamsDocumentableTypePaperItems       DocumentNewParamsDocumentableType = "paper_items"
	DocumentNewParamsDocumentableTypePaymentOrders    DocumentNewParamsDocumentableType = "payment_orders"
	DocumentNewParamsDocumentableTypeTransactions     DocumentNewParamsDocumentableType = "transactions"
	DocumentNewParamsDocumentableTypeDecisions        DocumentNewParamsDocumentableType = "decisions"
	DocumentNewParamsDocumentableTypeConnections      DocumentNewParamsDocumentableType = "connections"
)

type DocumentListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// The unique identifier for the associated object.
	DocumentableID param.Field[string] `query:"documentable_id"`
	// The type of the associated object. Currently can be one of `payment_order`,
	// `transaction`, `paper_item`, `expected_payment`, `counterparty`, `organization`,
	// `case`, `internal_account`, `decision`, or `external_account`.
	DocumentableType param.Field[DocumentListParamsDocumentableType] `query:"documentable_type"`
	PerPage          param.Field[int64]                              `query:"per_page"`
}

// URLQuery serializes [DocumentListParams]'s query parameters as `url.Values`.
func (r DocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of the associated object. Currently can be one of `payment_order`,
// `transaction`, `paper_item`, `expected_payment`, `counterparty`, `organization`,
// `case`, `internal_account`, `decision`, or `external_account`.
type DocumentListParamsDocumentableType string

const (
	DocumentListParamsDocumentableTypeCases            DocumentListParamsDocumentableType = "cases"
	DocumentListParamsDocumentableTypeCounterparties   DocumentListParamsDocumentableType = "counterparties"
	DocumentListParamsDocumentableTypeExpectedPayments DocumentListParamsDocumentableType = "expected_payments"
	DocumentListParamsDocumentableTypeExternalAccounts DocumentListParamsDocumentableType = "external_accounts"
	DocumentListParamsDocumentableTypeInternalAccounts DocumentListParamsDocumentableType = "internal_accounts"
	DocumentListParamsDocumentableTypeOrganizations    DocumentListParamsDocumentableType = "organizations"
	DocumentListParamsDocumentableTypePaperItems       DocumentListParamsDocumentableType = "paper_items"
	DocumentListParamsDocumentableTypePaymentOrders    DocumentListParamsDocumentableType = "payment_orders"
	DocumentListParamsDocumentableTypeTransactions     DocumentListParamsDocumentableType = "transactions"
	DocumentListParamsDocumentableTypeDecisions        DocumentListParamsDocumentableType = "decisions"
	DocumentListParamsDocumentableTypeConnections      DocumentListParamsDocumentableType = "connections"
)
