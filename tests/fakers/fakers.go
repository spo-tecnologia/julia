package fakers

import (
	"fmt"
	"time"

	"math/rand"
)

func RandomString(length ...int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	defaultLength := 20
	if len(length) > 0 {
		defaultLength = length[0]
	}

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, defaultLength)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func CpfCnpj() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(2) == 0 {
		return Cpf()
	}
	return Cnpj()
}

func Cnpj() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	baseCnpj := fmt.Sprintf("%08d0001", r.Intn(99999999))
	return fmt.Sprintf("%s%s", baseCnpj, calculateCnpjDigits(baseCnpj))
}

func calculateDigit(cnpj []int, weights []int) int {
	sum := 0
	for i := 0; i < len(weights); i++ {
		sum += cnpj[i] * weights[i]
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func calculateCnpjDigits(baseCnpj string) string {
	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	cnpj := make([]int, 12)
	for i := 0; i < 12; i++ {
		cnpj[i] = int(baseCnpj[i] - '0')
	}

	d1 := calculateDigit(cnpj, weights1)
	cnpj = append(cnpj, d1)
	d2 := calculateDigit(cnpj, weights2)
	cnpj = append(cnpj, d2)

	return fmt.Sprintf("%d%d", d1, d2)
}

func Cpf() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	baseCpf := fmt.Sprintf("%09d", r.Intn(999999999))
	return fmt.Sprintf("%s%s", baseCpf, calculateCpfDigits(baseCpf))
}

func calculateCpfDigit(cpf []int, weights []int) int {
	sum := 0
	for i := 0; i < len(weights); i++ {
		sum += cpf[i] * weights[i]
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func calculateCpfDigits(baseCpf string) string {
	weights1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	cpf := make([]int, 9)
	for i := 0; i < 9; i++ {
		cpf[i] = int(baseCpf[i] - '0')
	}

	d1 := calculateCpfDigit(cpf, weights1)
	cpf = append(cpf, d1)
	d2 := calculateCpfDigit(cpf, weights2)
	cpf = append(cpf, d2)

	return fmt.Sprintf("%d%d", d1, d2)
}
