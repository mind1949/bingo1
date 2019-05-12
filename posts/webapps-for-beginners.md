---
title: webapps-for-beginners
short: "Read this book at your own pace, and do exercises at your own pace.
"
date: 20.04.2017
tags:
    - webapp
    - rails

---

# Webapps for beginners

# Preface

Read this book at your own pace, and do exercises at your own pace.

<!--

- Note that we'll now start using the "new" syntax for Ruby hashes.
- Note that we'll now start omitting curly braces for hashes passed as the last argument.
  -->

# Using Libraries

Being able to stand on the shoulders of giants is one of the great advantages
when we pick a programming language that has a great ecosystem, and an even
more awesome community:

Being faced with a certain task that we want to accomplish we can often look
for other peoples' solutions, and simply re-use and build on their code.

In the chapter <a href="http://ruby-for-beginners.rubymonstas.org/bonus_2/libraries.html">Using Libraries</a>
you have learned that Ruby comes with lots of things available as soon as your
program starts, but others need to be loaded using the method `require`. For
example we did `require "digest"` and from then on there's a class
`Digest::SHA2` defined, which provides some handy tools (methods) that we
really wouldn't want to write ourselves.

But how exactly does that work?

In order to understand how external code is loaded in maybe 99% of all Ruby
applications out there nowadays you'll need to understand the following concepts
which we'll walk through:

- The Ruby Standard Library
- Rubygems and Bundler
- The Ruby load path

## Ruby Standard Library

*Batteries included*

Basically every programming language that you'll use in practice has some kind
of standard library that is shipped with the language itself.

When you install Ruby on your computer this will also install the <a
href="http://ruby-doc.org/stdlib-2.2.2/">Ruby Standard Library</a>. This means
you can simply `require` and use the things it includes.

Its documentation isn't the most pretty website on earth, but if you look at
the "package" (library) names on the left you see things mentioned like:

- `benchmark`: tools to test how performant your code is
- `debug`: tools to make debugging code easier
- `digest`, `openssl`, `securerandom`: tools for encryption and security
- `erb`: Ruby's standard templating system
- `net/imap`, `net/pop`, `net/smtp`: stuff for sending and receiving emails
- `zlib`: tools for compressing files (you know `zip` files, don't you?)

And a lot of other things. These are fairly low-level tools, of course. You do
not see a library for "logging in via Twitter or Google". They're more like
nuts and bolts, rather than bicycles or cars, but they're super useful
nonetheless.

In order to make one of these libraries available in your code you don't have
to do anything else other than use `require`. E.g. `require "zlib"` would make the
<a href="http://ruby-doc.org/stdlib-2.2.2/libdoc/zlib/rdoc/Zlib.html">zlib
library</a> available, which means you could now use the methods `deflate` and
`inflate` in order to compress and decompress files.

## Rubygems

*Ruby's package manager*

A gem is a library that has been prepared in a way so it can be distributed
(published and downloaded) via the internet.  Such libraries are commonly
refered to as "packages", and there are many different package managers for
different purposes. You can think of it as an app store that can be used to
download specific versions of libraries, which you can then use.

A package manager is a tool that knows how to install these packages on your
system so that they are then available for use. That means when you
install a gem using Rubygems, you can then `require` and use it, just as if it
was part of the standard library.

Gems are libraries, often quite small, sometimes quite big (such as Rails), that
provide certain tools for solving certain problems, just like the libraries
contained in the Ruby Standard Library. There are tons of gems <a href="#footnote-1">[1]</a>
for tons of purposes, so you'll often hear someone saying *"Hah, there's a Gem
for that!"* If you found the example of a library that helps to "sign in to a
web application via Google" hilarious, there actually is a
<a href="https://rubygems.org/gems/google-oauth/versions/0.0.2">gem for that</a>
:)

So, how do you use this?

Ruby has a built-in command line tool `gem`, which also is installed alongside
your Ruby installation, and it allows you to manage gems on your computer.

When you run `gem list` in your terminal you should see a list of all the
gems that are installed on your computer (for the currently selected Ruby
version if you use a Ruby version manager, such as RVM).

In order to install a certain gem you can run `gem install [the-gem-name]`.
E.g. `gem install middleman` would install the
<a href="https://middlemanapp.com/">Middleman</a> library, which is a super
handy tool for generating static web pages. This book is published using
Middleman.

Where does `gem` fetch all these gems (packages) from though?

Ruby gems are centrally hosted on <a href="https://rubygems.org">RubyGems.org</a>,
and Middleman, for example, has
<a href="https://rubygems.org/gems/middleman">an entry on this site</a>, too.
You can see the latest version number of this gem (which is `4.0.0.beta.2` as
of this writing), who the authors are, useful links to their homepage, source
code, documentation, and so on.

You can also see that the gem Middleman depends on a variety of other gems,
such as coffee-script, compass, execjs, and haml. This means that the authors
of Middleman themselves make use of code which in turn is packaged as gems,
too. This is a very common thing to do. We say *"Middleman's dependencies are
coffee-script, compass, ..."*

When you run `gem install middleman` you'll see that this not only installs
the Middleman gem, but also all of its dependencies, and all dependencies
that any of those dependencies might have. This means, when you install one
gem you'll get all the other stuff that this gem needs, too.

Also, each of these dependencies comes with a specification of a version number
or range of version numbers. For example `~> 2.2.0` means *"allow any version
of this gem that starts with 2.2"*.

Once installed on your computer, you can use a gem in your code in just the same
way as you use something from the Ruby Standard Library: You `require` it.

For example, in order to configure Middleman to generate this book we require
a gem called `middleman-toc`
<a href="https://github.com/rubymonsters/ruby-for-beginners/blob/master/config.rb#L1">here</a>.
This is an extension to Middleman that allows us to add a table of contents, and
we need to `require` it before Middleman can use it.

### Footnotes:

<a name="footnote-1">[1]</a> *As of this writing, there are <a href="https://rubygems.org/gems">6373 gems</a>
hosted on RubyGems.org.*

## Bundler

*Sandboxes of Gems*

Consider an application that relies on lots and lots of gems:

For example, a newly generated Rails application comes with a whopping 42 gems,
and this number usually grows pretty quickly with typical Rails applications.
The application <a href="http://speakerinnen.org/">Speakerinnen Liste</a>, at the time of this writing, installs no less than 138 gems. Only 40 of these
are defined in the Speakerinnen <a href="https://github.com/rubymonsters/speakerinnen_liste/blob/master/Gemfile">Gemfile</a>,
which are the ones that provide certain features used by the application.
The other 98 gems are dependencies that these 40 gems have.

Now, imagine working on an application like this over years. There are new
versions of gems being published all the time. Often they'll update the version
numbers of their dependencies when new versions come out, or add new gems as
dependencies, and remove old ones.

How likely do you think it is that any of the version number ranges specified
for all of these gems would conflict with any other version number range? In
practice it is very likely. Rubygems itself (the library `rubygems` that is part
of your Ruby installation) is not very smart at figuring out which versions
of certain gems play well with each other.

On top of this, with plain Ruby and Rubygems, when you `require` any gem in your
code, you'll simply get the latest version of this gem that is already installed
on your computer. Assuming these versions do not conflict with each other, you
still don't know if these are the exact same versions as the ones your
co-workers have installed. Again, in practice this is pretty unlikely.

All of this means that maintaining the right versions of the right gems on
your system manually would be a sheer nightmare. And if you get it wrong then
that can be a source of many, often subtle bugs that you don't really want to
deal with.

And that's where Bundler comes to the rescue:

Bundler allows you to define which gems your application depends on (in a file
called `Gemfile`. Here's the one we're using for
<a href="https://github.com/rubymonsters/ruby-for-beginners/blob/main/Gemfile">this book</a>).
When we run `bundle install`, this will figure out which
gem versions work well with each other (a task that Bundler is *great* at), and
store the solution to this riddle to a separate file (called `Gemfile.lock`).

The `Gemfile` and `Gemfile.lock` files are part of your code, and can be shared
with other developers. When they download your code, and run `bundle install`
themselves, they'll get exactly the same gem versions that you also have.

You can think of Rubygems as a tool to install gems on your computer. Over time
this may result in a collection of lots and lots of gems in various versions
that all sit somewhere on your filesystem.

Bundler on the other hand is a tool for picking *some* of these gem versions,
and restricting access to only these. You can think of it as a sandbox of the
few gems that your application really should use. Like a looking glass that
restricts the vision of your application to only see these few gem versions,
even though there maybe tons of other gem versions installed on your computer.

In order to use your application with Bundler, you'd prepend the command `bundle
exec` to whatever other command you execute in your terminal. Imagine you'd
normally execute your program like this:

```
ruby my_amazing_app.rb
```

In order to use it with Bundler and restrict the visible gem versions to the
ones defined in your `Gemfile.lock` file, you would run this instead:

```
bundle exec ruby my_amazing_app.rb
```

<p class="hint">
For Rails applications you do not have to prepend <code>bundle exec</code>, as
Rails does this itself, under the hood.
</p>

## Ruby load path

*Where to look for all the things*

Ruby is software, and software stores things somewhere on your file system.
In order to define places on your computer where interesting stuff is stored,
software often has the concept of a "load path".

If you are using a Unix based operating system such as Linux or Mac OS X, you
may have seen the environment variable `$PATH` in installation instructions.
This variable defines all the directories where executable files are stored.

Ruby has a load path, too. Inside your Ruby program you can print it out using:

```ruby
puts $LOAD_PATH
```

This will print out the array that is defined as the `$LOAD_PATH` when Ruby
starts your program. `puts` is smart enough to put each string in that array
on a separate line.

Each of these lines represents a directory on your computer where Ruby files
are stored. If you use `require` anywhere in your application (e.g. `require
"digest"`) then Ruby will look for a Ruby file with the same name (e.g.
`digest.rb`) in each of these directories. It will load the first file with
this name that it can find.

If you are curious, you can quickly check the default load path of your
Ruby installation like this:

```
ruby -e 'puts $LOAD_PATH'
```

The `-e` flag is a way to run some Ruby code without having to store it in
a file.

For me, this prints:

```
/Users/sven/.rbenv/versions/2.2.1/lib/ruby/site_ruby/2.2.0
/Users/sven/.rbenv/versions/2.2.1/lib/ruby/site_ruby/2.2.0/x86_64-darwin14
/Users/sven/.rbenv/versions/2.2.1/lib/ruby/site_ruby
/Users/sven/.rbenv/versions/2.2.1/lib/ruby/vendor_ruby/2.2.0
/Users/sven/.rbenv/versions/2.2.1/lib/ruby/vendor_ruby/2.2.0/x86_64-darwin14
/Users/sven/.rbenv/versions/2.2.1/lib/ruby/vendor_ruby
/Users/sven/.rbenv/versions/2.2.1/lib/ruby/2.2.0
/Users/sven/.rbenv/versions/2.2.1/lib/ruby/2.2.0/x86_64-darwin14
```

From this you can see that I am using <a href="https://github.com/sstephenson/rbenv">rbenv</a>
to manage my Ruby versions, that my currently active Ruby version is `2.2.1`,
and that I am running this on Mac OS X ("darwin").  All these paths are
directories somewhere in the directory where Ruby 2.2.1 is installed on my
computer.

Whenever there's a `require "something"` statement in some Ruby code that I run
on my computer, Ruby will check all these directories for a file
`something.rb`.

Now let's do a few exercises on
<a href="/exercises/using_rubygems.html">Rubygems</a> and
<a href="/exercises/using_bundler.html">Bundler</a>.

# HTML

<a href="http://en.wikipedia.org/wiki/HTML">HTML</a>, meaning "HyperText Markup
Language", is a document format used for defining the semantic structure of a
single web page.  One could say that HTML is what the internet is made of: All
the websites that we are looking at every day are all defined (described) as
HTML.

HTML was first proposed by <a href="http://en.wikipedia.org/wiki/Tim_Berners-Lee">Tim Berners-Lee</a>
in 1989, and layed grounds for the World Wide Web's huge breakthrough during
the 1990s. It is defined by the <a href="http://en.wikipedia.org/wiki/World_Wide_Web_Consortium">W3C</a>
and it's latest version is <a href="http://en.wikipedia.org/wiki/HTML5">HTML 5</a>
which added a bunch of exciting, and useful features.

The fundamental, primary feature of any web browser, such as Firefox, Chrome,
Safari, is to render (display) HTML documents:

Whenever you type a URL into the browser's address bar (or click a link, such
as <a href="http://rubymonstas.org">http://rubymonstas.org</a>
browser will send a request to this address (i.e. to some application running
on some computer that responds to this address), and it will (in most cases)
get an HTML document back as a response, which it will display to you.

What is meant by "semantic structure" though?

Maybe it is best to look at an example first. This is a fairly simple, but
valid HTML document:

```html
<html>
  <head>
    <title>Ruby Monstas HTML Example</title>
  </head>
  <body>
    <h1>Ruby Monstas HTML Example</h1>
    <p>One paragraph of text.</p>
    <p>Another paragraph of text, containing an emphasized <em>word</em>.</p>
    <h2>A list of items</h2>
    <ul>
      <li>First item</li>
      <li>Second item</li>
      <li>Third item</li>
    </ul>
  </body>
</html>
```

If you look at this document you'll notice the recurring pattern of "tags" that
start with `<something>` and then are closed with `</something>`. E.g. the
entire document starts with an opening `<html>` tag, and ends with a closing
`</html>` tag.

HTML entirely consists of these tags that have a certain meaning, can be nested,
and contain content.

Here's what the HTML tags used in this example mean:

- `<html>...</html>` - the HTML document as a whole
- `<head>...</head>` - the header of the document, containing meta information (i.e. information *about* the document, not *part of* the document itself)
- `<title>...</title>` - an example of one bit of meta information, the title of the page as displayed in your browser history, and the browser window title (or tab)
- `<body>...</body>` - the body of the document itself, i.e. the whole of its content
- `<h1>...</h1>`, `<h2>...</h2>` - a headline level 1, and level 2, containing the headline's text. HTML defines heading levels 1-6, which should be enough to define the structure even of large documents.
- `<p>...</p>` - a single paragraph, containing the paragraph's text.
- `<ul>...</ul>` - an unordered list (i.e. a list that uses bullet points, as opposed to, e.g., a numbered list)
- `<li>...</li>` - a single list item, must be contained in either a `<ul>` or `<ol>` tag

You can read `<html>` as *"Ok, let's start an HTML document!"*.
Then `<head>` says: *"Oh, btw, here's some extra meta information about the document ..."*
and `<title>` says: *"The title of this document is: [...]"*.
`</title>` meaning *"Ok, we're done with the title"*,
just like `</head>` meaning *"Ok, we're done with the meta information part."*

When the browser sees `<body>` it reads that as *"Ok, now here comes the real content part of the document."*

Inside that content section it first finds an `<h1>` tag, and reads *"Aha, the
main headline is [...]"*, so it will display this as a headline to you. It then
finds two `<p>` tags containing a few bits of plain text, so it displays two
paragraphs. The last paragraph has an emphasized word, so it makes that one
cursive (or something else, depending on style definitions defined elsewhere).
Next it finds a 2nd level headline `<h2>`, and then an "unordered list" `<ul>`,
i.e. a list with bullet points.

It is important to understand that HTML only defines the **semantic structure**
of a document. It does, as such, say nothing about the visual representation of,
for example, a headline, the spacing between paragraphs, the font family and size
used, any colors, borders, or even element placement: The visual representation
is defined in a different language called <a href="http://en.wikipedia.org/wiki/Cascading_Style_Sheets">CSS</a>
(or by the browser's defaults, should there not be any custom style definition
for this web page, as in our example above).

Why is the semantic structure of a document useful at all, even disregarding
its visual styling when presented in the browser window?

Maybe the simplest example is a link. In order to describe a link in text (HTML
is stored as plain text) we'll need to tell three pieces of information to the
browser:

- That we'd like to define a link,
- what text to display for the link, and
- what other web address to link to.

This is a valid example of a link in HTML, using the tag `a` (which means
"anchor"):

```html
<a href="http://rubymonstas.org">Ruby Monstas Homepage</a>
```

This HTML, when rendered in a web browser, looks like this:
<a href="http://rubymonstas.org">Ruby Monstas Homepage</a>, i.e. there's a link
with the text "Ruby Montas Homepage" and it links to the target URL ("href")
http://rubymonstas.org.

Does this make sense?

This is why HTML is called a "markup language". It is a language that "marks
up" certain each part of the content with its structural meaning, such as
*"This is a link, with this target URL"*, or *"This is a heading with this
level"*, *"This is a simple paragraph"*, and so on.

If you think about it, there are lots of useful applications, that leverage
the information that is provided by defining the **semantic structure** of a
document (as opposed to its visual representation).

For example, one could apply different "themes", as in styles, to the same HTML
page depending on people's preferences, or depending on the device this website
is displayed on (e.g. desktop browser windows versus mobile devices).

But also, the entire success of Google as a search engine relied on the fact
that the semantic structure of a document provides clues about the relevance of
certain search terms in this document: If a word is contained in the URL, the
document title, the toplevel headlines, then it's probably important. Even more
so, if other pages link to this document using link texts that also contain
that word, then this adds social proof to the mix, and search results should
probably list this page higher up in the results.

Another example is that, by inspecting the headline tags in an HTML document,
we can auto-generate a table of contents for this document quite easily,
without having authors maintain this manually. Many content management
systems (CMS) do this, including Wikipedia.

# Embedded Ruby

HTML is the main format that fuels the world wide web.

Web applications are applications that run somewhere on the internet, and that
web browsers talk to. E.g. if you check your Twitter or Facebook feed your
browser talks to the respective applications that are run by these companies.

So one could say that the primary purpose of web applications is to produce
HTML so it can be sent to web browsers.

Because generating HTML dynamically based on some data (e.g. the current user's
name, and their tweets, or emails) is such an important concern of applications
Ruby has some built-in support for making this easier: a library called
<a href="http://ruby-doc.org/stdlib-2.2.2/libdoc/erb/rdoc/ERB.html">ERB</a>,
short for "embedded Ruby".

## ERB Templates

The main idea behind ERB is to embed some Ruby code *into* an HTML document
(also called a template). <a href="#footnote-1">[1]</a>

Here's an example:

```erb
<html>
  <body>
    <h1>Messages for <%= name %></h1>
    <ul>
      <% messages.each do |message| %>
        <li><%= message %></li>
      <% end %>
    </ul>
  </body>
