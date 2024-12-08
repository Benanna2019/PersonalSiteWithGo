---
_id: "3f1d33c1-8cc1-41b0-9ec5-87f07101c32f"
title: "When to Use Static Generation v.s. Server-side Rendering"
published: "November 20, 2020"
slug: "when-to-use-static-generation-v-s-server-side-rendering"
description: "This is originally a pre-rendered blog post from Next JS. I have added my own thoughts at the end of this blog"
categories:
  ["NextJS", "React", "Javascript", "Coding Bootcamp", "Junior Developer"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

_\*I left both this post and the 'Two forms of Pre-rendering' post because personally, they are fascinating._

_Note: This is originally a pre-rendered blog post from Next JS. I have added my own thoughts at the end of this blog_

We recommend using **Static Generation** (with and without data) whenever possible because your page can be built once and served by CDN, which makes it much faster than having a server render the page on every request.

You can use Static Generation for many types of pages, including:

- Marketing pages
- Blog posts
- E-commerce product listings
- Help and documentation

You should ask yourself: "Can I pre-render this page **ahead** of a user's request?" If the answer is yes, then you should choose Static Generation.

On the other hand, Static Generation is **not** a good idea if you cannot pre-render a page ahead of a user's request. Maybe your page shows frequently updated data, and the page content changes on every request.

In that case, you can use **Server-Side Rendering**. It will be slower, but the pre-rendered page will always be up-to-date. Or you can skip pre-rendering and use client-side JavaScript to populate data.

As the article says above in things that can be generated statically, this blog is a an example of Static Generation. The plan is to use Static Generation along with Server-side Rendering in the future for more complex projects.
