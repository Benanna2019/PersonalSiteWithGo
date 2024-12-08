---
_id: "d9f01c1f-aea7-4198-8eea-9fe0a24e203e"
title: "Two Forms of Pre-rendering"
published: "November 20, 2020"
slug: "two-forms-of-pre-rendering"
description: "This is originally a pre-rendered blog post from Next JS. I have added my own thoughts at the end of this blog"
categories: ["Javascript", "Coding Bootcamp", "NextJS"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

_\*I left both this post and the 'Static Generation v.s Server-side Rendering' post because personally, they are fascinating._

_Note: This is originally a pre-rendered blog post from Next JS. I have added my own thoughts at the end of this blog_

Next.js has two forms of pre-rendering: **Static Generation** and **Server-side Rendering**. The difference is in **when** it generates the HTML for a page.

- **Static Generation** is the pre-rendering method that generates the HTML at **build time**. The pre-rendered HTML is then _reused_ on each request.
- **Server-side Rendering** is the pre-rendering method that generates the HTML on **each request**.

Importantly, Next.js lets you **choose** which pre-rendering form to use for each page. You can create a "hybrid" Next.js app by using Static Generation for most pages and using Server-side Rendering for others.

Static Generation is pretty incredible. In the the coding bootcamp that I just finished, this was one of the problems we faced as our apps became more and more complex, that of page loading/data fetching. Static Generation makes this so quick and clean campared to having to load the HTML every time. It is a very nice feature of any app/website, especially when it is your own, when navigating to a different page is quick by design. It provides a much nicer user experience.
