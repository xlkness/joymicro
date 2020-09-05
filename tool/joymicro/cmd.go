package main

type Cmd interface {
	Name() string
	Desc() string
	Help() string
	CheckArgs(...string) error
	Exec(...string)
}
