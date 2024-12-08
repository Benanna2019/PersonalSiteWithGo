---
_id: "ef515254-74b1-445a-b029-b41bedae16b4"
title: "Adding Typescript Goodies to Supabase select"
published: "March 15, 2023"
slug: "adding-typescript-goodies-to-supabase-select"
description: "Adding typescript autocomplete to supabase select function"
categories: ["Supabase", "Typescript"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

![image](https://cdn.sanity.io/images/nfspldzq/production/34e19071092e9a78b57770f2b4dd29d17f0909bb-500x500.png?w=800)

The past week I saw someone wanting to get autocomplete on Supabase's select function. So I got a curious as well.

**Important Notes**

The solution below currently only works with columns from one table. It does not:

- Allow for multiple tables
- Allow for join queries in the select
- Im working on it!
- If you just want the code, click [here](https://benapatton.com/posts/2023-03-15-adding-typescript-goodies-to-supabase-select#just-the-code)

## Problems

Here are the biggest struggles I had with this:

1. I could create a typescript union but I always got a full string of all possible options to send to the select function.

   - This isn't what I wanted because someone may only want to select one or two columns not all of them.

2. I was getting close with a few different typescript types but it was repeating a bunch of types like this

```typescript

type SelectTypes = 'content,id' | 'content,content' | 'id,content'.....

```

Clearly that is not what we want. We only want each pair once.

### Walkthrough

So if we have a type like this:

```typescript
type DBColumns = {
  id: number;
  content: string;
  created_at: string;
  user_id: string;
};
```

The types are not what is important. What we want is to take each column identifier and create a union of string types that are comma separated values for all possible combinations.

The result would look something like this:

```typescript

type UnionOfUnionTypes = 'id' | 'content' | 'created_at' | 'user_id' | 'id,content' | 'id,created_at' | 'id,user_id' | 'content,created_at' | ......

```

In the end (cue Linkin Park's _In The End_), a combination of Stackoverflow mixed with my own work & types generated from my database using the Supabase CLI.

First, we get our database types. I enjoy is using the supabase cli. The documentation on installing the supabase cli can be found [here](https://supabase.com/docs/guides/cli)

After that it installed you can run a command like this to generate your database types into a file. Below I am placing my types inside of my project `types` folder in a the file `db_types.ts`

`supabase gen types typescript --project-id [projectid_goes_here] > ./types/db_types.ts`

Next, we can create a file in the types folder called db_type_helpers.ts.

In that file start by importing the database types that we generated above.

```typescript
import { Database } from "./db_types";
```

Next, we want to get the keys, or columns in this case, of a selected table that we will pass in as a generic. We also know that we are going to need to get the columns of that table.

One thing we run into with this is that when you nest down into the database types generated from Supabase, the way you select a table is with array indexing, like this:

```typescript
type Profile = Database["public"]["Tables"]["profiles"]["Row"];
```

You have to hardcode the table name to index down into it. This is what we want to do in a more generic way.

So here is what we now have:

```typescript
export type SelectKeys<
  TableName extends string,
  DBTable extends Database["public"]["Tables"]
> = TableName extends keyof DBTable ? DBTable[TableName] : never;

export type GetColumns<TableName extends string> = keyof SelectKeys<
  TableName,
  Database["public"]["Tables"]
>["Row"];
```

Let's explain.

In SelectKeys we are passing a TableName as a generic and we are making sure our DBTable is indexing down into our Tables by extending `Database['public']['Tables']`.

Then we are going to check that the generic TableName is a table in our database by passing it as the 3rd index

Next in GetColumns, we are going to pass that TableName to the SelectKeys generic we just made and pass Database['public']['Tables'] as the second generic to SelectKeys. Then we want to get the ['Row'] off of that table if it does exist

If we stop there we will get something like this: \_this example is using a messages table that has id, content, user_id, and created_at columns.

![image](https://cdn.sanity.io/images/nfspldzq/production/b521d032feab8041aba948b8ae57444a9ab3d07c-942x278.png?w=800)

So we have a union type of all of our columns. Now, here comes the magic that would take another 4ish blog posts to explain well.

The following types, in summary, take a type and return our desired format of

`'id' | 'content' | 'created_at' | 'user_id' | 'id,content' | 'id,created_at' | 'id,user_id' | 'content,created_at' | ......`

Again, I did not write this, I got this from [this post](https://stackoverflow.com/questions/59471947/define-a-typescript-string-type-of-comma-separated-union-types) on stack overflow

```typescript
export type UnionToIntersection<U> = (
  U extends any ? (k: U) => void : never
) extends (k: infer I) => void
  ? I
  : never;

type UnionToOvlds<U> = UnionToIntersection<
  U extends any ? (f: U) => void : never
>;

type PopUnion<U> = UnionToOvlds<U> extends (a: infer A) => void ? A : never;

// I changed the name here to SelectQueryStringBuilder just for
// consistency but this is where whatever type we pass in as a generic // will get converted to the comma separated values we are looking for

export type SelectQueryStringBuilderType<
  U extends string,
  Sep extends string
> = PopUnion<U> extends infer SELF
  ? SELF extends string
    ? Exclude<U, SELF> extends never
      ? SELF
      :
          | `${SelectQueryStringBuilderType<
              Exclude<U, SELF>,
              Sep
            >}${Sep}${SELF}`
          | SelectQueryStringBuilderType<Exclude<U, SELF>, Sep>
          | SELF
    : never
  : never;
```

Lastly, we finally get to use it.

```typescript
export type MessageSelectCombinations =
  | SelectQueryStringBuilderType<GetColumns<"messages">, ",">
  | "*";
```

When we use it we see something like this:

![image](https://cdn.sanity.io/images/nfspldzq/production/1e2fa40dac64085a88cb574cdc51d3e88f8940af-1966x276.png?w=800)

### Just The Code

```typescript
import { Database } from "./db_types";

export type SelectKeys<
  TableName extends string,
  DBTable extends Database["public"]["Tables"]
> = TableName extends keyof DBTable ? DBTable[TableName] : never;

export type GetColumns<TableName extends string> = keyof SelectKeys<
  TableName,
  Database["public"]["Tables"]
>["Row"];

export type UnionToIntersection<U> = (
  U extends any ? (k: U) => void : never
) extends (k: infer I) => void
  ? I
  : never;

type UnionToOvlds<U> = UnionToIntersection<
  U extends any ? (f: U) => void : never
>;

type PopUnion<U> = UnionToOvlds<U> extends (a: infer A) => void ? A : never;

export type SelectQueryStringBuilderType<
  U extends string,
  Sep extends string
> = PopUnion<U> extends infer SELF
  ? SELF extends string
    ? Exclude<U, SELF> extends never
      ? SELF
      :
          | `${SelectQueryStringBuilderType<
              Exclude<U, SELF>,
              Sep
            >}${Sep}${SELF}`
          | SelectQueryStringBuilderType<Exclude<U, SELF>, Sep>
          | SELF
    : never
  : never;

export type MessageSelectCombinations =
  | SelectQueryStringBuilderType<GetColumns<"messages">, ",">
  | "*";

export type ProfilesSelectCombinations =
  | SelectQueryStringBuilderType<GetColumns<"profiles">, "," | ", ">
  | "*";
```
