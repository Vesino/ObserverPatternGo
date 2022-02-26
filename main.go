package main

import (
	"errors"
	"fmt"
	"strconv"
)

// In this example, subject could be, a conversation, a file sent, a price stock changes etc...
type Subject interface {
	Suscribe(o Observer) (bool, error)
	Unsuscribe(o Observer) (bool, error)
	Notify(o Observer) (bool, error)
}

type Observer interface {
	Update(string)
}

type StockObserver struct {
	name string
}

func (s *StockObserver) Update(t string) {
	// do something
	fmt.Println("Stock Observer:", s.name, "has been updated, received subject string:", t)
}

// This is a concrete subject
type StockMonitor struct {
	ticker string
	price  float64

	observers []Observer
}

func (s *StockMonitor) Suscribe(o Observer) (bool, error) {
	for _, observer := range s.observers {
		if observer == o {
			return false, errors.New("Observer already exixts")
		}
	}
	s.observers = append(s.observers, o)
	return true, nil
}

func (s *StockMonitor) Unsuscribe(o Observer) (bool, error) {
	for i, observer := range s.observers {
		if observer == o {
			// remove it from the slice of observers
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Observer not found")
}

func (s *StockMonitor) Notify() (bool, error) {
	for _, observer := range s.observers {
		observer.Update(s.String())
	}
	return true, nil
}

func (s *StockMonitor) SetPrice(price float64) {
	s.price = price
	s.Notify()
}

func (s *StockMonitor) String() string {
	convertFloatToString := strconv.FormatFloat(s.price, 'f', 2, 64)
	return "Stock Monitor: " + s.ticker + " $" + convertFloatToString
}

func main() {
	// Creation of a ner Stock for monitoring
	stockMonitor := &StockMonitor{
		ticker: "APPL",
		price:  0.0,
	}

	observerA := &StockObserver{
		name: "Stock Observer A",
	}

	observerB := &StockObserver{
		name: "Stock Observer B",
	}

	// Suscribe our observers to the Stock Monitor
	stockMonitor.Suscribe(observerA)
	stockMonitor.Suscribe(observerB)

	stockMonitor.SetPrice(500)

	stockMonitor.Unsuscribe(observerA)

	stockMonitor.SetPrice(528)
}

// Output
// Stock Observer: Stock Observer A has been updated, received subject string: Stock Monitor: APPL $500.00
// Stock Observer: Stock Observer B has been updated, received subject string: Stock Monitor: APPL $500.00
// Stock Observer: Stock Observer B has been updated, received subject string: Stock Monitor: APPL $528.00
