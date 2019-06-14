package engine
import (
  "crawler/fetcher"
  "log"
)

func Run(seeds ...Request) {
  var requests []Request

  for _, r := range seeds {
    requests = append(requests,r)
  }

  for len(requests) > 0 {
    request := requests[0]
    requests := requests[1:]
    log.Printf("Fetching %s\n", request.Url)
    content, err := fetcher.Fetch(request.Url)
    if err != nil {
      log.Printf("Fetch error, Url: %s %v\n", request.Url, err)
      continue
    }

    parseResult := request.ParseFunc(content)

    requests := append(requests, parseResult.Requests...)

    for _, item := range parseResult.Items {
      log.Printf("Got item %v\n", item)
    }

  }

}
