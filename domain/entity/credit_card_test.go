package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("400000000000000000", "Guilherme Rigotti", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5321254943663934", "Guilherme Rigotti", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5321254943663934", "Guilherme Rigotti", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5321254943663934", "Guilherme Rigotti", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5321254943663934", "Guilherme Rigotti", 10, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard("5321254943663934", "Guilherme Rigotti", 10, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}
