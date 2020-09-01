// Code generated by qtc from "http_readers.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/http_readers.qtpl:1
package templates

//line templates/http_readers.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/http_readers.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/http_readers.qtpl:1
func StreamHistoryHTML(qw422016 *qt422016.Writer, hyphaName, tbody string) {
//line templates/http_readers.qtpl:1
	qw422016.N().S(`
		<main>
			<nav>
				<ul>
					<li><a href="/page/`)
//line templates/http_readers.qtpl:5
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:5
	qw422016.N().S(`">Hypha</a></li>
					<li><a href="/edit/`)
//line templates/http_readers.qtpl:6
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:6
	qw422016.N().S(`">Edit</a></li>
					<li><a href="/text/`)
//line templates/http_readers.qtpl:7
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:7
	qw422016.N().S(`">Raw text</a></li>
					<li><b>History</b></li>
				</ul>
			</nav>
			<table>
				<thead>
					<tr>
						<th>Time</th>
						<th>Hash</th>
						<th>Message</th>
					</tr>
				</thead>
				<tbody>
				`)
//line templates/http_readers.qtpl:20
	qw422016.N().S(tbody)
//line templates/http_readers.qtpl:20
	qw422016.N().S(`
				</tbody>
			</table>
		</main>
`)
//line templates/http_readers.qtpl:24
}

//line templates/http_readers.qtpl:24
func WriteHistoryHTML(qq422016 qtio422016.Writer, hyphaName, tbody string) {
//line templates/http_readers.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/http_readers.qtpl:24
	StreamHistoryHTML(qw422016, hyphaName, tbody)
//line templates/http_readers.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line templates/http_readers.qtpl:24
}

//line templates/http_readers.qtpl:24
func HistoryHTML(hyphaName, tbody string) string {
//line templates/http_readers.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/http_readers.qtpl:24
	WriteHistoryHTML(qb422016, hyphaName, tbody)
//line templates/http_readers.qtpl:24
	qs422016 := string(qb422016.B)
//line templates/http_readers.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/http_readers.qtpl:24
	return qs422016
//line templates/http_readers.qtpl:24
}

//line templates/http_readers.qtpl:26
func StreamRevisionHTML(qw422016 *qt422016.Writer, hyphaName, naviTitle, contents, tree, revHash string) {
//line templates/http_readers.qtpl:26
	qw422016.N().S(`
		<main>
			<nav>
				<ul>
					<li><a href="/page/`)
//line templates/http_readers.qtpl:30
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:30
	qw422016.N().S(`">Hypha</a></li>
					<li><a href="/edit/`)
//line templates/http_readers.qtpl:31
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:31
	qw422016.N().S(`">Edit</a></li>
					<li><a href="/text/`)
//line templates/http_readers.qtpl:32
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:32
	qw422016.N().S(`">Raw text</a></li>
					<li><a href="/history/`)
//line templates/http_readers.qtpl:33
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:33
	qw422016.N().S(`">History</a></li>
					<li><b>`)
//line templates/http_readers.qtpl:34
	qw422016.E().S(revHash)
//line templates/http_readers.qtpl:34
	qw422016.N().S(`</b></li>
				</ul>
			</nav>
			<article>
				<p>Please note that viewing binary parts of hyphae is not supported in history for now.</p>
				`)
//line templates/http_readers.qtpl:39
	qw422016.N().S(naviTitle)
//line templates/http_readers.qtpl:39
	qw422016.N().S(`
				`)
//line templates/http_readers.qtpl:40
	qw422016.N().S(contents)
//line templates/http_readers.qtpl:40
	qw422016.N().S(`
			</article>
			<hr/>
			<aside>
				`)
//line templates/http_readers.qtpl:44
	qw422016.N().S(tree)
//line templates/http_readers.qtpl:44
	qw422016.N().S(`
			</aside>
		</main>
`)
//line templates/http_readers.qtpl:47
}

//line templates/http_readers.qtpl:47
func WriteRevisionHTML(qq422016 qtio422016.Writer, hyphaName, naviTitle, contents, tree, revHash string) {
//line templates/http_readers.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/http_readers.qtpl:47
	StreamRevisionHTML(qw422016, hyphaName, naviTitle, contents, tree, revHash)
//line templates/http_readers.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line templates/http_readers.qtpl:47
}

