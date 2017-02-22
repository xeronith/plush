package vv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ContentForOf(t *testing.T) {
	r := require.New(t)
	input := `
	<%= contentFor("buttons") { %><button>hi</button><% } %>
	<b1><%= contentOf("buttons") %></b1>
	<b2><%= contentOf("buttons") %></b2>
	`
	s, err := Render(input, NewContext())
	r.NoError(err)
	r.Contains(s, "<b1><button>hi</button></b1>")
	r.Contains(s, "<b2><button>hi</button></b2>")
}
