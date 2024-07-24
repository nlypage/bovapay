package main

import (
	"github.com/nlypage/bovapay/bovapay"
	"log"
)

func main() {
	bovaPayClient := bovapay.NewClient(bovapay.Options{
		APIKey: "9842080356958f88769019f1e992ddb208f7e0f5",
		UserID: "f4eb6120-33de-4a9b-a972-95488690ac23",
	})
	deposit, err := bovaPayClient.CreateDeposit(bovapay.CreateDepositRequest{
		MerchantID:       "1342w5e6r7",
		Amount:           90,
		PayeerIdentifier: "123",
		PayeerIP:         "127.0.0.1",
		CallbackURL:      "https://cryptoexchange01.info/bovapay/webhook",
		PayeerType:       bovapay.Primary,
		Lifetime:         10000,
		Currency:         bovapay.RUB,
		RedirectURL:      "https://t.me/Liberty_ExchangeBot",
	})

	if err != nil {
		log.Println(err)
	} else {
		log.Println(deposit)
	}
}
