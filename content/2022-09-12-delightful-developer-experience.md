---
_id: "1700db16-d213-4e1b-ad59-0f970b83a413"
title: "Delightful Developer Experience"
published: "September 12, 2022"
slug: "delightful-developer-experience"
description: "Building My Client's Website With A T3 Stack + Sanity.io"
categories: ["Developer Relations", "Developer Experience", "NextJS"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

![image](https://cdn.sanity.io/images/nfspldzq/production/f6f7a5296291b53f582c793de9c2803ede808284-1600x840.png?w=800)

There was a post/tweet/words on the internet by [@swyx](https://twitter.com/swyx) a while back on what he called _The Minimum Spanning Stack_.

I feel like I have found a good 'minimum spanning stack' for most of my freelance/client work. Most client work I take on, to this point, hasn't really had the need for complex user login and interactions. And adding that would be overkill. So hurray! No database for me. Just a CMS 🎉

## Up And Running Quick

There are many stacks out there, this just happens to be the one that I have chosen.

If you haven't heard of the [t3 stack](https://create.t3.gg/) from [Theo - ping.gg](https://twitter.com/t3dotgg), I would highly recommend checking it out. With the simply command `npx create-t3-app@latest` you can have any combination of the following tech ready to go in your app:

- [tRPC](https://trpc.io/docs/v9/)
- [Prisma](https://www.prisma.io/)
- [Tailwind](https://tailwindcss.com/)
- [Next Auth](https://next-auth.js.org/)

So for client projects, I don't need auth and I don't need a database. So I can just select tRPC and Tailwind when I am on the command line and have an app ready to go.

### Choosing Sanity As A CMS

I have used Sanity a fair amount at work. So I am familiar with it. That is the main reason I chose it. But there are a lot of cool things you can do with Sanity and for some reason I really love the GROQ Query Language.

Probably the nicest thing about using Sanity is the ability to pass it off to the client. Granted, I will still be asked to do a decent amount CMS stuff since a lot of my clients prefer to not jump into really technical stuff and some of them are older and don't want to mess around with computers. They just want to ask something to be setup.

So Sanity makes this really, really easy.

### Sanity With tRPC

If you are not familiar with tRPC, I would recommend watching this [Learn With Jason](https://www.youtube.com/watch?v=GryES84SSEU) video where he has the creator of tRPC on. Also watch this [Jack Herrington](https://www.youtube.com/watch?v=Lam0cYOEst8) youtube video on tRPC.

The beauty of tRPC is that it sits on top of [React-Query](https://tanstack.com/query/v4/?from=reactQueryV3&original=https://react-query-v3.tanstack.com/) or Tanstack-Query.

To simply quote from [tRPC's documentation](https://trpc.io/docs/v9/):

> tRPC allows you to easily build & consume fully typesafe APIs, without schemas or code generation.

tRPC removes a lot of boiler plate for code generation and schemas make your code base, much smaller 🎉

So, in combination with Sanity, I just create my tRPC router's and then in my next routes I can query them I have typesafe api requests to Sanity.

Here are a few screenshots of what a router looks like and what a query looks like:

#### Router Creation

This is what a getAllPosts router looks like:

![Screen Shot 2022-09-12 at 4.13.29 PM.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1663013615461/PugEsLzWF.png align="left")

#### Exporting Routers

![Screen Shot 2022-09-12 at 4.14.39 PM.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1663013764675/otbYztfac.png align="left")

#### Using Routers in Next Pages

![Screen Shot 2022-09-12 at 4.15.31 PM.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1663013738056/igWrlcQ2q.png align="left")

### Innovating Quickly

Due to the preconfigured, opinionated nature of the t3 stacks, I have so much ready to go for me. All I have to do is add in the sanity client and functions that get the data I want to get.

Other things that I use along with the stack that I think are worth mentioning:

- [Zod](https://github.com/colinhacks/zod)
- Components from [Brian Lovin's](https://brianlovin.com/) website. These are wonderful components. One thing I found out is that Sanity has a markdown editor which stores its content as a string. So if it is a post body that I want to render, I can just pass that `post.body` to the markdown component like `<Markdown children={post.body} />` and it will render the markdown for me.
- [Nextjs](https://nextjs.org/) - The t3 stacks run on Nextjs. This was nice for me to finally get back into some development with Next.
- [Tailwind UI](https://tailwindui.com/)

As always, and as all good developers do, I would like to claim that this is the stack I will stick too. But I don't know. I tend to find things week-to-week that I want to try and love something so much I want to integrate.

So here's to keeping it simple.
