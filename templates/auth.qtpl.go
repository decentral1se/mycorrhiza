// Code generated by qtc from "auth.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/auth.qtpl:1
package templates

//line templates/auth.qtpl:1
import "github.com/bouncepaw/mycorrhiza/user"

//line templates/auth.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/auth.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/auth.qtpl:3
func StreamLoginHTML(qw422016 *qt422016.Writer) {
//line templates/auth.qtpl:3
	qw422016.N().S(`
<main>
	<section>
	`)
//line templates/auth.qtpl:6
	if user.AuthUsed {
//line templates/auth.qtpl:6
		qw422016.N().S(`
		<h1>Login</h1>
		<form method="post" action="/login-data" id="login-form" enctype="multipart/form-data">
			<p>Use the data you were given by an administrator.</p>
			<fieldset>
				<legend>Username</legend>
				<input type="text" required autofocus name="username" autocomplete="on">
			</fieldset>
			<fieldset>
				<legend>Password</legend>
				<input type="password" required name="password" autocomplete="on">
			</fieldset>
			<p>By submitting this form you give this wiki a permission to store cookies in your browser. It lets the engine associate your edits with you. You will stay logged in until you log out.</p>
			<input type="submit">
			<a href="/">Cancel</a>
		</form>
	`)
//line templates/auth.qtpl:22
	} else {
//line templates/auth.qtpl:22
		qw422016.N().S(`
		<p>Administrator of this wiki have not configured any authorization method. You can make edits anonymously.</p>
		<p><a href="/">← Go home</a></p>
	`)
//line templates/auth.qtpl:25
	}
//line templates/auth.qtpl:25
	qw422016.N().S(`
	</section>
</main>
`)
//line templates/auth.qtpl:28
}

//line templates/auth.qtpl:28
func WriteLoginHTML(qq422016 qtio422016.Writer) {
//line templates/auth.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/auth.qtpl:28
	StreamLoginHTML(qw422016)
//line templates/auth.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line templates/auth.qtpl:28
}

//line templates/auth.qtpl:28
func LoginHTML() string {
//line templates/auth.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/auth.qtpl:28
	WriteLoginHTML(qb422016)
//line templates/auth.qtpl:28
	qs422016 := string(qb422016.B)
//line templates/auth.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/auth.qtpl:28
	return qs422016
//line templates/auth.qtpl:28
}

//line templates/auth.qtpl:30
func StreamLoginErrorHTML(qw422016 *qt422016.Writer, err string) {
//line templates/auth.qtpl:30
	qw422016.N().S(`
<main>
	<section>
	`)
//line templates/auth.qtpl:33
	switch err {
//line templates/auth.qtpl:34
	case "unknown username":
//line templates/auth.qtpl:34
		qw422016.N().S(`
		<p class="error">Unknown username.</p>
	`)
//line templates/auth.qtpl:36
	case "wrong password":
//line templates/auth.qtpl:36
		qw422016.N().S(`
		<p class="error">Wrong password.</p>
	`)
//line templates/auth.qtpl:38
	default:
//line templates/auth.qtpl:38
		qw422016.N().S(`
		<p class="error">`)
//line templates/auth.qtpl:39
		qw422016.E().S(err)
//line templates/auth.qtpl:39
		qw422016.N().S(`</p>
	`)
//line templates/auth.qtpl:40
	}
//line templates/auth.qtpl:40
	qw422016.N().S(`
		<p><a href="/login">← Try again</a></p>
	</section>
</main>
`)
//line templates/auth.qtpl:44
}

//line templates/auth.qtpl:44
func WriteLoginErrorHTML(qq422016 qtio422016.Writer, err string) {
//line templates/auth.qtpl:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/auth.qtpl:44
	StreamLoginErrorHTML(qw422016, err)
//line templates/auth.qtpl:44
	qt422016.ReleaseWriter(qw422016)
//line templates/auth.qtpl:44
}

//line templates/auth.qtpl:44
func LoginErrorHTML(err string) string {
//line templates/auth.qtpl:44
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/auth.qtpl:44
	WriteLoginErrorHTML(qb422016, err)
//line templates/auth.qtpl:44
	qs422016 := string(qb422016.B)
//line templates/auth.qtpl:44
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/auth.qtpl:44
	return qs422016
//line templates/auth.qtpl:44
}

//line templates/auth.qtpl:46
func StreamLogoutHTML(qw422016 *qt422016.Writer, can bool) {
//line templates/auth.qtpl:46
	qw422016.N().S(`
<main>
	<section>
	`)
//line templates/auth.qtpl:49
	if can {
//line templates/auth.qtpl:49
		qw422016.N().S(`
		<h1>Log out?</h1>
		<p><a href="/logout-confirm"><strong>Confirm</strong></a></p>
		<p><a href="/">Cancel</a></p>
	`)
//line templates/auth.qtpl:53
	} else {
//line templates/auth.qtpl:53
		qw422016.N().S(`
		<p>You cannot log out because you are not logged in.</p>
		<p><a href="/login">Login</a></p>
		<p><a href="/login">← Home</a></p>
	`)
//line templates/auth.qtpl:57
	}
//line templates/auth.qtpl:57
	qw422016.N().S(`
	</section>
</main>
`)
//line templates/auth.qtpl:60
}

//line templates/auth.qtpl:60
func WriteLogoutHTML(qq422016 qtio422016.Writer, can bool) {
//line templates/auth.qtpl:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/auth.qtpl:60
	StreamLogoutHTML(qw422016, can)
//line templates/auth.qtpl:60
	qt422016.ReleaseWriter(qw422016)
//line templates/auth.qtpl:60
}

//line templates/auth.qtpl:60
func LogoutHTML(can bool) string {
//line templates/auth.qtpl:60
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/auth.qtpl:60
	WriteLogoutHTML(qb422016, can)
//line templates/auth.qtpl:60
	qs422016 := string(qb422016.B)
//line templates/auth.qtpl:60
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/auth.qtpl:60
	return qs422016
//line templates/auth.qtpl:60
}
