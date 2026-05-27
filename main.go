package main
import "fmt"
import "github.com/disintegration/imaging"
import _ "golang.org/x/image/webp"
import "os"
import "strings"

var sb strings.Builder

func main() {
	imageSize := 200
	ASCIIChars := []string{
	" ", ".", "'", "`", "^", "\"", ",", ":",
	";", "!", "i", "l", "I", "?", "]", "[",
	"}", "{", "1", "(", ")", "|", "\\", "/",
	"t", "f", "j", "r", "x", "n", "u", "v",
	"c", "z", "X", "Y", "U", "J", "C", "L",
	"Q", "0", "O", "Z", "m", "w", "q", "p",
	"d", "b", "k", "h", "a", "o", "*", "#",
	"M", "W", "&", "8", "%", "B", "@", "$",
}
	Imagens, err := os.ReadDir("images")
	if err != nil {
			fmt.Println("Error opening image:", err)
			return
		}
	for image:= 0; image < len(Imagens); image ++{ 
		Path := fmt.Sprint("images/",Imagens[image].Name())
		img, err := imaging.Open(Path)
		if err != nil {
			fmt.Println("Error opening image:", err)
			return
		}
		// Resize da imagem
		alturaOriginal := img.Bounds().Dy()
		larguraOriginal := img.Bounds().Dx()
		novaAltura := float32(alturaOriginal) / float32(larguraOriginal) * float32(imageSize)
		novaAltura *= 0.55 // Ajuste para compensar a proporção dos caracteres ASCII
		img = imaging.Resize(img, imageSize, int(novaAltura), imaging.Box)

		// Converter para escala de cinza
		img = imaging.Grayscale(img) 

		// Imprimir os caracteres ASCII 
		for y:= 0; y <img.Bounds().Dy(); y ++{
			for x:= 0; x <img.Bounds().Dx(); x ++{
				r, g ,b, _ := img.At(x, y).RGBA()
				media := float32(((r >> 8) + (g >> 8) + (b >> 8)) / 3)  // media dos valores RGB max de 255
				ASCIIValue := int(media * float32(len(ASCIIChars)-1) / 255.0) // valores de 0 a Tamanho da lista de caracteres ASCII - 1
				ASCIIChar := ASCIIChars[ASCIIValue]
				sb.WriteString(ASCIIChar)

			}
			sb.WriteByte('\n')
		}
		fmt.Println(sb.String())
		sb.Reset()
	}
	fmt.Scanln()
}