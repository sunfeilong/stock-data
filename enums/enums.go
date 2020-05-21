package enums

type enum interface {
    codeToName(code string) string
    nameToCode(name string) string
}
