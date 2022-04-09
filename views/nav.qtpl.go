// Code generated by qtc from "nav.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/nav.qtpl:1
package views

//line views/nav.qtpl:1
import "strings"

//line views/nav.qtpl:2
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/nav.qtpl:3
import "github.com/bouncepaw/mycorrhiza/backlinks"

//line views/nav.qtpl:4
import "github.com/bouncepaw/mycorrhiza/l18n"

//line views/nav.qtpl:5
import "github.com/bouncepaw/mycorrhiza/user"

//line views/nav.qtpl:6
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/nav.qtpl:7
import "github.com/bouncepaw/mycorrhiza/viewutil"

//line views/nav.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/nav.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/nav.qtpl:9
func streamhyphaInfoEntry(qw422016 *qt422016.Writer, h hyphae.Hypha, u *user.User, action, displayText string) {
//line views/nav.qtpl:9
	qw422016.N().S(`
`)
//line views/nav.qtpl:10
	if u.CanProceed(action) {
//line views/nav.qtpl:10
		qw422016.N().S(`
<li class="hypha-info__entry hypha-info__entry_`)
//line views/nav.qtpl:11
		qw422016.E().S(action)
//line views/nav.qtpl:11
		qw422016.N().S(`">
	<a class="hypha-info__link" href="/`)
//line views/nav.qtpl:12
		qw422016.E().S(action)
//line views/nav.qtpl:12
		qw422016.N().S(`/`)
//line views/nav.qtpl:12
		qw422016.E().S(h.CanonicalName())
//line views/nav.qtpl:12
		qw422016.N().S(`">`)
//line views/nav.qtpl:12
		qw422016.E().S(displayText)
//line views/nav.qtpl:12
		qw422016.N().S(`</a>
</li>
`)
//line views/nav.qtpl:14
	}
//line views/nav.qtpl:14
	qw422016.N().S(`
`)
//line views/nav.qtpl:15
}

//line views/nav.qtpl:15
func writehyphaInfoEntry(qq422016 qtio422016.Writer, h hyphae.Hypha, u *user.User, action, displayText string) {
//line views/nav.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:15
	streamhyphaInfoEntry(qw422016, h, u, action, displayText)
//line views/nav.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:15
}

//line views/nav.qtpl:15
func hyphaInfoEntry(h hyphae.Hypha, u *user.User, action, displayText string) string {
//line views/nav.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:15
	writehyphaInfoEntry(qb422016, h, u, action, displayText)
//line views/nav.qtpl:15
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:15
	return qs422016
//line views/nav.qtpl:15
}

//line views/nav.qtpl:17
func streamhyphaInfo(qw422016 *qt422016.Writer, meta viewutil.Meta, h hyphae.Hypha) {
//line views/nav.qtpl:17
	qw422016.N().S(`
`)
//line views/nav.qtpl:19
	u := meta.U
	lc := meta.Lc
	backs := backlinks.BacklinksCount(h)

//line views/nav.qtpl:22
	qw422016.N().S(`
<nav class="hypha-info">
	<ul class="hypha-info__list">
		`)
//line views/nav.qtpl:25
	streamhyphaInfoEntry(qw422016, h, u, "history", lc.Get("ui.history_link"))
//line views/nav.qtpl:25
	qw422016.N().S(`
		`)
//line views/nav.qtpl:26
	streamhyphaInfoEntry(qw422016, h, u, "rename", lc.Get("ui.rename_link"))
//line views/nav.qtpl:26
	qw422016.N().S(`
		`)
//line views/nav.qtpl:27
	streamhyphaInfoEntry(qw422016, h, u, "delete", lc.Get("ui.delete_link"))
//line views/nav.qtpl:27
	qw422016.N().S(`
		`)
//line views/nav.qtpl:28
	streamhyphaInfoEntry(qw422016, h, u, "text", lc.Get("ui.text_link"))
//line views/nav.qtpl:28
	qw422016.N().S(`
		`)
//line views/nav.qtpl:29
	streamhyphaInfoEntry(qw422016, h, u, "media", lc.Get("ui.media_link"))
//line views/nav.qtpl:29
	qw422016.N().S(`
		`)
//line views/nav.qtpl:30
	streamhyphaInfoEntry(qw422016, h, u, "backlinks", lc.GetPlural("ui.backlinks_link", backs))
//line views/nav.qtpl:30
	qw422016.N().S(`
	</ul>
</nav>
`)
//line views/nav.qtpl:33
}

//line views/nav.qtpl:33
func writehyphaInfo(qq422016 qtio422016.Writer, meta viewutil.Meta, h hyphae.Hypha) {
//line views/nav.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:33
	streamhyphaInfo(qw422016, meta, h)
//line views/nav.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:33
}

