---
"_id": "cm4owmhyq00000cmthtbif1df"
title: "Beginning with Local First"
published: "December 14, 2024"
slug: "beginning-with-local-first"
description: "Some thoughts on what it looks like to get started with Local First."
categories: ["Software Development", "Local First", "Web Development"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

Right now Local First development is in a very interesting spot. On one hand you have many people coming to the Local First space because it is highly academic. I mean I didn't know what CRDT's were and outside of the high level of a CRDT, I don't exactly know the inner workings of them.

So you have an academic side of the Local First community. You also have the idealized side of the Local First community. You can find a list of the 7 ideals of Local First development [here](https://www.inkandswitch.com/local-first/#seven-ideals-for-local-first-software/). Odds are you will find that some ideals resonate with you.

However, an increasing number of developers are coming to Local First due to realizing how simple it can make writing code.

The simplicity is due to syncing. In the web dev world syncing means doing reads and writes based off data in the browser and then that syncs to the server. There is infinitely more to it than that, but that is the basic idea.

What that is really solving for you is the problem of the Network. The network always throws an app into some state of trying to sync the browser/client with the database. So sync puts data in the browser and then makes sure it is up to date with the database by performing sync operations against the source of truth, the server. There are other implementations that also don't need a server, but we are keeping it simple.

## A Path Forward

I am writing this as part of a advent blogging thing so it probably deserves more time.

But I think a good path forward for anyone looking to start building with local first is something like a simple counter or todo app, but through all the stages we currently use to build apps:

A rought list would include:

- SSR Version - a form with a counter button (could probably just return it on the session/cookie)
- Client JavaScript version (state)
- Local storage for client persistence
- IndexDB for client persistence (for learning IndexDB)
- Syncing from IndexDB to server
