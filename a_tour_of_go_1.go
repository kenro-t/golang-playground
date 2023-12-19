package main

// インポートは()で括って複数選択するか、 import "fmt"のように1行ずつすることも可能
import (
	"fmt"
	"math"
)

func main() {
	// 外部パッケージからエクスポートされたものを参照するとき、頭文字が大文字になる(piではなくPiになる)
	fmt.Println(math.Pi)
}