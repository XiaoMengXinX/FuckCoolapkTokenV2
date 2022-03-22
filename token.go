package token

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/XiaoMengXinX/crypto/bcrypt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var manufactor = "Google"
var brand = "Google"
var model = []string{
	"Pixel 3",
	"Pixel 3 XL",
	"Pixel 3a",
	"Pixel 4",
	"Pixel 4 XL",
	"Pixel 4a",
	"Pixel 5",
	"Pixel 5a",
}
var buildNumber = []string{
	"SQ1A.220105.002",
	"SQ1D.220105.007",
}

// GetToken Generate a token with random device info
func GetToken() (randDeviceCode, token string) {
	//rand.Seed(time.Now().UnixNano())
	randDeviceCode = CreateDeviceCode(randHexString(16), randMacAdress(), manufactor, brand, model[rand.Intn(len(model))], buildNumber[rand.Intn(len(buildNumber))])
	return randDeviceCode, GetTokenWithDeviceCode(randDeviceCode)
}

// GetTokenWithDeviceCode Generate a token with your device code
func GetTokenWithDeviceCode(deviceCode string) string {
	var timeStamp int64 = time.Now().Unix()
	base64TimeStamp := base64.RawStdEncoding.EncodeToString([]byte(strconv.FormatInt(timeStamp, 10)))
	md5TimeStamp := md5.Sum([]byte(strconv.FormatInt(timeStamp, 10)))
	md5DeviceCode := md5.Sum([]byte(deviceCode))

	token := fmt.Sprintf("token://com.coolapk.market/dcf01e569c1e3db93a3d0fcf191a622c?%s$%s&com.coolapk.market", fmt.Sprintf("%x", md5TimeStamp), fmt.Sprintf("%x", md5DeviceCode))
	base64Token := base64.RawStdEncoding.EncodeToString([]byte(token))
	md5Base64Token := md5.Sum([]byte(base64Token))
	md5Token := md5.Sum([]byte(token))

	bcryptSalt := fmt.Sprintf("%s/%xu", base64TimeStamp[:14], md5Token[:3])
	bcryptresult, _ := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%x", md5Base64Token[:])), []byte(bcryptSalt), 10)

	appToken := "v2" + base64.RawStdEncoding.EncodeToString(bcryptresult)

	return appToken
}

// CreateDeviceCode Generace your custom device code
func CreateDeviceCode(aid, mac, manufactor, brand, model, buildNumber string) string {
	return reverseString(base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("%s; ; ; %s; %s; %s; %s; %s", aid, mac, manufactor, brand, model, buildNumber))))
}

func randMacAdress() string {
	rand.Seed(time.Now().UnixNano())
	var macAdress string
	for i := 0; i < 6; i++ {
		macAdress += fmt.Sprintf("%02x", rand.Intn(256))
		if i != 5 {
			macAdress += ":"
		}
	}
	return macAdress
}

func randHexString(n int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = byte(rand.Intn(256))
	}
	return strings.ToUpper(hex.EncodeToString(bytes))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
