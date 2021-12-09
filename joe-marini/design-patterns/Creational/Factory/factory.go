package main

import "fmt"

// newPublication is a factory function that creates the specified publication type
func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	// Create the right kind of publication based on the given type
	switch {
	case pubType == "magazine":
		return createMagazine(name, pg, pub), nil
	case pubType == "newspaper":
		return createNewspaper(name, pg, pub), nil
	default:
		return nil, fmt.Errorf("no such publication type")
	}
}
