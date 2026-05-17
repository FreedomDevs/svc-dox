package main

import (
	"log"
	"net/netip"
	"strings"
	"svc-dox/responses"
	"svc-dox/responses/codes"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

func getLocalName(lang string, names map[string]string) string {
	// Если мапы нет в базе (город/страна не найдены), возвращаем заглушку
	if names == nil {
		return "Unknown"
	}

	// 1. Проверяем запрошенный язык (например, "ru")
	if name, exists := names[lang]; exists && name != "" {
		return name
	}

	// 2. Откатываемся на английский, если запрошенного языка не нашлось
	if name, exists := names["en"]; exists && name != "" {
		return name
	}

	// 3. Крайний случай: если нет ни "lang", ни "en", возвращаем любой имеющийся перевод
	for _, name := range names {
		if name != "" {
			return name
		}
	}

	return "Unknown"
}

func main() {
	cityDB, err := geoip2.Open("dbs/city.mmdb")
	if err != nil {
		log.Fatalf("Не удалось открыть City DB: %v", err)
	}
	defer cityDB.Close()

	asnDB, err := geoip2.Open("dbs/asn.mmdb")
	if err != nil {
		log.Fatalf("Не удалось открыть ASN DB: %v", err)
	}
	defer asnDB.Close()

	r := gin.Default()

	r.GET("/dox/ip", func(c *gin.Context) {
		ip := c.Query("ip")
		if ip == "" {
			responses.SendErrorResponse(codes.ErrIpNotProvided, nil, c)
		}

		lang := c.DefaultQuery("lang", "ru")

		rawIPs := strings.Split(ip, ",")
		response := make([]gin.H, 0, len(rawIPs))
		for _, rawIP := range rawIPs {
			addr, err := netip.ParseAddr(strings.TrimSpace(rawIP))
			if err != nil {
				response = append(response, gin.H{"result": "invalid ip format"})
				continue
			}

			if addr.IsLoopback() {
				response = append(response, gin.H{"result": "localhost"})
				continue
			}

			if addr.IsPrivate() {
				response = append(response, gin.H{"result": "private ip"})
				continue
			}

			if addr.IsLinkLocalUnicast() {
				response = append(response, gin.H{"result": "link local"})
				continue
			}

			if addr.IsMulticast() {
				response = append(response, gin.H{"result": "multicast"})
				continue
			}

			if !addr.IsGlobalUnicast() {
				response = append(response, gin.H{"result": "ip not exists"})
				continue
			}

			netIP := addr.AsSlice()
			netresponse := gin.H{}

			cityRecord, err := cityDB.City(netIP)
			if err != nil {
				response = append(response, gin.H{"result": "ip not found"})
				continue
			}

			netresponse["continent_code"] = cityRecord.Continent.Code
			netresponse["continent"] = getLocalName(lang, cityRecord.Continent.Names)
			netresponse["is_in_european_union"] = cityRecord.Country.IsInEuropeanUnion
			netresponse["country_code"] = cityRecord.Country.IsoCode
			netresponse["country"] = getLocalName(lang, cityRecord.Country.Names)

			subdivisions := make([]string, 0, len(cityRecord.Subdivisions))
			for _, subdivision := range cityRecord.Subdivisions {
				subdivisions = append(subdivisions, getLocalName(lang, subdivision.Names))
			}

			netresponse["subdivisions"] = subdivisions
			netresponse["city"] = getLocalName(lang, cityRecord.City.Names)
			netresponse["latitude"] = cityRecord.Location.Latitude
			netresponse["longitude"] = cityRecord.Location.Longitude

			asnRecord, err := asnDB.ASN(netIP)
			if err != nil {
				response = append(response, gin.H{"result": "ip not found"})
				continue
			}

			netresponse["ASN"] = asnRecord.AutonomousSystemNumber
			netresponse["ASO"] = asnRecord.AutonomousSystemOrganization

			response = append(response, netresponse)
		}

		responses.SendSuccessResponse(codes.SuccessIpDox, response, c)
	})

	r.Run("[::]:80")
}