</html>
```

Can you guess what this means?

Everything inside the so called ERB tags `<% ... %>` is considered Ruby code.
Everything outside of them is just some static text, in our case HTML code,
into which the results of the Ruby code will be embedded whenever the ERB tag
also has an equals sign, as in `<%= ... %>`.

Imagine stripping everything outside the ERB tags, and the opening and closing
tags themselves from the code above. And imagine replacing the `=` equals sign
with `puts` statements. You'd then end up with this code:

```ruby
puts name

messages.each do |message|
  puts message
end
```

That's some code you understand, right?

ERB, when executed, does exactly this, except that `=` as part of the ERB tag
`<%= ... %>` will not output things to the terminal, but capture it, and insert
it to the surounding text (HTML code, in our case) in place of this tag.

Ruby code in ERB tags that do not have an equal sign, such as `<% ... %>` will
be executed, but any return values won't be captured, and just discarded.

### Footnotes:

<a name="footnote-1">[1]</a>
*This idea actually predates Ruby's ERB library and
became popular with PHP, a language that originally was meant to be used
exactly this way: by embedding some code into an HTML template file.*

## Rendering ERB

The Ruby code embedded in our template uses two local variables (or methods)
with the names `name` and `messages`. Also, `each` is called on `messages`, so
this should probably be an array:

```ruby
puts name

messages.each do |message|
  puts message
end
```

So how do we provide these objects to the ERB template? And how do we execute
the whole thing?

It's probably best to look at an example:

```ruby
require "erb"

template = %(
  <html>
    <body>
      <h1>Messages for <%= name %></h1>
      <ul>
        <% messages.each do |message| %>
          <li><%= message %></li>
        <% end %>
      </ul>
    </body>
  </html>
)

name = "Ruby Monstas"
messages = [
  "We meet every Monday night at 7pm",
  "We've almost completed the beginners course!",
  "Feel free to come by and join us!"
]

html = ERB.new(template).result(binding)
puts html
```

Does this code make sense to you?

Let's walk through it:

- On the first line we require the erb from the
  <a href="http://ruby-doc.org/stdlib-2.2.2/libdoc/erb/rdoc/ERB.html">Ruby Standard Library</a>.
- As you've learned in <a href="http://ruby-for-beginners.rubymonstas.org/bonus_1/alternative-syntax.html">this chapter</a>
  the syntax `%(something)` defines a string. So, `template` is just one, long
  string that contains our ERB template.
- Next we define two local variables `name` and `messages`, which hold a simple
  string, as well as an array with 3 strings.
- On the next line we create an instace of the class `ERB` and pass our
  `template` (the string defined earlier) to it.
- On this instance we then call the method `result` with something that is
  called `binding`.
- This method call to `result` returns something that we assign to the variable
  `html`, so, as you might guess, this should be the HTML we were after.
- We'll then just output the result to the terminal on the last line, using `puts`.

If you run this code it will output something like this:

```html
<html>
  <body>
    <h1>Messages for Ruby Monstas</h1>
    <ul>
      <li>We meet every Monday night at 7pm</li>
      <li>We've almost completed the beginners course!</li>
      <li>Feel free to come by and join us!</li>
    </ul>
  </body>
</html>
```

... which is a valid HTML document that a browser would render (display) like
this:

![38A2B041-42CE-4CF2-B562-E6F44BC8D0A1](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.3IK7Ul/38A2B041-42CE-4CF2-B562-E6F44BC8D0A1.png)
Looks good?

To recap, all that our code above does is the following:

- Define a template with some HTML and some embedded ERB tags.
- Define a couple of variables that contain things that our template wants (`name` and `messages`).
- Create an instance of the class ERB, passing this template.
- On this instance, call the method `result`.

In other words, it executes (we say "renders") the ERB template using the
`name` and `messages` objects, and returns the HTML as a result.

 Bindings

Now, what's this `binding` thing used in our code?

```ruby
html = ERB.new(template).result(binding)
```

Obviously we do not define a local variable, so it needs to be a method. Let's
google that: <a href="http://www.google.com/?q=Ruby+binding">Ruby binding</a>

The first result goes to the <a href="http://ruby-doc.org/core-2.2.0/Binding.html">Ruby documentation</a>,
saying (some rather cryptic stuff stripped for our purpose):

*"Objects of class Binding encapsulate the execution context at some particular
place in the code and retain this context for future use. The variables,
methods, [...] that can be accessed in this context are all retained. Binding
objects can be created using Kernel#binding [...]. These binding objects can be
passed [around]."*

Hmmmmm, let's re-read that, and think about it.

First of all, <a href="http://ruby-doc.org/core-2.2.2/Kernel.html">Kernel</a>
is something we have never mentioned so far. It is a
<a href="http://ruby-for-beginners.rubymonstas.org/bonus_2/modules">Ruby module</a>
that is included into the class `Object`. That means that, whenever you create
any object, this module, and thus all of its methods, will be included.  Since
Ruby's top level scope is also an object, `binding` is defined there too. This
is also, by the way, the secret reason why methods like `p`, `puts`, and so
on are available everywhere: they're defined in `Kernel`.

Now, what's this execution context the documentation is talking about?

Remember when we talked about <a href="http://ruby-for-beginners.rubymonstas.org/methods/scopes.html">scopes</a>
in the context of methods? That's the same as the Ruby documentation means by
"execution context".  It is that empty room or space that Ruby enters whenever
it starts your program, or enters a method. If you define a *local* variable,
then this variable will be visible within this scope, or room, or "execution
context", but not outside of it.

Now the `binding` keeps exactly *this knowledge*: what variables are defined,
and what objects they are referring to. And the `binding` itself is an object
that can be passed around.

If you find this confusing don't worry:

All of this is knowledge that many Ruby programmers don't need, or at least
very rarely use themselves. Advanced programming techniques aside (like "meta
programming") rendering ERB templates is the one single situation where you'd
ever need it. And, on top of this, as a Sinatra or Rails developer you also
never need it because both Sinatra and Rails ship tools that hide this weird
stuff from you.

For now, the one thing you can remember is that by calling `binding` and
passing the result to the ERB instance, you simply pass *access* to the two
local variables `name` and `messages`, so they can be used inside your ERB
template.

Does this make sense?

Then maybe it's a good time for an
<a href="/exercises/mailbox_erb.html">exercise on ERB</a>.

# HTTP

*Just text messages*

When computers talk to each other via the internet all they do is send around
messages.

You can think of these as simple text messages: They may be encrypted, they may
have some binary content attached to them (such as an image, or a movie), but
essentially they're just text messages.

HTTP is the protocol that browsers and web applications speak in order to
know how exactly to formulate these messages. If you're not familiar with the
term "protocol", it means a clearly specified language, a set of formal rules
for how to talk to each other.

When you open your browser, type a URL into the address bar, and hit return,
your browser will send a text message to another computer associated
with the hostname that is part of your URL. The computer on the other side
will receive this message (called the request), and send another message back
(called the response).

For example, when you go to
<a href="http://rubymonstas.org/location.html">http://rubymonstas.org/location.html</a>
this is (a simplified version of) the request that is going to be sent by your
browser:

```http
GET /location.html HTTP/1.1
Host: rubymonstas.org
User-Agent: Mozilla/5.0
```

In essence, this message says *"Get me the page /location.html"*. Simple as
that.

The webserver that hosts our homepage then looks up the content of that page,
and sends another message back to your browser containing the content of
the page, which is, of course, HTML:

```http
HTTP/1.1 200 OK
Content-Type: text/html

<html>
  </head>
    <title>Ruby Monstas - Where we meet</title>
  <head>
  ...
</html>
```

That's it. Your browser will extract the HTML that is sent as the body of the
message and display it.

Let's look at the request and response in more detail. What does all of this
stuff mean exactly?

## An HTTP Request

HTTP defines that the first line of a request must contain three bits: a
**verb**, a **resource**, and the version of the protocol. Out of these, we
only really care about the first two bits. The version is only relevant in way
more advanced contexts.

The **verb**, also called the "method", defines an operation. `GET` of course
means just that: *"Get me this thing, please."* The "thing" we're looking for
is defined by the second bit on this line: the **resource**, also called the
path. In our case that's a web page, i.e. some HTML.

Other common verbs are <a href="#footnote-1">[1]</a>

- `POST` means (in modern web applications, such as Rails applications):
  *"Create a new instance of the resource, using this data."* (E.g. `POST` to
  `/users` means: create a new user.)
- `PUT` means: update the instance with this new data. E.g. `PUT` to `/users/1`
  means: for the user (with the id) `1`, update their attributes with the given
  data.
- And `DELETE` means what you think it does. E.g. `DELETE` to `/users/1` means:
  Delete the user (with the id) `1`.

HTTP also defines that after this first line there may come any number of
**headers**, i.e. key/value pairs, containing meta information about the request.

In our example these are the hostname that was used in the URL, and the
(simplified) name of the browser that was used to send the request. Both of
these actually are mostly for information, and not a core part of the request.

## An HTTP Response

Now, let's have a look at the response.

Again, HTTP defines that on the first line there need to be three bits of
information: The HTTP version number, a **status code** number, and a status
description. And again, we can just ignore the version number.

The status code is a way for the responding server to tell our browser what
*type of response* this is. Here are some common examples <a href="#footnote-2">[2]</a>:

- `200 OK` means: Everything's cool, here's the thing you were looking for.
- `301 Moved Permanently` means: The thing you are looking for does not exist
  here any more. Here's the place where it has been moved to, look over there.
- `404 Not Found` means: Dunno, this thing does not exist.
- `500 Internal Server Error` means: Ouch, I have run into an internal error on
  my side.

The status code indicates the type of response, and most of the time,
for most `GET` requests, it will be `200 OK`.

On the following lines, again, there will be some **headers**: key/value pairs
that contain meta information about the response. In our case there is a header
that says the content type of the body is `text/html` (which just means it is
HTML).

Now, when there's a **body** ("main content") on a request (sometimes requests
come with a body, too) or response (as in our example), then it needs to be
separated from the header with an extra, blank line. This tells the browser
that we're done with the headers, and the rest of the message is the response
body (the actual content).

## Summary

To summarize:

- A request has, on its first line, a verb (method) and a resource (path). A
  response has, on its first line, a status code and description.
- Both requests and responses can have any number of header key/value pairs
  that contain meta information.
- And both requests and responses can include a body, which needs to be
  separated from the headers with an extra, blank line.

There is, of course, a *lot* more to say about HTTP, but these are the
essentials of how computers talk to each other when they speak HTTP, as
browsers and webservers do. And these also are the basics of what you need
to know in order to build a simple web application, receive requests, and
send back responses.

If you are interested in more details, here is a
<a href="https://speakerdeck.com/rkh/http-rubymonsters-edition">presentation</a>
that Konstantin Haase, maintainer of Sinatra, gave for Ruby Monstas in early

1. Also, the
   <a href="http://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol">Wikipedia page about HTTP</a>
   is an interesting read.

## Footnotes:

<a name="footnote-1">[1]</a> There's a lot more to be said about HTTP verbs than
this, and some developers might argue that these explanations are wrong. Let's
say these verbs mean these things in applications like Rails applications.
The HTTP standard defines these verbs in more abstract, and generic terms,
while Rails has its own, more practical definition of these HTTP verbs.

E.g. `POST` actually means more "do whatever you think is right" in terms
of the HTTP specification, while in the context of Rails it means "create this
thing".

For us, in practice, for the time being, it is enough to remember that `GET`
means "get", `POST` means "create", `PUT` means "update", and `DELETE` means
"delete".

<a name="footnote-2">[2]</a> There are lots of other status codes that servers
can respond with, here's <a href="http://en.wikipedia.org/wiki/List_of_HTTP_status_codes">a list</a>.
Many of them are very rarely used, and maybe the funniest one is `418 I'm a
teapot`.

# Rack

<a href="http://rack.github.io/">Rack</a> is the most basic way to build a very
simple web application in Ruby, and we are going to start with it.

In practice you may never build an application with Rack directly, but you'll use
it when you work on Sinatra or Rails applications, since these use Rack under
the hood.

## Your first Rack app

Let's jump right in.

In a new directory `rack` create a file `config.ru` with the following content:

```ruby
class Application
  def call(env)
    status  = 200
    headers = { "Content-Type" => "text/html" }
    body    = ["Yay, your first web application! <3"]

    [status, headers, body]
  end
end

run Application.new
```

Now, make sure you have the gem Rack installed: In your terminal check `gem
list rack`. Does that show something like `rack (1.6.1)` (or any other version
number)? If it doesn't, install Rack with the command `gem install rack`.

This gem comes with a little executable (command line program) called `rackup`.
This command looks for a file `config.ru` in the current directory, and starts
a web server using it, on your local computer.

Make sure you have `cd`ed to your `rack` directory, and then run `rackup`. You
should see something like:

```
$ rackup
[2015-05-15 18:37:42] INFO  WEBrick 1.3.1
[2015-05-15 18:37:42] INFO  ruby 2.2.1 (2015-02-26) [x86_64-darwin14]
[2015-05-15 18:37:42] INFO  WEBrick::HTTPServer#start: pid=17588 port=9292
```

Of course the version numbers may be different, but the important bit that
you want to look for is the port. In our case that's `9292`.

Now your web application has started you can point your browser to
<a href="http://localhost:9292">http://localhost:9292</a>
something like this:

![87F7BCAB-71FD-41C6-B722-34613988BAAA](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.AOdl1L/87F7BCAB-71FD-41C6-B722-34613988BAAA.png)

Pretty cool, isn’t it? With just a few lines of simple Ruby code you have just
written an actual web application, and started a fully functional web server
with it.

Now, let's have a closer look at the code. Here's our class again:

```ruby
class Application
  def call(env)
    status  = 200
    headers = { "Content-Type" => "text/html" }
    body    = ["Yay, your first web application! <3"]

    [status, headers, body]
  end
