---
_id: "166ad786-8146-49d4-91c8-1a29fa1383f0"
title: "Getting Started With Tanstack Router"
published: "September 21, 2023"
slug: "getting-started-with-tanstack-router"
description: "Get A Quick Start with Tanstack Router"
categories: ["Typescript", "React", "Developer Experience"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

## Initialize App

First we can run `npm create vite@latest my-tanstack-app -- --template react-ts`

After you run this command you make sure you `cd my-tanstack-app` and then run `npm install`

This has created a react app with vice that has configured typescript for us.

Now you can open up this app in your code editor.

Now to test, let's just add some code to our `src/App.tsx` file. Go ahead and add "Tanstack Router" to the end of the `<h1>"Vite + React"</h1>` like this `<h1>"Vite + React + Tanstack Router"</h1>`

Sweet! Onward!

## Install remaining dependencies

### Tailwind (OPTIONAL)

I like tailwind. So we will use tailwind in this tutorial

This is straight from the tailwind docs for installing tailwind with vite. You can check those out [here](https://tailwindcss.com/docs/guides/vite)

Here is what you need to run.

`npm install -D tailwindcss postcss autoprefixer`
`npx tailwindcss init -p`

Now modify the newly created tailwind config to look like this

```javascript
/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {},
  },
  plugins: [],
};
```

Lastly, alter your index.css file and add the tailwind directives

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

Now finally, since we are using React, go ahead and remove the `App.css` file and start using tailwind in your `App.tsx` file.

Your app should now have no styles and look less good than it did before 🎉

_Note: We are not going to style anything but I love tailwind and don't like jumping back and forth between files to do any css changes_

## THE TANSTACK ROUTER!!

Alright, finally, we are ready to install the router

Run `npm install @tanstack/react-router@beta`

Let's go ahead and delete our App.tsx. Back in our main.tsx file, go ahead and remove references to the App.tsx component.

Now add these imports in the top of your main.tsx file

```javascript
import {
  Outlet,
  RouterProvider,
  Link,
  Router,
  Route,
  RootRoute,
} from "@tanstack/react-router";
```

We will talk about each as we use them

### RootRoute

First, we need to create a RootRoute. The RootRoute is the entry point for your app.

Metaphorically, if we think of Roots, the Root is the supplier to the rest of the plant, or our route 'tree'.

So there must be a RootRoute. Let's look at some code

```typescript
const rootRoute = new RootRoute({
  component: Root,
});

function Root() {
  return (
    <>
      <div>
        <Link to="/">Home</Link> <Link to="/about">About</Link>
      </div>
      <hr />
      <Outlet />
    </>
  );
}
```

Our new RootRoute is basically that entry point for our application.

First you will see that our RootRoute has a component property. Routes have a handful of properties but this is what you would expect.

When someone lands on our app, example.com, we want them to see something. And, for the most part, when the user navigates through our app, at various 'routes' we want them to see different ui.

Alright, we get that. But what is this `<Outlet />`.

Here, our metaphor changes a bit. Instead of thinking about a root and plants, think of the root route as a store front.

You have a basic layout and idea of the application but you haven't gone in and experience anything yet. The door to the store front, that is an `<Outlet />`.

The building and store front are the root that wraps the inside of the store. You enter into the store through the door, or in our case, you enter into the rest of the application through the root routes `<Outlet />`

_The root route will wrap the entire application_

### Register our router

Before we move on we need to register our router and make sure things are working. That way we can see our app progress as we add routes

Add this below your Root function

`const routeTree = rootRoute;`

We are being premeditative by calling this a routeTree because we only have a root route. So don't worry about this yet. We will add to this. For now just know we are creating a route tree and our tree will expand as we add more routes

Next we need to actually create our Router and then Register it.

Add this code:

```typescript
// Create the router using your route tree
const router = new Router({ routeTree });

// Register your router for maximum type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}
```

All we have done is created a router that is based off of our route tree. We will not have to change anything about our router in the future. All we will alter is our routeTree

Next we registered the router and this gives the type-safety that we expect with Tanstack/router.

Lastly, change your createRoot to look like this

```typescript
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
```

You may notice you are getting an error in your component. If you look at your Link component that is a link to `/about` you will notice the error is on the `to` prop.

Why is this happening? Well, type safety. We don't have an about route. "Yeah but why is our other route not throwing an error?"

RootRoutes, without a specified index route, meaning a route with the path "/", will both evaluate to that root path of "/"

Just hold on! We will add a route right...about...

### Now!

Alright, so let's create more routes and make our Link happy

Let's go ahead and add this code to our file about the route tree variable

```typescript
// Create an index route
const indexRoute = new Route({
  getParentRoute: () => rootRoute,
  path: "/",
  component: Index,
});

function Index() {
  return (
    <div>
      <h3>Welcome Home!</h3>
    </div>
  );
}

const aboutRoute = new Route({
  getParentRoute: () => rootRoute,
  path: "/about",
  component: About,
});

function About() {
  return <div>Hello from About!</div>;
}
```

Okay, sweet. We have an indexRoute and an aboutRoute. But how do we add them to our route tree?

Well, here again our metaphors get a little complicated but that's alright, we are used to complication.

Look at your `const routeTree = rootRoute;`

As we said before, the rootRoute wraps the application. So now, our metaphor changes, due to api naming, and we can think of our rootRoute as the parent of your whole application.

Here is how add a route:

`const routeTree = rootRoute.addChildren([indexRoute, aboutRoute])`

Now if you save and go back to your app, you will probably notice something weird. Things are probably rendering double. Go ahead and remove React.StrictMode or comment it out and that problem will go away.

_**And if you see red error boxes after a change locally, do a hard refresh and they will go away, at least for this basic setup**_

### End of Basic Setup

That's it for this basic setup. In the next part we will add more routes, nested routes, and look at data loading.
