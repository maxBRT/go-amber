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
title: "Base Post"
date: 2000-01-01
author: "John Doe"
description: "This is a base post"
image: "/satic/post.png"
draft: false
---
# Post 1

This is the first blog post.`

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
    <title>{{.Title}}</title>
</head>
<body>
    <header>
        <nav>
            <!-- Your navbar content -->
            <a href="/">Home</a>
            <a href="/blog">Blog</a>
            <a href="/projects">Projects</a>
        </nav>
    </header>
    <main>
        {{.Content}}
		{{.List}}
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
