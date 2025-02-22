// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.16
// +build go1.16

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/website"
	"golang.org/x/website/internal/godoc"
)

// Test that the release history page includes expected entries.
//
// At this time, the test is very strict and checks that all releases
// from Go 1 to Go 1.14.2 are included with exact HTML content.
// It can be relaxed whenever the presentation of the release history
// page needs to be changed.
func TestReleaseHistory(t *testing.T) {
	origFS, origPres := fsys, pres
	defer func() { fsys, pres = origFS, origPres }()
	fsys = website.Content
	pres = godoc.NewPresentation(godoc.NewCorpus(fsys))
	readTemplates(pres)
	mux := registerHandlers(pres)

	req := httptest.NewRequest(http.MethodGet, "/doc/devel/release", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	resp := rr.Result()
	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("got status code %d %s, want %d %s", got, http.StatusText(got), want, http.StatusText(want))
	}
	if got, want := resp.Header.Get("Content-Type"), "text/html; charset=utf-8"; got != want {
		t.Errorf("got Content-Type header %q, want %q", got, want)
	}
	if !strings.Contains(foldSpace(rr.Body.String()), foldSpace(wantGo114HTML)) {
		t.Errorf("got body that doesn't contain expected Go 1.14 release history entries")
		println("HAVE")
		println(rr.Body.String())
		println("WANT")
		println(wantGo114HTML)
	}
	if !strings.Contains(foldSpace(rr.Body.String()), foldSpace(wantGo113HTML)) {
		t.Errorf("got body that doesn't contain expected Go 1.13 release history entries")
	}
	if !strings.Contains(foldSpace(rr.Body.String()), foldSpace(wantOldReleaseHTML)) {
		t.Errorf("got body that doesn't contain expected Go 1.12 and older release history entries")
	}
}

// foldSpace returns s with each instance of one or more consecutive
// white space characters, as defined by unicode.IsSpace, replaced
// by a single space ('\x20') character, with leading and trailing
// white space removed.
func foldSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

const wantGo114HTML = `
<h2 id="go1.14">go1.14 (released 2020-02-25)</h2>

<p>
Go 1.14 is a major release of Go.
Read the <a href="/doc/go1.14">Go 1.14 Release Notes</a> for more information.
</p>

<h3 id="go1.14.minor">Minor revisions</h3>

<p>
go1.14.1 (released 2020-03-19) includes fixes to the go command, tools, and the runtime. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.14.1+label%3ACherryPickApproved">Go
1.14.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.14.2 (released 2020-04-08) includes fixes to cgo, the go command, the runtime,
and the <code>os/exec</code> and <code>testing</code> packages. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.14.2+label%3ACherryPickApproved">Go
1.14.2 milestone</a> on our issue tracker for details.
</p>
`

const wantGo113HTML = `
<h2 id="go1.13">go1.13 (released 2019-09-03)</h2>

<p>
Go 1.13 is a major release of Go.
Read the <a href="/doc/go1.13">Go 1.13 Release Notes</a> for more information.
</p>

<h3 id="go1.13.minor">Minor revisions</h3>

<p>
go1.13.1 (released 2019-09-25) includes security fixes to the
<code>net/http</code> and <code>net/textproto</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.1+label%3ACherryPickApproved">Go
1.13.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.2 (released 2019-10-17) includes security fixes to the
compiler and the <code>crypto/dsa</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.2+label%3ACherryPickApproved">Go
1.13.2 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.3 (released 2019-10-17) includes fixes to the go command,
the toolchain, the runtime, and the <code>syscall</code>, <code>net</code>,
<code>net/http</code>, and <code>crypto/ecdsa</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.3+label%3ACherryPickApproved">Go
1.13.3 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.4 (released 2019-10-31) includes fixes to the <code>net/http</code> and
<code>syscall</code> packages. It also fixes an issue on macOS 10.15 Catalina
where the non-notarized installer and binaries were being
<a href="https://golang.org/issue/34986">rejected by Gatekeeper</a>.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.4+label%3ACherryPickApproved">Go
1.13.4 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.5 (released 2019-12-04) includes fixes to the go command, the runtime,
the linker, and the <code>net/http</code> package. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.5+label%3ACherryPickApproved">Go
1.13.5 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.6 (released 2020-01-09) includes fixes to the runtime and
the <code>net/http</code> package. See
the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.6+label%3ACherryPickApproved">Go
1.13.6 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.7 (released 2020-01-28) includes two security fixes to
the <code>crypto/x509</code> package. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.7+label%3ACherryPickApproved">Go
1.13.7 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.8 (released 2020-02-12) includes fixes to the runtime, and the
<code>crypto/x509</code> and <code>net/http</code> packages. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.8+label%3ACherryPickApproved">Go
1.13.8 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.9 (released 2020-03-19) includes fixes to the go command, tools, the runtime, the
toolchain, and the <code>crypto/cypher</code> package. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.9+label%3ACherryPickApproved">Go
1.13.9 milestone</a> on our issue tracker for details.
</p>

<p>
go1.13.10 (released 2020-04-08) includes fixes to the go command, the runtime,
and the <code>os/exec</code> and <code>time</code> packages. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.13.10+label%3ACherryPickApproved">Go
1.13.10 milestone</a> on our issue tracker for details.
</p>
`