//line views/nav.qtpl:33
func hyphaInfo(meta viewutil.Meta, h hyphae.Hypha) string {
//line views/nav.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:33
	writehyphaInfo(qb422016, meta, h)
//line views/nav.qtpl:33
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:33
	return qs422016
//line views/nav.qtpl:33
}

//line views/nav.qtpl:35
func streamsiblingHyphae(qw422016 *qt422016.Writer, siblings string, lc *l18n.Localizer) {
//line views/nav.qtpl:35
	qw422016.N().S(`
`)
//line views/nav.qtpl:36
	if cfg.UseSiblingHyphaeSidebar {
//line views/nav.qtpl:36
		qw422016.N().S(`
<aside class="sibling-hyphae layout-card">
	<h2 class="sibling-hyphae__title layout-card__title">`)
//line views/nav.qtpl:38
		qw422016.E().S(lc.Get("ui.sibling_hyphae"))
//line views/nav.qtpl:38
		qw422016.N().S(`</h2>
	`)
//line views/nav.qtpl:39
		qw422016.N().S(siblings)
//line views/nav.qtpl:39
		qw422016.N().S(`
</aside>
`)
//line views/nav.qtpl:41
	}
//line views/nav.qtpl:41
	qw422016.N().S(`
`)
//line views/nav.qtpl:42
}

//line views/nav.qtpl:42
func writesiblingHyphae(qq422016 qtio422016.Writer, siblings string, lc *l18n.Localizer) {
//line views/nav.qtpl:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:42
	streamsiblingHyphae(qw422016, siblings, lc)
//line views/nav.qtpl:42
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:42
}

//line views/nav.qtpl:42
func siblingHyphae(siblings string, lc *l18n.Localizer) string {
//line views/nav.qtpl:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:42
	writesiblingHyphae(qb422016, siblings, lc)
//line views/nav.qtpl:42
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:42
	return qs422016
//line views/nav.qtpl:42
}

//line views/nav.qtpl:44
func StreamSubhyphae(qw422016 *qt422016.Writer, subhyphae string, lc *l18n.Localizer) {
//line views/nav.qtpl:44
	qw422016.N().S(`
`)
//line views/nav.qtpl:45
	if strings.TrimSpace(subhyphae) != "" {
//line views/nav.qtpl:45
		qw422016.N().S(`
<section class="subhyphae">
	<h2 class="subhyphae__title">`)
//line views/nav.qtpl:47
		qw422016.E().S(lc.Get("ui.subhyphae"))
//line views/nav.qtpl:47
		qw422016.N().S(`</h2>
	<nav class="subhyphae__nav">
		<ul class="subhyphae__list">
		`)
//line views/nav.qtpl:50
		qw422016.N().S(subhyphae)
//line views/nav.qtpl:50
		qw422016.N().S(`
		</ul>
	</nav>
</section>
`)
//line views/nav.qtpl:54
	}
//line views/nav.qtpl:54
	qw422016.N().S(`
`)
//line views/nav.qtpl:55
}

//line views/nav.qtpl:55
func WriteSubhyphae(qq422016 qtio422016.Writer, subhyphae string, lc *l18n.Localizer) {
//line views/nav.qtpl:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:55
	StreamSubhyphae(qw422016, subhyphae, lc)
//line views/nav.qtpl:55
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:55
}

//line views/nav.qtpl:55
func Subhyphae(subhyphae string, lc *l18n.Localizer) string {
//line views/nav.qtpl:55
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:55
	WriteSubhyphae(qb422016, subhyphae, lc)
//line views/nav.qtpl:55
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:55
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:55
	return qs422016
//line views/nav.qtpl:55
}

//line views/nav.qtpl:57
func streamcommonScripts(qw422016 *qt422016.Writer) {
//line views/nav.qtpl:57
	qw422016.N().S(`
`)
//line views/nav.qtpl:58
	for _, scriptPath := range cfg.CommonScripts {
//line views/nav.qtpl:58
		qw422016.N().S(`
<script src="`)
//line views/nav.qtpl:59
		qw422016.E().S(scriptPath)
//line views/nav.qtpl:59
		qw422016.N().S(`"></script>
`)
//line views/nav.qtpl:60
	}
//line views/nav.qtpl:60
	qw422016.N().S(`
`)
//line views/nav.qtpl:61
}

//line views/nav.qtpl:61
func writecommonScripts(qq422016 qtio422016.Writer) {
//line views/nav.qtpl:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:61
	streamcommonScripts(qw422016)
//line views/nav.qtpl:61
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:61
}

//line views/nav.qtpl:61
func commonScripts() string {
//line views/nav.qtpl:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:61
	writecommonScripts(qb422016)
//line views/nav.qtpl:61
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:61
	return qs422016
//line views/nav.qtpl:61
}
