---
_id: "d2eb0603-310a-4cf5-9131-9c20aafd0230"
title: "Svelte - More Thoughts"
published: "February 14, 2021"
slug: "svelte-more-thoughts"
description: "Learning Svelte "
categories: ["Svelte", "Junior Developer", "React"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

This is a follow up blog to a blog I wrote a couple of weeks ago where I shared some of my initial thoughts on Svelte. You can read that one [here](https://benapatton.com/posts/2021-01-31-svelte-initial-thoughts).

I have played around with Svelte a good bit over the past week and again, all I can say is wow. First, let me clarify some things and then I will go into detail on all of the new things I have tinkered with within Svelte.

## Some Clarification

I want my blog and the things I say to be helpful to whoever reads them. So let me make some clarifying points.

1. If you are just starting out and are looking to 'break into' the tech industry, I think you should primarily look at what jobs are in your area/what company you want to work for and learn those skills. Svelte is not mainstream enough yet for a job application.

   - Key Word: YET!

2. If you have already learned React or you already have a job, I think you should learn Svelte. There are many reasons for this, but primarily to start using it and also because it will be a big thing.

3. Lastly, and potentially most importantly, build projects using Svelte and show them off. If you are looking for a job, start with number 1. If there is a company you want to work for, learn their language. If you do not have a specific company in mind, I think you can learn Svelte and build your projects using Svelte. It is a much cleaner frontend framework to learn if you have the basics of html, css, and javascript down.

So that being said, we now move into more features of Svelte.

## Svelte Features

One of the really nice things, if you are coming from html, css, and javascript, or even if you have already learned React and Nextjs or Gatsby, Svelte does a few things which are really nice.

1. Everything is component based, which we are very familiar with in web development. However, with Svelte everything comes down to the individual component. So instead of writing classes on all your divs and then going to a css file (Home.module.css), all of your css is written in the component file itself. You write your script tags at the top, then your style tags, and then your html. You do not have to declare html at the top of the file. This makes for nice tracking of your components.

2. You can pass props but instead of setting props={props} you just pass {props} into the component. Then in that component you "export let props" at the top of the component within the script tags. Very nice and clean.

3. Baked in reactivity - This was the most time consuming piece of many of my projects and did not make it in other projects of mine because of the headache of implementing it within other frameworks. With Svelte, all I have to do is declare a variable in my script tag (for instance let name;) and then if I want to make it reactive, add a div{name}div and make an input with a property bind:value={name} and then as I type, the name value will be updated on the screen. No more long functions or useState, useEffect or useRefs. How beautiful 😭

4. Conditionals and Loops - If you have spent any time in React, a very common way to iterate over a component and produce more than one of those is to create the component and then when you are wanting it/them to show up, you can use the map function. Svelte simplifies a lot of this, even though you need to learn a little bit of syntax, it is still really straight forward. If you are wanting something to make something show up based off of a condition, you can simply write if statements. It would look like this:

```html
{#if condition}
<div>Do this thing</div>
{:else}
<div>Do this other thing</div>
{/if}
```

If you are familiar with if statements, the syntax is pretty easy to understand. if starts the statement, else checks for another condition, and if ends the if statement.

    You can also have 'each' statements which is basically a map function

but as you use them it becomes a more intuitive than a map statement. So say you have fetched some data, you could write `{each data.results as result}` which is like items.map(item => do this thing). So maybe a picture will be more helpful

![image](https://cdn.sanity.io/images/nfspldzq/production/c4cd683fa4313d0b6a3cc110a4144206d11de155-1280x440.png?w=800)

So as you can see, this is a very short code block but a lot of things
are happening.

- We are awaiting a response that we have stored in a variable quiz at the top of the code (you cannot see it in the picture).We then set that to data. Then in the each statement we find each data.results and set it to a sudo name, question. Then you can nest an if statement to do some action. Finally we break out of all of that. This is, to me, is really intuitive. I like it 🤯🤓

5. Animations/Transitions - Now, I would have never considered myself a 'visual' kind of guy but being married has taught me that I am much more of a 'visual' aesthetic personal than I thought myself to be.

   _I mean I married my wife so that ought to tell me enough! This is a Valentine's Day post after all._

   That being said, one of the things I wanted to work with and add to my portfolio was some animation. This was incredibly difficult to figure out. I had just finished bootcamp and I was trying to learn some new things to add to my portfolio. With React, my gosh it was complicated for me. It was a good struggle, meaning that I learned something through struggling through it. But then, Svelte.

   Sarcasm: I don't know whether to be frustrated or sad at how easy Svelte is.

   Svelte comes with transition and animation baked in. Meaning, all I have to do is import it the animations/transitions from svelte and then call them as a property on the element I want them to alter.

   Setting up animation in React = 8hrs of work as a newbie
   Setting up animation in Svelte = 7min!!!!! While learning Svelte.

   I had plenty of use with React and it took me that long. I am going through tutorials and building my own stuff with Svelte and it is a simple import fade from svelte and then add it to an element. No more DOM stuff with useRef and useEffects. Cue the dramatic music and camera zooming in on my face.

Anyways, Svelte continues to be 🔥

The next steps for me with Svelte are:

- Routing in Svelte using tinro or sapper just to test out some different ways to do routing.
- Refactor this blog to use Svelte
- Build a 'fullstack' application with Svelte
- Build my site that 'is' my portfolio and has other things on it, such as the blog, projects section, interactive page with visitors, etc.
- Learn Sapper, which will become SvelteKit, for future projects.
- Play around with Snowpack and Svelte (a blog to come soon as I understand it more in depth).

This is just the beginning for Svelte and I am excited to see where it goes.

Till next time,

Ben
