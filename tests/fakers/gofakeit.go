package fakers

import (
	"github.com/brianvoe/gofakeit"
)

func Test() {
	gofakeit.Email()
}
func Word() string {
	return gofakeit.Word()
}
func UUID() string {
	return gofakeit.UUID()
}
func Company() string {
	return gofakeit.Company()
}
func Name() string {
	return gofakeit.Name()
}
func NamePrefix() string {
	return gofakeit.NamePrefix()
}
func NameSuffix() string {
	return gofakeit.NameSuffix()
}
func FirstName() string {
	return gofakeit.FirstName()
}
func LastName() string {
	return gofakeit.LastName()
}
func Gender() string {
	return gofakeit.Gender()
}
func SSN() string {
	return gofakeit.SSN()
}
func Email() string {
	return gofakeit.Email()
}
func Phone() string {
	return gofakeit.Phone()
}
func PhoneFormatted() string {
	return gofakeit.PhoneFormatted()
}
func City() string {
	return gofakeit.City()
}
func Country() string {
	return gofakeit.Country()
}
func CountryAbr() string {
	return gofakeit.CountryAbr()
}
func State() string {
	return gofakeit.State()
}
func StateAbr() string {
	return gofakeit.StateAbr()
}
func Street() string {
	return gofakeit.Street()
}
func StreetName() string {
	return gofakeit.StreetName()
}
func StreetNumber() string {
	return gofakeit.StreetNumber()
}
func StreetPrefix() string {
	return gofakeit.StreetPrefix()
}
func StreetSuffix() string {
	return gofakeit.StreetSuffix()
}
func Zip() string {
	return gofakeit.Zip()
}
func Latitude() float64 {
	return gofakeit.Latitude()
}
func Longitude() float64 {
	return gofakeit.Longitude()
}
func FullAddress() string {
	return gofakeit.Address().Address + ", " + gofakeit.Address().City + " - " + gofakeit.Address().State + ", " + gofakeit.Address().Country
}
func URL() string {
	return gofakeit.URL()
}
func DomainName() string {
	return gofakeit.DomainName()
}
func DomainSuffix() string {
	return gofakeit.DomainSuffix()
}
func IPv4Address() string {
	return gofakeit.IPv4Address()
}
func IPv6Address() string {
	return gofakeit.IPv6Address()
}
func MacAddress() string {
	return gofakeit.MacAddress()
}
func HTTPMethod() string {
	return gofakeit.HTTPMethod()
}
func UserAgent() string {
	return gofakeit.UserAgent()
}
func ChromeUserAgent() string {
	return gofakeit.ChromeUserAgent()
}
func FirefoxUserAgent() string {
	return gofakeit.FirefoxUserAgent()
}
func OperaUserAgent() string {
	return gofakeit.OperaUserAgent()
}
func SafariUserAgent() string {
	return gofakeit.SafariUserAgent()
}
func Bool() bool {
	return gofakeit.Bool()
}
