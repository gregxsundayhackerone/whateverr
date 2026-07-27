package main

// XSSFILE1 marker
func Vulnerable(x string) string {
	return "<div>" + x + "</div>" // no escaping
}
