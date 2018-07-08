package parser

import (
	"fmt"
	"regexp"

	"github.com/gohah/go-learn/crawler/engine"
	"github.com/gohah/go-learn/crawler_distributed/config"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	// compile := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	compile := regexp.MustCompile(cityListRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		// result.Items = append(
		// 	result.Items, "City "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	// fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
