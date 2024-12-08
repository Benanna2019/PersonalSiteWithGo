---
_id: "63544218-5dbf-4149-a4f4-76bec2b7791d"
title: "The State of the Application"
published: "August 22, 2022"
slug: "the-state-of-the-application"
description: "Seeking a More Perfect Understanding of State in Application Development"
categories: ["React", "Javascript", "Developer Relations"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

![image](https://cdn.sanity.io/images/nfspldzq/production/b5e3763ef17476d7e947ce642905a485a10c5b44-1600x840.png?w=800)

State always seems to be an elusive necessity in application development.

Maybe you are reading this because you are learning. Maybe you are reading this to see if I get it right. Maybe you are reading this because, like me, you have actually been developing for a little while and want to actually get a firm grasp on what application state is.

Whatever your reason for reading this article, I hope we can move toward a greater understanding.

## The Metaphors

As with all things that we seek to understand, it helps to have a reference point. That is, a mental model, a frame of mind where we can grasp 'this thing' conceptually and then move to a firm understanding.

[Maggie Appleton](https://maggieappleton.com/) has my favorite [article](https://maggieappleton.com/xstate) with metaphors for state. Her article is specifically looking at how xState is used to 'manage' state.

For our purposes, we can look at her light switch example.

If we think of a typical light switch, there are 2 states. Off and on.

This is fairly simple. If the switch is in the 'off' position, the state is 'off'. If you are looking to connect this to programming, think of toggles.

If you turn the light on, the switch has be 'toggled' to the 'on' position. And the state of the light is now on.

But let's add some complexity. Let's add the ability to dim the lights.

We can think of dimming as putting the light into a 'dim' state. Now a dimmer has multiple levels of 'dim', either on the brighter side or darker side. So the state is probably something like 'dim' with another state value on the dim scale.

Let's just think of the dim scale as being 2-9 since 1 would be 'off' and 10 would be 'on'.

_*Why the heck are we into all this with lights and dimmers?*_

Well the point is that state starts simple but easily becomes nuanced.

The larger point however, is that state is the 'state of a thing'. That may seem really confusing, using the word in its definition. But let's look quickly back at our light switch.

If the switch is 'on', the state of the light is on.

### Applications

When designing applications (the all encompassing term for user interfaces and user experiences), we think about state all the time, most likely without ever actually thinking about 'state' formally.

When a user interacts with something on the site, we want something to change. That is the simplest way to understand application state. If they are not logged in, we want a 'sign in pop up.'

You can experience this right now. If you have a twitter, go logout (if you remember your password to log back in). If you don't have one, perfect.

Just go to Twitter and start looking around. Stop and start scrolling on twitter. You don't get very far down the page before you get a pop up that basically shuts twitter down until you sign in or refresh to keep browsing. But you only ever get a certain amount of scroll time.

This is application state.

If we were to write comments about code we needed to write for this to happen, we can think about it like this:

```javascript
// If user is NOT logged in, when they scroll to this point,
// set pop up modal open state too 'true'.
// If there is no modal, the modal open state is 'false' by default.
// Or if we think about our light switch:
// no modal? open state is 'off'
// see a modal? open state is 'on'

// If user IS logged in, let them scroll to their hearts desire,
// be distracted, and get no work done today
```

### Conclusion

This may seem very trivial, but it has helped me tremendously.

I will elaborate on this further in the future because there is an important distinction to make with 'state'.

That future article will be on 'client state' vs 'server state'.

But for next time we will look at code examples and how they indicate changes in state. We will do this using simple javascript and then look at some react examples as well.