const wantOldReleaseHTML = `
<h2 id="go1.12">go1.12 (released 2019-02-25)</h2>

<p>
Go 1.12 is a major release of Go.
Read the <a href="/doc/go1.12">Go 1.12 Release Notes</a> for more information.
</p>

<h3 id="go1.12.minor">Minor revisions</h3>

<p>
go1.12.1 (released 2019-03-14) includes fixes to cgo, the compiler, the go
command, and the <code>fmt</code>, <code>net/smtp</code>, <code>os</code>,
<code>path/filepath</code>, <code>sync</code>, and <code>text/template</code>
packages. See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.1+label%3ACherryPickApproved">Go
1.12.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.2 (released 2019-04-05) includes fixes to the compiler, the go
command, the runtime, and the <code>doc</code>, <code>net</code>,
<code>net/http/httputil</code>, and <code>os</code> packages. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.2+label%3ACherryPickApproved">Go
1.12.2 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.3 (released 2019-04-08) was accidentally released without its
intended fix. It is identical to go1.12.2, except for its version
number. The intended fix is in go1.12.4.
</p>

<p>
go1.12.4 (released 2019-04-11) fixes an issue where using the prebuilt binary
releases on older versions of GNU/Linux
<a href="https://golang.org/issues/31293">led to failures</a>
when linking programs that used cgo.
Only Linux users who hit this issue need to update.
</p>

<p>
go1.12.5 (released 2019-05-06) includes fixes to the compiler, the linker,
the go command, the runtime, and the <code>os</code> package. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.5+label%3ACherryPickApproved">Go
1.12.5 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.6 (released 2019-06-11) includes fixes to the compiler, the linker,
the go command, and the <code>crypto/x509</code>, <code>net/http</code>, and
<code>os</code> packages. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.6+label%3ACherryPickApproved">Go
1.12.6 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.7 (released 2019-07-08) includes fixes to cgo, the compiler,
and the linker.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.7+label%3ACherryPickApproved">Go
1.12.7 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.8 (released 2019-08-13) includes security fixes to the
<code>net/http</code> and <code>net/url</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.8+label%3ACherryPickApproved">Go
1.12.8 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.9 (released 2019-08-15) includes fixes to the linker,
and the <code>os</code> and <code>math/big</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.9+label%3ACherryPickApproved">Go
1.12.9 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.10 (released 2019-09-25) includes security fixes to the
<code>net/http</code> and <code>net/textproto</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.10+label%3ACherryPickApproved">Go
1.12.10 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.11 (released 2019-10-17) includes security fixes to the
<code>crypto/dsa</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.11+label%3ACherryPickApproved">Go
1.12.11 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.12 (released 2019-10-17) includes fixes to the go command,
runtime, and the <code>syscall</code> and <code>net</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.12+label%3ACherryPickApproved">Go
1.12.12 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.13 (released 2019-10-31) fixes an issue on macOS 10.15 Catalina
where the non-notarized installer and binaries were being
<a href="https://golang.org/issue/34986">rejected by Gatekeeper</a>.
Only macOS users who hit this issue need to update.
</p>

<p>
go1.12.14 (released 2019-12-04) includes a fix to the runtime. See
the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.14+label%3ACherryPickApproved">Go
1.12.14 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.15 (released 2020-01-09) includes fixes to the runtime and
the <code>net/http</code> package. See
the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.15+label%3ACherryPickApproved">Go
1.12.15 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.16 (released 2020-01-28) includes two security fixes to
the <code>crypto/x509</code> package. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.16+label%3ACherryPickApproved">Go
1.12.16 milestone</a> on our issue tracker for details.
</p>

<p>
go1.12.17 (released 2020-02-12) includes a fix to the runtime. See
the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.17+label%3ACherryPickApproved">Go
1.12.17 milestone</a> on our issue tracker for details.
</p>

<h2 id="go1.11">go1.11 (released 2018-08-24)</h2>

<p>
Go 1.11 is a major release of Go.
Read the <a href="/doc/go1.11">Go 1.11 Release Notes</a> for more information.
</p>

<h3 id="go1.11.minor">Minor revisions</h3>

<p>
go1.11.1 (released 2018-10-01) includes fixes to the compiler, documentation, go
command, runtime, and the <code>crypto/x509</code>, <code>encoding/json</code>,
<code>go/types</code>, <code>net</code>, <code>net/http</code>, and
<code>reflect</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.1+label%3ACherryPickApproved">Go
1.11.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.2 (released 2018-11-02) includes fixes to the compiler, linker,
documentation, go command, and the <code>database/sql</code> and
<code>go/types</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.2+label%3ACherryPickApproved">Go
1.11.2 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.3 (released 2018-12-12) includes three security fixes to "go get" and
the <code>crypto/x509</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.3+label%3ACherryPickApproved">Go
1.11.3 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.4 (released 2018-12-14) includes fixes to cgo, the compiler, linker,
runtime, documentation, go command, and the <code>net/http</code> and
<code>go/types</code> packages.
It includes a fix to a bug introduced in Go 1.11.3 that broke <code>go</code>
<code>get</code> for import path patterns containing "<code>...</code>".
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.4+label%3ACherryPickApproved">Go
1.11.4 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.5 (released 2019-01-23) includes a security fix to the
<code>crypto/elliptic</code> package.  See
the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.5+label%3ACherryPickApproved">Go
1.11.5 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.6 (released 2019-03-14) includes fixes to cgo, the compiler, linker,
runtime, go command, and the <code>crypto/x509</code>, <code>encoding/json</code>,
<code>net</code>, and <code>net/url</code> packages. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.6+label%3ACherryPickApproved">Go
1.11.6 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.7 (released 2019-04-05) includes fixes to the runtime and the
<code>net</code> package. See the
<a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.7+label%3ACherryPickApproved">Go
1.11.7 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.8 (released 2019-04-08) was accidentally released without its
intended fix. It is identical to go1.11.7, except for its version
number. The intended fix is in go1.11.9.
</p>

<p>
go1.11.9 (released 2019-04-11) fixes an issue where using the prebuilt binary
releases on older versions of GNU/Linux
<a href="https://golang.org/issues/31293">led to failures</a>
when linking programs that used cgo.
Only Linux users who hit this issue need to update.
</p>

<p>
go1.11.10 (released 2019-05-06) includes fixes to the runtime and the linker.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.10+label%3ACherryPickApproved">Go
1.11.10 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.11 (released 2019-06-11) includes a fix to the <code>crypto/x509</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.11+label%3ACherryPickApproved">Go
1.11.11 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.12 (released 2019-07-08) includes fixes to the compiler and the linker.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.12+label%3ACherryPickApproved">Go
1.11.12 milestone</a> on our issue tracker for details.
</p>

<p>
go1.11.13 (released 2019-08-13) includes security fixes to the
<code>net/http</code> and <code>net/url</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.11.13+label%3ACherryPickApproved">Go
1.11.13 milestone</a> on our issue tracker for details.
</p>

<h2 id="go1.10">go1.10 (released 2018-02-16)</h2>

<p>
Go 1.10 is a major release of Go.
Read the <a href="/doc/go1.10">Go 1.10 Release Notes</a> for more information.
</p>

<h3 id="go1.10.minor">Minor revisions</h3>

<p>
go1.10.1 (released 2018-03-28) includes fixes to the compiler, runtime, and the
<code>archive/zip</code>, <code>crypto/tls</code>, <code>crypto/x509</code>,
<code>encoding/json</code>, <code>net</code>, <code>net/http</code>, and
<code>net/http/pprof</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.10.1+label%3ACherryPickApproved">Go
1.10.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.10.2 (released 2018-05-01) includes fixes to the compiler, linker, and go
command.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.10.2+label%3ACherryPickApproved">Go
1.10.2 milestone</a> on our issue tracker for details.
</p>

<p>
go1.10.3 (released 2018-06-05) includes fixes to the go command, and the
<code>crypto/tls</code>, <code>crypto/x509</code>, and <code>strings</code> packages.
In particular, it adds <a href="https://go.googlesource.com/go/+/d4e21288e444d3ffd30d1a0737f15ea3fc3b8ad9">
minimal support to the go command for the vgo transition</a>.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.10.3+label%3ACherryPickApproved">Go
1.10.3 milestone</a> on our issue tracker for details.
</p>

<p>
go1.10.4 (released 2018-08-24) includes fixes to the go command, linker, and the
<code>net/http</code>, <code>mime/multipart</code>, <code>ld/macho</code>,
<code>bytes</code>, and <code>strings</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.10.4+label%3ACherryPickApproved">Go
1.10.4 milestone</a> on our issue tracker for details.
</p>

<p>
go1.10.5 (released 2018-11-02) includes fixes to the go command, linker, runtime,
and the <code>database/sql</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.10.5+label%3ACherryPickApproved">Go
1.10.5 milestone</a> on our issue tracker for details.
</p>

<p>
go1.10.6 (released 2018-12-12) includes three security fixes to "go get" and
the <code>crypto/x509</code> package.
It contains the same fixes as Go 1.11.3 and was released at the same time.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.10.6+label%3ACherryPickApproved">Go
1.10.6 milestone</a> on our issue tracker for details.
</p>

<p>
go1.10.7 (released 2018-12-14) includes a fix to a bug introduced in Go 1.10.6
that broke <code>go</code> <code>get</code> for import path patterns containing
"<code>...</code>".
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.10.7+label%3ACherryPickApproved">
Go 1.10.7 milestone</a> on our issue tracker for details.
</p>

<p>
go1.10.8 (released 2019-01-23) includes a security fix to the
<code>crypto/elliptic</code> package.  See
the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.10.8+label%3ACherryPickApproved">Go
1.10.8 milestone</a> on our issue tracker for details.
</p>

<h2 id="go1.9">go1.9 (released 2017-08-24)</h2>

<p>
Go 1.9 is a major release of Go.
Read the <a href="/doc/go1.9">Go 1.9 Release Notes</a> for more information.
</p>

<h3 id="go1.9.minor">Minor revisions</h3>

<p>
go1.9.1 (released 2017-10-04) includes two security fixes.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.9.1+label%3ACherryPickApproved">Go
1.9.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.9.2 (released 2017-10-25) includes fixes to the compiler, linker, runtime,
documentation, <code>go</code> command,
and the <code>crypto/x509</code>, <code>database/sql</code>, <code>log</code>,
and <code>net/smtp</code> packages.
It includes a fix to a bug introduced in Go 1.9.1 that broke <code>go</code> <code>get</code>
of non-Git repositories under certain conditions.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.9.2+label%3ACherryPickApproved">Go
1.9.2 milestone</a> on our issue tracker for details.
</p>

<p>
go1.9.3 (released 2018-01-22) includes fixes to the compiler, runtime,
and the <code>database/sql</code>, <code>math/big</code>, <code>net/http</code>,
and <code>net/url</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.9.3+label%3ACherryPickApproved">Go
1.9.3 milestone</a> on our issue tracker for details.
</p>

<p>
go1.9.4 (released 2018-02-07) includes a security fix to "go get".
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.9.4+label%3ACherryPickApproved">Go
1.9.4 milestone</a> on our issue tracker for details.
</p>

<p>
go1.9.5 (released 2018-03-28) includes fixes to the compiler, go command, and the
<code>net/http/pprof</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.9.5+label%3ACherryPickApproved">Go
1.9.5 milestone</a> on our issue tracker for details.
</p>

<p>
go1.9.6 (released 2018-05-01) includes fixes to the compiler and go command.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.9.6+label%3ACherryPickApproved">Go
1.9.6 milestone</a> on our issue tracker for details.
</p>

<p>
go1.9.7 (released 2018-06-05) includes fixes to the go command, and the
<code>crypto/x509</code> and <code>strings</code> packages.
In particular, it adds <a href="https://go.googlesource.com/go/+/d4e21288e444d3ffd30d1a0737f15ea3fc3b8ad9">
minimal support to the go command for the vgo transition</a>.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.9.7+label%3ACherryPickApproved">Go
1.9.7 milestone</a> on our issue tracker for details.
</p>

<h2 id="go1.8">go1.8 (released 2017-02-16)</h2>

<p>
Go 1.8 is a major release of Go.
Read the <a href="/doc/go1.8">Go 1.8 Release Notes</a> for more information.
</p>

<h3 id="go1.8.minor">Minor revisions</h3>

<p>
go1.8.1 (released 2017-04-07) includes fixes to the compiler, linker, runtime,
documentation, <code>go</code> command and the <code>crypto/tls</code>,
<code>encoding/xml</code>, <code>image/png</code>, <code>net</code>,
<code>net/http</code>, <code>reflect</code>, <code>text/template</code>,
and <code>time</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.8.1">Go
1.8.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.8.2 (released 2017-05-23) includes a security fix to the
<code>crypto/elliptic</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.8.2">Go
1.8.2 milestone</a> on our issue tracker for details.
</p>

<p>
go1.8.3 (released 2017-05-24) includes fixes to the compiler, runtime,
documentation, and the <code>database/sql</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.8.3">Go
1.8.3 milestone</a> on our issue tracker for details.
</p>

<p>
go1.8.4 (released 2017-10-04) includes two security fixes.
It contains the same fixes as Go 1.9.1 and was released at the same time.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.8.4">Go
1.8.4 milestone</a> on our issue tracker for details.
</p>

<p>
go1.8.5 (released 2017-10-25) includes fixes to the compiler, linker, runtime,
documentation, <code>go</code> command,
and the <code>crypto/x509</code> and <code>net/smtp</code> packages.
It includes a fix to a bug introduced in Go 1.8.4 that broke <code>go</code> <code>get</code>
of non-Git repositories under certain conditions.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.8.5">Go
1.8.5 milestone</a> on our issue tracker for details.
</p>

<p>
go1.8.6 (released 2018-01-22) includes the same fix in <code>math/big</code>
as Go 1.9.3 and was released at the same time.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.8.6">Go
1.8.6 milestone</a> on our issue tracker for details.
</p>

<p>
go1.8.7 (released 2018-02-07) includes a security fix to "go get".
It contains the same fix as Go 1.9.4 and was released at the same time.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.8.7">Go
1.8.7</a> milestone on our issue tracker for details.
</p>

<h2 id="go1.7">go1.7 (released 2016-08-15)</h2>

<p>
Go 1.7 is a major release of Go.
Read the <a href="/doc/go1.7">Go 1.7 Release Notes</a> for more information.
</p>

<h3 id="go1.7.minor">Minor revisions</h3>

<p>
go1.7.1 (released 2016-09-07) includes fixes to the compiler, runtime,
documentation, and the <code>compress/flate</code>, <code>hash/crc32</code>,
<code>io</code>, <code>net</code>, <code>net/http</code>,
<code>path/filepath</code>, <code>reflect</code>, and <code>syscall</code>
packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.7.1">Go
1.7.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.7.2 should not be used. It was tagged but not fully released.
The release was deferred due to a last minute bug report.
Use go1.7.3 instead, and refer to the summary of changes below.
</p>

<p>
go1.7.3 (released 2016-10-19) includes fixes to the compiler, runtime,
and the <code>crypto/cipher</code>, <code>crypto/tls</code>,
<code>net/http</code>, and <code>strings</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.7.3">Go
1.7.3 milestone</a> on our issue tracker for details.
</p>

<p>
go1.7.4 (released 2016-12-01) includes two security fixes.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.7.4">Go
1.7.4 milestone</a> on our issue tracker for details.
</p>

<p>
go1.7.5 (released 2017-01-26) includes fixes to the compiler, runtime,
and the <code>crypto/x509</code> and <code>time</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.7.5">Go
1.7.5 milestone</a> on our issue tracker for details.
</p>

<p>
go1.7.6 (released 2017-05-23) includes the same security fix as Go 1.8.2 and
was released at the same time.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.8.2">Go
1.8.2 milestone</a> on our issue tracker for details.
</p>

<h2 id="go1.6">go1.6 (released 2016-02-17)</h2>

<p>
Go 1.6 is a major release of Go.
Read the <a href="/doc/go1.6">Go 1.6 Release Notes</a> for more information.
</p>

<h3 id="go1.6.minor">Minor revisions</h3>

<p>
go1.6.1 (released 2016-04-12) includes two security fixes.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.6.1">Go
1.6.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.6.2 (released 2016-04-20) includes fixes to the compiler, runtime, tools,
documentation, and the <code>mime/multipart</code>, <code>net/http</code>, and
<code>sort</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.6.2">Go
1.6.2 milestone</a> on our issue tracker for details.
</p>

<p>
go1.6.3 (released 2016-07-17) includes security fixes to the
<code>net/http/cgi</code> package and <code>net/http</code> package when used in
a CGI environment.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.6.3">Go
1.6.3 milestone</a> on our issue tracker for details.
</p>

<p>
go1.6.4 (released 2016-12-01) includes two security fixes.
It contains the same fixes as Go 1.7.4 and was released at the same time.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.7.4">Go
1.7.4 milestone</a> on our issue tracker for details.
</p>

<h2 id="go1.5">go1.5 (released 2015-08-19)</h2>

<p>
Go 1.5 is a major release of Go.
Read the <a href="/doc/go1.5">Go 1.5 Release Notes</a> for more information.
</p>

<h3 id="go1.5.minor">Minor revisions</h3>

<p>
go1.5.1 (released 2015-09-08) includes bug fixes to the compiler, assembler, and
the <code>fmt</code>, <code>net/textproto</code>, <code>net/http</code>, and
<code>runtime</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.5.1">Go
1.5.1 milestone</a> on our issue tracker for details.
</p>

<p>
go1.5.2 (released 2015-12-02) includes bug fixes to the compiler, linker, and
the <code>mime/multipart</code>, <code>net</code>, and <code>runtime</code>
packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.5.2">Go
1.5.2 milestone</a> on our issue tracker for details.
</p>

<p>
go1.5.3 (released 2016-01-13) includes a security fix to the <code>math/big</code> package
affecting the <code>crypto/tls</code> package.
See the <a href="https://golang.org/s/go153announce">release announcement</a> for details.
</p>

<p>
go1.5.4 (released 2016-04-12) includes two security fixes.
It contains the same fixes as Go 1.6.1 and was released at the same time.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.6.1">Go
1.6.1 milestone</a> on our issue tracker for details.
</p>

<h2 id="go1.4">go1.4 (released 2014-12-10)</h2>

<p>
Go 1.4 is a major release of Go.
Read the <a href="/doc/go1.4">Go 1.4 Release Notes</a> for more information.
</p>

<h3 id="go1.4.minor">Minor revisions</h3>

<p>
go1.4.1 (released 2015-01-15) includes bug fixes to the linker and the <code>log</code>, <code>syscall</code>, and <code>runtime</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.4.1">Go 1.4.1 milestone on our issue tracker</a> for details.
</p>

<p>
go1.4.2 (released 2015-02-17) includes bug fixes to the <code>go</code> command, the compiler and linker, and the <code>runtime</code>, <code>syscall</code>, <code>reflect</code>, and <code>math/big</code> packages.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.4.2">Go 1.4.2 milestone on our issue tracker</a> for details.
</p>

<p>
go1.4.3 (released 2015-09-22) includes security fixes to the <code>net/http</code> package and bug fixes to the <code>runtime</code> package.
See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.4.3">Go 1.4.3 milestone on our issue tracker</a> for details.
</p>

<h2 id="go1.3">go1.3 (released 2014-06-18)</h2>

<p>
Go 1.3 is a major release of Go.
Read the <a href="/doc/go1.3">Go 1.3 Release Notes</a> for more information.
</p>

<h3 id="go1.3.minor">Minor revisions</h3>

<p>
go1.3.1 (released 2014-08-13) includes bug fixes to the compiler and the <code>runtime</code>, <code>net</code>, and <code>crypto/rsa</code> packages.
See the <a href="https://github.com/golang/go/commits/go1.3.1">change history</a> for details.
</p>

<p>
go1.3.2 (released 2014-09-25) includes bug fixes to cgo and the crypto/tls packages.
See the <a href="https://github.com/golang/go/commits/go1.3.2">change history</a> for details.
</p>

<p>
go1.3.3 (released 2014-09-30) includes further bug fixes to cgo, the runtime package, and the nacl port.
See the <a href="https://github.com/golang/go/commits/go1.3.3">change history</a> for details.
</p>

<h2 id="go1.2">go1.2 (released 2013-12-01)</h2>

<p>
Go 1.2 is a major release of Go.
Read the <a href="/doc/go1.2">Go 1.2 Release Notes</a> for more information.
</p>

<h3 id="go1.2.minor">Minor revisions</h3>

<p>
go1.2.1 (released 2014-03-02) includes bug fixes to the <code>runtime</code>, <code>net</code>, and <code>database/sql</code> packages.
See the <a href="https://github.com/golang/go/commits/go1.2.1">change history</a> for details.
</p>

<p>
go1.2.2 (released 2014-05-05) includes a
<a href="https://github.com/golang/go/commits/go1.2.2">security fix</a>
that affects the tour binary included in the binary distributions (thanks to Guillaume T).
</p>

<h2 id="go1.1">go1.1 (released 2013-05-13)</h2>

<p>
Go 1.1 is a major release of Go.
Read the <a href="/doc/go1.1">Go 1.1 Release Notes</a> for more information.
</p>

<h3 id="go1.1.minor">Minor revisions</h3>

<p>
go1.1.1 (released 2013-06-13) includes several compiler and runtime bug fixes.
See the <a href="https://github.com/golang/go/commits/go1.1.1">change history</a> for details.
</p>

<p>
go1.1.2 (released 2013-08-13) includes fixes to the <code>gc</code> compiler
and <code>cgo</code>, and the <code>bufio</code>, <code>runtime</code>,
<code>syscall</code>, and <code>time</code> packages.
See the <a href="https://github.com/golang/go/commits/go1.1.2">change history</a> for details.
If you use package syscall's <code>Getrlimit</code> and <code>Setrlimit</code>
functions under Linux on the ARM or 386 architectures, please note change
<a href="//golang.org/cl/11803043">11803043</a>
that fixes <a href="//golang.org/issue/5949">issue 5949</a>.
</p>

<h2 id="go1">go1 (released 2012-03-28)</h2>

<p>
Go 1 is a major release of Go that will be stable in the long term.
Read the <a href="/doc/go1.html">Go 1 Release Notes</a> for more information.
</p>

<p>
It is intended that programs written for Go 1 will continue to compile and run
correctly, unchanged, under future versions of Go 1.
Read the <a href="/doc/go1compat.html">Go 1 compatibility document</a> for more
about the future of Go 1.
</p>

<p>
The go1 release corresponds to
<code><a href="weekly.html#2012-03-27">weekly.2012-03-27</a></code>.
</p>

<h3 id="go1.minor">Minor revisions</h3>

<p>
go1.0.1 (released 2012-04-25) was issued to
<a href="//golang.org/cl/6061043">fix</a> an
<a href="//golang.org/issue/3545">escape analysis bug</a>
that can lead to memory corruption.
It also includes several minor code and documentation fixes.
</p>

<p>
go1.0.2 (released 2012-06-13) was issued to fix two bugs in the implementation
of maps using struct or array keys:
<a href="//golang.org/issue/3695">issue 3695</a> and
<a href="//golang.org/issue/3573">issue 3573</a>.
It also includes many minor code and documentation fixes.
</p>

<p>
go1.0.3 (released 2012-09-21) includes minor code and documentation fixes.
</p>

<p>
See the <a href="https://github.com/golang/go/commits/release-branch.go1">go1 release branch history</a> for the complete list of changes.
</p>
`