end
```

We define a class `Application`, and, on the last line, create an instance
of it, which we pass to the method `run`. The method `run` is defined by
Rack, and expects to be passed *something* that responds to the method `call`.
<a href="#footnote-1">[1]</a>

That's why we defined a method `call` on our class. This method takes one
argument `env`. It does not use the `env` (whatever that is), yet, but instead
just returns the same static array whenever it is called.

This array contains 3 things:

- The number 200, which represents the status code,
- a hash that contains a single header (the content type), and
- an array containing a single string, which is the body.

So the method `call` returns something that represents an HTTP response in Rack!

Rack makes it so that whenever there's a request coming in (on the computer
that is `localhost`, i.e. your own, local computer, and on the port `9292`),
it will turn this request into a hash `env`. It will then hand us this hash by
calling our method `call`.  I.e. the hash `env` that is passed to us as an
argument contains the request information. We'll have a look at that in a
minute.

Rack then expects us (our method `call`) to return an array containing those
three elements:

- The HTTP response code
- A hash of headers
- The response body, which must respond to each (i.e. we can just use an array)

In other words, that's also a kind of protocol (programmers also use the term
"interface" here). Rack defines how we can interact with it in a formal way in
terms of Ruby. The protocol is defined as something like:

<p class="hint">
A Rack application implements a method <code>call</code> that takes a hash
representing the request, and is supposed to return an array containing the
status code, a hash containing the headers, and an array containing the request
body.
</p>

Once Rack got these three things back from our method `call` it will create
the respective response (text) message out of it, and send it back to the
browser, so the browser can handle it (and in our case display the body).

Great!

If you've paid attention close enough you may have noticed that our little
Rack application actually is lying to the browser. Can you spot where?

Our response header hash defines a `Content-Type` header. And in that header we
claim that our response body has the content type `text/html`. But then we
return a body that isn't HTML, but just plain text. So that's wrong. Luckily
browsers are pretty forgiving. They try to do their best to still display
useful information to the user, and fix things for us.

We could fix our application by specifying that the `Content-Type` is plain
text by setting the value to `text/plain`. But instead we can also simply
turn the body into HTML like so:

```ruby
class Application
  def call(env)
    status  = 200
    headers = { "Content-Type" => "text/html" }
    body    = ["<html><body><h1>Yay, your first web application! <3</h1></body></html>"]

    [status, headers, body]
  end
end

run Application.new
```

### Footnotes:

<a name="footnote-1">[1]</a> Most examples for Rack applications will use
a `Proc` or `lambda`, which can be called using their method `call`. Here's an
example, using a lambda:

```ruby
application = lambda do |env|
  [200, { "Content-Type" => "text/html" }, ["Yay, your first web application! <3"]]
end

run application
```

This is why the author of Rack picked `call` as the main method: Our web
application will be "called" by Rack, and so it can be just an anonymous `Proc`
or `lambda`. Pretty slick.

However, in our example we use a class, so we can add more methods to it later.

## The Rack Env

Let's have a look at the `env` data that is passed along with the request.
Let's just print it out to the terminal as follows:

```ruby
  def call(env)
    p env
    [200, { "Content-Type" => "text/html" }, ["Yay, your first web application! <3"]]
  end
