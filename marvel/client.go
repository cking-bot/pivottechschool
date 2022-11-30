package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var pubkey, privkey = keys()
var BaseURL = "https://gateway.marvel.com/v1/public"
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func keys() (public, private string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	pubkey := os.Getenv("MARVEL_PUBLIC_KEY")
	privkey := os.Getenv("MARVEL_PRIVATE_KEY")
	return pubkey, privkey
}

type marvelClient struct {
	baseURL    string
	pubkey     string
	privkey    string
	httpClient *http.Client
}

func NewClient(url string) marvelClient {

	return marvelClient{
		url,
		pubkey,
		privkey,
		httpClient,
	}
}

func (c *marvelClient) md5Hash(ts int64) string {
	tsForHash := strconv.Itoa(int(ts))
	hash := md5.Sum([]byte(tsForHash + c.privkey + c.pubkey))
	return hex.EncodeToString(hash[:])
}

func (c *marvelClient) signURL(url string) string {
	ts := time.Now().Unix()
	hash := c.md5Hash(ts)
	return fmt.Sprintf("%s&ts=%d&apikey=%s&hash=%s", url, ts, c.pubkey, hash)
}

func (c *marvelClient) GetCharacters() ([]Character, error) {
	url := c.baseURL + "/characters?limit=25"
	url = c.signURL((url))

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var characterRes CharactersResponse
	if err := json.NewDecoder(res.Body).Decode(&characterRes); err != nil {
		return nil, err
	}
	return characterRes.Data.Results, nil
}
