package entity

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestNegative(t *testing.T) {
	g := gomega.NewWithT(t)

	f := Books{
		Title: "Nawinza155",
		Price: -5,
		Code:  "BK123456",
	}
	g.Expect(f)
}
