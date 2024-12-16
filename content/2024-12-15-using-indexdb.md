---
"_id": "cm4pxfuup00000cjr7sav98op"
title: "Using IndexedDB"
published: "December 15, 2024"
slug: "using-indexdb"
description: "Using IndexedDB in a custom element"
categories:
  ["Software Development", "Local First", "Web Development", "IndexedDB"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
customElementKeys: ["counter"]
type: "Blog"
---

Click the button above to increment the count that is stored in IndexedDB on your browser.

You should be able to open the dev tools, go to the application tab, find the IndexedDB tab on the left sidebar, open it and you should see this in IndexedDB.

Below is the code for the ssr'd `<custom-counter></custom-counter>` custom element. This is learning to use IndexedDB so I can work towards a very basic sync example.

```js
function CustomCounter({ html, state }) {
  return html` <script type="module">
    import { openDB } from "https://cdn.jsdelivr.net/npm/idb@8/+esm";

    // Open/create the database
    const dbPromise = openDB("CounterDB", 1, {
      upgrade(db) {
        // Create an object store if it doesn't exist
        if (!db.objectStoreNames.contains("counter")) {
          db.createObjectStore("counter");
        }
      },
    });

    // Function to get the current count
    async function getCount() {
      const db = await dbPromise;
      return (await db.get("counter", "count")) || 0;
    }

    // Function to set the count
    async function setCount(val) {
      const db = await dbPromise;
      await db.put("counter", val, "count");
    }

    // Initialize the display
    async function initCounter() {
      const count = await getCount();
      document.getElementById("count-value").textContent = count;
    }

    // Handle increment
    document.getElementById("increment").addEventListener("click", async () => {
      const currentCount = await getCount();
      const newCount = currentCount + 1;
      await setCount(newCount);
      document.getElementById("count-value").textContent = newCount;
    });

    // Initialize on load
    initCounter();
  </script>`;
}
```
