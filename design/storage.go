package design

import "github.com/goadesign/gorma"

var sg = StorageGroup("Profile Storage Group", func() {
	Description("Global storage group.")
	Store("postgres", gorma.Postgres, func() {
		Description("Postgres relational store")
		Model("User", func() {

		})
	})
})
