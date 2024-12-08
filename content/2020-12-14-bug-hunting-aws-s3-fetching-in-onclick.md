---
_id: "12ed51bf-559e-4c8c-80f6-88e7459b066d"
title: "Bug Hunting - AWS S3 Fetching in onClick"
published: "December 14, 2020"
slug: "bug-hunting-aws-s3-fetching-in-onclick"
description: "My final project for my coding bootcamp is an Invoice App"
categories: ["Career Changer", "Coding Bootcamp"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

My final project for my coding bootcamp is an Invoice App. The app is still being built, slowly but surely it will get to where it needs to be. You can read about Hiron (The Invoice Application) on my portfolio site in the projects section. _*The reference portfolio site is no longer active nor the app*_

All I want to do in this blog post today is explain an issue I was running into and how I seem to have solved it, even though the solution is not perfect.

One of the key features of the application is creating, storing and fetching Invoice PDF's from AWS S3 buckets. Creation of the pdf was easy enough. Storage was a little more complicated due to trying to understand blob's and passing an exported function a component. _If you have questions feel free to find me on Twitter (@BenAPatton)[https://twitter.com/BenAPatton]._

The more complicated part was retrieving the pdf and having it display.

I have a table that updates as you add invoices. In the table there is a pdf icon that has an async function attached to it to call that invoice id from S3 and display it. When you call that function it returns a url to you that is the url to the pdf. I tried a lot of different things, some I have not tried that may also work, such as ReactPortals, but here is what worked for me and it seems incredibly simple.

Solution: window.open("retrievedS3url", "\_blank")

So this does and doesn't work. It will open a new tab, close it, and give you a downloaded file. When you click the file, the new tab will open and your correct invoice will be displayed. So again, it is not perfect, but it does work for what I need it to do and it is not the least intuitive thing for user experience. It could be a lot better, I recognize that but for me, this was a big win.

Reflection: If you cannot figure out an issue to a bug, keep working on other things. Most likely something you do in another project will have direct or indirect application to whatever bug you are trying to solve. Be patient with yourself as you learn and grow in your skills.

Hope this helps someone out. Again if you have questions or comments feel free to email be at bass41992ben@gmail.com. Or you can send me a message on my portfolio site.

_The invoice app mentioned in this post is not in production yet. When it is I will update this blog._
