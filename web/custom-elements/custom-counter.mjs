function CustomCounter({ html, state }) {
  return html`
    <div class="flex gap-3 items-center h-10">
      <div class="counter-display" role="status" aria-live="polite">
        Count: <span id="count-value" class="w-10"></span>
      </div>
      <div>
        <button
          id="increment"
          aria-label="Increment counter"
          title="Increment counter"
        >
          <span aria-hidden="true">+</span>
        </button>
      </div>
    </div>

    <script type="module">
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
      document
        .getElementById("increment")
        .addEventListener("click", async () => {
          const currentCount = await getCount();
          const newCount = currentCount + 1;
          await setCount(newCount);
          document.getElementById("count-value").textContent = newCount;
        });

      // Initialize on load
      initCounter();
    </script>
  `;
}
