package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicTemplating(t *testing.T) {
	x, err := templatedSafe(
		"{{first}} is {{second}} templated.",
		"this",
		"basic",
	)

	assert.Nil(t, err)
	assert.Equal(t, "this is basic templated.", x)

	x, err = templatedSafe(
		"This is a regular string.",
	)

	assert.Nil(t, err)
	assert.Equal(t, "This is a regular string.", x)

	x, err = templatedSafe(
		"{{first}} is {{second}} templated.",
		"this",
	)

	assert.NotNil(t, err)
	assert.Equal(t, "", x)

	x, err = templatedSafe(
		"{{first}} is {{second}} templated.",
		"this",
		"basic",
		"too many",
	)

	assert.NotNil(t, err)
	assert.Equal(t, "", x)
}

func TestTemplatedWrapper(t *testing.T) {
	x := Templated(
		"{{first}} is {{second}} templated.",
		"this",
		"basic",
	)

	assert.Equal(t, "this is basic templated.", x)

	assert.Panics(t, func() {
		Templated(
			"{{first}} is {{second}} templated.",
			"this",
		)
	})

	assert.Panics(t, func() {
		Templated(
			"{{first}} is {{second}} templated.",
			"this",
			"basic",
			"additional",
		)
	})
}

func TestTemplatingWithReuseArguments(t *testing.T) {
	assert.Equal(t,
		"foo/foo/bar",
		Templated(
			"{{first}}/{{first}}/{{another}}",
			"foo",
			"bar",
		),
	)

	assert.Equal(t,
		"foo/bar/foo/bar/foo/foo",
		Templated(
			"{{first}}/{{another}}/{{first}}/{{another}}/{{first}}/{{first}}",
			"foo",
			"bar",
		),
	)
}

func TestTemplatingWithEmptyNames(t *testing.T) {
	assert.Equal(t,
		"foo/bar/baz",
		Templated(
			"{{first}}/{{}}/{{another}}",
			"foo",
			"bar",
			"baz",
		),
	)

	assert.Equal(t,
		"foo/bar/baz",
		Templated(
			"{{}}/{{}}/{{}}",
			"foo",
			"bar",
			"baz",
		),
	)
}

func TestTemplatingWithEscaping(t *testing.T) {
	assert.Equal(t,
		"{}/{{}}/{{{}}}",
		Templated(
			"{}/{{{}}}/{{{{}}}}",
		),
	)

	assert.Equal(t,
		"{name}/{{name}}/{{{name}}}",
		Templated(
			"{name}/{{{name}}}/{{{{name}}}}",
		),
	)

	assert.Equal(t,
		"{name}/foo/{{name}}/{{{name}}}",
		Templated(
			"{name}/{{name}}/{{{name}}}/{{{{name}}}}",
			"foo",
		),
	)

	assert.Equal(t,
		"foo/{first}/{{}}/{{another}}/bar/{{{escaped}}}/{{first}}/baz",
		Templated(
			"{{first}}/{first}/{{{}}}/{{{another}}}/{{}}/{{{{escaped}}}}/{{{first}}}/{{two}}",
			"foo",
			"bar",
			"baz",
		),
	)

	assert.Equal(t,
		"{{foo",
		Templated(
			"{{{{name}}",
			"foo",
		),
	)

	assert.Equal(t,
		"foo}}",
		Templated(
			"{{name}}}}",
			"foo",
		),
	)

	assert.Equal(t,
		"{{foo",
		Templated(
			"{{{{}}",
			"foo",
		),
	)

	assert.Equal(t,
		"{name}}/{{foo/{barname}}}/{{{name}}}",
		Templated(
			"{name}}/{{{{name}}/{{{}}name}}}/{{{{name}}}}",
			"foo",
			"bar",
		),
	)

}
