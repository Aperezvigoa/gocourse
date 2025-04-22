package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~~o, Base64 Encoding")

	// Encode to Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Endoded 64:", encoded)

	// Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error while decoding:", err)
		return
	}
	fmt.Printf("Decoded: %s\n", decoded)

	// URL Safe Encoding, avoid '/' and '+'
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encoded:", urlSafeEncoded)

	urlSafeDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoded)
	if err != nil {
		fmt.Println("Error decoding URL:", urlSafeDecoded)
		return
	}

	fmt.Println("URL Safe Decoded:", string(urlSafeDecoded))
}
