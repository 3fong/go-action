package main

import (
	"fmt"
	"log"

)

type user interface {
	login() string
}

type sysuser struct {
	name string
	age  int
}

func (sys sysuser) login() string {
	log.Println("sysuser")
	return fmt.Sprintf("%s %d", sys.name, sys.age)
}

type tenantuser struct {
	name        string
	age         int
	enablelogin bool
}

func (tenant tenantuser) login() string {
	log.Println("tenantuser")
	return fmt.Sprintf("%s %d", tenant.name, tenant.age)
}

type role int

func (r role) Value() (u user) {
	switch r {
	case 1:
		u = sysuser{
			name: "dalong",
			age:  333,
		}
	case 2:
		u = tenantuser{
			name:        "dalongrong",
			age:         30,
			enablelogin: true,
		}
	}
	return u
}
func main() {
	var id int = 2
	r := role(id)
	if r.Value() != nil {
		log.Printf("%s", r.Value().login())
	} else {
		log.Println("is nil")
	}
}