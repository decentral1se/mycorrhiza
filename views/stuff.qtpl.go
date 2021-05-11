// Code generated by qtc from "stuff.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/stuff.qtpl:1
package views

//line views/stuff.qtpl:1
import "path/filepath"

//line views/stuff.qtpl:2
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/stuff.qtpl:3
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/stuff.qtpl:4
import "github.com/bouncepaw/mycorrhiza/user"

//line views/stuff.qtpl:5
import "github.com/bouncepaw/mycorrhiza/util"

//line views/stuff.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/stuff.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/stuff.qtpl:7
func StreamBaseHTML(qw422016 *qt422016.Writer, title, body string, u *user.User, headElements ...string) {
//line views/stuff.qtpl:7
	qw422016.N().S(`
<!doctype html>
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta charset="utf-8">
		<link rel="stylesheet" type="text/css" href="/static/common.css">
		<title>`)
//line views/stuff.qtpl:14
	qw422016.E().S(title)
//line views/stuff.qtpl:14
	qw422016.N().S(`</title>
		`)
//line views/stuff.qtpl:15
	for _, el := range headElements {
//line views/stuff.qtpl:15
		qw422016.N().S(el)
//line views/stuff.qtpl:15
	}
//line views/stuff.qtpl:15
	qw422016.N().S(`
	</head>
	<body>
		<header>
			<nav class="header-links main-width">
				<ul class="header-links__list">
`)
//line views/stuff.qtpl:21
	for _, link := range cfg.HeaderLinks {
//line views/stuff.qtpl:21
		qw422016.N().S(`					<li class="header-links__entry"><a class="header-links__link" href="`)
//line views/stuff.qtpl:22
		qw422016.E().S(link.Href)
//line views/stuff.qtpl:22
		qw422016.N().S(`">`)
//line views/stuff.qtpl:22
		qw422016.E().S(link.Display)
//line views/stuff.qtpl:22
		qw422016.N().S(`</a></li>
`)
//line views/stuff.qtpl:23
	}
//line views/stuff.qtpl:23
	qw422016.N().S(`					`)
//line views/stuff.qtpl:24
	qw422016.N().S(UserMenuHTML(u))
//line views/stuff.qtpl:24
	qw422016.N().S(`
				</ul>
			</nav>
		</header>
		`)
//line views/stuff.qtpl:28
	qw422016.N().S(body)
//line views/stuff.qtpl:28
	qw422016.N().S(`
	</body>
</html>
`)
//line views/stuff.qtpl:31
}

//line views/stuff.qtpl:31
func WriteBaseHTML(qq422016 qtio422016.Writer, title, body string, u *user.User, headElements ...string) {
//line views/stuff.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:31
	StreamBaseHTML(qw422016, title, body, u, headElements...)
//line views/stuff.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:31
}

//line views/stuff.qtpl:31
func BaseHTML(title, body string, u *user.User, headElements ...string) string {
//line views/stuff.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:31
	WriteBaseHTML(qb422016, title, body, u, headElements...)
//line views/stuff.qtpl:31
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:31
	return qs422016
//line views/stuff.qtpl:31
}

//line views/stuff.qtpl:33
func StreamUserListHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:33
	qw422016.N().S(`
<div class="layout">
<main class="main-width user-list">
	<h1>List of users</h1>
`)
//line views/stuff.qtpl:38
	var (
		admins     = make([]string, 0)
		moderators = make([]string, 0)
		editors    = make([]string, 0)
	)
	for u := range user.YieldUsers() {
		switch u.Group {
		case "admin":
			admins = append(admins, u.Name)
		case "moderator":
			moderators = append(moderators, u.Name)
		case "editor", "trusted":
			editors = append(editors, u.Name)
		}
	}

//line views/stuff.qtpl:53
	qw422016.N().S(`
	<section>
		<h2>Admins</h2>
		<ol>`)
//line views/stuff.qtpl:56
	for _, name := range admins {
//line views/stuff.qtpl:56
		qw422016.N().S(`
			<li><a href="/page/`)
//line views/stuff.qtpl:57
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:57
		qw422016.N().S(`/`)
//line views/stuff.qtpl:57
		qw422016.E().S(name)
//line views/stuff.qtpl:57
		qw422016.N().S(`">`)
//line views/stuff.qtpl:57
		qw422016.E().S(name)
//line views/stuff.qtpl:57
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:58
	}
