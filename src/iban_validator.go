package src

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type IBANValidatorService struct {
	countryCodesAndLength []string
}

func NewIBANValidatorService() *IBANValidatorService {
	return &IBANValidatorService{
		countryCodesAndLength: []string{
			"AL-28", "AD-24", "AT-20", "AZ-28", "BH-22", "BY-28", "BE-16",
			"BA-20", "BR-29", "BG-22", "CR-22", "HR-21", "CY-28", "CZ-24",
			"DK-18", "DO-28", "EG-29", "SV-28", "EE-20", "FO-18", "FI-18",
			"FR-27", "GE-22", "DE-22", "GI-23", "GR-27", "GL-18", "GT-28",
			"VA-22", "HU-28", "IS-26", "IQ-23", "IE-22", "IL-23", "IT-27",
			"JO-30", "KZ-20", "XK-20", "KW-30", "LV-21", "LB-28", "LY-25",
			"LI-21", "LT-20", "LU-20", "MT-31", "MR-27", "MU-30", "MD-24",
			"MC-27", "ME-22", "NL-18", "MK-19", "NO-15", "PK-24", "PS-29",
			"PL-28", "PT-25", "QA-29", "RO-24", "LC-32", "SM-27", "ST-25",
			"SA-24", "RS-22", "SC-31", "SK-24", "SI-19", "ES-24", "SD-18",
			"SE-24", "CH-21", "TL-23", "TN-24", "TR-26", "UA-29", "AE-23",
			"GB-22", "VG-24", "DZ-26", "AO-25", "BJ-28", "BF-28", "BI-16",
			"CM-27", "CV-25", "CF-27", "TD-27", "KM-27", "CG-27", "DJ-27",
			"GQ-27", "GA-27", "GW-25", "HN-28", "IR-26", "CI-28", "MG-27",
			"ML-28", "MA-28", "MZ-25", "NI-32", "NE-28", "SN-28", "TG-28",
		},
	}
}

func (s *IBANValidatorService) ValidateIBAN(iban string) *Response {
	if !s.checkIBANCountryCodeAndLength(iban) {
		return &Response{
			StatusCode: http.StatusBadRequest,
			Body:       iban[:2] + " is not a valid IBAN country code or has an invalid length.",
		}
	}

	if !s.checkForModulusOfOne(iban) {
		return &Response{
			StatusCode: http.StatusBadRequest,
			Body:       iban + "'s checksum does not have a modulus of 1.",
		}
	}

	return &Response{
		StatusCode: http.StatusOK,
		Body:       iban + " is a valid IBAN.",
	}
}

func (s *IBANValidatorService) checkIBANCountryCodeAndLength(iban string) bool {
	for _, countryCodeAndLength := range s.countryCodesAndLength {
		if strings.HasPrefix(iban, countryCodeAndLength[:2]) {
			if len(iban) == parseLength(countryCodeAndLength[3:]) {
				return true
			}
		}
	}
	return false
}

func parseLength(lengthStr string) int {
	length, _ := strconv.Atoi(lengthStr)
	return length
}

func (s *IBANValidatorService) checkForModulusOfOne(iban string) bool {
	convertedIBAN := s.convertAllAlphaCharactersToNumericCharacters(
		strings.ToUpper(iban[4:]) + strings.ToUpper(iban[:4]))

	for len(convertedIBAN) > 9 {
		modulus := parseModulus(convertedIBAN[:9])
		convertedIBAN = strconv.Itoa(modulus) + convertedIBAN[9:]
	}

	modulus := parseModulus(convertedIBAN)
	return modulus == 1
}

func parseModulus(s string) int {
	modulus, _ := strconv.Atoi(s)
	return modulus % 97
}

func (s *IBANValidatorService) convertAllAlphaCharactersToNumericCharacters(iban string) string {
	re := regexp.MustCompile("[A-Z]")
	return re.ReplaceAllStringFunc(iban, func(match string) string {
		return strconv.Itoa(int(match[0]) - 'A' + 10)
	})
}
