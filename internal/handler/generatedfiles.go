package handler

// Initial content for the static website.

// HomePage
var homeIndex string = `---
title: "index"
date: 2000-01-01
author: "John Doe"
description: "This is the index page"
image: ""
draft: false
---

# Welcome to My Website

Welcome to the homepage of my static website. Here, you can find information about my projects, blog posts, and more.

## Features
- **Simple**: A clean and minimalistic layout.
- **Static**: No servers, just fast-loading static pages.
- **Flexible**: Easily customizable content.

## About Me
I'm passionate about web development and love creating beautiful, fast, and accessible websites.

## Explore
- [Projects](projects): View my work and experiments.
- [Blog](blog): Read my thoughts and tutorials.

`

// BlogIndex
var blogIndex string = `---
title: "index"
date: 2000-01-01
author: "John Doe"
description: "This is the index page"
image: ""
draft: false
---
# My Articles

Welcome to my blog! Here, I share my thoughts, tutorials, and experiences in web development.
`

// Post
var basePost string = `---
title: "Lorem ipsum"
date: 2000-01-01
author: "John Doe"
description: " Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras lacinia ut velit a ultricies. Etiam at rutrum mauris. Fusce nibh ex, commodo ut dui eu, blandit mattis massa. Curabitur fringilla est a turpis porta tempor. Nunc aliquam faucibus magna, vel."
image: "/satic/post.png"
draft: false
---
# Lorem ipsum

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean lacinia pretium convallis. Maecenas vel porta mi. Nullam auctor a nibh ut venenatis. Sed ut vehicula tellus. Donec turpis elit, gravida in faucibus eget, laoreet non lorem. Nullam mattis nulla eget eleifend interdum. In auctor felis sapien, vitae bibendum diam mollis sit amet. Fusce maximus pulvinar nisl at consectetur. Fusce eu tortor ac ante congue porta et et dui. Curabitur eget hendrerit dui. Donec congue semper ligula, vel ornare quam imperdiet vel.

Nulla viverra lacus a augue ullamcorper, ut porttitor lectus vestibulum. Morbi consectetur tortor sit amet arcu tempus, nec sollicitudin nisi interdum. Aenean in sem ac diam vulputate tincidunt. Aenean in venenatis lacus, ultrices tempus massa. Sed ex mi, convallis ut placerat in, posuere eget lorem. Aliquam ac justo felis. Etiam imperdiet, leo id aliquet convallis, nunc tellus condimentum sem, vitae pellentesque nulla orci ac lacus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Integer a ex a nulla venenatis tincidunt non suscipit tortor. Donec venenatis bibendum sodales. Praesent eu nisi eget tellus malesuada eleifend cursus molestie lacus. Phasellus molestie lorem ante, tristique commodo tortor convallis non. Nam rutrum erat mi, sed auctor mauris rutrum in. Donec et convallis sapien, tristique faucibus eros.

Integer imperdiet metus eget ipsum interdum, in volutpat quam vestibulum. Donec egestas risus nec tempor accumsan. Phasellus malesuada ligula gravida, luctus dolor a, varius lacus. Vivamus venenatis dolor ut turpis lobortis aliquam. Duis sollicitudin tincidunt odio, finibus lobortis magna suscipit a. Vivamus efficitur eleifend turpis quis tincidunt. Suspendisse vestibulum varius semper. Fusce molestie interdum lectus vitae commodo. Pellentesque eu mauris sed enim sollicitudin aliquet.

Donec lectus libero, sagittis ac faucibus et, finibus eget lectus. Proin varius lacinia odio, ac maximus dui faucibus in. Donec vel risus dui. Mauris sodales odio ut felis lacinia, a gravida magna vulputate. Vivamus sagittis enim sit amet ipsum accumsan bibendum. Donec imperdiet in turpis ac tristique. Mauris egestas, lectus sed dictum laoreet, leo ligula malesuada elit, at semper mauris nunc a justo. Mauris condimentum consequat ligula, quis gravida nisl sodales et. Integer posuere lacinia nibh et lacinia. Morbi sed mattis libero. Nulla malesuada, odio a lobortis lobortis, nulla purus pulvinar erat, vitae vestibulum ligula nulla nec felis.

Ut vel rhoncus turpis, sit amet maximus enim. Praesent at efficitur lacus, ac dapibus diam. Suspendisse a tellus at dui ultrices aliquam at eget nisl. Proin iaculis finibus vulputate. Aliquam pretium nec lectus vel pulvinar. Aliquam felis nibh, fermentum ut volutpat quis, egestas quis mauris. Maecenas elementum arcu nec tortor tincidunt, sit amet sagittis massa elementum. Nunc porttitor ornare enim ut congue. Nulla facilisi. Mauris eu porttitor nibh. Integer molestie pharetra velit quis scelerisque.
`

