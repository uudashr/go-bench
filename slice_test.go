package bench_test

import (
	"testing"
)

const paragraph = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean mollis, urna ac vestibulum commodo, justo arcu eleifend velit, sit amet accumsan erat lorem non ligula. Praesent tincidunt rutrum est, et posuere mauris dapibus ut. Nunc cursus purus ut posuere sagittis. Sed placerat lectus id magna rhoncus, quis iaculis felis faucibus. Pellentesque a ligula sem. Duis faucibus nunc nibh, in rutrum ex consequat rhoncus. Vestibulum auctor lacus a metus fringilla, id condimentum nibh tincidunt. Ut eu urna vel lectus consequat maximus. Maecenas efficitur imperdiet erat eget posuere."

func BenchmarkSlice_Append(b *testing.B) {
	var runes []rune
	for _, c := range paragraph {
		runes = append(runes, c)
	}
	_ = runes
}

func BenchmarkSlice_AppendAll(b *testing.B) {
	var runes []rune
	runes = append(runes, []rune(paragraph)...)
	_ = runes
}

func BenchmarkSlice_PreAllocated_Assign(b *testing.B) {
	runes := make([]rune, len(paragraph))
	for i, c := range paragraph {
		runes[i] = c
	}
	_ = runes
}

func BenchmarkSlice_PreAllocated_Append(b *testing.B) {
	runes := make([]rune, 0, len(paragraph))
	for _, c := range paragraph {
		runes = append(runes, c)
	}
	_ = runes
}
