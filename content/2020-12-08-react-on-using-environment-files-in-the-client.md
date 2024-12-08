---
_id: "5565d121-5b67-4cfa-a919-57da147af789"
title: "React - On Using Environment Files in the Client"
published: "December 8, 2020"
slug: "react-on-using-environment-files-in-the-client"
description: "I needed environment files on a client side app"
categories: ["React", "Coding Bootcamp", "Junior Developer", "Career Changer"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

So I use React a lot and this week I have been working on my portfolio website, which you can view ---here--- (_that portfolio site has been taken down, you are now on the portfolio site da da duh!_). So here I want to share a specific issue and the solution that I found.

Summary of Issues & Solution:

Issue 1: I needed environment files on a client side app
Issue 2: You cannot store secret credentials in an normal js file. For security purposes you add the file to the .gitignore. However, when you deploy your app, most likely it will throw an error because the file you need is not in github, if you are using something like vercel or netlify that can connect with github.
Issue 3: I had to go back to using the environment but was not really sure how to get them since the normal way of "process.env.YOUR_ENV_SECRET" does not run.

Solution: Create React App Docs. Apparently, within a React app, you can add 'REACT_APP_YOUR_ENV_SECRET' and set it '=' whatever, and that is how you solve the problem. So when you 'call the information you need from your environment file, you call it 'process.env.REACT_APP_YOUR_ENV_SECRET'. If 'REACT_APP' is not at the front of whatever it is you have set your secrets equal to, then it will not be able to deply correctly.

USE CASE:

I think the uses of this can be many, and mine was is just one of those.

Like I said, I was working on my portfolio this week. I was adding the 'Contact Me' section to my portfolio and I was using Emailjs for managing those who so graciously would decide to send me a message. For Emailjs to work you need a 'service_id', 'TEMPLATE_ID' & 'USER_ID'.

I had this information safely stored in my .env file. However, when I went to deploy the update portfolio site it failed. Something was wrong with the .env information. So my first method was to change up where I was keeping my information. So I created a credentials.js file and stored the information there as an object that I could reference as 'credentials.USER_ID'.

This led to the second snag. For my credentials to be save I have to add the file to my .gitignore but here is the issue. Say credentials was in 'src/assests/credentials.js' my deployment scripts tell me that the file 'src/assests/credentials.js' is not defined, because they are in the .gitignore, and therefore it quits the deployment.

So, the fix I found was relatively simple, and as explained above, to use environment information (.env) in a Create-react-app, all you have to do is add 'REACT*APP*' in front of the 'variable' or 'key' that you are trying to access within your React App.

And Eureaka! After adding that it works flawlessly, well, that is until the next flaw.

Hope you find this helpful.

Till next time,

Ben
