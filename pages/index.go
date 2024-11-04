package pages

import (
	"github.com/maddalax/htmgo/framework/h"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	tasks_details := []string{"Do the laundry", "Do the dishes"}

	tasks := h.List(tasks_details, func(item string, index int) *h.Element {
		return h.Tr(h.Td(h.Text(item)), h.Td(h.Button(h.Text("Save"))))
	})

	page := h.NewPage(
		RootPage(
			h.Table(
				h.Tr(
					h.Th(
						h.Text("Task"),
					),
					h.Th(
						h.Text("Save"),
					),
				),
				h.Tr(tasks),
			),
		),
	)
	return page
}
