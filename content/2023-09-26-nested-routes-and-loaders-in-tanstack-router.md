---
_id: "de16e7a6-ab0a-41e1-8700-c68d6f333c11"
title: "Nested Routes & Loaders in Tanstack/Router"
published: "September 26, 2023"
slug: "nested-routes-and-loaders-in-tanstack-router"
description: "Understanding nested routes and loader patterns in tanstack/router"
categories: ["Typescript", "React"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

In the [last article](https://benapatton.com/posts/2023-09-21-getting-started-with-tanstack-router) we created our app and created some simple routes. Now we want to start seeing how we turn our app into something more usable.

## Nested Routes

_Note: The code here is taken from the tanstack router docs from the basic example_

You may be used to nested routes mostly through file structure routing or through React Router's Route components.

Here is how you setup a nested route in tanstack router.

### Declare Routes

Let's first create a route for some posts. Would this be a walkthrough/tutorial if it didn't have posts?

```typescript
const postsRoute = new Route({
  getParentRoute: () => rootRoute,
  path: "posts",
  key: false,
  loader: fetchPosts,
  component: ({ useLoader }) => {
    const posts = useLoader();

    return (
      <div className="p-2 flex gap-2">
        <ul className="list-disc pl-4">
          {posts?.map((post) => {
            return (
              <li key={post.id} className="whitespace-nowrap">
                <Link
                  to={postRoute.to}
                  params={{
                    postId: post.id,
                  }}
                  className="block py-1 text-blue-800 hover:text-blue-600"
                  activeProps={{ className: "text-black font-bold" }}
                >
                  <div>{post.title.substring(0, 20)}</div>
                </Link>
              </li>
            );
          })}
        </ul>
        <hr />
        <Outlet />
      </div>
    );
  },
});
```

What is happening? Well, it is pretty straight forward.

1. We are declaring a new route like we have already done
2. This route has a parent route and we declare that by passing a property called `getParentRoute` and then passing it a function that references another route. So here we are doing like we did with our about and index routes in the first article
3. Now we see our first instance of something new, _THE LOADER_

Having worked in Nextjs and Remix for a while, the loader in tanstack router is closest to a Remix loader. The main difference is that Remix's loader, due to being a framework, is on your server.

This means that loaders in tanstack router need to interact with your server through use of a fetch, axios, or trpc function or something similar.

So here we simply has a fetchPosts function that reaches out to a json placeholder api.

So great, we use that and have some data that is returned from a the fetch function. Now we need to get that data.

4. The component property, assuming you use a loader function, comes with a useLoader hook that you can use to get the data.

Whew! We made it. Now, let's setup a nested route.

```typescript
const postRoute = new Route({
  getParentRoute: () => postsRoute,
  path: "$postId",
  loader: async ({ params: { postId } }) => fetchPost(postId),
  errorComponent: ({ error }) => {
    if (error instanceof NotFoundError) {
      return <div>{error.message}</div>;
    }

    return <ErrorComponent error={error} />;
  },
  component: ({ useLoader }) => {
    const post = useLoader();

    return (
      <div className="space-y-2">
        <h4 className="text-xl font-bold underline">{post.title}</h4>
        <div className="text-sm">{post.body}</div>
      </div>
    );
  },
});
```

### Dynamic Routes

Everything here should look familiar from the first example. But let's narrow down on the path and the loader.

If you look at the path definition, the path definition is a 'dynamic' parameter. So the literal path for this route, because the parent route is `postsRoute`, would look like this `example.com/posts/$postId`.

This $postId portion of the url though will always be a different value. So you would see something like `example.com/posts/first-post` and so on.

So whenever we hit that url in our application we need to get that parameter value and then use it to display only the post information that we want to display.

Now if we look at our loader property, we see something different:

`async ({ params: { postId } }) => fetchPost(postId)`

On a dynamic route we get access to the params. The params is whatever follows the $. So here it would be postId. Thus we can de-structure the postId off of the params and then use that to fetch a single post information

🎉 We have created a nested route.....BUT WAIT!

We need to register them. As we have already seen we register routes in our route tree. So here we need to first register our postsRoute then our postRoute.

This is how we do this using the addChildren function on a route

```typescript
const routeTree = rootRoute.addChildren([
  postsRoute.addChildren([postRoute]),
  indexRoute,
]);
```

That is a pretty basic overall example. Here is a much larger route tree for an example app using tanstack router I built.

```typescript
const routeTree = rootRoute.addChildren([
  indexRoute,
  salesRoute.addChildren([
    invoicesRoute.addChildren([invoiceIdRoute, newInvoiceRoute]),
    customersRoute.addChildren([customerIdRoute, newCustomerRoute]),
    depositsRoute.addChildren([depositIdRoute]),
    subscriptionsRoute,
    salesIndexOverviewRoute,
  ]),
]);
```

You can nest routes as deep as you want to!

This post is longer than intended so in the next one I will look at the different type of loader using the Loader api and how to reference loaders across your application.
