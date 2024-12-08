---
_id: "8a6071cc-c1c5-4518-a88e-2332a859ed60"
title: "Trekking In Serverless Territory"
published: "November 1, 2022"
slug: "trekking-in-serverless-territory"
description: "Making sense of serverless as a frontend dev"
categories: ["Serverless", "Draft"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

_A point of clarification to begin_

_I don't know many people personally working at AWS but I want to clarify I am grateful for everything being done at AWS. If you are reading this and you work at AWS I understand you are pouring your heart and soul into the vision and mission is and we are all grateful cause it is helping us move forward and ship more reliably._

_One last clarifier is this, my observation over 2 years has been to see so many frontend developers believe in the promise of AWS, serverless, cloud, etc, and still choose to use different technologies._

_The goal of this article is to get at the why and hopefully have a path forward._

In tech twitter, occasionally words start to pop up over and over. Whether it's the algorithm's due to one off searches or me finally noticing, I don't know. But what I do know is, like a dial turned to the top, eventually it gets so loud I cannot ignore it anymore.

_Serverless_ is the current radio station of my brain.

## My first run in with serverless

Ironically, my first run in with serverless came when I was going through my bootcamp. Our instructor loved/loves AWS and was/is really bullish on what their services will mean for the future of web dev.

And this was my first run in with the AWS console. If you know, you know.

For those that don't know. The console has the aesthetic of a legacy application from the late 1990's early 2000's. Think something like Microsoft Works with a more modern edging and better spacing.

Enough ragging on the console. But the point here stands, AWS does not do itself any favors with developers (maybe more specifically frontend devs) adopting it because its services are clouded...😏...in eery UI and deep technical language.

So, I want to go back. I want to go to what I see as a core principle of the cloud that I don't know is referenced that much, other than in AWS service names.

## Understanding Elasticity

Elasticity is not a hard idea to grasp. We have it in many things. Elasticity, or a flexible & expandable-ness, has grown in it's use though.

You may first think of rubber bands. That's pretty common. They are flexible, expandable, and can take on shapes.

But now we have this elasticity in fabric. I remember when I first bought a pair of pants that had this elastic material built into the fabric. I was so happy I could squat and the pants not fall down at all. They were elastic and could stretch and move with my motions in space.

But remember for a moment, what was life like before that. Pants and clothing were stiff. They may look nice when you are standing or in pictures but they are uncomfortable.

If this analogy hasn't gone far enough, let's look to the Incredible Hulk.

Early on in the Bruce Banner's tenure as the Hulk, I have to imagine it was incredibly annoying to transform back to Bruce Banner from the Hulk to find your clothes have been shredded and no longer fit.

As soon as Banner transformed into the Hulk, shirt ripped and pants ripped below the knee. Banner had to find new outfits.

Fast forward a bit too when Banner and Hulk start collaborating, more specifically when we have Professor Hulk, we see that he has Elastic clothing. Now he can transform either way and clothes still fit! Worth a celebration 🎉

### Amazon Elastic .....

There are a slew of services provided by AWS that contain Elastic in their name.

- Elastic Beanstalk
- Elastic Cloud Compute
- Elastic Load Balancer
- Elastic Block Store
- Elastic File System
- Elastic Container Service
- Elastic Transcoder

And many more!

The point is, you see elasticity everywhere in AWS. Why?

_Scale on demand_

All I have said so far is to get to this. Think one more time of clothing.

Before you had elastic built into your clothing, you had to have multiple pairs of jeans, shirts, shorts, etc depending on weight gain/loss (like all analogies, this fails at some point).

And with that, we get to serverl...wait, we have to talk about servers.

### Servers

So, in a super simple, silly trite example, we can think of servers as pants pre-elasticity.

But here is the very nuanced issue I ran into as a 'Fullstack Javascript Developer'. I only thought of servers in terms of a node server, or whatever web server we had that we were running.

Now this isn't wrong, it is just incomplete and it led to a terrible time figuring out this whole serverless thing.

#### What is serverless?

The sooner you can realize that there are still physical items, in the physical universe, called servers that exist in a real place on a map in a Data Center, the sooner you will have an easier time understanding serverless.

The problem, as a new traveler on this web developer journey, is that in the world of frontend, backend, database, etc, the only run in I had with a server was my node server.

So when we talk about 'serverless' I am thinking that I don't have to have a node server.

ERRR! Wrong!

Serverless has everything to do with physical servers, or infrastructure, and who owns/runs/handles those servers.

I wasn't ever really around for this or in a space where this was necessary, but before the cloud, companies had to buy physical servers and take care of that infrastructure.

So say a company launched their product and demand was off the charts. They would have to go physically buy servers, provision them, and then have them sitting around just in case demand went wild again.

But then what happens if the demand was low? They were paying for both the servers they owned as well as paying for personel (physical people living life needing a job who would show up to work and be security for the servers, etc, etc).

So, as many have said before, _serverless is a bit of a misnomer_ and dang it, that is pretty frustrating.

But here is the promise or beauty of serverless:

> You, the developer, the company, the startup, do not have to own any physical servers. You can rent the compute in an on demand model so cost rises and falls with said demand.

"So", you say, "I don't care about those things already." Which, I concur, I didn't believe I did either.

But the more I looked into this and tried to understand companies business models, you/I (the frontend devs) would probably 'care' if you/I understood what we pay for with hosting.

#### Hosted Servers

This is a section where I am a little out of my league and I believe I could really do to have some constructive feedback. but here is the question that I ask myself, "When I deploy and app, how do I pay?"

For many places/things, if you want to do anything other than have a random small, free app that does virtually nothing, you have to pay a monthly price.

Think Heroku before they got rid of their free tier. To do anything meaningful you had to pay a based amount to "provision dyno's".

I still don't know what dyno's are but the point is that this is not necessarily serverless because of the lack of on demand. You paid a pre-provisioned amount of money and then you can scale from their with some level of on demand.

One key difference between AWS and all of these other platforms is that with AWS you pay based off of usage, always. You can select other payment models but that we don't have time to dive into that.

Many companies offer free tiers up to a certain amount of usage or they offer free 'personal' plans with paid business/enterprise plans.

However, the big selling point of many of these hosting companies is that they 'wrap the cloud' or take care of cloud deployments so you don't have too. So now you are paying for a service rather than a thinking about scale and demand. And since someone is providing convenience for you by handling cloud stuff, you will be charged significantly more.

"Ben, what does all of this have to do with serverless and me as a frontend dev?"

Well, honestly, its probably more detail than you want but the point is, to understand serverless, we have to be a little clearer these days than just saying, you don't pay for your own infrastructure.

To someone who was in IT, I think that makes a lot of sense. To all the new devs entering the scene from bootcamps and colleges, I don't know that they understand these things. Not in an intelligence capacity but more so what we frontend devs know and spend our time doing from bootcamp onward.

_I know because I was/am one of these developers_

### Some more context

I have hung around, in a fly on the wall capacity, some of those who really are pushing the infrastructure as code paradigm. The pitch for infrastructure as code is something like:

> AWS is cheap and is going to keep getting cheaper as the other providers grow and add features. Going serverless will be the cheapest option with on demand scalability where you only pay for function invocation

Hire me amazon! I can give a sales pitch. 😂 I kid.

On a serious note, I have been wondering this very perplexing question:

> The promise of serverless is a beautiful promise. Why do so many of us in the frontend developer ecosystem not seem to care?

Reasons abound and we have all either used them or heard them ad nauseam.

- AWS is evil.
- Bezos is industrializing tech and industrialization is the enemy of innovation
- Bleeding eyes on the console
- AWS billing horror stories

And on and on it goes.

But here is the ever living tension:

> AWS really can give you virtually any scale at fractions of the cost. Yet we choose other technologies because I can do something fast _without_ having to use AWS, regardless of downstream effects.

Why? My conclusion is this:

_*The way we learn frontend is through incredibly fast feedback loops and when faced with serverless, therefore AWS, we are like a child standing before a Titan*_

We are small and the complexity of AWS makes it incredibly hard see the benefits as worth it.

But let's narrow down a bit more.

We now have frameworks and abstractions. And I do value this. I have looked into [serverless.com](https://serverless.com) and [arc.codes](https://arc.codes).

However, you still must be somewhat familiar with AWS. Maybe you, like I do, feel this pain. You have started down the AWS path. You have attempted to use something. But then you still have to jump into the console and know something about using the console, even if it is just to navigate to billing or cloudwatch to understand what is happening with your application.

More than that, frontend devs are incredibly uncomfortable, in my experience, with AWS to due horror stories of billing/costs, regardless of the fact that it is really cheap. Why? Mostly due to experiments gone wrong. So if someone abstracts AWS for me there is this nagging thought that I am going to end up with hidden costs.

_Fear_, or in a real sense, poor developer experience.

Maybe that is the sum of all of this so far, that AWS is shrouded in fear for frontend developers.