// ProjectIndex
var projectIndex string = `---
title: "Project index"
date: 2000-01-01
author: "John Doe"
description: "This is the projects index page"
image: ""
draft: false
---
# My Projects

Here is a list of the projects i built.
`

// Project
var baseProject string = `---
title: "Base Project"
date: 2000-01-01
author: "John Doe"
description: "This is a base post"
image: "/static/images/project.png"
draft: false
---
# Project 1

Details about the project

[Link To The project](https://www.github.com/)`

// Base template for HTML files

var baseTemplate string = `<!-- base.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="/static/css/styles.css">
    <title>{{.Title}}</title>
</head>
<body>
    <header>
<h1>GoAmber</h1>
    </header>
<nav>
<img src="/static/images/icon.jpg" alt="Logo" class="icon">
<h2>Loren Ipsum</h2>
<p>Web Developer</p>
<div class="nav-container">
           <ul class="nav-items">
            <li><a href="/">Home</a></li>
            <li><a href="/blog">Blog</a></li>
            <li><a href="/projects">Projects</a></li>
			</ul>
</div>
        </nav>
    <main>
<div class="content">
        {{.Content}}
		{{.List}}
</div>
    </main>
    <footer>
        <!-- Your footer content -->
        <p>&copy; 2025 My Website</p>
    </footer>
</body>
</html>
`

// fileserver for dev is the content of the serve.go file
var serveFile string = `package main

import (
	"log"
	"net/http"
)

func main() {
	// Define the directory to serve
	fs := http.FileServer(http.Dir("./output"))
	http.Handle("/", fs)

	log.Println("Serving on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
`
var styles string = `:root {
  --amber: #FFA726;
  --amber-dark: #F57C00;
  --amber-light: #FFCC80;
  --bg-light: #FFFFFF;
  --bg-dark: #333333;
  --bg-soft: #555555;
  --text-light: #F5F5F5;
  --amber-gradient-start: #FFB300;
  --amber-gradient-end: #FB8C00;
}

.icon {
    border-radius: 50%;
    width: 150px;
    height: 150px;
}

body {
    display: grid;
    grid-template-columns: 350px 1fr 120px;
    grid-template-rows: 70px 1fr 70px;
    grid-template-areas:
        "header header header"
        "nav main aside"
        "footer footer footer";
    background-color: var(--bg-dark);
    color: var(--text-light);
    height: 100vh;
}

header {
    grid-area: header;
    background-color: var(--bg-dark);
    color: var(--text-light);
    display: flex;
    justify-content: center;
    align-items: center;
}

nav {
    grid-area: nav;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    background-color: var(--bg-dark);
    color: var(--text-light);
    height: 100%;
}

nav p {
    margin: 0 20px;
    text-align: center;
}

nav a {
    text-decoration: none;
    text-align: center;
    color: var(--text-light);
    padding: 5px 50px;
    border-radius: 5px;
    margin: 5px 0px;
    width: 100%;
    background-color: var(--amber-dark);
}

nav li {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    list-style: none;
    margin: 5px auto;
 }


main {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    padding: 20px;
    grid-area: main;
    background-color: var(--bg-dark);
    color: var(--text-light);
    border-radius: 10px;
}

.content {

    display: flex;
    border-top-left-radius: 10px;
    border-bottom-left-radius: 10px;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    grid-area: content;
    padding: 20px;
}

.content-item {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: left;
    padding: 20px ;
    width: 80%;
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px;
    border: 1px solid var(--amber);
}

.content-item a {
    text-decoration: none;
    text-align: center;
    color: var(--text-light);
    padding: 5px 50px;
    border-radius: 5px;
    margin: 5px 0px;
    width: 20%;
    background-color: var(--amber-dark);
}

.project-image {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    width: 80%;
    height: 20%;
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;
    border: 1px solid var(--amber);
    padding: 20px;
}

footer {
    grid-area: footer;
    background-color: var(--bg-dark);
    color: var(--text-light);
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 0.8em;
    padding: 10px;
    position: relative;
    bottom: 0;
    left: 0;
    right: 0;
    width: 100%;
    height: 70px;
}`