//line views/stuff.qtpl:58
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>Moderators</h2>
		<ol>`)
//line views/stuff.qtpl:62
	for _, name := range moderators {
//line views/stuff.qtpl:62
		qw422016.N().S(`
			<li><a href="/page/`)
//line views/stuff.qtpl:63
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:63
		qw422016.N().S(`/`)
//line views/stuff.qtpl:63
		qw422016.E().S(name)
//line views/stuff.qtpl:63
		qw422016.N().S(`">`)
//line views/stuff.qtpl:63
		qw422016.E().S(name)
//line views/stuff.qtpl:63
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:64
	}
//line views/stuff.qtpl:64
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>Editors</h2>
		<ol>`)
//line views/stuff.qtpl:68
	for _, name := range editors {
//line views/stuff.qtpl:68
		qw422016.N().S(`
			<li><a href="/page/`)
//line views/stuff.qtpl:69
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:69
		qw422016.N().S(`/`)
//line views/stuff.qtpl:69
		qw422016.E().S(name)
//line views/stuff.qtpl:69
		qw422016.N().S(`">`)
//line views/stuff.qtpl:69
		qw422016.E().S(name)
//line views/stuff.qtpl:69
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:70
	}
//line views/stuff.qtpl:70
	qw422016.N().S(`</ol>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:74
}

//line views/stuff.qtpl:74
func WriteUserListHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:74
	StreamUserListHTML(qw422016)
//line views/stuff.qtpl:74
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:74
}

//line views/stuff.qtpl:74
func UserListHTML() string {
//line views/stuff.qtpl:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:74
	WriteUserListHTML(qb422016)
//line views/stuff.qtpl:74
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:74
	return qs422016
//line views/stuff.qtpl:74
}

//line views/stuff.qtpl:76
func StreamHyphaListHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:76
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>List of hyphae</h1>
	<p>This wiki has `)
//line views/stuff.qtpl:80
	qw422016.N().D(hyphae.Count())
//line views/stuff.qtpl:80
	qw422016.N().S(` hyphae.</p>
	<ul class="hypha-list">
		`)
//line views/stuff.qtpl:82
	for h := range hyphae.YieldExistingHyphae() {
//line views/stuff.qtpl:82
		qw422016.N().S(`
		<li class="hypha-list__entry">
			<a class="hypha-list__link" href="/hypha/`)
//line views/stuff.qtpl:84
		qw422016.E().S(h.Name)
//line views/stuff.qtpl:84
		qw422016.N().S(`">`)
//line views/stuff.qtpl:84
		qw422016.E().S(util.BeautifulName(h.Name))
//line views/stuff.qtpl:84
		qw422016.N().S(`</a>
			`)
//line views/stuff.qtpl:85
		if h.BinaryPath != "" {
//line views/stuff.qtpl:85
			qw422016.N().S(`
			<span class="hypha-list__amnt-type">`)
//line views/stuff.qtpl:86
			qw422016.E().S(filepath.Ext(h.BinaryPath)[1:])
//line views/stuff.qtpl:86
			qw422016.N().S(`</span>
			`)
//line views/stuff.qtpl:87
		}
//line views/stuff.qtpl:87
		qw422016.N().S(`
		</li>
		`)
//line views/stuff.qtpl:89
	}
//line views/stuff.qtpl:89
	qw422016.N().S(`
	</ul>
</main>
</div>
`)
//line views/stuff.qtpl:93
}

//line views/stuff.qtpl:93
func WriteHyphaListHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:93
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:93
	StreamHyphaListHTML(qw422016)
//line views/stuff.qtpl:93
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:93
}

//line views/stuff.qtpl:93
func HyphaListHTML() string {
//line views/stuff.qtpl:93
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:93
	WriteHyphaListHTML(qb422016)
//line views/stuff.qtpl:93
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:93
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:93
	return qs422016
//line views/stuff.qtpl:93
}

