---
_id: "a4358d57-183e-43e3-8dd9-c0747f19b426"
title: "HTML forms with Supabase, RSC, & Nextjs"
published: "March 27, 2023"
slug: "html-forms-with-supabase-rsc-and-nextjs"
description: "React Server Components let us use old school forms with actions"
categories: ["Supabase", "NextJS", "Draft"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

![image](https://cdn.sanity.io/images/nfspldzq/production/4a7941fa92371ec9c9f2f52d26c15dcac5922ceb-1080x1080.png?w=800)

_**Update 1 - March 28, 2023**_: After talking with Dan Abramov on twitter, this post is now a draft, very much so in a draft state 😅. There are some cool things on the way for RSC to get this post to where it needs to be.

## What are we working on?

As the title might suggest, I want to talk about html forms in the new Nextjs app directory with server components.

I honestly am unsure if normal React allows you to post to an api route from a form but being that we have always had to the the whole `e.preventDefault()` thing, I am pretty sure it's at least not common.

A week or so ago, I read/heard something by Dan Abramov about React Server Components. He said:

> SSR is the 'client' to React Server Components.

Being that I have been using Remix for a while I thought about what makes forms unique, outside of their Form component with all the great stuff mixed into it.

I remember it being one of the early great things about Remix, that you could just add an action to a form and put a method on the form as well and not have to do the `e.preventDefault()` thing. But all of this is because things are much closer to the server. With React we had to do that because React runs in the client so we had to prevent the default behavior of forms because it wasn't as simple as just sending some action to a handler.

### Enter Server Components

So this is where my mind started racing. I asked something like, "aren't react server components basically like a Remix route.

Now before you continue, the answer is sort of and definitely not at the same time. You can do a lot of amazing things in a remix route with it's loaders and actions that you cannot do with react server components, yet. The RSC story for mutations is still being written and I think we will see a lot of things progress in the future.

### What does this all look like?

Well, I'm glad you asked. I am using [Supabase](https://supabase.com) in this example.

_Also important to note, I took learnings from [this walkthrough](https://supabase.com/docs/guides/auth/auth-helpers/nextjs-server-components) in Supabase's docs and [this tutorial](https://egghead.io/courses/build-a-realtime-chat-app-with-remix-and-supabase-d36e2618) by [Jon Meyers](https://twitter.com/jonmeyers_io) on egghead_

First we create our messages route.

In the new Nextjs app directory, for every route that is not the root route, you have to use the convention of creating a folder with the route name and a file called page.tsx. So it would look like this:

```bash
app/
   messages/
           page.tsx
   layout.tsx
   page.tsx

```

So we create a messages route. Here is what this file looks like:

```javascript
import { createClient } from "@/lib/playground/utils/supabase-server";
import { redirect } from "next/navigation";

export default async function Messages() {
  const supabase = createClient();

  const {
    data: { session },
  } = await supabase.auth.getSession();

  if (!session) {
    redirect("/");
  }

  const { data: messages, error } = await supabase.from("messages").select("*");

  return (
    <>
      <div>
        <h1>Messages</h1>
        <pre>{JSON.stringify(messages, null, 2)}</pre>
      </div>
      <form method="POST" action="/api/message-actions">
        <input type="text" name="message" />
        <button type="submit">Submit</button>
      </form>
    </>
  );
}
```

In this file we are doing a few things:

1. Check if there is a session. If not, we will redirect using the new Nextjs redirect function from `next/navigation`.
2. Next, we want to get the messages from the database.

   - If you have Row Level Security on this will return an empty array, so remember to write a policy that let's users (in my case authenticated users) read all messages from the database

3. In the jsx, you will see a form that looks different than typical jsx forms. There is no action! Golly gee willikers Batman, no onSubmit? Nope!

But why though?

Well forms with an onSubmit wouldn't work in a server component because that is a distinctly client side thing we have to do in React components.

So we can slap an action on this sneaky form and give it a send the form data to the action.

So in the Next app directory, we now have the `app/api` folder which allows us to create api routes. So in this case, we want to create a message-actions api route like this:

```
app/
   api/
      message-actions/
         route.ts
```

In that route this is what our code looks like:

```javascript
import { createClient } from "@/lib/playground/utils/supabase-server";
import { redirect } from "next/navigation";
import { NextRequest } from "next/server";
import invariant from "tiny-invariant";

export async function POST(request: NextRequest) {
  const supabase = createClient();
  const {
    data: { session },
  } = await supabase.auth.getSession();
  if (!session) {
    return new Response("Unauthorized", { status: 401 });
  }

  const formData = await request.formData();
  const message = formData.get("message");
  invariant(typeof message === "string", "name is required");
  if (!message || message === null) {
    return new Response("Bad Request", { status: 400 });
  }

  const { data, error } = await supabase
    .from("messages")
    .insert([{ content: message, user_id: session.user.id }]);
  console.log("data", data);

  if (error) {
    return new Response(error.message, { status: 500 });
  }

  return redirect("/messages");
}
```

Again, we are doing a few things here:

1. we are making sure that we are using our server instance of supabase.
2. we are checking for a session. We only want active session users to be able to send anything to our api route
3. we are getting our message from our form data and doing a series of checks to make sure that it is defined, not null, and that it is a string.
4. then we are sending the message from our form data as the content field and attaching the session.user.id to the user_id field.
5. assuming everything goes as planned and there are no errors, we want to redirect back to our messages page which will populate the messages with any newly added messages.

_Clarification_

You do not have to send the user_id. If you followed the other tutorials, particularly the egghead tutorial, you will have set up the user_id column to auto generate based off of the signed in user's id, which is Supabase is the `auth.uid()`

Then you will also have set up a RLS policy to only allow authenticated users to insert rows to the messages table if the user_id = auth.uid()

I found that the issue I was having where I thought it was a problem with supabase was a good ole typical case of not having restarted my development server after I made changes to the database.

After I did that, everything works smoothly.

### Conclusion

Hopefully all of that makes sense. If you are interested in a video, here is one I made as well to go along with this. Video is a little rough for me but hopefully it can be helpful if you need/want it - [Server Component forms w/ Nextjs app dir & Supabase](https://youtu.be/mSpiHtzmDIM)

### The second wrap up

Did you go through the word search above? Which words did you find? Here is what should all be there.

- [ ] - Supabase
- [ ] - React
- [ ] - Server
- [ ] - Components
- [ ] - Nextjs
- [ ] - Api
- [ ] - Route
- [ ] - Action
