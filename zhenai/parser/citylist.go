package parser
import (
  "crawler/engine"
  "regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(bytes []byte) engine.ParseResult {
  re := regexp.MustCompile(cityListRe)
  submatch := re.FindAllSubmatch(bytes, -1)
  result := engine.ParseResult{}

  for _,item := range submatch {
    result.Items = append(result.Items, "City:"+string(item[2]))
    result.Requests = append(result.Requests, engine.Request{
      Url: string(item[1]),
      ParseFunc: func([]byte) ParseResult {
      	return ParseResult{}
      },
    })
  }
  return result
}