//line views/stuff.qtpl:95
func StreamAboutHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:95
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section>
		<h1>About `)
//line views/stuff.qtpl:99
	qw422016.E().S(cfg.WikiName)
//line views/stuff.qtpl:99
	qw422016.N().S(`</h1>
		<ul>
			<li><b><a href="https://mycorrhiza.lesarbr.es">MycorrhizaWiki</a> version:</b> 1.2.0 indev</li>
`)
//line views/stuff.qtpl:102
	if user.AuthUsed {
//line views/stuff.qtpl:102
		qw422016.N().S(`			<li><b>User count:</b> `)
//line views/stuff.qtpl:103
		qw422016.N().D(user.Count())
//line views/stuff.qtpl:103
		qw422016.N().S(`</li>
			<li><b>Home page:</b> <a href="/">`)
//line views/stuff.qtpl:104
		qw422016.E().S(cfg.HomeHypha)
//line views/stuff.qtpl:104
		qw422016.N().S(`</a></li>
			<li><b>Administrators:</b>`)
//line views/stuff.qtpl:105
		for i, username := range user.ListUsersWithGroup("admin") {
//line views/stuff.qtpl:106
			if i > 0 {
//line views/stuff.qtpl:106
				qw422016.N().S(`<span aria-hidden="true">, </span>
`)
//line views/stuff.qtpl:107
			}
//line views/stuff.qtpl:107
			qw422016.N().S(`				<a href="/page/`)
//line views/stuff.qtpl:108
			qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:108
			qw422016.N().S(`/`)
//line views/stuff.qtpl:108
			qw422016.E().S(username)
//line views/stuff.qtpl:108
			qw422016.N().S(`">`)
//line views/stuff.qtpl:108
			qw422016.E().S(username)
//line views/stuff.qtpl:108
			qw422016.N().S(`</a>`)
//line views/stuff.qtpl:108
		}
//line views/stuff.qtpl:108
		qw422016.N().S(`</li>
`)
//line views/stuff.qtpl:109
	} else {
//line views/stuff.qtpl:109
		qw422016.N().S(`			<li>This wiki does not use authorization</li>
`)
//line views/stuff.qtpl:111
	}
//line views/stuff.qtpl:111
	qw422016.N().S(`		</ul>
		<p>See <a href="/list">/list</a> for information about hyphae on this wiki.</p>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:117
}

//line views/stuff.qtpl:117
func WriteAboutHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:117
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:117
	StreamAboutHTML(qw422016)
//line views/stuff.qtpl:117
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:117
}

//line views/stuff.qtpl:117
func AboutHTML() string {
//line views/stuff.qtpl:117
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:117
	WriteAboutHTML(qb422016)
//line views/stuff.qtpl:117
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:117
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:117
	return qs422016
//line views/stuff.qtpl:117
}

//line views/stuff.qtpl:119
func StreamAdminPanelHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:119
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>Administrative functions</h1>
	<section>
		<h2>Safe things</h2>
		<ul>
			<li><a href="/about">About this wiki<a></li>
			<li><a href="/user-list">User list</a></li>
			<li><a href="/update-header-links">Update header links</a></li>
		</ul>
	</section>
	<section>
		<h2>Dangerous things</h2>
		<form action="/admin/shutdown" method="POST" style="float:left">
			<fieldset>
				<legend>Shutdown wiki</legend>
				<input type="submit">
			</fieldset>
		</form>
		<form action="/reindex" method="GET" style="float:left">
			<fieldset>
				<legend>Reindex hyphae</legend>
				<input type="submit">
			</fieldset>
		</form>
		<form action="/admin/reindex-users" method="POST" style="float:left">
			<fieldset>
				<legend>Reindex users</legend>
				<input type="submit">
			</fieldset>
		</form>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:154
}

//line views/stuff.qtpl:154
func WriteAdminPanelHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:154
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:154
	StreamAdminPanelHTML(qw422016)
//line views/stuff.qtpl:154
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:154
}

//line views/stuff.qtpl:154
func AdminPanelHTML() string {
//line views/stuff.qtpl:154
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:154
	WriteAdminPanelHTML(qb422016)
//line views/stuff.qtpl:154
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:154
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:154
	return qs422016
//line views/stuff.qtpl:154
}
