---
"_id": "clnanxvkk000008ju1ww5d389"
title: "The Enhance Web Framework - Part 1: Background"
published: "October 3, 2023"
slug: "ehance-web-framework"
description: "Funtional web apps with apis as first class citizens"
categories: ["Framework Series", "Web Frameworks"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

I am web dev. I have altered that statement over time. I used to just say, "I am a react dev" because it felt better to specialize. But overtime I realized saying I am some specific framework/library dev makes you prone to take moral stances on a something that is simply a tool towards doing a job.

I love web development. It is so incredibly fun to think up of things and create. You think and you have ways to make something exist. That is an absolutely ridiculously awesome reality.

So awesome that my sentences make someone who cares deeply about word choice and structure probably cringe.

I remember a mentor telling me, "the most important thing you can learn/understand in web development is/are api's". At the time that did not make much sense to me. But it has become clearer and clearer over time that web development is api's all the way down.

None of this sufficiently paints the background I want but I want to share about a web framework I have been using some and where a lot of ideas, patterns, and concepts converge.

## The Convergence

One of the least well received Marvel films, Thor 2, has a cool concept in it called the convergence. It is when all the 9 realms align and the barriers between them are weakened.

This analogy goes further than I initially thought because in web dev world there are 'realms', like serverless, pure functions, frontend, backend, api's, and so on that each have traditionally had their own barriers to understanding.

Just take serverless for instance. I went down the serverless rabbit hole about a year ago. I still think it is fascinating. One of the first things you learn about serverless is that no one actually agrees on the meaning. The same goes for each of the other "realms". And we have reached a point in web development where the separation, or barriers if you will, are weakening and allowing for a convergence of paradigms.

## A little background

When I got my first fulltime developer we were using [architect](https://arc.codes). I remember at that time getting on a call with [Brian Leroux](https://twitter.com/brianleroux) and he walked me through what architect could do. Needless to say, I did not understand.

I would honestly say, for the first few years of web development, I have had this unbearable feeling that I am playing "catchup". I was always trying to learn something while never learning enough and simultaneously forgetting whatever it was I just spent time learning. Nothing clicked. I knew architect was powerful but I did not know why it was powerful.

Fast forward to that serverless journey I said I went on. I explored all the major ways to implement serverless architecture and I then understood why architect was so powerful. I honestly think it is so simple and powerful most doubt its capabilities.

I started looking into it more and Brian actually jumped on another call with me and talked to me about [Functional Web Apps](https://fwa.dev), again, something I did not fully appreciate at the time. But all of these things over time have started to make sense. I just had to give myself a bit more time and exposure to patterns and current practices to understand them. So here is the convergence of ideas I am referring too in graphic form.

![The Convergence of Ideas](/static/images/convergence/nine-realms.png)

[_Globe Image By brgfx on Freepik_](https://www.freepik.com/free-vector/earth-globe-cartoon-sticker-white-background_21196645.htm#page=2&query=globe%202d&position=3&from_view=search&track=ais)

As you will see with the metaphorical realms converging, there are probably more you want to add, I am just keeping it really basic. But the point here is that as technology has progressed there has been a lot more overlap between 'concerns'. For instance we used to, and many still do, separate concerns of their apps with different architecture patterns, like Model View Controler (MVC).

But as we have innovated, at least in the Javascript world, concerns have gotten a lot less 'separated'. I think this is a good thing because it used to be a lot of cognitive overhead to go find your model, view, controller, etc for any given thing you were working on especially as the application grew in size and inevitably complexity. So separation of concerns was nice but it ended up being _more_ work in many cases. And the lessening of separation is not just happening in the Javascript ecosystem, just look at Microsofts Blazor.

So what the heck does this all have to do with Enhance? Well, I have spent more time thinking about Enhance than I have any other framework on this framework series journey. The reason is due to the Functional Web App approach mentioned above.

## Functional Web Apps

I believe going into depth here is worth it for understanding how an Enhance app works. So let's look at each piece of what makes up a functional web app and try to digest it.

### Cloud functions

Functional web apps "are primarily composed of single-responsibility cloud functions". So first, what is a cloud function? Here we are referring to lambda functions. Lambda functions are a way to run some amount of code, large and small, in the cloud and you not have to provision or manage a server. You write the code, package it up, put it into a lambda function, and deploy it. And then whenever the function is invoked, the lambda starts up, runs the code, and then shuts down. You don't have to provision a server to handle load the lambda handles the load for you and then goes away.

As for payment, typically you would pay to use a server at some base cost. With lambda you pay for how long the function runs.

Up above I said, "Lambda functions are a way to run some amount of code, large and small", which is true. But if you notice, the primary purpose of a cloud function is _single responsibility_. So what does this mean?

I have tended to think of 'functions' in terms of something like helper functions in my app. But if we take a step back and look at how many people are deploying applications, they are deploying full applications through serverless/cloud functions. So function takes a bigger scope when you look at it in terms of where your application is deployed.

And if we think about typical applications, we will take a react app as an example, the app has components, which are composed together to build a large UI and then we may have a framework that sits on top of that and glues everything together. And when we are deploying a large application we go on adventures to find bugs and add features because portions of our app have no clear responsibility and it all bleeds together.

Just think for example, what do you use for error handling? Odds are as your app has grown you have an error handling service and you have a long list of bugs/errors that get reported because everything has bleed together and we, in a noble attempt to address, prevent, or find bugs early, have added error logging all over.

And then we have thousands upon thousands of errors in an application and so our adventures begin in the code base.

Single responsibility functions mean that your app is predictable. If you have error logging, you have it in each single responsibility function and you know exactly where that error happened.

So you have predictable code, you have better error handling, and you have code whose infrastructure (the lambda function), is managed for you. Friction starts to fall away 📉

### Managed Database

I would venture to say that most people love the thought of the first aspect of functional web apps, outside of those who get hung up on "serverless isn't really serverless".

But the managed database, ie Dynamo DB, is where most people will get hung up, at least frontend devs who are used to syntactic sugar like what Prisma, Drizzle, or Kysley might afford them when working with relational databases.

But fear not noble frontend dev, Ehance comes with begin's data packages `@begin/data` which is a wrapper around Dynamo DB that makes it actually a joy to work with. No more `aws-sdk` for dynamo! 🎉

So with that out of the way, what is a managed database? Well, just like how a lambda function is 'serverless' in that it is managed for you, Dynamo is a "serverless" database in that the database is managed for you. So no database administrators, no cost per database, etc. You pay for what you use and the database scales up and down with the demand of your application, ie _Managed Database_.

So if you can get passed the hurdle of working with Dynamo, of which I think the `@begin/data` wrapper for Dynamo should get you there, you are working with a managed database.

### Deterministic Deployment

Think of deterministic deployment as Infrastructure as Code (IaC). Infrasture as Code means that you define the cloud resources that are used through code. Most often this will be some high level format that sits on top of Cloud Formation. So in the case of Enhance you will notice that you have a `app.arc` file. Typically this file is where you define your cloud resources. When you use Enhance you may feel, like I have recently, that you are authoring code and not defining cloud resources. I ask the Ehance discord about this and got a really helpful response which I have summed up below.

With Enhance:

> "it is still IaC but it uses conventions on top of configuration with the ability to add to the configuration as you see fit"

The larger point here is that your cloud resources will be deterministic which means it is predictable. No more differences between dev, staging, and production environments. You have one environment and it is predictable.

## Conclusion

This article turned out to be much longer than I had originally intended. I do believe all the information here is good to have as background going into learning enhance though.

The convergence of spaces like cloud, frontend, backend, apis, all the things, along with the functional web app approach paint a good background to understand what Enhance is doing under the hood. And that is what I will write about in the next part.
