---
_id: "e4b8fd6d-3e40-45c5-8fa4-acf860bae08d"
title: "Coding Bootcamp - Week 3"
published: "August 31, 2020"
slug: "coding-bootcamp-week-3"
description: "This week was one of the more difficult conceptually"
categories: ["Career Changer", "Coding Bootcamp"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

_This is an article that I originally wrote for my Wordpress site but have moved here._

This week was one of the more difficult conceptually. I say that because I also started writing my functional programming blog which takes the concepts of conceptuality and gets more complex conceptually (if that was hard to follow….you understand).

_\*If someone reads this, wants to correct me/help me with this, feel free._

Anyways, this week we looked at the call-stack and asynchronous functions.

What I am about to say is my best attempt at explaining async functions as a newbie. You’ve been warned.

So, remember, all of this is within Javascript. Javascript reads all of its code synchronously. That means it reads it top down. This never changes within Javascript.

So async programming is basically programming that is being run somewhere else while Javascript finishes running its synchronous code. Asynchronous code is executed in the browser and even though it maybe finished running before all the synchronous code is finished, it is placed into the callback queue. This means that it will not run until all the synchronous code is finished and then the first async code in the queue will be executed first (FIFO – first in first out).

There are also these things called promises within all of this. Promises are more important than your typical async code and will be placed into the 'microtask' queue which is run before the callback queue.

So right now, as we are still very new to learning, our code can get a bit muggy trying to think through this whole call-stack, call-back, 'microtask' queue flow. But it’s part of the fun of learning to code (look for that blog to come soon).

The next blog in the bootcamp series will be the reflection on week 4 followed by “The Fun of Learning to Code.”

Til next time,

– Ben
