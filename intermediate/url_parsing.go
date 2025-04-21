package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Full URL:", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	rawURL1 := "https://example.com:8080/path?name=John&age=30"

	parsedURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// RawQuery returns a string
	// Query() returns a Map and to access we use Get(k)

	queryParams := parsedURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	// Building URL
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseUrl.Query()
	query.Set("name", "John")
	query.Set("age", "30")
	baseUrl.RawQuery = query.Encode() // Encode format query as URL format -> name=John

	fmt.Println("Built URL:", baseUrl.String())

	values := url.Values{}

	// Add key-value pairs to the value object
	values.Add("name", "Gopher")
	values.Add("age", "18")
	values.Add("color", "Blue")
	values.Add("city", "London")
	values.Add("country", "UK")

	// Encode
	encodedQuery := values.Encode()
	fmt.Println("EncodedQuery:", encodedQuery)

	// Build a URL
	baseURL1 := "https://example.com/search"
	fullUrl := baseURL1 + "?" + encodedQuery

	fmt.Println(fullUrl)
}
