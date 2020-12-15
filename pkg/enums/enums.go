package enums

type enum interface {
    CodeToName(code string) string
    NameToCode(name string) string
}
