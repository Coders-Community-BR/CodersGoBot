package util

import "math/rand"

// Código hexadecimal para as cores
var colors = [16]int{0x0066ff, 0x00c1ff, 0x00fff9, 0x00ff9c, 0x00ff2c, 0x8fff00, 0xf1ff00,
	0xffbb00, 0xff7000, 0xff3800, 0x00a31c, 0xc4390b, 0x6e01ff, 0x790cbc, 0xa600ff, 0xf100ff}

func GetRandomColor() int {
	return colors[rand.Intn(len(colors))]
}
