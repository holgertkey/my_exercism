package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
    validLogPattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
    re := regexp.MustCompile(validLogPattern)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    re := regexp.MustCompile(`<[-~*=]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    count := 0
    re := regexp.MustCompile(`"[^"]*([Pp][Aa][Ss][Ss][Ww][Oo][Rr][Dd])[^"]*"`)
    for _, line := range lines {
        if re.MatchString(line) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+(\S+)`)
    result := make([]string, len(lines))
    for i, line := range lines {
        if match := re.FindStringSubmatch(line); match != nil {
            result[i] = "[USR] " + match[1] + " " + line
        } else {
            result[i] = line
        }
    }
    return result
}
