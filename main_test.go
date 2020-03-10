package main

import "testing"

func benchmarkDNS(domainName string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FetchIPForDomain(domainName)
	}
}

func BenchmarkDNSUnknown(b *testing.B)  { benchmarkDNS("apple.com", b) }
func BenchmarkDNSKnown(b *testing.B)  { benchmarkDNS("google.com", b) }