/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package adyen

import (
	"fmt"
	"net/http"

	"github.com/adyen/adyen-go-api-library/v6/src/balanceplatform"
	"github.com/adyen/adyen-go-api-library/v6/src/checkout"
	"github.com/adyen/adyen-go-api-library/v6/src/legalentity"
	"github.com/adyen/adyen-go-api-library/v6/src/management"
	"github.com/adyen/adyen-go-api-library/v6/src/recurring"
	"github.com/adyen/adyen-go-api-library/v6/src/transfers"
	"github.com/adyen/adyen-go-api-library/v6/src/webhook"

	binlookup "github.com/adyen/adyen-go-api-library/v6/src/binlookup"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/adyen/adyen-go-api-library/v6/src/disputes"
	"github.com/adyen/adyen-go-api-library/v6/src/payments"
	"github.com/adyen/adyen-go-api-library/v6/src/payout"
	"github.com/adyen/adyen-go-api-library/v6/src/platformsaccount"
	"github.com/adyen/adyen-go-api-library/v6/src/platformsfund"
	"github.com/adyen/adyen-go-api-library/v6/src/platformshostedonboardingpage"
	"github.com/adyen/adyen-go-api-library/v6/src/platformsnotificationconfiguration"
	"github.com/adyen/adyen-go-api-library/v6/src/posterminalmanagement"
	"github.com/adyen/adyen-go-api-library/v6/src/storedvalue"
)

// Constants used for the client API
const (
	EndpointTest                      = "https://pal-test.adyen.com"
	EndpointLive                      = "https://pal-live.adyen.com"
	EndpointLiveSuffix                = "-pal-live.adyenpayments.com"
	MarketpayEndpointTest             = "https://cal-test.adyen.com/cal/services"
	MarketpayEndpointLive             = "https://cal-live.adyen.com/cal/services"
	CheckoutEndpointTest              = "https://checkout-test.adyen.com/checkout"
	CheckoutEndpointLiveSuffix        = "-checkout-live.adyenpayments.com/checkout"
	BinLookupPalSuffix                = "/pal/servlet/BinLookup/"
	TerminalAPIEndpointTest           = "https://terminal-api-test.adyen.com"
	TerminalAPIEndpointLive           = "https://terminal-api-live.adyen.com"
	DisputesEndpointTest              = "https://ca-test.adyen.com/ca/services/DisputeService"
	DisputesEndpointLive              = "https://ca-live.adyen.com/ca/services/DisputeService"
	BalancePlatformEndpointTest       = "https://balanceplatform-api-test.adyen.com/bcl"
	BalancePlatformEndpointLive       = "https://balanceplatform-api-live.adyen.com/bcl"
	TransfersEndpointTest             = "https://balanceplatform-api-test.adyen.com/btl"
	TransfersEndpointLive             = "https://balanceplatform-api-live.adyen.com/btl"
	ManagementEndpointTest            = "https://management-test.adyen.com"
	ManagementEndpointLive            = "https://management-live.adyen.com"
	LegalEntityEntityTest             = "https://kyc-test.adyen.com/lem"
	LegalEntityEntityLive             = "https://kyc-live.adyen.com/lem"
	PosTerminalManagementEndpointTest = "https://postfmapi-test.adyen.com/postfmapi/terminal"
	PosTerminalManagementEndpointLive = "https://postfmapi-live.adyen.com/postfmapi/terminal"
)

// also update LibVersion in src/common/configuration.go when a version is updated and a major lib version is released
const (
	MarketpayAccountAPIVersion      = "v6"
	MarketpayFundAPIVersion         = "v6"
	MarketpayNotificationAPIVersion = "v6"
	MarketpayHopAPIVersion          = "v6"
	PaymentAPIVersion               = "v68"
	RecurringAPIVersion             = "v68"
	CheckoutAPIVersion              = "v70"
	BinLookupAPIVersion             = "v50"
	EndpointProtocol                = "https://"
	DisputesAPIVersion              = "v30"
	StoredValueAPIVersion           = "v46"
	BalancePlatformAPIVersion       = "v2"
	TransfersAPIVersion             = "v3"
	ManagementAPIVersion            = "v1"
	LegalEntityAPIVersion           = "v3"
	PosTerminalManagementAPIVersion = "v1"
)

