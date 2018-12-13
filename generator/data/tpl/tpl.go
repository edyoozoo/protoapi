// Code generated by "esc -o generator/data/tpl/tpl.go -modtime 0 -pkg=tpl generator/template"; DO NOT EDIT.

package tpl

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return []os.FileInfo(fis[0:limit]), nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/generator/template/echo_enum.gogo": {
		name:    "echo_enum.gogo",
		local:   "generator/template/echo_enum.gogo",
		size:    491,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3SQMWvDMBSEZ79fcYQMEjTOnuKpaSFLUmjoEjIo9qsxtSUjy0MQ+u9FtuNQSLZD3H36
eOs13kzBKFmzVY4LXK5orXFGtdUrtgfsD0e8b3fHlKhV+a8qGd6nn2MMgchd2+FprxoOAZV2RLnRnYOg
xPsVrNIlI/2ouC46hBBfb23vl7eYRci3qnseKyuwLkIgSfTT6xwij6LzVOLL2UqXQqIbAjwlWjXcYZOh
Ue1prp7Hgqfkic9daIPFnBcv02D0SAJRYtn1VmP45xSFzhSe+cXDChkPEtWmpai0k0NVUqCHOv9xy5m3
62a4kLgYU8MTAEzoYZBld4XpA9ZFxP4FAAD//9bZDJPrAQAA
`,
	},

	"/generator/template/echo_service.gogo": {
		name:    "echo_service.gogo",
		local:   "generator/template/echo_service.gogo",
		size:    2781,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7xWb2/bthN+LX2Kq3/9taRjy4q3IksCI8tiu8uA/MFi7MXCNJOls01EJg2KSp3K7Gcf
TrI9O033L8HeCNTd8bnnHh4PbLXgRCcIY1RoIosJDB9gZrTV0UweQvcCzi8G0OueDgLfn0XxXTRGKIrg
slo65/tyOtPGAvO9msExzmc13/dqY2kn+TCI9bSVRsPMRvFdC+OJrm37HrT+pHVrlXG9GOuaz30/1ior
oXvTSKaQWSPVGDpQ+8AYY9dR89Nx89ebhRDJ4vrV/4R4/f83b4WoC7EjRFOIVkeIIyE+3P5WCLFwn28W
10LMizA8Dl2TVt29ft/Ror8fLk397snS1O2vTP1e393wHSZE8J8n5XW+YEyIebvNqWhahQsh5uE+r9Nf
mNAn4vxoy7fDjxijzOEu4YXf0WdIn5g+SMbdkRDzvRHVMm/vljzb35Dj3bBi/S6hvz18Xg0LJsQmlf1t
KlWO0fNycM7r/1CdSlPOv/+ilZ5V68thbSM1F0IEi9vF52dh1l+OHudCBHxno+CXEu4lRXtBwZ4t1tHr
cqrdR4ZmmplXU60D1dwMzvLMnujpTKbIShen6FaLBu55NEXnQGZgJwhSWTSjKEaItbKRVBlEaVq6yGB0
mqLJfPsww83N612F7xVFE0ykxgjBGdqJTjJwjszBQNoUnWM0roMTrSzObQPqRRGcqlluBw8zdI4DI8tF
btemopAjUAhBzxhtyAa1mnPV1rWN4lAlzvGKA6qEEjvff5rRKFcx3K6LuP0xUkmKhmXmHori9dLMoWS7
dPZpT+F7Bm1uFBAEi2GzHg4MjQEkVpxCPangoAMKP7JHhfrkHFEodCAOfpAqYVLxw9LyqgNKpiWAZzCb
EcaJnk61KgsuKLpcHcCb9bo4wyyLxnhAEJUyjJP43opxHPx0dXHOvm2HDSBY7nueWxK5j9KeMZRIquCX
KJVJZJHxw5XjryittixpVbv+VnKd2yfPGOiQSymXR0spM3MfbDRT3ACpCIlO+WmMtc76cRFbtMJwmazi
5a26yJUct4PbFKxzy32PWmyj3+ha/YxjmVk0W9crzzABq2EoVQJG55YuUtmEX4QzhHrZVL14ohtQdeS6
IQvfa7Ug+yhtPCFAerzEFojX8eUpNQOaBkmRZ/SoIaC3GXRxFOWprdw+CXLbAH1HimJQWQNWZd0K5YcU
RXqtwqDq5j9eVMFW6lKSPxkD61O6QnMvY6zO6fLialCdFQbvewNWK9+CduJcrfGVW8q3L/rXsd/31tCU
5l9gPxoovwcAAP//F5mGs90KAAA=
`,
	},

	"/generator/template/echo_struct.gogo": {
		name:    "echo_struct.gogo",
		local:   "generator/template/echo_struct.gogo",
		size:    924,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7SSUY+TQBDHn7ufYiSmgeaO+oypyeWghqTX01r7YszdWqZ1lS7csDXWzX53M1s4qNFH
eSDDf2bn/5tZplO4rQqEPWokabCALyeoqTKVrNVrSO9heb+GLM3XsRC13H6XewRr43fn0DkhplMWbkvZ
NEt5YMmcavxDg8bQcWvAipG110BS7xHiucKyaMA5VuO1MiWXcniqOXr81lQ6CayNz12Cx/Nx1AUfcuKv
vcTuqLcQEkysfTmEiOAtmt4ojAZWVgAAqB0QzGagVdkq/PyQBL+Qqo0s+xPPWUJzJN0VePmcbBMU95Yt
csv/DHq5qwg2slSFNBhGMOnijKgi3h8SNZDM4NPniZ/ZJ6z792Kveay467PCp6Mi9AAjnneAx6MHAZuM
kC3mebZIH1bZ+4/5KktZZe8ZyLpGXYT8dQXjAYUPeYrkou8V+DQvLoExukiM3MVNdpT41IPOKzpIAwEe
pCqDDvcF/cxYiO+k2X79YEjpfTj0inr8fLm5WeTpQ3Z3ky/+N30Xqx2UqH33CN7AK4/T/grji8u0/t0k
wLXON23rtCqFE78DAAD//3AqAjycAwAA
`,
	},

	"/generator/template/go/enum.gogo": {
		name:    "enum.gogo",
		local:   "generator/template/go/enum.gogo",
		size:    490,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3SQQWv6QBTEz3mfYhAPu/BX7/7JqbbgRQuVXsTDmryGUN0Nm81BlvfdyyYxUqi34TEz
78esVnhxJaNiy94ELnG+ofEuONPU/7HZY7c/4HWzPSyJGlN8m4oR4/J9kCJE4db0p525sghqG4gKZ9sA
RVmMC3hjK8byreZL2UIkXe/uGOd3maeST3PpeLAswLYUIU301dkCqkigU1TjI/jaVkqj7QUiZdZcucU6
x9U0x8l6GgyRsic8D6A1ZpOe/RsDA0cmRJnn0HmL/s8xAZ1InvGlYZVOgyS0MalqG3Rv1ST0J87vuvnU
t22ncqVxdu6CSAAwVveBPH8gjA/YlhD5CQAA//+hO0wK6gEAAA==
`,
	},

	"/generator/template/go/service.gogo": {
		name:    "service.gogo",
		local:   "generator/template/go/service.gogo",
		size:    2283,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6RWUW/jNgx+tn4FFxw2u8vJxWF7SZGHrc3uOmBNsQbbY6HYTCzUkVxJbq5n6L8PlBM7
TpNuwB4CGORH8eNHikqawrXOEdao0AiHOSxfoTLaaVHJyVpfwc0c7uYLmN3cLjhjlciexBqhafh9++k9
a5oPFiZT4N4zuam0cRCzaLSWrqiXPNObtBRL60T2lGJW6NHQ96r1N63TfdLuY61HLGEsTSnZndig9yAt
uAJBKodmJTKETCsnpLIgyjK4yGB0WaKxzL1WeBjcRTUsapqPYIRaI/A/0BU6t+A9IztfSFei93EGxJZf
a+XwqxuDwWe4aBp+q6raLV4r9D6B2KCtgnleu87eNHIFCoHPjNGGbDAaeT+GpfwWTCGic1IAqpwAaAz9
tElY1Ja+pzMk85bIf+ZwOnXSaoIqJyE8Y6cVWtUqg8dO1McvQuUlmtiaF2iaDztz0kq3c/5GMQ2LDLra
KKAjjsRNIO4LJ2hEYk+moHAbH1XKWBTJVRBqChn/Vao8NvicXAXTd1NQsgxHREGBD5Z/EfZabzZaETZU
7j35Q+smU+JteYvYeZsOOYHvu+8GjWmVi5P9CaGijP/+ML+Lf/p0SVNiq6TNjqXFU7ifLy9Do/cwkp9F
kWfpBTvB+i9Rylw47JnLFbyIcmYMsTf4zPeQOLnaew6FeKfQweGTXfC/1eZZz/siZW2Ck/MGh0M/GPLJ
FKx54Yf3LVyxpG/vu7086GOaAk6mFMLpFhzXmMBWliVUQskM5AqEtWic1ApWQpZj2BYyK2i1KO1gWwgH
W4StUI7OJipj0E/wXoIrAgSap1TD0Gc/bHYPfHBGqnU3FPsBa3UOQn+E08q2SnUrZdjxAQ06e4/rGxhu
e9h6R/hPl32zaRccLAbax3/iWlqHZrCXa4s5OA1LqXIwuna0gcO2eAOPES7C7Z9lhR5Duzq6zdGw6E3E
39IV9wZX8muMIWAMo1HC/Ek2PfgcL9hKV0BWW6c3UAXoGaqHic+THu8OARt6GWpIU7Bb6bKCkpM9c0Di
/nJ/SwsFzZi6Wlup1mEV/mDhBleiLl3rZtTbx270eGvlcUtiAO3nbw+DdnP2TykfpA5tPfcGHg7cA5oX
mWE7cvfzh0U7dsg/zxbxruYfYRT+DbjC+9H4zNuQDJ+X80k+z7oclO//JDl6z/4JAAD//0nOB8brCAAA
`,
	},

	"/generator/template/go/struct.gogo": {
		name:    "struct.gogo",
		local:   "generator/template/go/struct.gogo",
		size:    1055,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7RTT2/aThA9408xP+snZKPE9OyKShGYyhIhLaFcqirZ4sHd1uw646UqXe13r2axg6nS
3soBjd/8eW/fzo7HMNUFQokKSRgs4PMRatJGi1qmpX4NsztY3q0hm+XrJAhqsf0mSgRrk3en0LkgGI8Z
mFaiaZZiz5A51vgbBo2hw9aADQbWXgMJVSIkc4lV0YBzjCZraSou5fBYc/T4tdEqDa1NTlPCx1M7qoKb
XPDirGB3UFuICEbW/t8XEcNbNGeiKO5R2QAAQO6AYDIBJasW4d93QfATSW9Ede54zhKaA6muwMOnZJug
5EzZSm71B9bKHSQbUclCGFzh00ESFs49n+DSxBi60iiGURdnRJrYWCRqIJ3Ax08jb4ZPWPdnx6/hJXqf
ZCN6utmTMGSSATLFPM8Ws4dV9v5DvspmjDL3BERdoyoi/rqCYU+FD/kU6cXcK/BpdjSFIbo4GLiLK+5U
4tNZ6FzTXhgIcS9kFXZy/6MfGQPJrTDbL/eGpCqjPld8lp8vNzeLfPaQ3d7ki3+tvovlDipUfnoMb+CV
l9PuyPDiMq3/b1LgWueHtnVKVn6FUPGWnNYnb6Z6v9fKN937Z/aXDfJVUczvUaqSNbSjQ58Je+N/BQAA
//+U2FP6HwQAAA==
`,
	},

	"/generator/template/go_client.gogo": {
		name:    "go_client.gogo",
		local:   "generator/template/go_client.gogo",
		size:    2379,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6yVX2/bNhfGr8VPcV7BeCGurlwU240BXzSJNwRYkyJxdlMUKC0d2VwlUiOpZo7A7z7w
jyTbWbeg2B1N8hz+znMeHy0WsNlzDVwDg4rXCDsUqJjBErYHaJU0krUcsq+oNJci77qnLi9ksxiOKFks
4JcxiJkl9D3kG94gWOsOr27h5nYD66vrTU5Iy4ovbIfQ9/mHsLSWEN60UhnISJJuDwZ1SpIURSFLLnaL
37UUfkMpqfwRlwsuO8Nr90OgWeyNaVNCCTGH1uU23NQI+Q1r0FrQRnWFgZ4AALCWP9z96va42BFLSNWJ
ArIWfjgLo3CP5p2/nXWqjhEUepK0ecyygk7VLkl8N7+UzVqpZw8nff8aFBM7hFkFyxUMF3/mWJfaKeXY
zsH73uedVdbCZyfDMu37eJh+jiGvAUVp7VQKulJOSSisnXgZjVU4JIWmUwLSQjaNFODVTV2WiTV/j1qz
HXrAf5f2WZH/aXVuzSuQCjKuL/jTWqmQirqNUG7csPZYirOufluJLX+aZJiePl1FZdaia3TcmqHL7SsO
jwxahQNrgQtDSCGF9h4/SjMplIy1HweunKl+Y3WH4cpAQodmF7L0xhs966s6KU+wBrWja1j7cbz6KVzo
SfINngloCUeNmceAwJFYMurn3/nogD5NZjznu5QlZtQJcqR8xoWh/iolJypPOCfZBnkoXOsxd0ZhK2Ud
vRgz+4DVaiKYOgv2bDl63uxlqcc3/WSYjQWc+SlT+McVM8z/565F25nNofUXM4V6PLntzHg0dyYLRvPT
xFn/3qiwvVyB+52/Z0rvWT2kpyThlb/wvxUI7qpMRt157WNDM9yoWq5gHFGvfPMe7q6tTZ3genzHDc38
g9TGjbc5pKxta14ww6UIM3cOfhjnN/h40VUVqiyi0pfiJCVWqEChzi9kecgva6kxoyQUfXEwONKEgZ7f
ISvf1a7uEPLil/QjN8XeP3VvmOn0pfcKSQqmEd6+ebP0gaElyxX8/6wrvbO8eyc24EE0sQUTagynJPkb
qOdUDmvYjaFzd+wN6rF+DFjbMM0CFRcl/gn5besaoSF89tKXAYZE38Pnd0P4Ed5bj1eE2RrwZqdflpdQ
hfjvpgrhE9VPQy9PAqXyRs1SLgwqwWrQqL5i/KPBElJ4FYfiiOZ8XGLFutr8U8ZOfBHyUYD2vvIzJaXO
dMffhr8CAAD//1ut0G9LCQAA
`,
	},

	"/generator/template/php_client.gophp": {
		name:    "php_client.gophp",
		local:   "generator/template/php_client.gophp",
		size:    3889,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RWbW/qNhT+nl9xhpCAiqa924ddldGpt2VbpbZc3dFJ0+0VMuFQvCV2Gju3pZH/+2Qn
MSFxgOVDVezz8pzH5+2XX+N17J2dwWxNBVABBFY0RHhGhgmRuITFBuKES05iCv3vmAjKmZ+m76kf8Ois
vBpoG79bJSIvIMvAn9EIQSl9eTOFh+kMJje3M9/zGIlQxCRAyDL/gUT4p/6h1MjzUoHwN+fvnD991tav
Yjoyh/eb6zuyEE8Tlkbmz8jzzk5O4B6FIM8o4OTkzMuyU0gIe0bwy3OlvCAkQkCWSSpDBONQKSNLV0DF
J/o+SZLyHPBNIlsKKN0/5fc8mbwFGEvKmVHFUGCuf82jvfrXPIo4c5lgS6WARnGIETJZUSnAe5kHAFCJ
6jeK4VJoUvWF5h8DTXm3IFJzWKoY814umC5CGsAqZYH2DpRR2SdJQjbQTVDEnAkc5Irm716v+qMr6FMh
UPat/teOxdD5NhhULJXW6ArwBfw7ssAQOndXnyZ38y+Tz5Or2eSmUzGuv65cU3F6aU3CGAze/mC0I7fi
CZJgDS04gIgKN3VQFWBUTBf/YCDBvyGSzDYxQg1RjiqKYQwMX7fpVMorVcdWapxeGr4rONrkvpOQLolE
t6VdRr5+g7FRGzlj0vnpjMBlpZY9DVsmk6qnyjvC39HUOh77SIrrmiXV7qQ8pNtGf0t8Dthuxw5rDUab
p8pz37XU9Bb8/y3kH4pKroXTqGG5TvireRnbqPKmH9q+1u/0rH5PTxTGJeAbFbJT4bQ9svzM3exq8QqU
c+urn2XOTFOq1vYhy4y/SjHW6HK9aq0+Wl7geQdR3W6CMk1Yw/xoJ/ItF3Xjks+LBui2m18e9+h7i7Mm
V8ljGF86imYLbHhM0RywNzxQFM2TQcmgHun5ld4L9IrQWArMYdtGYCe3ljo0egPOhIRqkmSZ/xcJU8cI
3gIr/dodpD7G11LG1yFFJkfuHJvPjeckDWS/uyACHxMKY+h9+PFn/9w/9z9cfDz/eN5rSeqt+aLF2kL+
w970d16gllbl19Ou52lCe+YRCyDDppykEfJUGrGfzncFBs03LInrBnppilAvTXAxBj9fsopHULVGcY9y
ze1xnbJqk/BvWZzKPM11s35pELUmbBliAuOtge04GUJ3Qd8NqmGBkdW7pOmoGMVys9UbuDYPfWsnnT9N
pUXmHP8Jivp0axXbt0eUfShBsXup7E5r8ZfBuvEvisXYBkHZEt/An5pRIKBjdDst8ZTaZVDWV1M2HzxW
4yDs4mHc616eMzsJdupYlCobbm5u/35rw2ouMQc2xF2tkg2XT7duQc6eFa5gqGhC2R4rB2b7I/uX8VcG
OTjIyZObGC+g4+/Lyn3L41GeHzgsiSRF8uLS31knil7ZnLLbfnd6GZAwvIqprp6XoRlDuufdy6VSnfz3
45db83/ZA2xXKphT3n8BAAD//3TiBMQxDwAA
`,
	},

	"/generator/template/spring_service.gojava": {
		name:    "spring_service.gojava",
		local:   "generator/template/spring_service.gojava",
		size:    856,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6yRTY7iMBCF9zlFKauwGHOAbBADQiwgaMgFKkkRLKDssSuNUOS7t8JfsuhuqRFZRX5+
z++rGo/hr6kIamJyKFRBcQHrjBi0OoVZBussh/lsmasoslgesCZoW7W5/YaQRpE+WeMEjKuVt05zvXN4
orNxB3WmQhWaK4XMRlC0YbUgWaG1muv0t9aN8S97/5G3hj1NTXV5wfy/IS83b2Sb4qhLwMKLw1KgPKL3
3VTWeKIQpugJ2ggAoG3/gEOuCdSKZG8qDyE8Fb0DJlBbch+6pPxiCeLFPI8fdyYD3iS+Tl32IcSjmzok
up7ce7WtyhqxjXSJIfTFurhkMmDppCX3NzWP7sW7z5E0jnt7onmUXtUegbjq2v6ItMm2A6Z+/e9AWpAk
b4V4LvXbJ794L71HPZJC9BkAAP//3o5aN1gDAAA=
`,
	},

	"/generator/template/spring_struct.gojava": {
		name:    "spring_struct.gojava",
		local:   "generator/template/spring_struct.gojava",
		size:    585,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5SPwWoCMRCG73mKOdqD8QGWQkFbqBT1sC8w7o7baDYJk1mphLx7ibqtUmhpTjOB+f7v
n81g7luCjhwxCrWwPUFgLx6DqWCxhtW6hufFa62VCtgcsCNISW8uY86VUqYPngUa3+sdRiH+6K3eY3OI
3ml0zguK8U4vo3dzJhTP1b+ONuwDsZy+s/Z4RD2IsfrNRKmUCsPWmgYaizEWv3kZVthTzpAUAEBKU2B0
HYF+MWTbmPP5P7A5ohDsjENbTpd4xPoUzpcJdIFA6TlCyLWQszrvTzedLrSLx73BpKzeCQ+NbJCxz/nh
D6vy5N1EfaMAj78KlfWqldIdE6Zj1y+7m5IdScHWRmzhTka18phkYPczNl9jxuysPgMAAP//ioLHmkkC
AAA=
`,
	},

	"/generator/template/ts/helper.gots": {
		name:    "helper.gots",
		local:   "generator/template/ts/helper.gots",
		size:    3553,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xXT2/bRha/81O8FXZDSVYobzYBAjmM4zhO7IVjeWX7soaxGItP1sTUkDszlK1VBGzR
pmmLps0h7aE9tEBRIEUOyakoEgT9Mnacj1G8GZKibDk1D4L53u/93t95Q9erVacKm12uoMNDBK5gDwVK
pjGA3QG4sYx0xGLuGhhaVDsSmnGhoIthjBI6iWhrHgkFuss0HERJGMAuQqIwAC4KlCzmhqIGiUpYGA5I
7Xp1FnNPKxciSW+Hh4cbKPu8jSRMPY9J2lFgIj2QXGsURLE5iHGjLXmsnSpcvvjjVOH9y5/fffv46M2v
J89+ePfZ06PXX2U5O1WwmuMvHx0/fXH02/+PXn9//Pjtu29eHT/5+viLX6zFyXefHH/+5OTFy/evPj55
9nxhfWVs+Omjozc/nfz40dHb30+ePS9GWXeGQ94B7x4Kb5mpxajXi8SSlJEcjRzeiyOpYQg9Fhc0ZA8j
6MioR2UaDr3FkCm1xno4GjV3HyjXGQ5RBKORQ42FKtzBDhcYwLLWMSxS3TqRBIkqjoRC6DIRhFzsOVCt
O3honKJIetDVOjbwoQMAcGfp7sLW6ib4MFszgrVm6/7CKvhwZTaV3F7593+WWq1mC3y4mgkXm/fvN9fG
8iupfGVtc6m1trCaa66RRRY1BX4rZpL1YJiFOhoHrbsISOXIRcXws1m0kOU0wTJKWWnAuox6XOENgX2U
N9PkeIe0Xs7v+z4kIrCVq6QgenRXRgfEO2dEI/NrftQB1+0uTPB4SjOdqCJBm1HN09J6ecUaOYAeiTqR
IgvVk/gA23qSOWCaVT40P9MdFrsx6TNEnfpdkhL8KWM3JYC5C4Sdk1byybSlcyYKOsk854ycs6PQZyEw
UFraeYVbllsBE4CCdkKQau0WajNBO4jFMYoAdGR2TyLD6cNiCMp9FjZSkkr2R9q/ND8L3GqtLEa9OBIo
NBlVPGdchzhkbSzX/3Z1tr7Ha+Decqeq/7Fg1Y3p6itX63s1cP96jnbRGtfOUc8a45np2mu3rfH2Oeo7
Vr3jVgqdgNsJDwNgsNVapZvB1pXqY9qjqMLmZIqg0LREhtAw8l0aRXqlBUQCZTd8AZwSWbzE/yaoNES7
NEgeoLfnGdtlDMOolWrzq4ZMLBSiDmhalEVgcWAakE8OdJIwhC0ZZr2+bFzcW9qkFPdxUO+zMEGIGZfK
cOAh68UhNuiFUiJbH0p0xhr1ehi1WdiNlG5cn70+WyIQk3vgw1CwHjagdIBi7wB5qQaCt/etQHMmSiPC
ZnH5N2Eq4Tyx+DnJJeLwc4qpg53dmv9KUA62ZHhj82Y5kfmU19KiN2Dz9MDzDpT/YrXFHZaeg0SGc8Wj
TBskZlKrjGV7B3zY3plzrJ7KWibQPpo7/ywxKfssHG8VCqBpu29uZGqq12WqeSDWZRSj1IPyPg4qRRJ6
aFH4qYPtfRzsjCnTYDN2g/R9EDQFDx+asYk6kInd/BZwT/tIq+C6U7lNnjXiUWN9vU4+uQImJRucjsPT
0YapW7lifW+n07xA6J0zAeyDb0o5A+72TjEKwFDhdPDpKinqUJ+F0wtEAK8TySXW7pb7NJOTpHk+AdM4
oTEpnZvQXR7i2XyMR/CB7FY2mpnp3BlQ7tbSTejT5Ml/1krr22LPdfrPjeaaZ8eWdwbl/im3o4k3M+Ve
nKhuOb019ivUBt+FmfweqVTGJa0UPxdoNBRKzkL+PwzW7cbzU9IHERdl9xLt3fwIngYXc6B1OuMDHWiP
iwAPm52yO+/agl/+O8yDO+9CA4gSZs74LcY1PtZTL1/ydMEtnkroo7SIANpdBVi2oXJcvrJyYLprL7xq
i559KJn9n/4fYfQTLn0obbCBwZQusnvrRTovM/3gzp22bgsxjoXFwP7k+8O0HEr1ku1nnu8MlDySFano
9v4jAAD//+sjJp7hDQAA
`,
	},

	"/generator/template/ts/objs.gots": {
		name:    "objs.gots",
		local:   "generator/template/ts/objs.gots",
		size:    1210,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xSz2sTTRi+z1/xkEOTLl+Te+GDr5+NVShatHopRSabd5Mxm9l1ZtYahgHBihasFGw9
9KAnoVBoexK0/jlN0j9DZtMkG1s89CWE2Xmf55nn/VELAhZgvS00IhEThEaLJCluqIlGD+VUJSbhqSjP
wMJEGi6kBo9jmDahyQ2HNioLTaYIDRKyhUxTE0LmgKmq0dCkXoqQNAuwcJtgAS5Pvw0+v7s4/z7c/zp4
v3fx8+PYKQswyvQ/vO3vHS+t3e/v7A6PTy/P3gz3jwY7r/u/Dob7R8PD7cHB2WD3ZHj+afBlu39yePFj
hwU1xmo1kMy6mlm7AMVli1Ct+ws4x+hVmiiTA2Bt9QHvknOwDAAK+LuC4mZOGCXGwH/9+SmPM3LunwmJ
ZNNDHbP26lirjVpqeinNGFnmhq/7y4IZIQ2piId0C0eLsBZGe8mpuHPWigj0AtVV3qAYpdWl/+urzx7V
1+pL6/XlknMbm7nXidy1IhacYyyXqa6QrN7j+k7S7SayrlSifIVBwOB/+C/lindhi3mE4498e8hfImk8
p9AwBLVx5VEmQyMSiS5PC3RfQ2Wi4EvMPcwYmF9ERRslZOt69nHWyHvs3PxVI6NEoRKTQYd6fqUn4mOA
DxFh+mq1zfXDLbmmkpSU6VU61JvH3NyUudGh3maR7kNvCRO2MULb4uC9xRUyI5eU25wdaTFCrglla5FP
Gc6VF69hfCgymZJ/eALXfikm2+AfuBrwDU81KeJZbP6qX3oiOzLZQu66NIN0jE2P0/+bicXd+h0AAP//
LkxTtboEAAA=
`,
	},

	"/generator/template/ts/service_axios.gots": {
		name:    "service_axios.gots",
		local:   "generator/template/ts/service_axios.gots",
		size:    2142,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xVXY8URRR9pn/FzYRkPjL0sCYSnHFIVkVdE9mNi0/Gh9ru2zOFPVVNVc3CpK0ExAVJ
FtiEDVF3FYxiiHysJorswsqf6Z4ZnvgLprp6PpbRxFhP3XXPPffcc2/P1CoVpwKn21RCQEMEKqGFDAVR
6MNKD4qR4IqTiBYzGFqUx5kilEkIBGcKmQ/zSwvgcR9BtYmCc1x8BueoaoNqI4R0RRDRgyI5T7ksVs2l
wIALrAJVRQkCz3apQN8m5zAjhTKpSBiiD5RlVJHgZ9BTuZaJ0qw0lXBOUKWQGfjpXoTLnqDRCJ1hIsFX
qY8SCKwQST3oStJCCLiwLZAwBMJ86JAeMEQfiH+mK1UHmQLieVz4lLVAcZARejSg3kjSqAmLZD5IqrpE
Uc6cChz578epwHDnbv/WleTp48Hm7f5XG8ne9dEQnArYSLq+lm7cT69eG9zfGf56abB5b35pYfDtl8nT
Hwd3Lr58tp7uPk72nw827w0ePEieXO3f2k33bmbGvny2DunWnf7Dn15sXxj+fDF5/t1w52IWSh9+nW7f
S/auv/hhd7D1KHnycFLw8prltqwHzB3u3LVKc+iNX/o3NpK/tqy8+aUFqzC9vdfffjBR+Oj7F9+s9S+t
pZf/SG/sDC/tWz39O7v9a4/StT+T/ZtWh302od++SJ5cSzfWhxfWk/3t9Mpef+v3/uauU6k5tBNxoSDr
owoxzJuHJcE7VCJos6id0WY1RuDYAQCIY0FYC+Gw6kVYhcMrnIdQb0KphWohA75DFDEdS3Df7TLPDFWW
tc6zj9hM0Lqa3yDztXZGRd1aHLtvh0TKU6SDWi+unJnWMN7ij0VYBRSCi/cJ80OzZxOKNoYRimLDcVaJ
MKtr4NCEQlupqF6rzb3xmjt37Lg7N/e6e+xo/fjR40cLDcfB81mRIFcNy6jesrmlrgjrIJWgrFXOnZjQ
dkXYcLRTq0FXojXVcbJWvVEjxqJJW3BEa4uwZk6M0tomdlC1uW+yChGXqpBlmAgNAM9mbi+jWKUefqh8
cBcjazMU3jt5egKepmnhmMU6ngEyB008ozxp3szsphi1nvEljl07nFJEBOnIurlZYFFXmVyty3XId+nN
OHYXu2ocgc+B4SqKE7mHISqYshaa0/Mt5Q5XoRDHEyu1LmQ3uYZCuZFRmUl7nAXUsFh2cwpKECYDLjof
oYw4k1iAOnwy7mUmXPKJImbGhw4JVF3BwFzYGvpT3XCypzyUzdqN49xorc2iVCGOaQAMx/5b87W2dsUx
hhK1jsG+g85HUs0bKI/Vux5RXruEQkDzxFRX5tRq0Darj/YzAJE3cACU6zzwoRi6SQlddlUbWUmgnK1B
AyiZz5UHht41TkCz2YSinVex/AreHCV6/3A7GpGlgA+WF0+5ERESSyPicm7tqydvId8oV6Dk4aodExAJ
r2xYeYZDQ2YilLD8L7pmKpi/qGldM5TOwTfnf7HpsvnVyGaffZh/BwAA///3089mXggAAA==
`,
	},

	"/generator/template/ts/service_fetch.gots": {
		name:    "service_fetch.gots",
		local:   "generator/template/ts/service_fetch.gots",
		size:    1570,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xT3W4bRRi936f4ZFXyj9x1gkQVnCZS+FPNRW0p5gHGu5/tCeuZZWbWwVpGaghuQUpI
JCIkmogWoaKKogYkKCQl8DJe27niFdDOrO2kcIGvdj3nO+fM+c5WSiWnBM0uldCmAQKV0EGGgij0oTWA
fCi44iSkeQNDi/I4U4QyCW3BmULmw0ajBh73EVSXKNjm4gPYpqoLbVRe15y2uYA7zWYDIkk6KDO6hZiZ
phK2BVUKGVAGzUGIm56gocrQBhMK3qc+SiDQIpJ6ltDwGxckCIAwH3pkAAzRB+JvRVL1kCkgnseFT1kH
FAcZokfb1EsZt9BTIPDDiAq0SOaDpCoiinLmlODm//85JZiePhl/9WD08sXk6NH4s8PR+RezHJ0S2JNk
b5gcPks+3588O53+tDs5errRqE0efjp6+d3k8c7ff+wlZy9GF39Njp6aDJOdh2l6FrkguT+0eIu8Ftj0
9IlVz6AHP4wPDkd/HlvJjUbNciWPzscnPy5Un39z+fVwvDtM7v+aHJxOdy8uT+5Nv98ZPz4b7z9Phr+N
Lr68/PZscpw9p0c/fzL6fT853Jve2xtdnCQPzsfHv4yPzpxSxaG9kAsFsQMAEMeCsA7CDTUIsQw3WpwH
UF2DQgdVzQDfJoqkl5DgvhsxL81eFrXOpm/aSdC6nP2DzNfa0WkPe5B3K3HsvhUQKe+SHmpdb23J/Orc
w7xs74ugDCgEF3cI84O0DguKLgYhivyq4/SJSBuWwmENcl2lwmqlsvzGa+7yrRV3efl199ZSdWVpZSm3
6jj4kRFpZ65hE9WbdrYQiaAKUgnKOsUsiQVtJIJVRzuOuZ03856msrgJ3NTacSoViCTaL8qZ66R1v11j
TZNoPVLpw3pBouhTD2eyZeih6nJ/8R4SQXqyCnayWIWG4D0q8XZGAR8Dwz6K9cxwgAqu3APWroZZyK5T
hkx3pldcdcy0QBUJZq2ncZQhnjvKN+qbzXwZWtwfVOG9zfpd12rQ9qBgbRZBF13VRVYQKGFt5ukKc+be
FSh50McU525JzgrFooHqouuRVByF+E+Ca21IUbO5+XJsdRe11NruzEym+zI1fid9MwG69dDWF7T+Vzvi
2LUVLcwWEcdujYU2fK2vLCSO3Xqk5ievLibzb2pwnaMMr4yuF3JxvOiY1rky5OZOcrNSmCubT8v07p8A
AAD//6nV5A0iBgAA
`,
	},

	"/generator/template/yii2/Module.gophp": {
		name:    "Module.gophp",
		local:   "generator/template/yii2/Module.gophp",
		size:    1375,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RTXW/aMBR996+4qpBCETTvYbCPatNWiVJBN6la0GSSG/Dm2J7t7ENR/vtkJ4QQqKrl
heTec8+xzz28eq32ihBBczSKJghleXNPc1y7j6qaElIYhCfGpv7lL2Pxb9zGKzRKCoPH6pYajN9JaY3V
VH0SFnVGE5wSEo5GBEZAFYNcpgVHSDFjglkmBSScGkNgFBL/BosagX8sitRA3FI3DZYrjjkKa+BcjJQE
AMALumcE5Rsm9qiZTWVSNcXQ/6piy1kCg0QKqyXnqO9bD2YQnNoQH1EmmJL/VskKkfjrumsPr32rPqsH
UY3CRlHdnLb1J8aiaECVmsx14/dknkmdUwszOKwgij4sV4u3j9/u1sv75my9aYP2VuZKCufb8GsLcU+g
8WeBxgYwm8Npy7f9WnwzPi7fT0SR743PZxTVxjl1mdJDqFKcJdSZEn43UnhscFC4M1I8eJLgnH5zQREF
3XK8lfIHwy+Us9Qze9KMcoPPjxidvTzQkwxQa6k/UpFy1C8b1w/Tvh408fsuzbOKm+vOVgfNMMxgYPfM
TOY7tMPTA3UiFJ+moAcct3SdkUPJhW7HjEV9yGRFLmZ6e/gfDp1OP9y19g7tZ80XVNCd45vMaZquCo79
NJblBAYG9S+WoLMsmt2s6y/vYFWdoTUVO4QWtUC7l6npAa8elutHKMsudVXdNHupqiu3pqt+P6SKhUfM
+EwbRQodpc24zk5rV0X+BQAA//9/+Ld9XwUAAA==
`,
	},

	"/generator/template/yii2/RequestHandler.gophp": {
		name:    "RequestHandler.gophp",
		local:   "generator/template/yii2/RequestHandler.gophp",
		size:    373,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xPwUrEMBS85yvewYMWbD9gBT3qQQXXYy+hGW2hec3mJaKE9++y6VpQPOzpMW9mmJmb
2zAGY9h6SLADqJT2yXrsj0B1Z0yWv8/eLw6z7IwZZitCLzhkSLq37GZEwmcCO6FxxdL/5oshIirlmqLl
d1C7R/yYBrSPSOPiRLUKuqaplxq6CzZaT2tqX0pNPfah9oFDTq9fAap0EXHYLBEpR/7X85zTZjrpu3rf
Mg9pWvhnrurlOZlXtE6qrTuafJjhwYlGRFRGT5PBTtWo+Q4AAP//qCrWrHUBAAA=
`,
	},

	"/generator/template/yii2/controllers/ApiController.gophp": {
		name:    "ApiController.gophp",
		local:   "generator/template/yii2/controllers/ApiController.gophp",
		size:    1195,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/2xTTY+bPBC+8ytG0UpA9CZ7hze0aVW1PbRdtb1UyypyzGSxRGyvPeSjyP+9whASkviC
PfPM1zMP/7/TpQ4CybZoNeMITTP/zrb4q304l3MlyaiqQmPTIKjtDWCrCqxs6l1/hOguRyHyPa7zj0N0
D1Dqr1L5k1GkllqkQcArZi0stThDAQ+EsrBwNgVNAACgjdgxQnhYlUwWPmtnr9eV4LCpJSehJAgpKIq9
qwtszwOVws6yUygsQOIe8qtxfuJbjZa+dKAoTn286+o8Tqdduik074Us0QgqFHe98fFuM2ss2U4oY287
GlywAM0MSkqSC3g6IMUGIs/UCg/Cko0med40aDnTCKP+82VNZd/9JI7PxcYFn0NWU4mSBGekTPgCC3ge
YdsT+pohLLIbni7KJImHtd4o/m+U5OU8ghtuBqk28qKZE8eencYw+Yow/4ZUqsI6d5dU5j9NQ4KqngHn
bgk2+AaLVpZJ8sC0nmWmW286grSWXg6dmvOmGUaC+Vepa/p91L7CbeQs83Jrn7PsFemDKo5PzLCtjeK7
8B2rRMEIr5K1IrjS6CzrWXcuOsVfqcIHCmmJSY5qc3eAHzUNE8RwJQmD9n5Ll6vyIFIrZgw7XoLOW6XS
qL3n8PR3559RomHVpwNH3a4rmvT56KgR1AbCYbwQhOTKGOQ0n8RjQaAsnAtc8C8AAP//YJBkmqsEAAA=
`,
	},

	"/generator/template/yii2/handlers/ErrorHandler.gophp": {
		name:    "ErrorHandler.gophp",
		local:   "generator/template/yii2/handlers/ErrorHandler.gophp",
		size:    819,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6RQXWvqQBB9318xiA8Grhe53L6YmmJF6EtLwb4IgTJuRrs02Vl2NuAH+e8lBmusVFqc
p5k95+w5M7d37s0pZbEgcagJdjv4+4QFzfZTVcVKlUIwNyZuGuYtc/rsOfDYmVgpnaMITL1n/4A2y8kD
rQPZTCDdGJMuUChtw2qnAABcuciNhmVpdTBswZPNyE/Xmlw997p0aKM9v1HVZZbQQsFYCWg18RIOsdJ7
s91bfn4XtfR1zY0ZDrvoXD/xJI6tUD+RgKGUCWcEI/g/GMQnkm5NhBEcrftJ4Ff0Hje96MitgHKhn6Sc
cFGwvTrov+uC/tbv5rvDNAYnUF2dgkRwRZ1R0k60ovDYAL3oz7lIAur3c8mLR01jmQVv7OqrsL2burhQ
hgHrC9Uv8WWqkM0OV6tUpT4CAAD//zSBgW8zAwAA
`,
	},

	"/generator/template/yii2/handlers/RequestHandler.gophp": {
		name:    "RequestHandler.gophp",
		local:   "generator/template/yii2/handlers/RequestHandler.gophp",
		size:    241,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/0yOMU4EMQxFe5/CBQUUzAUGCUEFBQgBDdI03sSwI2XibJwUYPnuiKw0ovR/tt+/uS3H
Apk21kKB0QynZ9r4bUzuM0DXc3xPyv/RsknkpPNY+BD5EVleqjS5K+sMQAdtlULDkEgVX/nUWdsD5Zi4
ggEiotk1VspfjNMTt6NEdR9gPy79kNaAnz2HtkpGs1HP/fJsX8zG/78Qp8dcenv/LuyOF5VPV/Ou4Rzd
weE3AAD//+j458PxAAAA
`,
	},

	"/generator/template/yii2/models/enum.gophp": {
		name:    "enum.gophp",
		local:   "generator/template/yii2/models/enum.gophp",
		size:    146,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/zyLsQrCQBBE+/2K+QHzA1EsRCu1tLJZc4sKlzW4OVCW/Xe5A9O8GZh56+30mKiY4PTd
Hflm172WsaEnGjKbwb3lmUdBVxkB+cyiyVCP5AQA7iu8We+C7vCUnAwRbRheajPc/+6m9gvnIhH9ooqm
CAr6BQAA//8rGSPnkgAAAA==
`,
	},

	"/generator/template/yii2/models/error.gophp": {
		name:    "error.gophp",
		local:   "generator/template/yii2/models/error.gophp",
		size:    1999,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RU0W4aOxB991fMXSFlkQI/wIWKKNu+0DZq81IlETLLAK4We2ubhsTyv1d4s4vxegPZ
h0TYc86cOeOZ/z+Vm5IQTreoSpojGAPDb3SLP90va0eE7BTCLyFehXi8k0KLaclGhOQFVQqMcf8PiAoH
1gLuNfKlgjr68Ya9ZlIKme1zLDUTHNi2LHCLXHtRX1EpukZiCACAMQOQlK8Rhp8ZFksF1rqLUgqNucYl
9IxxOQ8qawjypbWkCtwtCpbDasfzKidnOqVS0hfoSVSl4Ar7FdD9fTfr4WMrSJlSqNMG/5A0GpKnft9j
qtnYCvAPDGd0gQUks+lNNpv/yO6y6X12m3jkh6+nN0wNJg0ljMHpTfujk7iVkEjzDXToAKo8b0JRnjCm
vi9+Y65heEs1vX8pEQJFlaptCWPg+Hza7hpjbaivRg0mznNPS1fcX1qwJdUYZzp15eEJxg42itaFhcJo
FTGW4AW1uNxr8k8tuSDfxfZGGv4Bm0N0bXf8cZ7DdrWgo8aI9HjiCFvL1fapJfG7jtk+iv/oQP/3NtFB
Oa1Z1hspnl13moX1BTlKWjRbLU2uGvwVMAVcaMA9UzrxPO2urDqLL72gXoV63uRKjYm+Nmtb69laMMbl
9IYysCzW2WBOOrqwPlEV8krUO8lb9KOT6o9+hORazN+WYZy3urys8e8OaRDnvWUYTyKDcxR2fcngnOG7
PjMY7ZN+7aAl/wIAAP//fzbP1M8HAAA=
`,
	},

	"/generator/template/yii2/models/message.gophp": {
		name:    "message.gophp",
		local:   "generator/template/yii2/models/message.gophp",
		size:    1963,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RUUW/aPBR996+4X4TUIBX+AB9MTGV7YVu19WVqK2TCBTwF27PN2s7yf5/iNCFxnELz
0Arb5/jcc+/x/x/kXhJOD6glzRCshfFXesAf/pdzE0KOGuGnEH+FeLhVwoi5ZBNCspxqDdb6/wWixIFz
wA4yxwNyo6ECPHxBrekOiSUAANaOQFG+Qxh/YphvNDjnN6QSBjODGxhY6wkLCRUE+cY5Uh48rnOWwfbI
M8MEB8aZSalS9AUGCrUUXOOwBPq/b95afGwLKdMaTVrj75NaQ/I4HDaYKja2BfwN4yVdYw7Jcv5xsVx9
X9wu5neLm6RBXnwDs2d6NKspYQpebzqctM5thUKa7aFHB1Dd8CYU1RDG9Lf1L8wMjG+ooXcvEiFQVKo6
SJgCx6d2LyuMc6G+CjWaec8bWvrO/aE521CDcaa2K/ePMPWwSbQuzDVGq4ixBBPU4fLT1Fx15IL7LrY3
0vB32ByiK7vjw3kO29eCnhoj0uMXR9g6rnZXHYnv9WT7JP69gf7vNdFBOZ0sm70ST7479YP1GTkqmi+e
M5SFijS5qvFXwDRwYQCfmTZJw9P+ysq1+KMX1KvRrOq7Umuj0+Zc5+11Dqz1dzZCGVgW62yQk54u7Fqq
Ql6F5qh4h37Sqv7kR0huxOr1MYzzlpuXNf7NkAbnGrMM01kkOCdh15cE5wzf9ZlgdFeGlYOO/AsAAP//
ZvO2hqsHAAA=
`,
	},

	"/generator/template": {
		name:  "template",
		local: `generator/template`,
		isDir: true,
	},

	"/generator/template/go": {
		name:  "go",
		local: `generator/template/go`,
		isDir: true,
	},

	"/generator/template/ts": {
		name:  "ts",
		local: `generator/template/ts`,
		isDir: true,
	},

	"/generator/template/yii2": {
		name:  "yii2",
		local: `generator/template/yii2`,
		isDir: true,
	},

	"/generator/template/yii2/controllers": {
		name:  "controllers",
		local: `generator/template/yii2/controllers`,
		isDir: true,
	},

	"/generator/template/yii2/handlers": {
		name:  "handlers",
		local: `generator/template/yii2/handlers`,
		isDir: true,
	},

	"/generator/template/yii2/models": {
		name:  "models",
		local: `generator/template/yii2/models`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"generator/template": {
		_escData["/generator/template/echo_enum.gogo"],
		_escData["/generator/template/echo_service.gogo"],
		_escData["/generator/template/echo_struct.gogo"],
		_escData["/generator/template/go"],
		_escData["/generator/template/go_client.gogo"],
		_escData["/generator/template/php_client.gophp"],
		_escData["/generator/template/spring_service.gojava"],
		_escData["/generator/template/spring_struct.gojava"],
		_escData["/generator/template/ts"],
		_escData["/generator/template/yii2"],
	},

	"generator/template/go": {
		_escData["/generator/template/go/enum.gogo"],
		_escData["/generator/template/go/service.gogo"],
		_escData["/generator/template/go/struct.gogo"],
	},

	"generator/template/ts": {
		_escData["/generator/template/ts/helper.gots"],
		_escData["/generator/template/ts/objs.gots"],
		_escData["/generator/template/ts/service_axios.gots"],
		_escData["/generator/template/ts/service_fetch.gots"],
	},

	"generator/template/yii2": {
		_escData["/generator/template/yii2/Module.gophp"],
		_escData["/generator/template/yii2/RequestHandler.gophp"],
		_escData["/generator/template/yii2/controllers"],
		_escData["/generator/template/yii2/handlers"],
		_escData["/generator/template/yii2/models"],
	},

	"generator/template/yii2/controllers": {
		_escData["/generator/template/yii2/controllers/ApiController.gophp"],
	},

	"generator/template/yii2/handlers": {
		_escData["/generator/template/yii2/handlers/ErrorHandler.gophp"],
		_escData["/generator/template/yii2/handlers/RequestHandler.gophp"],
	},

	"generator/template/yii2/models": {
		_escData["/generator/template/yii2/models/enum.gophp"],
		_escData["/generator/template/yii2/models/error.gophp"],
		_escData["/generator/template/yii2/models/message.gophp"],
	},
}
