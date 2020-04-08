# Color Thief (Golang)

Grab the color palette from an image.

## Download

```
go get -u github.com/cascax/colorthief-go
```

## Usage

```go
import "github.com/cascax/colorthief-go"

colors, err := colorthief.GetPaletteFromFile("img/image-1.png", 6)

baseColor, err := colorthief.GetColorFromFile("img/image-1.png")
```