```

In order for the server to pick up this change we need to restart it. Go to
your terminal window where `rackup` (WEBrick) is running our app, and hit
`ctrl-c`.  Then hit `cursor-up` to get the last command back (or type
`rackup`), and hit `return`.

If you now refresh the page in your browser (hit `cmd-r` or `ctrl-r` depending
on your operating system) you should then see ... wow, quite a bit of messy
output in the logs in your terminal. For me it looks like this (some less
interesting bits removed):

```
{"GATEWAY_INTERFACE"=>"CGI/1.1", "PATH_INFO"=>"/", "QUERY_STRING"=>"", "REMOTE_ADDR"=>"127.0.0.1", "REM
OTE_HOST"=>"localhost", "REQUEST_METHOD"=>"GET", "REQUEST_URI"=>"http://localhost:9292/", "SCRIPT_NAME"
=>"", "SERVER_NAME"=>"localhost", "SERVER_PORT"=>"9292", "SERVER_PROTOCOL"=>"HTTP/1.1", "SERVER_SOFTWAR
E"=>"WEBrick/1.3.1 (Ruby/2.2.1/2015-02-26)", "HTTP_HOST"=>"localhost:9292", "HTTP_ACCEPT_LANGUAGE"=>"en
-US,en;q=0.8,de;q=0.6", "HTTP_CACHE_CONTROL"=>"max-age=0", "HTTP_ACCEPT_ENCODING"=>"gzip", "HTTP_ACCEPT
"=>"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8", "HTTP_USER_AGENT"=>"Mo
zilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.1
35 Safari/537.36", "rack.version"=>[1, 3], "rack.url_scheme"=>"http", "HTTP_VERSION"=>"HTTP/1.1", "REQU
EST_PATH"=>"/"}
```

Woha! What does all of this stuff mean?

First of all, you can see that it's a Ruby hash with lots of keys that are all
strings. Then, most of these strings are all upper case with an underscore `_`
as a word separator. Some of these start with `HTTP_`, while others don't.
However, there also are some keys that start with `rack`, and are all
lowercase.

Rack uses a simple convention for these keys:

- All headers that have been part of the actual HTTP request are prefixed with
  `HTTP` and uppercased. For example the request header `host: localhost:9292`
  will be translated to the hash key `HTTP_HOST` with the value `localhost:9292`.
  I.e. these are the actual HTTP headers that our browser has sent.
- All other uppercase keys represent additional information that has been
  passed (added) from the webserver that has received the request (in this case
  WEBrick, which runs our little Rack application). For example, WEBrick adds
  the key `PATH_INFO` with the resource (path), as well as the key
  `REQUEST_METHOD` with the verb (method) from the HTTP request. These weren't
  headers in the request, but obvioulsy part of it. On top of this, WEBrick
  also adds other things, such as the `SERVER_SOFTWARE` key (telling us which
  WEBrick and Ruby version we are using), and so on.
- All keys that are prefixed with `rack.` represent internal additions that
  Rack adds.

Let's write a little bit of code to make this easier for us to inspect:

```ruby
class Application
  def call(env)
    puts inspect_env(env)
    [200, { "Content-Type" => "text/html" }, ["Yay, your first web application! <3"]]
  end

  def inspect_env(env)
    puts format('Request headers', request_headers(env))
    puts format('Server info', server_info(env))
    puts format('Rack info', rack_info(env))
  end

  def request_headers(env)
    env.select { |key, value| key.include?('HTTP_') }
  end

  def server_info(env)
    env.reject { |key, value| key.include?('HTTP_') or key.include?('rack.') }
  end

  def rack_info(env)
    env.select { |key, value| key.include?('rack.') }
  end

  def format(heading, pairs)
    [heading, "", format_pairs(pairs), "\n"].join("\n")
  end

  def format_pairs(pairs)
    pairs.map { |key, value| '  ' + [key, value.inspect].join(': ') }
  end
end
```

Again, after changing your code, you'll need to restart your server application.

For me this outputs the following (again, stripping some of the less interesting bits):

```
Request headers

  HTTP_HOST: "localhost:9292"
  HTTP_REFERER: "http://localhost:9292/"
  HTTP_ACCEPT_LANGUAGE: "en-US,en;q=0.8,de;q=0.6"
  HTTP_ACCEPT_ENCODING: "gzip"
  HTTP_USER_AGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36"
  HTTP_ACCEPT: "*/*"
  HTTP_VERSION: "HTTP/1.1"

Server info

  GATEWAY_INTERFACE: "CGI/1.1"
  PATH_INFO: "/"
  QUERY_STRING: ""
  REMOTE_ADDR: "127.0.0.1"
  REMOTE_HOST: "localhost"
  REQUEST_METHOD: "GET"
  REQUEST_URI: "http://localhost:9292/"
  SCRIPT_NAME: ""
  SERVER_NAME: "localhost"
  SERVER_PORT: "9292"
  SERVER_PROTOCOL: "HTTP/1.1"
  SERVER_SOFTWARE: "WEBrick/1.3.1 (Ruby/2.2.1/2015-02-26)"
  REQUEST_PATH: "/"

Rack info

  rack.version: [1, 3]
  rack.url_scheme: "http"
```

Now, that's way easier to read, right?

Luckily, we can just ignore most of these things.

At the moment, the only interesting keys for us are `REQUEST_METHOD` and
`PATH_INFO`: They're the relevant bits from the HTTP request.

<p class="hint">
The most interesting bits in the <code>env</code> hash are the
<code>REQUEST_METHOD</code>, and <code>PATH_INFO</code>.
</p>

Ok, let's do something with them.

## Method and Path

We saw that the `env` hash that Rack passes to the method `call` contains the
keys `REQUEST_METHOD` and `PATH_INFO`.

Let's modify our app a little so we can make use of it:

```ruby
class Application
  def call(env)
    handle_request(env['REQUEST_METHOD'], env['PATH_INFO'])
  end

  private

    def handle_request(method, path)
      if method == "GET"
        get(path)
      else
        method_not_allowed(method)
      end
    end

    def get(path)
      [200, { "Content-Type" => "text/html" }, ["You have requested the path #{path}, using GET"]]
    end

    def method_not_allowed(method)
      [405, {}, ["Method not allowed: #{method}"]]
    end
end
```

Reading the code closely, do you understand what it does, and why?

We have changed the method `call` to extract the values for the keys
`REQUEST_METHOD` and `PATH_INFO`. And we then pass these two values to a new
method `handle_request`, which checks the request `method`. Keep in mind that
`method` here is just a variable name that refers to the HTTP concept of a
"request method".  This, of course, is not the same as a Ruby method, it's just
a variable that will hold a string such as `GET`, `POST` depending on the HTTP
request.

- If `method` is `GET`, then we call another method `get`, passing the `path`.
  The method `get` complies with Rack's convention for returning a response: It
  returns an array that has the response status, headers, and a body. We've
  changed the body a little bit so it displays the `path` that was requested.
- If `method` is not `GET`, then we call another method `method_not_allowed`.
  This method also complies with Rack's convention, but returns a different
  response. This time we use the status code `405` which means exactly this:
  *"Method Not Allowed"*. Our little application just does not support any other
  methods.

Because these response arrays are the return values of these two methods,
they'll also be the return value of the method `handle_request`, and it turn
the method `call`. So they'll be passed back to Rack, and turned into the
actual HTTP response that is returned to your browser.

If you restart your server, and point your browser to
[http://localhost:9292/ruby/monstas](http://localhost:9292/ruby/monstas)
you should now see something like this:

![15F4DCC5-6EA4-49DA-BD7E-A78947ACFE23](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.V5KQNI/15F4DCC5-6EA4-49DA-BD7E-A78947ACFE23.png)

Congratulations!

You have just written your first web application that responds to different
requests with (albeit only slightly) different responses.

Imagine working on this application more, and returning different HTML pages
based on the `path` that is part of the request: You could use the ERB
rendering method from the [previous chapters about ERB](/erb.html) in order to
render different ERB templates.

And guess what, this is exactly what [Sinatra](/sinatra.html) makes super easy :)

In case you are wondering how to test the response for request methods other
than `GET`, that's not so easy without knowing more about [HTML forms](/forms.html).
However, if your computer has `curl` installed (a commandline tool to execute
HTTP requests) you can try this:

```
curl -i -X POST http://localhost:9292
```

For me this outputs:

```
HTTP/1.1 405 Method Not Allowed
```

As well as a bunch of other response headers.

# Sinatra

<a href="http://www.sinatrarb.com/">Sinatra</a> is a quite popular, small
framework for writing web applications. It sits somewhere between Rack and
Rails in terms of how many features it provides, what kind of things it
makes easy for developers, and what kind of things one has to implement on
your own.

We'll use Sinatra for introducing quite a few concepts that are common to
all web applications, before we later move on to learning how to
build a Rails appplication.

## Sinatra versus Rails

Why do we start with Sinatra, and not Ruby on Rails? How do the two compare?

You could compare Sinatra to a bicycle, and Rails to a car. Both do
the job of getting you from A to B.

Sinatra is much more lightweight, needs less resources, and does fewer things
out of the box. Rails on the other hand is packed with features, comes with a
ton of code, and makes it very easy to build complicated web applications in
limited time, if you know how to use it.

With Rails every little piece of code in your application has a precisely
defined place. So when you look at any other Rails application you know exactly
what's were, and you can get started immediately. Sinatra on the other hand defines
almost nothing for you. People come up with their own structures, and different
Sinatra applications can look very different.

Of course, because it's lightweight, Sinatra also is much faster than Rails,
and this is where our little metaphor ends fairly quickly :) Rack, by the way,
would be a pair of shoes. Sinatra and Rails both use Rack under the hood.

So which one is better for a beginner?

Both are great :) Sinatra is great for learning the basics of HTTP and routing.
Rails on the other hand is better for learning how to use a database and
actually store things, and learning the concept of a "resource".

Also, next to Sinatra and Rails there are, of course, lots of other frameworks
for building web applications. For example, <a href="http://hanamirb.org/">Hanami</a>
has gained popularity quite quickly recently.

## Domain specific languages

Sinatra, on its homepage, does not call itself a framework. Instead it calls
itself a DSL, which is quite a common term in the Ruby world.

So let's talk about that for a moment, too.

DSL is short for <a href="http://en.wikipedia.org/wiki/Domain-specific_language">"Domain specific language"</a>.
"Domain" in this case refers to the "problem domain", i.e. the "problem at
hand", or rather the context of the problem. The domain where a solution or
tool can be applied.

What does that mean?

When you think about a hammer as a concept, then the domain it is relevant to
is "building physical things". In contrast, it is completely irrelevant to the
domain of mathmatics. Likewise, the concept of an operating system, is
something that is relevant in the domain of using computers, while it is
entirely irrelevant in the domain of baking pizza.

Applications are built to solve problems in a certain domain. In a commercial
context the domain often is what a business' customers care about.

Consider a book shop application, such as Amazon, back when it still did
nothing else but selling books. The domain of this application is the entirety
of concepts that their users have in mind, and that they care about when they
use it. In our example the domain would include concepts like: books,
categories, a shopping cart, orders, payment methods, delivery addresses, and
so on.

A domain specific language is a language that includes terms to speak about
these concepts: *"Books can be placed into a shopping cart."* or *"A shopping cart
can be checked out, which will place an order."*

In the context of Ruby code the term "domain specific language" is used to
describe a piece of code or library that provides classes and methods that
allow us to "speak about them", or implement them, in the form of code.

The problem domain that Sinatra lives in is building web applications. And web
applications which "speak" HTTP with browsers.

It therefor has methods like, for example, `get`, `post`, `put`, and `delete`.
You can use these methods in order to describe how your application responds to
HTTP requests. It also has methods like `headers`, `session`, `cookies`, and
other things that relate to concepts from HTTP.

So, instead of writing verbose code like this:

```ruby
def handle_request(method, path)
  if method == "GET"
    [200, { "Content-Type" => "text/html" }, ["You have requested the path #{path}, using GET"]]
  else
    [405, {}, ["Method not allowed: #{method}"]]
  end
end
```

Sinatra allows us to write code like this:

```ruby
get "/some/path" do
  "You have requested the path /some/path"
end

post "*" do
  status 405
end
```

As you can see this code uses a "language" (i.e. methods provided by Sinatra)
that is specific to the domain HTTP.

Does that make sense?

<p class="hint">
The term DSL is used for libraries that allow you to write descriptive,
narrative Ruby code that "speaks" about the solution to a problem using
terms that are specific to the given problem domain.
</p>

## Your first Sinatra app

Let's get started looking at Sinatra.

Make sure you have the Sinatra gem installed. Use `gem list sinatra` to
check if it's there. If it's not install it using `gem install sinatra`.

Now, let's steal the intro example from their homepage, and adopt it. Make a
new directory `sinatra`, `cd` into it, create a file `monstas.rb` and add this
code:

```ruby
require "sinatra"

get "/" do
  "OMG, hello Ruby Monstas!"
end
```

Now you can run your little app using `ruby monstas.rb`. You should see
something like this:

```
$ ruby monstas.rb
[2015-05-15 21:37:41] INFO  WEBrick 1.3.1
[2015-05-15 21:37:41] INFO  ruby 2.2.1 (2015-02-26) [x86_64-darwin14]
== Sinatra (v1.4.6) has taken the stage on 4567 for development with backup from WEBrick
[2015-05-15 21:37:41] INFO  WEBrick::HTTPServer#start: pid=27182 port=4567
```

Again, there are lots of version numbers, that we can ignore, and it
also tells us the port that it's running on. This time it's `4567`. For some
reason Sinatra finds it important to use a different port number `¯\_(ツ)_/¯`

So let's point the browser to <a href="http://localhost:4567">http://localhost:4567</a>

You should see something like this:

![B74656CC-A90E-4490-AE30-A9BC5E05F3FC](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.VZcwsy/B74656CC-A90E-4490-AE30-A9BC5E05F3FC.png)

That  was easy, wasn't it.

If you've read the chapters about [Rack](/rack.html) it is interesting to know
that Sinatra uses Rack under the hood, but it deals with the nitty gritty
details of looking at the `REQUEST_METHOD` and `REQUEST_PATH` for you.

It allows you to use the methods `get`, `post`, `put`, and `delete`
with a path argument, and simply specify a block that will be called whenever a
request matches the request method (verb) and path.

The Ruby code that makes up our little Sinatra application reads very well,
and it only focusses on the bits and pieces we care about (as opposed to our
Rack application, which had to include the knowledge about the `env` etc.)

Sinatra also allows you to simply return a string from this block, which will
then be used as the response body, and sets things like the status code and
headers for you (when it returns the Rack style response array to Rack). Since
the vast majority of requests will want to return 200 as a status code Sinatra
just assumes you want that too, unless you specify something else.

## Routes

In Sinatra calls to methods like `get`, `post`, `put`, and `delete` are called
*routes*. They take a path, and a block that handles a request.

What's up with that term?

The request is being picked up by the webserver, and then *routed* to a piece
of code that handles the request, and specifies the response. They're like the
info desk at a shopping mall: You (as an HTTP request) walk up to them, and
tell them `GET /something`. They'll route (send) you to a certain shop on a
certain floor where you can find what you're after.

So, this would be three Sinatra routes:

```ruby
get "/" do
  "OMG, hello Ruby Monstas!"
end

get "/signup" do
  "Here you can sign up for the next beginners course."
end

post "/signin" do
  # do something to sign in the user
end
```

When a request comes in Sinatra will look at the request method and path, and
match it against the first route:

- If it matches it will run (call) the block and return a response to the
  browser.
- If it does not match, it will look at the next route, and so on.
- If no route matches then Sinatra responds with a `404 Not Found`.

## Params

Now, sometimes the request path contains dynamic data.

For example, the path of the URL <a href="https://rubygems.org/gems/rack">https://rubygems.org/gems/rack</a>
is `/gems/rack`. The path of the details page for the gem Sinatra on
RubyGems.org is `/gems/sinatra`, the path for Middleman is `/gems/middleman`,
and so on.

Obviously we don't want to hardcode ("write out literally") all these names in
our application code: We don't want to change our code for each and every new
gem that is added: At the time of this writing RubyGems.org has 122,037.
Instead we want to be able to express *"a path that starts with `/gems`
followed by another, second segment"*.

In Sinatra we can do this by specifying a pattern as a path. Sinatra will then
match the pattern against the path, and see if it applies.

Let's try that out.

Add the following route (request handler) to your program, at the end of the
file:

```ruby
get "/monstas/:name" do
  "Hello #{params["name"]}!"
end
```

Restart your Sinatra application, and point your browser to
<a href="http://localhost:4567/monstas/monstas">http://localhost:4567/monstas/monstas</a>

How does this work?

`params` is a hash that Sinatra makes available for you in your route blocks, and
it will automatically include relevant data from the request.

In our case our route specifies a path that is a pattern: the last part of the
path starts with a colon `:`. This tells Sinatra that we'd like to accept any
string here, and that we'd like to call this string `name`.

Sinatra therefore adds the key `"name"` to the `params` hash, and sets the
given string from the path (i.e. from the URL) to it.

When you point your browser to the URL
<a href="http://localhost:4567/monstas/Elizabeth">http://localhost:4567/monstas/Elizabeth</a>
your application will say *"Hello Elizabeth!"*, when you go to
<a href="http://localhost:4567/monstas/Juliane">http://localhost:4567/monstas/Juliane</a>
your application will say *"Hello Juliane!"*, and so on.

Let's inspect the params hash, and return this string as the response body:

```ruby
get "/monstas/:name" do
  params.inspect
end
```

If you restart your application, and reload the page in your browser, then it
should display something like this:

```
{"splat"=>[], "captures"=>["monstas"], "name"=>"monstas"}
```

So this confirms that `params` is a hash, and the key `"name"` has the value
`"monstas"` set.  `splat` and `captures` are for building more complicated
routes, and we can ignore these for now.

This is pretty cool.

The `params` hash can contain more than matches from the URL. You'll later see
that it also contains any data sent from HTML forms as part of the HTTP
request. As well as any query params that can be part of the URL (separated
with a question mark `?`).

But for now it's good to know that Sinatra adds matches from the path pattern
to the `params` hash.

## Rendering templates

So far, our application does not actually return HTML, it returns just plain
text.

Let's fix that.

For this we'll want to re-use what we've learned about rendering ERB templates.
In order to use ERB we need to require it, define an ERB template, and make any
variables used in the template known as local variables in our route:

```ruby
require "sinatra"
require "erb"

get '/monstas/:name' do
  ERB.new("<h1>Hello <%= params[:name] %></h1>").result(binding)
end
```

This code is familiar to you, isn't it?

We get use `params[:name]` in the template because `params` is "known" in the
scope that is passed as part of the `binding`. The rest is just the same
as in our examples in the chapters about <a href="/erb.html">ERB</a>

If you restart the server, and reload the page in your browser, it should now
look like this:

![FB50B34D-9B06-4CF6-8830-3EEE9A3349B8](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.WFMDi1/FB50B34D-9B06-4CF6-8830-3EEE9A3349B8.png)

Awesome, we're now serving HTML, not just plain text.

However, Sinatra also has built-in support for ERB templates ("views"). We can
achieve exactly the same thing without spelling out the
`ERB.new(...).result(binding)` noise:

```ruby
get '/monstas/:name' do
  erb "<h1>Hello <%= name %></h1>", { :locals => { :name => params[:name] } }
end
```

I.e. Sinatra has a method `erb` that hides all the details of rendering the
template from us, and also accepts a template.

On top of this, it also accepts a hash that allows us to specify various
options. If we specify a key `:locals` and give it another hash, then Sinatra
will make each key/value pair on this hash available as *local* variables (thus
"locals") in our ERB template.

Of course, since `params` already is a hash, and it already has the key `name`
defined, we can also just say:

```ruby
get '/monstas/:name' do
  erb "<h1>Hello <%= name %></h1>", { :locals => params }
end
```

Nice, isn't it.

## Layouts

Our HTML still isn't quite valid: `<html>` and `<body>` tags are mandatory,
even though browsers still render our HTML happily without them.

We could add these wrapping tags to each and every one of our templates,
every time we create a new one. However, that's quite some repetition, and
should we ever want to change anything about them (which is likely) we'd
have to change all of our templates.

For building web applications it is handy to have "wrapping" templates. And
that's something Sinatra supports, too. They're called "layout" templates.

Here's how it works:

```ruby
get '/monstas/:name' do
  template = "<h1>Hello <%= name %></h1>"
  layout   = "<html><body><%= yield %></body></html>"
  erb template, { :locals => params, :layout => layout }
end
```

Don't forget to restart your application and refresh the page. If you inspect
the source code of the web page (right click on the page, and select "View Page
Source", or whatever that's called in your browser) you'll see that it now has
the `<html>` and `<body>` tags from our layout template.

The only real change that we have made is that we've added the `:layout` option
to the options hash. Because this is too much stuff to fit on one line we have
also sticked the template and layout into local variables.

As you can see the rendered "content" (from our main template) is being wrapped
by, or inserted into, our layout template.

What about that `yield` thing in our layout template though?

`yield` is a keyword in Ruby that calls a block that was given to a method.
Whenever you pass a block to a method (such as `each`, `collect`, `select`, and
so on) this method can then call the block by using the keyword `yield`.
Because we've never implemented a method that took a block we've also never
discussed this keyword.

Hmmmmm. Ok, that's how it works under the hood, yes. However, in this context,
all you need to remember is that, in a layout template, `<%= yield %>` marks
the place where the other template (the one that is being wrapped) is supposed
to be inserted.

Does this make sense?

Imagine you have an application with, say, 10 `get` routes. Each of them
renders a different template. Say, there's one for the homepage, one for
a user signup page, one for a user profile page, and so on. Each of these
templates is supposed to be wrapped into the same layout, which has the
enclosing `<html>`, `<body>`, and other tags, which are all common to each
of these pages. Maybe there's also a common header menu at the top, and a
common footer at the end of the page.

Each route will then render its own template, and specify the layout template
to be used, which will replace the `<%= yield %>` tag with the template, and
wrap it.

That's pretty handy.

## Template files

So far we've defined our templates as strings right inside our route. That
worked well because our templates were ridiculously small. Of course, any
real application will have much bigger templates. Managing these inside
our routes would get pretty messy pretty quickly.

It is, therefore, better to store them in separate files. And again, Sinatra
has built-in support for that: when we call the method `erb` with a symbol
instead of a string, then Sinatra assumes this is part of a filename, and
it will look for a template file in a directory `views`.

Let's create a new directory `views`, and add a file `monstas.erb`, containing
our template:

```erb
<h1>Hello <%= name %></h1>
```

Also, add a file `layout.erb` with our layout template:

```erb
<html><body><%= yield %></body></html>
```

Now we are ready to change our route as follows:

```ruby
get '/monstas/:name' do
  erb :monstas, { :locals => params, :layout => :layout }
end
```

Restart your application, and reload the page. You should see the same result.

But our code looks much better this way, doesn't it?

Interestingly, we don't even need the name the layout. Sinatra looks for
this filename by default (we could specify a different name though, in case
we need different layout templates in different contexts):

```ruby
get '/monstas/:name' do
  erb :monstas, { :locals => params, :layout => true }
end
```

And finally, we can also even totally omit the option, because Sinatra
assumes we want a layout and finds one in the `views` directory:

```ruby
get '/monstas/:name' do
  erb :monstas, { :locals => params }
end
```

If we don't want a layout, for some reason, then we can pass `:layout => false`
instead.

Neat.

## Using instance variables

So far we've passed data to our templates using the `:locals` option key which
holds a hash.

Sinatra supports a second way of passing data, which uses instance variables.
We mention this mostly because this is also the "Rails way" of passing data to
your templates (views).

Let's change our template to use an instance variable `@name`, like so:

```erb
<h1>Hello <%= @name %></h1>
```

If we now assign the same instance variable in our route, then Sinatra will
make it available to the template, too:

```ruby
get '/monstas/:name' do
  @name = params["name"]
  erb :monstas
end
```

This also is a little bit more concise, and spares a few keystrokes.

So, which way is the better one?

On one hand there's an argument that using the `:locals` way is the cleaner,
and "right" way of doing it: These two objects (our route, and the template)
should be separated clearly, and not simply share things.  On the other hand
using instance variables is much more common due to the fact that Rails
encourages it.

As always, you should just use whatever feels better to you, and maybe ask your
friends and fellow developers for their opinions and reasons.

# Forms

*Talking back*

So far, all that our little application can do is return some HTML when asked.
I.e. it can respond to `GET` requests by returning HTML.

What if we want to talk back to our application, and send some data from the
browser to it? You have seen this a million times, there are those little forms
on websites that let you enter data and somehow submit it.

So, how does that work?

## HTML Form Tags

HTML defines a couple tags for describing forms that are part of a document
(web page).

Here's how a simple form looks in HTML:

```html
<form>
  <input type="text">
  <input type="submit">
</form>
```

If you copy this HTML code to a file, and open the file in your browser it will
look similar to this:

![CFE77E90-6798-4DC8-B840-2C15B4086C49](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.Q5D7Ea/CFE77E90-6798-4DC8-B840-2C15B4086C49.png)

Nice.

You can see that the two `<input>` tags are rendered (displayed) in different
ways because they have two different types: `text` and `submit`. One is a text
input field, and one is a button to submit the form. In modern HTML there are
lots of other input element types. You can find a
<a href="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input">full list here</a>.

Let's make our little Sinatra application serve this form.

Make a new template file `monstas.erb` in your `view` directory and add the the
HTML above. Then replace your `get '/monstas/:name'` route with the following:

```ruby
get "/monstas" do
  erb :monstas
end
```

Restart the server and go to <a href="http://localhost:4567/monstas">http://localhost:4567/monstas</a>
You should see something similar to our screenshot above.

Now, this form isn't very useful, yet. If you submit it (enter something, and
click "Submit") nothing much will happen. It clears the input element, and
appends a question mark to the URL for some reason, but that's it.

What's happening here?

## Submitting Data

When you click a button that is a "submit" input element then your browser will
submit the form that this button is part of.

But to where? If we do not specify anything else the browser will submit it to
the same URL that we are currently looking at, which is, in our case the path
`/monstas`.

We might as well specify it: this helps make clear what's going on.
Curiously, the attribute is called `action` (and not `target` or `url` as one
might expect):

```html
<form action="/monstas">
  <input type="text">
  <input type="submit">
</form>
```

Now, what does "submit" mean, exactly, in this context?

We said browsers speak HTTP when they talk to a web server (our Sinatra
application), so "submitting" a form means making another HTTP request. Again,
if we do not specify anything else, the default method will be `GET`.

So our browser makes another `GET` request with the path `/monstas`. This
basically just reloads the page, and displays the same form again, served from
our `get "/monstas"` route. This also explains why the text in the input field
goes away.

Ok, where does the question mark come from though?

When the browser submits the form it collects all the data from the form input
elements, and sends it along with the HTTP request. In the case of a `GET`
request <a href="#footnote-1">[1]</a> it will append it to the URL, after a
question mark, as name/value pairs. These are called "query parameters" in
HTTP.

However, our form input text element does not have a name, and so the browser cannot
pass it in a meaningful way.

So let's change that, and specify a name for our text input. Since we want
the user to input their name, the name of our input should be `name`:

```html
  <input type="text" name="name">
```

If you restart the application, reload the page, and again enter some text and
click submit ... you should see that the URL changes to something like
`http://localhost:4567/monstas?name=Monstas`.

Aha!

So that's how the browser passes our input to the application. It just appends
it to the URL as query parameters (name/value pairs).

Now, how can we make use of this form data in our application?

### Footnotes:

<a name="#footnote-1">[1]</a> You'll see later that form data is passed as part
of the request body in case of `POST` requests. The request body is able to
hold much more data than the URL, which is limited in its length.

## Accessing Form Data

We mentioned earlier that form data will be available in the `params` hash.

Let's check that by outputting the `params` in our route:

```ruby
get "/monstas" do
  p params
  erb :monstas
end
```

Again, restart the application, and load the URL
<a href="http://localhost:4567/monstas?name=Monstas">http://localhost:4567/monstas?name=Monstas</a>
In your terminal you should then see something like:

```
{"name"=>"Monstas"}
```

So this is how we can access the data that has been passed from the browser as
part of the HTTP request, when it submitted the form.

Let's do something with it!

Change your route like so:

```ruby
get "/monstas" do
  @name = params["name"]
  erb :monstas
end
```

And add this code at the top of your `monstas.erb` view:

```erb
<% if @name %>
  <h1>Hello <%= @name %>!</h1>
<% end %>
```

When you restart the server, and reload the page you should see something like
this:

![FF527BCE-452A-4196-99CB-41782D5733AE](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.8fxW41/FF527BCE-452A-4196-99CB-41782D5733AE.png)

Awesome!

So this is how you can pass (submit) data from the brower to your application,
and use it in some way.

Let's make one more improvement to our form. It currently always wipes out the
text that has been submitted. Let's make sure we preseve it in the input element.

We can do that by specifying the `value` attribute on the `input` tag, like so:

```html
<form action="/monstas">
  <input type="text" name="name" value="<%= @name %>">
  <input type="submit">
</form>
```

When you restart your application, and reload the page, it should now put the
name to the input element:

![32EEB2C3-18CC-4D6F-BC90-63F1C64E9BE0](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.SZFdc1/32EEB2C3-18CC-4D6F-BC90-63F1C64E9BE0.png)

Perfect.

# Storing data

So far our little application doesn’t do anything with the data that we submit,
except for displaying it on the same page right away. Once you go to another
URL that previous data is gone.

Instead, let's store it to a file. Normally, web applications would use some
kind of database for storing data, but for our purpose using a file will be
good enough.

What we want to achive is that every name ever submitted to our application
is appended to a file, let's say `names.txt`. So we'll keep a long running
list of all these names.

## Writing to a file

How do you append a string to a file in Ruby?

If you google for <a href="http://google.com?q=Ruby+file+append">"Ruby file append"</a>
you'll find a bunch of answers that all basically look like this:

```ruby
File.open(filename, "a+") do |file|
  file.puts(string)
end
```

This uses two local variables `filename` and `string`, which in our case would
be `names.txt` and the name passed as the param.

The slightly weird looking second argument `"a+"` tells the `open` method that
we want to use the file for appending something (thus `a`), and that we'd like
it to create a new file unless it already exists (thus `+`).

Also, `File.open` takes a block, and passes an object, an instance of the class
<a href="http://ruby-doc.org/core-2.2.0/File.html">File</a> to it.

It does this because the file needs to be closed once we're done with it. The
`open` method makes sure we don't forget this, and closes the file once it has
run our block.  Pretty handy.

Inside of the block we simply call `puts` on the file object, which will append
the string that we pass, and also add a newline (just like `puts` does when you
output a string to the terminal).

Alright. Ready to go?

## Storing the name

Now that we know how to append something to a file, let's use that in our
application and store all those names to a file.

We could change our route like so:

```ruby
get "/monstas" do
  @name = params["name"]

  File.open("names.txt", "a+") do |file|
    file.puts(@name)
  end

  erb :monstas
end
```

However, that stuffs a lot of clutter into our route, and we'd like to keep
these readable.

So let's extract that to method right away:

```ruby
def store_name(filename, string)
  File.open(filename, "a+") do |file|
    file.puts(string)
  end
end

get "/monstas" do
  @name = params["name"]
  store_name("names.txt", @name)
  erb :monstas
end
```

Better. Our route now describes what it does, instead of telling how exactly
it is done.

If you restart your Sinatra application, and reload the page, you should see
a file `names.txt` created in the same directory, and it should contain the
name from the form.

You can check this using command line tools like this:

```
# check if the file is there
$ ls names.txt
names.txt

# look at the content of the file
$ cat names.txt
Monstas
```

Of course you can also just look at the file in your editor :)

