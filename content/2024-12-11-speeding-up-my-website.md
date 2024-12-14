---
"_id": "cm4kg7uxx00000cjtfmod41u7"
title: "Speeding up my website"
published: "December 11, 2024"
slug: "speeding-up-my-website"
description: "Using cache control headers and prefetching"
categories: ["Web", "Performance", "Prefetching", "Cache Control"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

_Just a fun little note, the snow fall effect is from [Zach Leatherman](https://www.zachleat.com/web/snow-fall/)._

## Cache Control Headers

Ok, so I moved my website to Go and I was initially pumped cause I was just under this default impression everything would be fast. However, there are some things that I had not taken into account and honestly had never needed or understood when to use them before.

My first experience with speeding things up was by using Cache Control headers. To be honest, most of my website is now becoming a blog version of [enhance-movies](https://enhance-movies.com). Since this content is mostly static, I can set the cache control headers to be an hour or more. I should probably set it for something around 24 hours since I don't write or update things that often but it doesn't really matter.

But I now have cache control headers set for my content so when you click and navigate you should notice they are quick.

## Prefetching

Sadly, Enhance always seemed a bit confusing to me but now that I have started learning a very different language from Node and building a website from scratch with it and server rendering content, I can see how things were working.

So I have added a browser folder in my code that contains code that will only run when js is enabled. So for this basic example I have added a prefetch.mjs, again taken from Enhance Movies code, and when js is present in the browser, it will prefetch the next page.

Combine that with Cache Control headers and you have cached content and prefetching so things speed up a good bit!

### Service Worker

The last thing worth mentioning is that I am working on adding service workers. Frankly, service workers are their own realm of web development and performance.

My only experience thus far with using them has been building out a Local First site that uses a database in the browser. But I am using sveltekit for that and sveltekit treats service workers as more of a first class citizen so it is not as convoluted.

Service workers are really well suited, at least to my current understanding, for offline support. This isn't the biggest deal for my website cause it probably isn't being accessed by anyone offline, other than me just for playing/testing around.

But for the Local First site, I am building a journaling app that needs to allow for offline reads and writes. So for something like adding a journal entry, you want to be able to add it and then dynamically update the UI or dynamically add routes so you can go see the newly added content.

### Local First

Speaking of Local First, if you are interested in learning more about it, or are interested in learning how to build in a Local First way, I am collaborating with a friend, [Kevin Cunningham](https://kevincunningham.co.uk/), to build a platform called [Local First Academy](https://localfirstacademy.com).

Feel free to sign up for our [newsletter](https://newsletter.localfirstacademy.com). We have some fun things coming in the holidays and looking into 2025.

If you have any thoughts or questions around Local First, please feel free to reach out on [Bluesky](https://bsky.app/profile/benapatton.bsky.social)
