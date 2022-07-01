package controllers

import (
	"io"
	"net/http"
)

type customResponseWriter struct {
	http.ResponseWriter
	io     io.Writer
	status int
	length int
}

func (c *customResponseWriter) Write(data []byte) (int, error) {
	c.length += len(data)
	return c.io.Write(data)
}

func (c *customResponseWriter) WriteHeader(status int) {
	c.ResponseWriter.WriteHeader(status)
	c.status = status
}
