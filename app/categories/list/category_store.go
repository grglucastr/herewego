package main

type CategoryStore interface {
	List() (string, error)
}
