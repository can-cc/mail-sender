package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/mailgun/mailgun-go"
	"io"
	"log"
	"os"
)

const (
	chunksize int = 1024
)

func openFile(name string) (byteCount int, buffer *bytes.Buffer) {
	var (
		data  *os.File
		part  []byte
		err   error
		count int
	)

	data, err = os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunksize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		log.Fatal("Error Reading ", name, ": ", err)
	} else {
		err = nil
	}

	byteCount = buffer.Len()
	return
}

func main() {
	box := packr.NewBox("./")
	s, _ := box.FindString("help.txt")

	flag.Usage = func() {
		fmt.Println(s)
	}

	mailgunApiKey := flag.String("mailgun-api-key", "", "mailgun api")
	domain := flag.String("domain", "", "mailgun domain")
	noReply := flag.String("noreply", "", "mailgun noreply")
	recipient := flag.String("recipient", "", "recipient address")
	attachmentPath := flag.String("attachment-path", "", "attachment-path")
	subject := flag.String("subject", "", "subject")
	body := flag.String("body", "", "body")
	flag.Parse()

	if len(*mailgunApiKey) <= 0 {
		flag.Usage()
		return
	}
	if len(*domain) <= 0 {
		flag.Usage()
		return
	}
	if len(*noReply) <= 0 {
		flag.Usage()
		return
	}
	if len(*recipient) <= 0 {
		flag.Usage()
		return
	}
	if len(*subject) <= 0 {
		flag.Usage()
		return
	}
	if len(*body) <= 0 {
		flag.Usage()
		return
	}

	fmt.Println()

	gun := mailgun.NewMailgun(*domain, *mailgunApiKey)

	// _, attachmentBuffer := openFile(*attachmentPath)
	// rc := ioutil.NopCloser((attachmentBuffer))

	m := mailgun.NewMessage(*noReply, *subject, *body, *recipient)
	// m.AddReaderAttachment(*attachmentPath, rc)
	m.AddAttachment(*attachmentPath)
	// m.AddHeader("From", *recipientName)

	response, id, _ := gun.Send(m)
	fmt.Printf("Response ID: %s\n", id)
	fmt.Printf("Message from server: %s\n", response)

	fmt.Println()
}
