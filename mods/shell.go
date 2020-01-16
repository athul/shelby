package mods

var zsh = `
autoload -Uz
precmd() {
	prompt_symbol='❯'
	PROMPT=$'%(?.%{\e[92m%}.%{\e[91m%})${prompt_symbol}%f
}`

var bash = `
export SHELBY_SHELL="bash"
prompt_shelby_load() {
if [ $? != 0 ]; then
    local prompt_symbol="\[\e[0;91m\]❯\[\e[0m\]"
  else
    local prompt_symbol="\[\e[0;92m\]❯\[\e[0m\]"
  fi

  PS1="$(~/.local/bin/shelby)\n${prompt_symbol} " 
}
PROMPT_COMMAND=prompt_shelby_load
`

//UseShell checks for the Shell Nature and returns the Structure
func UseShell(shell string) string {
	shells := map[string]string{
		"zsh":  zsh,
		"bash": bash,
	}
	return shells[shell]
}

/*
import (
	"bytes"

	"io"
	"log"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/context"
	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileInitBash is "init.bash"
var FileInitBash = []byte("\x50\x52\x4f\x4d\x50\x54\x5f\x43\x4f\x4d\x4d\x41\x4e\x44\x3d\x62\x72\x6f\x6e\x7a\x65\x5f\x70\x72\x6f\x6d\x70\x74\x0a\x62\x72\x6f\x6e\x7a\x65\x5f\x70\x72\x6f\x6d\x70\x74\x28\x29\x20\x7b\x0a\x09\x50\x53\x31\x3d\x22\x24\x28\x63\x6f\x64\x65\x3d\x24\x3f\x20\x6a\x6f\x62\x73\x3d\x24\x28\x6a\x6f\x62\x73\x20\x2d\x70\x20\x7c\x20\x77\x63\x20\x2d\x6c\x29\x20\x62\x72\x6f\x6e\x7a\x65\x20\x70\x72\x69\x6e\x74\x20\x22\x24\x7b\x42\x52\x4f\x4e\x5a\x45\x5b\x40\x5d\x7d\x22\x29\x20\x22\x0a\x7d\x0a")

// FileInitFish is "init.fish"
var FileInitFish = []byte("\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x66\x69\x73\x68\x5f\x70\x72\x6f\x6d\x70\x74\x3b\x20\x65\x6e\x76\x20\x63\x6f\x64\x65\x3d\x24\x73\x74\x61\x74\x75\x73\x20\x6a\x6f\x62\x73\x3d\x28\x63\x6f\x75\x6e\x74\x20\x28\x6a\x6f\x62\x73\x20\x2d\x70\x29\x29\x20\x63\x6d\x64\x74\x69\x6d\x65\x3d\x7b\x24\x43\x4d\x44\x5f\x44\x55\x52\x41\x54\x49\x4f\x4e\x7d\x6d\x73\x20\x62\x72\x6f\x6e\x7a\x65\x20\x70\x72\x69\x6e\x74\x20\x24\x42\x52\x4f\x4e\x5a\x45\x3b\x20\x65\x63\x68\x6f\x20\x2d\x6e\x20\x27\x20\x27\x3b\x20\x65\x6e\x64\x0a")

// FileInitZsh is "init.zsh"
var FileInitZsh = []byte("\x42\x52\x4f\x4e\x5a\x45\x5f\x53\x54\x41\x52\x54\x3d\x24\x28\x64\x61\x74\x65\x20\x2b\x25\x73\x25\x33\x4e\x29\x0a\x75\x6e\x73\x65\x74\x6f\x70\x74\x20\x70\x72\x6f\x6d\x70\x74\x5f\x73\x75\x62\x73\x74\x0a\x0a\x70\x72\x65\x65\x78\x65\x63\x28\x29\x20\x7b\x0a\x09\x42\x52\x4f\x4e\x5a\x45\x5f\x53\x54\x41\x52\x54\x3d\x24\x28\x64\x61\x74\x65\x20\x2b\x25\x73\x25\x33\x4e\x29\x0a\x7d\x0a\x0a\x70\x72\x65\x63\x6d\x64\x28\x29\x20\x7b\x0a\x09\x50\x52\x4f\x4d\x50\x54\x3d\x22\x24\x28\x63\x6f\x64\x65\x3d\x24\x3f\x20\x6a\x6f\x62\x73\x3d\x24\x28\x6a\x6f\x62\x73\x20\x7c\x20\x77\x63\x20\x2d\x6c\x29\x20\x63\x6d\x64\x74\x69\x6d\x65\x3d\x24\x28\x28\x24\x28\x64\x61\x74\x65\x20\x2b\x25\x73\x25\x33\x4e\x29\x2d\x24\x42\x52\x4f\x4e\x5a\x45\x5f\x53\x54\x41\x52\x54\x29\x29\x6d\x73\x20\x62\x72\x6f\x6e\x7a\x65\x20\x70\x72\x69\x6e\x74\x20\x22\x24\x7b\x42\x52\x4f\x4e\x5a\x45\x5b\x40\x5d\x7d\x22\x29\x20\x22\x0a\x7d\x0a")

func init() {
	if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}

	var err error

	var f webdav.File

	f, err = FS.OpenFile(CTX, "init.bash", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(FileInitBash)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "init.fish", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(FileInitFish)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "init.zsh", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(FileInitZsh)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
} */
