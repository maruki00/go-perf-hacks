package zerocopy

import (
	"io"

	"golang.org/x/exp/mmap"
)


//in this example we reuse buf each time. 
func StreamData(src io.Reader, dst io.Writer) error {
	buf := make([]byte, 4096) 
	_, err := io.CopyBuffer(dst, src, buf)
	return err
}


//her we return a slice without copying it.
func process(buffer []byte) []byte {
	return buffer[128:256] 
}


//here we access the file directly
func ReadFileZeroCopy(path string) ([]byte, error) {
	r, err := mmap.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	data := make([]byte, r.Len())
	_, err = r.ReadAt(data, 0)
	return data, err
}