## Listing all names

Let's make it so that we can look at the list of the names in the browser
though. This is a web application, right.

For this we'll need to read the names from the file. Again, if you ask Google
for "Ruby read file" you'll find it's as simple as this:

```ruby
File.read(filename)
```

This returns a single, long string, which represents the content of the entire
file. Because we store every name on a new line we can split this string with
the newline character `"\n"` in order to get our names as an array:

```ruby
def read_names
  File.read("names.txt").split("\n")
end
```

However, this would break if no file with this name exists yet. So let's add
a little safeguard, and return an empty array if the file does not exist:

```ruby
def read_names
  return [] unless File.exist?("names.txt")
  File.read("names.txt").split("\n")
end
```

Does this make sense? If the file does not exist we return an empty array `[]`.
If it exists we read it, and split the content into lines. Even if the file
exists, but it is empty, we'll still get an array.

Also, let's store the names on an instance variable in our route, so we can
then use it in the template later:

```ruby
get "/monstas" do
  @name = params["name"]
  @names = read_names
  store_name("names.txt", @name)
  erb :monstas
end
```

Now we can output the names as an unordered list (`<ul>`) in our `monstas.erb`.

So let's add this at the end of your file (we now want to display the full list):

```erb
<ul>
  <% @names.each do |name| %>
    <li><%= name %></li>
  <% end %>
<ul>
```

![9670BB1B-58DC-432B-B05C-32CFF79967DC](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.s4UD7u/9670BB1B-58DC-432B-B05C-32CFF79967DC.png)

Wheeeee! Pretty cool.

The tag `ul` means "unordered list", and it is supposed to have one or many
`li` tags, which means "list item". Yeah, HTML tag names are a little weird,
their naming dates back a while.

## Using POST

Ok, our little application is already pretty cool, isn't it?

However, there's one big problem that we definitely want to fix: The HTTP
specification says that `GET` requests should be safe, and idempotent.

What does that mean? <a href="http://en.wikipedia.org/wiki/Idempotence">Wikipedia says</a>:

*"Idempotence is the property of certain operations in mathematics and computer
science, that can be applied multiple times without changing the result beyond
the initial application."*

Obviously, whenever we reload the URL our application adds the name again, and
again, and again. Since we do this in response to a `GET` request we do not
comply with the HTTP specification: we do change the result, the HTML that
is returned.

When we store, modify, or delete data in our application we also say that we
"change the state" of the application: It goes from "3 names stored" to "4
names stored".

`GET` requests should not modify the state of our application according to the
HTTP specification. They should only "get" what's already there, and not change
it. So what do we do?

The appropriate HTTP verb (request method) to use for this kind of request is
`POST`. The result of a `POST` request does not need to be idempotent, and it's
basically up to the application to decide what to do with it.  In modern
applications `POST` usually means "add this thing to the collection", where
"the collection" is defined by the path: In our case we want to add a name to
the collection `monstas`.

In order to tell the browser to send a `POST` request instead of a `GET`
request we add this as an attribute to the `<form>` tag like so:

```html
<form action="/monstas" method="post">
  <input type="text" name="name" value="<%= @name %>">
  <input type="submit">
</form>
```

When you reload the page, and try to submit the form again, you'll get
Sinatra's 404 (Not Found) page though: "Sinatra doesn’t know this ditty."

Of course!

We do not have a route for `POST` requests, yet. Remember how the HTTP verb is
a key part of the HTTP request? And Sinatra wants us to use these verbs in
order to define our route.

So let's add one, and move the logic for storing the name to the new
route:

```ruby
get "/monstas" do
  @name = params["name"]
  @names = read_names
  erb :monstas
end

post "/monstas" do
  @name = params["name"]
  store_name("names.txt", @name)
end
```

Cool!

Our `get` route is now idempotent (it does not change any state), and we also
have a `post` route for the same path, and we store the name if there is one.

However, what should we now send back to the browser in response?

For starters, we could just send a little confirmation:

```ruby
post "/monstas" do
  @name = params["name"]
  store_name("names.txt", @name)
  "Ok!"
end
```

![E8424B43-6173-4916-85F5-63DB93E4A7C0](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.kc3bZn/E8424B43-6173-4916-85F5-63DB93E4A7C0.png)

Hmmmm. Ok, this isn't too bad, we could make this a proper HTML template.

However, where would the user go next? Shouldn't we display the updated
list of all the names next?

## Using Redirects

After handling a `POST` (or `PUT` or `DELETE`) request, in today's web
applications it is common to redirect the browser to the page that displays the
changed state:

Our application would respond: *"Alright, I've done this! Go here, check it out!"*

The way of expressing this in terms of HTTP is returning the status code `303
See Other`, and adding a `location` header to the response that tells the
browser where to go (what to `GET`) next.

```http
HTTP/1.1 303 See Other
Location: http://localhost:4567/monstas
```

Hmm, ok. But how do we achieve that in Sinatra?

Remember Sinatra calls itself a DSL? It has a method `redirect` for that:

```ruby
post "/monstas" do
  @name = params[:name]
  store_name("names.txt", @name)
  redirect "/monstas"
end
```

The `redirect` method will make sure our application responds with the status
code `303` and add a `location` header to the response with the value
`http://localhost:4567/monstas` (since we have passed the path `/monstas`, but the
`location` header needs a full URL).

We also want to see the the Welcome message when we post a new name. For that to
work, we can redirect to `"/monstas?name=#{@name}"`. Our `GET` route will then
include ("render") the name into the HTML page:

```ruby
post "/monstas" do
  @name = params[:name]
  store_name("names.txt", @name)
  redirect "/monstas?name=#{@name}"
end
```

If you restart the server, go to <a href="http://localhost:4567/monstas">http://localhost:4567/monstas</a>
and submit the form you'll see that your browser will be redirected, and make
a new `GET` request to the same URL.

So this is pretty cool. We now have two routes: One for displaying (`GET`ing)
the current state of our application, and one for storing (`POST`ing) new data.

Awesome!

# Sessions

HTTP is a so called <a href="http://en.wikipedia.org/wiki/Stateless_protocol">stateless protocol</a>.
Wikipedia says:

*"In computing, a stateless protocol is a communications protocol that treats
each request as an independent transaction that is unrelated to any previous
request so that the communication consists of independent pairs of request and
response."*

Imagine talking to a person with a very shortlived memory:

- Hey, who are you? - *I'm a web application.*
- Hey, who are you? - *I'm a web application.*
- Hey, who are you? - *I'm a web application.*
- How often have I just asked you who you are? - *I don't know, I don't keep track.*

That's what's meant by HTTP being stateless: The web application just responds
to the request at hand, but has no way to identify how these requests relate to
each other.

Imagine thousands of users clicking around and using your web application. The
application would just respond to each of these requests individually, but does
not know from the protocol where these requests are coming from, or how many
people (browsers) it is talking to: There is no concept of a conversation, or
"session" in HTTP.

Basically.

However, of course there are ways to identify your users. We all know that we
can sign in to a web application, and it would recognize who we are. Right?

When you sign into Gmail it will display *your* emails, not anyone else's
emails.  Obviously it needs to know who you are, so it can find *your* emails.
The same happens for basically every useful, modern web application.

In order to identfy who is making a request web applications often use cookies
for this. There are other techniques, but using cookies is by far the most
common one.

So let's check that out.

## Cookies

A cookie is a little piece of information that a web application can send along
with a response that will be stored by the browser. From then on, when the
browser makes another request to the same application, it will include the cookie
to the request, sending it back to the application.

For example, an HTTP response that sets a cookie for a user's prefered visual
theme could look like this:

```
HTTP/1.0 200 OK
Content-type: text/html
Set-Cookie: theme=light
```

From now on, the browser would then include the cookie to subsequent requests:

```
GET /blog.html HTTP/1.1
Host: rubymonstas.org
Cookie: theme=light
```

This way the application could apply the "light" theme to the blog, because the
user has selected it in some previous request.

<p class="hint">
Cookies are a way to persist (keep) state (data) across multiple HTTP requests.
</p>

## Sessions in Sinatra

Building on top of cookies, even though HTTP itself does not have a concept of a
"session" (or conversation), Sinatra, like basically any web application
tool or framework, supports sessions.

A session is a way for a web application to set a cookie that persists
an identifier across multiple HTTP requests, and then relate these requests
to each other: If you've signed in before the web application will be able to
know that you're still the same user you've identified as a couple requests
earlier. If you've done something else before, and something has been stored
to your session, then the web application will be able to use it later.

In our example we'll persist (store) a short confirmation message across two
requests: The `POST` request will store the message, and redirect the browser
to another URL. The browser will make that `GET` request, and we'll display
the confirmation message.

Let's have a look how this works.

Say the message we'd like to pass from our `POST` route to the next request is
*"Successfully stored the name [name]"*. I.e. after storing the name to the
file, we'd like to pass a message to the `GET` request that the browser is
going to be redirected to later.

In order to make this work we first need to enable the `:sessions` feature in
Sinatra. You can do that by adding this line above of your routes:

```ruby
enable :sessions
```

Now, we want to store the message in our `post` route:

```ruby
post "/monstas" do
  @name = params["name"]
  store_name("names.txt", @name)
  session[:message] = "Successfully stored the name #{@name}."
  redirect "/monstas?name=#{@name}"
end
```

Ok, cool.

The `session` looks like a simple Ruby hash, but if we store something to it
then Sinatra will set a cookie for us. It does so by sending a `Set-Cookie`
header along the reponse. This header will have a long, messy looking,
encoded string as a value.

In my browser it looks like this:

```
Set-Cookie: rack.session=BAh7CUkiD3Nlc3Npb25faWQGOgZFVEkiRWI4OTdhMDJlNDBkMDFlNjcxNWUw%0AZGI1ZWU5MzQ0YTQyMjAzYjFiZTE2YzYxNzgwMWQxYjI3NzhiOWNhYTQ4YzUG%0AOwBGSSIJY3NyZgY7AEZJIiU2ZjdjN2Y0ZmM0MTdmMGJkNjBkNmY5MmQ1NDEx%0ANGQ4ZgY7AEZJIg10cmFja2luZwY7AEZ7B0kiFEhUVFBfVVNFUl9BR0VOVAY7%0AAFRJIi03NGNlNDIxYTczNjMwZDY3MWViNTlkYzIzN2YyN2M5NGU3ZWU4NTRm%0ABjsARkkiGUhUVFBfQUNDRVBUX0xBTkdVQUdFBjsAVEkiLTA3NjBhNDRjMzU0%0AODIxMzJjZjIyNDQyYTBkODhjMDhiYjg1NTYyNTAGOwBGSSIIZm9vBjsARkki%0ACGJhcgY7AFQ%3D%0A; path=/; HttpOnly
```

Wow. Ok, the name of the cookie gives us a hint that this is a session, and it
is managed by Rack, which is what Sinatra uses under the hood to persist the
session.

Luckily we do not need to understand how exactly it does this. All we need to
know is that we can now use this data in the next request (the `GET` request)
like so:

```ruby
get "/monstas" do
  @message = session[:message]
  @name = params["name"]
  @names = read_names
  erb :monstas
end
```

I.e. we grab `:message` from our session, and stuff it into the instance variable `@message`.
Doing so we can then disply it in our view:

```erb
<% if @message %>
  <p><%= @message %></p>
<% end %>
```

Let's try it out. Restart your application, and go to <a href="http://localhost:4567/monstas">http://localhost:4567/monstas</a>
If you enter a name, and click submit you should then see something like this:

![952E85B9-0D50-4D9F-B4B8-E6B7FAE8DADC](/var/folders/2k/q63b1kcj38j4cfrl_fjx6z1c0000gn/T/net.shinyfrog.bear/BearTemp.xHOHvc/952E85B9-0D50-4D9F-B4B8-E6B7FAE8DADC.png)

How does this work?

In our `post` route we store the message to the session hash. This is
something Sinatra provides to us as developers. When we enable this
feature Sinatra will, after every request, store this hash to a cookie
with the name `rack.session`, in the encoded form that you saw above.

