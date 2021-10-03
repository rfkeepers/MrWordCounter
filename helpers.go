package main

import (
	"fmt"
	"sort"
	"strings"
)

// variants of the asynchronous version
type version int

const (
	unknown version = iota
	vChanneled
	vSliced
)

// the most minimal concern cases
func runBasicSet() {
	run("empty string", "")
	run("minimal case", "hello world")
	run("one duplicate", "hello hello")
	run("multiple duplicates", "jabberwocky wabberjocky jabberwocky wabberjocky")
}

// a more complicated set with mixed case and non-alphanumeric characters
func runComplicatedSet() {
	run("complicated string", "the dev dev loves to dev code in dev the golang")
	run("mixed case", "The DEV dev loves to dEv code in dev the Golang")
	run("with punctuation", "!The DEV dev loves? to dev... code in dev. the@#$%^&*()_+=-/ Golang.")
}

// simple cases for the async runners
func runAsyncSet(ver version) {
	goRoutineCount := *routines
	runAsync(ver, goRoutineCount, "nil input", nil)
	runAsync(ver, goRoutineCount, "zero len input", []string{})
	runAsync(ver, goRoutineCount, "single input", []string{"!The DEV dev loves? to dev... code in dev. the@#$%^&*()_+=-/ Golang."})
	runAsync(ver, goRoutineCount, "duplicate inputs", []string{
		"!The DEV dev loves? to dev... code in dev. the@#$%^&*()_+=-/ Golang.",
		"!The DEV dev loves? to dev... code in dev. the@#$%^&*()_+=-/ Golang.",
		"!The DEV dev loves? to dev... code in dev. the@#$%^&*()_+=-/ Golang.",
		"!The DEV dev loves? to dev... code in dev. the@#$%^&*()_+=-/ Golang.",
		"!The DEV dev loves? to dev... code in dev. the@#$%^&*()_+=-/ Golang.",
	})
	runAsync(ver, goRoutineCount, "random inputs", []string{
		"El veloz murciélago hindú comía feliz cardillo y kiwi. La cigüeña tocaba el saxofón detrás del palenque de paja.",
		"Съешь же ещё этих мягких французских булок, да выпей чаю",
		"Zażółć gęślą jaźń",
		"استنكار  النشوة وتمجيد الألم نشأت بالفعل، وسأعرض لك التفاصيل لتكتشف حقيقة وأساس تلك",
		"누구든지 체포 또는 구속을 당한 때에는 즉시 변호인의 조력을 받을 권리를 가진다, 새로운 회계연도가 개시될 때까지 예산안이 의",
	})
}

func runLoremSet(ver version) {
	goRoutineCount := *routines
	runAsync(ver, goRoutineCount, "one huge string", []string{lorems(256)})
	runAsync(ver, goRoutineCount, "many large strings", []string{
		lorems(32),
		lorems(32),
		lorems(32),
		lorems(32),
		lorems(32),
	})
	thousandLorems := []string{}
	for i := 0; i < 1024; i++ {
		thousandLorems = append(thousandLorems, lorems(1))
	}
	runAsync(ver, goRoutineCount, "hundred lorem ipsums", thousandLorems)
}

func run(title, input string) {
	fmt.Printf("\n--- %s\nInput: %s\n", title, input)
	prettyPrint(mrWordCount(input))
}

var versionNames = map[version]string{
	vChanneled: "Channeled",
	vSliced:    "Sliced",
}

func runAsync(ver version, goRoutineCount int, title string, inputSet []string) {
	fmt.Printf("\n--- %s\nversion: %s\ngoroutines: %d\n", title, versionNames[ver], goRoutineCount)
	switch ver {
	case vChanneled:
		prettyPrint(mrChanneledCounter(goRoutineCount, inputSet))
	case vSliced:
		prettyPrint(mrSlicedCounter(goRoutineCount, inputSet))
	}
}

func prettyPrint(wordsAndCounts wordCounts) {
	arr := []string{}
	for w, c := range wordsAndCounts {
		arr = append(arr, spaceBetween(w, c))
	}
	sort.Strings(arr)
	fmt.Println(strings.Join(arr, "\n"))
}

func spaceBetween(word string, count int) string {
	tabs := "\t"
	wlen := len(word+":") / 4
	if wlen < 3 {
		tabs += "\t"
	}
	if wlen < 2 {
		tabs += "\t"
	}
	return fmt.Sprintf("%s:%s%d", word, tabs, count)
}

func lorems(count int) string {
	li := ""
	for i := 0; i < count; i++ {
		li += loremIpsum
	}
	return li
}

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis leo pellentesque, porta urna nec, consectetur nisi. Nullam vel tincidunt arcu. Nulla pellentesque nisi rhoncus turpis placerat interdum. Vivamus eleifend risus arcu, sagittis congue ex laoreet vel. Fusce eu posuere felis. Proin tincidunt ac elit ut aliquam. Sed auctor scelerisque lectus, tempor vehicula eros maximus in. Sed elementum eu lorem vitae lacinia.

Pellentesque massa lorem, blandit sed mauris sit amet, laoreet lobortis leo. Phasellus at tincidunt turpis. Interdum et malesuada fames ac ante ipsum primis in faucibus. Mauris eu justo justo. Integer posuere purus mi, sit amet semper ligula facilisis ut. Integer blandit ipsum vel neque euismod pellentesque. Praesent hendrerit diam quis pretium viverra. Aliquam erat volutpat.

Praesent quis felis quis neque accumsan fringilla sed ac ligula. Sed tempor sagittis nisl, at lobortis ipsum iaculis elementum. Suspendisse vitae viverra eros. Proin maximus sed lorem vel aliquet. Pellentesque vel accumsan elit. Integer orci ex, tempor vel arcu vel, vulputate commodo augue. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis et enim at quam ornare laoreet. Pellentesque luctus venenatis efficitur. Nulla a ultrices lorem.

Curabitur ac consequat justo, id scelerisque purus. Sed porta arcu faucibus, lobortis lorem ut, gravida nibh. Aliquam congue leo viverra feugiat eleifend. Nam non fringilla libero. Vestibulum in sagittis sapien. In euismod nec sapien at feugiat. Sed quis orci eu urna iaculis rhoncus a eget massa. Fusce hendrerit odio ac lorem vehicula placerat. Nulla placerat eget turpis vitae condimentum. Morbi bibendum, neque maximus convallis tempus, ante ante finibus diam, ornare pretium neque sapien sed risus. In hendrerit molestie dolor ac maximus. Pellentesque varius metus arcu, id sollicitudin est volutpat eu. Praesent laoreet ex vitae malesuada scelerisque. Praesent aliquet maximus pretium.

Nulla dictum dapibus nunc, vel sodales magna vulputate quis. Cras velit nisi, iaculis id tempor vitae, sagittis vel orci. Nunc ullamcorper ipsum ut sollicitudin interdum. Aenean auctor, justo id ultrices fringilla, ante neque consequat eros, at sagittis tellus leo id odio. Etiam vitae elit a justo lobortis elementum posuere sit amet nibh. Suspendisse non ex at mi feugiat convallis. Cras eu orci sit amet purus placerat dapibus. Ut in semper velit, in egestas augue. Etiam consequat leo nec arcu auctor dictum quis vel ante. In feugiat arcu metus, in tristique dui efficitur porttitor. Etiam eu iaculis mi. Phasellus vel consequat elit. Nulla dictum quam eu leo gravida, non consectetur arcu semper.`
