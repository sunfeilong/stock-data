package token

import (
    "io/ioutil"
    "net/http"
    "strings"
)

func GetHKToken(url string) string {

    response, err := http.Get(url)
    if err != nil {
        logger.Panicw("Get HangKong Stock Market Data Exception", "err", err)
    }
    allBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        logger.Panicw("Get HangKong Stock Market Read Data Exception", "err", err)
    }

    split := strings.Split(string(allBytes), "\n")
    startGetToken := false
    for _, line := range split {
        if strings.Contains(line, TokenLineIndex) {
            startGetToken = true
            continue
        }
        if startGetToken {
            line = strings.TrimSpace(line)
            if strings.HasPrefix(line, "return") {
                line = strings.ReplaceAll(line, "return", "")
                line = strings.ReplaceAll(line, "\"", "")
                line = strings.ReplaceAll(line, ";", "")
                line = strings.ReplaceAll(line, " ", "")
                line = strings.TrimSpace(line)
                return line
            }
        }
    }
    return ""
}
