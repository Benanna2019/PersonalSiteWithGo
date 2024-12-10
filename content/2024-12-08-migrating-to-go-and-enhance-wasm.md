---
"_id": "cm4fo14be00000cmk3gd92rnc"
title: "Building a blog with Go, Enhance WASM, Extism & Fly.io"
published: "December 8, 2024"
slug: "building-a-blog-with-go-enhance-wasm-extism-fly-io"
description: "Why Migrate from JS to Go?"
categories: ["Go", "Enhance", "Extism", "Fly.io"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

![enhance](/static/images/enhance/GoEnhanceWASM.png)

One thing in software that is constantly intriguing to me is where software paradigms mix. Add to that some level of fringe software that offers some possibility and I cannot stop thinking about it.

Time passes both too quickly and too slowly for me to be sure how long ago I heard of Enhance WASM. Enhance WASM is a way for you to write Enhance SSR Web Components but then use them in any language, particulary through a tool called Extism.

Extism is essentially a way to compile WASM modules into a shared library that I can use wherever I want. In other words, Extism allows me to write plugins for different languages.

Originally I did this with [Elixir, Phoenix LiveView, and Extism](https://www.youtube.com/watch?v=LVlDhNxsSTQ). Doing this blew my mind. Being able to write code in two languages for one project in the same file that wasn't seperated into client and server projects was so fun. There is no real definitely to being a "real software engineer" but this felt like a personal definition I didn't know I had that was being defined as I built.

## Web and Performance

My primary language when coding is JavaScript. More than that it is using React and Typescript. Mentors and those who have been doing this for a long time have been cautioning me more and more about React's performance issues.

A key moment for me that really made me wrestle and think about what I know as a web developer was when I first started trying to use [Enhance](https://enhance.dev). Enhance is just web standards. You write server rendered html as custom elements and then you progressively enhance them with JavaScript where needed.

In a weird way, even writing that last sentence and understanding what I said is a win. Why? The easiest way to say it is that when I started writing code with enhance, I really struggled and felt like I was taking a step backwards.

### Fast Sites

Recently I have been paying a lot of attention to blog posts and things coming out of [Frontend Masters](https://frontendmasters.com). Frontend Masters is built with Vanilla JS for the client and Go for the backend. I don't know what their code looks like, I would like to, but know that is enough.

Combined this with the excellent timing that I was learning a little bit of Go while learning [Datastar](https://data-star.dev), it occurred to me that I could combine what I knew while working/learning Datastar with what I knew about writing Enhance SSR Web Components and Extism.

Again, I know this does't exactly feel fast at the moment but I am starting with the belief and end goal in mind that the platform will win. So I want to be all in on the platform and build for it to the best of my ability.

## Enhance, Extism, and Go

To be honest, at this point I don't know exactly what to share other than tech I used and some of how this works. My site is not as fast as I would like but the great part about this is that I think it will be overtime. This is a build in public adventure for me to build a very performant/fast site.

So, here are the things that I used:

- Enhance WASM - This allows me to write server rendered html as custom elements but then pass them to a templ component so that I can render what I need to when I need to. These two things actually work really well together, templ and custom elements. There was a bit of a headache initially trying to debug my custom elements when they were not rendering but you can reach for the trust "poor man's debugger", also known as the `<pre>` tag and write it like so `<pre>${JSON.stringify(someValue, null, 2)}</pre>` and then you can see the error.
- Extism - This is just an incredible tool made by the [Dylibso](https://dylibso.com) team for writing/using plugins in any language.
- Go - I need to and owe a much deeper dive into Go. Coming from a Node background, I think there are some really cool comparison opportunities between Go and Node. All in all, things are very very similar, at least to the scale I have right now.
- Datastar - Datastar isn't directly used yet but will be when I add things like filtering. Datastar is a hypermedia library that allows me to build my app with a hypermedia approach.
- Fly.io - Fly.io is where I have deployed my app. This is probably why my app doesn't feel fast currently I think I need to scale it to more machines so that one is closer to me but this is a good start.
