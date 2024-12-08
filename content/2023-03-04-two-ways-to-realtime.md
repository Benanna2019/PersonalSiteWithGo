---
_id: "b7340045-7412-4685-b280-a55941039c1e"
title: "Two Ways to Realtime"
published: "March 4, 2023"
slug: "two-ways-to-realtime"
description: "Server sent events allow for realtime updates"
categories: ["Remix", "Supabase"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

![image](https://cdn.sanity.io/images/nfspldzq/production/95121332b3078455389aaedd6cfdc991a2a270a5-1380x776.png?w=800)

_Photo by [Garry Killian](https://www.freepik.com/author/garrykillian)_

_If you want to skip to the code/implementation, click [here](#examples)_

What comes into your mind when you think of realtime sites/applications?

Generally I believe we tend to think multiplayer applications. And by multiplayer, we can have multiple users on a screen doing a thing.

So like google docs, users in a waiting room, etc.

But there is another element with streaming data in realtime to the client, ie the user interface where users interact with data.

This problem seemed complicated until really recently, for me.

## Two ways of handling realtime

First, shoutout to [Supabase](https://supabase.com). Supabase, through its realtime set of features allows for you to 'listen' for postgres changes on a certain table and update that for everyone on the client. The other two features of Supabase Realtime are Presence and Broadcast. Both of these will be explored in a later blog post.

But suffice it to say, Supabase made this really easy.

Around the same time that I discovered Supabase's realtime offerings, I also discovered that my favorite framework for development, [Remix](https://remix.run) allowed you to do this through 'server sent events'.

So what is the difference between the two?

With Supabase, you listen for changes on your database directly. Wherever that database information is relevant in your client application (react application, svelte, etc) it will update.

Supabase relies on using realtime features on the client. So in React you use Supabase realtime methods/helpers to run code in a component.

With Remix, you set up a server file that receives and event.

So let's think of messaging.

In Supabase you might have a 'messages' table. With the supabase realtime methods you would listen or 'track' changes in that table.

In Remix, you still may have a messages table, but instead of calling code in the client, everytime someone submits an action where they add a message to the 'messages' table, you send an event.

Then, you just use a Remix helper to revalidate the data that is being loaded in the server and your data that's handed from server to client is refreshed and updated wherever relevant.

## Examples

### Supabase

_The example below is largely from [Jon Meyers](https://twitter.com/jonmeyers_io) [egghead course](https://egghead.io/courses/build-a-realtime-chat-app-with-remix-and-supabase-d36e2618)_

With Supabase, in your react app we have a component that handles tracking changes in our messages table.

```javascript
import { useOutletContext } from '@remix-run/react'
import { useEffect, useState } from 'react'
import { SupabaseOutletContext } from '~/root'

interface Message {
  content: string
  created_at: string
  id: string
  user_id: string
}

export default function Messages({ messages }: { messages: Message[] }) {
  const [messagesToShow, setMessagesToShow] = useState(messages)
  const { supabase } = useOutletContext<SupabaseOutletContext>()

  useEffect(() => {
    setMessagesToShow(messages)
  }, [messages])

  useEffect(() => {
    const channel = supabase
      .channel('*')
      .on(
        'postgres_changes',
        { event: 'INSERT', schema: 'public', table: 'messages' },
        (payload) => {
          setMessagesToShow([...messagesToShow, payload.new as Message])
        },
      )
      .subscribe()

    return () => {
      supabase.removeChannel(channel)
    }
  }, [supabase, messagesToShow, setMessagesToShow])

  return (
    <div>
      <ul>
        {messagesToShow.map((message) => (
          <li key={message.id}>{message.content}</li>
        ))}
      </ul>
    </div>
  )
}

```

Above, the main idea here is that we are passing this Messages component some messages we are getting from our Server. In this case we are getting them from a loader in a Remix route.

You want to have a supabase client in some way, here we are having to use remix auth helpers from Supabase. 🙏 Supabase

Overall, we have a component that when it mounts checks two things.

1. It keeps a running tab on the incoming messages. Thats the first useEffect. As those messages change it will setMessagesToShow to the incoming messages

2. We have a second useEffect that checks for postgres changes to our messages table. Specifically it checks for newly added messages to the message table, `event: 'INSERT'` and then sets our messagesToShow to the previous messages plus the new message that was inserted.

This is nice because all we listen for is if the messages we are loading on a server somewhere have changed, we set those, and then we are updating the UI state to reflect those server changes in realtime and send it to all our clients who have access to the messages table.

### Remix

Remix handles this totally server-side. The only client updates we do are revalidate if a server event is sent.

First we create an Event Emitter.

_Note: I had to do this by adding a global variable as seen below. Other examples don't do this but I could not get it to work otherwise._

This is in a file called services/emitter.server.ts

```javascript
import { EventEmitter } from 'events'

let emitter: EventEmitter

declare global {
  var __emitter: EventEmitter | undefined
}

if (process.env.NODE_ENV === 'production') {
  emitter = new EventEmitter()
} else {
  if (!global.__emitter) {
    global.__emitter = new EventEmitter()
  }
  emitter = global.__emitter
}

export { emitter }
```

Then, in as a route in remix we can create a route called `sse.messages.ts`

In routes/sse.messages.ts

```javascript
import type { LoaderArgs } from "@remix-run/node";
import { eventStream } from "remix-utils";

import { emitter } from "~/services/emitter.server";

export function loader({ request }: LoaderArgs) {
  return eventStream(request.signal, function setup(send) {
    function listener(value: string) {
      send({ data: value });
    }

    emitter.on("message", listener);

    return function cleanup() {
      emitter.off("message", listener);
    };
  });
}
```

Then in somewhere in our app we have a route where a user can submit a message to the database. In this route we are also displaying those messages as json.

```javascript
import { ActionArgs, json, LoaderArgs } from '@remix-run/node'
import { Form, useLoaderData, useRevalidator } from '@remix-run/react'
import { useEffect } from 'react'
import { useEventSource } from 'remix-utils'
import { emitter } from '~/services/emitter.server'
import createServerSupabase from '~/utils/supabase.server'

export let loader = async ({ request }: LoaderArgs) => {
  const response = new Response()
  const supabase = createServerSupabase({ request, response })

  const { data } = await supabase.from('messages').select()
  return json({ messages: data ?? [] }, { headers: response.headers })
}

export let action = async ({ request }: ActionArgs) => {
  const response = new Response()
  const supabase = createServerSupabase({ request, response })

  let { message } = Object.fromEntries(await request.formData())

  const { error } = await supabase
    .from('messages')
    .insert({ content: message as string })

  if (error) {
    console.log('error', error)
  }

  emitter.emit('message', message)

  return json(null, { status: 201 })
}

export default function HomeIndexPage() {
  let { messages } = useLoaderData<typeof loader>()
  let revalidator = useRevalidator()

  let lastMessageId = useEventSource('../sse/messages')

  useEffect(() => {
    revalidator.revalidate()
  }, [lastMessageId])

  return (
    <>
      <h2>Our messages will be below</h2>
      <pre>{JSON.stringify(messages, null, 2)}</pre>
      <Form method="post">
        <input className=" m-2 rounded-md ring-2" type="text" name="message" />
        <button
          className="rounded-lg border-2 border-slate-700 px-2"
          type="submit"
        >
          Submit
        </button>
      </Form>
    </>
  )
}
```

The flow is a user submits a message. The message goes to the action function. If the message is submitted correctly to supabase/our database, we 'emit' an event on our event emitter.

Then in our page, we load the messages from the server

Our lastMessageId is where we are subscribing to events from the sse.messages.ts route. Then in that useEffect, we are revalidating the server/loader data whenever the subscribed to server event (sse.messages.ts) receives an event.

Hopefully you can see how to subscribe to realtime events in both these ways.

The Remix version was initially a little harder for me to understand but both are really great ways to subscribe to database changes and provide realtime experiences to your users.

_Thanks [Jon Meyers](https://twitter.com/jonmeyers_io) for your [egghead course](https://egghead.io/courses/build-a-realtime-chat-app-with-remix-and-supabase-d36e2618). Thanks [Sergio](https://twitter.com/sergiodxa) for your [Remix-Utils](https://github.com/sergiodxa/remix-utils) package. Thanks [Supabase](https://supabase.com) and [Remix](https://remix.run) for an amazing developer experience._

_- Ben_
