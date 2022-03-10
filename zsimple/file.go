package zsimple

import (
	"fmt"
)

type File struct {
	Name string
}

func (f *File) Close() {
	fmt.Println("Close file", f.Name)
}

// provider ini kita tambabhkan return value berupa closure function yang jika dipanggil menjalankan fungsi close file
func NewFile(name string) (*File, func()){
	file := &File{Name: name} 
	return file,func(){
		file.Close()
	}
}