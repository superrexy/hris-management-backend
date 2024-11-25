package utils

import (
	"fmt"
	"strings"

	customValidation "hris-management/utils/validation"

	"github.com/go-playground/validator/v10"
)

const TAG = "Utils::PayloadValidation> "

var validation = validator.New()

type errorValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func PayloadValidation(s interface{}) (errors []errorValidation) {
	validation.RegisterValidation("timeformat", customValidation.TimeFormatValidation)
	validation.RegisterValidation("datetime", customValidation.ISO8601DateTimeValidation)
	if err := validation.Struct(s); err != nil {

		errors := make([]errorValidation, 0)
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, errorValidation{
				Field:   fieldToSnakeCase(err.Field()),
				Message: messageForTag(err),
			})
		}
		return errors
	}

	return nil
}

func messageForTag(fe validator.FieldError) string {
	field := fieldToSnakeCase(fe.Field())

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("The %s is required", field)
	case "email":
		return fmt.Sprintf("The %s should be a valid email address", field)
	case "timeformat":
		return fmt.Sprintf("The %s should be in 24-hour format (HH:MM)", field)
	case "gte":
		return fmt.Sprintf("The %s should be greater than or equal to %s", field, fe.Param())
	case "lte":
		return fmt.Sprintf("The %s should be less than or equal to %s", field, fe.Param())
	case "gt":
		return fmt.Sprintf("The %s should be greater than %s", field, fe.Param())
	case "lt":
		return fmt.Sprintf("The %s should be less than %s", field, fe.Param())
	case "oneof":
		return fmt.Sprintf("The %s should be one of %s", field, fe.Param())
	case "eqcsfield":
		return fmt.Sprintf("The %s should be equal to the %s field", field, fe.Param())
	case "eqfield":
		return fmt.Sprintf("The %s should be equal to the %s field", field, fe.Param())
	case "fieldcontains":
		return fmt.Sprintf("The %s should contain the characters '%s'", field, fe.Param())
	case "fieldexcludes":
		return fmt.Sprintf("The %s should not contain the characters '%s'", field, fe.Param())
	case "gtcsfield":
		return fmt.Sprintf("The %s should be greater than the %s field", field, fe.Param())
	case "gtecsfield":
		return fmt.Sprintf("The %s should be greater than or equal to the %s field", field, fe.Param())
	case "gtefield":
		return fmt.Sprintf("The %s should be greater than or equal to the %s field", field, fe.Param())
	case "gtfield":
		return fmt.Sprintf("The %s should be greater than the %s field", field, fe.Param())
	case "ltcsfield":
		return fmt.Sprintf("The %s should be less than the %s field", field, fe.Param())
	case "ltecsfield":
		return fmt.Sprintf("The %s should be less than or equal to the %s field", field, fe.Param())
	case "ltefield":
		return fmt.Sprintf("The %s should be less than or equal to the %s field", field, fe.Param())
	case "ltfield":
		return fmt.Sprintf("The %s should be less than the %s field", field, fe.Param())
	case "necsfield":
		return fmt.Sprintf("The %s should not be equal to the %s field", field, fe.Param())
	case "nefield":
		return fmt.Sprintf("The %s should not be equal to the %s field", field, fe.Param())
	case "cidr":
		return fmt.Sprintf("The %s should be a Classless Inter-Domain Routing (CIDR)", field)
	case "cidrv4":
		return fmt.Sprintf("The %s should be a Classless Inter-Domain Routing (CIDRv4)", field)
	case "cidrv6":
		return fmt.Sprintf("The %s should be a Classless Inter-Domain Routing (CIDRv6)", field)
	case "datauri":
		return fmt.Sprintf("The %s should be a Data URL", field)
	case "fqdn":
		return fmt.Sprintf("The %s should be a Full Qualified Domain Name (FQDN)", field)
	case "hostname":
		return fmt.Sprintf("The %s should be a Hostname (RFC 952)", field)
	case "hostname_port":
		return fmt.Sprintf("The %s should be a HostPort", field)
	case "hostname_rfc1123":
		return fmt.Sprintf("The %s should be a Hostname (RFC 1123)", field)
	case "ip":
		return fmt.Sprintf("The %s should be an Internet Protocol Address (IP)", field)
	case "ip4_addr":
		return fmt.Sprintf("The %s should be an Internet Protocol Address (IPv4)", field)
	case "ip6_addr":
		return fmt.Sprintf("The %s should be an Internet Protocol Address (IPv6)", field)
	case "ip_addr":
		return fmt.Sprintf("The %s should be an Internet Protocol Address (IP)", field)
	case "ipv4":
		return fmt.Sprintf("The %s should be an Internet Protocol Address (IPv4)", field)
	case "ipv6":
		return fmt.Sprintf("The %s should be an Internet Protocol Address (IPv6)", field)
	case "mac":
		return fmt.Sprintf("The %s should be a Media Access Control Address (MAC)", field)
	case "tcp4_addr":
		return fmt.Sprintf("The %s should be a Transmission Control Protocol Address (TCPv4)", field)
	case "tcp6_addr":
		return fmt.Sprintf("The %s should be a Transmission Control Protocol Address (TCPv6)", field)
	case "tcp_addr":
		return fmt.Sprintf("The %s should be a Transmission Control Protocol Address (TCP)", field)
	case "udp4_addr":
		return fmt.Sprintf("The %s should be a User Datagram Protocol Address (UDPv4)", field)
	case "udp6_addr":
		return fmt.Sprintf("The %s should be a User Datagram Protocol Address (UDPv6)", field)
	case "udp_addr":
		return fmt.Sprintf("The %s should be a User Datagram Protocol Address (UDP)", field)
	case "unix_addr":
		return fmt.Sprintf("The %s should be a Unix domain socket end point Address", field)
	case "uri":
		return fmt.Sprintf("The %s should be a URI String", field)
	case "url":
		return fmt.Sprintf("The %s should be a URL String", field)
	case "http_url":
		return fmt.Sprintf("The %s should be an HTTP URL String", field)
	case "url_encoded":
		return fmt.Sprintf("The %s should be URL Encoded", field)
	case "urn_rfc2141":
		return fmt.Sprintf("The %s should be a Urn RFC 2141 String", field)
	case "alpha":
		return fmt.Sprintf("The %s should contain only alphabetic characters", field)
	case "alphanum":
		return fmt.Sprintf("The %s should contain only alphanumeric characters", field)
	case "alphanumunicode":
		return fmt.Sprintf("The %s should contain only alphanumeric Unicode characters", field)
	case "alphaunicode":
		return fmt.Sprintf("The %s should contain only alphabetic Unicode characters", field)
	case "ascii":
		return fmt.Sprintf("The %s should contain only ASCII characters", field)
	case "boolean":
		return fmt.Sprintf("The %s should be a boolean", field)
	case "contains":
		return fmt.Sprintf("The %s should contain the substring '%s'", field, fe.Param())
	case "containsany":
		return fmt.Sprintf("The %s should contain any of the characters '%s'", field, fe.Param())
	case "containsrune":
		return fmt.Sprintf("The %s should contain the rune '%s'", field, fe.Param())
	case "endsnotwith":
		return fmt.Sprintf("The %s should not end with '%s'", field, fe.Param())
	case "endswith":
		return fmt.Sprintf("The %s should end with '%s'", field, fe.Param())
	case "excludes":
		return fmt.Sprintf("The %s should not contain the substring '%s'", field, fe.Param())
	case "excludesall":
		return fmt.Sprintf("The %s should not contain any of the characters '%s'", field, fe.Param())
	case "excludesrune":
		return fmt.Sprintf("The %s should not contain the rune '%s'", field, fe.Param())
	case "lowercase":
		return fmt.Sprintf("The %s should be in lowercase", field)
	case "multibyte":
		return fmt.Sprintf("The %s should contain multi-byte characters", field)
	case "number":
		return fmt.Sprintf("The %s should be a number", field)
	case "numeric":
		return fmt.Sprintf("The %s should be numeric", field)
	case "printascii":
		return fmt.Sprintf("The %s should contain only printable ASCII characters", field)
	case "startsnotwith":
		return fmt.Sprintf("The %s should not start with '%s'", field, fe.Param())
	case "startswith":
		return fmt.Sprintf("The %s should start with '%s'", field, fe.Param())
	case "uppercase":
		return fmt.Sprintf("The %s should be in uppercase", field)
	case "base64":
		return fmt.Sprintf("The %s should be a Base64 String", field)
	case "base64url":
		return fmt.Sprintf("The %s should be a Base64URL String", field)
	case "base64rawurl":
		return fmt.Sprintf("The %s should be a Base64RawURL String", field)
	case "bic":
		return fmt.Sprintf("The %s should be a Business Identifier Code (ISO 9362)", field)
	case "bcp47_language_tag":
		return fmt.Sprintf("The %s should be a language tag (BCP 47)", field)
	case "btc_addr":
		return fmt.Sprintf("The %s should be a Bitcoin Address", field)
	case "btc_addr_bech32":
		return fmt.Sprintf("The %s should be a Bitcoin Bech32 Address (segwit)", field)
	case "credit_card":
		return fmt.Sprintf("The %s should be a Credit Card Number", field)
	case "mongodb":
		return fmt.Sprintf("The %s should be a MongoDB ObjectID", field)
	case "mongodb_connection_string":
		return fmt.Sprintf("The %s should be a MongoDB Connection String", field)
	case "cron":
		return fmt.Sprintf("The %s should be a Cron expression", field)
	case "spicedb":
		return fmt.Sprintf("The %s should be a SpiceDb ObjectID/Permission/Type", field)
	case "datetime":
		return fmt.Sprintf("The %s should be a Datetime", field)
	case "e164":
		return fmt.Sprintf("The %s should be an e164 formatted phone number", field)
	case "eth_addr":
		return fmt.Sprintf("The %s should be an Ethereum Address", field)
	case "hexadecimal":
		return fmt.Sprintf("The %s should be a Hexadecimal String", field)
	case "hexcolor":
		return fmt.Sprintf("The %s should be a Hexcolor String", field)
	case "hsl":
		return fmt.Sprintf("The %s should be an HSL String", field)
	case "hsla":
		return fmt.Sprintf("The %s should be an HSLA String", field)
	case "html":
		return fmt.Sprintf("The %s should contain HTML Tags", field)
	case "html_encoded":
		return fmt.Sprintf("The %s should be HTML Encoded", field)
	case "isbn":
		return fmt.Sprintf("The %s should be an International Standard Book Number", field)
	case "isbn10":
		return fmt.Sprintf("The %s should be an International Standard Book Number 10", field)
	case "isbn13":
		return fmt.Sprintf("The %s should be an International Standard Book Number 13", field)
	case "issn":
		return fmt.Sprintf("The %s should be an International Standard Serial Number", field)
	case "iso3166_1_alpha2":
		return fmt.Sprintf("The %s should be a two-letter country code (ISO 3166-1 alpha-2)", field)
	case "iso3166_1_alpha3":
		return fmt.Sprintf("The %s should be a three-letter country code (ISO 3166-1 alpha-3)", field)
	case "iso3166_1_alpha_numeric":
		return fmt.Sprintf("The %s should be a numeric country code (ISO 3166-1 numeric)", field)
	case "iso3166_2":
		return fmt.Sprintf("The %s should be a country subdivision code (ISO 3166-2)", field)
	case "iso4217":
		return fmt.Sprintf("The %s should be a currency code (ISO 4217)", field)
	case "json":
		return fmt.Sprintf("The %s should be a JSON", field)
	case "jwt":
		return fmt.Sprintf("The %s should be a JSON Web Token (JWT)", field)
	case "latitude":
		return fmt.Sprintf("The %s should be a Latitude", field)
	case "longitude":
		return fmt.Sprintf("The %s should be a Longitude", field)
	case "luhn_checksum":
		return fmt.Sprintf("The %s should pass the Luhn Algorithm Checksum", field)
	case "postcode_iso3166_alpha2":
		return fmt.Sprintf("The %s should be a Postcode", field)
	case "postcode_iso3166_alpha2_field":
		return fmt.Sprintf("The %s should be a Postcode", field)
	case "rgb":
		return fmt.Sprintf("The %s should be an RGB String", field)
	case "rgba":
		return fmt.Sprintf("The %s should be an RGBA String", field)
	case "ssn":
		return fmt.Sprintf("The %s should be a Social Security Number (SSN)", field)
	case "timezone":
		return fmt.Sprintf("The %s should be a Timezone", field)
	case "uuid":
		return fmt.Sprintf("The %s should be a Universally Unique Identifier (UUID)", field)
	case "uuid3":
		return fmt.Sprintf("The %s should be a Universally Unique Identifier (UUID v3)", field)
	case "uuid3_rfc4122":
		return fmt.Sprintf("The %s should be a Universally Unique Identifier (UUID v3 RFC4122)", field)
	case "uuid4":
		return fmt.Sprintf("The %s should be a Universally Unique Identifier (UUID v4)", field)
	case "uuid4_rfc4122":
		return fmt.Sprintf("The %s should be a Universally Unique Identifier (UUID v4 RFC4122)", field)
	case "uuid5":
		return fmt.Sprintf("The %s should be a Universally Unique Identifier (UUID v5)", field)
	case "uuid5_rfc4122":
		return fmt.Sprintf("The %s should be a Universally Unique Identifier (UUID v5 RFC4122)", field)
	case "uuid_rfc4122":
		return fmt.Sprintf("The %s should be a Universally Unique Identifier (UUID RFC4122)", field)
	case "md4":
		return fmt.Sprintf("The %s should be an MD4 hash", field)
	case "md5":
		return fmt.Sprintf("The %s should be an MD5 hash", field)
	case "sha256":
		return fmt.Sprintf("The %s should be a SHA256 hash", field)
	case "sha384":
		return fmt.Sprintf("The %s should be a SHA384 hash", field)
	case "sha512":
		return fmt.Sprintf("The %s should be a SHA512 hash", field)
	case "ripemd128":
		return fmt.Sprintf("The %s should be a RIPEMD-128 hash", field)
	case "ripemd160":
		return fmt.Sprintf("The %s should be a RIPEMD-160 hash", field)
	case "tiger128":
		return fmt.Sprintf("The %s should be a TIGER128 hash", field)
	case "tiger160":
		return fmt.Sprintf("The %s should be a TIGER160 hash", field)
	case "tiger192":
		return fmt.Sprintf("The %s should be a TIGER192 hash", field)
	case "semver":
		return fmt.Sprintf("The %s should be a Semantic Versioning 2.0.0", field)
	case "ulid":
		return fmt.Sprintf("The %s should be a Universally Unique Lexicographically Sortable Identifier (ULID)", field)
	case "cve":
		return fmt.Sprintf("The %s should be a Common Vulnerabilities and Exposures Identifier (CVE id)", field)
	case "eq":
		return fmt.Sprintf("The %s should be equal to %s", field, fe.Param())
	case "eq_ignore_case":
		return fmt.Sprintf("The %s should be equal to %s (case insensitive)", field, fe.Param())
	case "ne":
		return fmt.Sprintf("The %s should not be equal to %s", field, fe.Param())
	case "ne_ignore_case":
		return fmt.Sprintf("The %s should not be equal to %s (case insensitive)", field, fe.Param())
	case "dir":
		return fmt.Sprintf("The %s should be an existing directory", field)
	case "dirpath":
		return fmt.Sprintf("The %s should be a directory path", field)
	case "file":
		return fmt.Sprintf("The %s should be an existing file", field)
	case "filepath":
		return fmt.Sprintf("The %s should be a file path", field)
	case "image":
		return fmt.Sprintf("The %s should be an image", field)
	case "isdefault":
		return fmt.Sprintf("The %s should be the default value", field)
	case "len":
		return fmt.Sprintf("The %s should have a length of %s", field, fe.Param())
	case "max":
		return fmt.Sprintf("The %s should be at most %s", field, fe.Param())
	case "min":
		return fmt.Sprintf("The %s should be at least %s", field, fe.Param())
	case "required_if":
		return fmt.Sprintf("The %s is required if %s", field, fe.Param())
	case "required_unless":
		return fmt.Sprintf("The %s is required unless %s", field, fe.Param())
	case "required_with":
		return fmt.Sprintf("The %s is required with %s", field, fe.Param())
	case "required_with_all":
		return fmt.Sprintf("The %s is required with all of %s", field, fe.Param())
	case "required_without":
		return fmt.Sprintf("The %s is required without %s", field, fe.Param())
	case "required_without_all":
		return fmt.Sprintf("The %s is required without all of %s", field, fe.Param())
	case "excluded_if":
		return fmt.Sprintf("The %s should be excluded if %s", field, fe.Param())
	case "excluded_unless":
		return fmt.Sprintf("The %s should be excluded unless %s", field, fe.Param())
	case "excluded_with":
		return fmt.Sprintf("The %s should be excluded with %s", field, fe.Param())
	case "excluded_with_all":
		return fmt.Sprintf("The %s should be excluded with all of %s", field, fe.Param())
	case "excluded_without":
		return fmt.Sprintf("The %s should be excluded without %s", field, fe.Param())
	case "excluded_without_all":
		return fmt.Sprintf("The %s should be excluded without all of %s", field, fe.Param())
	case "unique":
		return fmt.Sprintf("The %s should be unique", field)
	case "iscolor":
		return fmt.Sprintf("The %s should be a valid color (hexcolor|rgb|rgba|hsl|hsla)", field)
	case "country_code":
		return fmt.Sprintf("The %s should be a valid country code (iso3166_1_alpha2|iso3166_1_alpha3|iso3166_1_alpha_numeric)", field)
	}

	return fe.Error()
}

func fieldToSnakeCase(field string) string {
	var result string
	for i, letter := range field {
		if 'A' <= letter && letter <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(letter + 32)
		} else {
			result += string(letter)
		}
	}
	return strings.ToLower(result)
}
