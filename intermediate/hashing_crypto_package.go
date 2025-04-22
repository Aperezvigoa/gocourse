package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	// Creating the worst password possible
	password := "Password123"

	// Hashing with SHA256
	// hash256 := sha256.Sum256([]byte(password))
	// fmt.Println("Byte slice SHA256 of password:", hash256)
	// fmt.Printf("SHA-256: %x\n", hash256)

	// Hashing with SHA512
	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println("Byte slice SHA512 of password:", hash512)
	// fmt.Printf("SHA-512: %x\n", hash512)

	/*
	 * Salting adds an extra layer of security by combining the password
	 * with a unique random value. So it can be random, or you can store
	 * a string as a salt and use that with every password that you are
	 * hashing.
	 * This practice of salting, helps to protect against dictionary
	 * attacks and rainbow table attacks.
	 *
	 * To sum all together, salt is a value added to the password before
	 * hashing, and its purpose is that it ensures that even if two users
	 * have the same password, theis hashed values will be different
	 */

	salt, err := generateSalt()
	if err != nil {
		fmt.Println(err)
		return
	}
	signUpHash := HashPassword(password, salt)
	fmt.Printf("Password: %s\n", signUpHash)

	// Storing salt and password in database
	saltStr := base64.StdEncoding.EncodeToString(salt)

	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt: %x\n", salt)
	fmt.Println("Salt:", saltStr)    // simulate as storing in data base
	fmt.Println("Hash:", signUpHash) // simulate as storing in data base

	// verify
	// retrieve the saltStr and decode
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	loginHash := HashPassword("Password12345", decodedSalt)

	if loginHash == signUpHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}
}

// Function to get salt
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt) // Readfull gets randomBytes from rand.Reader, and reads it into salt
	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Function to hash password
func HashPassword(p string, salt []byte) string {
	saltedPassword := append(salt, []byte(p)...)
	hashedPassword := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hashedPassword[:])
}
