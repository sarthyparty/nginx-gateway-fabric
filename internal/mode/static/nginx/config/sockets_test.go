package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetSocketNameTLS(t *testing.T) {
	res := getSocketNameTLS(800, "*.cafe.example.com")

	g := NewGomegaWithT(t)
	g.Expect(res).To(Equal("unix:/var/run/nginx/*.cafe.example.com800.sock"))
}

func TestGetSocketNameHTTPS(t *testing.T) {
	res := getSocketNameHTTPS(800)

	g := NewGomegaWithT(t)
	g.Expect(res).To(Equal("unix:/var/run/nginx/https800.sock"))
}