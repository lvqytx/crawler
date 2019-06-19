package fetcher

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"

  "golang.org/x/net/html/charset"
  "golang.org/x/text/encoding"
  "golang.org/x/text/encoding/unicode"
  "golang.org/x/text/transform"
)

func Fetch (url string)  ([]byte,error) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatalln(err)
  }
  req.Header.Set("User-Agant", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")

  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("wrong stat code: %d", resp.StatusCode)
  }

  bodyReader := bufio.NewReader(resp.Body)
  e := determineEncoding(bodyReader)
  utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

  return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
  bytes, err := r.Peek(1024)
  if err != nil {
    log.Printf("Fetch error %v\n", err)
    return unicode.UTF8
  }
  e, _, _ := charset.DetermineEncoding(bytes, "")
  return e
}