// APIClient manages communication with the Adyen Checkout API API v51
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	client *common.Client
	// API Services
	Payments                           *payments.Payments
	payout                             *payout.APIClient
	Recurring                          *recurring.GeneralApi
	BinLookup                          *binlookup.BinLookup
	Notification                       *webhook.NotificationService
	PlatformsAccount                   *platformsaccount.PlatformsAccount
	PlatformsFund                      *platformsfund.PlatformsFund
	PlatformsHostedOnboardingPage      *platformshostedonboardingpage.PlatformsHostedOnboardingPage
	PlatformsNotificationConfiguration *platformsnotificationconfiguration.PlatformsNotificationConfiguration
	PosTerminalManagement              *posterminalmanagement.GeneralApi
	Disputes                           *disputes.Disputes
	StoredValue                        *storedvalue.StoredValue
	BalancePlatform                    *balanceplatform.APIClient
	Transfers                          *transfers.GeneralApi
	Management                         *management.APIClient
	LegalEntity                        *legalentity.APIClient
}

// NewClient creates a new API client. Requires Config object.
//
// @param config              A reference to common.Config object
//
// create a new API client based on provided api key & url prefix for LIVE environment
//
//	client := NewClient(&common.Config{
//	   ApiKey:                "apiKey",
//	   Environment:           common.LiveEnv,
//	   LiveEndpointURLPrefix: "liveEndpointURLPrefix",
//	})
//
//	ApiKey                Defines the api key that can be retrieved by back office
//	Environment           This defines the payment environment live or test
//	LiveEndpointURLPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
//
// create a new API client based on provided api key for TEST environment
//
//	client := NewClient(&common.Config{
//	   ApiKey:                "apiKey",
//	   Environment:           common.TestEnv,
//	})
//
//	ApiKey                Defines the api key that can be retrieved by back office
//	Environment           This defines the payment environment live or test
//
// creates a new API client based on provided credentials &  url prefix for LIVE environment
//
//	client := NewClient(&common.Config{
//	    Username:              "username",
//	    Password:              "password",
//	    ApplicationName:       "applicationName",
//	    Environment:           common.LiveEnv,
//	    LiveEndpointURLPrefix: "liveEndpointURLPrefix",
//	})
//
//	Username              Your merchant account Username
//	Password              Your merchant accont Password
//	Environment           This defines the payment environment live or test
//	ApplicationName       Application name to be used in user agent
//	LiveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
//
// creates a new API client based on provided credentials for TEST environment
//
//	client := NewClient(&common.Config{
//	    Username:              "username",
//	    Password:              "password",
//	    ApplicationName:       "applicationName",
//	    Environment:           common.TestEnv,
//	})
//
//	Username              Your merchant account Username
//	Password              Your merchant accont Password
//	Environment           This defines the payment environment live or test
//	ApplicationName       Application name to be used in user agent
//
// optionally a custom http.Client can be passed via the Config allow for advanced features such as caching.
func NewClient(cfg *common.Config) *APIClient {

	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}
	if cfg.ConnectionTimeoutMillis != 0 {
		cfg.HTTPClient.Timeout = cfg.ConnectionTimeoutMillis
	}
	if cfg.DefaultHeader == nil {
		cfg.DefaultHeader = make(map[string]string)
	}
	if cfg.UserAgent == "" {
		cfg.UserAgent = fmt.Sprintf("%s/%s", common.LibName, common.LibVersion)
	}

	c := &APIClient{}
	c.client = &common.Client{}
	c.client.Cfg = cfg

	if cfg.Environment != "" {
		c.SetEnvironment(cfg.Environment, cfg.LiveEndpointURLPrefix)
	}

	// API Services
	c.Payments = &payments.Payments{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/pal/servlet/Payment/%s", c.client.Cfg.Endpoint, PaymentAPIVersion)
		},
	}

	c.Recurring = &recurring.GeneralApi{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/pal/servlet/Recurring/%s", c.client.Cfg.Endpoint, RecurringAPIVersion)
		},
	}

	c.BinLookup = &binlookup.BinLookup{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/%s/%s", c.client.Cfg.Endpoint, BinLookupPalSuffix, BinLookupAPIVersion)
		},
	}

	c.PlatformsAccount = &platformsaccount.PlatformsAccount{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/Account/%s", c.client.Cfg.MarketPayEndpoint, MarketpayAccountAPIVersion)
		},
	}

	c.PlatformsFund = &platformsfund.PlatformsFund{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/Fund/%s", c.client.Cfg.MarketPayEndpoint, MarketpayFundAPIVersion)
		},
	}

	c.PlatformsHostedOnboardingPage = &platformshostedonboardingpage.PlatformsHostedOnboardingPage{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/Hop/%s", c.client.Cfg.MarketPayEndpoint, MarketpayHopAPIVersion)
		},
	}

	c.PlatformsNotificationConfiguration = &platformsnotificationconfiguration.PlatformsNotificationConfiguration{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/Notification/%s", c.client.Cfg.MarketPayEndpoint, MarketpayNotificationAPIVersion)
		},
	}

	c.Disputes = &disputes.Disputes{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/%s", c.client.Cfg.DisputesEndpoint, DisputesAPIVersion)
		},
	}

	c.StoredValue = &storedvalue.StoredValue{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/pal/servlet/StoredValue/%s", c.client.Cfg.Endpoint, StoredValueAPIVersion)
		},
	}

	c.PosTerminalManagement = &posterminalmanagement.GeneralApi{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/%s", c.client.Cfg.PosTerminalManagementEndpoint, PosTerminalManagementAPIVersion)
		},
	}

	c.Transfers = &transfers.GeneralApi{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/%s", c.client.Cfg.TransfersEndpoint, TransfersAPIVersion)
		},
	}

	c.BalancePlatform = balanceplatform.NewAPIClient(c.client)
	c.Management = management.NewAPIClient(c.client)
	c.LegalEntity = legalentity.NewAPIClient(c.client)

	return c
}

