package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"localhost/lib/reciever"
	"localhost/lib/secure"
	"localhost/lib/sender"
	"localhost/lib/server"
	"localhost/lib/user"
	"log"
	"os"
	"sync"

	flatbuffers "github.com/google/flatbuffers/go"
	"gopkg.in/yaml.v2"
)

type userSenderReceiver struct {
	reciever *reciever.ReceiveData
	sender   *sender.SendData
	wg       *sync.WaitGroup
}

func main() {
	// get Current working directory
	cwdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// create a server
	server := server.NewServer("8080")
	usr := &userSenderReceiver{
		reciever: nil,
		sender:   nil,
		wg:       &sync.WaitGroup{},
	}
	// Start running the server
	server.ConfigureAndStart(usr.reciever, usr.sender, usr.wg)
	defer server.Stop()

	// read the new users profile
	testdata := cwdir + "/../lib/user/testdata/user.yaml"
	userdata := readYamlFile(testdata)
	userDetails := readUserProfile(data)
	b := flatbuffers.NewBuilder(0)
	fbuf := writeUser(b, userdata)

	// It is assumed that the reciver will do a sha512 hash -----------------
	data := hashData(userDetails.Profile.Verification.EncryptedPass)
	// Code above -----------------------------------------------------------
	userDetails.Profile.Verification.EncryptedPass = secure.HashValue(data)
	fmt.Println(userDetails.Profile.Verification.EncryptedPass)
	// Apply Argon2i encryption on use details
	email := userDetails.Profile.Verification.EncryptedEmail
	password := userDetails.Profile.Verification.EncryptedPass
	mobileNo := userDetails.Profile.Verification.EncryptedMobileNo
	encryptUserProfile(&userDetails.Profile)
	fmt.Println(verifyUserProfile(email, password, mobileNo, &userDetails.Profile))
	fmt.Println(verifyUserProfile("email", password, mobileNo, &userDetails.Profile))
	fmt.Println(verifyUserProfile(email, "password", mobileNo, &userDetails.Profile))
	fmt.Println(verifyUserProfile(email, password, "mobileNo", &userDetails.Profile))
	usr.wg.Wait()
}

// encryptUserProfile apply argon2 encryption
func encryptUserProfile(user *user.Profile) {
	user.Verification.EncryptedEmail = secure.HashValue(user.Verification.EncryptedEmail)

	user.Verification.EncryptedMobileNo = secure.HashValue(user.Verification.EncryptedMobileNo)

	user.Verification.EncryptedPass = secure.HashValue(user.Verification.EncryptedPass)
}

func verifyUserProfile(email, password, number string, user *user.Profile) bool {
	// This can be done better ...
	if validator, err := secure.VerifyValue(email, user.Verification.EncryptedEmail); err != nil || validator == false {
		// send a response back, saying invalid email
		return validator
	}
	if validator, err := secure.VerifyValue(password, user.Verification.EncryptedPass); err != nil || validator == false {
		user.Verification.IncorrectCounts++
		// send a responce back saying invalid password
		return validator
	}
	if validator, err := secure.VerifyValue(number, user.Verification.EncryptedPass); err != nil || validator == false {
		// invalid number
		return validator
	}
	return true
}

// hashData Assumed that receiver first hashes the data before proceeding
func hashData(data string) string {
	password := sha512.New()
	password.Sum(([]byte(data)))
	return hex.EncodeToString(password.Sum(nil))
}

func readYamlFile(path string) []byte {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	return yamlFile
}

func readUserProfile(yamlFile []byte) *user.User {
	userDetails := &user.User{}
	err := yaml.Unmarshal(yamlFile, userDetails)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return userDetails
}