We say the hash is being <a href="http://en.wikipedia.org/wiki/Serialization">serialized</a>,
which is a fancy way of saying it is turned into some kind of format that
can be stored in text form. Sinatra (actually, Rack, under the hood) then also
encrypts and signs this data, so it is safe to send it over the internet (in
case we keep any sensitive data in it). Thus hackers cannot easily tamper
with it, it is a shared secret between our web application (Sinatra) and us
(our browser).

Ok, so the `post` route includes the `Set-Cookie` header with this session
cookie in its response, and sends it to the browser. The browser will, from
now on, pass this cookie back to our application as part of every subsequent
request. That's how cookies work: once set, they'll be included into every
request that is being made from now on ... and our web application can use
it.

When our browser is now redirected to `GET` the same URL again, it passes the
cookie, and Sinatra will, because we have the `:sessions` feature enabled,
*deserialize* (i.e. decrypt and read) the data, and put it back into the
hash that is returned by the method `session`, so we can work with it.

In our `get` route, if we find something in `session[:message]` we will display
it in the view. If nothing is stored on that key in the session then the view
won't display anything either.

Does that make sense?

Awesome :)

### Transient state

However, there's a little problem with our approach. Have you noticed?

Let's recap what our application does:

- On the `GET /monstas` route we render a view that includes a form.
- This form, when submitted, makes a `POST` request to the same path `/monstas`, and includes the `name` variable.
- On the `POST /monstas` route we find the `name` data in the `params` hash.
- We store it to the file.
- We set the confirmation message to `session[:message]`.
- We redirect the browser to `GET /monstas`.
- The browser requests `GET /monstas`.
- We find the confirmation message on `session[:message]`.
- We display the message.

This works great.

However, the message is now stored in our session cookie. And that means that
from now on, whenever you browse (make a `GET` request) to the path `/monstas`
the browser will always include the same cookie (data) to the request. And our
application will always find it, and always display the same confirmation
message ... even though we haven't actually added any new names this time.

Instead, what we really want to do is display the confirmation message only
once: on the `GET` request that is made right after the `POST` request
redirected to `/monstas`. When we then reload the page (or close and reopen
the browser and go to `/monstas` tomorrow) the confirmation message should
be gone.

Right?

This is called "transient state": State that is only there for a brief moment,
and then goes away. And a session is a great place to keep it.

So how can we fix that?

All we have to do is delete the message from the session right before we
display it:

```ruby
get "/monstas" do
  @message = session.delete(:message)
  @name = params["name"]
  @names = read_names
  erb :monstas
end
```

Deleting it from the `session` will return the value that was stored on this
key, and we assign it to the instance variable `@message`, which makes it
available to our template.

In other words, if anything is stored on this key it will be assigned to the
`@message` instance variable, and the view will display it. If nothing's stored
on the key, then deleting the key will simply return nil, and nothing will be
displayed in the view.

Problem solved :)

# Validations

Our little application now already makes use of a bunch of things that you will
regularly find in web applications.

We have a form that posts data to another route. The `post` route picks up the
data, and stores it, then redirects to another route, which displays the data.
We also use a session, and a query parameter to pass data from one route to
another.

Pretty cool. You'll see a lot of these very same concepts in use when you start
building your first Rails application. All of these things will work pretty
much the same way in Rails. Except, you've now built them manually yourself,
so you know how this stuff works under the hood.

Let's look at one other concept that Rails helps with, too: Validating user
input.

If you look at our little application we still are a little naive in accepting
whatever data comes in to our `post` route: We simply store whatever the user
sends, whenever they send it.

Do we really want to store duplicate names? What if the same name is being
submitted over and over again? And what if there's no name submitted at all?

What we really want to do is validate the incoming data (in our case the name),
and only accept and store it when we find it's valid. If it's not, then we want
to display a message to the user, and ask them to submit the form again.

We could change our `post` route like so:

```ruby
post "/monstas" do
  @name = params["name"]

  if @name.nil? or @name.empty?
    session[:message] = "You need to enter a name."
  elsif read_names.include?(@name)
    session[:message] = "#{@name} is already included in our list."
  else
    store_name("names.txt", @name)
    session[:message] = "Successfully stored the name #{@name}."
  end

  redirect "/monstas?name=#{@name}"
end
```

This is a valid implementation, and if you restart your application you
can try it out.

The `if` statement first checks if the `@name` is empty. If it is we simply
store a message to the session, and then redirect.

Note the duplicate condition `@name.nil? or @name.empty?`. The name parameter
could either be missing (and thus, be `nil`), or it could be an empty string,
so we need to check both cases.

This could be simplified to one condition like so:

```ruby
  if @name.to_s.empty?
```

If it's `nil`, then `nil.to_s` would return an empty string. If it's an empty
string, then `"".to_s` returns the same empty string again.

The `if` statement then also checks if the given `@name` already is included in
the names in our file: the array returned by `read_names`. If we already have
the name, then, again, we just add a message to the session, and redirect.

Only in the last case, when the name is not empty, and not already known,
we do store it, and add the respective message to the session, and redirect.

Alright, this works.

However, it's worth considering that this adds quite a bit of stuff to our
route. And if we have a lot of routes then that's a lot of code to add.

So what if we extract this to a separate class?

Extracting code to a class (or method) is a useful technique to keep your
code clean and readable. The routes can focus on their job (passing data
from the request to the view, reading data from our file, storing new
data to the file etc), and the new class can focus on a different task:
Finding out if the passed data is valid.

So let's try that, and implement a little Ruby class that hides some of the
logic from the route. Here's how we could do it:

```ruby
class NameValidator
  def initialize(name, names)
    @name = name.to_s
    @names = names
  end

  def valid?
    validate
    @message.nil?
  end

  def message
    @message
  end

  private

    def validate
      if @name.empty?
        @message = "You need to enter a name."
      elsif @names.include?(@name)
        @message = "#{@name} is already included in our list."
      end
    end
end

post "/monstas" do
  @name = params["name"]
  validator = NameValidator.new(@name, read_names)

  if validator.valid?
    store_name("names.txt", @name)
    session[:message] = "Successfully stored the name #{@name}."
  else
    session[:message] = validator.message
  end

  redirect "/monstas?name=#{@name}"
end
```

This adds quite a bit of code that we need to figure out, and type. But it
seems worth it: Our route is now much shorter, and way easier to understand
from just a quick look at it. We could move the `NameValidator` to a separate
file.

Cool, this works great.

## Rerendering the form

However, imagine we'd now have a much bigger form, with lots of fields. And
we'd validate each of these fields and have several validation messages.
Imagine we'd have like 20 form fields, and the user has made mistakes on 5 of
them.

We'd want our application to display the same form again, and display the
validation messages alongside with it. The original data should be, again,
prefilled to the form, so the user does not have to type it all again.

In order to preserve the form data that had been entered by the user we'd
need to append it all to the URL (as we do with the name in `"/monstas?name=#{@name}"`).
And we'd need to store all the validation messages to the session.

This is a lot of stuff. And it actually also might break because URLs cannot be
very long.

For this reason modern web applications usually follow a different pattern:

If the submitted data is invalid, instead of redirecting the user, we would
simply re-render the same template right there, with the same data.

Here's how we could do that:

```ruby
post "/monstas" do
  @name = params["name"]
  @names = read_names
  validator = NameValidator.new(@name, @names)

  if validator.valid?
    store_name("names.txt", @name)
    session[:message] = "Successfully stored the name #{@name}."
    redirect "/monstas?name=#{@name}"
  else
    @message = validator.message
    erb :monstas
  end
end
```

As you can see we now only store a message to the session, and redirect, only
if the given data is valid. If it's not then we store the validation message to
the `@message` instance variable and render our template again.

Cool. If you restart your application you can try how it works.

However, this now displays an empty "Hello" at the top. Why's that?

We assign `params["name"]` to the instance variable `@name` which we then check
in our template: `<% if @name %>`. However, since this has been submitted by a
form, with an input element called `name`, what we get is an empty string, not
`nil`.

We can fix this by changing our view like so:

```erb
<% unless @name.to_s.empty? %>
  <h1>Hello <%= @name %></h1>
<% end %>
```

Awesome.

With this completed you have now walked through an important pattern for web
applications, which is also used in Rails applications by default:

- There is an HTML form which is being retrieved via a `GET` request.
- This form posts to another route, which validates the submitted data.
- If the data is valid, it does something with it (in our case we store it) and
  redirects, passing a message via the session.
- If it's not valid, it renders an error message, as well as the form,
  with the form fields populated with the given data.

Make sure to remember this pattern. Maybe write it down to a cheatsheet,
formulate it in your own words. Maybe try turning it into a comic.

Next, let's take this a step further and have a look at something that Rails
calls "resources".

# Resources

In HTTP a "resource" is something that is identified by a URL.

In Rails a "resource" is something slightly different: a collection of 7 routes
that belong together in the sense that they allow listing, viewing, and
managing a collection of things. It is, essentially, a convention of grouping
routes, which has been encoded in Rails, but also is being used in other
applications, and generally a great practice.

Let's have a look at an example.

## Groups of routes

Imagine we are building an application that allows managing the members of our
study group. There's a way to

- list all the members
- look at a member's details
- add a new member
- update a member's details
- remove a member

That's exactly what a *lot* of web applications do, with certain variations.

Almost all web applications that you use on a daily basis will, in some way or
the other, have lists of things. Amazon has lists of goods, and so has Ebay.
Facebook has lists of posts, and so has Twitter. Gmail obviously has lists of
emails, and so on, and so on. Usually there's a way to display details for
these things, to create new instances (e.g. by publishing a post on Facebook),
editing, and deleting them.

Because this is such a common scheme, Rails has added a first class concept for
this. And, as we do, we are going to re-implement it in Sinatra. This way you'll
have a great understanding of what a "resource" is in Rails once you get started
with that framework.

Also, for our purposes we'll deviate from Rails just a little bit. And we'll
explain that later.

In a nutshell, Rails defines a `resource` as a collection of routes that deal
with the same "thing". In our example we deal with "members", so our resource
will be `members`.

Here's what the routes that make up our resouce look like:

- `GET /members` displays a list of all members
- `GET /members/:id` displays the details for a single member
- `GET /members/new` displays a form for creating a new member
- `POST /members` creates a new member from that form
- `GET /members/:id/edit` displays a form for editing a member's details
- `PUT /members/:id` updates a member's details from that form
- `GET /members/:id/delete` asks for confirmation to delete the member
- `DELETE /members/:id` deletes the member

Here's the same information as a table. Note that we also give names to these
routes [1]:

| Name    | Method | Path                | Function                            |
| ------- | ------ | ------------------- | ----------------------------------- |
| index   | GET    | /members            | Display all members                 |
| show    | GET    | /members/:id        | Display a single member             |
| new     | GET    | /members/new        | Display a form for a new member     |
| create  | POST   | /members            | Create that new member              |
| edit    | GET    | /members/:id/edit   | Display a form for editing a member |
| update  | PUT    | /members/:id        | Update that member                  |
| delete  | GET    | /members/:id/delete | Ask for a confirmation to delete    |
| destroy | DELETE | /members/:id        | Delete a member                     |

These names are the same as the ones that Rails uses, too. We'll use these
names for our templates, in case we need a template. Naming them the same for
every resource that we write, ever, helps others to understand what we're
talking about in an instant.

So lets look at the routes more. If you look at the purpose of our routes,
there are 4 groups:

- `index` and `show` are used to display existing data.
- `new` and `create` are used to create a new member.
- `edit` and `update` are used to update a member.
- `delete` and `destroy` are used to delete a member

The two pairs `new` and `create` as well as `edit` and `update` would follow
the same pattern that we've discussed in the chapter about
<a href="/validations">Validations</a>:

- The first one request `GET`s an HTML form for the user to enter some data.
- This form is then submitted as another request, using `POST` or `PUT`, to the
  second route.
- The second route validates the data.
- If the data is valid it creates/updates the member, and redirects to the show
  view, passing a confirmation message.
- If the data is not valid it re-renders the form with an error message.

For the last pair there's no validation, of course. Instead we just delete the
object and redirect to the index view.

However there's a problem with all this:

Today's browsers still do not allow sending HTTP requests using any other verb
than `GET` and `POST`.

Ouch.

So what do we do?

We fake that. Let's see ...

### Footnotes

[1] For our purposes here we have deviated from Rails a little, and added an
8th route: `GET /members/:id/delete`. This route displays a page that asks for
a confirmation to delete. In Rails this is solved with a little Javascript box
that pops up and asks this question. However, we don't want to get into
Javascript too much just yet, and adding this route is a just as valid
solution, too. In fact, some web applications out there prefer this over
the Javascript solution.

## Faking HTTP verbs

Remember how we said that in an HTML form we can specify the HTTP verb that
is supposed to be used for making the request like so:

```html
<form action="/monstas" method="post">
  ...
</form>
```

This makes the form `POST` to `/monstas`, instead of the default `GET`.

Now, it's probably fair to say that every sane person in the world would expect
that it is also possible to make that a `PUT`, or `DELETE` request. Like so:

```html
<form action="/monstas" method="put">
  ...
</form>
```

Except that ... it's not. Today's browsers still do not allow sending HTTP
requests using any other verb than `GET` and `POST`.

The reasons for why that still is the case in 2015 are either fascinating or
sad, depending how you look at it <a href=#footnote-1">[1]</a>
But for now we'll just need to accept that, and work around it.

Sinatra (as well as Rails, and other frameworks) therefore support "faking"
requests to look as if they were `PUT` or `DELETE` requests on the application
side, even though in reality they're all `POST` requests.

This works by adding a hidden form input tag to the form, like so:

```html
<input name="_method" type="hidden" value="put" />
```

A hidden input tag is just that, it is hidden, meaning that it is not displayed
to the user. However, it is there, and it will be sent to the application as
part of the request just like any other input tag.

In order to make this work we also need to tell Sinatra that we want this
kind of behaviour. Developers say we need to "opt in" to it. Like so:

```ruby
use Rack::MethodOverride
```

Now, whenever Sinatra receives a `POST` request that has a parameter with the
name `_method` it will treat this request as if it was a request with the HTTP
method (verb) given by this parameter. This way one can add data to the form
that isn't relevant to the user, but relevant to the application.

Sinatra will treat any `POST` request that has the parameter `_method` set to
`put` as a `PUT` HTTP request: it will use a route that was defined with `put`.
Likewise, it will treat a `POST` request that has the paramter set to `delete`
as a `DELETE` HTTP request, and use the respective route.

This way we can write our application code *as if* browsers support
sending forms as `PUT` or `DELETE` requests, even though they don't. The only
thing we need to do is add that little hidden input form field.

The parameter name `_method`, with an underscore in front, has been chosen to
indicate that it's a "private" concern: it is something that Sinatra manages
for us. Also, by choosing an unusual name like this it won't clash with our own
form input fields, in case we ever need to add a form field named `method`,
e.g. for, maybe, a payment method, or a shipping method, or whatever else.

### Footnotes

<a name=footnote-1">[1]</a> You can read more about this, for example,
<a href="http://programmers.stackexchange.com/questions/114156/why-are-there-are-no-put-and-delete-methods-on-html-forms">here</a>.

## Writing resources

Writing web applications can be defined as implementing such resources.

Rails helps a lot with the boilerplate code, code that has to be written over
and over again, when you define a lot of resources for your web application.

However, in this course, we'll want to implement our `members` resource
ourselves, and do all the legwork that is required.

