---
"_id": "cm4m3qp6h00000cmndand9ryl"
title: "Plugins Everywhere"
published: "December 12, 2024"
slug: "plugins-everywhere"
description: "Once you see a plugin, you'll see them everywhere."
categories: ["Web", "Design", "Design Systems", "Plugins", "WASM", "Extism"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

This is an initial entry just to start on a simple journey of explaining how I have been using Extism.

I don't find many opportunities, none in fact, to use WASM at work. But the more I have thought about it the more opportunities I see or feel that WASM could be used in all sorts of situations.

## Extending Software w/ Extism

I have written about this blog migration to Go using Extism and Enhance WASM before. I won't go into incredible detail but the experience has me thinking, dangerous I know.

Simply put, Enhance WASM lets you write server rendered custom elements. By using Extism you can use these in any runtime currently supported by Extism, and there are many. I have written them in Go, Extism, and JavaScript. On the docket are Zig, Rust, C++, and C which will take a great deal of time.

But the point here is that through the use of WASM custom elements can be written and then read anywhere. With React 19 supporting custom elements and other major frameworks already having support, the ability is there.

So what can you write in one language that you want to use in another? If it can compile to WASM, you can use it anywhere. Create a plugin and make it available to any runtime that supports it.

## A Future Dream/Idea for Design Systems

The team behind Extism has created something called [XTP](https://www.getxtp.com/). XTP is a way to essentially create a registry of plugins and serve them to your app/software.

The idea is to create a enhance plugin/xtp plugin that anyone can install. But then they would be supplied with some type of admin UI or studio where they can create custom elements and upload them to the registry.

There are a few rough edges around the enhance WASM plugin, but the idea is that you can author elements that are uploaded to a registry of elements and then those elements can be called wherever. This is one of the rough edges. Currently custom elements used through a plugin cannot really make use of slots, at least in another language, and from my minimum experience.

The reason slots cannot be utilized is due to the element being returned as a string so the element cannot exactly be called like a typical custom element so content cannot be slotted.

Having a registry of custom elements through a plugin also opens the doors to write wrappers for apps, like `withXTPCustomElements` for a React app. Or you could load it via a script tag and then they are availble for use in the app.

What could you do through a plugin?
