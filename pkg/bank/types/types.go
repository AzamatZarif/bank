package types

// Money представляеть собой денежную сумму в мин. единицах (центы, копейки, дирамы и т.д.).
type Money int64

// Category представляет собой категориюб в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
type Category string

// Status представляет статус платежа.
type Status string

// Payment представляет информацию о платеже.
type Payment struct {
	ID int
	Amount Money
	Category Category
	Status Status
}