//line templates/http_readers.qtpl:47
func RevisionHTML(hyphaName, naviTitle, contents, tree, revHash string) string {
//line templates/http_readers.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/http_readers.qtpl:47
	WriteRevisionHTML(qb422016, hyphaName, naviTitle, contents, tree, revHash)
//line templates/http_readers.qtpl:47
	qs422016 := string(qb422016.B)
//line templates/http_readers.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/http_readers.qtpl:47
	return qs422016
//line templates/http_readers.qtpl:47
}

// If `contents` == "", a helpful message is shown instead.

//line templates/http_readers.qtpl:50
func StreamPageHTML(qw422016 *qt422016.Writer, hyphaName, naviTitle, contents, tree string) {
//line templates/http_readers.qtpl:50
	qw422016.N().S(`
		<main>
			<nav>
				<ul>
					<li><b>Hypha</b></li>
					<li><a href="/edit/`)
//line templates/http_readers.qtpl:55
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:55
	qw422016.N().S(`">Edit</a></li>
					<li><a href="/text/`)
//line templates/http_readers.qtpl:56
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:56
	qw422016.N().S(`">Raw text</a></li>
					<li><a href="/history/`)
//line templates/http_readers.qtpl:57
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:57
	qw422016.N().S(`">History</a></li>
				</ul>
			</nav>
			<article>
				`)
//line templates/http_readers.qtpl:61
	qw422016.N().S(naviTitle)
//line templates/http_readers.qtpl:61
	qw422016.N().S(`
				`)
//line templates/http_readers.qtpl:62
	if contents == "" {
//line templates/http_readers.qtpl:62
		qw422016.N().S(`
					<p>This hypha has no text. Why not <a href="/edit/`)
//line templates/http_readers.qtpl:63
		qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:63
		qw422016.N().S(`">create it</a>?</p>
				`)
//line templates/http_readers.qtpl:64
	} else {
//line templates/http_readers.qtpl:64
		qw422016.N().S(`
					`)
//line templates/http_readers.qtpl:65
		qw422016.N().S(contents)
//line templates/http_readers.qtpl:65
		qw422016.N().S(`
				`)
//line templates/http_readers.qtpl:66
	}
//line templates/http_readers.qtpl:66
	qw422016.N().S(`
			</article>
			<hr/>
			<form action="/upload-binary/`)
//line templates/http_readers.qtpl:69
	qw422016.E().S(hyphaName)
//line templates/http_readers.qtpl:69
	qw422016.N().S(`"
			      method="post" enctype="multipart/form-data">
				<label for="upload-binary__input">Upload new binary part</label>
				<br>
				<input type="file" id="upload-binary__input" name="binary"/>
				<input type="submit"/>
			</form>
			<hr/>
			<aside>
				`)
//line templates/http_readers.qtpl:78
	qw422016.N().S(tree)
//line templates/http_readers.qtpl:78
	qw422016.N().S(`
			</aside>
		</main>
`)
//line templates/http_readers.qtpl:81
}

//line templates/http_readers.qtpl:81
func WritePageHTML(qq422016 qtio422016.Writer, hyphaName, naviTitle, contents, tree string) {
//line templates/http_readers.qtpl:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/http_readers.qtpl:81
	StreamPageHTML(qw422016, hyphaName, naviTitle, contents, tree)
//line templates/http_readers.qtpl:81
	qt422016.ReleaseWriter(qw422016)
//line templates/http_readers.qtpl:81
}

//line templates/http_readers.qtpl:81
func PageHTML(hyphaName, naviTitle, contents, tree string) string {
//line templates/http_readers.qtpl:81
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/http_readers.qtpl:81
	WritePageHTML(qb422016, hyphaName, naviTitle, contents, tree)
//line templates/http_readers.qtpl:81
	qs422016 := string(qb422016.B)
//line templates/http_readers.qtpl:81
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/http_readers.qtpl:81
	return qs422016
//line templates/http_readers.qtpl:81
}