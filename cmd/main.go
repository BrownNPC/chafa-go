package main

import (
	"fmt"

	"github.com/ploMP4/chafa-go"
)

func main() {
	fmt.Println(chafa.GetNActualThreads())
	c := chafa.CanvasNew(nil)
	_ = chafa.CanvasNewSimilar(c)

	fmt.Println(c.Refs)
	chafa.CanvasRef(c)
	fmt.Println(c.Refs)
	chafa.CanvasUnRef(c)
	fmt.Println(c.Refs)

	fmt.Println(chafa.CanvasPrint(c, nil))
}
