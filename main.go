package main

import (
	"fmt"

	s1 "github.com/hughluo/semver"
	s2 "github.com/hughluo/semver/v2"
)

func main() {
	fmt.Printf("the semantic version is %s\n", s1.SemVer())
	fmt.Printf("the semantic version is %s\n", s2.SemVer())
}
