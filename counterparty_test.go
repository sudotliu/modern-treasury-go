package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestCounterpartyNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.New(context.TODO(), moderntreasury.CounterpartyNewParams{Name: moderntreasury.F("string"), Accounts: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccounts{{AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash), PartyType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness), PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Name: moderntreasury.F("string"), AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetails{{AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}}), RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetails{{RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}}), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), PartyName: moderntreasury.F("string"), PartyIdentifier: moderntreasury.F("string"), PlaidProcessorToken: moderntreasury.F("string"), ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetails{{ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}})}, {AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash), PartyType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness), PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Name: moderntreasury.F("string"), AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetails{{AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}}), RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetails{{RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}}), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), PartyName: moderntreasury.F("string"), PartyIdentifier: moderntreasury.F("string"), PlaidProcessorToken: moderntreasury.F("string"), ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetails{{ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}})}, {AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash), PartyType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness), PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Name: moderntreasury.F("string"), AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetails{{AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}}), RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetails{{RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}}), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), PartyName: moderntreasury.F("string"), PartyIdentifier: moderntreasury.F("string"), PlaidProcessorToken: moderntreasury.F("string"), ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetails{{ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}})}}), Email: moderntreasury.F("dev@stainlessapi.com"), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), SendRemittanceAdvice: moderntreasury.F(true), Accounting: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccounting{Type: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountingTypeCustomer)}), LedgerType: moderntreasury.F(moderntreasury.CounterpartyNewParamsLedgerTypeCustomer), TaxpayerIdentifier: moderntreasury.F("string")})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyGet(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.Get(
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

func TestCounterpartyUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.Update(
		context.TODO(),
		"string",
		moderntreasury.CounterpartyUpdateParams{Name: moderntreasury.F("string"), Email: moderntreasury.F("dev@stainlessapi.com"), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), SendRemittanceAdvice: moderntreasury.F(true), TaxpayerIdentifier: moderntreasury.F("string")},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.List(context.TODO(), moderntreasury.CounterpartyListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), Name: moderntreasury.F("string"), Email: moderntreasury.F("dev@stainlessapi.com"), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), CreatedAtLowerBound: moderntreasury.F(time.Now()), CreatedAtUpperBound: moderntreasury.F(time.Now())})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyDelete(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.Counterparties.Delete(
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

func TestCounterpartyCollectAccountWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.CollectAccount(
		context.TODO(),
		"string",
		moderntreasury.CounterpartyCollectAccountParams{Direction: moderntreasury.F(moderntreasury.CounterpartyCollectAccountParamsDirectionCredit), SendEmail: moderntreasury.F(true), Fields: moderntreasury.F([]moderntreasury.CounterpartyCollectAccountParamsFields{moderntreasury.CounterpartyCollectAccountParamsFieldsName, moderntreasury.CounterpartyCollectAccountParamsFieldsName, moderntreasury.CounterpartyCollectAccountParamsFieldsName}), CustomRedirect: moderntreasury.F("https://example.com")},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}