So how about heading over to the exercise [Sinatra Resource](http://localhost:4567/exercises/sinatra_resource.html).

Give it a shot!

# Databases

Databases are applications that help us store data in a flexible and powerful
way, and they have been around ever since the 1960s.

Almost all web applications use database in one way or the other in order to
store all the data they need to deal with. The allow us to store and find data
in flexible, and very efficient ways.

Nowadays there are lots of different types of databases, but we'll focus on the
most prevalent, traditional type of databases. These are called "relational
databases" (or [RDBMS](http://en.wikipedia.org/wiki/Relational_database_management_system)).

## Tables

At its core, you can imagine a database as a bunch of spreadsheets. Except
they're called "tables", and they can be huge, containing tons and tons of
rows. More than any spreadsheet ever could handle.

Each table has a bunch of columns, and can have an arbitrary number of rows
(also referred to as "records"). Columns (also referred to as "fields") have a
name and a type. Their type specifies which kind of data can be stored. Each
row has a number of cells , and each one of the cells can hold some value with
the type defined by the column.

Let's have a look at an example.

Imagine we have a table called `members`, and it contains our member data. So
it could look like this:

| id   | name    | joined_on  |
| ---- | ------- | ---------- |
| 1    | Anja    | 2013-06-24 |
| 2    | Carla   | 2013-06-24 |
| 3    | Rebecca | 2013-06-31 |

The `id` column would be a running number, so it would have the type `integer`.
The `name` column is a `string` (databases call it this type a `varchar`), and
the `joined_on` column is a `date`.

For a table like this, the columns (with their name and type) are referred to
as the database *structure*, whereas the rows represent the *data* that we
store.  Rows change all the time: new members sign up, existing members change
their details, or remove their profile etc. The structure remains the same,
unless we, as developers, need another column, or table in order to store more
data.

Databases usually contain many tables. And often data from one table relates
to data in other tables.

For example, we could add the ability to post status updates to our members
application. Maybe we would have a table `statuses`:

| id   | member_id | message                                  |
| ---- | --------- | ---------------------------------------- |
| 1    | 1         | Finished the search feature for speakerinnen.org |
| 2    | 1         | Working on the CSS cleanup with Maren next |
| 3    | 3         | Created some new designs for our stickers! |

The `id` column, again, contains a running number that allows to identify a
single status update. The column `message` obviously contains the status update
message.

What about the `member_id` column though?

It references a row in a different table: our `members` table. This means that,
in this example, Anja has posted two status updates, Rebecca one, and Carla
hasn't posted yet.

If you look at the two columns `id` and `member_id` you notice that the column
`id` is special: It must never contain duplicate values, because we want to use
the `id` to identify a certain message (or member). This is called a *primary
key*, and the column is called a "unique" one. Also, it usually auto-increments
the id for us: Whenever we store a new row to this table then the database will
assign an id, make sure we get the next number, and never get duplicate values.

These "features", or special properties of the `id` column also are considered
part of the structure, alongside with the column name and type: we define
these things when we create or modify the database structure.

The column `member_id` on the other hand should not be unique: We want to be
able to store many messages that all belong to the same member row, in the
`members` table. Therefor we need to be able to have multiple rows with the
same `member_id` in the `messages` table.

Does that make sense? This is how we can store data in a database, give it a
certain structure, and relate a piece of data (a row) in one table to data
(rows) in other tables.

Now, how can we talk to a database like this? How can we actually add data to a
table, or retrieve it?

## SQL

*Talk to databases*

Using SQL (Structured Query Language) we can talk to (relational) databases.

For example we can ask (query) the database to retrieve certain bits of
information from tables, or we can insert, update or delete data. And SQL is
also the language that is used to define the database structure in the first
place.

SQL was invented in the 1970s, and it's quite ugly to look at. However, lots
of database systems support it, and so it's quite common to use SQL in web
applications in some way.

For example, we could retrieve all fields in the first row from our `members`
table like this:

```sql
SELECT * FROM members WHERE id = 1;
```

The statement `SELECT` tells the database that we'd like to *retrieve* data (as
opposed to, for example, `INSERT` which inserts a new row, or `UPDATE` which
updates an existing row). The star `*` means "all fields". `FROM` specifies the
table that we want to look at, and `WHERE` specifies a condition that this row
needs to match: We'd like to retrieve the row where the value in the `id`
column equals `1`.

So our query above would return a result containing one row:

```
1 | Anja | 2013-06-24
```

However, when we ask for all rows that have the `joined_on` date `2013-06-24`
we'd get back two rows:

```sql
SELECT * FROM members WHERE joined_on = '2013-06-24';
```

This would return:

```
1 | Anja  | 2013-06-24
2 | Carla | 2013-06-24
```

Instead of asking for all fields per row, we could also just ask for a certain
column that we are interested in:

```sql
SELECT name FROM members WHERE joined_on = '2013-06-24';
```

This would return just the names:

```
Anja
Carla
```

In order to insert a new row to the table we could use an SQL statement like
this (assumning our `id` column auto-increments, i.e. automatically assigns
the next number to the new row):

```sql
INSERT INTO members (name, joined_at) VALUES('Maren', '2013-06-24');
```

Updating looks like this:

```sql
UPDATE members SET joined_on = '2013-06-24' WHERE id = 3;
```

And deleting like this:

```sql
DELETE FROM members WHERE id = 3;
```

As you can see these statements all look somewhat similar, starting with a certain
command, naming the table, and ending with a semicolon. However, they also don't
really look very consistent. For example, why does the `INSERT` statement separate
the column names and inserted values, while the `UPDATE` statement pairs them?

On the other hand, even though it's a little weird, it's also a very powerful
language, and being able to figure out some SQL and manually writing it can be
very useful when you have access to a database, and you want to find out some
bits of information that cannot be retrieved with the application that is using
the database: You'd just directly talk to the database, and ask it for the
information you need.

Of course there are tools for this. We'll look at libraries that make it easy
to talk to databases in the chapter about [ORMs](/databases/orm.html).

Let's play with a real database, and run some SQL statements next.

## SQLite

