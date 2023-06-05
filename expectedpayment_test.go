package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestExpectedPaymentNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.New(context.TODO(), moderntreasury.ExpectedPaymentNewParams{
		AmountLowerBound:      moderntreasury.F(int64(0)),
		AmountUpperBound:      moderntreasury.F(int64(0)),
		Direction:             moderntreasury.F(moderntreasury.ExpectedPaymentNewParamsDirectionCredit),
		InternalAccountID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CounterpartyID:        moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Currency:              moderntreasury.F(shared.CurrencyAed),
		DateLowerBound:        moderntreasury.F(time.Now()),
		DateUpperBound:        moderntreasury.F(time.Now()),
		Description:           moderntreasury.F("string"),
		LineItems:             moderntreasury.F([]moderntreasury.ExpectedPaymentNewParamsLineItems{{Amount: moderntreasury.F(int64(0)), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), Description: moderntreasury.F("string"), AccountingCategoryID: moderntreasury.F("string")}, {Amount: moderntreasury.F(int64(0)), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), Description: moderntreasury.F("string"), AccountingCategoryID: moderntreasury.F("string")}, {Amount: moderntreasury.F(int64(0)), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), Description: moderntreasury.F("string"), AccountingCategoryID: moderntreasury.F("string")}}),
		Metadata:              moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}),
		RemittanceInformation: moderntreasury.F("string"),
		StatementDescriptor:   moderntreasury.F("string"),
		Type:                  moderntreasury.F(moderntreasury.ExpectedPaymentTypeACH),
		IdempotencyKey:        moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentGet(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.Get(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.Update(
		context.TODO(),
		"string",
		moderntreasury.ExpectedPaymentUpdateParams{
			AmountLowerBound:      moderntreasury.F(int64(0)),
			AmountUpperBound:      moderntreasury.F(int64(0)),
			CounterpartyID:        moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:              moderntreasury.F(shared.CurrencyAed),
			DateLowerBound:        moderntreasury.F(time.Now()),
			DateUpperBound:        moderntreasury.F(time.Now()),
			Description:           moderntreasury.F("string"),
			Direction:             moderntreasury.F(moderntreasury.ExpectedPaymentUpdateParamsDirectionCredit),
			InternalAccountID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata:              moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}),
			RemittanceInformation: moderntreasury.F("string"),
			StatementDescriptor:   moderntreasury.F("string"),
			Type:                  moderntreasury.F(moderntreasury.ExpectedPaymentTypeACH),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.List(context.TODO(), moderntreasury.ExpectedPaymentListParams{
		AfterCursor:         moderntreasury.F("string"),
		CounterpartyID:      moderntreasury.F("string"),
		CreatedAtLowerBound: moderntreasury.F(time.Now()),
		CreatedAtUpperBound: moderntreasury.F(time.Now()),
		Direction:           moderntreasury.F(moderntreasury.ExpectedPaymentListParamsDirectionCredit),
		InternalAccountID:   moderntreasury.F("string"),
		Metadata:            moderntreasury.F(map[string]string{"foo": "string"}),
		PerPage:             moderntreasury.F(int64(0)),
		Status:              moderntreasury.F(moderntreasury.ExpectedPaymentListParamsStatusArchived),
		Type:                moderntreasury.F(moderntreasury.ExpectedPaymentListParamsTypeACH),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentDelete(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.Delete(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
