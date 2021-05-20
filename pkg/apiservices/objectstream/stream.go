package objectstream

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cloudfunny/funio/pkg/apiservices/locate"
)

type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

type GetStream struct {
	reader io.Reader
}

func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		request, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		client := http.Client{}
		r, e := client.Do(request)
		if e != nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("data server return http code %d", r.StatusCode)
		}
		c <- e
	}()
	return &PutStream{writer, c}
}

func newGetStream(url string) (*GetStream, error) {
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("data server return http code %d", r.StatusCode)
	}
	return &GetStream{r.Body}, nil
}

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return NewGetStream(server, object)
}

func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, object)
	}
	return newGetStream("http://" + server + "/objects/" + object)
}

func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c
}

func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}
