package main

import "fmt"

func main() {
	creditPay := CreditCardPayment{CreditCardNumber: 123456789}
	paypalPay := PaypalPayment{AccountName: "Tom Laplace"}
	cryptoPay := CryptoPayment{WalletAddress: "x890a189"}

	Pay(creditPay, 500)
	Pay(paypalPay, 100)
	Pay(cryptoPay, 50)
}

type Payment interface {
	DoPayment(amount float64) error
}

type CreditCardPayment struct {
	CreditCardNumber int
}

func (cb CreditCardPayment) DoPayment(amount float64) error {
	fmt.Println("Paiement par carte bleue", cb.CreditCardNumber, ":", amount, "€")
	return nil
}

type PaypalPayment struct {
	AccountName string
}

func (p PaypalPayment) DoPayment(amount float64) error {
	fmt.Println("Paiement par Paypal par", p.AccountName, ":", amount, "€")
	return nil
}

type CryptoPayment struct {
	WalletAddress string
}

func (crypto CryptoPayment) DoPayment(amount float64) error {
	fmt.Println("Paiement par crypto", crypto.WalletAddress, ":", amount, "€")
	return nil
}

func Pay(p Payment, amount float64) {
	err := p.DoPayment(amount)
	if err != nil {
		fmt.Println("Erreur lors du paiement", err)
	}
}
