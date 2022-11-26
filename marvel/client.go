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

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func NewClient() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pubkey := os.Getenv("MARVEL_PUBLIC_KEY")
	privkey := os.Getenv("MARVEL_PRIVATE_KEY")

	client := marvelClient{
		baseURL: "https://gateway.marvel.com/v1/public",
		pubkey:  pubkey,
		privkey: privkey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	characters, err := client.getCharacters()
	if err != nil {
		log.Fatal(err)
	}

	for _, character := range characters {
		fmt.Println(character.Name, character.Description)
	}
}

type marvelClient struct {
	baseURL    string
	pubkey     string
	privkey    string
	httpClient *http.Client
}

func (c *marvelClient) md5Hash(ts int64) string {
	tsForHash := strconv.Itoa(int(ts))
	hash := md5.Sum([]byte(tsForHash + c.privkey + c.pubkey))
	return hex.EncodeToString(hash[:])
}

func (c *marvelClient) signURL(url string) string {
	ts := time.Now().Unix()
	hash := c.md5Hash(ts)
	return fmt.Sprintf("%s&ts=%d&apikey=%s&has=%s", url, ts, c.pubkey, hash)
}

func (c *marvelClient) getCharacters() ([]Characters, error) {
	url := c.baseURL + "/events?limit=25"
	url = c.signURL(url)
	spew.Dump(url)

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	spew.Dump(res.Status, res.StatusCode)

	var characterResponse CharactersResponse
	if err := json.NewDecoder(res.Body).Decode(&characterResponse); err != nil {
		return nil, err
	}
	return characterResponse.Data.Results, nil
}