func (c *APIClient) Checkout() *checkout.APIClient {
	return checkout.NewAPIClient(c.client)
}

func (c *APIClient) Payout() *payout.APIClient {
	if c.payout == nil {
		c.payout = payout.NewAPIClient(c.client)
		c.payout.InitializationApi.BasePath = func() string {
			return fmt.Sprintf("%s/pal/servlet/Payout/%s", c.client.Cfg.Endpoint, PaymentAPIVersion)
		}
	}
	return c.payout
}

/*
SetEnvironment This defines the payment environment for live or test

  - @param environment           This defines the payment environment live or test
  - @param liveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
*/
func (c *APIClient) SetEnvironment(env common.Environment, liveEndpointURLPrefix string) {
	if env == common.LiveEnv {
		c.client.Cfg.Environment = env
		c.client.Cfg.MarketPayEndpoint = MarketpayEndpointLive
		c.client.Cfg.DisputesEndpoint = DisputesEndpointLive
		if liveEndpointURLPrefix != "" {
			c.client.Cfg.Endpoint = EndpointProtocol + liveEndpointURLPrefix + EndpointLiveSuffix
			c.client.Cfg.CheckoutEndpoint = EndpointProtocol + liveEndpointURLPrefix + CheckoutEndpointLiveSuffix
		} else {
			c.client.Cfg.Endpoint = EndpointLive
			c.client.Cfg.CheckoutEndpoint = ""
		}
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointLive
		c.client.Cfg.BalancePlatformEndpoint = BalancePlatformEndpointLive
		c.client.Cfg.TransfersEndpoint = TransfersEndpointLive
		c.client.Cfg.ManagementEndpoint = ManagementEndpointLive
		c.client.Cfg.LegalEntityEndpoint = LegalEntityEntityLive
		c.client.Cfg.PosTerminalManagementEndpoint = PosTerminalManagementEndpointLive
	} else {
		c.client.Cfg.Environment = env
		c.client.Cfg.Endpoint = EndpointTest
		c.client.Cfg.MarketPayEndpoint = MarketpayEndpointTest
		c.client.Cfg.CheckoutEndpoint = CheckoutEndpointTest
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointTest
		c.client.Cfg.DisputesEndpoint = DisputesEndpointTest
		c.client.Cfg.BalancePlatformEndpoint = BalancePlatformEndpointTest
		c.client.Cfg.TransfersEndpoint = TransfersEndpointTest
		c.client.Cfg.ManagementEndpoint = ManagementEndpointTest
		c.client.Cfg.LegalEntityEndpoint = LegalEntityEntityTest
		c.client.Cfg.PosTerminalManagementEndpoint = PosTerminalManagementEndpointTest
	}

	c.client.Cfg.CheckoutEndpoint += "/" + CheckoutAPIVersion
	c.client.Cfg.BalancePlatformEndpoint += "/" + BalancePlatformAPIVersion
	c.client.Cfg.ManagementEndpoint += "/" + ManagementAPIVersion
	c.client.Cfg.LegalEntityEndpoint += "/" + LegalEntityAPIVersion
}

// GetConfig Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *common.Config {
	return c.client.Cfg
}