[SQLite](http://en.wikipedia.org/wiki/SQLite) is a minimalistic implementation
of a relational database that supports most of SQL, although not all of it. It
is less powerful than, for example, [PostgreSQL](http://en.wikipedia.org/wiki/PostgreSQL) and
[MySQL](http://en.wikipedia.org/wiki/MySQL), but it's also super lightweight,
and great for learning, experiments, and getting started quickly.

We are going to use the database SQLite in our examples, because it's the least
hassle to set it up. So you want to make sure you have it installed, many
operating systems have it preinstalled.

Check if it's installed by running this in your terminal:

```
$ sqlite3 --version
```

If that outputs a version string then SQLite is installed. If it's not then
you'll see something like `command not found: sqlite3`. In that case Mislav has
written up some nice instructions [over here](http://mislav.uniqpath.com/rails/install-sqlite3/).

SQLite comes with a handy command line tool that one can use to create
databases and interact with them. It also has an interactive
[shell](http://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop):
Just like with [IRB](http://ruby-for-beginners.rubymonstas.org/your_tools/irb.html)
(where we can execute Ruby code interactively) we can log into the SQLite
database, and execute [SQL](/databases/sql.html) statements interactively.

Let's try it:

```
$ sqlite3 members
```

This should put you into the interactive SQLite shell, and look something
like this:

```
SQLite version 3.8.5 2014-08-15 22:37:57
Enter ".help" for usage hints.
sqlite>
```

The prompt `sqlite>` waits for your input. If you type an SQL command and hit
return it will execute it.

Let's create our `members` table:

```
sqlite> CREATE TABLE members (id INTEGER PRIMARY KEY, name VARCHAR(255), joined_on DATE);
```

If this does not output an error, then the command was successful.

The command says that we'd like to create a table `member` with 3 columns `id`,
`name`, and `joined_on`. The column `id` is supposed to be an integer column,
and we'd like to use it as our primary key (meaning that it will be unique, and
it will autoincrement the id for us). The column `name` is a string column, and
values can be 255 characters long. And `joined_at` is a date column.

Cool. So we've just created a table in our database.

We can list our tables like so:

```
sqlite> .tables
members
```

And we can check the structure (schema) of our `members` table like so:

```
sqlite> .schema members
CREATE TABLE members (id INTEGER PRIMARY KEY, name VARCHAR(255), joined_on DATE);
```

Nice.

Now let's insert a row to the table:

```
sqlite> INSERT INTO members(name, joined_on) VALUES('Anja', '2013-06-24');
```

Again, if this does not output anything, that means our command was successful.

We can now retrieve the data using a `SELECT` statement like so:

```
sqlite> SELECT * FROM members;
1|Anja|2013-06-24
```

So this displays on row.

Now should be a good time to do some [exercises on SQL](/exercises/sql.html).

## Object Relational Mappers

*Map database contents to Ruby objects*

There are lots of reasons, most of them historical, why SQL reads weird, and
it's quite unlikely that SQL will become a more pleasant language, or replaced
anytime soon.

For this reason programmers have written lots and lots of tools (libraries),
which make talking to, and working with relations databases a little bit
easier.

One class of such tools is called ORM: [Object Relational Mappers](http://en.wikipedia.org/wiki/Object-relational_mapping).

An ORM is a library that "maps" data, stored in a database, to objects, and
usually has methods such as `save`, in order to save an object as a database
row, `create` to insert new data, `update` to change it, and so on. In other
words, they usually provide a [DSL](/sinatra/dsl.html) for working with the
database.

Your data is stored in the database as rows, because that's what databases
do. However in your Ruby application (or whatever language you use for this)
you would see, and use this data as objects: because that's what Ruby is
great at. The ORM is a tool that transforms your database data to Ruby
objects and vice versa.

Does that make sense?

Let's look at an example.

If we have a database table `members` like this:

| id   | name    | joined_on  |
| ---- | ------- | ---------- |
| 1    | Anja    | 2013-06-23 |
| 2    | Carla   | 2013-06-24 |
| 3    | Rebecca | 2013-06-31 |

Then in our Ruby code, using an ORM, we could communicate with it like so:

```ruby
class Member # we'd need to somehow include the ORM functionality
end

# Find one member

member = Member.find(id: 1)
puts "#{member.name} has joined on #{member.joined_on}."

# Change the member's joined_on date:

member.joined_on = '2013-06-24'
member.save

puts
puts "Correction!"
member = Member.find(id: 1)
puts "#{member.name} has joined on #{member.joined_on}."

# Find several members based on their joined_on date:

puts
puts "Who joined on 2013-06-24?"
members = Member.where(joined_on: '2013-06-24')
members.each do |member|
  puts "#{member.name} has joined on #{member.joined_on}."
end
```

And this would output:

```
Anja has joined on 2013-06-23.

Correction!
Anja has joined on 2013-06-24.

Who joined on 2013-06-24?
Anja has joined on 2013-06-24.
Carla has joined on 2013-06-24.
```

Of course the details of this Ruby code might vary, depending on the concrete
ORM tool that we are using.

But the basic idea is that we can use Ruby classes and objects to retrieve some
data from the database (as in `Member.find(id: 1)`), which would then appear in
our application as a normal Ruby object. We can call methods to look up fields
(such as in `member.name`, which returns the value from the `name` column).
And we can use the same object to modify, and save the data back to the
database.

Two widely used ORMs in Ruby are [ActiveRecord](https://github.com/rails/rails/tree/master/activerecord),
which is part of Rails, and [Sequel](https://github.com/jeremyevans/sequel),
which is more modern, slick, and performant.

Since you will anyway get to know ActiveRecord later when you learn Rails, we
will introduce Sequel first. This way you will get to know two libraries and
can later compare them.

You can install the Sequel gem like so:

```
$ gem install sequel
```

# Migrations

*Persist executable changes to the database structure*

In programming, when we store something to a medium that survives when the
program terminates (or when the computer is shut down, and rebooted), we use
the term [persistence](http://en.wikipedia.org/wiki/Persistence_(computer_science))
for that.

Databases are probably the most common way to persist data in webapplications.
Your web application would read data from the database, and store data to it.
E.g. our application could read members, and store them. We use the term
*runtime* for this: Our application does all of this "at runtime", meaning:
after it has been started, and while it is running.

This data is stored as rows in our database tables. What about the structure
though?

The database structure normally does not change at runtime. The users of
your application do not get to add more columns to the `members` table,
right? That's something you, as a developer, want to define, basically
while you write your code.

So even though the database structure itself is not exactly part of your code,
it highly relates to your code: Your Ruby code makes assumptions about what
tables exist, and what columns each of the tables has. In other words, your
code would not work unless the database has a certain structure.

E.g. this Ruby code from our previous example makes the assumption that there
is a table `members` and it has at least the columns `id`, `name`, and
`joined_on`:

```ruby
member = Member.find(id: 1)
puts "#{member.name} has joined on #{member.joined_on}."
```

Because defining the database structure is a task that is separate from running
the actual application (starting your web server), there usually is some tool
to create a database, and apply the exact structure that your application (Ruby
code) requires.

One could, of course, create the database, and its structure manually, just
like you've just defined a table in the interactive SQLite shell.  However that
obviously is tedious, makes it difficult to collaborate with others, and
install your application in different environments: People would have to create
the database manually over and over, so that's not really a practical option.

Instead web applications use a concept called "database migrations", and we'll
explore it here because it's going to be a big topid in Rails, too.

## Database structure

A better way to share, and programmatically load a database schema is to store
the respective SQL commands to a file, and then make the database execute it.

For example, for our `members` database we could have a file `schema.sql` that
contains the following SQL command (we only have one table, so a single command
is sufficient):

```
CREATE TABLE members (id INTEGER PRIMARY KEY, name VARCHAR(255), joined_on DATE);
```

We could then import this structure definition into a new database named
`members-2` using a command like this:

```
cat schema.sql | sqlite3 members-2
```

This already works much better than asking our fellow developers setup their
database manually.

We could share this file with them, check it into our version control system
(e.g. Git), and they could create the same database with the same command,
making sure they'd always end up with the same database that we have, and that
our application (Ruby code) requires.

However, there's another problem with this.

Imagine we're all working on this shared Ruby Monstas members app. Someone
starts by defining a first, simple version of the `members`, much like the
one that we currently have. Then later, someone else picks up work, and adds
another column. Maybe someone else will then also rename one of the columns,
or drop it, because it's no longer required.

Everytime we'd export a new version of our database schema to the file `schema.sql`
and share it with each other.

However, since the database schema only describes the database structure in its
final state, we could not just run it: Our local database already has a certain
table structure defined, and (e.g.) the `CREATE TABLE` command would fail.

The only solution to this would be to delete our local database, and re-create
it using the new schema. This would work fine in development. However, what if
we already have this application running publicly, and people are already using
it? We'd then loose all the data that has been stored previously.

Obviously that wouldn't be an acceptable solution.

The solution to this is a concept that is referred to as *database migrations*.

Let's have a look.

## Incremental changes

Instead of sharing the final database structure, as defined in our `schema.sql`
file, we would share many files where each of them defines a single change to
our database structure. We also somehow number these files so we can run each
change one after another in the precise same order.

Imagine our first step is to define the table members. So "change 1" would
contain our SQL code from above. We'd store this in a file `db-change-1.sql`:

```
CREATE TABLE members (id INTEGER PRIMARY KEY, name VARCHAR(255), joined_on DATE);
```

Now, the next person adds another change to the database structure, which adds
the table `messages` from our example above. They store this in a file
`db-change-2.sql`:

```
CREATE TABLE messages (id INTEGER PRIMARY KEY, member_id INTEGER, message TEXT);
```

And then, a couple days later, another person figures we should also keep track
of a timestamp which tells us when a message has been sent. We store this change
in a file `db-change-3.sql`

```
ALTER TABLE messages ADD COLUMN sent_at DATETIME;
```

Now we have 3 changes, stored in 3 separate files. We also know the order in
which these changes have been created: we can defer this from their filenames.
If we keep track of changes that we already have applied to our database
structure, then we can easily run the ones our collaborators have added just
recently, and continue working on our code.

This concept is called *database migrations*: We migrate the database structure
from one state to the next one, by applying one change after another.

Sounds useful?

## Incremental rollbacks

So far you've learned how to, on a conceptual level, keep track of incremental
changes to your database structure. You learned how to store such changes in
files that can be shared, and, when executed will change the database structure
one step after another.

That's pretty useful.

However, there's more to migrations.

What if we make a change to our database structure, apply it to our
"production" application (i.e. the one that is publicly running on the
internet, and actively used by our users), and then notice we've made a
mistake? We'd then need a way to quickly roll back our change. How would we do
that?

This is why we would, per change (or per "migration"), not only define a way
to *apply* a certain change. We would also define a way to roll it back, i.e.
undo it.

With our naive, homegrown example migrations system, where we store plain
SQL in files, we could do this by using file names like `db-change-1-up.sql`
and `db-change-2-down.sql`.

In the terminology of migrations the terms `up` and `down` are used to describe
changes that are *applied* and *undone*: Migrating *up* means applying a
change, and migrating *down* means undoing a change (rolling it back).

This is because traditionally migrations (changes) have been numbered sequentially,
just as we do this in our example. That means applying a new change migrates
the database "up" to a higher verion (change) number. E.g. if we'd apply our
change `db-change-up-3.sql` we'd "migrate up to version 3". If we then undo it
we "migrate down from version 3".

Does this make sense?

## Migration tools

Of course, managing all of this manually is a little tedius so far:

- We need to deal with plain SQL.
- We need to execute individual SQL (migration) files with a manual command on the command line.
- We do not have a good way of keeping track of the change (version) number that has been applied to the database structure last.

For that reason ORM tools usually come with a way to not only automatically
track changes that still need to be applied to the database structure (and
skip the ones that already were applied in the past).

They also allow us to define the database structure in a Ruby
[DSL](/sinatra/dsl.html), so we do not have to deal with the gory details of
SQL, but can use some nice, readable Ruby code instead.

Here's an example of a migration written for Sequel:

```ruby
Sequel.migration do
  up do
    create_table(:members) do
      primary_key :id
      string :name
      datetime :joined_on
    end
  end

  down do
    drop_table(:members)
  end
end
```

And here's an example of a migration in ActiveRecord:

```ruby
class CreateMembers < ActiveRecord::Migration
  def up
    create_table(:members) do |table|
      table.string :name
      table.datetime :joined_on
    end
  end
end
```

As you can see the Ruby code looks quite a bit different. That's because Sequel
and ActiveRecord define different DSLs for applying changes to your database
structure. You'll just need to learn whatever tool you're using, and look things
up from the documentation.

However, they both have in common that there's a bit of code for applying the
change (and they both use the method name `up` for describing it), and a bit of
code for reverting the change (undoing it, using the method name `down`).
And they have in common that you can use Ruby code, instead of plain SQL.

One other advantage of this is that the SQL that we need to run actually sometimes
looks a tiny bit different depending what database system we use (e.g. SQLite,
Postgres, MySQL). ORMs abstract these changes away from us, meaning that they
take care of generating different SQL depending what database system we use,
and let us describe changes in the same Ruby code no matter what.

If you find some of this confusing then don't worry. All of this will start
making sense as soon as you use migrations in praxis, in order to create
tables, and then apply more changes to them later.

# Exercises

## Using Rubygems

A data format is a way to represent rich data types (such as Strings, Numbers,
Arrays, Hashes, and so on) in form of a simple text document that can be stored
on the local file system, or transferred over the internet.

<a href="http://en.wikipedia.org/wiki/JSON">JSON</a> is a data format that is
very common in modern web application development, and it is useful for
computers to talk to each other, and pass around data. JSON is human readable,
and pretty similar to, for example, Javascript and Ruby code.

For example, this is a valid piece of JSON:

```json
{
  "emails": [
    {
      "subject": "Hi there, Ruby Monstas!",
      "date": "2015-01-02",
      "from": "Ferdous"
    },
    {
      "subject": "Keep on coding!",
      "date": "2015-01-03",
      "from": "Dajana"
    }
  ]
}
```

As you can see the syntax is similar to Ruby hashes and arrays: There is an
outer hash that has a single key `"emails"`, which is a string. This key's
value is an array that has two elements, each of which is another hash with the
keys `"subject"`, `"date"`, and `"from"`.

So, this piece of JSON represents data for a collection of two emails.

In order to read (we say "parse") this piece of JSON data in a Ruby application
we can use the `json` gem. (You can find it on
<a href="https://rubygems.org/gems/json">Rubygems.org</a>, too).

You can run `gem list json` in order to check if this gem is already installed
on your system. `gem install json` will install it, and `gem update json` will
look for a newer version and install that one if there's one.

Make a new file `libraries-1.rb`, and add the following lines:

```ruby
require "json"

data = '{
  "emails": [
    {
      "subject": "Hi there, Ruby Monstas!",
      "date": "2015-01-02",
      "from": "Ferdous"
    },
    {
      "subject": "Keep on coding!",
      "date": "2015-01-03",
      "from": "Dajana"
    }
  ]
}'
data = JSON.parse(data)

p data.keys
```

If you then run `ruby libaries-1.rb` you'll see that the method `JSON.parse` converts
the string that is stored in the variable `data` to a Ruby hash, so we can call `keys`
on it, and it will output an array with the single key defined on the outer hash:

```ruby
["emails"]
```

Likewise, if you add another line at the end, as follows:

```ruby
p data["emails"].first["subject"]
```

You see that `data["emails"]` returns the array stored on that key, so we can call `first`
on it, which returns the first hash in that array, and we can finally fetch the value
for the key `"subject"` stored on that hash.

So this prints out `"Hi there, Ruby Monstas!"`.

By the way, if you've read the bonus chapter on
<a href="http://ruby-for-beginners.rubymonstas.org/bonus_1/alternative-syntax.html">Alternative Syntax</a>
then you may have notices that we also have a perfect example of a usecase for
the `%(..)` string syntax here:

JSON data will often contain double quotes (for any string), and often contain
single quotes. Since JSON data, as a whole, needs to be a plain Ruby string
when defined inside of your Ruby code, this is a great usecase for this syntax.

Like this:

```ruby
require "json"

data = %({
  "emails": [
    {
      "subject": "Here's this week's homework!!",
      "date": "2015-01-04",
      "from": "Ariane"
    }
  ]
})
data = JSON.parse(data)

p data["emails"].first["subject"]
```

## Using Bundler

Imagine you'd want, for some reason, use an older version of a particular gem,
rather than the latest one.

Maybe the authors of this gem have made a change that breaks your application.
Of course they'd normally try to avoid this at all costs, but sometimes it
still happens. (Another, maybe more common, case is that you are working on a
Rails application that has been started a year ago, or is even older than that.
Now a new major version of Rails comes out, and ships lots of changes that
would break this application. So you want to stick to the older version of
Rails until you get around making it ready for the newer one.)

One way to use an older version of a particular gem is to simply uninstall
any newer versions from your system. E.g. you could run:

```
gem uninstall rails
gem install rails --version '~> 3'
```

This would uninstall all rails gems, and then install the latest version of
Rails that still starts with a `3` (i.e. ignoring all the newer versions that,
as time of this writing, start with a `4`).

However, this is quite cumbersome. You'd need to figure out what versions of
all the gems that you want to use are compatible with each other. And if you
ever switch from one application to another you'd have to uninstall and
reinstall all those gems over and over.

Instead, Bundler provides a much better way of picking certain gem versions
out of the plethora of gems that might be installed on your computer.

In this example we are going to use the Chronic gem, which provides a great
way to convert natural language date formats to Ruby `Time` objects.

Inside a new directory `chronic` create a file `Gemfile`, and add the following
lines:

```
source "https://rubygems.org"

gem "chronic", "~> 0.9.0"
```

The `source` directive says that we'd like to use the standard RubyGems.org
as a source for our gems. There are other sources, such as private Rubygems
servers run by companies to host their own, internal gems, or Rubygems servers
that people run in order to sell paid versions of their gems.

The `gem` directive declares that we'd like to use the gem `chronic`, and
we'd like to use any version that starts with `0.9` (for whatever reason).

Now, make sure you have `cd`ed to that directory, and run:

```
bundle install
```

You'll see that Bundler first fetches some meta data from RubyGems.org. This
means it checks what versions of the Chronic gem are available, what dependencies
it has, and what dependencies these might have.

In our case Chronic doesn't have any other dependencies, and so Bundler just
installs it. You will see that it installs a version starting with `0.9` (as
time of this writing that would be `0.9.1`).

In order to check what gems are part of your bundle you can run `bundle show`.
You can also open the file `Gemfile.lock` that Bundler has created.

With this we've created a little gems sandbox that makes, no matter what other
gems are available on your system, just two gems available to our Ruby code:
Chronic, and Bundler itself.

Lets have a look at the load path to confirm that:

```
bundle exec ruby -e 'puts $LOAD_PATH'
```

You'll se a bunch of directories that belong to your Ruby installation, and make
its standard library available. But you'll also see two directories that end with
something similar to this:

```
gems/chronic-0.9.1/lib
gems/bundler-1.9.6/lib
```

This means that any code that we run using our bundle (i.e. standing in this
directory, and using `bundle exec` as a prefix) will be able to require, and
use the Chronic library.

Let's try that. Create a new file `19-bundler.rb` in this directory, and add the
following lines:

```ruby
require "chronic"

time = Chronic.parse('tomorrow')
p time.class
p time
```

This will output something like this:

```
Time
2015-05-15 12:00:00 +0200
```

You can see that the method `parse` takes a natural language string, and converts
it into an instance of the Ruby class `Time` (see here for the
<a href="http://ruby-doc.org/core-2.2.0/Time.html">documentation</a>), so we can
now use this object's methods to do more interesting stuff. E.g. let's look up
which month that is:

```ruby
require "chronic"

time = Chronic.parse('tomorrow')
puts time.month
```

Let's say for some reason we'd now want to upgrade the Chronic gem, and use the
latest version in the newer `0.10` series. Maybe the authors have added some
new functionality that we'd like to use, or they've fixed a certain bug.

In order to do this we need to allow Bundler using the `0.10` series of versions.
So in your Gemfile change `~> 0.9.0` to `~> 0.10.0`.

If you now run:

```
bundle update chronic
```

You'll see that it now installs a newer version of this gem. Also, it has changed
the version number in the `Gemfile.lock` file.

Bundler is a really cool piece of software. It is one of those things that,
once you've used it for a while, wonder how life was possible before it existed
(*"How on earth did people meet before there were cellphones?!?"*). If you ask
any seasoned Ruby developer how their life was before Bundler, they'll probably
respond *"Quite miserable"*, or maybe give you a blank stare and say *"Ummm, I
forgot"*.

## Generating HTML with ERB

### Exercise 3.1

With what you've learned from the <a href="/erb">Embedded Ruby</a>
chapter, your goal now is to generate the same HTML as in the exercise
<a href="http://ruby-for-beginners.rubymonstas.org/exercises/mailbox_html.html">The Mailbox Html Formatter</a> -
except that you now use ERB for generating the HTML, not the
MailboxHtmlFormatter class.

Make a new file `mailbox_erb-1.rb` and add the following code:

```ruby
require "erb"

class Email
  # your class from Ruby for Beginners, exercise 13.1
end

class Mailbox
  # your class from Ruby for Beginners, exercise 13.1
end

emails = [
  Email.new("Homework this week", date: "2014-12-01", from: "Ferdous"),
  Email.new("Keep on coding! :)", date: "2014-12-01", from: "Dajana"),
  Email.new("Re: Homework this week", date: "2014-12-02", from: "Ariane")
]
mailbox = Mailbox.new("Ruby Study Group", emails)

html = # complete this code ...

puts html
```

Now complete it so that it outputs exactly the same HTML code as in Ruby for
Beginners, exercise 13.1:

```html
<html>
  <head>
    <style>
      table {
        border-collapse: collapse;
      }
      td, th {
        border: 1px solid black;
        padding: 1em;
      }
    </style>
  </head>
  <body>
    <h1>Ruby Study Group</h1>
    <table>
      <thead>
        <tr>
          <th>Date</th>
          <th>From</th>
          <th>Subject</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>2014-12-01</td>
          <td>Ferdous</td>
          <td>Homework this week</td>
        </tr>
        <tr>
          <td>2014-12-01</td>
          <td>Dajana</td>
          <td>Keep on coding! :)</td>
        </tr>
        <tr>
          <td>2014-12-02</td>
          <td>Ariane</td>
          <td>Re: Homework this week</td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
```

### Exercise 3.2

Implement a class `MailboxErbRenderer` which takes a mailbox, and a filename to an ERB template:

```ruby
class MailboxErbRenderer
  # fill in this class body
end

mailbox = Mailbox.new("Ruby Study Group", emails)
renderer = MailboxErbRenderer.new(mailbox, "mailbox.erb")
html = renderer.render

puts html
```

## Mailbox via Sinatra

### Exercise 4.1

After you've read the chapters about <a href="/sinatra.html">Sinatra</a> write
a little Sinatra application that reads emails from a CSV file, and renders
an HTML table using ERB.

Feel free to copy your solution to
<a href="http://ruby-for-beginners.rubymonstas.org/exercises/mailbox_csv.html">Ruby for Beginners, exercise 15.1</a>
for this, and wrap it into a new Sinatra application.

However, make it so that whenever the CSV file is changed, the application
picks this up (and displays the change) without having to restart the
application.

## Sinatra Resource

After reading the chapter about <a href="/resources.html">Resources</a> your
objective is to implement a resource `members` in Sinatra.

### Exercise 5.1

Start by writing a Sinatra application that has an `index` and a `show` route:

1. On `GET` to `/members` display a list of member names, which are stored
   in a file `members.txt`. The erb template name is `index.erb`.
2. Each of the member names is a link that points to `/members/:name` (`:name`
   being the given member's name)
3. On `GET` to `/members/:name` display a details page for this member
   (i.e. just show their name), and a link "All Members" that goes back to
     `/members`. The erb template name is `show.erb`.

### Exercise 5.2

Now add the `new` and `create` routes:

1. Add a link "New member" to the `index.erb` view, and point it to
   `/members/new`.
2. On `GET` to `/members/new` display a form that `POST`s to `/members`.
   This form has one input element called `name` and a submit button. Also,
   add a link "Back" that goes to `/members`.
3. On `POST` to `/members` validate that the given name is not empty, and
   not already in our list. If the validation succeeds, redirect the user
   to `/members/:name` and pass a success message by using the session.
   If the validation fails re-render the form and display an error message.
4. Make sure the success message is displayed in the `show.erb` view.

### Exercise 5.3

Next add the `edit` and `update` routes:

1. In the `index.erb` view add a link "Edit" next to each of the listed names,
   and point it to `/members/:name/edit`.
2. Also, add the same link to the `show.erb` view.
3. On `GET` to `/members/:name/edit` display a form that `PUT`s to
   `/members/:name`. This form has the same elements as the form on `new.erb`.
   Also, add a link "Back" that goes to `/members`.
4. On `PUT` to `/members/:name` validate the given name. If the validation
   succeeds redirect the user to `/members/:name` and pass a success message
   by using the session. If the validation fails re-render the form and display
   an error message.

### Exercise 5.4

Finally add the `delete` and `destroy` routes:

1. In the index view add a link "Delete" next to each of the "Edit" links,
   and point it to `/members/:name/delete`.
2. Also, add the same link to the `show.erb` view.
3. On `GET` to `/members/:name/delete` prompt the user for confirmation:
   "Do you really want to remove the member [name]?", and add a form that sends
   a `DELETE` request to `/members/:name`, with a button "Remove Member".
   Also add a link "Back" that goes to `/members/.
4. On `DELETE` to `/members/:name` remove the name from the file `names.txt`,
   and redirect to `/members`

<p class="hint">
For the two forms on the <code>edit.erb</code> and <code>delete.erb</code>
views you'll need to apply the trick from the
<a href="/resources/fake_methods.html">Faking HTTP verbs</a> chapter.
</p>

<p class="hint">
If you have a hard time figuring out why a certain request does not work as
expected try reading the logs of your Sinatra application (in your terminal).
If that still doesn't give you a good hint try inspecting the request in your
browser's web inspector on the network tab.
</p>

## SQL

Interactive SQL training:
http://www.sqlcourse.com/ƒ

