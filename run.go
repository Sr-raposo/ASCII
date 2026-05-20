package main
import "fmt"
import "os"
import "os/exec"
import "path/filepath"
func main() {
	
	OsLocation, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	OsDir := filepath.Dir(OsLocation)
	AsciiPath := filepath.Join(OsDir, "ASCII.exe")
	exec.Command(
		"cmd",
		"/c",
		"start",
		"/max",
		"wt",
		"-p",
		"ASCII",
		AsciiPath,
	).Run()

}