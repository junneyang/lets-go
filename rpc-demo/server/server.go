package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"sync"

	"example.com/rpc-demo/codec"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve(lis net.Listener) error {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc accept error: ", err)
			return err
		}
		go s.ServeConn(conn)
	}
}

func (s *Server) ServeConn(conn net.Conn) {
	defer conn.Close()

	opt := &codec.Option{}
	if err := json.NewDecoder(conn).Decode(opt); err != nil {
		log.Println("rpc server options error: ", err)
		return
	}
	if opt.MagicNumber != codec.MagicNumber {
		log.Println("rpc server invalid magic number: ", opt.MagicNumber)
		return
	}
	codecFunc := codec.CodeCMap[opt.CodecType]
	if codecFunc == nil {
		log.Printf("rpc server invalid codec type %s", opt.CodecType)
		return
	}
	s.ServeCodec(codecFunc(conn))
}

func (s *Server) ServeCodec(cc codec.Codec) {
	defer cc.Close()
	request, err := s.readRequest(cc)
	if err != nil {
		s.sendResponse(cc, request.header, nil)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go s.handleRequest(cc, request, wg)
	wg.Wait()
}

type Request struct {
	header *codec.Header
	args   reflect.Value
}

type Response struct {
	header *codec.Header
	body   reflect.Value
}

func (s *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	header := &codec.Header{}
	if err := cc.ReadHeader(header); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server read header error: ", err)
		}
		return nil, err
	}
	return header, nil
}

func (s *Server) readRequest(cc codec.Codec) (*Request, error) {
	header, err := s.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	request := &Request{header: header}
	request.args = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(request.args.Interface()); err != nil {
		log.Println("rpc server read args err: ", err)
	}
	return request, nil
}

func (s *Server) handleRequest(cc codec.Codec, request *Request, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("header=>%v, args=>%v", request.header, request.args.Elem())
	response := &Response{header: request.header}
	response.body = reflect.ValueOf(fmt.Sprintf("rpc resp %d", request.header.Seq))
	log.Printf("header=>%v, body=>%v", response.header, response.body)
	s.sendResponse(cc, response.header, response.body.Interface())
}

func (server *Server) sendResponse(cc codec.Codec, header *codec.Header, body interface{}) {
	if err := cc.Write(header, body); err != nil {
		log.Println("rpc server write response error: ", err)
	}
}
