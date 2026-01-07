package entity

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestFarmHasCow(t *testing.T) {
	g := gomega.NewWithT(t)

	f := Books{
		Title: "Nawinza155",
		Price: 150,
		Code:  "BK123456",
	}
	g.Expect(f)
}
