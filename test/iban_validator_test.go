package ibanvalidator_test

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"iban-validator-golang/src"
)

func TestValidateIBAN(t *testing.T) {
	t.Run("with valid IBANs", func(t *testing.T) {
		ibans, err := readLinesFromFile("./data/list_of_valid_IBANs.txt")
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		for _, iban := range ibans {
			t.Run("IBAN: "+iban, func(t *testing.T) {
				status := src.NewIBANValidatorService().ValidateIBAN(iban)
				assert.Equal(t, 200, status.StatusCode, "valid IBAN: "+iban)
			})
		}
	})

	t.Run("with invalid IBANs", func(t *testing.T) {
		ibans, err := readLinesFromFile("./data/list_of_invalid_IBANs.txt")
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		for _, iban := range ibans {
			t.Run("IBAN: "+iban, func(t *testing.T) {
				status := src.NewIBANValidatorService().ValidateIBAN(iban)
				assert.Equal(t, 400, status.StatusCode, "invalid IBAN: "+iban)
			})
		}
	})

	t.Run("with invalid country code", func(t *testing.T) {
		iban := "ZE7280000810340009783242"

		status := src.NewIBANValidatorService().ValidateIBAN(iban)

		assert.Equal(t, 400, status.StatusCode)
		assert.Equal(t, "ZE is not a valid IBAN country code or has an invalid length.", status.Body)
	})

	t.Run("with too long length", func(t *testing.T) {
		iban := "SE72800008103400097832422"

		status := src.NewIBANValidatorService().ValidateIBAN(iban)

		assert.Equal(t, 400, status.StatusCode)
		assert.Equal(t, iban[0:2]+" is not a valid IBAN country code or has an invalid length.", status.Body)
	})

	t.Run("with invalid checksum", func(t *testing.T) {
		iban := "SE7280700810340009783242"

		status := src.NewIBANValidatorService().ValidateIBAN(iban)

		assert.Equal(t, 400, status.StatusCode)
		assert.Equal(t, iban+"'s checksum does not have a modulus of 1.", status.Body)
	})
}

func readLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Println("error closing file", err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines, scanner.Err()
}
